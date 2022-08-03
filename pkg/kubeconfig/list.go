package kubeconfig

import (
	"fmt"
	"os"

	"gopkg.pigeonligh.com/easygo/pretty/table"
	"gopkg.pigeonligh.com/easygo/pretty/text"
	"gopkg.pigeonligh.com/kube-helper/pkg/core"
)

func PrintList(c *core.Context, suggest bool) {
	config := c.ReadConfig()
	if suggest {
		for _, conf := range config.SwitchConfig.Kubeconfigs {
			fmt.Fprintln(c.Writer(), conf.Name)
		}
		fmt.Fprintln(c.Writer(), "none")
		return
	}

	if len(config.SwitchConfig.Kubeconfigs) == 0 {
		fmt.Fprintln(c.Writer(), "No configs")
	} else {
		fmt.Fprintln(c.Writer(), "Configs:")

		t := table.NewByHeaders([]table.TableHeader{
			{Text: "#", HeaderAlign: text.AlignCenter, TextAlign: text.AlignCenter},
			{Text: "name", HeaderAlign: text.AlignCenter, TextAlign: text.AlignLeft},
			{Text: "filepath", HeaderAlign: text.AlignCenter, TextAlign: text.AlignLeft},
		})

		nowKubeconfigPath := os.Getenv("KUBECONFIG")
		for i, conf := range config.SwitchConfig.Kubeconfigs {
			if conf.Path == nowKubeconfigPath {
				t.AddRow(table.Row{"->", conf.Name, conf.Path})
			} else {
				t.AddRow(table.Row{i + 1, conf.Name, conf.Path})
			}
		}
		fmt.Fprint(c.Writer(), t.Render())
	}
}
