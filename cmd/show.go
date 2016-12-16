package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show hhmm",
	Long:  "show hhmm",
	Run: func(cmd *cobra.Command, args []string) {
		mode := "all"

		if len(args) > 0 {
			mode = args[0]
		}

		switch mode {
		case "after", "a":
			fmt.Printf("%04d\n", hhmm.After)
		case "before", "b":
			fmt.Printf("%04d\n", hhmm.Before)
		default:
			for _, v := range hhmm.All {
				fmt.Printf("%04d\n", v)
			}
		}
	},
}
