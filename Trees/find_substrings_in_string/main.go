package main

import (
	"log/slog"
	"os"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode), // Initialize the children map
		isEnd:    false,
	}
}

func (t *TrieNode) Insert(char rune) {
	t.children[char] = NewTrieNode()
}

func (t *TrieNode) Has(char rune) bool {
	_, ok := t.children[char]
	return ok
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if !node.Has(char) {
			node.Insert(char)
		}
		node = node.children[char]
	}

	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if !node.Has(char) {
			return false
		}
		node = node.children[char]
	}

	return node.isEnd
}

func main() {
	args := os.Args[1:]
	trie := NewTrie()

	if len(args) < 2 {
		slog.Error("usage: go run main.go <word> <word1> <word2> ... <wordN>")
		os.Exit(1)
	}

	for _, word := range args[1:] {
		trie.Insert(word)
		slog.Info("word inserted: ", "word", word)
	}

	stringsFound := make([]string, 0)
	word := args[0]

	for i := 0; i < len(word); i++ {
		node := trie.root
		for j := i; j < len(word); j++ {
			char := rune(word[j])
			if !node.Has(char) {
				break
			}
			node = node.children[char]

			if node.isEnd {
				stringsFound = append(stringsFound, word[i:j+1])
			}
		}
	}

	println("Printing all the strings found in the trie:")
	for _, str := range stringsFound {
		println("\t- " + str)
	}

}
