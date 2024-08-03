package method

// Trie 字典树
type Trie struct {
	root *TrieNode
}

// TrieNode 树节点
type TrieNode struct {
	spell    byte
	isEnd    bool
	children map[byte]*TrieNode
}

// newTrieNode 新建节点
func newTrieNode(c byte) *TrieNode {

	return &TrieNode{c, false, make(map[byte]*TrieNode)}

}

// newTrie 初始化一棵字典树
func newTrie() Trie {

	return Trie{newTrieNode('/')}

}

// Insert 在字典树中加入一个拼音
func (t *Trie) Insert(word string) {
	curr := t.root

	for _, v := range []byte(word) {
		if _, ok := curr.children[v]; !ok { // 不存在
			curr.children[v] = newTrieNode(v)
		}

		curr = curr.children[v]
	}

	curr.isEnd = true

}

// PrefixList 获取树中包含该前缀的所有单词
func (t *Trie) PrefixList(prefix string) []string {
	res := make([]string, 0)

	curr := t.root

	for _, v := range []byte(prefix) {
		if _, ok := curr.children[v]; !ok { // 不存在
			return res
		}

		curr = curr.children[v]
	}

	// startWithList 前缀全部匹配单词
	var startWithList func(prefix string, curr *TrieNode)
	startWithList = func(prefix string, curr *TrieNode) {
		if curr.isEnd {
			res = append(res, prefix)
		}

		for _, v := range curr.children {
			startWithList(prefix+string(v.spell), v)
		}
	}
	startWithList(prefix, curr)
	return res

}

// IsExist 判断字典树中是否有该单词
func (t *Trie) IsExist(word string) bool {
	curr := t.root

	for _, v := range []byte(word) { // 不存在
		if _, ok := curr.children[v]; !ok {
			return false
		}

		curr = curr.children[v]
	}

	return curr.isEnd

}

// StartsWith 判断树中是否有该前缀，用于校验输入是否合法
func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root

	for _, v := range []byte(prefix) {
		if _, ok := curr.children[v]; !ok { // 不存在
			return false
		}

		curr = curr.children[v]
	}

	return true

}