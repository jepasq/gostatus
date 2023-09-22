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
