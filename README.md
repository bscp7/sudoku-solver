Sudoku solver
=============

Sudoku solver runs very basic algorithm to solve Sudoku puzzle. It can take single or multiple sudoku puzzles in input file. The algorithm is tested against 10,000 easy and medium sudoku puzzles.

# Help

```bash
go run . -h
Usage of /var/sudoku-solver:
  -d int        add delay in processing
  -f string     sudoku puzzle file (9 X 9)
  -m	        multiple sudoku puzzles (81 X any)
```

# Solve single sudoku puzzle

```bash
❯ go run . -d 0 -f ./examples/1.txt
filepath: ./examples/1.txt
 ------------------------------------ 
|   | 8 |   |   |   | 1 | 7 | 2 |   |
 ------------------------------------ 
| 4 |   |   | 7 | 3 | 8 |   | 6 | 1 |
 ------------------------------------ 
|   |   | 5 |   | 9 | 6 | 8 |   | 4 |
 ------------------------------------ 
| 5 | 2 |   |   |   |   |   |   |   |
 ------------------------------------ 
|   |   | 9 | 4 | 8 | 3 | 1 |   |   |
 ------------------------------------ 
|   |   |   |   |   |   |   | 9 | 7 |
 ------------------------------------ 
| 6 |   | 8 | 1 | 2 |   | 3 |   |   |
 ------------------------------------ 
| 2 | 1 |   | 8 | 7 | 5 |   |   | 6 |
 ------------------------------------ 
|   | 5 | 7 | 3 |   |   |   | 1 |   |
 ------------------------------------ 
Solved: false
TimeTaken: 0s
 ------------------------------------ 
| 3 | 8 | 6 | 5 | 4 | 1 | 7 | 2 | 9 |
 ------------------------------------ 
| 4 | 9 | 2 | 7 | 3 | 8 | 5 | 6 | 1 |
 ------------------------------------ 
| 1 | 7 | 5 | 2 | 9 | 6 | 8 | 3 | 4 |
 ------------------------------------ 
| 5 | 2 | 4 | 9 | 1 | 7 | 6 | 8 | 3 |
 ------------------------------------ 
| 7 | 6 | 9 | 4 | 8 | 3 | 1 | 5 | 2 |
 ------------------------------------ 
| 8 | 3 | 1 | 6 | 5 | 2 | 4 | 9 | 7 |
 ------------------------------------ 
| 6 | 4 | 8 | 1 | 2 | 9 | 3 | 7 | 5 |
 ------------------------------------ 
| 2 | 1 | 3 | 8 | 7 | 5 | 9 | 4 | 6 |
 ------------------------------------ 
| 9 | 5 | 7 | 3 | 6 | 4 | 2 | 1 | 8 |
 ------------------------------------ 
Solved: true
TimeTaken: 361.986µs
```

# Solve many sudoku puzzles in sequence

```bash
❯ go run . -d 10 -f ./examples/multi-puzzle.txt -m
```
