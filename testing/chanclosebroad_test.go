package testing
//此代码说明如何传递消息到chan，以及关闭routing
import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool  {
	select {
	case <-cancelChan :
		return true
	default:
		return  false
	}
}

func TestCancel(t *testing.T){
	chan1:= make(chan struct{})
	for i:=0;i<5;i++{
		go func(i int,chancel chan struct{}) {
			for{
				if isCancelled(chan1){
					break
				}
				time.Sleep(time.Microsecond *15)
			}
			fmt.Println(i,":cancel is done!")
		}(i,chan1)
	}
	cancel(chan1)
	time.Sleep(time.Second)

}

func cancel(chan1 chan struct{}) {

	//方式一：给一个空结构:给空相当于消息chan结尾，但结果只有一个routing关闭
	//chan1 <- struct{}{}

    //方式二：使用close方式（正解），会关闭所有的routing
    close(chan1)

}
