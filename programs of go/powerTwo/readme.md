Golang Program to check whether given positive number is power of 2 or not, without using any branching or loop

Examples
Consider n = 16(00010000)

Now find x = n-1 => 15(00001111) => x & n => 0

Approach to solve this problem
Step 1: Define a method, where n and is an argument, returns type is int.
Step 2: Perform x = n & n-1.
Step 3: If x is 0, then the given number is power of 2; else not.