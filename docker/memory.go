package docker

import (
	"errors"
	"github.com/docker/docker/api/types"
)

var (
	errInvalidMemoryLimit = errors.New("invalid memory limit")
)

func calculateMemPercent(v *types.StatsJSON) (float64, error) {
	if v.MemoryStats.Limit == 0 {
		return 0, errInvalidMemoryLimit
	}

	percent := float64(v.MemoryStats.Usage) / float64(v.MemoryStats.Limit) * 100.0
	return percent, nil
}
