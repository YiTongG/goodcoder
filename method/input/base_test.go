package input

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPrime(t *testing.T) {
	word1 := Word{
		Chinese : "站" ,
		Frequency : 10,
		Spell : "zhan",
	}
	word2 := Word{
		Chinese : "占" ,
		Frequency : 8,
		Spell : "zhan",
	}
	word3 := Word{
		Chinese : "张" ,
		Frequency : 10,
		Spell : "zhang",
	}

	word4 := Word{
		Chinese : "长" ,
		Frequency : 8,
		Spell : "zhang",
	}

	flag := isPrime(word1,word2)
	assert.Equal(t, flag, true)

	flag = isPrime(word1,word3)
	assert.Equal(t, flag, true)

	flag = isPrime(word4,word3)
	assert.Equal(t, flag, false)

}