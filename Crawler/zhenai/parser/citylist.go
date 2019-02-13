package parser

import (
	"GoStudy/Crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1) // return [][][]byte

	result := engine.ParseResult{}
	for _,m:= range matches{
		result.Items = append(result.Items, "City" + string(m[2])) // city name
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parsefunc: ParseCity,
		})
		//for _, subMatch := range m{
		//	fmt.Printf("%s", subMatch)
		//}
		//fmt.Printf("City:%s, URL:%s\n",m[2],m[1])
	}
	return result
	//fmt.Printf("Matches found: %d\n", len(matches))

}