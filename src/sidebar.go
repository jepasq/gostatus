package main

import (
	"errors"
	"fmt"
)

type SidebarItem struct {
	Label  string
	Active bool
	Slur   string
}

func SidebarArray() (Sidebar) {
	home  := SidebarItem{"Home",  true,  "/"}
	hello := SidebarItem{"Hello", false, "hello"}
	form  := SidebarItem{"Form",  false, "form"}

	arr := []SidebarItem{home, hello, form}
	
	return arr;
}

type Sidebar []SidebarItem

func SidebarArray_GetItemByLabel(list []SidebarItem, label string) (SidebarItem, error) {
	for _, it := range list {
		if it.Label == label {
			return it, nil
		}
		
	}
	return SidebarItem{}, errors.New("Not found")
}

func (arr Sidebar) MakeActive(label string) (error) {
	found := false
	for i, it := range arr {
		fmt.Printf(" > : '%s'\n", it.Label)

		if it.Label == label {
			// see https://stackoverflow.com/a/15952415
			arr[i].Active = true 
			found = true
		} else {
			arr[i].Active = false
		}
		
	}
	if found == false {
		return errors.New(fmt.Sprintf("Can't find irem '%s' in '%v'",
			label, arr))
	} else {
		return nil
	}
}

func (arr Sidebar) GetItemByLabel(label string) (SidebarItem, error) {
	for _, it := range arr {
		if it.Label == label {
			return it, nil
		}
		
	}
	return SidebarItem{},
		errors.New("Failed to get Sidebar item from label")

}
