package algorithms

import "leetcode/algorithms/graph"

func cloneGraph1(node *graph.Node) *graph.Node {
	visited := make(map[*graph.Node]*graph.Node)

	return clone(node, visited)
}

func clone(node *graph.Node, visited map[*graph.Node]*graph.Node) *graph.Node {
	if node == nil {
		return nil
	}

	if currentNode, ok := visited[node]; ok {
		return currentNode
	} else {
		visited[node] = &graph.Node{Val: node.Val}
	}

	for _, neighbour := range node.Neighbors {
		visited[node].Neighbors = append(visited[node].Neighbors, clone(neighbour, visited))
	}

	return visited[node]
}
