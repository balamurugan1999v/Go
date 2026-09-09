package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Entering into If file")
	printMethod("PrintCase 1")
	value1 := collectInput("Enter the value 1")
	value2 := collectInput("Enter the value 2")
	flag := ifCondtionOne(value1, value2)
	strCoversion := fmt.Sprintf("%t", flag)
	printMethod(strCoversion)
}

func printMethod(inp string) {
	log.Println(inp)
}

func collectInput(inp string) int {
	log.Println(inp)
	var input int
	fmt.Scanln(&input)
	return input
}

func ifCondtionOne(value1 int, value2 int) bool {
	if value1 > value2 {
		printMethod("Value1 is greater than value2")
		return false
	} else if value1 < value2 {
		printMethod("Value1 is less than value2")
		return false
	} else {
		printMethod("Both value1 and value2 are equal")
		return true
	}
}

/*


output:
2026/09/09 14:34:03 Entering into If file
2026/09/09 14:34:03 PrintCase 1
2026/09/09 14:34:03 Enter the value 1
5
2026/09/09 14:34:06 Enter the value 2
10
2026/09/09 14:34:07 Value1 is less than value2
2026/09/09 14:34:07 false

*/
