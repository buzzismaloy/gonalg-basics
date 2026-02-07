package foo

// privateFoo - non-exported type
type privateFoo struct {
	Value string
}

// NewPrivateFoo - constructor of privateFoo type
// this func is public
func NewPrivateFoo() privateFoo {
	return privateFoo{Value: "some data"}
}

func main() {
}
