package docker

import (
	"encoding/json"
	"errors"
	"strconv"

	"gitlab.yxapp.in/kaizhang33/sysmonitor/item"
	"gitlab.yxapp.in/kaizhang33/sysmonitor/lain"
)

const (
	intBase      = 10
	floatFmt     = 'f'
	floatPrec    = 2
	floatBitSize = 64
)

var (
	errInvalidNetworks = errors.New("invalid networks")
)

// Monitor 监控容器信息
func Monitor(containerName string, key item.Key) (value string, err error) {
	switch key {
	case item.CPUTotalUsage:
		v, err := stats(containerName)
		if err != nil {
			return "", err
		}
		value = strconv.FormatUint(v.CPUStats.CPUUsage.TotalUsage, intBase)
	case item.CPUPercent:
		v, err := stats(containerName)
		if err != nil {
			return "", err
		}
		cpuPercent := calculateCPUPercentUnix(
			v.PreCPUStats.CPUUsage.TotalUsage,
			v.PreCPUStats.SystemUsage,
			v,
		)
		value = strconv.FormatFloat(cpuPercent, floatFmt, floatPrec, floatBitSize)
	case item.MemUsage:
		v, err := stats(containerName)
		if err != nil {
			return "", err
		}
		value = strconv.FormatUint(v.MemoryStats.Usage, intBase)
	case item.MemLimit:
		v, err := stats(containerName)
		if err != nil {
			return "", err
		}
		value = strconv.FormatUint(v.MemoryStats.Limit, intBase)
	case item.MemPercent:
		v, err := stats(containerName)
		if err != nil {
			return "", err
		}
		memPercent, err := calculateMemPercent(v)
		if err != nil {
			return "", err
		}
		value = strconv.FormatFloat(memPercent, floatFmt, floatPrec, floatBitSize)
	case item.NetworkRxBytes:
		v, err := stats(containerName)
		if err != nil {
			return "", err
		}
		rxBytes := calculateNetworkRxBytes(v.Networks)
		value = strconv.FormatUint(rxBytes, intBase)
	case item.NetworkTxBytes:
		v, err := stats(containerName)
		if err != nil {
			return "", err
		}
		txBytes := calculateNetworkTxBytes(v.Networks)
		value = strconv.FormatUint(txBytes, intBase)
	case item.NetworkIP:
		appName, err := lain.ParseAppName(containerName)
		if err != nil {
			return "", err
		}

		info, err := inspect(containerName)
		if err != nil {
			return "", err
		}

		network, ok := info.NetworkSettings.Networks[appName]
		if !ok {
			return "", errInvalidNetworks
		}
		value = network.IPAddress
	case item.ConImage:
		info, err := inspect(containerName)
		if err != nil {
			return "", err
		}

		value = info.Config.Image
	case item.ConVolumes:
		info, err := inspect(containerName)
		if err != nil {
			return "", err
		}

		bs, err := json.Marshal(info.Config.Volumes)
		if err != nil {
		}
		value = string(bs)
	default:
		return "", item.ErrInvalidKey
	}

	return value, nil
}
