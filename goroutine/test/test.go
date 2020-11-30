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
	doSomething(1)
	doSomething(2)
	doSomething(3)
	//阻塞队列，整体耗时9秒
}
