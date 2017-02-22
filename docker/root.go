package docker

import (
	"context"
	"fmt"
	"io/ioutil"

	// "github.com/docker/docker/api/types"
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

	stats, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return
	}
	defer func() {
		if err = stats.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	b, err := ioutil.ReadAll(stats.Body)
	if err != nil {
		return
	}
	value = string(b)

	return
}
