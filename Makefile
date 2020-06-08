GO=go

run: simple-log
	./simple-log

simple-log: *.go
	$(GO) build

vet:
	$(GO) fmt
	$(GO) vet
