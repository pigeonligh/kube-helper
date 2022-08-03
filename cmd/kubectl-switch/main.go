package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"gopkg.pigeonligh.com/kube-helper/cmd/kubectl-switch/commands"
)

func main() {
	logrus.SetOutput(os.Stderr)

	if err := commands.Command().Execute(); err != nil {
		return
	}
}
