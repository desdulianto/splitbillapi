all: test web

web:
	cd app/web; GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../../bin/web

test:
	cd handler; go test

clean:
	rm bin/*
