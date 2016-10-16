package goraph

import (
	"container/heap"
	"fmt"
	"math"
)

type nodeDistance struct {
	id       ID
	distance float64
}
type nodeDistanceHeap []nodeDistance

func (h nodeDistanceHeap) Len() int           { return len(h) }
func (h nodeDistanceHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h nodeDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(nodeDistance))
}
func (h *nodeDistanceHeap) Pop() interface{} {
	heapSize := len(*h)
	lastNode := (*h)[heapSize-1]
	*h = (*h)[0 : heapSize-1]
	return lastNode
}
func (h *nodeDistanceHeap) updateDistance(id ID, val float64) {
	for i := 0; i < len(*h); i++ {
		if (*h)[i].id == id {
			(*h)[i].distance = val
			break
		}
	}
}

func Dijkstra(g Graph, source, target ID) ([]ID, map[ID]float64, error) {
	minHeap := &nodeDistanceHeap{}
	distance := make(map[ID]float64)
	distance[source] = 0.0
	for id := range g.GetNodes() {
		if id != source {
			distance[id] = math.MaxFloat64
		}
		nds := nodeDistance{}
		nds.id = id
		nds.distance = distance[id]
		heap.Push(minHeap, nds)
	}
	heap.Init(minHeap)
	prev := make(map[ID]ID)
	for minHeap.Len() != 0 {
		u := heap.Pop(minHeap).(nodeDistance)
		if u.id == target {
			break
		}
		cmap, err := g.GetTargets(u.id)
		if err != nil {
			return nil, nil, err
		}
		for v := range cmap {
			weight, err := g.GetWeight(u.id, v)
			if err != nil {
				return nil, nil, err
			}
			alt := distance[u.id] + weight
			if distance[v] > alt {
				distance[v] = alt
				prev[v] = u.id
				minHeap.updateDistance(v, alt)
			}
		}
		heap.Init(minHeap)
	}
	path := []ID{}
	u := target
	for {
		if _, ok := prev[u]; !ok {
			break
		}
		path = append(path, u)
		//temp := make([]ID, len(path)+1)
		//temp[0] = u
		//copy(temp[1:], path)
		//path = temp
		u = prev[u]
	}
	//temp := make([]ID, len(path)+1)
	//temp[0] = source
	//copy(temp[1:], path)
	//path = temp
	path = append(path, source)
	//reverse path now
	n := len(path)
	for i := 0; i < n/2; i += 1 {
		path[i], path[n-i-1] = path[n-i-1], path[i]
	}
	return path, distance, nil
}

func BellmanFord(g Graph, source, target ID) ([]ID, map[ID]float64, error) {
	distance := make(map[ID]float64)
	distance[source] = 0.0
	for id := range g.GetNodes() {
		if id != source {
			distance[id] = math.MaxFloat64
		}
	}
	prev := make(map[ID]ID)
	for i := 1; i <= g.GetNodeCount()-1; i += 1 {
		for id := range g.GetNodes() {
			cmap, err := g.GetTargets(id)
			if err != nil {
				return nil, nil, err
			}
			u := id
			for v := range cmap {
				weight, err := g.GetWeight(u, v)
				if err != nil {
					return nil, nil, err
				}
				alt := distance[u] + weight
				if distance[v] > alt {
					distance[v] = alt
					prev[v] = u
				}
			}
			pmap, err := g.GetSources(id)
			if err != nil {
				return nil, nil, err
			}
			v := id
			for u := range pmap {
				weight, err := g.GetWeight(u, v)
				if err != nil {
					return nil, nil, err
				}
				alt := distance[u] + weight
				if distance[v] > alt {
					distance[v] = alt
					prev[v] = u
				}
			}
		}
	}
	for id := range g.GetNodes() {
		cmap, err := g.GetTargets(id)
		if err != nil {
			return nil, nil, err
		}
		u := id
		for v := range cmap {
			weight, err := g.GetWeight(u, v)
			if err != nil {
				return nil, nil, err
			}

			// alt = distance[u] + weight(u, v)
			alt := distance[u] + weight

			// if distance[v] > alt:
			if distance[v] > alt {
				return nil, nil, fmt.Errorf("there is a negative-weight cycle: %v", g)
			}
		}

		pmap, err := g.GetSources(id)
		if err != nil {
			return nil, nil, err
		}
		v := id
		for u := range pmap {
			// edge (u, v)
			weight, err := g.GetWeight(u, v)
			if err != nil {
				return nil, nil, err
			}

			// alt = distance[u] + weight(u, v)
			alt := distance[u] + weight

			// if distance[v] > alt:
			if distance[v] > alt {
				return nil, nil, fmt.Errorf("there is a negative-weight cycle: %v", g)
			}
		}
	}
	path := []ID{}
	u := target
	for {
		if _, ok := prev[u]; !ok {
			break
		}
		path = append(path, u)
		u = prev[u]
	}
	path = append(path, source)
	n := len(path)
	for i := 0; i < n/2; i += 1 {
		path[i], path[n-i-1] = path[n-i-1], path[i]
	}
	return path, distance, nil
}
