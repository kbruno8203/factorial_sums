package factorial_sums

import (
    "errors"
    "math/big"
    "strconv"
    "strings"
)

// GetFactorialSums calculates the sum of the digits of the factorial of n.
// It returns -1 if n is less than or equal to 0.
// It raises an error if the factorial has more than 100,000 digits.
func GetFactorialSums(n int) (int, error) {
    if n <= 0 {
        return -1, nil
    }

    fact := big.NewInt(1)
    for i := 1; i <= n; i++ {
        fact.Mul(fact, big.NewInt(int64(i)))
    }

    // Check if the factorial has more than 100,000 digits
    factStr := fact.String()
    if len(factStr) > 100000 {
        return 0, errors.New("factorial has more than 100,000 digits")
    }

    // Calculate the sum of the digits
    sum := 0
    for _, digit := range strings.Split(factStr, "") {
        d, err := strconv.Atoi(digit)
        if err != nil {
            return 0, err
        }
        sum += d
    }

    return sum, nil
}