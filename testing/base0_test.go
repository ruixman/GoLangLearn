package testing

import (
	"fmt"
	"testing"
)

func TestDefine(t *testing.T){

	x := "hello"
	for _, x1 := range x {    //此处x1也可能为x
		x2 := x1 + 'A' - 'a'
		fmt.Printf("%c", x2) // "HELLO" (one letter per iteration)

	}
	t.Log(x)

	//与上面的等效
	for _, x := range x {        //此处x为重新定义，
		x := x + 'A' - 'a'       //此处x为重新定义
		t.Logf("%c", x)  // "HELLO" (one letter per iteration)
		t.Log(x)
	}
}
