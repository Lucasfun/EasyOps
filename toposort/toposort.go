package toposort

import (
	"fmt"
	"github.com/Lucasfun/EasyOps/testYAML"
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

func TopoSort(inDegree map[string] int, outDegree *map[string][]string) {
	out := *outDegree

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
		for index := range out[v_de]{ //遍历某个value
			inDegree[out[v_de][index]] -= 1
			if inDegree[out[v_de][index]] == 0{
				q.Enqueue(out[v_de][index])
			}
		}
		fmt.Print(v_de)
	}
}

func concurrent_tppo(inDegree map[string] int, outDegree *map[string][]string) {
	out := *outDegree

	q := Queue{}
	q.Init(len(inDegree))

	level := 1 //记录并表示vertex层数

	//第一次将入度为0的点入队
	for key,value := range inDegree{  //O(N) = len(inDegree),实际执行 m <= N次
		if value == 0{
			q.Enqueue(key)//入度为0入栈
		}
	}

	//队空前执行 ： 同一层次点打印结束后，依次将各点出队并检查其next是否符合入队条件
	for q.Head != q.Tail  {
		fmt.Printf("第%d层安装 : ",level)
		level ++

		count := 0 // 记录本层次出队unit个数
		for index := range q.Items{
			head_v := q.Items[index]
			if value,isExist := inDegree[head_v];isExist == true{
				if value == 0{
					count ++
					fmt.Print(head_v)
				}
			}
		}
		fmt.Println()

		for i := 0;i < count;i ++{
			v_de := q.Dequeue()
			for index := range out[v_de]{ //遍历某个value
				inDegree[out[v_de][index]] -= 1
				if inDegree[out[v_de][index]] == 0{
					q.Enqueue(out[v_de][index])
				}
			}
		}
		//fmt.Print(v_de)
	}
}

func countIndegree(outDegree *map[string][]string) map[string]int {
	out := *outDegree
	inDegree := map[string]int{}
	for key,_ := range out{
		inDegree[key] = 0
	}

	for _,value := range out{
		for index := range value{
			inDegree[value[index]] ++
		}
	}
	return inDegree
}

func Test_toposort() {
	//yaml文件获得组件依赖--Graph
	graph := testYAML.GraphConfig{}
	graph.Init()

	//Graph转化为出度表 -- map[string][]string
	outDegree := graph.New_outDegree()

	inDegree := countIndegree(outDegree)

	//TopoSort(inDegree,outDegree)
	concurrent_tppo(inDegree,outDegree)

}

