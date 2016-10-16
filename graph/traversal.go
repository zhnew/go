package goraph

func BFS(g Graph, id ID) []ID {
	if g.GetNode(id) == nil {
		return nil
	}
	q := []ID{id}
	visited := make(map[ID]bool)
	visited[id] = true
	rs := []ID{id}
	for len(q) != 0 {
		u := q[0]
		q = q[1:len(q):len(q)]
		cmap, _ := g.GetTargets(u)
		for _, w := range cmap {
			if _, ok := visited[w.ID()]; !ok {
				q = append(q, w.ID())
				visited[w.ID()] = true
				rs = append(rs, w.ID())
			}
		}
		pmap, _ := g.GetSources(u)
		for _, w := range pmap {
			if _, ok := visited[w.ID()]; !ok {
				q = append(q, w.ID())
				visited[w.ID()] = true
				rs = append(rs, w.ID())
			}
		}
	}
	return rs
}
func DFS(g Graph, id ID) []ID {
	if g.GetNode(id) == nil {
		return nil
	}
	visited := make(map[ID]bool)
	rs := []ID{}
	dfs(g, id, &rs, visited)
	return rs
}
func dfs(g Graph, id ID, rs *[]ID, visited map[ID]bool) {
	if _, ok := visited[id]; ok {
		return
	}
	visited[id] = true
	*rs = append(*rs, id)
	cmap, _ := g.GetTargets(id)
	for _, w := range cmap {
		if _, ok := visited[w.ID()]; !ok {
			dfs(g, w.ID(), rs, visited)
		}
	}
	pmap, _ := g.GetSources(id)
	for _, w := range pmap {
		if _, ok := visited[w.ID()]; !ok {
			dfs(g, w.ID(), rs, visited)
		}
	}
}
func DFSWithoutRecursion(g Graph, id ID) []ID {
	if g.GetNode(id) == nil {
		return nil
	}
	visited := make(map[ID]bool)
	rs := []ID{}
	stack := []ID{id}
	for len(stack) > 0 {
		id := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[id]; ok {
			continue
		}
		visited[id] = true
		rs = append(rs, id)
		cmap, _ := g.GetTargets(id)
		for _, w := range cmap {
			if _, ok := visited[w.ID()]; !ok {
				stack = append(stack, w.ID())
			}
		}
		pmap, _ := g.GetSources(id)
		for _, w := range pmap {
			if _, ok := visited[w.ID()]; !ok {
				stack = append(stack, w.ID())
			}
		}
	}
	return rs
}
