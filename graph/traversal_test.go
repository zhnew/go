package goraph

import (
	"fmt"
	"os"
	"testing"
)

func TestBFS(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	rs := BFS(g, StringID("S"))
	fmt.Println("BFS:", rs)
	if len(rs) != 8 {
		t.Errorf("should be 8 vertices but %s", g)
	}
}

func TestDFS(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	rs := DFS(g, StringID("S"))
	fmt.Println("DFS:", rs)
	if len(rs) != 8 {
		t.Errorf("should be 8 vertices but %s", g)
	}
}
func TestDFSWithoutRecursion(t *testing.T) {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	g, err := NewGraphFromJSON(f, "graph_00")
	if err != nil {
		t.Error(err)
	}
	rs := DFSWithoutRecursion(g, StringID("S"))
	fmt.Println("DFS:", rs)
	if len(rs) != 8 {
		t.Errorf("should be 8 vertices but %s", g)
	}
}
