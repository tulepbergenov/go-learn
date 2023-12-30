package main

import "fmt"

func main() {
	var v1 int // v1 == 0
	fmt.Println(v1)

	var v2 int = 42
	fmt.Println(v2)

	v3 := 5
	//v3 :=	5 // Error, for short syntax
	v3 = 10
	fmt.Println(v3)

	var v5, v6 = 1, 2
	v5, v6 = 3, 4
	v5, v6 = v6, v5

	v5, v6, v7 := 1, 2, 3

	fmt.Println(v7)

	var (
		v01 = 1
		v02 = "string"
	)

	_ = v01
	_ = v02

	// Style
	var GrettingText string = "Hello World"

	fmt.Println(GrettingText)

	// const earthRadiusInMeter int = 6371000
	// const earthRadiusInMeter = 6371000

	const (
		earthRadiusInMeter = 6371000
		year               = 2023
	)

	_ = earthRadiusInMeter

	x := 42
	ptr := &x

	*ptr = 21

	fmt.Println(x)
}
