package Crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code", resp.StatusCode)
		return
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",all)

}
