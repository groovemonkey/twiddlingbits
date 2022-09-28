package twiddlingbits

import (
	"fmt"
)

// A binary slice is a slice of ints with a max length of 63, consisting of only 1s and 0s
type BinarySlice struct {
	Bits []int
}

// A BinarySlice may contain only ones and zeros
func Valid(b *BinarySlice) bool {
	// Limit of 63 bits
	if len(b.Bits) > 63 {
		return false
	}

	// only ones and zeros allowed
	for _, v := range b.Bits {
		if v != 1 && v != 0 {
			return false
		}
	}
	return true
}

// MakeInt creates an int from a BinarySlice, e.g. BinarySlice{[]int{0, 1, 0, 0, 1}} => 9
func MakeInt(b BinarySlice) (int, error) {
	var number int

	for i, d := range b.Bits {
		// Already initialized with zeros, only ones need to move
		if d == 1 {
			// Start shifting from the end of the string (lowest-order bit)
			// right-most bit doesn't shift; each successive bit shifts one further
			numShifts := len(b.Bits) - i - 1
			current := 0b1 << numShifts
			// bitwise OR them together
			number = number | current
		} else if d != 0 {
			// Only ones and zeros allowed
			return 0, fmt.Errorf("invalid BinarySlice: %v", b)
		}
	}
	return number, nil
}

// BinaryStringToInt takes a string of zeros and treats them as a Binary number, returning an int
// "00011" => 3, "01010" => 10
func BinaryStringToInt(s string) (int, error) {
	var number int

	for i, d := range s {
		// Already initialized with zeros, only ones need to move
		if string(d) == "1" {
			// Start shifting from the end of the string (lowest-order bit)
			// right-most bit doesn't shift; each successive bit shifts one further
			numShifts := len(s) - i - 1
			current := 0b1 << numShifts
			number = number | current
		} else if string(d) != "0" {
			// Only ones and zeros allowed
			return 0, fmt.Errorf("invalid binary string: %s", s)
		}
	}
	return number, nil
}
