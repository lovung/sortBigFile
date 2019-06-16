# I. Environment

ENV: Macbook Pro 2014 15 inch

- MacOS: 10.13.6 (17G7024)
- CPU: 2.8 GHz Intel Core i7
- RAM: 16 GB 1600 MHz DDR3
- GPU: NVIDIA GeForce GT 750M 2048 MB - Intel Iris Pro 1536 MB

# II. Tool 1: GNU Time (gtime -v ...)
## Resutls

| Program | Langauge | Data | Items | CPU (%) | RAM (KBs) | Time (s) |
|---------|----------|------|-------|-----|-----|------|
| gpu-bitonicsort | C & Cuda | bit_list.txt | 1048576 | 95 | 20784 | 0.72 |
| gpu-mergesort | C & Cuda | list.txt | 1000000| 98 | 24368 | 2.36 |
| cpu-mergesort | C | list.txt | 1000000 | 98 | 14176 | 0.58 |
| cpu-mergesort | Go | list.txt | 1000000 | 114 | 51172 | 0.34 |
| cpu-quicksort | C |list.txt | 1000000 | 99 | 4776 | 0.33 | 
| cpu-quicksort | Go |list.txt | 1000000 | 108 | 35628 | 0.31 | 

## Details
### gpu-bitonicsort

```
    Command being timed: "./gpu-bitonicsort/main resources/bit_list.txt out/out.txt"
    User time (seconds): 0.63
    System time (seconds): 0.05
    Percent of CPU this job got: 95%
    Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.72
    Average shared text size (kbytes): 0
    Average unshared data size (kbytes): 0
    Average stack size (kbytes): 0
    Average total size (kbytes): 0
    Maximum resident set size (kbytes): 20784
    Average resident set size (kbytes): 0
    Major (requiring I/O) page faults: 166
    Minor (reclaiming a frame) page faults: 5253
    Voluntary context switches: 2
    Involuntary context switches: 171
    Swaps: 0
    File system inputs: 0
    File system outputs: 0
    Socket messages sent: 0
    Socket messages received: 0
    Signals delivered: 0
    Page size (bytes): 4096
    Exit status: 0
```
### gpu-mergesort
```
    Command being timed: "./gpu-mergesort/mergesort -i resources/list.txt -o out/out.txt"
    User time (seconds): 2.27
    System time (seconds): 0.05
    Percent of CPU this job got: 98%
    Elapsed (wall clock) time (h:mm:ss or m:ss): 0:02.36
    Average shared text size (kbytes): 0
    Average unshared data size (kbytes): 0
    Average stack size (kbytes): 0
    Average total size (kbytes): 0
    Maximum resident set size (kbytes): 24368
    Average resident set size (kbytes): 0
    Major (requiring I/O) page faults: 170
    Minor (reclaiming a frame) page faults: 6145
    Voluntary context switches: 0
    Involuntary context switches: 661
    Swaps: 0
    File system inputs: 0
    File system outputs: 0
    Socket messages sent: 0
    Socket messages received: 0
    Signals delivered: 0
    Page size (bytes): 4096
    Exit status: 0
```

### cpu-mergesort/c
```
    Command being timed: "./cpu-mergesort/main"
    User time (seconds): 0.55
    System time (seconds): 0.02
    Percent of CPU this job got: 98%
    Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.58
    Average shared text size (kbytes): 0
    Average unshared data size (kbytes): 0
    Average stack size (kbytes): 0
    Average total size (kbytes): 0
    Maximum resident set size (kbytes): 14176
    Average resident set size (kbytes): 0
    Major (requiring I/O) page faults: 0
    Minor (reclaiming a frame) page faults: 3674
    Voluntary context switches: 0
    Involuntary context switches: 269
    Swaps: 0
    File system inputs: 0
    File system outputs: 0
    Socket messages sent: 0
    Socket messages received: 0
    Signals delivered: 0
    Page size (bytes): 4096
    Exit status: 0
```

### cpu-mergesort/go
```
    Command being timed: "./cpu-mergesort/go/main -i=resources/list.txt"
    User time (seconds): 0.32
    System time (seconds): 0.06
    Percent of CPU this job got: 114%
    Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.34
    Average shared text size (kbytes): 0
    Average unshared data size (kbytes): 0
    Average stack size (kbytes): 0
    Average total size (kbytes): 0
    Maximum resident set size (kbytes): 51172
    Average resident set size (kbytes): 0
    Major (requiring I/O) page faults: 0
    Minor (reclaiming a frame) page faults: 12926
    Voluntary context switches: 0
    Involuntary context switches: 11155
    Swaps: 0
    File system inputs: 0
    File system outputs: 0
    Socket messages sent: 0
    Socket messages received: 0
    Signals delivered: 0
    Page size (bytes): 4096
    Exit status: 0
```


### cpu-quicksort/c
```
    Command being timed: "./cpu-quicksort/c/main"
    User time (seconds): 0.31
    System time (seconds): 0.01
    Percent of CPU this job got: 99%
    Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.33
    Average shared text size (kbytes): 0
    Average unshared data size (kbytes): 0
    Average stack size (kbytes): 0
    Average total size (kbytes): 0
    Maximum resident set size (kbytes): 4776
    Average resident set size (kbytes): 0
    Major (requiring I/O) page faults: 0
    Minor (reclaiming a frame) page faults: 1323
    Voluntary context switches: 0
    Involuntary context switches: 45
    Swaps: 0
    File system inputs: 0
    File system outputs: 0
    Socket messages sent: 0
    Socket messages received: 0
    Signals delivered: 0
    Page size (bytes): 4096
    Exit status: 0
```
### cpu-quicksort/go
```
    Command being timed: "./cpu-quicksort/go/main -i=resources/list.txt"
    User time (seconds): 0.28
    System time (seconds): 0.05
    Percent of CPU this job got: 108%
    Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.31
    Average shared text size (kbytes): 0
    Average unshared data size (kbytes): 0
    Average stack size (kbytes): 0
    Average total size (kbytes): 0
    Maximum resident set size (kbytes): 35628
    Average resident set size (kbytes): 0
    Major (requiring I/O) page faults: 0
    Minor (reclaiming a frame) page faults: 9043
    Voluntary context switches: 0
    Involuntary context switches: 6846
    Swaps: 0
    File system inputs: 0
    File system outputs: 0
    Socket messages sent: 0
    Socket messages received: 0
    Signals delivered: 0
    Page size (bytes): 4096
    Exit status: 0
```

# II. Tool 2: Hyperfine
## Resutls

| Program | Langauge | Data | Items | CPU (%) | RAM (KBs) | Time (s) |
|---------|----------|------|-------|-----|-----|------|
| gpu-bitonicsort | C & Cuda | bit_list.txt | 1048576 | 95 | 20784 | 0.72 |
| gpu-mergesort | C & Cuda | list.txt | 1000000| 98 | 24368 | 2.36 |
| cpu-mergesort | C | list.txt | 1000000 | 98 | 14176 | 0.58 |
| cpu-mergesort | Go | list.txt | 1000000 | 114 | 51172 | 0.34 |
| cpu-quicksort | C |list.txt | 1000000 | 99 | 4776 | 0.33 | 
| cpu-quicksort | Go |list.txt | 1000000 | 108 | 35628 | 0.31 | 

## References

- https://github.com/kevin-albert/cuda-mergesort
- https://github.com/khaman1/GPU-QuickSort-Algorithm
- https://gist.github.com/mre/1392067
- https://www.geeksforgeeks.org/insertion-sort/
- https://www.geeksforgeeks.org/insertion-sort-for-singly-linked-list/
- https://www.researchgate.net/profile/Neetu_Faujdar2/publication/283353990_Performance_Evaluation_of_Merge_and_Quick_Sort_using_GPU_Computing_with_CUDA/links/56b9cefb08ae9d9ac67f3145/Performance-Evaluation-of-Merge-and-Quick-Sort-using-GPU-Computing-with-CUDA.pdf
- https://gist.github.com/mycodeschool/9678029
- https://www.tutorialspoint.com/data_structures_algorithms/merge_sort_program_in_c.htm
- https://hackernoon.com/parallel-merge-sort-in-go-fe14c1bc006
- https://gist.github.com/mcvoid/7079156
- https://github.com/shawnsmithdev/zermelo
- https://austingwalters.com/radix-sort-in-go/
- https://golangbot.com/read-files/
- http://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html
- https://github.com/arisath/Benchmarking-Sorting-Algorithms
- http://www.golangprograms.com/golang-program-for-implementation-of-insertionsort.html