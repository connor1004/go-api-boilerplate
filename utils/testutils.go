package utils

import (
	"testing"
)

// CheckResponseCode - check if the expected code is equal to actual code
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
