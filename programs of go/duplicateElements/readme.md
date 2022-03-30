Write a Golang program to find duplicate elements in a given range

We can solve this problem in two different ways. 

Method 1: 

Examples
Input Array = [1, 2, 3, 4, 4] => Range is from 1 to 5 but 4 is a duplicate element in this range.

Approach to solve this problem
Step 1: Define a method that accepts an array.
Step 2: Declare a visited map.
Step 3: Iterate the given array. If the element exists in the visited map, then return that element.
Step 4: Else, return -1.

Method 2: Using XOR operation

Examples
Input Array = [1, 2, 3, 4, 4] => Range is from 1 to 5 but 4 is duplicate in that range.

Range is from 1 to 5. => XOR => 0^1^2^3^4^4^0^1^2^3^4 => 4 (since 0^1=1).

Approach to solve this problem
Step 1: Define a method that accepts an array.
Step 2: Find the range value from the given array and define a variable xor, initialize with 0.
Step 3: Iterate the given array and do a xor operation with the array’s elements.
Step 4: Also perform xor operation from the lower range value to the higher range value.
Step 5: At the end, return the xor variable, non-zero value for duplicate element.