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
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/input"
	"icode.baidu.com/baidu/goodcoder/gongyitong/constants"
	"icode.baidu.com/baidu/goodcoder/gongyitong/method/log"

	"os"
	//"path/filepath"
	"strings"
)

func main() {
	// 若实现加载指定目录下全部词典，并判断词典名是否合法
	log.SetEnv(constants.ENV_DEVBOX)
	log.SetDebugMode(true)
	stdin := bufio.NewReader(os.Stdin)
	for {
		spell, err := stdin.ReadString('\n')
		if err != nil {
			break
		}
		spell = strings.TrimRight(spell, "\n")
		words := input.FindWords(spell)
		fmt.Println(strings.Join(words, ", "))
	}
}
