#include <stdlib.h>

void swap(int *a, int *b) 
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

void insertion_sort(int *arr, int len) 
{
    for (int i = 1; i < len; ++i) {
        for (int j = i; j > 0 && arr[j] < arr[j-1]; --j) {
            swap(&arr[j], &arr[j-1]);
        }
    }
}

void selection_sort(int *arr, int len) 
{
    int i, j, min;
    for (i = 0; i < len; ++i) {
        min = i;
        for (j = i+1; j < len; ++j) {
            if (arr[j] < arr[min]) min = j;
        }
        swap(&arr[i], &arr[min]);
    } 
}

int partition(int *arr, int l, int r) 
{
    int pivot = arr[r];
    int partition_index = l;
    for (int i = l; i < r; ++i) {
        if (arr[i] <= pivot) {
            swap(&arr[i], &arr[partition_index]);
            partition_index++;
        }
    }
    swap(&arr[partition_index], &arr[r]);
    return partition_index;
}

void quick_sort_rec(int *arr, int l, int r) 
{
    if (l < r) {
        int p = partition(arr, l, r);
        quick_sort_rec(arr, l, p-1);
        quick_sort_rec(arr, p+1, r);
    }
}

void quick_sort(int *arr, int len)
{
    quick_sort_rec(arr, 0, len-1);
}

void merge(int *dest, const int *src1, int len1,
                      const int *src2, int len2) 
{
    int pos1 = 0;
    int pos2 = 0;
    for (int i = 0; i < len1 + len2; ++i) {
        if (pos1 == len1 || (pos2 < len2 && src2[pos2] < src1[pos1])) {
            dest[i] = src2[pos2]; 
            ++pos2;
        } else {
            dest[i] = src1[pos1];
            ++pos1;
        }
    }
} 

void merge_sort(int*arr, int len) 
{
    if (len <= 1) return;
    int llen = len/2;
    int rlen = len - llen;

    int *left_arr = malloc(sizeof(int)*llen);
    int *right_arr = malloc(sizeof(int)*rlen);

    for (int i = 0; i < llen; ++i) left_arr[i] = arr[i];
    for (int i = 0; i < rlen; ++i) right_arr[i] = arr[i];

    merge_sort(left_arr, llen);
    merge_sort(right_arr, rlen);

    merge(arr, left_arr, llen, right_arr, rlen);

    free(left_arr);
    free(right_arr);
}
