package item

import (
	"errors"
)

type key int

const (
	// CPUTotalUsage 表示 CPU 的总使用量
	CPUTotalUsage key = iota
	// CPUPercent 表示 CPU 使用率
	CPUPercent
	// MemUsage 表示内存的使用量
	MemUsage
	// MemLimit 表示最大可用内存
	MemLimit
	// MemPercent 表示内存的使用率
	MemPercent
	// NetworkRxBytes 表示通过网络接收的字节数
	NetworkRxBytes
	// NetworkTxBytes 表示通过网络发出的字节数
	NetworkTxBytes
	// NetworkIP 表示容器 IP
	NetworkIP
	// ConImage 表示容器镜像
	ConImage
	// ConVolumes 表示容器使用的 Volumes
	ConVolumes
)

var (
	// ErrInvalidKey 表示非法的监控项
	ErrInvalidKey = errors.New("invalid key")
)

// Key 表示监控项
type Key interface {
	get() key
}

func (k key) get() key {
	return k
}

func (k key) String() string {
	switch k {
	case 0:
		return "cpu_total_usage"
	case 1:
		return "cpu_percent"
	case 2:
		return "mem_usage"
	case 3:
		return "mem_limit"
	case 4:
		return "mem_percent"
	case 5:
		return "network_rx_bytes"
	case 6:
		return "network_tx_bytes"
	case 7:
		return "network_ip"
	case 8:
		return "con_image"
	case 9:
		return "con_volumes"
	default:
		return "undefined"
	}
}

// Parse 从字符串里解析出 Key
func Parse(s string) (Key, error) {
	switch s {
	case "cpu_total_usage":
		return CPUTotalUsage, nil
	case "cpu_percent":
		return CPUPercent, nil
	case "mem_usage":
		return MemUsage, nil
	case "mem_limit":
		return MemLimit, nil
	case "mem_percent":
		return MemPercent, nil
	case "network_rx_bytes":
		return NetworkRxBytes, nil
	case "network_tx_bytes":
		return NetworkTxBytes, nil
	case "network_ip":
		return NetworkIP, nil
	case "con_image":
		return ConImage, nil
	case "con_volumes":
		return ConVolumes, nil
	default:
		return nil, ErrInvalidKey
	}
}
