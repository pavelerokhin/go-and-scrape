package business

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitWordsOK(t *testing.T) {
	input1 := "a b cc ddd eeeee"
	input2 := "abcd"
	output1 := []string{"a", "b", "cc", "ddd", "eeeee"}
	output2 := []string{"abcd"}

	tOutput1 := splitWords(input1)
	tOutput2 := splitWords(input2)

	require.Equal(t, tOutput1, output1)
	require.Equal(t, tOutput2, output2)

}

func TestStripPunctuationOK(t *testing.T) {
	input := "a,  b:,, cc... ddd - eeeee"
	output := "a  b cc ddd  eeeee"

	tOutput := stripPunctuation(input)
	require.Equal(t, tOutput, output)
}

func TestCountWordsOK(t *testing.T) {
	input := []string{"a", "a", "b", "c", "d", "b", "e", "b", "e", "d", "e", "e"}
	output := map[string]int{
		"a": 2,
		"b": 3,
		"c": 1,
		"d": 2,
		"e": 4,
	}

	tOutput := countWords(input)
	require.Equal(t, tOutput, output)
}
