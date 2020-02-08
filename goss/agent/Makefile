GOCMD=go
GOBUILD=$(GOCMD) build -ldflags="-s -w"

build_goss:
	GOOS=linux GOARCH=amd64  $(GOBUILD) -o goss
