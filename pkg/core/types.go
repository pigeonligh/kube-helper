package core

type KubeConfig struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type SwitchConfig struct {
	Kubeconfigs []KubeConfig `yaml:"kubeconfigs"`
}

type Config struct {
	SwitchConfig SwitchConfig `yaml:"switch"`
}
