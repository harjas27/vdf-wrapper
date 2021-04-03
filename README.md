# VDF Wrapper
---
This repository contains a go wrapper for VDF implementation in rust(taken from https://github.com/poanetwork/vdf)

## Usage
### Build
* On linux based OS:
```sh
make build
```
* On macOS:
```sh
make buildmacos
```

### Testing
This will run the tests listed in `vdf_test.go` which cover the basic testcases of generating and verifying
```sh
make test
```

## Benchmarking
### Generating the performance benchmarks
```sh
make bench
```
* generate and verify testcase:
```text
Version         |   Difficulty   |  Operations  |   NanoSeconds/op  |   Bytes/op    |   Allocs/op   |
----------------|----------------|--------------|-------------------|---------------|---------------|
Wrapper(Rust)   |   100          |  1           |   1427637682      |   1632        |   4           | 
                |   1000         |  1           |   1666627794      |   1632        |   4           |  
                |   10000        |  1           |   2283725710      |   1632        |   4           |
                |   100000       |  1           |   8857166929      |   1728        |   5           |
----------------|----------------|--------------|-------------------|---------------|---------------|
Go              |   100          |  2           |   776854280       |   322415756   |   3154336     | 
                |   1000         |  2           |   2526245854      |   2698829064  |   26177189    |  
                |   10000        |  2           |   17168590230     |   22612145312 |   221400958   | 
                |   100000       |  1           |   158626758580    |   216820251240|   2131684085  |
```
* verify testcase with 100 difficulty
```text
Version         |   Difficulty   |  Operations  |   NanoSeconds/op  |   Bytes/op    |   Allocs/op   |
----------------|----------------|--------------|-------------------|---------------|---------------|
Wrapper(Rust)   |   100          |  2           |   772468279       |   576         |   1           |
----------------|----------------|--------------|-------------------|---------------|---------------|
Go              |   100          |  3           |   373219950       |   95635781    |   946215      | 
```
### Analysis
* Time taken per operation - Time taken for verification is of the similar order. In case of generate, as the difficulty increases the time taken by the wrapper version is significantly less than
that by the go version. For a difficulty of 10000, time taken by the go version is around `17 seconds`. However the wrapper version takes only `3 seconds` which is similar to the time taken by the pure rust library for the same inputs

* Memory per operation - Both the metrics bytes used and allocations show that the wrapper version is significanlty better than the go version

### Profiling
There are tests written in `vdf_test.go` which test the generate and verify functions for both this wrapper version and using the go version. In order to
create the cpu and memory profiles, run:
```sh
make prof
```
This will generate the CPU and memory profiles for the tests. In order to visualize these profiles:
* install Graphviz(https://graphviz.org/download/)
* this will create png files for the profiles 
* for samples, see `profile001.png` for cpu and `profile002.png` for memory

## Integration 
The file `vdf.go` exposes a public function `New(difficulty int, input [32]byte)` which can be used to initialize
a VDF. The `Execute` function can be used to run the created VDF and `Verify` function can be used to verify the proof

### With the Harmony client
This repository exposes the similar functions to create, execute and verify as done by the https://github.com/harmony-one/vdf repo. All the 
public functions have been implemented with the same signature as that to the current go VDF repo. Thus, by changing the imported go module, this repo can be integrated with the client 
An integration with the harmony client that builds successfully can be found here - https://github.com/harjas27/harmony/tree/use_rust_vdf