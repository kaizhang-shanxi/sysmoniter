package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.yxapp.in/kaizhang33/sysmonitor/docker"
	"gitlab.yxapp.in/kaizhang33/sysmonitor/item"
)

const (
	dockerUsage = `Usage:
  sysmonitor docker containerName key
key's choices are:
{
    cpu_total_usage,
    cpu_percent,
    mem_usage,
    mem_limit,
    mem_percent,
    network_rx_bytes,
    network_tx_bytes,
    network_ip,
    con_image,
    con_volumes
}
`
)

var dockerCommand = &cobra.Command{
	Use:   "docker",
	Short: "monitor docker container",
	RunE:  monitorDocker,
}

func init() {
	dockerCommand.SetUsageTemplate(dockerUsage)
}

func monitorDocker(cmd *cobra.Command, args []string) error {
	fmt.Printf("args: %v.", args)
	if len(args) != 2 {
		return errInvalidArgsNumber
	}

	key, err := item.Parse(args[1])
	if err != nil {
		return err
	}

	value, err := docker.Monitor(args[0], key)
	if err != nil {
		return err
	}

	fmt.Println(value)
	return nil
}
