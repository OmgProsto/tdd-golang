package golang_tdd

func Hello() string {
	return "Hello"
}

func HelloName(name string) string {
	return "Hello " + name
}

func HelloDefault(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello " + name
}
