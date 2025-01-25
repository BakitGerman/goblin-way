package main

import qeque "github.com/BakitGerman/goblin-way/go-data-structures/queue"

func person_is_seller(name string) bool {
	return len(name) == 4 || name[len(name)-1] == 'y'
}

var graph = make(map[string][]string)

func main() {
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	searchWithDequeSlice("you")
	searchWithDequeLinkedList("you")
}

// O(V+E) - V кол-во ребер, E - кол-во людей
func searchWithDequeSlice(name string) bool {
	var search_queue []string
	search_queue = append(search_queue, graph[name]...)
	var searched []string
	var person string
	for len(search_queue) != 0 {
		person = search_queue[0]
		search_queue = search_queue[1:]
		if person_not_in_searched(person, searched) {
			if person_is_seller(person) {
				println("Deque (slice realization) is work! ", person+" is mango seller!")
				return true
			}

			search_queue = append(search_queue, graph[person]...)
			searched = append(searched, person)

		}
	}
	return false
}

// O(V+E) - V кол-во ребер, E - кол-во людей
func searchWithDequeLinkedList(name string) bool {
	deq := qeque.New()
	for idx := range graph[name] {
		deq.Push_back(graph[name][idx])
	}
	var searched []string
	var person string
	for deq.Len() != 0 {
		person = (deq.Pop_front()).(string)
		if person_not_in_searched(person, searched) {
			if person_is_seller(person) {
				println("Deque (double linked list realization) is work! ", person+" is mango seller!")
				return true
			}

			for idx := range graph[person] {
				deq.Push_back(graph[person][idx])
			}

			searched = append(searched, person)

		}
	}
	return false
}

func person_not_in_searched(person string, searched []string) bool {
	for _, n := range searched {
		if n == person {
			return false
		}
	}
	return true
}
