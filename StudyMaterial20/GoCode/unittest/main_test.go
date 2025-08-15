
package main 

import "testing"

func TestPalindrome( t *testing.T ) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsParlindrome("detartrated") = false`)
	}

	if !IsPalindrome("kayak") {
		t.Error(`IsParlindrome("kayak") = false`)
	}
}

func TestNotPalindrome( t *testing.T ) {
	if IsPalindrome("Human") {
		t.Error(`IsParlindrome("Human") = true`)
	}
}


