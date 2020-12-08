package topologicalSort

import (
	"fmt"
	"sync"
)

type Queue struct {
	mux sync.Mutex    //互斥锁，确保数据安全
	Items []*UnitType // 数组:items  数组大小:n
	Cnt   int
	Head  int    // Head 队头下标  Tail 队尾下标
	Tail  int
}

// 初始化一个容量为 capacity 的数组
func (q *Queue) Init(capacity int) {
	q.Items = make([]*UnitType, capacity)
	q.Cnt = capacity
	q.Head, q.Tail = 0, 0
}

// 入队
func (q *Queue) Enqueue(item *UnitType) {
	if q.Tail == q.Cnt { // Tail 到尾部了
		if q.Head == 0 { // TODO 实际情况不会出现队列空间不足的情景，初始化时队列容量应该就为所有组件的个数
			fmt.Println("The queue is full!")
			return
		} else { // 数据搬移
			for i := q.Head; i < q.Tail; i++ {
				q.Items[i-q.Head] = q.Items[i]
			}
			q.Tail = q.Tail - q.Head
			q.Head = 0
		}

	}
	q.Items[q.Tail] = item
	q.Tail++
}

// 出队
func (q *Queue) Dequeue() *UnitType {
	if q.Head == q.Tail {
		fmt.Println("The queue has no item!")
		return nil  // TODO 一般不会出现，若返回空Unit出现逻辑未知错误
	}
	ret := q.Items[q.Head]
	q.Items[q.Head] = nil
	q.Head++
	return ret
}
