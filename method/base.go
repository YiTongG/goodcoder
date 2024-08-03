package method

import (
	"bufio"
	"errors"
	"icode.baidu.com/baidu/goodcoder/gongyitong/log"
	"io"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// word 单词结构体
type word struct {
	chinese   string
	frequency rune
	spell     string
}

// resource 一个基本资源
type resource struct {
	path string
}

// readFile 读取字典中数据
func (t *resource) readFile(reader *bufio.Reader) []word {
	var wordlist []word
	for {
		line, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		stringline := string(line)
		arrayline := strings.Split(stringline, " ")
		if len(arrayline) != 2 {
			continue
		}
		if !isChineseChar(arrayline[0]) || isNumRange(arrayline[1]) == 0 {
			log.Warning("invalid content.")
			continue
		}
		chinese := arrayline[0]

		item := word{
			spell:     getSpell(t.path),
			chinese:   chinese,
			frequency: rune(isNumRange(arrayline[1])),
		}
		wordlist = append(wordlist, item)
	}

	return wordlist
}

// sortWordList 给字排序
func sortWordList(wordlist []word) []word {
	for k := range wordlist {
		for j := len(wordlist) - 1; j >= k+1; j-- {
			if isPrime(wordlist[j], wordlist[j-1]) {
				tmp := wordlist[j-1]
				wordlist[j-1] = wordlist[j]
				wordlist[j] = tmp
			}

		}
	}
	newWordlist := make(map[string]int)

	var result []word
	for t := range wordlist {
		if _, ok := newWordlist[wordlist[t].chinese]; ok {
			continue
		} else {
			newWordlist[wordlist[t].chinese] = 0
			result = append(result, wordlist[t])
		}
	}
	return result
}

// chineseList 转换成返回的汉字列表
func chineseList(wordlist []word) []string {
	chinese := make([]string, 0, 50)
	i := 0
	for _, tmp := range wordlist {
		chinese = append(chinese, tmp.chinese)
		i++
	}

	return chinese
}

// isPrime 当且仅当word1优先于word2 返回true
func isPrime(word1 word, word2 word) bool {
	if word1.frequency > word2.frequency {
		return true
	}
	if word1.frequency == word2.frequency {
		if strings.Compare(word1.spell, word2.spell) == -1 {
			return true
		}
	}
	return false
}

// getSpell 由输入目录获得文件代表的拼写
func getSpell(input string) string {
	var spell string
	r1 := filepath.Base(input)
	r, _ := regexp.Compile("([a-z]+).dat")
	if r.FindStringSubmatch(r1) != nil {
		spell = r.FindStringSubmatch(r1)[1]
	}

	return spell
}

// isChineseChar 判断是否汉字
func isChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u4e00-\u9fa5]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

// isNumRange 判断汉字评分是否符合标准
func isNumRange(str string) int {
	var num int
	var err error
	if num, err = strconv.Atoi(str); err != nil {
		return 0
	}
	if num > 10 || num < 1 {
		return 0

	}
	return num
}