// Copyright 2021 Baidu Inc. All rights reserved.
// Use of this source code is governed by a xxx
// license that can be found in the LICENSE file.

/*
modification history
--------------------
2021/09/24 16:25:47, by gongyitong@baidu.com, create
*/


/*
gongyitong@baidu.com, create
*/

package main

import (
	"bufio"
	"fmt"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method"
	"os"
	"strings"
)

// SpellTree 包含全部拼音的树，用于校验

// loop 循环读取输入拼音
func loop(im method.MyInputMethod) {
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

// main 实现加载指定目录下全部词典，并判断词典名是否合法
func main() {
	//log.SetEnv(constants.EnvDevbox)
	//log.SetDebugMode(true)
	im := method.NewInputMethod(os.Args[1:])
	//args := []string{"./your-goodcoder-code", "./dict/zhan.dat", "./dict/zhang.dat", "./dict/zha.dat"}
	//im := method.NewInputMethod(args[1:], SpellTree)

	loop(im)

}