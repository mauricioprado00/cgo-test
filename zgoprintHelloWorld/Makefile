GO = go
BINARY = goprintHelloWorld
LIBPATH = ../libprintHelloWorld/
LIBSTATIC = libprintHelloWorld.a

all: $(BINARY)

$(BINARY): goprintHelloWorld.go
	# need to copy because go runs inside docker
	cp $(LIBPATH)$(LIBSTATIC) .
	$(GO) build -o $(BINARY) goprintHelloWorld.go

.PHONY: clean

clean:
	rm -f $(BINARY) $(LIBSTATIC)
