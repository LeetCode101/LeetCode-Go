package algorithms

import "leetcode/algorithms/graph"

func cloneGraph2(node *graph.Node) *graph.Node {
	if node == nil {
		return nil
	}

	queue := []*graph.Node{node}
	visited := make(map[*graph.Node]*graph.Node)
	visited[node] = &graph.Node{Val: node.Val}

	for len(queue) > 0 {
		oldNode := queue[0]
		queue = queue[1:]

		for _, neighbour := range oldNode.Neighbors {
			if _, ok := visited[neighbour]; !ok {
				visited[neighbour] = &graph.Node{Val: neighbour.Val}
				queue = append(queue, neighbour)
			}

			visited[oldNode].Neighbors = append(visited[oldNode].Neighbors, visited[neighbour])
		}
	}

	return visited[node]
}
