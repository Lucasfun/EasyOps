package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chanSelect()
	//ContextSingleGoroutine()
	//ContextMultiGoroutine()
}

func chanSelect() {//select->case <- chan，通过该方式通知goroutine停止
	stop := make(chan bool)
	go func() {
		for{
			select {
			case <-stop:  //只能保证与下方的close()、stop <- true等动作同步，不保证接下来的代码也同步
				fmt.Println("监控退出【1】，停止...")
				//time.Sleep(time.Second * 1)
				//fmt.Println("监控退出【1】动作先执行...")
				return
			default:
				fmt.Println("goroutine 【1】监控中....")
				time.Sleep(2 * time.Second)
			}
		}
	}()


	//同时用一个channel stop 做不到同时停止若干个goroutine
	//go func() {
	//	for{
	//		select {
	//		case <-stop:
	//			time.Sleep(time.Second)
	//			fmt.Println("监控退出【1】，停止...")
	//			return
	//		default:
	//			fmt.Println("goroutine 【2】监控中....")
	//			time.Sleep(2 * time.Second)
	//		}
	//	}
	//}()

	time.Sleep(4 * time.Second)
	fmt.Println("可以了，通知监控停止")

	stop <- true
	//close(stop) //Context实现的核心原理
	fmt.Println("stopping")


	//为了检测监控是否停止，如果没有监控输出，就表示停止了
	time.Sleep(time.Second * 5)

	//新建goroutine，同样的原理
}

//Context控制单个goroutine
func ContextSingleGoroutine() {
	// context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
	// 然后我们使用context.WithCancel(parent)函数，创建一个可取消的子Context，然后当作参数传给 --
	// goroutine使用，这样就可以使用这个子Context跟踪这个goroutine
	ctx,cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for{
			// 在goroutine中，使用select调用<-ctx.Done()判断是否要结束，如果接受到值的话，
			// 就可以返回结束goroutine了；如果接收不到，就会继续进行监控
			select {
			case <- ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine 监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	// 那么是如何发送结束指令的呢？这就是示例中的cancel函数啦，它是我们调用context.WithCancel(parent)函数生成子Context的
	// 时候返回的，第二个返回值就是这个取消函数，它是CancelFunc类型的。我们调用它就可以发出取消指令，然后我们的监控goroutine就
	// 会收到信号，就会返回结束
	cancel()

	// 为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

//Context控制多个goroutine

func ContextMultiGoroutine()  {
	ctx,cancel := context.WithCancel(context.Background())
	// 示例中启动了3个监控goroutine进行不断的监控，每一个都使用了Context进行跟踪，当我们使用cancel函数通知取消时，这3个goroutine都会
	// 被结束。这就是Context的控制能力，它就像一个控制器一样，按下开关后，所有基于这个Context或者衍生的子Context都会收到通知，这时就可以
	// 进行清理操作了，最终释放goroutine，这就优雅的解决了goroutine启动后不可控的问题

	go watch(ctx,"监控【1】")
	go watch(ctx,"监控【2】")
	go watch(ctx,"监控【3】")

	time.Sleep(10*time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5*time.Second)
}

func watch(ctx context.Context,name string)  {
	for{
		select {
		case <- ctx.Done():
			fmt.Println(name,"监控退出，停止了...")
			return
		default:
			fmt.Println(name,"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}