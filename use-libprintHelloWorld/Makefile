CC = gcc
CFLAGS = -Wall
LIBDIR = -L/home/mauricio/www/signyourname-project/ocr/c-test/libprintHelloWorld/
LIBS =  -lprintHelloWorld
# this also works
#LIBS =  -l:libprintHelloWorld.a
# this also works but needs the .so to be in the same directory later.
#LIBS =  -l:printHelloWorld.so

use: use.c
	$(CC)  $(CFLAGS) -o use use.c $(LIBDIR) $(LIBS)

.PHONY: clean

clean:
	rm -f use
