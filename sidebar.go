package main

type sidebarItem struct {
	label  string
	active bool
	slur   string
}

func SidebarArray() ([]sidebarItem) {
	home  := sidebarItem{"Home",  true,  "/"}
	hello := sidebarItem{"Hello", false, "hello/"}
	form  := sidebarItem{"Form",  false, "hello/"}

	arr := []sidebarItem{home, hello, form}
	
	return arr;
}

func SidebarArray_GetItemByLabel(list []sidebarItem, label string) (sidebarItem) {
	for _, it := range list {
		if it.label == label {
			return it
		}
		
	}
	//	return nil
}
