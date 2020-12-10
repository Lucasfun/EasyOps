package unit_install

import (
	"log"
)

type Queue struct {
	Items []*Unit // units数组:items  数组大小:n
	Cnt   int     // 容量
	Head  int     // Head 队头下标  Tail 队尾下标
	Tail  int
}

// 初始化一个容量为 capacity 的数组
func (q *Queue) Init(capacity int) {
	q.Items = make([]*Unit, capacity)
	q.Cnt = capacity
	q.Head, q.Tail = 0, 0
}

//返回一个Queue变量引用
func GetQueue(capacity int) *Queue {
	return &Queue{
		Items: make([]*Unit, capacity),
		Cnt:   capacity,
		Head:  0,
		Tail:  0,
	}
}

// 入队
func (q *Queue) Enqueue(item *Unit) {
	if q.Tail == q.Cnt { // Tail 到尾部了
		for i := q.Head; i < q.Tail; i++ { //数据搬迁
			q.Items[i-q.Head] = q.Items[i]
		}
		q.Tail = q.Tail - q.Head
		q.Head = 0
	}
	q.Items[q.Tail] = item
	q.Tail++
}

// 出队
func (q *Queue) Dequeue() *Unit {
	if q.Empty() {
		log.Fatal("Dequeue when queue is empty!")
	}
	ret := q.Items[q.Head]
	q.Items[q.Head] = nil
	q.Head++
	return ret
}

//队空？
func (q *Queue) Empty() bool {
	if q.Head == q.Tail {
		return true
	}
	return false
}
