package commands

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use: "kubectl-switch",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	cmd.AddCommand(listKubeConfigCommand())
	cmd.AddCommand(setKubeConfigCommand())
	cmd.AddCommand(unsetKubeConfigCommand())
	cmd.AddCommand(useKubeConfigCommand())
	cmd.AddCommand(currentKubeConfigCommand())

	cmd.InitDefaultHelpFlag()
	_ = cmd.Flags().MarkHidden("help")
	return cmd
}
