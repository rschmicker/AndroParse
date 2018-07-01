all:
	@cd ./plugins/;	make; cd ..;
	go install .

install: all

configure:
	glide install

unittest: configure
	! gofmt -l $(shell glide novendor -no-subdir) | grep .
	go vet $(shell glide novendor)
	go test -cover $(shell find ./plugins/* -type d | grep -v cmd)

update:
	glide update

clean:
	rm -rf ~/bin/androparse
	cd ./plugins/;	make clean; cd ..;