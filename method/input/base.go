package input

import (
	"bufio"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/log"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/utils"
	"io"
	"strings"
)

type Address struct {
	From string
	Path string
}
type  Word struct{
	Chinese  string
	Frequency  rune
	Spell string
}

type ResourceMethod interface {
	GetResource() []Word
}
type BaseAddress struct {
	Address Address
}
// 读取字典中数据
func (t *BaseAddress) ReadFile(reader *bufio.Reader) []Word {
	var wordlist []Word
	for {
		line, _,err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		strline := string(line)
		arrline := strings.Split(strline, " ")
		if len(arrline) != 2 {
			continue
		}
		if !utils.IsChineseChar(arrline[0])  ||utils.IsNumRange(arrline[1])==0 {
			log.Warning("invalid content.")
			continue
		}
		chinese := arrline[0]

		item := Word{
			Spell:     utils.GetSpell(t.Address.Path),
			Chinese:   chinese,
			Frequency: rune(utils.IsNumRange(arrline[1])),
		}
		wordlist = append(wordlist, item)
	}

	return wordlist
}


func SortWordList(wordlist []Word) []Word {
	for k, _ := range wordlist {
		for j := k+1;j<len(wordlist);j++ {
			if isPrime(wordlist[j],wordlist[k]){
				tmp := wordlist[k]
				wordlist[k] = wordlist[j]
				wordlist[j] = tmp
			}
			// 去重
			if wordlist[j].Chinese==wordlist[k].Chinese  {
				if j == len(wordlist)-1{
					wordlist = wordlist[:j-1]
				}else {
					wordlist = append(wordlist[:j], wordlist[j+1:]...)

				}
			}
		}
	}
	return wordlist
}
//test ok
func ChineseList(wordlist []Word) []string {
	var chinese []string
	//chinese := make([]string, 10)

	i := 0
	for _,tmp := range wordlist {
		if tmp.Chinese ==""{
			break
		}
		chinese = append(chinese,tmp.Chinese)
		i++
	}
	if len(chinese)>10 {
		return chinese[:10]
	}
	return chinese
}
// isPrime 当且仅当word1优先于word2 返回true
func isPrime(word1 Word,word2 Word) bool {
	if word1.Frequency>word2.Frequency {
		return true
	}
	if word1.Frequency==word2.Frequency {
		if strings.Compare(word1.Spell,word2.Spell) == -1 {
			return true
		}
	}
	return false
}