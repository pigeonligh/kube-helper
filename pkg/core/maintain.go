package core

import "sort"

func MaintainKubeconfigs(configs []KubeConfig) []KubeConfig {
	configMap := make(map[string]KubeConfig)
	names := make([]string, 0, len(configs))

	for _, c := range configs {
		configMap[c.Name] = c
		names = append(names, c.Name)
	}

	sort.Strings(names)

	ret := make([]KubeConfig, 0, len(configs))
	for _, name := range names {
		ret = append(ret, configMap[name])
	}
	return ret
}
