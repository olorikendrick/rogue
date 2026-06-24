package store

import (
	"runtime"
)

type Host struct {
	Architecture string `json:"architecture"`
	WordSize     byte   `json:"WordSize"`
}

func GetHost() Host {
	arch := runtime.GOARCH

	word := 64

	if ^uintptr(0) == uintptr(^uint32(0)) {
		word = 32
	}

	return Host{
		Architecture: arch,
		WordSize:     byte(word),
	}
}
