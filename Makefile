all: ex00

ex00:
	cd ./ex00/coinsort && go test -v .

ex01:
	cd ./cmd/site_server && go run .

clean:
	