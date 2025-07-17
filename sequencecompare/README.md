# Sequence Comparison

The challenge in this directory is focused on comparing three int slices. The following rules apply:

1. The input slices are defined as follows:
   1. A full slice must exist which is the first input to any function and is a range of numbers between any two positive values
   2. A partial slice must exist with all of the values from the full slice except a one value
   3. A partial slice must exist with all the values from the full slice except two values
2. The order of both the second and the third slice is irrelevant and the effective results should not change based on 
   their order
3. Design two functions, both of which take the full slice followed by the partial slice:
   1. The first function FindOneMissingSliceValue will compare the first slice with the second slice and return the missing integer
   2. The second function FindTwoMissingSliceValues will compare the first slice with the third slice and return both missing integers
4. Make no changes to any file other than challenge.go
5. Do not use any libraries including stdlib

The objective is to ensure the functions produce the least allocations as possible and are O(n) if that's possible.