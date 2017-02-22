package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var (
	errInvalidArgsNumber = errors.New("invalid number of args")
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sysmonitor",
	Short: "Monitor the system",
	Long: `Sysmonitor can monitor the status of docker or virtual machine,
including cpu, memory and network, etc.`,
}

func init() {
	RootCmd.AddCommand(dockerCommand)
	RootCmd.AddCommand(vmCommand)
}
