package cmd

import (
	"fmt"
	"os"

	"github.com/YOwatari/hhmm/params"
	"github.com/spf13/cobra"
)

var number int64

var n params.DevideNumber
var hhmm params.HHMM

var RootCmd = &cobra.Command{
	Use:   "hhmm",
	Short: "hhmm is a CLI",
	Long:  "hhmm is a CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := n.Set(number); err != nil {
			fmt.Print(err)
			os.Exit(-1)
		}

		if err := hhmm.Set(n.Value); err != nil {
			fmt.Print(err)
			os.Exit(-1)
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize()
	RootCmd.PersistentFlags().Int64VarP(&number, "devide", "d", 288, "devide number")
}
