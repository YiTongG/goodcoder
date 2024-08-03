package method

import (
	"icode.baidu.com/baidu/goodcoder/gongyitong/log"
	"regexp"
	"strings"
)

// NewMyInputMethod 一个输入法接口
type NewMyInputMethod interface {
	FindWords(string) []string
}

// MyInputMethod 建立字典用到的资源
type MyInputMethod struct {
	trie *Trie                // 输入dict资源建立的树
	dict map[string]*resource // 拼音与资源关系的映射
}

// NewInputMethod 新建一个输入法
func NewInputMethod(dicts []string) MyInputMethod {
	fileTree := newTrie()
	var spell string
	m := MyInputMethod{}
	m.dict = make(map[string]*resource)

	for _, row := range dicts {
		spell = getSpell(row)
		r, err := regexp.Compile("([a-z]+)")
		if err != nil {
			log.Err("invalid resource")
			continue
		}
		if r.FindStringSubmatch(spell) != nil {
			spell = r.FindStringSubmatch(spell)[1]
		}
		m.dict[spell] = &resource{path: row}
		fileTree.Insert(spell)
		m.trie = &fileTree

	}
	return m
}

// FindWords 在字典中找相应字符
func (mim *MyInputMethod) FindWords(spell string) []string {

	var wordlist []word
	if !mim.trie.StartsWith(spell) {
		log.Warning("invalid method")
		return nil
	}

	if mim.trie.IsExist(spell) {
		wordlist = mim.readWordList(spell)
	} else if mim.trie.StartsWith(spell) {
		filelist := mim.trie.PrefixList(spell)
		for _, tmp := range filelist {
			wordlist = append(wordlist, mim.readWordList(tmp)...)
		}
	}
	wordlist = sortWordList(wordlist)
	list := chineseList(wordlist)
	if !mim.trie.IsExist(spell) {
		if len(list) > 10 {
			list = list[:10]
		}
	}
	return list
}

// readWordList 从http/https 和 filesystem读文件资源
func (mim *MyInputMethod) readWordList(spell string) []word {
	path := mim.dict[spell].path
	var wordlist []word
	r := &resource{}
	r.path = path

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resource := &httpResource{r}
		wordlist = resource.GetResource()

	} else {

		resource := &fsResource{r}
		wordlist = resource.GetResource()
	}

	return wordlist
}