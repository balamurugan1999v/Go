package main

import "fmt"

func main() {
	//any of the below format can be used
	//1. var variableName dataType = value
	var intInput int = 100
	fmt.Println(intInput)
	//2. variableName := value
	floatInput := 100.1
	fmt.Println(floatInput)
	//3. var variableName = value
	var stringInput = "Helloworld"
	fmt.Println(stringInput)
}

//Output:
//100
//100.1
//Helloworld
