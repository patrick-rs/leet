package test

import (
	"testing"

	"github.com/patrick-rs/leet/pkg/palindrom"
)

func TestStringWithOnePalindrom(t *testing.T) {
	pal := "abcd"
	expectedResult := false
	result := palindrom.IsPalindrome(pal)

	if result == !expectedResult {
		t.Errorf("Expected %v but got %v when passing '%s'", expectedResult, result, pal)
	}
}

func TestSingleChar(t *testing.T) {
	pal := "a"
	expectedResult := true
	result := palindrom.IsPalindrome(pal)

	if result == !expectedResult {
		t.Errorf("Expected %v but got %v when passing '%s'", expectedResult, result, pal)
	}
}

func TestLongPalindrom(t *testing.T) {
	pal := "abccba"
	expectedResult := true
	result := palindrom.IsPalindrome(pal)

	if result == !expectedResult {
		t.Errorf("Expected %v but got %v when passing '%s'", expectedResult, result, pal)
	}
}

func TestLongestPalindrom(t *testing.T) {
	str := "abacdabcddcba"
	expectedResult := "abcddcba"
	result := palindrom.LongestPalindrome(str)

	if result != expectedResult {
		t.Errorf("Expected %s but got %s when passing '%s'", expectedResult, result, str)
	}
}

func TestLongestPalindromWithNoLongPalindrom(t *testing.T) {
	str := "asdfbsdiadfg"
	result := palindrom.LongestPalindrome(str)
	// every string contains a palindrom because a single character is a palindrom
	expectedLengthOfResult := 1

	if len(result) != expectedLengthOfResult {
		t.Errorf("Expected length of result to be '%d' but got '%d' when passing '%s'", expectedLengthOfResult, len(result), str)
	}
}

func TestLongestPalindromWith2CharacterPalindrom(t *testing.T) {
	str := "asdfbssdiadfg"
	result := palindrom.LongestPalindrome(str)
	// every string contains a palindrom because a single character is a palindrom
	expectedResult := "ss"

	if result == expectedResult {
		t.Errorf("Expected result to be '%s' but got '%s' when passing '%s'", expectedResult, result, str)
	}
}
