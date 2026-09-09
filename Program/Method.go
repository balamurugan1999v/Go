package main

import "fmt"

func main() {
	printHelloWorld()
	name := collectInformation("Enter your name")
	printTheValue(name)
}

func printHelloWorld() {
	fmt.Println("Helloworld")
}

func collectInformation(inp string) string {
	fmt.Println(inp)
	var value string
	fmt.Scanln(&value)
	return value
}

func printTheValue(in string) {
	fmt.Println("Here is the final answer", in)
}

/*

Output:
Helloworld
Enter your name
Balamurugan
Here is the final answer Balamurugan

*/
