package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x :=0; x < 100 ;x ++{
			naturals <- x //send
		}
		close(naturals)
	}()

	//squarer
	go func() {
		for x := range naturals{
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares{  //这里无缓存的Channel吗，每次循环都输出一个？ -- >每次Channel有货久被拿出来输出
		//fmt.Printf("%d ",len(squares))
		fmt.Println(x)
	}

}
