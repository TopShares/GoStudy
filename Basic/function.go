package Basic

import "fmt"

//闭包， 匿名函数

func main() {
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)                      //值copy
		defer func() { fmt.Println("defer_closure i = ", i) }() //refer copy
		fs[i] = func() { fmt.Println("closure i = ", i) }       // refer copy
	}
	for _, f := range fs {
		f()
	}
}
