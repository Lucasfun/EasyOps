package main

import (
	"log"
	"time"
)

func doSomething(id int) {
	log.Printf("before do job:(%d)\n",id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d)\n",id)
}

func main() {

	go doSomething(1)
	go doSomething(2)
	go doSomething(3)

	time.Sleep(4 * time.Second)
	//并发队列，总体用时3秒
}
