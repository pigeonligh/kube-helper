package kubeconfig

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.pigeonligh.com/kube-helper/pkg/utils"
)

func GetDefaultKubeconfigPath(name string) string {
	err := os.MkdirAll(path.Join(utils.UserDir(), ".kube-helper"), os.ModePerm)
	if err != nil {
		logrus.WithError(err).Fatalln("Failed to create config directory")
	}

	err = os.MkdirAll(path.Join(utils.UserDir(), ".kube-helper", "kubeconfigs"), os.ModePerm)
	if err != nil {
		logrus.WithError(err).Fatalln("Failed to create config directory")
	}

	return path.Join(utils.UserDir(), ".kube-helper", "kubeconfigs", name)
}
