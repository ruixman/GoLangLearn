package testing

import (
	"fmt"
	"sync"
	"testing"
)

func dataProduce(ch chan int,group *sync.WaitGroup){

	go func(){
		for i:=0;i<10;i++{
			ch <- i
		}
		close(ch)
		group.Done()   //此处显式标注下
	}()
}

func dataConsume(ch chan int, group *sync.WaitGroup){

	go func() {
		for {

		/**：如果chan 关闭了，则不需要判断是否访chan还有没有值(第15行），chan会返回一个该类型的默认值（如果是int，则为0）*/
			if data,ok:= <-ch;ok{
				fmt.Println(data)
			}else {
				break
			}

		}
		group.Done()
	}()
}

func TestCloseChannel(t *testing.T)  {
	var wg sync.WaitGroup
	ch:= make(chan int)
	wg.Add(1)
	dataProduce(ch,&wg)
	wg.Add(1)
	dataConsume(ch,&wg)
	wg.Wait()
}