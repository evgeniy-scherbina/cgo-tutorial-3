#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main () {
    for (int i = 0; i < 1; i++) {
        char *str;

        /* Initial memory allocation */
        str = (char *) malloc(15);
        strcpy(str, "tutorialspoint");
        printf("String = %s,  Address = %u\n", str, str);

        free(str);
    }

    return(0);
}