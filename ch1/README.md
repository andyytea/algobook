# The Role of Algorithms in Computing

An **algorithm** is any well-defined computational procedure that takes some value, or set of values, as **input** and produces some value, or set of values, as **output**.

An algorithm can also be viewed as a tool for solving a well-specified **computational problem**.

In practice, algorithms are only interesting if they are able to solve a general, yet well-specified problem.

Since many programs use sorting as an intermediate step, sorting is a fundamental operation in CS. Which algorithm is best for a given application depends on the number of items to be sorted, the extent to which the items are already sorted, possible restrictions on item values, the architechture of the computer, and the kind of storage devices to be used (main memory, disks, or even tapes).

An algorithm is said to be **correct** if, for every input instance, it halts with the correct output. Additionally, we say that a correct algorithm **solves** the given computational problem.

However, *incorrect* algorithms can be situationally useful.

In all applications of algorithms, the two common threads are as follows:

1. They have many possible solutions, yet none of which can be identified to be the "best" solution available.
2. They have practical applications for real life problems.

A **data structure** is a way to store and organize data in order to facilitate access and modifications. Each of which have situationally advantageous/disadvantageous traits.

## Hard Problems

The book is mainly about efficiency of algorithms, however there are problems for which no efficient solution is known. Problems that are `NP-complete` have no efficient algorithms. However, nobody has ever been able to prove that an efficient algorithm for one cannot exist.

## Parallelism

In order to perform more computations per second, chips are designed to contain numerous processing cores. In order to get the best performance from multi-core computers, we need to design algorithms with parallelism in mind.

## Performance

Computers are fast, but not infinitely fast. Additionally, memory can be inexpensive, but it is not free. Thus, computing time is bounded by the resources available (mainly time and memory).

The total system performance depends on choosing efficient algorithms as much as on choosing fast hardware.
