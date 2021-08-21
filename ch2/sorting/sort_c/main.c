#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "sorting.h"

void print_arr(int *arr, int len) {
    printf("[ ");
    for (int i = 0; i < len; ++i) {
        printf("%d ", arr[i]);
    }
    printf("]\n");
}

int main(void) {
    char s[20];
    int num;
    int len;
    while (scanf("%s", s) == 1) {
        scanf("%d", &len);
        int *arr = malloc(len*sizeof(int));

        if (len > 0) {
            for (int i = 0; i < len; ++i) {
                scanf("%d", &num);
                arr[i] = num;
            }
        }

        printf("%s: ", s);
        print_arr(arr,len);
        if (strcmp(s, "insertion_sort")) {
            insertion_sort(arr, len);
        } else if (strcmp(s, "selection_sort")) {
            selection_sort(arr, len);
        } else if (strcmp(s, "merge_sort")) {
            merge_sort(arr, len);
        }

        printf("sorted: ");
        print_arr(arr,len);
        printf("\n");
        free(arr);
    }
}
