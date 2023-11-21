package main

import (
	"testing"
)

func TestSidebarArrayCtor(t *testing.T) {
	arr := SidebarArray()
	if (arr == nil) {
		t.Fatalf(`SidebarArray ctor failed`)
	}
}

func TestSidebarArrayCtor_GetItemByName(t *testing.T) {
	arr := SidebarArray()
	it, err := SidebarArray_GetItemByLabel(arr, "Home")
	if (it.label != "Home" || err != nil) {
		t.Fatalf("Can't get 'Home' item")
	}

}

