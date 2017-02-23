package docker

import (
	"github.com/docker/docker/api/types"
)

func calculateNetworkRxBytes(networks map[string]types.NetworkStats) (rx float64) {
	for _, v := range networks {
		rx += float64(v.RxBytes)
	}

	return
}

func calculateNetworkTxBytes(networks map[string]types.NetworkStats) (tx float64) {
	for _, v := range networks {
		tx += float64(v.TxBytes)
	}

	return
}
