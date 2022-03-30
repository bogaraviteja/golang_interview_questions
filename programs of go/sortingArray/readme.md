Write a Golang program to search an element in a sorted array

Approach to solve this problem
Step 1: Iterate the array from the 0th index to n-1, where n is the size of the given array.
Step 2: Declare low=0th index and high=n-1. Start a for loop till low is lesser than high.
Step 3: Find mid=(low+high)/2, if the element at the middle is equal to key, then return mid index.
Step 4: If the element at mid is greater than the key, then make high = mid.
Step 5: If the element at mid is lesser than the key, then make low = mid + 1.
Step 6: If key is not present in the given array, then return -1.

Time Complexity: log2(n)