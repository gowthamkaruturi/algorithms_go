package graphs

import "fmt"

type Graph struct {
	vertices     int
	adjacentList []DoublyLinkedList
}

func InitGraph(vertices int) *Graph {
	return &Graph{vertices: vertices,
		adjacentList: make([]DoublyLinkedList, vertices),
	}
}

func (g *Graph) AddEdge(source, destination int) {
	if source < g.vertices && destination < g.vertices {
		g.adjacentList[source].AddFromFront(destination)
		//g.adjacentList[destination].AddFromEnd(source)
	}
}

func (g *Graph) PrintGraph() {
	for i := 0; i < g.vertices; i++ {
		if !g.adjacentList[i].IsEmpty() {
			fmt.Printf("{%d} =>", i)
			g.adjacentList[i].Print()
		} else {
			fmt.Printf("{%d} => {%s} <-->", i, "null")
		}

	}

}
