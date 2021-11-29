// Copyright 2021 Baidu Inc. All rights reserved.
// Use of this source code is governed by a xxx
// license that can be found in the LICENSE file.

/*
modification history
--------------------
2021/09/24 16:25:47, by gongyitong@baidu.com, create
*/


package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"path/filepath"
	"strings"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method"
)

type InputMethod interface {
	FindWords(string) []string
}

func loop(im InputMethod) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		spell, err := stdin.ReadString('\n')
		if err != nil {
			break
		}
		spell = strings.TrimRight(spell, "\n")
		words := im.FindWords(spell)
		fmt.Println(strings.Join(words, ", "))
	}
}
//建立全部合法拼音的字典树，用于判断词典名、拼音名是否合法
func setup() method.TrieTree{
	var spell_list []string
	spell, err := os.Open("./pinyin_spell.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		}
	defer spell.Close()

	br := bufio.NewReader(spell)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		spell_list = append(spell_list,string(a))

	}
	spellTree := method.CreateTree(spell_list)
	return spellTree
}
	//var dictpath = "./dict"
	//var dict_list []string
	//files,_ := filepath.Glob(dictpath+"/*")
	//检查是否拼音+，.dat格式

	//fmt.Println(files) // contains a list of all files in the current directory


func main() {
	// 若实现加载指定目录下全部词典，并判断词典名是否合法
    setup()


}