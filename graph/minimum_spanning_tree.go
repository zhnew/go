package goraph

import (
	"container/heap"
	"fmt"
	"sort"
)

func Kruskal(g Graph) (map[Edge]struct{}, error) {
	A := make(map[Edge]struct{})
	forests := NewForests()
	for _, nd := range g.GetNodes() {
		MakeDisjointSet(forests, nd.String())
	}
	edges := g.GetEdges()[:]
	sort.Sort(EdgeSlice(edges))
	for _, edge := range edges {
		if FindSet(forests, edge.Source().String()).represent != FindSet(forests, edge.Target().String()).represent {
			A[edge] = struct{}{}
			Union(forests, FindSet(forests, edge.Source().String()), FindSet(forests, edge.Target().String()))
		}
	}
	return A, nil
}

func Prim(g Graph, src ID) (map[Edge]struct{}, error) {
	fmt.Printf("this is Prim alg start at %s \n", src.String())
	edges := make(map[Edge]struct{})
	//nodeIDs record the node already in mini tree
	nodeIDs := make(map[ID]bool)
	for id := range g.GetNodes() {
		nodeIDs[id] = false
	}
	nodeIDs[src] = true
	//store all edges
	minHeap := &edgeHeap{}
	for _, e := range g.GetOutEdge(src) {
		heap.Push(minHeap, e)
	}
	for len(edges) < g.GetNodeCount()-1 {
		e := heap.Pop(minHeap).(Edge)
		uID := e.Target().ID()

		if nodeIDs[uID] == true {
			continue
		}

		edges[e] = struct{}{}
		nodeIDs[uID] = true
		//edge (u,v) is added to minHeap
		for _, e := range g.GetOutEdge(uID) {
			vID := e.Target().ID()
			if nodeIDs[vID] != true {
				heap.Push(minHeap, e)
			}
		}
	}
	return edges, nil
}

type edgeHeap []Edge

func (h edgeHeap) Len() int           { return len(h) }
func (h edgeHeap) Less(i, j int) bool { return (h[i]).Weight() < h[j].Weight() }
func (h edgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *edgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *edgeHeap) Pop() interface{} {
	heapSize := len(*h)
	lastNode := (*h)[heapSize-1]
	*h = (*h)[0 : heapSize-1]
	return lastNode
}
