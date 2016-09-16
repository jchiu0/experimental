package benchhash

import (
	//	"log"
	//	"runtime"
	"testing"
)

const (
	hashMapSize   = 100000
	numGoRoutines = 1000
)

type hashPair struct {
	label   string
	newFunc func() HashMap
}

var (
	hashPairs = []hashPair{
		hashPair{"GoMap", NewGoMap},
		hashPair{"GotomicMap", NewGotomicMap},
		hashPair{"ShardedGoMap4", NewShardedGoMap4},
		hashPair{"ShardedGoMap8", NewShardedGoMap8},
	}
)

func BenchmarkSingleRead(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			SingleRead(hashMapSize, p.newFunc, b)
		})
	}
}

func BenchmarkMultiReadChan(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			MultiReadChan(hashMapSize, numGoRoutines, p.newFunc, b)
		})
	}
}

func BenchmarkMultiReadNoChan(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			MultiReadNoChan(hashMapSize, numGoRoutines, p.newFunc, b)
		})
	}
}

func BenchmarkSingleWrite(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			SingleWrite(hashMapSize, p.newFunc, b)
		})
	}
}

func BenchmarkMultiWriteChan(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			MultiWriteChan(hashMapSize, numGoRoutines, p.newFunc, b)
		})
	}
}

func BenchmarkMultiWriteNoChan(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			MultiWriteNoChan(hashMapSize, numGoRoutines, p.newFunc, b)
		})
	}
}

func BenchmarkReadWriteChan(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			ReadWriteChan(hashMapSize, numGoRoutines, numGoRoutines, p.newFunc, b)
		})
	}
}
