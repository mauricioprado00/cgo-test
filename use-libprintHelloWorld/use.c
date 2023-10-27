#include "../libprintHelloWorld/printHelloWorld.h"
#include <stdio.h>

int main() {
    printHelloWorld();
    printHelloYou("C User");
    printf("%s\n", getHelloWorld());
    printHelloWorldXTimes(3);
    return 0;
}
