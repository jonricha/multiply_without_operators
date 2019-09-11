# multiply_without_operators

This repo is showing implementations of multiplying 2 numbers in Go without using the * or / operators.

RecursiveMultiply uses recursion and has some issues with large numbers due to stack overflow
LoopMultiply is similar to RecursiveMultiply except that it uses a loop instead of recursion.  This can take a long time and is again, not ideal.
MemoMultiply is a better implementation using recursion and memoization to take advantage of previous calculations.
