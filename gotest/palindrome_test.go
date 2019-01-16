package gotest

import "testing"


func TestPalindrome(t *testing.T) {
	if !IsPalindrome("636") {
		t.Error(`IsPalindrome("TongTu") = false`)
	}
	if !IsPalindrome("636") {
		t.Error(`IsPalindrome("6361") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("TongTu") {
		t.Error(`IsPalindrome("363") = true`)
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 1", args{"121"}, true},
		{"test 2", args{"1213"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.text); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
