
package main 

import "testing"

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


