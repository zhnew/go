package goraph

import (
	"sync"
)

type DisjointSet struct {
	represent string
	members   map[string]struct{}
}
type Forests struct {
	mu   sync.Mutex
	data map[*DisjointSet]struct{}
}

func NewForests() *Forests {
	set := &Forests{
		data: make(map[*DisjointSet]struct{}),
	}
	return set
}
func MakeDisjointSet(forests *Forests, name string) {
	members := make(map[string]struct{})
	members[name] = struct{}{}
	newDS := &DisjointSet{
		represent: name,
		members:   members,
	}
	forests.mu.Lock()
	defer forests.mu.Unlock()
	forests.data[newDS] = struct{}{}
}
func FindSet(forests *Forests, name string) *DisjointSet {
	forests.mu.Lock()
	defer forests.mu.Unlock()
	for data := range forests.data {
		if data.represent == name {
			return data
		}
		for k := range data.members {
			if k == name {
				return data
			}
		}
	}
	return nil
}

func Union(forests *Forests, ds1, ds2 *DisjointSet) {
	for k, v := range ds2.members {
		ds1.members[k] = v
	}
	forests.mu.Lock()
	defer forests.mu.Unlock()
	delete(forests.data, ds2)
}
