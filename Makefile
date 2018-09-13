DDIR = deploy
ODIR = $(DDIR)/_output
SERVICES ?= web

all: test

test:
	go test ./...

dep:
	dep ensure

compile:
	@$(foreach var, $(SERVICES), GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/$(var)/bin/$(var) app/$(var)/main.go;)
