#include <stdio.h>
#include "lib.h"

// во второй версии библиотеки придумали новую алгебру
// и заменили сложение умножением
int add(int a, int b) {
    printf("add_v2");
    return a * b;
}
