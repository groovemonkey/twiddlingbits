package twiddlingbits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryStringToInt(t *testing.T) {
	tt := []struct {
		desc   string
		input  string
		expect int
	}{
		{
			desc:   "All zeros",
			input:  "00000",
			expect: 0,
		},
		{
			desc:   "All ones",
			input:  "11111",
			expect: 31,
		},
		{
			desc:   "just a number",
			input:  "01001",
			expect: 9,
		},
		{
			desc:   "test number - 1",
			input:  "01010",
			expect: 10,
		},
		{
			desc:   "test number - 2",
			input:  "00011",
			expect: 3,
		},
		{
			desc:   "test number - 3",
			input:  "11100",
			expect: 28,
		},
		{
			desc:   "length one, zero",
			input:  "0",
			expect: 0,
		},
		{
			desc:   "length one, one",
			input:  "1",
			expect: 1,
		},
		{
			desc:   "zero value",
			input:  "",
			expect: 0,
		},
		{
			desc:   "63-bit value",
			input:  "101010101010100101010111110101010101010010101011110000110010100",
			expect: 6148728415636021652,
		},
		{
			desc:   "64-bit value (overflow)",
			input:  "1010101010101001010101111101010101010100101010111100001100101000",
			expect: -6149287242437508312,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := BinaryStringToInt(tc.input)
			assert.Nil(t, err)
			assert.Equal(t, tc.expect, result, tc.input)
			// assert.Equal(t, tc.input, fmt.Sprintf("%b", result))
		})
	}
}

func TestMakeInt(t *testing.T) {
	tt := []struct {
		desc   string
		input  BinarySlice
		expect int
	}{
		{
			desc:   "All zeros",
			input:  BinarySlice{[]int{0, 0, 0, 0, 0}},
			expect: 0,
		},
		{
			desc:   "All ones",
			input:  BinarySlice{[]int{1, 1, 1, 1, 1}},
			expect: 31,
		},
		{
			desc:   "just a number",
			input:  BinarySlice{[]int{0, 1, 0, 0, 1}},
			expect: 9,
		},
		{
			desc:   "test number - 1",
			input:  BinarySlice{[]int{0, 1, 0, 1, 0}},
			expect: 10,
		},
		{
			desc:   "test number - 2",
			input:  BinarySlice{[]int{0, 0, 0, 1, 1}},
			expect: 3,
		},
		{
			desc:   "test number - 3",
			input:  BinarySlice{[]int{1, 1, 1, 0, 0}},
			expect: 28,
		},
		{
			desc:   "length one, zero",
			input:  BinarySlice{[]int{0}},
			expect: 0,
		},
		{
			desc:   "length one, one",
			input:  BinarySlice{[]int{1}},
			expect: 1,
		},
		{
			desc:   "zero value",
			input:  BinarySlice{[]int{}},
			expect: 0,
		},
		{
			desc:   "63-bit value",
			input:  BinarySlice{[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}},
			expect: 6148728415636021652,
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			result, err := MakeInt(tc.input)
			assert.Nil(t, err)
			assert.Equal(t, tc.expect, result, tc.input)
		})
	}
}
