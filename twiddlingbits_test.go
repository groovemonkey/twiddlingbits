package twiddlingbits_test

import (
	"fmt"

	"github.com/groovemonkey/twiddlingbits"
)

func ExampleBinaryStringToInt() {
	r, _ := twiddlingbits.BinaryStringToInt("00101")
	fmt.Println(r)
}

func ExampleBinarySlice() {
	bits := []int{0, 1, 0, 0, 1}
	bs := twiddlingbits.BinarySlice{bits}

	r, _ := twiddlingbits.MakeInt(bs)
	fmt.Println(r)
}
