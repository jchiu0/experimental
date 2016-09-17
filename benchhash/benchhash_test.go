package benchhash

import (
	"log"
	"testing"
)

const (
	hashMapSize   = 100000
	numGoRoutines = 1000
)

func init() {
	log.Printf("n=%d q=%d", hashMapSize, numGoRoutines)
}

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

func BenchmarkRead(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			MultiRead(hashMapSize, numGoRoutines, p.newFunc, b)
		})
	}
}

func BenchmarkWrite(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			MultiWrite(hashMapSize, numGoRoutines, p.newFunc, b)
		})
	}
}

func benchmarkReadWrite(b *testing.B, fracRead float64) {
	numReadGoRoutines := int(fracRead * float64(numGoRoutines))
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			ReadWrite(hashMapSize, numReadGoRoutines, numGoRoutines-numReadGoRoutines,
				p.newFunc, b)
		})
	}
}

func BenchmarkReadWrite1(b *testing.B) { benchmarkReadWrite(b, 0.1) }
func BenchmarkReadWrite3(b *testing.B) { benchmarkReadWrite(b, 0.3) }
func BenchmarkReadWrite5(b *testing.B) { benchmarkReadWrite(b, 0.5) }
func BenchmarkReadWrite7(b *testing.B) { benchmarkReadWrite(b, 0.7) }
func BenchmarkReadWrite9(b *testing.B) { benchmarkReadWrite(b, 0.9) }
