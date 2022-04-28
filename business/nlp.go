package business

import (
	"regexp"
	"sort"
	"strings"

	"github.com/liderman/rustemmer"
)

func countWords(words []string) map[string]int {
	count := make(map[string]int)
	for _, w := range words {
		if count[w] == 0 {
			count[w] = 1
		} else {
			count[w]++
		}
	}

	return count
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func stem(word string) string {
	return rustemmer.GetWordBase(word)
}

func splitWords(s string) []string {
	return strings.Split(s, " ")
}

func stripPunctuation(s string) string {
	r := regexp.MustCompile(`[()\[\].,\-"':;«»—!?]`)
	return r.ReplaceAllLiteralString(s, "")
}
