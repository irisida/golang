package main

import (
	"fmt"
	"math"
)

func main() {

	// demonstration with ints but the same principle is applied to
	// other data types too
	var (
		val8  int8
		val16 int16
		val32 int32
		val64 int64
	)

	// declared but uninitialised
	fmt.Println(val8, val16, val32, val64)

	// initialise with max value associated to the datatype
	val8 = math.MaxInt8
	val16 = math.MaxInt16
	val32 = math.MaxInt32
	val64 = math.MaxInt64

	fmt.Println(val8, val16, val32, val64)

	// add an increment
	val8++
	val16++
	val32++
	val64++

	// Go boomerangs the value back to the minimal value
	fmt.Println(val8+1, val16+1, val32+1, val64+1)

	// double check minimal values
	val8 = math.MinInt8
	val16 = math.MinInt16
	val32 = math.MinInt32
	val64 = math.MinInt64

	fmt.Println(val8, val16, val32, val64)

}
