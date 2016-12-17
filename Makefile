.PHONY: all depends test install clean

all: qrserve

depends:
	go get github.com/skip2/go-qrcode

qrserve: 
	go build

test:
	go test

install:
	go install

clean:
	go clean
