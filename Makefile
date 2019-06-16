.phony: build hyperfine gtime

build: 
	# @cd cpu-insertionsort/go && CGO_ENABLED=1 go build -o main main.go && cd -
	@cd cpu-mergesort/go && CGO_ENABLED=1 go build -o main main.go && cd -
	@cd cpu-mergesort/c && gcc main.c -o main && cd -
	@cd cpus-mergesort/go && CGO_ENABLED=1 go build -o main main.go && cd -

hyperfine:
	# @hyperfine --warmup 3 './cpu-insertionsort/go/main'
	@hyperfine --warmup 3 './cpu-mergesort/go/main'
	@hyperfine --warmup 3 './cpu-mergesort/c/main'	
	@hyperfine --warmup 3 './cpu-quicksort/go/main'
	@hyperfine --warmup 3 './cpus-mergesort/go/main'