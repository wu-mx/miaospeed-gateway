package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/miaokobot/miaospeed/interfaces"
	"miaospeed-gateway/config"
	"miaospeed-gateway/log"
	"miaospeed-gateway/utils"
	"net"
	"net/http"
	"os"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Socket establishing error: %s", err)
		return
	}
	defer conn.Close()

	for {
		sr := interfaces.SlaveRequest{}
		err = conn.ReadJSON(&sr)
		if err != nil {
			if !strings.Contains(err.Error(), "EOF") && !strings.Contains(err.Error(), "reset by peer") {
				log.Errorf("Receiving error: %s", err.Error())
			}

			return
		}

		slave, ok := config.GConf.Slaves[sr.Basics.Slave]
		if !ok {
			conn.WriteJSON(interfaces.SlaveResponse{
				Error: "Gateway: Slave not found.",
			})
			return
		}
		if slave.Disable {
			conn.WriteJSON(interfaces.SlaveResponse{
				Error: "Gateway: Slave disabled.",
			})
			return
		}

		if !slave.SkipTokenVerify {
			if !utils.VerifyRequest(&sr, config.GConf.Slaves[sr.Basics.Slave].Token, config.OFFICIAL_BUILD_TOKEN) {
				log.Debugf("Challenge failed: %s", sr.Basics.Slave)
				conn.WriteJSON(interfaces.SlaveResponse{
					Error: "Gateway: Could not verify the request,please check your token.",
				})
				return
			}
		}

		if len(config.GConf.Whitelist) > 0 {
			if !utils.HasStringVal(config.GConf.Whitelist, sr.Basics.Invoker) {
				log.Debugf("Invoker %s not in whitelist", sr.Basics.Invoker)
				conn.WriteJSON(interfaces.SlaveResponse{
					Error: "Gateway: You are not in whitelist.",
				})
				return
			}
		}

		client, err := NewClient(slave)
		if err != nil {
			log.Errorf("Error when creating client: %s", err.Error())
			conn.WriteJSON(interfaces.SlaveResponse{
				Error: fmt.Sprintf("Gateway: Error when creating client: %s", err.Error()),
			})
			return
		}

		if slave.Invoker != "" {
			sr.Basics.Invoker = slave.Invoker
		}

		utils.Resign(&sr, slave.Token, slave.BuildToken)

		client.conn.WriteJSON(sr)
		log.Info(fmt.Sprintf("Forwarded task %s from %s to %s(%s)", sr.Basics.ID, sr.Basics.Invoker, sr.Basics.Slave, sr.Basics.SlaveName))

		for {
			var sr interfaces.SlaveResponse
			err = client.conn.ReadJSON(&sr)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					return
				}
				log.Errorf("Error when reading from remote slave: %s", err.Error())
				conn.WriteJSON(interfaces.SlaveResponse{
					Error: fmt.Sprintf("Gateway: Error when reading from remote slave: %s", err.Error()),
				})
				return
			}
			if sr.Error != "" {
				sr.Error = "Remote Slave: " + sr.Error
			}
			conn.WriteJSON(sr)
		}
	}
}

func LaunchServer() {
	server := http.Server{
		Handler:   http.HandlerFunc(handler),
		TLSConfig: config.MakeSelfSignedTLSServer(),
	}

	if strings.HasPrefix(config.GConf.Listen, "/") {
		unixListener, err := net.Listen("unix", config.GConf.Listen)
		if err != nil {
			log.Errorf("Server listen faild: %s", err.Error())
			os.Exit(1)
		}
		server.Serve(unixListener)
	} else {
		netListener, err := net.Listen("tcp", config.GConf.Listen)
		if err != nil {
			log.Errorf("Server listen faild: %s", err.Error())
			os.Exit(1)
		}

		log.Successf("Server listen on %s", config.GConf.Listen)
		if config.GConf.TLS {
			server.ServeTLS(netListener, "", "")
		} else {
			server.Serve(netListener)
		}
	}
}
