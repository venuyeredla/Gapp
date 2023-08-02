package consts

import (
	"Gapp/dsa/ds/slice"
)

type Flag uint16

const BLOCKSIZE = 4096

const (
	INTERNAL Flag = 1 << iota
	LEAF
	VARCHAR_CTRL
	VARCHAR_FREE
	VARCHAR_RUN
	VARCHAR_KEYS
	VARCHAR_VALS
	LIST_CTRL
	LIST_IDX
)

func AsFlag(bytes []byte) Flag {
	back := slice.AsSlice(&bytes)
	return *(*Flag)(back.Array)
}
