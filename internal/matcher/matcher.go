package matcher

import (
	"fmt"
	"sort"
	"strings"
)

type node struct {
	// map with children
	children map[rune]*node

	// defines word has finished
	isEnd bool
}

func newNode() *node {
	return &node{
		children: make(map[rune]*node),
		isEnd:    false,
	}
}

type Matcher interface {

	// append new word to tree
	// returns occured error
	Insert(str string) error

	// appends entire collection
	// returns occured error
	InsertAll(lst []string) error

	// returns matched words for input or error occured during operation
	Match(prefix string) []string

	// Fixes input
	Suggestions(input string, maxDistance int) []string

	// dumps words or letters
	dump(node *node, prefix string) []string
	normalize(s string) string
}

type matcherImpl struct {
	root *node
}

func NewMatcher() Matcher {
	return &matcherImpl{
		root: newNode(),
	}
}

// normalize string by trimming spaces and lowered
func (m *matcherImpl) normalize(s string) string {
	lowered := strings.ToLower(s)
	removeSpaces := strings.TrimSpace(lowered)
	return removeSpaces
}

func (m *matcherImpl) dump(node *node, prefix string) []string {
	var words []string = []string{}

	if node.isEnd {
		words = append(words, prefix)
	}

	for k, v := range node.children {
		words = append(words, m.dump(v, fmt.Sprint(prefix, string(k)))...)
	}

	return words
}

func (m *matcherImpl) Match(prefix string) []string {
	if prefix == "" {
		return []string{}
	}

	prefix = m.normalize(prefix)
	var node = m.root

	for _, key := range prefix {
		if _, ok := node.children[key]; !ok {
			return []string{}
		}
		node = node.children[key]
	}

	words := m.dump(node, prefix)

	return words
}

func (m *matcherImpl) Insert(str string) error {
	if m.root == nil {
		return fmt.Errorf("There is no root")
	}

	if str == "" {
		return fmt.Errorf("String is empty")
	}

	str = m.normalize(str)

	var node = m.root

	for _, key := range str {
		if _, ok := node.children[key]; !ok {
			node.children[key] = newNode()
		}
		node = node.children[key]
	}

	node.isEnd = true

	return nil
}

func (m *matcherImpl) InsertAll(lst []string) error {
	if m.root == nil {
		return fmt.Errorf("InsertAll failed: root is nil")
	}

	for _, word := range lst {
		if err := m.Insert(word); err != nil {
			return fmt.Errorf("InsertAll failed for word %s: %v", word, err)
		}
	}
	return nil
}

func (m *matcherImpl) Suggestions(input string, maxDistance int) []string {
	type suggestion struct {
		word     string
		distance int
	}

	var results []suggestion = []suggestion{}

	words := m.dump(m.root, "")
	for _, word := range words {
		distance := Dist(input, word)
		if distance <= maxDistance {
			results = append(results, suggestion{word: word, distance: distance})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].distance < results[j].distance
	})

	suggestions := make([]string, len(results))
	for i, result := range results {
		suggestions[i] = result.word
	}

	return suggestions
}
