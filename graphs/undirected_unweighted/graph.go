package main

import "fmt"

type graph map[string][]string

func (g graph) addVertexToGraph(vtx string) {
	if g[vtx] != nil {
		fmt.Println("\n-- Vertex already exists. --")
		return
	}
	g[vtx] = make([]string, 0)
}

func (g graph) addEdgeToGraph(fromVtx, toVtx string) {
	if g[fromVtx] == nil { // check if initial vertex exists
		fmt.Println("\n-- Initial vertex " + fromVtx + " does not exist. --")
		return
	}
	for i := range g[fromVtx] { // check if edge already exists
		if g[fromVtx][i] == toVtx {
			fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " already exists. --")
			return
		}
	}
	if g[toVtx] == nil { // create new destination vertext if it does not exists
		g[toVtx] = make([]string, 0)
		fmt.Println("\n-- Destination vertex " + toVtx + " created. --")
	}

	g[fromVtx] = append(g[fromVtx], toVtx)
	g[toVtx] = append(g[toVtx], fromVtx)
	return
}

func (g graph) removeVertexFromGraph(vtx string) {
	length := len(g[vtx]) - 1
	for length != -1 {
		g.removeEdgeFromGraph(vtx, g[vtx][length])
		length--
	}
	delete(g, vtx)
}

func (g graph) removeEdgeFromGraph(fromVtx, toVtx string) {
	if g[fromVtx] == nil || g[toVtx] == nil {
		fmt.Println("\n-- Edge between " + fromVtx + " and " + toVtx + " does not exist. --")
		return
	}

	for i := range g[fromVtx] {
		if g[fromVtx][i] == toVtx {
			if i == 0 {
				g[fromVtx] = g[fromVtx][1:len(g[fromVtx])]
			} else if i == (len(g[fromVtx]) - 1) {
				g[fromVtx] = g[fromVtx][0:(len(g[fromVtx]) - 1)]
			} else {
				initial := g[fromVtx][0:i]
				final := g[fromVtx][i+1 : len(g[fromVtx])]
				g[fromVtx] = append(initial, final...)
			}
			break
		}
	}

	for i := range g[toVtx] {
		if g[toVtx][i] == fromVtx {
			if i == 0 {
				g[toVtx] = g[toVtx][1:len(g[toVtx])]
			} else if i == (len(g[toVtx]) - 1) {
				g[toVtx] = g[toVtx][0:(len(g[toVtx]) - 1)]
			} else {
				initial := g[toVtx][0:i]
				final := g[toVtx][i+1 : len(g[toVtx])]
				g[toVtx] = append(initial, final...)
			}
			break
		}
	}
}

func main() {
	i := 0
	g := make(graph)
	for i == 0 {
		fmt.Println("\n1. ADD A VERTEX")
		fmt.Println("2. ADD AN EDGE")
		fmt.Println("3. REMOVE VERTEX")
		fmt.Println("4. REMOVE AN EDGE")
		fmt.Println("5. DISPLAY USING DFS")
		fmt.Println("6. DISPLAY USING BFS")
		fmt.Println("7. SIMPLE DISPLAY")
		fmt.Println("8. EXIT")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)
		switch choice {
		case 1:
			g.addVertex()
		case 2:
			g.addEdge()
		case 3:
			g.removeVertex()
		case 4:
			g.removeEdge()
		case 5:
			result := g.displayDFS()
			fmt.Println(result)
		case 6:
			result := g.displayBFS()
			fmt.Println(result)
		case 7:
			g.simpleDisplay()
		case 8:
			i = 1
		default:
			fmt.Println("Command not recognized.")
		}
	}
}

func (g graph) addVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s\n", &vtxName)
	g.addVertexToGraph(vtxName)
}

func (g graph) addEdge() {
	var fromVtx, toVtx string
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s\n", &toVtx)
	g.addEdgeToGraph(fromVtx, toVtx)
}

func (g graph) removeVertex() {
	var vtxName string
	fmt.Print("Enter the name of vertex: ")
	fmt.Scanf("%s\n", &vtxName)
	g.removeVertexFromGraph(vtxName)
}

func (g graph) removeEdge() {
	var fromVtx, toVtx string
	fmt.Print("Enter the initial vertex name: ")
	fmt.Scanf("%s\n", &fromVtx)
	fmt.Print("Enter the destination vertex name: ")
	fmt.Scanf("%s\n", &toVtx)
	g.removeEdgeFromGraph(fromVtx, toVtx)
}

func (g graph) displayDFS() []string {
	var startVtx string
	fmt.Print("Enter the start vertex name: ")
	fmt.Scanf("%s\n", &startVtx)
	result := g.DFS(startVtx)
	return result
}

func (g graph) displayBFS() []string {
	var startVtx string
	fmt.Print("Enter the start vertex name: ")
	fmt.Scanf("%s\n", &startVtx)
	result := g.BFS(startVtx)
	return result
}

func (g graph) simpleDisplay() {
	fmt.Println("")
	for i := range g {
		fmt.Print(i, " => ")
		for j := range g[i] {
			fmt.Print(g[i][j])
		}
		fmt.Println("")
	}
}
