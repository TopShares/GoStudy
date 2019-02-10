package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string)([]byte, error) {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		return nil,err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("worong status code: %d",resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)

}
