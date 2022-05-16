package modules

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripPunctuationOK(t *testing.T) {
	input := "a,  b:,, cc... ddd - eeeee"
	output := "a  b cc ddd  eeeee"

	tOutput := stripPunctuation(input)
	require.Equal(t, tOutput, output)
}
