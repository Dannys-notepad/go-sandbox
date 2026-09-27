# Student Management Program

A small command-line learning project that models students and course grades. It uses structs and methods to represent records, calculates a weighted CGPA from course units and scores, generates sample records, and includes helpers for displaying and searching students.

## Concepts practiced

- Structs, methods, and slices
- Pointer receivers and pointer parameters
- Iteration and conditional logic
- Binary search over student IDs
- Random sample data with `math/rand/v2`

## Requirements

Go 1.22 or newer, because the program imports `math/rand/v2`.

## Run

From the repository root:

```sh
go run Student_Management_Program/main.go
```

Or from this folder:

```sh
go run main.go
```

The current `main` function generates three sample students, calculates their CGPAs, and prints the record slice. Sample ages and scores are randomized, so output can vary between runs.