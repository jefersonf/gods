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
			hasFound := trie.SearchWord(tc.word)
			if hasFound != tc.searchResult {
				t.Errorf("Got %t, want %t", hasFound, tc.searchResult)
			}
		})
	}
}
