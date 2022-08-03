package kubeconfig

import (
	"fmt"

	"gopkg.pigeonligh.com/kube-helper/pkg/core"
	"gopkg.pigeonligh.com/kube-helper/pkg/utils"
)

func Use(c *core.Context, name string) error {
	if name == "none" {
		utils.PrintSource("unset KUBECONFIG")
		return nil
	}

	config := c.ReadConfig()
	var currentConfig core.KubeConfig
	for _, conf := range config.SwitchConfig.Kubeconfigs {
		if conf.Name == name {
			currentConfig = conf
		}
	}
	if currentConfig.Name != name {
		return fmt.Errorf("config %s not found", name)
	}
	utils.PrintSource("export KUBECONFIG=" + currentConfig.Path)
	return nil
}
