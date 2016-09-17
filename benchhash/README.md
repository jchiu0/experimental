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

For the `ReadWrite` setup, we ask the user for an addition parameter `fracRead`. There will be `fracRead * q` goroutines which do reads and `(1-fracRead) * q` goroutines which do writes. To be clear, each goroutine scans `n/q` elements.

We did not use `RunParallel` because it doesn't fit the structure of our test setup.

# Results

```
2016/09/17 09:20:03 n=100000 q=10
testing: warning: no tests to run
BenchmarkRead/GoMap-8   	     300	   4777112 ns/op
BenchmarkRead/GotomicMap-8         	     300	   4999495 ns/op
BenchmarkRead/ShardedGoMap4-8      	     300	   4023751 ns/op
BenchmarkRead/ShardedGoMap8-8      	     500	   3533274 ns/op
BenchmarkWrite/GoMap-8             	      50	  21076068 ns/op
BenchmarkWrite/GotomicMap-8        	      20	  61367851 ns/op
BenchmarkWrite/ShardedGoMap4-8     	     100	  19192543 ns/op
BenchmarkWrite/ShardedGoMap8-8     	     100	  13161005 ns/op
BenchmarkReadWrite1/GoMap-8        	      50	  26428365 ns/op
BenchmarkReadWrite1/GotomicMap-8   	      20	  54971271 ns/op
BenchmarkReadWrite1/ShardedGoMap4-8         	      50	  20199121 ns/op
BenchmarkReadWrite1/ShardedGoMap8-8         	     100	  14794162 ns/op
BenchmarkReadWrite3/GoMap-8                 	      50	  24263125 ns/op
BenchmarkReadWrite3/GotomicMap-8            	      30	  43635552 ns/op
BenchmarkReadWrite3/ShardedGoMap4-8         	     100	  19724739 ns/op
BenchmarkReadWrite3/ShardedGoMap8-8         	     100	  15082278 ns/op
BenchmarkReadWrite5/GoMap-8                 	      50	  22199468 ns/op
BenchmarkReadWrite5/GotomicMap-8            	      50	  31865706 ns/op
BenchmarkReadWrite5/ShardedGoMap4-8         	     100	  17662395 ns/op
BenchmarkReadWrite5/ShardedGoMap8-8         	     100	  14077359 ns/op
BenchmarkReadWrite7/GoMap-8                 	      50	  21696984 ns/op
BenchmarkReadWrite7/GotomicMap-8            	      50	  21396830 ns/op
BenchmarkReadWrite7/ShardedGoMap4-8         	     100	  22103478 ns/op
BenchmarkReadWrite7/ShardedGoMap8-8         	     100	  13461778 ns/op
BenchmarkReadWrite9/GoMap-8                 	     100	  17888208 ns/op
BenchmarkReadWrite9/GotomicMap-8            	     100	  12637082 ns/op
BenchmarkReadWrite9/ShardedGoMap4-8         	     100	  14613906 ns/op
BenchmarkReadWrite9/ShardedGoMap8-8         	     100	  12212447 ns/op
```

```
2016/09/17 09:21:37 n=100000 q=100
testing: warning: no tests to run
BenchmarkRead/GoMap-8   	     300	   4615845 ns/op
BenchmarkRead/GotomicMap-8         	     300	   4791639 ns/op
BenchmarkRead/ShardedGoMap4-8      	     500	   3552482 ns/op
BenchmarkRead/ShardedGoMap8-8      	     500	   3030253 ns/op
BenchmarkWrite/GoMap-8             	      50	  33950352 ns/op
BenchmarkWrite/GotomicMap-8        	      20	  62495558 ns/op
BenchmarkWrite/ShardedGoMap4-8     	     100	  19541584 ns/op
BenchmarkWrite/ShardedGoMap8-8     	     100	  12773103 ns/op
BenchmarkReadWrite1/GoMap-8        	      50	  30339647 ns/op
BenchmarkReadWrite1/GotomicMap-8   	      20	  59281297 ns/op
BenchmarkReadWrite1/ShardedGoMap4-8         	     100	  19697308 ns/op
BenchmarkReadWrite1/ShardedGoMap8-8         	     100	  12942587 ns/op
BenchmarkReadWrite3/GoMap-8                 	      50	  31352920 ns/op
BenchmarkReadWrite3/GotomicMap-8            	      30	  47015773 ns/op
BenchmarkReadWrite3/ShardedGoMap4-8         	     100	  17926981 ns/op
BenchmarkReadWrite3/ShardedGoMap8-8         	     100	  12052382 ns/op
BenchmarkReadWrite5/GoMap-8                 	     100	  15386086 ns/op
BenchmarkReadWrite5/GotomicMap-8            	      30	  35690441 ns/op
BenchmarkReadWrite5/ShardedGoMap4-8         	     100	  14828539 ns/op
BenchmarkReadWrite5/ShardedGoMap8-8         	     100	  10471614 ns/op
BenchmarkReadWrite7/GoMap-8                 	     100	  13387133 ns/op
BenchmarkReadWrite7/GotomicMap-8            	      50	  21134461 ns/op
BenchmarkReadWrite7/ShardedGoMap4-8         	     100	  12635970 ns/op
BenchmarkReadWrite7/ShardedGoMap8-8         	     200	   9206044 ns/op
BenchmarkReadWrite9/GoMap-8                 	     100	  10234650 ns/op
BenchmarkReadWrite9/GotomicMap-8            	     100	  11980098 ns/op
BenchmarkReadWrite9/ShardedGoMap4-8         	     200	   9035206 ns/op
BenchmarkReadWrite9/ShardedGoMap8-8         	     200	   6597306 ns/op
```

2016/09/17 09:21:37 n=100000 q=100
testing: warning: no tests to run
BenchmarkRead/GoMap-8   	     300	   4615845 ns/op
BenchmarkRead/GotomicMap-8         	     300	   4791639 ns/op
BenchmarkRead/ShardedGoMap4-8      	     500	   3552482 ns/op
BenchmarkRead/ShardedGoMap8-8      	     500	   3030253 ns/op
BenchmarkWrite/GoMap-8             	      50	  33950352 ns/op
BenchmarkWrite/GotomicMap-8        	      20	  62495558 ns/op
BenchmarkWrite/ShardedGoMap4-8     	     100	  19541584 ns/op
BenchmarkWrite/ShardedGoMap8-8     	     100	  12773103 ns/op
BenchmarkReadWrite1/GoMap-8        	      50	  30339647 ns/op
BenchmarkReadWrite1/GotomicMap-8   	      20	  59281297 ns/op
BenchmarkReadWrite1/ShardedGoMap4-8         	     100	  19697308 ns/op
BenchmarkReadWrite1/ShardedGoMap8-8         	     100	  12942587 ns/op
BenchmarkReadWrite3/GoMap-8                 	      50	  31352920 ns/op
BenchmarkReadWrite3/GotomicMap-8            	      30	  47015773 ns/op
BenchmarkReadWrite3/ShardedGoMap4-8         	     100	  17926981 ns/op
BenchmarkReadWrite3/ShardedGoMap8-8         	     100	  12052382 ns/op
BenchmarkReadWrite5/GoMap-8                 	     100	  15386086 ns/op
BenchmarkReadWrite5/GotomicMap-8            	      30	  35690441 ns/op
BenchmarkReadWrite5/ShardedGoMap4-8         	     100	  14828539 ns/op
BenchmarkReadWrite5/ShardedGoMap8-8         	     100	  10471614 ns/op
BenchmarkReadWrite7/GoMap-8                 	     100	  13387133 ns/op
BenchmarkReadWrite7/GotomicMap-8            	      50	  21134461 ns/op
BenchmarkReadWrite7/ShardedGoMap4-8         	     100	  12635970 ns/op
BenchmarkReadWrite7/ShardedGoMap8-8         	     200	   9206044 ns/op
BenchmarkReadWrite9/GoMap-8                 	     100	  10234650 ns/op
BenchmarkReadWrite9/GotomicMap-8            	     100	  11980098 ns/op
BenchmarkReadWrite9/ShardedGoMap4-8         	     200	   9035206 ns/op
BenchmarkReadWrite9/ShardedGoMap8-8         	     200	   6597306 ns/op
```

```
2016/09/17 09:22:54 n=100000 q=1000
testing: warning: no tests to run
BenchmarkRead/GoMap-8   	     300	   4631208 ns/op
BenchmarkRead/GotomicMap-8         	     200	   7331226 ns/op
BenchmarkRead/ShardedGoMap4-8      	     500	   3557570 ns/op
BenchmarkRead/ShardedGoMap8-8      	     500	   3032456 ns/op
BenchmarkWrite/GoMap-8             	     100	  18611536 ns/op
BenchmarkWrite/GotomicMap-8        	      20	  69230224 ns/op
BenchmarkWrite/ShardedGoMap4-8     	      50	  20495975 ns/op
BenchmarkWrite/ShardedGoMap8-8     	     100	  13947165 ns/op
BenchmarkReadWrite1/GoMap-8        	     100	  18366579 ns/op
BenchmarkReadWrite1/GotomicMap-8   	      20	  61766970 ns/op
BenchmarkReadWrite1/ShardedGoMap4-8         	     100	  19582341 ns/op
BenchmarkReadWrite1/ShardedGoMap8-8         	     100	  13659815 ns/op
BenchmarkReadWrite3/GoMap-8                 	     100	  16403908 ns/op
BenchmarkReadWrite3/GotomicMap-8            	      20	  55717562 ns/op
BenchmarkReadWrite3/ShardedGoMap4-8         	     100	  17180934 ns/op
BenchmarkReadWrite3/ShardedGoMap8-8         	     100	  12518772 ns/op
BenchmarkReadWrite5/GoMap-8                 	     100	  13356154 ns/op
BenchmarkReadWrite5/GotomicMap-8            	      30	  42300278 ns/op
BenchmarkReadWrite5/ShardedGoMap4-8         	     100	  14127336 ns/op
BenchmarkReadWrite5/ShardedGoMap8-8         	     100	  11109806 ns/op
BenchmarkReadWrite7/GoMap-8                 	     100	  12033057 ns/op
BenchmarkReadWrite7/GotomicMap-8            	      50	  27611885 ns/op
BenchmarkReadWrite7/ShardedGoMap4-8         	     100	  11935605 ns/op
BenchmarkReadWrite7/ShardedGoMap8-8         	     200	   9450262 ns/op
BenchmarkReadWrite9/GoMap-8                 	     200	   9522226 ns/op
BenchmarkReadWrite9/GotomicMap-8            	     100	  17649796 ns/op
BenchmarkReadWrite9/ShardedGoMap4-8         	     200	   9463748 ns/op
BenchmarkReadWrite9/ShardedGoMap8-8         	     200	   6575452 ns/op
```