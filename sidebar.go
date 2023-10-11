package main

import (
	"errors"
)

type SidebarItem struct {
	label  string
	active bool
	slur   string
}

func SidebarArray() ([]SidebarItem) {
	home  := SidebarItem{"Home",  true,  "/"}
	hello := SidebarItem{"Hello", false, "hello/"}
	form  := SidebarItem{"Form",  false, "hello/"}

	arr := []SidebarItem{home, hello, form}
	
	return arr;
}

func SidebarArray_GetItemByLabel(list []SidebarItem, label string) (SidebarItem, error) {
	for _, it := range list {
		if it.label == label {
			return it, nil
		}
		
	}
	return SidebarItem{}, errors.New("Not found")
}
