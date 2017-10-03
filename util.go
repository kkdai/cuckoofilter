package cukoofilter

import metro "github.com/dgryski/go-metro"

func hash(data []byte) uint64 {
	return metro.Hash64(data, 1337)
}
