#include <stdio.h>
#include "lib.h"

// в первой версии библиотеки сложение - это тривиальная операция
int add(int a, int b) {
    fprintf(stderr, "add_v1\n");
    return a + b;
}
