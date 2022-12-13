package assignments

import (
	"fmt"
	"strconv"
	"time"
)

/**
Create a program that accepts a number of vertices “V”, creates an undirected complete graph with “V” vertices using
an adjacency list data structure of your creation and then output the graph in the format below.
Do not include the outputting of the graph in your timing analysis.
*/
func RunHw3ProgramOne(v int) {
	if v < 2 {
		panic("v should greater than 1")
	}
	var completeGraph = make([]*Vertex, v)
	for i := 1; i <= v; i++ {
		completeGraph[i-1] = initVertex(i)
		for j := 1; j <= v; j++ {
			if i != j {
				completeGraph[i-1].addAdjacentVertex(j)
			}
		}
	}
	outPutGraph(completeGraph, "RunHw3ProgramOneExample")
}

/**
Create a program that accepts a number of vertices “V”, creates a cycle with “V” vertices using an adjacency list data
structure of your creation and then output the graph in the format below.
Do not include the outputting of the graph in your timing analysis
*/
func RunHw3ProgramTwo(v int) {
	if v < 2 {
		panic("v should greater than 1")
	}
	var cycleGraph = make([]*Vertex, v)
	for i := 1; i <= v; i++ {
		cycleGraph[i-1] = initVertex(i)
		left := i%v + 1
		right := i - 1
		if right == 0 {
			right = v
		}
		cycleGraph[i-1].addAdjacentVertex(left)
		if right != left {
			cycleGraph[i-1].addAdjacentVertex(right)
		}
	}
	outPutGraph(cycleGraph, "RunHw3ProgramTwoExample")
}

/**
Create a program that accepts a number of vertices “V” and a number of edges “E”. You program should create a graph
with “V” vertices and “E” edges between random pairs of vertices. Store your graph using an adjacency list data structure
of your creation and then output the graph in the format below. THETA(V^2 ) is fine. Do not include the outputting
 of the graph in your timing analysis. Be sure to include the case where all possible edges are created as well
  as the case with no edges.
*/
func RunHw3ProgramThree(v int, e int) {
	maxNumberOfEdges := v * (v - 1) / 2
	if e > maxNumberOfEdges {
		panic("wrong argument of e, it should not greater than v*(v-1)/2")
	}
	graph, _ := initGraph(v)
	possibleEdges := RunHw2ProgramTwo(maxNumberOfEdges, maxNumberOfEdges) //this function is used for generating non-repeat random number in Hw2
	index := 0
	for i := 0; i < v; i++ {
		for j := i + 1; j < v; j++ {
			if possibleEdges[index] <= e {
				graph[i].addAdjacentVertex(j + 1)
				graph[j].addAdjacentVertex(i + 1)
			}
			index++
		}
	}
	outPutGraph(graph, "RunHw3ProgramThreeExample")
}

/**
Create a program that accepts a number of vertices “V” and a number of edges “E”. You program should create a graph
with “V” vertices and “E” edges between random pairs of vertices. Store your graph using an adjacency list data structure
of your creation and then output the graph in the format below. THETA(V^2 ) is fine. Do not include the outputting
 of the graph in your timing analysis. Be sure to include the case where all possible edges are created as well
  as the case with no edges.
  linear time
*/
func RunHw3ProgramThree2(v int, e int) {
	maxNumberOfEdges := v * (v - 1) / 2
	if e > maxNumberOfEdges {
		panic("wrong argument of e, it should not greater than v*(v-1)/2")
	}
	graph := make([]*Vertex, v)
	pickNum := e
	picked := make([]int, e*2)
	possibleEdges := RunHw2ProgramTwo(v*v, v*v) //this function is used for generating non-repeat random number in Hw2
	for i := 0; i < v; i++ {
		graph[i] = initVertex(i + 1)
	}
	for i := 0; i < v*v; i++ {
		index := possibleEdges[i] - 1 //index from 1 to v*v
		left := index/v + 1
		right := index%v + 1
		// fmt.Println("possibility: " + fmt.Sprint(possibleEdges[i]))
		// fmt.Println("left: " + fmt.Sprint(left))
		// fmt.Println("right: " + fmt.Sprint(right))
		conflict := false
		if left == right || left > v {
			continue
		}
		for j := 0; j < e*2; j++ {
			if index == picked[j] {
				conflict = true
				break
			}
		}
		if conflict {
			continue
		}
		graph[left-1].addAdjacentVertex(right)
		graph[right-1].addAdjacentVertex(left)
		picked[pickNum-1] = index
		picked[pickNum-1+e] = (right-1)*v + left - 1
		pickNum--
		if pickNum <= 0 {
			break
		}
	}
	// outPutGraph(graph, "tt")
}

func RunHw3Main() {
	// chs := make([]chan int, 2)
	// chs[0] = make(chan int)
	// chs[1] = make(chan int)
	// chs[2] = make(chan int)

	// go RuntestHw3ProgramOne(chs[0])
	// go RuntestHw3ProgramTwo(chs[1])
	// go RuntestHw3ProgramThree(chs[2])

	// for _, ch := range chs {
	// 	<-ch
	// }
	// RunHw3ProgramThree2(6, 14)
	// RunHw3ProgramThree(6, 12)

	// RuntestHw3ProgramThreeSingle()
	// RuntestHw3ProgramTwoSingle()
	// RuntestHw3ProgramOneSingle()

	// RunHw3ProgramOne(6)
	// RunHw3ProgramOne(12)
	// RunHw3ProgramTwo(12)
	RunHw3ProgramTwo(6)

	fmt.Println("done")
}

func RuntestHw3ProgramOneSingle() {
	i := 10
	for i <= 2500 {
		start1 := time.Now()
		RunHw3ProgramOne(i)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw3ProgramOne-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
}

func RuntestHw3ProgramTwoSingle() {
	i := 10
	for i <= 2500 {
		start1 := time.Now()
		RunHw3ProgramTwo(i)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw3ProgramTwo-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
}

func RuntestHw3ProgramThreeSingle() {
	i := 10
	for i <= 2500 {
		start1 := time.Now()
		RunHw3ProgramThree(i, i*(i-1)/2-2)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw3ProgramThree-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
}

func RuntestHw3ProgramTwo(ch chan int) {
	i := 10
	for i <= 2500 {
		start1 := time.Now()
		RunHw3ProgramTwo(i)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw3ProgramTwo-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
	ch <- 1
}

func RuntestHw3ProgramThree(ch chan int) {
	i := 10
	for i <= 2500 {
		start1 := time.Now()
		RunHw3ProgramThree(i, i*(i-1)/2-2)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw3ProgramThree-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
	ch <- 1
}

func RuntestHw3ProgramOne(ch chan int) {
	i := 10
	for i <= 2500 {
		start1 := time.Now()
		RunHw3ProgramOne(i)
		end1 := time.Since(start1).Microseconds()
		writeFile("RunHw3ProgramOne-", fmt.Sprint(i)+"\t"+strconv.FormatInt(end1, 10)+"\n")
		i++
	}
	ch <- 1
}
