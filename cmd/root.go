package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var silent bool

// Public getter so main.go can check it
func IsSilent() bool {
	return silent
}

var rootCmd = &cobra.Command{
	Use:   "ipharvester",
	Short: "Aggressive public search engine IP harvester",
	Long:  "IPHarvester v1.0 — Shodan/ZoomEye/ViewDNS harvester. No API. No mercy.",

	Run: func(cmd *cobra.Command, args []string) {
		if v, _ := cmd.Flags().GetBool("version"); v {
			fmt.Println("v1.0")
			return
		}
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version and exit")
	rootCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "Suppress banner")
}
