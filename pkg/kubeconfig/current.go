package kubeconfig

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd"

	"gopkg.pigeonligh.com/kube-helper/pkg/core"
	"gopkg.pigeonligh.com/kube-helper/pkg/utils"
)

func parseNamespace(path string) string {
	ns, _, _ := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: path},
		&clientcmd.ConfigOverrides{}).Namespace()
	return ns
}

func joinNameAndNs(name, ns string) string {
	if ns == "" {
		return name
	}
	return name + "/" + ns
}

func PrintCurrent(c *core.Context) {
	config := c.ReadConfig()
	if len(config.SwitchConfig.Kubeconfigs) > 0 {
		nowKubeconfigPath := utils.KubeConfig()
		for _, conf := range config.SwitchConfig.Kubeconfigs {
			if conf.Path == nowKubeconfigPath {
				fmt.Fprintln(c.Writer(), joinNameAndNs(conf.Name, parseNamespace(nowKubeconfigPath)))
				return
			}
		}
	}
	fmt.Fprintln(c.Writer())
}
