package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/config"
)

type ServiceConfig struct {
	Common    config.CommonConfig    `yaml:"commonConfig"`
	Cockroach config.DataStoreConfig `yaml:"cockroachConfig"`
	Mongo     config.DataStoreConfig `yaml:"mongoConfig"`
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
