package tree

import (
	"bufio"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/log"
	"io"
	"os"
)

func InitTree(filename string)(Trie,error) {
	Tree := Constructor()
	spell, err := os.Open(filename)
	if err != nil {
		log.Err("Open spellfile failed.Error: %v\n", err)
	}
	defer spell.Close()

	br := bufio.NewReader(spell)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		Tree.Insert(string(a))
	}
    return Tree,nil
}
func InitWordTree(spell []string)(Trie) {
	Tree := Constructor()
	for _, tmp := range spell {
		Tree.Insert(string(tmp))
	}
	return Tree
}