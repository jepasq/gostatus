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
	for _, it := range arr {
		fmt.Printf(" > : '%s'\n", it.Label)

		if it.Label == label {
			it.Active = true
		} else {
			it.Active = false
		}
		
	}
	return errors.New(fmt.Sprintf("Can't find irem '%s' in '%v'",
		label, arr))
}

func (arr Sidebar) GetItemByLabel(label string) SidebarItem {
	si, err := SidebarArray_GetItemByLabel(arr, label)
	
	if err != nil {
		fmt.Printf("Failed to get Sidebar item from label '%s'\n",
			label)
	}

	return si
}
