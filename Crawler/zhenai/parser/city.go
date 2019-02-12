package parser

import (
	"GoStudy/Crawler/engine"
	"regexp"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{

	re:= regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	//fmt.Println(string(contents))
	//fmt.Println(len(matches))


	result := engine.ParseResult{}
	for _,m := range matches{
		result.Items = append(result.Items, "User" + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parsefunc: ParseProfile,
		})
	}
	return result
}