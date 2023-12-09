package config

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
)

//var RootCAs = make(map[string]*x509.CertPool)

// MakeSelfSignedTLSServer https://github.com/miaokobot/miaospeed/blob/master/preconfigs/certs.go
func MakeSelfSignedTLSServer() *tls.Config {
	cert, _ := tls.X509KeyPair([]byte(OFFICIAL_TLS_PUB_KEY), []byte(OFFICIAL_TLS_PRIV_KEY))

	// Construct a tls.config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// Other options
	}

	return tlsConfig
}

func LoadSlaveCert(slave *Slave) x509.Certificate {
	block, _ := pem.Decode([]byte(slave.TLSPubKey))
	cr, _ := x509.ParseCertificate(block.Bytes)
	return *cr
}
