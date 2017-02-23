package docker

import (
	"github.com/docker/docker/api/types"
)

func calculateNetworkRxBytes(networks map[string]types.NetworkStats) (rx uint64) {
	for _, v := range networks {
		rx += v.RxBytes
	}

	return
}

func calculateNetworkTxBytes(networks map[string]types.NetworkStats) (tx uint64) {
	for _, v := range networks {
		tx += v.TxBytes
	}

	return
}
