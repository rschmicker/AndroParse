all:
	go install -ldflags "-s -X utils.Version=$(shell date -u +%Y-%m-%d-%H-%M-%S)" ./...

install: all

configure:
	glide install

unittest: configure
	! gofmt -l $(shell glide novendor -no-subdir) | grep .
	go vet $(shell glide novendor)
	go test -ldflags -s -cover $(shell glide novendor | grep -v cmd)

update:
	glide update

clean:
	rm -rf ~/bin/webserver