package fetcher

import (
	"GoStudy/Go-Spider/infra/errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//var rateLimiter = time.Tick(10*time.Millisecond)

func Fetch(url string)([]byte, error) {
	//<-rateLimiter
	log.Printf(url)
	//request, err := http.Get(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return nil,err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60")


	//defer resp.Body.Close()
	//
	//if request.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("  Wrong status code: %d",request.StatusCode)
	//}

	client := http.DefaultClient

	response, err := client.Do(request)

	if err != nil {
		return nil, fmt.Errorf("%v",errors.ErrorResponse)
	}

	//fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, fmt.Errorf("  Wrong status code: %v",errors.ErrorStatusCode)
	}

	//utf8Content := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	return ioutil.ReadAll(response.Body)
	//return ioutil.ReadAll(request.Body)

	//if ok {
	//
	//	utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	//	return ioutil.ReadAll(utf8Content)
	//} else {
	//	return ioutil.ReadAll(response.Body)
	//}
}
