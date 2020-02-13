package testing
//this test is communication sequence process
//
import (
	"fmt"
	"testing"
)

func service() string{
	//time.Sleep(time.Microsecond * 50)
	return "-----service Done"
}

func otherTask()  {
	fmt.Println("working on other something else")
	//time.Sleep(time.Microsecond * 100)
	fmt.Println("other task is done")
}

func Test1(t *testing.T){
	fmt.Println(service())
	//time.Sleep(time.Second)
	otherTask()
}

func AsyncService() chan string{
	retCh:= make(chan string,0)  //后面的参数是设定通道的buffer,可选，设定为0也可以？？
    go func() {
    	ret:= service()
    	fmt.Println("return result.")
    	retCh <- ret
    	fmt.Println("service exit.")
	}()
	return retCh
}

func TestAsync(t *testing.T){
	retCh:= AsyncService()
	otherTask()
	fmt.Println(<-retCh)   //取值的时候才到
}


