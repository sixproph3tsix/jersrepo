package main

import (
	"testing"
)

func TestGetBool(t *testing.T) {
	bools := []struct {
		input    string
		expected bool
	}{
		{"TRUE", true},
		{"true", true},
		//{"  TRUE ", true},
		{"FALSE", false},
		//{" FALSE  ", false},
	}

	for _, bool := range bools {
		result := GetBool(bool.input)
		expected := bool.expected

		if result != bool.expected {
			t.Errorf("expected: %v, got: %v", expected, result)
		}

	}
}

func TestGetInt(t *testing.T) {
	ints := []struct {
		input string
		expected int
	}{
		{"2", 2},
		{"3", 3},
		{"1", 1},
	}

	for _, int := range ints {
		result := GetInt(int.input)
		expected := int.expected

		if result != expected {
			t.Errorf("got: %v, wanted: %v", result, expected)
		}
	}
}

func TestGetFloat64(t *testing.T) {
	floats := []struct {
		input string
		expected float64
	}{
		{"1.23", 1.23},
		{"111.2222", 111.2222},
	}

	for _, float := range floats {
		result := GetFloat64(float.input)
		expected := float.expected

		if result != expected {
			t.Errorf("got: %v, expected: %v\n", result, expected)
		}
	}
}
