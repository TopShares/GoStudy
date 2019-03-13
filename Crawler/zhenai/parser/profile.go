package parser

import (
	"GoStudy/Crawler/engine"
	"GoStudy/Crawler/model"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//匹配换行  ([\s\S])
var infoRe  = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>(.*?)</div>`)
var nameRe = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>(.*?)</h1>`)
//<div class="des f-cl" data-v-3c42fade>北京 | 27岁 | 硕士 | 未婚 | 177cm | 12001-20000元</div>

var genderRe = regexp.MustCompile(`<div class="m-title" data-v-bff6f798>(.*?)</div>`)

func ParseProfile(contents []byte) engine.ParseResult{
	profile := model.Profile{}

	info := extractString(contents,infoRe)
	//"北京 | 27岁 | 硕士 | 未婚 | 177cm | 12001-20000元"
	profile.Name = extractString(contents,nameRe)
	matches := genderRe.FindAllSubmatch(contents,-1)
	//fmt.Println(len(matches))
	//fmt.Println(string(matches[0][1]))
	//fmt.Println(string(matches[1][1]))
	//fmt.Println(string(matches[2][1]))
	//fmt.Println(string(matches[3][1]))
	//fmt.Println(string(matches[4][1]))
	if len(matches) >4 {
		sex := strings.Replace(string(matches[4][1]),"的动态","",-1)
		if sex == "他" {
			profile.Gender = "male"
		} else {
			profile.Gender = "female"
		}
	}
	arr:=strings.Split(info,"|")

	if len(arr)>5 {
		profile.City = arr[0]
		arr[1] = strings.Replace(arr[1],"岁","",-1)
		profile.Age,_ = strconv.Atoi(strings.Replace(arr[1], " ", "", -1))
		profile.Education = strings.Replace(arr[2],"","",-1)
		profile.Marriage = strings.Replace(arr[3],"","",-1)
		profile.Height = strings.Replace(arr[4],"","",-1)
		profile.Income = strings.Replace(arr[5],"","",-1)
	}

	fmt.Println(profile)
	result := engine.ParseResult{Items:[]interface{}{profile}}
	//fmt.Printf("%s",profile)
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string  {
	match := re.FindSubmatch(contents)

	if len(match) >=2{
		return string(match[1])
	}else {
		return ""
	}
}