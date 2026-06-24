package store

import (
	"crypto/sha256"
	"encoding/hex"
)

type Dep struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (d *Dep) Key() string {
	hasher := sha256.New()
	data := d.Name + d.Version
	hasher.Write([]byte(data))

	return hex.EncodeToString(hasher.Sum(nil))
}
