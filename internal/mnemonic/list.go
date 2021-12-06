package mnemonic

import "github.com/tyler-smith/go-bip39/wordlists"

var (
	wordList []string
	wordMap  map[string]int
)

func init() {
	SetWordList(wordlists.English)
}

func SetWordList(list []string) {
	wordList = list

	wordMap = make(map[string]int)

	for i, v := range wordList {
		wordMap[v] = i
	}
}
