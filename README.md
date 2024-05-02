# Prime Number Checker

This is a command-line tool written in Go to check if a given number is prime or to find prime numbers up to a specified limit.

## Overview

This program provides two main functionalities:

1. **Prime Number Checker**: It checks if a given number is prime or not. If the number is not prime, it also provides its factors.
2. **Prime Number Generator**: It generates all prime numbers up to a specified limit.

## Usage

### 1. Prime Number Checker

To check if a number is prime and find its factors (if not prime), run the following command:

```
go run main.go <number>
```

Replace `<number>` with the integer you want to check.

Example:

```
go run main.go 17
```

Output:

```
17 is a prime number
```

To find factors of a non-prime number:

```
go run main.go <number>
```

Example:

```
go run main.go 24
```

Output:

```
24 is not a prime number, its factors are 2, 3, 4, 6, 8, 12, 24
```

### 2. Prime Number Generator

To generate all prime numbers up to a specified limit, use the `--seq` flag along with the limit:

```
go run main.go <limit> --seq
```

Replace `<limit>` with the integer limit.

Example:

```
go run main.go 30 --seq
```

Output:

```
Prime numbers upto 30 are 2, 3, 5, 7, 11, 13, 17, 19, 23, 29
```