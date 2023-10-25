CC = gcc
CFLAGS = -Wall -shared -fPIC
LDFLAGS = -shared

LIBRARY_SRC = printHelloWorld.c
LIBRARY_OBJ = $(LIBRARY_SRC:.c=.o)
LIBRARY_HEADER = printHelloWorld.h
LIBRARY_NAME = libprintHelloWorld.so

all: $(LIBRARY_NAME)

$(LIBRARY_NAME): $(LIBRARY_OBJ)
	$(CC) $(LDFLAGS) -o $@ $(LIBRARY_OBJ)

%.o: %.c
	$(CC) $(CFLAGS) -c -o $@ $<

headers: $(LIBRARY_HEADER)

clean:
	rm -f $(LIBRARY_OBJ) $(LIBRARY_NAME)

.PHONY: all headers clean
