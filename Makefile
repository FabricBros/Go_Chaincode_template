TAG = 0.10

test:
	go test -cover -test.v

clean:
	go clean

build:
	go build .