package main

import "fmt"

func main() {
	var max_uint8 uint8 = 255
	var max_uint16 uint16 = 65535
	var max_uint32 uint32 = 4294967295
	var max_uint64 uint64 = 18446744073709551615

	var max_int8 int8 = 127
	var max_int16 int16 = 32767
	var max_int32 int32 = 2147483647
	var max_int64 int64 = 9223372036854775807

	var min_int64 int64 = -9223372036854775808

	var same_as_uint8 byte
	var same_as_int32 rune
	var same_as_uint32_or_uint64 uint
	var same_size_as_uint int

	/* Using variables in print just to avoid red underlines */
	fmt.Print(max_uint8, max_uint16, max_uint32, max_uint64)
	fmt.Print(max_int8, max_int16, max_int32, max_int64)
	fmt.Print(min_int64)
	fmt.Print(same_as_uint8, same_as_int32, same_as_uint32_or_uint64, same_size_as_uint)
}
