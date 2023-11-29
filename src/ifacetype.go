package main

/// The interface used in this session



// IfacetypeToStr as int
type Ifacetype int

// The possible Ifacetype values
const (
	Web = iota    // The interface is a web one
	Cli           // The interface is console-based
	Tui           // The interface is terminal-based
)


// Returns an string that represent the given interface type
func IfacetypeToStr(t Ifacetype) (string) {
	switch t {
	case Web:
		return "Web"
	case Cli:
		return "Cli"
	case Tui:
		return "Tui"
	default:
		return "Unset"
	}
}

