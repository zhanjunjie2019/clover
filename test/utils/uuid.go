package utils

import (
	"encoding/binary"
	"github.com/google/uuid"
	"strings"
)

func UUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func TinyUUID() (rs string) {
	bs, _ := uuid.New().MarshalBinary()
	u641 := binary.LittleEndian.Uint64(bs[:8])
	u642 := binary.LittleEndian.Uint64(bs[8:])
	s1 := toString(u641)
	s2 := toString(u642)
	return s1 + s2
}

func toString(i uint64) string {
	if i < cl {
		return cs[i]
	} else {
		return toString(i/cl) + cs[i%cl]
	}
}

const cl = 63

var cs = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o",
	"p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x",
	"c", "v", "b", "n", "m", "1", "2", "3", "4", "5", "6", "7",
	"8", "9", "0", "_", "Q", "W", "E", "R", "T", "Y", "U", "I",
	"O", "P", "A", "S", "D", "F", "G", "H", "J", "K", "L", "Z",
	"X", "C", "V", "B", "N", "M"}
