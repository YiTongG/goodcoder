package file

import (
	"bufio"
	"fmt"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/log"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)
type  Word struct{
	Chinese  string
	Frequency  rune
	Spell string
}


func FileName(dictpath string) ([]string, error) {
	//
	file, err := filepath.Glob(dictpath + "/*")
	if err != nil {
		log.Err("Read filelist failed.Error:%v",err)
		return nil, err
	}

	//
	return file, nil
}
func ReadWordList(spell string) []Word  {
	var wordlist []Word
    path := "./"+spell+".dat"
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		arr := strings.Split(line, " ")
		if len(arr) != 2 {
			continue
		}
		chinese := arr[0]

		frequency, err1 := strconv.Atoi(arr[1])
		if err1 != nil {
			continue
		}
		item := Word{
			Spell:     spell,
			Chinese:   chinese,
			Frequency: rune(frequency),
		}
		wordlist = append(wordlist, item)
	}
	return wordlist
}
func SortWordList(wordlist []Word) []Word{
	//t := reflect.TypeOf(wordlist)
	fmt.Println(wordlist)
	v := reflect.ValueOf(wordlist[0])

	for k := 0; k < v.NumField(); k++ {
		for j := k+1 ;j< v.NumField()-1 ;j++{
			if wordlist[k].Frequency<wordlist[j].Frequency {
				tmp := wordlist[k]
				wordlist[k] = wordlist[j]
				wordlist[j] = tmp
			}
		}
	}
	return wordlist
}
func ChineseList(wordlist []Word) []string {
	var chinese []string
	i := 0
	for _,tmp := range wordlist {
		chinese[i] = tmp.Chinese
		i++
	}
	if len(chinese)>10 {
		return chinese[:10]
	}
	return chinese
}
