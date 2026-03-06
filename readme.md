# Go Training – 

This repository contains my own exercises, practice work, and peer-reviewed assignments from the
Programming with Google Go Specialization from Coursera.

The main goal of this repository is to learn Go.

- Go language fundamentals
- Software design using functions, methods, and interfaces
- Concurrency using goroutines and channels

---

## Repository Structure

The repository structure follows the organization of the specialization and my own modules:

.
├── main/
│   ├── course2/
│   │   ├── module1/
│   │   ├── module2/
│   │   │   └── peer/
│   │   │       └── main.go
│   │   ├── module3/
│   │   └── module4/
│   │
│   ├── course3/
│   │   ├── module1/
│   │   ├── module2/
│   │   ├── module3/
│   │   │   └── peer/
│   │   │       └── main.go
│   │   └── module4/
│   │       └── main.go
├── concurrency/
│   ├── worker-poool/
│
├── go.mod
└── README.md

---

## Conventions

- Each courseX directory represents a course in the specialization.
- Each moduleX directory corresponds to a weekly module.
- The peer directory contains peer-reviewed assignments.
- Each exercise is self-contained and executable using go run.
- my own exercises will be in its folder category

---

## Example Exercises

Some of the topics implemented in this repository include:

- Idiomatic use of functions and methods in Go
- Interface-based design
- Concurrency patterns using goroutines, channels, and mutexes
- Classic concurrency problems (e.g. Dining Philosophers)
- Coordination between goroutines using safe synchronization patterns

---

## Running the Exercises

From the root of the repository, you can run any exercise like this:

go run main/course3/module3/peer/main.go

Each main.go file can be executed independently.

---

## Personal Goal

This repository serves as a learning log, a future reference, and a foundation for deeper exploration of Go concurrency and robust system design.

The code here is not meant to be perfect or production-ready.
It intentionally reflects learning progression throughout the course.

---

## License

Educational and personal use only.
Course content belongs to Coursera and the University of California, Irvine.
The code written in this repository is free to use for learning purposes.
