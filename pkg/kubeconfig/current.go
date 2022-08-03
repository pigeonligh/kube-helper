package kubeconfig

import (
	"fmt"

	"gopkg.pigeonligh.com/kube-helper/pkg/core"
	"gopkg.pigeonligh.com/kube-helper/pkg/utils"
)

func wrapKubeconfig(s string) string {
	return ` | %{[32m%}kube:%{[32m%}` + s + `%{[32m%}%{[00m%}`
}

func PrintCurrent(c *core.Context, format bool) {
	config := c.ReadConfig()
	if len(config.SwitchConfig.Kubeconfigs) > 0 {
		nowKubeconfigPath := utils.KubeConfig()
		for _, conf := range config.SwitchConfig.Kubeconfigs {
			if conf.Path == nowKubeconfigPath {
				if format {
					fmt.Fprintln(c.Writer(), wrapKubeconfig(conf.Name))
				} else {
					fmt.Fprintln(c.Writer(), conf.Name)
				}
				return
			}
		}
	}
	fmt.Fprintln(c.Writer())
}
