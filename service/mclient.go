package service

import (
	"crypto/tls"
	"errors"
	"github.com/gorilla/websocket"
	"miaospeed-gateway/config"
	"strings"
)

type MSClient struct {
	conn *websocket.Conn
}

func NewClient(slave *config.Slave) (*MSClient, error) {
	client := &MSClient{}

	if !(strings.HasPrefix(slave.Address, "mwss://") || strings.HasPrefix(slave.Address, "wss://")) {
		dialer := websocket.DefaultDialer

		conn, _, err := dialer.Dial(slave.Address, nil)
		if err != nil {
			return nil, err
		}

		client.conn = conn
		return client, nil
	} else {
		dialer := &websocket.Dialer{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: strings.HasPrefix(slave.Address, "mwss://") || slave.SkipTLSVerify,
			},
		}

		conn, _, err := dialer.Dial(strings.Replace(slave.Address, "mwss://", "wss://", 1), nil)
		if err != nil {
			return nil, err
		}
		client.conn = conn

		if strings.HasPrefix(slave.Address, "mwss://") {
			scert := config.LoadSlaveCert(slave)
			state := conn.UnderlyingConn().(*tls.Conn).ConnectionState()

			// Verify the server's certificate public key
			yes := scert.Equal(state.PeerCertificates[0])
			if !yes {
				return nil, errors.New("cannot verify the server's security")
			}
		}

		return client, nil
	}
}
