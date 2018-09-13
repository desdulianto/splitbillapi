REGISTRY  = registry.bukalapak.io/bukalapak
DDIR      = deploy
ODIR      = $(DDIR)/_output
SERVICES ?= web
VERSION   = $(shell git show -q --format=%h)
NOCACHE   = --no-cache

all: test

test:
	go test ./...

dep:
	dep ensure

compile:
	@$(foreach var, $(SERVICES), GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/$(var)/bin/$(var) app/$(var)/main.go;)

build:
	@$(foreach var, $(SERVICES), docker build $(NOCACHE) -t $(REGISTRY)/splitbillapi/$(var):$(VERSION) -f ./deploy/$(var)/Dockerfile .;)
