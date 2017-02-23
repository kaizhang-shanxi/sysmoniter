package main

import (
	"os"

	"gitlab.yxapp.in/kaizhang33/sysmonitor/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
