package main

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/uber-go/zap"
)

const (
	// DockerAPIVersion 表示默认的 DOCKER_API_VERSION
	DockerAPIVersion = "1.24"
)

var (
	log = zap.New(
		zap.NewTextEncoder(zap.TextTimeFormat(time.RFC3339)),
		zap.AddCaller(),
	)
)

func init() {
}

func main() {
	cli, err := client.NewClient(client.DefaultDockerHost, DockerAPIVersion, nil, nil)
	if err != nil {
		log.Fatal("client.NewEnvClient...", zap.Error(err))
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal("cli.ContainerList...", zap.Error(err))
	}

	for _, container := range containers {
		log.Info("container.",
			zap.String("ID", container.ID[:10]),
			zap.String("Image", container.Image),
		)
	}
}
