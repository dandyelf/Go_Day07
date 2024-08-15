.PHONY: ex01 ex00 ex02
all: tests

tests: ex01 ex00

ex00:
	cd ./ex00/coinsort && go test -v .

ex01: bench cputop
	
bench:
	cd ./ex01/coinsortbench && go test -benchmem -bench=.

cputop:
	cd ./ex01 && go build . && ./ex01 && go tool pprof -top cpu.prof | head -n 17 > top10.txt

ex02:
	cd ex02/coinsortdocument && godoc -url "/pkg/leftrana/moneybag/ex02/coinsortdocument/" > doc.html