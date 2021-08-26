# Elementary Data Structures

We examine the representation of dynamic sets by simple data structures that use pointers. The rudimentary data structures include: stacks, queues, linked lists, and rooted trees.

## Stacks and Queues

Stacks and queues are dynamic sets in which the element removed from the set by the `Delete` operation is pre-specified. In a **stack**, the element removed from the set is the one most recently inserted (i.e. implements LIFO policy). 

Similarly in a **queue** , the element removed from the set is the one that has been in the set for the longest time (i.e. implements FIFO policy).

### Stack

The `Insert` operation on a stack is often referred to as `Push` instead. The deletion operation, which does not take an argument, is often called `Pop`.

## Benchmarking Implementations

[Stack as Linked List vs. Array](eds_go)
