package core

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"gopkg.pigeonligh.com/kube-helper/pkg/utils"
)

func GetDefaultConfigPath() string {
	err := os.MkdirAll(path.Join(utils.UserDir(), ".kube-helper"), os.ModePerm)
	if err != nil {
		logrus.WithError(err).Fatalln("Failed to create config directory")
	}
	return path.Join(utils.UserDir(), ".kube-helper", "config")
}

func (c *Context) ReadConfig() *Config {
	config := &Config{}

	data, err := ioutil.ReadFile(c.configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config
		}
		logrus.WithError(err).Fatalln("Failed to open config file.")
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		logrus.WithError(err).Fatalln("Failed to parse config file.")
	}

	return config
}

func (c *Context) SaveConfig(config *Config) {
	config.SwitchConfig.Kubeconfigs = MaintainKubeconfigs(config.SwitchConfig.Kubeconfigs)

	data, err := yaml.Marshal(config)
	if err != nil {
		logrus.WithError(err).Fatalln("Failed to marshal config.")
	}

	err = ioutil.WriteFile(c.configPath, data, os.ModePerm)
	if err != nil {
		logrus.WithError(err).Fatalln("Failed to save config.")
	}
}
