package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/golang/glog"
	"github.com/timoth-y/scrapnote-api/lib.common/config"
)

type ServiceConfig struct {
	Common config.CommonConfig     `yaml:"commonConfig"`
	Events config.ConnectionConfig `yaml:"eventsConfig"`
	RPC    config.ConnectionConfig `yaml:"rpcConfig"`
	Auth   AuthConfig              `yaml:"authConfig"`
}

type AuthConfig struct {
	IssuerName           string `yaml:"issuerName"`
	TokenExpirationDelta int    `yaml:"tokenExpirationDelta"`
	PrivateKeyPath       string `yaml:"privateKeyPath"`
	PublicKeyPath        string `yaml:"publicKeyPath"`
}

func ReadServiceConfig(filename string) (sc ServiceConfig, err error) {
	file, err := ioutil.ReadFile(filename); if err != nil {
		glog.Fatalln(err)
		return
	}

	err = yaml.Unmarshal(file, &sc); if err != nil {
		glog.Fatalln(err)
		return
	}
	return
}
