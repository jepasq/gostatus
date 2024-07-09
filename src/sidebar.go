package main

import (
	"errors"
	"fmt"
)

/// IconClass can be find here https://icons.getbootstrap.com/
type SidebarItem struct {
	Label     string
	Active    bool
	Slur      string
	IconClass string
}

func SidebarArray() (Sidebar) {
	home  := SidebarItem{"Home",  true,  "/",     "house"}
	admin := SidebarItem{"Admin", false, "admin", "gear"}
	form  := SidebarItem{"Form",  false, "form",  "input-cursor-text"}

	arr := []SidebarItem{home, admin, form}
	
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
