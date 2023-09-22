package main

import (
	"testing"
)

func TestSidebarArrayCtor(t *testing.T) {
	_ = SidebarArray()
	t.Fatalf(`SidebarArray`)
	/*	if !want.MatchString(string(data)) {
		t.Fatalf(`SidebarArray`)
	}
	*/
}
