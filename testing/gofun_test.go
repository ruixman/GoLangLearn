package testing

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGo(t *testing.T){
	for i:=0;i<=10;i++ {
		go func(i int){
			fmt.Print( " ",i,"  ")
		}(i)
	}

    time.Sleep(time.Second * 1)
}

func TestGo2(t *testing.T){  //注意局部的值变化
	for i:=0;i<=10;i++ {
		go func(){
			fmt.Print( " ",i,"  ")
		}()
	}
	time.Sleep(time.Second * 1)
}

func TestCounter(t *testing.T){
	counter:=0
	var mut sync.Mutex
	var wg  sync.WaitGroup
	for i:=0;i<=5000;i++{
		wg.Add(1)
	   go func() {
	   	defer func() {
	   		mut.Unlock()
		}()
	   	mut.Lock()
	   	counter++
	   	wg.Done()
	   }()
	}
	wg.Wait()
	//time.Sleep(1 * time.Second)
	t.Logf("counter value is: %d ",counter)
}

//func  TestCounterThreadSave(t *testing.T){
//	var mut sync.Mutex
//	counter:=0
//
//}
