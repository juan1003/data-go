package main

import (
    "fmt"
)

type Node struct {
    value int
    neighbors []*Node
}

type Graph struct {
    nodes map[int]*Node
}

// Inits an empty Graph
func (g *Graph) Initialize() {
    g.nodes = make(map[int]*Node)
}


// Adds a Node to the Graph
func (g *Graph) AddNode(value int) {
    node := &Node {
        value: value,
    }
    g.nodes[value] = node
}

// AddEdge adds an edge between two nodes
func (g* Graph) AddEdge(fromValue, toValue int) {
    fromNode := g.nodes[fromValue]
    toNode := g.nodes[toValue]
    if fromNode != nil && toNode != nil {
        fromNode.neighbors = append(fromNode.neighbors, toNode)
    }
}

// RemoveEdge removes an edge between two nodes
func (g* Graph) RemoveEdge(fromValue, toValue int) {
    fromNode := g.nodes[fromValue]
    toNode := g.nodes[toValue]
    if fromNode != nil && toNode != nil {
        for i, neighbor := range fromNode.neighbors {
            if neighbor == toNode {
                fromNode.neighbors = append(fromNode.neighbors[:i], fromNode.neighbors[i+1:]...)
                break
            }
        }
    }
}

// HasEdge checks if an edge exists between two nodes
func (g* Graph) HasEdge (fromValue, toValue int) bool {
    fromNode := g.nodes[fromValue]
    toNode := g.nodes[toValue]
    if fromNode != nil && toNode != nil {
        for _, neighbor := range fromNode.neighbors {
            if neighbor == toNode {
                return true
            }
        }
    }

    return false
}

// GetNeighbors returns the neighbors of a Node
func (g* Graph) GetNeighbors(value int) []*Node {
    node := g.nodes[value]
    if node != nil {
        return node.neighbors
    }

    return nil
}

func main() {
    graph := Graph{}
    graph.Initialize()

    graph.AddNode(1)
    graph.AddNode(2)
    graph.AddNode(3)
    graph.AddNode(4)

    graph.AddEdge(1, 2)
    graph.AddEdge(1, 3)
    graph.AddEdge(2, 4)
    graph.AddEdge(3, 4)

    fmt.Println(graph.HasEdge(1, 2)) // true
    fmt.Println(graph.HasEdge(2, 3)) // false

    neighbors := graph.GetNeighbors(1)
    for _, neighbor := range neighbors {
        fmt.Println(neighbor.value)
    }

    // Output:
    // 2
    // 3
}
