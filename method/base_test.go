package method

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	word1 := &word{
		chinese:   "站",
		frequency: 10,
		spell:     "zhan",
	}
	word2 := &word{
		chinese:   "占",
		frequency: 8,
		spell:     "zhan",
	}
	word3 := &word{
		chinese:   "张",
		frequency: 10,
		spell:     "zhang",
	}

	word4 := &word{
		chinese:   "长",
		frequency: 8,
		spell:     "zhang",
	}

	flag := isPrime(*word1, *word2)
	assert.Equal(t, flag, true)

	flag = isPrime(*word1, *word3)
	assert.Equal(t, flag, true)

	flag = isPrime(*word4, *word3)
	assert.Equal(t, flag, false)

}
func TestChineseList(t *testing.T) {
	wordlist := []word{
		{
			chinese:   "站",
			frequency: 10,
			spell:     "zhan",
		},
		{
			chinese:   "占",
			frequency: 8,
			spell:     "zhan",
		}, {
			chinese:   "张",
			frequency: 10,
			spell:     "zhang",
		}, {
			chinese:   "长",
			frequency: 8,
			spell:     "zhang",
		},
	}
	chineseList(wordlist)

}
func TestIschineseChar(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  bool
	}{{
		Name:  "yes",
		Input: "汉",
		Want:  true,
	}, {
		Name:  "no",
		Input: "eng",
		Want:  false,
	},
	}
	for _, tt := range testcases {
		flag := isChineseChar(tt.Input)
		if !reflect.DeepEqual(flag, tt.Want) {
			t.Errorf("ischineseChar verify fail: expect=[%+v] real=[%+v]", flag, tt.Want)
		}
	}
}
func TestIsNumRange(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  int
	}{{
		Name:  "yes",
		Input: "1",
		Want:  1,
	}, {
		Name:  "no",
		Input: "11",
		Want:  0,
	},
		{
			Name:  "no",
			Input: "5.6",
			Want:  0,
		},
	}
	for _, tt := range testcases {
		flag := isNumRange(tt.Input)
		if !reflect.DeepEqual(flag, tt.Want) {
			t.Errorf("IsNumRange verify fail: expect=[%+v] real=[%+v]", flag, tt.Want)
		}
	}
}
func TestGetspell(t *testing.T) {
	testcases := []struct {
		Name  string
		Input string
		Want  string
	}{{
		Name:  "yes",
		Input: "../testdata/zhan.dat",
		Want:  "zhan",
	}, {
		Name:  "no",
		Input: "../testdata/zhan.da",
		Want:  "",
	},
	}
	for _, tt := range testcases {
		flag := getSpell(tt.Input)
		if !reflect.DeepEqual(flag, tt.Want) {
			t.Errorf("getspell verify fail: expect=[%+v] real=[%+v]", flag, tt.Want)
		}
	}
}