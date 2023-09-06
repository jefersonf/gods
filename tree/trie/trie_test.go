package trie

import (
	"fmt"
	"testing"
)

func TestHeight(t *testing.T) {

	testcases := []struct {
		stringSet  []string
		trieHeight int
	}{
		{
			stringSet:  []string{},
			trieHeight: 0,
		},
		{
			stringSet:  []string{"a"},
			trieHeight: 1,
		},
		{
			stringSet:  []string{"foo"},
			trieHeight: 3,
		},
		{
			stringSet:  []string{"Pot", "Bowl"},
			trieHeight: 4,
		},
		{
			stringSet:  []string{"Paper and Crayons", "Notebook"},
			trieHeight: 15,
		},
		{
			stringSet:  []string{"Once upon a time in South America", "A good story"},
			trieHeight: 27,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("StringSet_%d", i), func(t *testing.T) {
			trie := New()
			for _, s := range tc.stringSet {
				trie.AddString(s)
			}
			trieHeight := trie.Height()
			if trieHeight != tc.trieHeight {
				t.Errorf("Incorrect trie size; got %d, want %d", trieHeight, tc.trieHeight)
			}
		})
	}
}

func TestSearchWord(t *testing.T) {

	testcases := []struct {
		wordSet      []string
		word         string
		searchResult bool
	}{
		{
			wordSet:      []string{},
			word:         "",
			searchResult: true,
		},
		{
			wordSet:      []string{""},
			word:         "",
			searchResult: true,
		},
		{
			wordSet:      []string{""},
			word:         "a",
			searchResult: false,
		},
		{
			wordSet:      []string{"a", "z"},
			word:         "",
			searchResult: true,
		},
		{
			wordSet:      []string{"a", "z"},
			word:         "a",
			searchResult: true,
		},
		{
			wordSet:      []string{"a", "z"},
			word:         "b",
			searchResult: false,
		},
		{
			wordSet:      []string{"apple", "banana", "grapes"},
			word:         "banana",
			searchResult: true,
		},
		{
			wordSet:      []string{"apple", "banana", "grapes"},
			word:         "strawberry",
			searchResult: false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("WordSet %d", i), func(t *testing.T) {
			trie := New()
			for _, w := range tc.wordSet {
				trie.AddString(w)
			}
			hasFound := trie.Search(tc.word)
			if hasFound != tc.searchResult {
				t.Errorf("Got %t, want %t", hasFound, tc.searchResult)
			}
		})
	}
}

func TestCountWord(t *testing.T) {

	testcases := []struct {
		wordSet     []string
		word        string
		countResult int
	}{
		{
			wordSet:     []string{},
			word:        "",
			countResult: 0,
		},
		{
			wordSet:     []string{""},
			word:        "",
			countResult: 0,
		},
		{
			wordSet:     []string{""},
			word:        "a",
			countResult: 0,
		},
		{
			wordSet:     []string{"a", "z"},
			word:        "",
			countResult: 0,
		},
		{
			wordSet:     []string{"a", "z"},
			word:        "a",
			countResult: 1,
		},
		{
			wordSet:     []string{"a", "z"},
			word:        "b",
			countResult: 0,
		},
		{
			wordSet:     []string{"apple", "banana", "grapes"},
			word:        "banana",
			countResult: 1,
		},
		{
			wordSet:     []string{"apple", "banana", "grapes"},
			word:        "strawberry",
			countResult: 0,
		},
		{
			wordSet:     []string{"apple", "banana", "grapes", "banana", "banana"},
			word:        "banana",
			countResult: 3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("WordSet %d", i), func(t *testing.T) {
			trie := New()
			for _, w := range tc.wordSet {
				trie.AddString(w)
			}
			wordCount := trie.CountWord(tc.word)
			if wordCount != tc.countResult {
				t.Errorf("Got %d, want %d", wordCount, tc.countResult)
			}
		})
	}
}

func TestWordCount(t *testing.T) {

	testcases := []struct {
		wordSet   []string
		wordCount int
	}{
		{
			wordSet:   []string{},
			wordCount: 0,
		},
		{
			wordSet:   []string{""},
			wordCount: 0,
		},
		{
			wordSet:   []string{"z"},
			wordCount: 1,
		},
		{
			wordSet:   []string{"f", "o", "o"},
			wordCount: 2,
		},
		{
			wordSet:   []string{"apple", "apples"},
			wordCount: 2,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("WordSet %d", i), func(t *testing.T) {
			trie := New()
			for _, w := range tc.wordSet {
				trie.AddString(w)
			}
			wordCount := trie.WordCount()
			if wordCount != tc.wordCount {
				t.Errorf("Got %d, want %d", wordCount, tc.wordCount)
			}
		})
	}
}
