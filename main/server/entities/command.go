package entities

type Command struct {
	Fields []string
	Result chan string
}

var Data = make(map[string]string)