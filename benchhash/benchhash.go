package benchhash

import (
	"log"
	"math/rand"
	"sync"
	"testing"
)

func init() {
	log.Printf("Nothing")
}

type HashMap interface {
	Get(key uint32) (uint32, bool)
	Put(key, val uint32)
}

type KeyValPair struct {
	Key, Val uint32
}

func intArray(n int) []uint32 {
	a := make([]uint32, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Uint32()
	}
	return a
}

func intPairArray(n int) []KeyValPair {
	a := make([]KeyValPair, n)
	for i := 0; i < n; i++ {
		a[i] = KeyValPair{rand.Uint32(), rand.Uint32()}
	}
	return a
}

//////////////////////////////////////////////////////////////////////////////////
// ReadTest(n)
// Read n times from an empty hash map.
//////////////////////////////////////////////////////////////////////////////////

// SingleRead does single-threaded read.
func SingleRead(n int, newFunc func() HashMap, b *testing.B) {
	work := intArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h := newFunc()
		for j := 0; j < n; j++ {
			h.Get(work[j])
		}
	}
}

// MultiReadChan gets n items using q Go routines. Use a channel to pass work.
func MultiReadChan(n, q int, newFunc func() HashMap, b *testing.B) {
	work := intArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ { // N reps.
		h := newFunc()
		c := make(chan uint32)
		var wg sync.WaitGroup
		for j := 0; j < q; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for key := range c { // Pull work from channel.
					h.Get(key)
				}
			}()
		}
		for _, key := range work {
			c <- key
		}
		close(c)
		wg.Wait()
	}
}

// MultiReadNoChan gets n items using q Go routines. Use a channel to pass work.
func MultiReadNoChan(n, q int, newFunc func() HashMap, b *testing.B) {
	work := intArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ { // N reps.
		var ptr int             // Next input is work[ptr].
		var ptrMutex sync.Mutex // Guards the variable "ptr".
		h := newFunc()
		var wg sync.WaitGroup
		for j := 0; j < q; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					ptrMutex.Lock()
					p := ptr
					ptr++
					ptrMutex.Unlock()
					if p >= n {
						break
					}
					h.Get(work[p])
				}
			}()
		}
		wg.Wait()
	}
}

//////////////////////////////////////////////////////////////////////////////////
// WriteTest(n)
// Write n times to an empty hash map.
//////////////////////////////////////////////////////////////////////////////////

// SingleWrite does single-threaded write.
func SingleWrite(n int, newFunc func() HashMap, b *testing.B) {
	work := intPairArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		h := newFunc()
		for _, kv := range work {
			h.Put(kv.Key, kv.Val)
		}
	}
}

// MultiWriteChan writes n items using q Go routines. Use a channel to pass work.
func MultiWriteChan(n, q int, newFunc func() HashMap, b *testing.B) {
	work := intPairArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ { // N reps.
		h := newFunc()
		c := make(chan KeyValPair)
		var wg sync.WaitGroup
		for j := 0; j < q; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for kv := range c { // Pull work from channel.
					h.Put(kv.Key, kv.Val)
				}
			}()
		}
		for _, kv := range work { // Push work into channel.
			c <- kv
		}
		close(c)
		wg.Wait()
	}
}

// MultiWriteNoChan gets n items using q Go routines. Use a channel to pass work.
func MultiWriteNoChan(n, q int, newFunc func() HashMap, b *testing.B) {
	work := intPairArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ { // N reps.
		var ptr int             // Next input is work[ptr].
		var ptrMutex sync.Mutex // Guards the variable "ptr".
		h := newFunc()
		var wg sync.WaitGroup
		for j := 0; j < q; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					ptrMutex.Lock()
					p := ptr
					ptr++
					ptrMutex.Unlock()
					if p >= n {
						break
					}
					h.Put(work[p].Key, work[p].Val)
				}
			}()
		}
		wg.Wait()
	}
}

//////////////////////////////////////////////////////////////////////////////////
// WriteRead(n)
// Reads and writes happen at the same time. Do n/3 writes, 2n/3 reads.
//////////////////////////////////////////////////////////////////////////////////

// ReadWriteChan does read and write concurrently using channels. Use only one
// goroutines for writes. Use q goroutines for reads.
func ReadWriteChan(n, qRead, qWrite int, newFunc func() HashMap, b *testing.B) {
	workWrite := intPairArray(n / 3)
	workRead := intArray(n - len(workWrite))
	b.StartTimer()
	for i := 0; i < b.N; i++ { // N reps.
		h := newFunc()
		cRead := make(chan uint32)
		cWrite := make(chan KeyValPair)
		var wg sync.WaitGroup
		for j := 0; j < qRead; j++ { // Read goroutines.
			wg.Add(1)
			go func() {
				defer wg.Done()
				for key := range cRead { // Pull work from read channel.
					h.Get(key)
				}
			}()
		}

		for j := 0; j < qWrite; j++ { // Write goroutines.
			wg.Add(1)
			go func() {
				defer wg.Done()
				for kv := range cWrite { // Pull work from write channel.
					h.Put(kv.Key, kv.Val)
				}
			}()
		}
		for _, key := range workRead {
			cRead <- key
		}
		for _, kv := range workWrite {
			cWrite <- kv
		}

		close(cRead)
		close(cWrite)
		wg.Wait()
	}
}
