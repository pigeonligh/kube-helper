package kubeconfig

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.pigeonligh.com/kube-helper/pkg/core"
)

func Unset(c *core.Context, name string) {
	config := c.ReadConfig()
	newConfigs := make([]core.KubeConfig, 0)
	for _, conf := range config.SwitchConfig.Kubeconfigs {
		if conf.Name == name {
			continue
		}
		newConfigs = append(newConfigs, conf)
	}
	config.SwitchConfig.Kubeconfigs = newConfigs
	c.SaveConfig(config)
}

func Set(c *core.Context, name, filepath string, fromStdin bool) {
	if name == "none" {
		logrus.Fatalln("Kubeconfig name cannot be none.")
	}

	if fromStdin {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			logrus.WithError(err).Fatalln("Failed to read kubeconfig from stdin.")
		}
		err = ioutil.WriteFile(filepath, data, os.ModePerm)
		if err != nil {
			logrus.WithError(err).Fatalln("Failed to save kubeconfig to file.")
		}
	}

	config := c.ReadConfig()

	currentIndex := -1
	for index, conf := range config.SwitchConfig.Kubeconfigs {
		if conf.Name == name {
			currentIndex = index
		}
	}
	newConfig := core.KubeConfig{
		Name: name,
		Path: filepath,
	}
	if currentIndex == -1 {
		config.SwitchConfig.Kubeconfigs = append(config.SwitchConfig.Kubeconfigs, newConfig)
	} else {
		config.SwitchConfig.Kubeconfigs[currentIndex] = newConfig
	}

	c.SaveConfig(config)
}
