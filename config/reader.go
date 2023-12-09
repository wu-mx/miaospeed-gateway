package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"miaospeed-gateway/log"
)

func Load(path string, conf *Config) error {
	cfile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(cfile, conf)
	if err != nil {
		return err
	}

	if conf.Slaves == nil {
		return errors.New("no slaves defined in config file")
	}

	for i, slave := range conf.Slaves {
		if slave.Address == "" {
			return errors.New(fmt.Sprintf("slave address is not defined for slave %s", i))
		}
		if slave.BuildToken == "" {
			slave.BuildToken = OFFICIAL_BUILD_TOKEN
		}

		if slave.TLSPubKey == "" {
			slave.TLSPubKey = OFFICIAL_TLS_PUB_KEY
		} else {
			cert, err := ioutil.ReadFile(slave.TLSPubKey)
			if err != nil {
				return errors.New(fmt.Sprintf("can not read key file %s. %s", slave.TLSPubKey, err.Error()))
			}
			slave.TLSPubKey = string(cert)
		}
		/*if slave.TLSPrivKey == "" {
			slave.TLSPrivKey = OFFICIAL_TLS_PRIV_KEY
		} else {
			cert, err := ioutil.ReadFile(slave.TLSPrivKey)
			if err != nil {
				return errors.New(fmt.Sprintf("can not readkey file %s. %s", slave.TLSPrivKey, err.Error()))
			}
			slave.TLSPrivKey = string(cert)
		}*/
	}
	log.Successf("Loaded %d slaves.", len(conf.Slaves))
	return nil
}
