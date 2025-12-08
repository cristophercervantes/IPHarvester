package main

import (
	"github.com/cristophercervantes/IPHarvester/banner"
	"github.com/cristophercervantes/IPHarvester/cmd"
)

func main() {
	banner.PrintBanner()
	cmd.Execute()
}
