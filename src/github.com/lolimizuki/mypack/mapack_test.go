package mypack

import (
	"testing"
)

func TestMyFunction(t *testing.T) {
	corrString := "String from my function!!!"
	if MyFunction() != corrString {
		t.Errorf("string is correct")
	}
}
