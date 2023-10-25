CC = gcc
CFLAGS = -Wall
#LIBS = -L ../libprintHelloWorld/ -l printHelloWorld
LIBS = /home/mauricio/www/signyourname-project/ocr/c-test/libprintHelloWorld/libprintHelloWorld.so

use: use.c
	$(CC) $(CFLAGS) -o use use.c $(LIBS)

.PHONY: clean

clean:
	rm -f use
