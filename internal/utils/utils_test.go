package utils

import "testing"

func TestToLowerCase(t *testing.T) {
	res := ToLowerCase("en-US")
	if res != "en-us" {
		t.Fatalf("Error")
	}
}

func TestToCleanCode(t *testing.T) {
	res := ToCleanCode("en-US")
	if res != "en" {
		t.Fatalf("Error")
	}
}
