package lain

import (
	"errors"
	"strings"
)

var (
	errInvalidContainerName = errors.New("invalid container name")
)

// ParseAppName 从容器名解析得到 LAIN 的 App Name
func ParseAppName(containerName string) (string, error) {
	if containerName == "" {
		return "", errInvalidContainerName
	}

	parts := strings.Split(containerName, ".")
	if len(parts) == 0 {
		return "", errInvalidContainerName
	}

	return parts[0], nil
}
