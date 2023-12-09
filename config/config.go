package config

var GConf Config
var Build string
var Version = "0.0.2"
var Commit string

type Slave struct {
	Address    string `yaml:"address"`
	Token      string `yaml:"token"`
	BuildToken string `yaml:"buildToken"` //Default: OFFICIAL_BUILD_TOKEN
	TLSPubKey  string `yaml:"tlsPubKey"`  //Default: OFFICIAL_TLS_PUB_KEY
	//TLSPrivKey      string `yaml:"tlsPrivKey"` //Default: OFFICIAL_TLS_PRIV_KEY
	Disable         bool   `yaml:"disable"`
	SkipTokenVerify bool   `yaml:"skipVerify"`
	SkipTLSVerify   bool   `yaml:"skipTLSVerify"`
	Invoker         string `yaml:"invoker"`
}
type Config struct {
	Slaves    map[string]*Slave `yaml:"slaves"`
	TLS       bool              `yaml:"serverTLS"`
	Listen    string            `yaml:"listen"`
	Whitelist []string          `yaml:"whitelist"`
}
