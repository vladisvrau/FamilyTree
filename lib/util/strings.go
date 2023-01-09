package util

import (
	"fmt"
	"strings"
)

func JoinSlice(sep string, list []any) string {
	var stringSlice []string

	for _, item := range list {
		stringSlice = append(stringSlice, fmt.Sprintf("%s", item))
	}

	return strings.Join(stringSlice, sep)
}
