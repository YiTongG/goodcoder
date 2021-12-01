package method

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func setup() (TrieTree,error){
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
			return err
		}
		spell_list = append(spell_list,string(a))

	}
	spellTree := CreateTree(spell_list)
	return spellTree
}