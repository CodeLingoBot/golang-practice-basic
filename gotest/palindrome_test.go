package gotest

import "testing"

func TestPalindrome(t *testing.T){
	if !IsPalindrome("636"){
		t.Error(`IsPalindrome("TongTu") = false`)
	}
	if !IsPalindrome("636"){
		t.Error(`IsPalindrome("6361") = false`)
	}
}

func TestNonPalindrome(t *testing.T){
	if IsPalindrome("TongTu"){
		t.Error(`IsPalindrome("363") = true`)
	}
}