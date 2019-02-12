package parser

import (
	"GoStudy/Go-Spider/infra/errors"
	//"GoStudy/Crawler/fetcher"
	"fmt"
	"net/http"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//contents, err := fetcher.Fetch("http://album.zhenai.com/u/95387448")
	//contents, err := ioutil.ReadFile("profile_test_data.html") //local
	url := "http://album.zhenai.com/u/95387448"
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}


	// @获取403禁止访问的网页
	//request.Header.Add("Host","Host: album.zhenai.com")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60")


	//defer resp.Body.Close()

	client := http.DefaultClient

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(errors.ErrorResponse)
	}

	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		fmt.Println(errors.ErrorStatusCode)
	}

	//result := ParseProfile(contents)
	fmt.Println("\n---------------------")
	//fmt.Println(result)
}