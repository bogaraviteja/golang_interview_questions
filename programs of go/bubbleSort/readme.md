Write a Golang program to check whether a given array is sorted or not (Using Bubble Sort Technique)

Examples
Input arr = [7, 15, 21, 26, 33] => Array is already sorted.
Input arr = [7, 5, 1, 6, 3] => Array is not sorted.

Approach to solve this problem
Step 1: Iterate the array from the 0th index to n-1.
Step 2: Iterate the array from the 0th index to n-1-i, where i is the index of the above loop.
Step 3: If swap is not taking place in the first iteration, then print that “array is already sorted”.
Step 4: If swap occurs, then print “array is not sorted”.