package main

type Node struct {
	Name  string
	Value int
}

func main() {
	// variant 1
	graph := make(map[string][]Node, 5)
	graph["Книга"] = []Node{{Name: "Постер", Value: 0}, {Name: "Редкая пластина", Value: 5}}
	graph["Постер"] = []Node{{Name: "Бас-гитара", Value: 30}, {Name: "Барабан", Value: 35}}
	graph["Редкая пластина"] = []Node{{Name: "Бас-гитара", Value: 15}, {Name: "Барабан", Value: 20}}
	graph["Бас-гитара"] = []Node{{Name: "Пианино", Value: 20}}
	graph["Барабан"] = []Node{{Name: "Пианино", Value: 10}}
	graph["Пианино"] = []Node{}

	for i := 0; i < len(graph); i++ {

	}
}
