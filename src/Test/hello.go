package main

import ("fmt")


var sharedVar = "Shared Variable"

var (
	a = 5
	b = 6
	c = 10
)

func main() {


	for i:=0; i<100; i++ {
		fmt.Println(i)
	}
}
