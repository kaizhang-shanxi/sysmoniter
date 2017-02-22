package main

import (
	"fmt"
	"os"

	"gitlab.yxapp.in/kaizhang33/sysmonitor/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
