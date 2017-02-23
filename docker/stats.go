package docker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"gitlab.yxapp.in/kaizhang33/sysmonitor/item"
)

const (
	// APIVersion 表示默认的 DOCKER_API_VERSION
	APIVersion = "1.24"
)

// Monitor 监控容器信息
func Monitor(containerID string, key item.Key) (value string, err error) {
	cli, err := client.NewClient(client.DefaultDockerHost, APIVersion, nil, nil)
	if err != nil {
		return
	}
	defer func() {
		if err = cli.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	resp, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var stats types.StatsJSON

	if err = json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return
	}

	fmt.Printf("CPUStats: %+v.\n", stats.CPUStats)
	fmt.Printf("MemoryStats: %+v.\n", stats.MemoryStats)
	fmt.Printf("Networks: %+v.\n", stats.Networks)

	return
}
