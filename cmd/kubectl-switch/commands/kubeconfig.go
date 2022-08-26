package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"gopkg.pigeonligh.com/kube-helper/pkg/core"
	"gopkg.pigeonligh.com/kube-helper/pkg/kubeconfig"
)

func listKubeConfigCommand() *cobra.Command {
	var suggest bool

	cmd := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			context := core.NewDefault()

			kubeconfig.PrintList(context, suggest)
		},
	}
	cmd.Flags().BoolVar(&suggest, "suggest", false, "print for suggest")

	cmd.InitDefaultHelpFlag()
	_ = cmd.Flags().MarkHidden("help")
	return cmd
}

func currentKubeConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "current",
		Run: func(cmd *cobra.Command, args []string) {
			context := core.NewDefault()

			kubeconfig.PrintCurrent(context)
		},
	}
	cmd.InitDefaultHelpFlag()
	_ = cmd.Flags().MarkHidden("help")

	return cmd
}

func setKubeConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "set name filepath",
		Run: func(cmd *cobra.Command, args []string) {
			context := core.NewDefault()

			if len(args) == 1 {
				kubeconfig.Set(context, args[0], kubeconfig.GetDefaultKubeconfigPath(args[0]), true)
			} else if len(args) == 2 {
				kubeconfig.Set(context, args[0], args[1], false)
			} else {
				cmd.HelpFunc()(cmd, args)
			}

		},
	}
	cmd.InitDefaultHelpFlag()
	_ = cmd.Flags().MarkHidden("help")
	return cmd
}

func unsetKubeConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "unset name",
		Run: func(cmd *cobra.Command, args []string) {
			context := core.NewDefault()

			if len(args) == 1 {
				kubeconfig.Unset(context, args[0])
			} else {
				cmd.HelpFunc()(cmd, args)
			}

		},
	}
	cmd.InitDefaultHelpFlag()
	_ = cmd.Flags().MarkHidden("help")
	return cmd
}

func useKubeConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "use name",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.HelpFunc()(cmd, args)
				return
			}
			context := core.NewDefault()

			if err := kubeconfig.Use(context, args[0]); err != nil {
				logrus.Fatalln(err)
			}
		},
	}
	cmd.InitDefaultHelpFlag()
	_ = cmd.Flags().MarkHidden("help")
	return cmd
}
