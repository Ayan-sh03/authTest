package util //

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP() (string, error) {
	// Define the range for the OTP (5 digits)
	min := int64(10000)
	max := int64(99999)

	// Generate a cryptographically secure random number within the defined range

	randomInt, err := rand.Int(rand.Reader, new(big.Int).Sub(big.NewInt(max), big.NewInt(min)))
	if err != nil {
		return "", err
	}

	// Add the minimum value to ensure a 5-digit OTP
	otpValue := randomInt.Int64() + min

	// Format the OTP as a string with leading zeros
	otp := fmt.Sprintf("%05d", otpValue)

	return (otp), nil
}
