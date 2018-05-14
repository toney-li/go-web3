package utils

import (
    "math/big"
    "fmt"
)

func IntToHex(n *big.Int) string {
    return fmt.Sprintf("0x%x", n)
}
