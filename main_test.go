package main

import "testing"

func TestBla(t *testing.T) {

	if s := Bla(); s == "" {
		t.Errorf("error: %v", s)
	}
}
