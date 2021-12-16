package utils

import (
	"regexp"
	"strconv"
	"unicode"
)

func GetSpell(input string) string {
	var  spell string
	r, _ := regexp.Compile("/([a-z]+).dat")
	if r.FindStringSubmatch(input)!=nil {
		spell = r.FindStringSubmatch(input)[1]
	}
	return spell
}

func DictType(input string) string {
	url, _ := regexp.Compile("(http|https):\\/\\/([\\w.]+\\/?)\\S*")
	tmpfilename := url.FindStringSubmatch(input)
	if tmpfilename != nil {
		return "remote"
	}
	return "local"
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
func IsNumRange(str string) int {
	var num int
	var err error
	if num, err = strconv.Atoi(str);
	err != nil {
		return 0
	}
	if num>10 || num<1 {
		return 0

	}
	return num
}




