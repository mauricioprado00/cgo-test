#include <stdio.h>

void printHelloWorld()
{
    printf("Hello, World!\n");
}

void printHelloYou(char *name)
{
    printf("Για σασ, %s!\n", name);
}

char *getHelloWorld()
{
    return "Hallo, Welt!";
}

void printHelloWorldXTimes(int times)
{
    for (int i = 0; i < times; i++)
    {
        printf("Bonjour le monde!\n");
    }
}