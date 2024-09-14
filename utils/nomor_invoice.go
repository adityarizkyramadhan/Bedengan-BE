package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateNomorInvoice() string {
	result := "INV-"
	result += RandomNumber()
	return result
}

func RandomNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate random number between 0 and 99999999, formatted to 8 digits
	return fmt.Sprintf("%08d", r.Intn(100000000))
}
