package trypkg

import "testing"

func Test(t *testing.T) {
	result := get_three()
	if result != "3" {
		t.Errorf("it's error")
	}
}
