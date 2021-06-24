package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const maxASCII = 127
const minASCII = 32

func main() {
	// Get a password length (8 characters is the default)
	var length uint64 = 8
	if len(os.Args) > 1 {
		/* 1. Transform argument into int64
		 *    Error out if the conversion fails
		 * 2. Ensure int64 is positive
		 * 3. Set the int64 as the length */

		l, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "You must input a valid number as an argument\n")
			os.Exit(1)
		}
		if l < 0 {
			l *= -1
		}
		length = uint64(l)
	}

	// Generate each character individually
	var pass string
	for i := uint64(0); i < length; i++ {
		// 1. Generate cryptographically secure number
		// 2. Ensure the number is within the predefined ASCII range
		// 3. Transform the number to a byte and append it to the password

		bSecNumber, _ := crand.Int(crand.Reader, big.NewInt(maxASCII-minASCII))
		bSecNumber.Add(bSecNumber, big.NewInt(minASCII))
		pass += string(byte(bSecNumber.Abs(bSecNumber).Uint64()))
	}

	// Print trailing new-line only if output is a terminal
	fileInfo, _ := os.Stdout.Stat()
	if (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		fmt.Printf("%v\n", pass)
	} else {
		fmt.Printf("%v", pass)
	}
}
