package method

import (
	"github.com/stretchr/testify/assert"
	"icode.baidu.com/baidu/goodcoder/gongyitong/log"
	"reflect"
	"testing"
)

// SpellTree 包含全部拼音的树，用于校验

// TestInputMethod 测试创建词典
func TestNewInputMethod(t *testing.T) {
	log.SetEnv(log.EnvDevbox)
	log.SetDebugMode(true)
	args := []string{"./your-goodcoder-code", "../testdata/zhan.dat", "../testdata/zhang.dat", "http://10.138.61.200:8001/orp/de.dat"}
	im := NewInputMethod(args[1:])
	testcases := []struct {
		Name      string
		spell     string
		wants     []string
		spellTree Trie
	}{
		{
			Name:  "normal_case1",
			spell: "z",
			wants: []string{"粘", "谵", "长", "展", "詹", "张", "瞻", "盏", "沾", "战"},
		},
		{
			Name:  "normal_case2",
			spell: "zhan",
			wants: []string{"粘", "谵", "展", "詹", "瞻", "盏", "沾", "战", "站", "斩", "颤", "崭", "旃"},
		},
		{
			Name:  "normal_case3",
			spell: "de",
			wants: []string{"的", "得", "地", "德"},
		},
		{
			Name:  "spell_error",
			spell: "diu",
			wants: []string(nil),
		},
	}
	for _, tt := range testcases {
		t.Run(tt.Name, func(t *testing.T) {
			res := im.FindWords(tt.spell)
			assert.Equal(t, res, tt.wants)

		})
	}
}
func TestFindWords(t *testing.T) {
	log.SetEnv(log.EnvDevbox)
	log.SetDebugMode(true)
	args := []string{"./your-goodcoder-code",
		"../testdata/zha.dat",
		"../testdata/zhan.dat",
		"../testdata/abc.dat",
		"../testdata/zhang.dat",
		"http://10.138.61.200:8001/orp/de.dat",
		"http://10.138.61.200:8001/orp/da.dat"}

	im := NewInputMethod(args[1:])

	// 测试样例
	testExamples := []struct {
		spell string
		res   []string
	}{
		{
			spell: "z",
			res:   []string{"粘", "谵", "长", "展", "詹", "张", "瞻", "盏", "沾", "战"},
		},

		{
			spell: "zhan",
			res:   []string{"粘", "谵", "展", "詹", "瞻", "盏", "沾", "战", "站", "斩", "颤", "崭", "旃"},
		},
		{
			spell: "dui",
			res:   []string(nil),
		},
		{
			spell: "de",
			res:   []string{"的", "得", "地", "德"},
		},
		{
			spell: "abbbbb",
			res:   []string(nil),
		},
	}
	for _, tt := range testExamples {
		res := im.FindWords(tt.spell)
		if !reflect.DeepEqual(res, tt.res) {
			t.Errorf("FindWords verify fail: expect=[%+v] real=[%+v]", tt.res, res)
		}
	}
}