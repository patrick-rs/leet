package test

import (
	"testing"

	"github.com/patrick-rs/leet/pkg/zigzagconversion"
)

func TestFourChar(t *testing.T) {
	str := "PAYP"
	rows := 3

	expectedResult := "PAPY"
	result := zigzagconversion.Convert(str, rows)

	if result != expectedResult {
		t.Errorf("Expected:\n %v\n but got:\n %v\n when passing string: '%s' with rows; %d", expectedResult, result, str, rows)
	}
}

func TestName(t *testing.T) {
	str := "PATRICK"
	rows := 4

	expectedResult := "PKACTIR"
	result := zigzagconversion.Convert(str, rows)

	if result != expectedResult {
		t.Errorf("Expected:\n %v\n but got:\n %v\n when passing string: '%s' with rows; %d", expectedResult, result, str, rows)
	}
}

func TestSingleChar(t *testing.T) {
	str := "A"
	rows := 1

	expectedResult := "A"
	result := zigzagconversion.Convert(str, rows)

	if result != expectedResult {
		t.Errorf("Expected:\n %v\n but got:\n %v\n when passing string: '%s' with rows; %d", expectedResult, result, str, rows)
	}
}

func TestTwoChar(t *testing.T) {
	str := "AB"
	rows := 1

	expectedResult := "AB"
	result := zigzagconversion.Convert(str, rows)

	if result != expectedResult {
		t.Errorf("Expected:\n %v\n but got:\n %v\n when passing string: '%s' with rows; %d", expectedResult, result, str, rows)
	}
}
