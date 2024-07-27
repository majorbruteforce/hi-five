package graph

import "container/list"

// BFS traverses the graph in a Breadth first order
// originating from start.
//
// The visit function is implemented upon visiting
// a vertex.
func (g *Graph) BFS(start *Vertex, visit func(v *Vertex)) {
	visited := make(map[*Vertex]bool)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		v := queue.Remove(queue.Front()).(*Vertex)
		visit(v)

		for adjacent := range g.Edges[v] {
			if !visited[adjacent] {
				visited[adjacent] = true
				queue.PushBack(adjacent)
			}
		}
	}
}

// DFS traverses the graph in a Depth first order
// beginning from start.
//
// The visit function is implemented upon visiting
// a vertex.
func (g *Graph) DFS(start *Vertex, visit func(v *Vertex)) {
	visited := make(map[*Vertex]bool)
	stack := list.New()
	stack.PushFront(start)
	visited[start] = true

	for stack.Len() > 0 {
		v := stack.Remove(stack.Front()).(*Vertex)
		visit(v)

		for adjacent := range g.Edges[v] {
			if !visited[adjacent] {
				visited[adjacent] = true
				stack.PushFront(adjacent)
			}
		}
	}
}
