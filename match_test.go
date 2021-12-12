package search

import "testing"

func TestSourceStore_MatchFullSpell(t *testing.T) {
	match := New()

	match.Store("123", [][]string{
		{"chang", "zhang"},
		{"fa"},
	})

	t.Log(match.MatchFullSpell("changfa"))
}
