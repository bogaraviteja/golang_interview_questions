Write a Golang program to find the sum of digits for a given number

Examples
num = 125 => 1 + 2 + 5 = 8
num = 101 => 1 + 0 + 1 = 2
num = 151 => 1 + 5 + 1 = 7

Approach to solve this problem
Step 1: Define a function that accepts numbers(num); type is int.
Step 2: Start a true loop until num becomes 0 and define res:=0.
Step 3: Find modulo and add to res.
Step 4: Divide num by 10.
Step 5: Return res.