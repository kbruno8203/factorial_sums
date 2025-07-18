package factorial_sums

import (
    "errors"
    "math/big"
    "strconv"
    "strings"
)

func GetFactorialSums(n int) (int, error) {
    if n <= 0 {
        return -1, nil
    }

    fact := big.NewInt(1)
    for i := 1; i <= n; i++ {
        fact.Mul(fact, big.NewInt(int64(i)))
    }

    factStr := fact.String()
    if len(factStr) > 100000 {
        return 0, errors.New("factorial has more than 100,000 digits")
    }

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