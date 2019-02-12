package main

import (
	"fmt"
	"regexp"
)

const text = `My email is test@test.com
email is sdi@si.cn
email is dkdd@22.cn.cn
`


func main() {
	pattern := `[a-zA-Z0-9]+@.+`
	re := regexp.MustCompile(pattern)
	math := re.FindAllString(text, -1)
	fmt.Println(math)
	pattern = `^([a-zA-Z0-9_-]+)@([a-zA-Z0-9_-]+)(\.[a-zA-Z0-9_-]+)$`
	for _, m := range math {
		fmt.Println(m)
	}
}
