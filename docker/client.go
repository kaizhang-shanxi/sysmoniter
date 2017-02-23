package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

const (
	// APIVersion 表示默认的 DOCKER_API_VERSION
	APIVersion = "1.24"
)

func stats(containerID string) (*types.StatsJSON, error) {
	cli, err := client.NewClient(client.DefaultDockerHost, APIVersion, nil, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = cli.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	resp, err := cli.ContainerStats(context.Background(), containerID, false)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var v types.StatsJSON
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}

	return &v, nil
}

func inspect(containerID string) (*types.ContainerJSON, error) {
	cli, err := client.NewClient(client.DefaultDockerHost, APIVersion, nil, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = cli.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	containerInfo, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return nil, err
	}

	return &containerInfo, nil
}
