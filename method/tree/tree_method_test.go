package tree

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	tree := Constructor()
	tree.Insert("zhang")
	tree.Insert("zhan")
	tree.Insert("zha")

	flag := tree.PrefixList("z")
	fmt.Println(flag)


}