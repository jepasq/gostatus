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

func TestSidebarArrayCtor_GetItemByName(t *testing.T) {
	arr = SidebarArray()
	it = SidebarArray_GetItemByLabel(arr, "Home")
	if (it.name != "Home") {
		t.Fatalf("Can't get 'Home' item")
	}

}

