package main

import (
	"testing"
	"hash/fnv"
	"fmt"
	farm "github.com/dgryski/go-farm"
)

func BenchmarkFnvHash(b *testing.B) {
	h := fnv.New64a()
	for i := 0; i < b.N; i++ {
		h.Write([]byte(key(i)))
		h.Sum64()
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
