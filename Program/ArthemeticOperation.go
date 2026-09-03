package main

import "fmt"

func main() {
	var value1 int = 101
	var value2 int = 20
	var value3 int = value1 + value2
	var value4 int = value1 - value2
	var value5 int = value1 * value2
	var value6 int = value1 / value2
	var value7 int = value1 % value2
	fmt.Println(value3)
	fmt.Println(value4)
	fmt.Println(value5)
	fmt.Println(value6)
	fmt.Println(value7)
	//output:
	//121
	//81
	//2020
	//5
	//1

	var fvalue1 float64 = 101.1
	var fvalue2 float64 = 1.11
	var fvalue3 float64 = fvalue1 + fvalue2
	var fvalue4 float64 = fvalue1 - fvalue2
	var fvalue5 float64 = fvalue1 * fvalue2
	var fvalue6 float64 = fvalue1 / fvalue2
	//var fvalue7 float64 = fvalue1 % fvalue2
	fmt.Println(fvalue3)
	fmt.Println(fvalue4)
	fmt.Println(fvalue5)
	fmt.Println(fvalue6)
	//fmt.Println(fvalue7)
	// output:
	// 102.21
	// 99.99
	// 112.221
	// 91.08108108108107
	//invalid operation: operator % not defined on fvalue1 (variable of type float64)
	var intValue1 = 1
	var floatValue1 = 1.1
	var outputValue1 = intValue1 + int(floatValue1)
	var outputValue2 = float64(intValue1) + floatValue1
	fmt.Println(outputValue1)
	fmt.Println(outputValue2)
	//output:
	// 2
	// 2.1

	var string1 = "Hello"
	var string2 = "World"
	var string3 = string1 + string2
	fmt.Println(string3)
	//output:
	//HelloWorld

}
