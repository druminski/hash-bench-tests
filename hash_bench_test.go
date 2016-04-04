package main

import (
	"testing"
	"hash/crc64"
	"hash/fnv"
	"fmt"
	farm "github.com/dgryski/go-farm"
)

func BenchmarkFnvaHash(b *testing.B) {
	h := fnv.New64a()
	for i := 0; i < b.N; i++ {
		h.Write([]byte(key(i)))
		h.Sum64()
	}
}

func BenchmarkFnvHash(b *testing.B) {
	h := fnv.New64()
	for i := 0; i < b.N; i++ {
		h.Write([]byte(key(i)))
		h.Sum64()
	}
}

func BenchmarkCRCHash(b *testing.B) {
	table := crc64.MakeTable(crc64.ISO)
	for i := 0; i < b.N; i++ {
		crc64.Checksum([]byte(key(i)), table)
	}
}

func BenchmarkFarmHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		farm.Hash64([]byte(key(i)))
	}
}

func key(i int) string {
	return fmt.Sprintf("key-%010d", i)
}
