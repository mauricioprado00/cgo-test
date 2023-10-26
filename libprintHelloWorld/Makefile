CC = gcc
CFLAGS = -Wall -shared -fPIC
LDFLAGS = -shared

LIBRARY_SRC = printHelloWorld.c
LIBRARY_OBJ = $(LIBRARY_SRC:.c=.o)
LIBRARY_HEADER = printHelloWorld.h
LIBRARY_NAME = printHelloWorld.so
# needs to be prefixed lib, otherwise linker needs to be specified with -l:printHelloWorld
LIBRARY_ARCHIVE = libprintHelloWorld.a

all: $(LIBRARY_NAME) $(LIBRARY_ARCHIVE)

$(LIBRARY_NAME): $(LIBRARY_OBJ)
	$(CC) $(LDFLAGS) -o $@ $(LIBRARY_OBJ)

$(LIBRARY_ARCHIVE): $(LIBRARY_OBJ)
	ar rcs $@ $(LIBRARY_OBJ)

%.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

headers: $(LIBRARY_HEADER)

clean:
	rm -f $(LIBRARY_OBJ) $(LIBRARY_NAME) $(LIBRARY_ARCHIVE) $(LIBRARY_ARCHIVE)

.PHONY: all headers clean
