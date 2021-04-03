ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	cd lib/vdflib && cargo build --release
	cp lib/vdflib/target/release/libvdflib.so lib/
	go build -ldflags="-r $(ROOT_DIR)lib"

buildmacos:
	cd lib/vdflib && cargo build --release
	cp lib/vdflib/target/release/libvdflib.dylib lib/
	go build -ldflags="-r $(ROOT_DIR)lib"

test:
	go test -ldflags="-r $(ROOT_DIR)lib"

bench:
	go test -bench=. -benchmem -run=XXX -cpuprofile profile.out  -ldflags="-r $(ROOT_DIR)lib"

prof:
	go test -cpuprofile cpu_profile.out -memprofile mem_profile.out -ldflags="-r $(ROOT_DIR)lib"
	go tool pprof -png cpu_profile.out
	go tool pprof -png mem_profile.out
