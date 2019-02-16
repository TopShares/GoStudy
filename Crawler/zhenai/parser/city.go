package parser

import (
	"GoStudy/Crawler/engine"
	"regexp"
)

var (
	profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`"href=(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult{

	matches := profileRe.FindAllSubmatch(contents, -1)

	//fmt.Println(string(contents))
	//fmt.Println(len(matches))


	result := engine.ParseResult{}
	for _,m := range matches{
		//result.Items = append(result.Items, "User" + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parsefunc: ParseProfile,
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents,-1)
	for _,m := range matches{
		result.Requests= append(result.Requests,
			engine.Request{
				Url:string(m[1]),
				Parsefunc:ParseCity,
			})
	}
	return result
}