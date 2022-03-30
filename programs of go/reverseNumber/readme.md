Write a Golang program to reverse a number

Examples
Input num = 1432 => output = 2341
Input num = 9878 => output = 8789
Input num = 6785 => output = 5876

Approach to solve this problem
Step 1: Define a function that accepts a number (num); type is int.
Step 2: Define res = 0 variable and start a loop until num becomes 0.
Step 3: Find remainder = num % 10 and make a number.
Step 4: Divide num by 10.
Step 5: Return res.