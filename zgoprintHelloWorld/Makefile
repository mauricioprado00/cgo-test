GO = go
BINARY = goprintHelloWorld

all: $(BINARY)

$(BINARY): goprintHelloWorld.go
	$(GO) build -o $(BINARY) goprintHelloWorld.go

.PHONY: clean

clean:
	rm -f $(BINARY)
