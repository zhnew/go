package main

import (
	"fmt"
	//"github.com/zhnew/go/graph/testgraph"
	graph "github.com/zhnew/go/graph"
	"os"
)

func main() {
	g1 := graph.NewGraph()
	fmt.Println("g1:", g1.String())
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	g2, err := graph.NewGraphFromJSON(f, "graph_00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("g2:", g2.String())
}
