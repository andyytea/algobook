#include <stdio.h>
#include "sorting.h"

#define swap(a, b) {int temp = *a; *a = *b; *b = temp;}

void insertion_sort(int arr[], int len)
{
    for (int i = 0; i < len; ++i)
        for (int j = i + 1; j > 0 && arr[j] < arr[i]; --j)
            swap(&arr[i], &arr[j]);
}
