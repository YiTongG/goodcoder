package tree

import "fmt"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children []element
}

type element struct {
	spell   string
	isEnd bool
	next  *TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: []element{},
	}
}

func (n *TrieNode) Push(v string, isEnd bool) *TrieNode {
	newNode := newTrieNode()
	n.children = append(n.children, element{
		spell:   v,
		isEnd: isEnd,
		next:  newNode,
	})
	return newNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	chars := []rune(word)
	curr := t.root
	for k, v := range chars {
		index := curr.searchChar(string(v))
		if index != -1 {
			if k == len(chars)-1 {
				curr.children[index].isEnd = true
			}
			curr = curr.children[index].next
			continue
		}
		// 插入
		if k == len(chars)-1 {
			curr = curr.Push(string(v), true)
		} else {
			curr = curr.Push(string(v), false)
		}
	}

}

func (tn *TrieNode) searchChar(c string) (index int) {
	for k, n := range tn.children {

		if n.spell == c {

			return k
		}

	}

	return -1
}
func (tn *TrieNode) searchLeafChar(c string) (children string) {
	for _, n := range tn.children {

		if n.spell == c {

			return tn.searchLeafChar(c)
		}

	}

	return ""
}

/** Returns if the word is in the trie. */
func (t *Trie) IsExist(word string) bool {
	chars := []rune(word)
	curr := t.root
	for k, v := range chars {
		index := curr.searchChar(string(v))
		if index == -1 {
			break
		}
		if k == len(chars)-1 {
			return curr.children[index].isEnd
		}
		curr = curr.children[index].next
	}

	return false
}

/** 校验输入是否合法 */
func (t *Trie) StartsWith(prefix string) bool {
	chars := []rune(prefix)
	curr := t.root
	for k, v := range chars {
		index := curr.searchChar(string(v))
		if index == -1 {
			break
		}
		if k == len(chars)-1 {
			return true
		}
		curr = curr.children[index].next
	}

	return false
}

func (t *Trie) PrefixList(prefix string) []string {
	var filelist  []string
	fmt.Println(prefix)
	chars := []rune(prefix)
	curr := t.root
	//var  char []rune
	//k := len(chars)-1
	//var v rune
	for k:=0 ;k<=5;k++{
		index := curr.searchChar(string(chars[k]))
		if index == -1 {
			break
		}
		////filelist = append(filelist,string(char))
		if k >= len(chars)-1 {
			//	filelist = append(filelist,string(char))
			//	fmt.Println(filelist)
			filelist = append(filelist,string(chars))
			chars = append(chars,chars[k])
		}
		curr = curr.children[index].next

	}
	//for k, v := range chars {
	//
	//	index := curr.searchChar(string(v))
	//	if index == -1 {
	//		break
	//	}
	//	////filelist = append(filelist,string(char))
	//	if k >= len(chars)-1 {
	//	//	filelist = append(filelist,string(char))
	//	//	fmt.Println(filelist)
	//	filelist = append(filelist,string(chars))
	//	chars = append(chars,v)
	//	}
	//	curr = curr.children[index].next
	//
	//}
	fmt.Println(filelist)
	return filelist
}