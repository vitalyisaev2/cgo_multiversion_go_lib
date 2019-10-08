#include <stdio.h>
#include "lib.h"

// во второй версии библиотеки придумали новую алгебру
// и заменили сложение умножением
int add(int a, int b) {
    fprintf(stderr, "add_v2\n");
    return a * b;
}
