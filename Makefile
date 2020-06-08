run:
	go build
	./simple-log

vet:
	go fmt
	go vet
