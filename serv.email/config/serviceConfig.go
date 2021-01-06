package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/golang/glog"
	"go.kicksware.com/api/service-common/config"
)

type ServiceConfig struct {
	Common       config.CommonConfig     `yaml:"commonConfig"`
	Events       config.ConnectionConfig `yaml:"eventsConfig"`
	RPC          config.ConnectionConfig `yaml:"rpcConfig"`
	Mail         MailConfig              `yaml:"mailConfig"`
	FallbackMail MailConfig              `yaml:"fallbackMailConfig"`
}

type MailConfig struct {
	Server                string `yaml:"server"`
	Address               string `yaml:"address"`
	Password              string `yaml:"password"`
	VerifyEmailTemplate   string `yaml:"verifyEmailTemplate"`
	ResetPasswordTemplate string `yaml:"resetPasswordTemplate"`
	NotificationTemplate  string `yaml:"notificationTemplate"`
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
