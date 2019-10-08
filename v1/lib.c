#include <stdio.h>
#include "lib.h"

// в первой версии библиотеки сложение - это тривиальная операция
int add(int a, int b) {
    printf("add_v1");
    return a + b;
}
