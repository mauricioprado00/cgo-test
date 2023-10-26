package main

// #cgo LDFLAGS: -L. -lprintHelloWorld
// void printHelloWorld();
import "C"

func main() {
    C.printHelloWorld()
}
