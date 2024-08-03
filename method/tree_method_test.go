package method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrieTree(t *testing.T) {
	tree := newTrie()
	tree.Insert("zhang")
	tree.Insert("zhan")
	tree.Insert("zha")

	tree.Insert("chang")
	tree.Insert("chu")
	tree.Insert("cheng")
	res := []string{"zha", "zhan", "zhang"}
	flag := tree.PrefixList("z")
	assert.Equal(t, flag, res)
	res = []string{"chang", "chu", "cheng"}
	flag = tree.PrefixList("c")
	flag1 := tree.IsExist("zhan")
	assert.Equal(t, flag1, true)
	flag2 := tree.IsExist("zzzzzzz")
	assert.Equal(t, flag2, false)

}