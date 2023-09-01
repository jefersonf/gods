package trie

import (
	"fmt"
	"testing"
)

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
