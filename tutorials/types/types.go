package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var print = fmt.Println

func main() {

	print(reflect.TypeOf(34.43))     // float64
	print(reflect.TypeOf("Blatter")) // string
	print(reflect.TypeOf(true))      // bool
	print(reflect.TypeOf(723))       // int

	floatStr := "4.23"
	floatValue, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		print("Error parsing float:", err)
	} else {
		print(floatValue, reflect.TypeOf(floatValue))
	}

	number := 43
	stringNumber := strconv.Itoa(number)
	if err != nil {
		print("Error parsing float:", err)
	} else {
		print(stringNumber, "has been converted into a", reflect.TypeOf(stringNumber))
	}
}
