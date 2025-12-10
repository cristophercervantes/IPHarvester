package main

import (
	"github.com/cristophercervantes/IPHarvester/banner"
	"github.com/cristophercervantes/IPHarvester/cmd"
)

func main() {
	// Show banner once, at startup (unless -s/--silent)
	if !cmd.IsSilent() {
		banner.PrintBanner()
	}
	cmd.Execute()
}
