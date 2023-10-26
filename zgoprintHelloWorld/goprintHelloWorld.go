package main

// #cgo CFLAGS: -Wall
// #cgo LDFLAGS: -L/home/mauricio/www/signyourname-project/ocr/c-test/libprintHelloWorld/ -lprintHelloWorld
// void printHelloWorld();
import "C"

func main() {
    C.printHelloWorld()
}
