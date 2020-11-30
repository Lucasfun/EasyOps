package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x := 0; x <= 10 ;x ++{
			naturals <- x //send
		}
		close(naturals)
	}()

	go func() {
		for{
			x,ok := <- naturals
			if !ok{
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for{
		//x, ok := <- squares
		//if !ok {
		//	break
		//}
		fmt.Println(<- squares)
		time.Sleep(time.Second * 1)
	}

}
