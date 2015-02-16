all:
	go build -o calc
test:
	go test ./compute
benchmark:
	go test -bench ./compute
