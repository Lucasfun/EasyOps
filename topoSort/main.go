package main

import (
	"fmt"
)

type Queue struct {
	Items   []string  // 数组:items  数组大小:n
	Cnt     int
	Head    int       // Head 队头下标  Tail 队尾下标
	Tail    int
}

// 初始化一个大小为 capacity 的数组
func (q *Queue) Init(capacity int) {
	q.Items = make([]string,capacity)
	q.Cnt = capacity
	q.Head,q.Tail = 0,0
}

// 入队
func (q *Queue) Enqueue(item string) {
	if q.Tail == q.Cnt {  // Tail 到尾部了
		if q.Head == 0 {  // 真的没空间了
			fmt.Println("The queue is full!")
			return
		}else {           // 数据搬移
			for i := q.Head;i < q.Tail; i++ {
				q.Items[i - q.Head] = q.Items[i]
			}
			q.Tail = q.Tail - q.Head
			q.Head = 0
		}

	}
	q.Items[q.Tail] = item
	q.Tail++
	return
}

// 出队
func (q *Queue) Dequeue() string{
	if q.Head == q.Tail {
		fmt.Println("The queue has no item!")
	}
	ret := q.Items[q.Head]
	q.Items[q.Head] = ""
	q.Head++
	return ret
}

func TopoSort(inDegree map[string] int, outDegree map[string][]string) {
	q := Queue{}
	q.Init(len(inDegree))

	for key,value := range inDegree{  //O(N) = len(inDegree),实际执行 m <= N次
		if value == 0{
			q.Enqueue(key)
			//delete(inDegree,key)
		}
	}

	for q.Head != q.Tail  {
		v_de := q.Dequeue()
		for index := range outDegree[v_de]{ //遍历某个value
			inDegree[outDegree[v_de][index]] -= 1
			if inDegree[outDegree[v_de][index]] == 0{
				q.Enqueue(outDegree[v_de][index])
			}
		}
		fmt.Print(v_de)

		//for key,value := range inDegree{ //O(N) = len(inDegree),实际执行 m <= N次
		//	if value == 0{
		//		q.Enqueue(key)
		//		delete(inDegree,key)
		//	}
		//}
	}
}

func countIndegree(outDegree map[string][]string) map[string]int {
	inDegree := map[string]int{}
	for key,_ := range outDegree{
		inDegree[key] = 0
	}

	for _,value := range outDegree{
		for index := range value{
			inDegree[value[index]] ++
		}
	}
	return inDegree
}

func main() {

	outDeree := map[string] []string{
		"A": {"D","E"},
		"B": {"E"},
		"C": {"F"},
		"D": {"G"},
		"E": {"G"},
		"F": {"H"},
		"G": {"H"},
	} //出度表

	inDegree := countIndegree(outDeree)

	TopoSort(inDegree,outDeree)
}
