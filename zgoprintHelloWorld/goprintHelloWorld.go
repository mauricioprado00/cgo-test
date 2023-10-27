package main

// #cgo LDFLAGS: -L. -lprintHelloWorld
// void printHelloWorld();
// void printHelloYou(char *name);
// char *getHelloWorld();
// void printHelloWorldXTimes(int times);
import "C"

func main() {
    C.printHelloWorld()
    C.printHelloYou(C.CString("Go User"))
    C.printHelloWorldXTimes(5)
    var hello = C.getHelloWorld()
    println(C.GoString(hello))
}
