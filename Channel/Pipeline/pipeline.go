package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x :=0; x < 50 ;x ++{
			naturals <- x //send
		}
	}()

	go func() {
		for{
			x := <- naturals
			squares <- x * x
		}
	}()

	for{
		fmt.Println(<- squares)
	}

}
