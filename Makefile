.PHONY: all depends test install clean

all: qrserve

depends:
	go get github.com/skip2/go-qrcode

qrserve: qrserve.go
	go build

test: qrserve.go qrserve_test.go
	go test

install:
	go install

clean:
	rm -f qrserve
