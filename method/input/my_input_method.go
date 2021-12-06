package input

import (
	"bufio"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/file"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/log"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/tree"
	"io"
	"os"
	"regexp"
)

func BuildSpellTree(filename string) tree.Trie {
	spellTree,err := tree.InitTree(filename)
	if err != nil {
		log.Err("init failed Error:%v\n",err)
	}
	return spellTree
}


//建立一个包含所有拼音的字典树，用于快速校验输入是否合法
var pinyin = "./pinyin_spell.txt"
var spellTree = BuildSpellTree(pinyin)

func FindWords(spell string) []string {
	if !spellTree.StartsWith(spell) {
		return nil
	}
	wordlist := readFromCache(spell)
	if wordlist == nil {
		wordlist = fetchWordList(spell)
	}
	return file.ChineseList(wordlist)
}

func readFromCache(spell string) []file.Word {
	filename, err := file.FileName("./history")
	if err != nil {
		log.Err("")
		return nil
	}
	tree := tree.InitWordTree(filename)
	if tree.IsExist(spell) {
		return file.ReadWordList(spell)
	}
	return nil
}
func fetchWordList (spell string) []file.Word{
	var filelist []string
	filename, err := file.FileName("./dict")
	if err != nil {
		log.Err("")
		return nil
	}

	r, _ := regexp.Compile("dict/([a-z]+).dat")
	var tmpfilename []string
	for _, tmp := range filename {
		tmpfilename = r.FindStringSubmatch(tmp)
		if tmpfilename != nil {
			if spellTree.IsExist(tmpfilename[1]) {
				filelist = append(filelist, tmpfilename[1])
			}
		}
	}
	var wordlist []file.Word
	filelist = spellTree.PrefixList(spell)
	for _, tmp := range filelist {
		wordlist = append(wordlist, file.ReadWordList(tmp)...)
	}
	//tree := tree.InitWordTree(filelist)

	sortword := file.SortWordList(wordlist)
	writeCache(sortword,spell)
	return sortword
}



func writeCache(wordlist []file.Word, spell string) {
	for _, tmp := range wordlist {
		path := "./history/" + spell + ".dat"
		f, err := os.Create(path)
		if err != nil {

		}
		f, err = os.Open(path)
		if err != nil {

		}
		defer f.Close()
		rd := bufio.NewWriter(f)
		for {
			_, err := rd.WriteString(tmp.Chinese + " ")
			if err == io.EOF {
				break
			}
			_, err = rd.WriteRune(tmp.Frequency)
			if err == io.EOF {
				break
			}
			_, err = rd.WriteString("\n")
			if err == io.EOF {
				break
			}

		}
	}
}
