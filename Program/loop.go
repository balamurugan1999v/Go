package main

import (
	"fmt"
)

func main() {
	case1()
	case2()
	case3()
	case4()
	case5()
}

func case3() {
	fmt.Println("Case3 custom testing")
	iter := 1
	max := 100
	for ; iter < max; iter++ {
		if iter%3 == 0 {
			fmt.Print(iter, " ")
		}
	}
	fmt.Println()
}

/*

Output:
Case3 custom testing
3 6 9 12 15 18 21 24 27 30 33 36 39 42 45 48 51 54 57 60 63 66 69 72 75 78 81 84 87 90 93 96 99

*/

func case4() {
	fmt.Println("Case4 custom testing")
	iter := 1
	max := 100
	for ; ; iter++ {
		if iter < max {
			break
		}
		if iter%3 == 0 {
			fmt.Print(iter, " ")
		}
	}
	fmt.Println()
}

/*
Output:
Case4 custom testing

*/

func case5() {
	fmt.Println("Case5 custom testing")
	iter := 1
	max := 100
	for iter < max {
		if iter%3 == 0 {
			fmt.Print(iter, " ")
		}
		iter++
	}
	fmt.Println()
}

/*
Output:
Case5 custom testing
3 6 9 12 15 18 21 24 27 30 33 36 39 42 45 48 51 54 57 60 63 66 69 72 75 78 81 84 87 90 93 96 99
*/

func case2() {
	fmt.Println("Enterring into case2")
	fmt.Println("Infinite loop")
	lastValue := 10
	iter := 0
	for {
		if iter == lastValue {
			break
		}
		fmt.Println(iter)
		iter++
	}
	fmt.Println("Existing from case2")
}

/*

Enterring into case2
Infinite loop
0
1
2
3
4
5
6
7
8
9
Existing from case2

*/

func case1() {
	var threshold int = 10
	var i int
	for i = 5; i < threshold; i++ {
		fmt.Println(i)
	}
}

/*

5
6
7
8
9

*/
