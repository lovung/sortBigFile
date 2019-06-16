# Performance test

ENV: Macbook Pro 2014 15 inch

- MacOS: 10.13.6 (17G7024)
- CPU: 2.8 GHz Intel Core i7
- RAM: 16 GB 1600 MHz DDR3
- GPU: NVIDIA GeForce GT 750M 2048 MB - Intel Iris Pro 1536 MB
- Test tool: gnu-time on MacOS (gtime -v ...)

## Resutls

| Program | Data | Items | CPU (%) | RAM (KBs) | Time (s) |
|---------|------|-------|-----|-----|------|
| gpu-bitonicsort | bit_list.txt | 1048576 | 95 | 20784 | 0.72 |
| gpu-mergesort | list.txt | 1000000| 98 | 24368 | 2.36 |
| cpu-mergesort | list.txt | 1000000| 98 | 14176 | 0.58 |
| cpu-quicksort | list.txt | 1000000| 99 | 4776 | 0.33 | 

## Detail
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

### cpu-mergesort
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

### cpu-quicksort
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

