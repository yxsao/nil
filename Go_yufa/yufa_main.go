package main //程序所属包

import (
	"fmt"
	"runtime"
	"sync"
	"time"
) //导入依赖包

const ChangLiang int = 0 //常量声明
var (
	Pri string = "全部结束" //一般类型声明
	WG  sync.WaitGroup
	xie int
)
var chaint chan int = make(chan int, 2)

type jishi2 int //结构声明

type Ixie interface { //声明接口
}

func Send(n int) { //函数定义
	chaint <- n
}

func Xiec1() {
	fmt.Println("开始协程1")
	x := 0
	for i := 0; i < 200; i++ {
		x += i
		time.Sleep(time.Millisecond * 1)
	}
	Send(1)
	//WG.Add(1)
	fmt.Println("协程1结束", x)
	xie = 1
}
func Xiec2() {
	fmt.Println("开始协程2")
	x := 0
	for i := 0; i < 100; i++ {
		x += i
		time.Sleep(time.Millisecond * 2)
	}
	Send(2)
	//WG.Add(1)

	fmt.Println("协程2结束", x)
	xie = 2
}

func main() { //main函数
	fmt.Println("当前CPU核心数:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() - 4)
	jishi1 := time.Now().Unix()
	fmt.Println("计时开始:", time.Unix(int64(jishi1), 0).Format("2006-01-02 15:04:05"))
	go Xiec2() //协程
	go Xiec1()

	for {
		//fmt.Println("等待")
		num := <-chaint //等待chan信号
		switch num {
		case 1:
			fmt.Println("协程1先完成")
			//break
		case 2:
			fmt.Println("协程2先完成")
			//break
		default:
			fmt.Println("等待协程")

		}
		//time.Sleep(time.Millisecond * 3000)
		if xie >= 0 {
			break
		}

	}
	fmt.Println(<-chaint, Pri)
}
