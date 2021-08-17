# Analyzing Algorithms

Most often, it's computational time that we want to measure, however we may situationally want to consider factors such as memory, communication bandwidth, or computer hardware generally.

Before we can even begin to analyze an algorithm, we must have some assumption about the model of the implementation technology that we will use (one that includes the resources of that technology and their costs). For all implementations here, we will assume the random-access machine (RAM) model of computation as our implementation technology (with primitives like add, subtract, load, store, etc.). In the RAM model, instructions execute one at a time, with no concurrent operations for the sake of simplicity. 

In general, the time taken to compute grows with the size of the input.

For a more precise definition of **running time**, the running time on a particular input is the number of primitive operations or "steps" executed.

In some cases, we will be concerned with **average-case** analysis rather than **worst-case**. Naturally, this introduces the concept of **probablistic analysis**.

A common abstraction we use for analyzing running time is using **order of growth** or more commonly known as **Big-O**.
