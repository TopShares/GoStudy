package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("profile_test_data.html") //local
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents)
	fmt.Println("---------------------")
	fmt.Println(result)
}