Bunch of little small explorations.

```
BenchmarkSingleRead/GoMap-8   	     200	   6854720 ns/op
BenchmarkSingleRead/GotomicMap-8         	     100	  16313303 ns/op
BenchmarkSingleRead/ShardedGoMap4-8      	     200	   7216706 ns/op
BenchmarkSingleRead/ShardedGoMap8-8      	     200	   7218582 ns/op
BenchmarkMultiReadChan/GoMap-8           	      30	  37406811 ns/op
BenchmarkMultiReadChan/GotomicMap-8      	      30	  44893949 ns/op
BenchmarkMultiReadChan/ShardedGoMap4-8   	      30	  37423172 ns/op
BenchmarkMultiReadChan/ShardedGoMap8-8   	      30	  37809290 ns/op
BenchmarkMultiReadNoChan/GoMap-8         	     100	  17097579 ns/op
BenchmarkMultiReadNoChan/GotomicMap-8    	      50	  23139782 ns/op
BenchmarkMultiReadNoChan/ShardedGoMap4-8 	     100	  16540913 ns/op
BenchmarkMultiReadNoChan/ShardedGoMap8-8 	     100	  16459923 ns/op
BenchmarkSingleWrite/GoMap-8             	     100	  16681813 ns/op
BenchmarkSingleWrite/GotomicMap-8        	      10	 104578655 ns/op
BenchmarkSingleWrite/ShardedGoMap4-8     	     100	  17553417 ns/op
BenchmarkSingleWrite/ShardedGoMap8-8     	     100	  17643992 ns/op
BenchmarkMultiWriteChan/GoMap-8          	      20	  54098654 ns/op
BenchmarkMultiWriteChan/GotomicMap-8     	      20	  90100367 ns/op
BenchmarkMultiWriteChan/ShardedGoMap4-8  	      30	  43700146 ns/op
BenchmarkMultiWriteChan/ShardedGoMap8-8  	      30	  42394817 ns/op
BenchmarkMultiWriteNoChan/GoMap-8        	      50	  26183198 ns/op
BenchmarkMultiWriteNoChan/GotomicMap-8   	      20	  69900641 ns/op
BenchmarkMultiWriteNoChan/ShardedGoMap4-8         	      50	  22647179 ns/op
BenchmarkMultiWriteNoChan/ShardedGoMap8-8         	      50	  20777076 ns/op
BenchmarkReadWriteChan/GoMap-8                    	      30	  43685500 ns/op
BenchmarkReadWriteChan/GotomicMap-8               	      20	  68394258 ns/op
BenchmarkReadWriteChan/ShardedGoMap4-8            	      30	  40731970 ns/op
BenchmarkReadWriteChan/ShardedGoMap8-8            	      30	  40154742 ns/op
```