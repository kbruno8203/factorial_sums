# Factorial Sums

**Description:** short program that takes the factorial of a number `n` and then sums the digits of the result
- Eg. For `n = 10` we have `10! = 3628800` and the result sum `3 + 6 + 2 + 8 + 8  + 0 + 0 = 27`

**Purpose:** basic practice with Go

**Specifications:**
- returns `-1` if `n` is less than or equal to `0`
- returns `-2` if `n!` exceeds the 100,000 digit upper bound
- uses Stirling's approximation to estimate digits in a factorial for large `n`

**Project Structure:**
```
factorial_sums/
├── factorial_sums.go - program contents
├── factorial_sums_test.go - comprehensive testing
└── README.md         - This file
```


