package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string)([]byte, error) {
	request, err := http.Get(url)
	if err != nil{
		return nil,err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60")


	//defer resp.Body.Close()

	if request.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("  worong status code: %d",request.StatusCode)
	}

	//utf8Content := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	return ioutil.ReadAll(request.Body)

	//if ok {
	//
	//	utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	//	return ioutil.ReadAll(utf8Content)
	//} else {
	//	return ioutil.ReadAll(response.Body)
	//}
}
