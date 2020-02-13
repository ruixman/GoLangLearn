package testing
//用方法代替函数，没有什么新意，只不过在运算时符号比较有新意而已
import (
	"testing"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}
type point struct{
	x string
	y int
}

func (a point) calcu(b int) int{
	return b
}

func Test2(t *testing.T)  {
	a1:= point{"hel",2}
	c:=a1.calcu(234)
	t.Log(c)
}

