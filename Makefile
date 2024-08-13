.PHONY: ex01 ex00
all: tests

tests: ex01 ex00

ex00:
	cd ./ex00/coinsort && go test -v .

ex01:
	cd ./ex01/coinsortbench && go test -benchmem -bench=.
	