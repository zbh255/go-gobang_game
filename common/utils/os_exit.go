package utils

import (
	"github.com/foxsuagr-sanse/go-gobang_game/common/db"
	"os"
	"os/signal"
	"syscall"
)

/*
信号	值	动作	说明
SIGHUP	1	Term	终端控制进程结束(终端连接断开)
SIGINT	2	Term	用户发送INTR字符(Ctrl+C)触发
SIGQUIT	3	Core	用户发送QUIT字符(Ctrl+/)触发
SIGILL	4	Core	非法指令(程序错误、试图执行数据段、栈溢出等)
SIGABRT	6	Core	调用abort函数触发
SIGFPE	8	Core	算术运行错误(浮点运算错误、除数为零等)
SIGKILL	9	Term	无条件结束程序(不能被捕获、阻塞或忽略)
SIGSEGV	11	Core	无效内存引用(试图访问不属于自己的内存空间、对只读内存空间进行写操作)
SIGPIPE	13	Term	消息管道损坏(FIFO/Socket通信时，管道未打开而进行写操作)
SIGALRM	14	Term	时钟定时信号
SIGTERM	15	Term	结束程序(可以被捕获、阻塞或忽略)
SIGUSR1	30,10,16	Term	用户保留
SIGUSR2	31,12,17	Term	用户保留
SIGCHLD	20,17,18	Ign	子进程结束(由父进程接收)
SIGCONT	19,18,25	Cont	继续执行已经停止的进程(不能被阻塞)
SIGSTOP	17,19,23	Stop	停止进程(不能被捕获、阻塞或忽略)
SIGTSTP	18,20,24	Stop	停止进程(可以被捕获、阻塞或忽略)
SIGTTIN	21,21,26	Stop	后台程序从终端中读取数据时触发
SIGTTOU	22,22,27	Stop	后台程序向终端中写数据时触发
*/


func ONEXIT()  {
	Sig := make(chan os.Signal,1)
	signal.Notify(Sig,syscall.SIGINT,syscall.SIGHUP,syscall.SIGQUIT,syscall.SIGILL,syscall.SIGFPE,syscall.SIGTERM)
	select {
	case <- Sig:
		// 捕获到信号
		var d db.DB = &db.SetData{}
		dblink := d.RedisInit(0)
		if _,err := dblink.FlushAll().Result(); err != nil {
			println(err)
		} else {

		}
		defer func() {
			dblink.Close()
			os.Exit(0)
		}()
	}
}
