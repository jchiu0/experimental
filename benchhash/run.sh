set -e

mkdir -p results
for q in 10 100 1000; do
  go test -cpu 2 -benchn 100000 -benchq $q -bench=. > results/benchhash.q$q.txt
done

