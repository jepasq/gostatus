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

func TestSidebar_GetItemByName(t *testing.T) {
	arr := SidebarArray()
	it, err := SidebarArray_GetItemByLabel(arr, "Home")
	if (it.Label != "Home" || err != nil) {
		t.Fatalf("Can't get 'Home' item")
	}
}

func TestSidebar_GetItemByLabel_error(t *testing.T) {
	arr := SidebarArray()
	_, err := arr.GetItemByLabel("UnknownTooLongLabel")
	if (err == nil) {
		t.Fatalf("Unknown label didn't raise any error")
	}
}

func TestSidebar_MakeActive(t *testing.T) {
	arr := SidebarArray()
	it, err := arr.GetItemByLabel("Form")
	if (it.Active && err != nil) {
		t.Fatalf("Item shouldn't be active")
	}

	err = arr.MakeActive("Form")
	if (err != nil) {
		t.Fatalf("MakeActive returned an error : %s", err)
	}
	
	it, err = arr.GetItemByLabel("Form")
	if (!it.Active && err != nil) {
		t.Fatalf("Item is not active")
	}
}

/// The home element has a house icon field
func TestSidebar_HomeHouse(t *testing.T) {
	arr := SidebarArray()
	it, err := arr.GetItemByLabel("Home")
	if (it.IconClass != "house" && err != nil) {
		t.Fatalf("Home item should have a bi-house icon")
	}
}
