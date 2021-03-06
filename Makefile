.phony: build hyperfine gtime


build: 
	# @cd cpu-mergesort/c && gcc main.c -o main && cd -

	# @cd cpu-mergesort/go && CGO_ENABLED=1 go build -o main main.go && cd -
	# @cd cpus-mergesort/go && CGO_ENABLED=1 go build -o main main.go && cd -
	@cd cpu-radixsort/go && CGO_ENABLED=1 go build -o main main.go && cd - 

prepare:
	@mkdir -p out && mkdir -p resources
	@go get ./...
	@go run lib/generate/main.go

hyperfine:
	# @hyperfine --warmup 3 './cpu-mergesort/c/main'

	# @hyperfine --warmup 3 './cpu-mergesort/go/main'
	# @hyperfine --warmup 3 './cpu-quicksort/go/main'
	# @hyperfine --warmup 3 './cpus-mergesort/go/main'
	@hyperfine --warmup=3 './cpu-radixsort/go/main'