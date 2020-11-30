package main

import (
	"time"
	"fmt"
)

func main() {
	//主程序结束退出时强行打断 spinner（goroutine)[打断goroutine只能主函数退出或者终结程序]
	go spinner(100 * time.Microsecond)
	const n  = 45
	fibN := fib(n)
	fmt.Println("feibonaqie(%d)",n,fibN)
}

func spinner(delay time.Duration)  {
	for {
		for _,r := range "ABCABC"{
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2{
		return x
	}else{
		return fib(x -1) + fib(x - 2)
	}
}