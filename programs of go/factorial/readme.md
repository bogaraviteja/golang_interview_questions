Write a Golang program to find the factorial of a given number (Using Recursion)

Examples
Factorial of 5 = 5*4*3*2*1 = 120

Factorial of 10 = 10*9*8*7*6*5*4*3*2*1 =

Approach to solve this problem
Step 1: Define a function that accepts a number (greater than 0), type is int.
Step 2: If the number is 1, then return 1.
Step 3: Otherwise, return num*function(num-1).