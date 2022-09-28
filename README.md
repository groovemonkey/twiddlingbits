# twiddlingbits

*This is a toy.* It's some binary-munging functionality I needed for $REASONS. I've made it into a tiny library and kept the name as `twiddlingbits` to discourage you from using this for production.

In Go, it's easy to go from an int to a string of the binary representation of that int, e.g.

```
i := 7
fmt.Sprintf("%b", i)
```

It's not super easy to go the other way, though. UNTIL NOW (cue sales music):

```
result, _ := BinaryStringToInt("00101") // 5
```

I also created something called a `BinarySlice`, which has an integer slice that contains only ones and zeroes (`BinarySlice.Bits`), with a max length of 63 because I didn't need to deal with overflows for this.

```
bits := []int{0, 1, 0, 0, 1}
bs := BinarySlice{bits}

result, _ := MakeInt(bs.Bits) // 9
```


## TODO

- switch from int to uint so we can have 64 bits.
