int linear_search(int arr[], int len, int target)
{
    int pos = -1;
    for (int i = 0; i < len; ++i)
        if (arr[i] == target)
            return (pos = i);
    return pos;
}
