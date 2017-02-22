package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var vmCommand = &cobra.Command{
	Use:   "vm",
	Short: "monitor virtual machine",
	RunE:  monitorVM,
}

func monitorVM(cmd *cobra.Command, args []string) error {
	fmt.Printf("args: %v.", args)
	return nil
}
