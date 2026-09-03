package main

import "fmt"

func main() {
	var inputStringValue string
	fmt.Scan(&inputStringValue)
	fmt.Println(inputStringValue)
	//Response
	//Hello world --> input
	//Hello       --> output

	//Response
	//Helloworld  --> input
	//Helloworld  --> output

	var inputIntValue int = 0
	fmt.Println(inputIntValue)
	fmt.Scan(&inputIntValue)
	fmt.Println(inputIntValue)
	//Response:
	//0 --> output
	//10 --> input
	//10 --> output

}
