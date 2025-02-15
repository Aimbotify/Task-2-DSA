# Wave Array Rearrangement in Go

This repository contains a Go program that rearranges an array of integers into a "wave pattern" based on a given integer `x`.

## Problem Description
Given an array of integers and an integer `x` (where `x ≥ 1`), the program rearranges the array into consecutive blocks of `2x+1` elements such that:
- The element at index `x` (1-based) is the maximum in its block.
- Elements to the left of the peak are in non-decreasing order.
- Elements to the right of the peak are in non-increasing order.

## Files Included
- `main.go`: Contains the Go code to perform the wave array rearrangement.

## How to Run the Program
1. Clone the repository.
2. Navigate to the project directory.
3. Run the program using:
   ```bash
   go run main.go
   ```

## Example Output
```
Wave array: [10 15 20 90 67 23 5 100 44 34 1]
```

## Explanation
- The array is divided into blocks of `2x+1` elements.
- Each block is sorted and rearranged to ensure that the `x`-th element is the peak, with elements on the left in ascending order and elements on the right in descending order.



