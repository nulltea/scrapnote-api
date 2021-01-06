package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/config"
)

type ServiceConfig struct {
	Common config.CommonConfig     `yaml:"commonConfig"`
	Auth   config.AuthConfig       `yaml:"authConfig"`
	Events config.ConnectionConfig `yaml:"eventsConfig"`
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
