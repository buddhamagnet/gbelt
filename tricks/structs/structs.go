package structs

// T is a package level struct.
var T struct {
	Name   string
	Number int
}

func init() {
	T.Name = "package struct"
	T.Number = 100
}
