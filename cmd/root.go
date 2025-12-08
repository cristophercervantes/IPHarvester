package cmd

import (
	"fmt"
	"os"

	"github.com/cristophercervantes/IPHarvester/banner"
	"github.com/spf13/cobra"
)

var silent bool

var rootCmd = &cobra.Command{
	Use:   "ipharvester",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {

		if v, _ := cmd.Flags().GetBool("version"); v {
			banner.PrintVersion()
			return
		}
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("version", "v", false, "Print the version of the tool and exit.")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Suppress banner output")

	cobra.OnInitialize(func() {
		if !silent {
			banner.PrintBanner()
		}
	})
}
