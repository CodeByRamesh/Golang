package main 

import "testing"

func IsParlindrome( s string ) bool {
	for i := range s {
		if s[i] != s[ len(s) - 1 - i ] {
			return false
		}
	}
	return true
}

func TestPalindrome( t *testing.T ) {
	if !IsParlindrome("detartrated") {
		t.Error(`IsParlindrome("detartrated") = false`)
	}

	if !IsParlindrome("kayak") {
		t.Error(`IsParlindrome("kayak") = false`)
	}
}

func TestNotPalindrome( t *testing.T ) {
	if IsParlindrome("Human") {
		t.Error(`IsParlindrome("Human") = true`)
	}
}
