package main

import (
	"os"

	"github.com/jockjiang/webexbeat/cmd"

	_ "github.com/jockjiang/webexbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
