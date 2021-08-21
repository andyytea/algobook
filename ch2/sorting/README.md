# Sorting
Formally, we define the **sorting problem** as:

**Input:** A sequence of `n` numbers `[a_1, a_2, ..., a_n]`

**Output:** A permutation (reordering) `[a'_1, a'_2, ... a'_n]` of the input sequence such that `a'_1 <= ... <= a'_n`

We call any given input sequence as an **instance** of a problem.

[Insertion Sort](sort_c/sorting.c#L10-L17)

We formally state that the loop invariant is:

```
At the start of each iteraion of the outer for loop, the subarray
A[1 ... j-1] consists of elements originally in A[1 ... j-1], but
in sorted order.
```

We use loop invariants to understand why an algorithm is correct by showing three things about it:

1. Initialization (loop invariant holds before first iteration)
2. Maintenance (true before an iteration of the loop and remains true afterwards)
3. Termination (when the loop terminates, the invariant gives us a useful property to show the algorithm is correct) 

Example:

`loop invariant`: There is no index `0 <= k <= i-1` so that `A[k] = target`\
Initialization: If the program exists before line 3, `pos = -1` is an acceptable value because it is not in `0 <= k <= i-1`.\
Maintenance: If the loop exits before exhausting all values, pos is set to the current index (which is equal) as desired.\
Termination: If the loop is exited after exhausting all values of i, then there is still no index in `0 <= k <= i-1` so returning -1 is acceptable.\
```c
int linear_search(int arr[], int len, int target)
{
    int pos = -1;
    for (int i = 0; i < len; ++i)
        if (arr[i] == target)
            return (pos = i);
    return pos;
}
```

# Divide and Conquer

Many of the most important algorithms are *recursive* in nature: to solve a given problem, they call themselves one or more times.

The algorithm for **merge sort** is the canonical example for the *divide-and-conquer* paradigm of algorithms.

The key to merge sort is the idea that a list with 0 or 1 elements is considered a sorted list. In the merge step, the left and right hand list are combined in the desired order such that the left and right input and single output at each step are all sorted list.

[Merge Sort](sort_c/sorting.c#L59-L94)

Note: my implementation of mergesort allocates memory to the heap. Depending on how you implement it, you might end up copying memory on every funciton call. Thus, `merge_sort()` requires enough heap memory to store the original array several times over.

Example: a list with `n = 8` elements

```
       8        [1 int array(s) w/ 8 element(s)]
      / \       
     /   \      
    /     \     
   4       4    [2 int array(s) w/ 4 element(s)]
  / \     / \   
 2   2   2   2  [4 int array(s) w/ 2 element(s)]
/ \ / \ / \ / \ 
1 1 1 1 1 1 1 1 [8 int array(s) w/ 1 element(s)]
```
Test sorting algorithms by compiling `gcc main.c sorting.c` and `cat test.txt | ./a.out`
