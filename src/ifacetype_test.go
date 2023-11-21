package main

import (
	"testing"
)


func TestIfaceToStr(t *testing.T) {
	iftype := IfacetypeToStr(Web);
	
	if iftype != "Web" {
		t.Fatalf(`Iface string doesn't match 'Web'`)
	}
}
