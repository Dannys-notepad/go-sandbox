# Go Learning Journal

This repository is a growing collection of projects and exercises built while learning Go. Each folder is a separate stop along the way, from small algorithm experiments to larger practice programs. The code may evolve as new concepts are explored.

## Projects

| Folder | What it contains |
| --- | --- |
| [Data Structures and Algorithms](Data_Structures-and-Algorithms/README.md) | Early experiments with searching and sorting algorithms. |
| [Student Management Program](Student_Management_Program/README.md) | A small student-record program practicing structs, methods, slices, and functions. |

## Requirements

- Go installed and available as `go` in your terminal.
- Go 1.22 or newer for `Student_Management_Program`, which imports `math/rand/v2`.

## Running the examples

Run an individual source file from the repository root:

```sh
go run Data_Structures-and-Algorithms/binarySearch.go
go run Data_Structures-and-Algorithms/selectionSort.go
go run Student_Management_Program/main.go
```

The algorithm examples each define their own `main` function, so run them one at a time. The folders are learning projects, not yet a single combined application.