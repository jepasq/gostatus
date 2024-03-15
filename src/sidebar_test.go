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
	if (it.label != "Home" || err != nil) {
		t.Fatalf("Can't get 'Home' item")
	}
}

func TestSidebar_GetItemByName_error(t *testing.T) {
	arr := SidebarArray()
	_, err := SidebarArray_GetItemByLabel(arr, "UnknownTooLongLabel")
	if (err == nil) {
		t.Fatalf("Unknown label didn't raise any error")
	}
}

func TestSidebar_MakeActive(t *testing.T) {
	arr := SidebarArray()
	it, err := SidebarArray_GetItemByLabel(arr, "Form")
	if (it.active && err != nil) {
		t.Fatalf("Item shouldn't be active")
	}

	err = arr.MakeActive("Form")
	if (err != nil) {
		t.Fatalf("MakeActive returned an error")
	}
	
	it, err = SidebarArray_GetItemByLabel(arr, "Form")
	if (!it.active && err != nil) {
		t.Fatalf("Item is not active")
	}
}

