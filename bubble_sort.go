package main

import (
	"fmt"
)

func main() {
	a := [...]int{3,5,8,1,9}
	fmt.Println(a)

	count := len(a)
	for i := 0; i < count; i++ {
		for j := i+1; j < count; j++ {
			if a[i] > a[j]{
				temp := a[j]
				a[j] = a[i]
				a[i] = temp
			}
		}
	}
	fmt.Println(a)
}