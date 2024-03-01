go install golang.org/x/perf/cmd/benchstat@latest
go test -bench=. -benchmem -count 10 | tee sorting.txt
benchstat sorting.txt