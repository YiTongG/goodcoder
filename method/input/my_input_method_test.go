package input

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	tree2 "icode.baidu.com/baidu/goodcoder/gongyitong/method/tree"
	"testing"
)

func TestTrieTree(t *testing.T) {
	tree := tree2.Constructor()
	tree.Insert("apple")
	tree.Insert("pear")
	flag := tree.IsExist("apple")
	fmt.Println(flag)
	assert.Equal(t, flag, true)
	flag = tree.IsExist("pear")
	fmt.Println(flag)
	assert.Equal(t, flag, true)

}
