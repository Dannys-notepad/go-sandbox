# Data Structures and Algorithms

Small Go exercises for learning how common search and sorting algorithms work. These are standalone experiments, each with its own `main` function.

## Examples

- `binarySearch.go` contains a binary-search function and a small sample invocation. Binary search requires the input slice to be sorted.
- `selectionSort.go` contains helpers for finding the largest and smallest values, plus a selection-sort implementation. Its current `main` demonstrates finding the largest value; the sort call is commented out.

## Run

From the repository root, run one file at a time:

```sh
go run Data_Structures-and-Algorithms/binarySearch.go
go run Data_Structures-and-Algorithms/selectionSort.go
```

Or from this folder:

```sh
go run binarySearch.go
go run selectionSort.go
```

Both files currently belong to `package main`, so they are intended to be run individually rather than compiled together as one package.