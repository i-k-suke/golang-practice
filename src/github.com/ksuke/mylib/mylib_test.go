package mylib

import "testing"

/**
* execute test code: g
* $ go test -v ./...
*/
func TestAverage(t *testing.T) {
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
