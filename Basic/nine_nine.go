package Basic

import (
	"fmt"
)

func main() {

	/*
		九九乘法表
	*/
	//外层控制行
	for i := 1; i <= 9; i++ {
		//内层控制列
		for j := 1; j <= i; j++ {
			// if j>i {
			// 	break
			// }
			fmt.Printf("%d*%d=%d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
