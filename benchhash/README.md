# General setup
There are three different setups `Read`, `Write`, `ReadWrite`. Here are the main parameters.

* `b.N`: This is the total number of reps for go bench to do its measurement.
* `n`: For each rep, we will read/write `n` elements to **one single hash** across multiple goroutines.
* `q`: For each rep, we will use `q` goroutines.

Each of the three setups has the same structure below:

* Assume `n` is divisible by `q`.
* Create an array `work` of size `n`. This is a list of key-value pairs.
* Start timer.
* For each of the `b.N` reps, do the following:
  * Create an empty hash map `h`.
  * Create `q` goroutines.
	* Each goroutine will either read or write to `h`.
	* Each goroutine works on an equal part of the array `work`.

Note that we do not use channels at all as that will cause unnecessary locking.

For the `ReadWrite` setup, we ask the user for an additional parameter `fracRead`. There will be `fracRead * q` goroutines which do reads and `(1-fracRead) * q` goroutines which do writes. To be clear, each goroutine scans `n/q` elements.

We did not use `RunParallel` because it doesn't fit the structure of our test setup.

Please see `benchhash.go` for the test setups.

# Results

Run `run.sh` to run the benchmarks. The results are found in the `results` subdirectory.

Here is the main line in `run.sh`:

```
go test -cpu 2 -benchn 100000 -benchq $q -bench=. > results/benchhash.q$q.txt
```

## Interpreting results

Note that `BenchReadWrite7` means `7/10` of the goroutines are doing reads. Generally, reads are cheaper, so `BenchReadWrite9` should take less time than `BenchReadWrite1`.

Note that `ShardedGoMap16` means we use 16 shards of GoMaps.
