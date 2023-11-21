package main

import "log"

// basic types  (numbers, strings, boolens)
var myInt int
var myUint uint
var myfloat float32
var myfloat64 float64

func main() {
	myInt = 10
	myUint = 20
	myfloat = 10.1
	myfloat64 = 100.1

	log.Println(myInt, myUint, myfloat, myfloat64)

	myString := "John"

	log.Println(myString)

	myString = "Roman"

	var myBool = true
	myBool = false

	log.Println(myBool)
}
