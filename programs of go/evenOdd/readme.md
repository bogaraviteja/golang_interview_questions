Write a Golang program to find odd and even numbers using bit operation

Examples
Input num = 5 => 101 & 1 = 001 => True, i.e., Odd; else num would be Even.

Approach to solve this problem
Step 1: Define a method that accepts a number.
Step 2: Perform & operation with that number.
Step 3: If the & operator returns a non-zero value, then that number would be odd.
Step 4: Else, the number would be even.