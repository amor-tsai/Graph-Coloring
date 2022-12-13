package assignments

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

var TestArr = map[string]int{
	"Test_project":                                  1,
	"Test_createCompleteGraphAndColoring":           0,
	"Test_AllColoring":                              0,
	"Test_coloringGraphWithRandomOrdering":          0,
	"Test_coloringGraphWithTheSmallestLastOrdering": 0,
	"Test_coloringGraphWithTheLargestLastOrdering":  0,
	"Test_initGraph":                                0,
	"Test_createCompleteGraph":                      0,
	"Test_createCycleGraph":                         0,
	"Test_createRandomGraphWithUniformDistribution": 0,
	"Test_createRandomGraphWithSkewedDistribution":  0,
	"Test_createRandomGraphWithAmorDistribution":    0,
}

func Test_project(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}

	graph, adjacentMatrix, _ := createCycleGraph(10)
	outPutGraph(graph,"testCycleColoring")
	_, _, orderList := coloringGraphWithRandomOrdering(graph, adjacentMatrix)
	TestOutput = 1
	startColoring(graph, adjacentMatrix, orderList)
	// markColoringInfo(graph, orderList, 1, 1, "1")

}

func Test_AllColoring(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	fmt.Println(getFuncName())
	for i := 0; i < 1; i++ {
		// graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(300)
		// graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(100, 500)
		graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(5, 6)
		// fmt.Println("under coloringGraphWithRandomOrdering")
		// coloringGraphWithRandomOrdering(graph, adjacentMatrix)
		// verifyColoringConflicts(graph, adjacentMatrix, t)
		// fmt.Println("under coloringGraphWithSmallestOriginalDegreeLast")
		_, _, orderList := coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
		startColoring(graph, adjacentMatrix, orderList)
		// verifyColoringConflicts(graph, adjacentMatrix)
		// fmt.Println("under coloringGraphWithLargestOriginalDegreeLast")
		// coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
		// verifyColoringConflicts(graph, adjacentMatrix)
		// fmt.Println("under coloringGraphWithAscendingOrderingofVertexID")
		// coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix)
		// verifyColoringConflicts(graph, adjacentMatrix, t)
		// fmt.Println("under coloringGraphWithTheSmallestLastOrdering")
		// coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
		// coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
		verifyColoringConflicts(graph, adjacentMatrix)
	}
}

func Test_createCompleteGraphAndColoring(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	fmt.Println("Test_createCompleteGraph")
	// t.Log("Test_createCompleteGraph")
	graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(100)
	// coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix)
	coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
	verifyColoringConflicts(graph, adjacentMatrix)
	// t.Log("done")
	// graph, _, _ := createCompleteGraph(5000)
	// verifyRandomGraph(graph, "Test_createCompleteGraph")
}

// func verifyColoringConflicts(verticesList []*Vertex, adjacentMatrix [][]int, t *testing.T) {
// 	for i := 0; i < len(verticesList); i++ {
// 		for j := 0; j < len(verticesList); j++ {
// 			if adjacentMatrix[i][j] == 1 && verticesList[i].color == verticesList[j].color {
// 				t.Fatal("color conflicts exist")
// 			}
// 		}
// 	}
// }

func Test_initGraph(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	graph, _ := initGraph(10)
	fmt.Println(graph[0])
}

func Test_createCompleteGraph(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	fmt.Println("Test_createCompleteGraph")
	// t.Log("Test_createCompleteGraph")
	createCompleteGraph(10000)
	t.Log("done")
	// graph, _, _ := createCompleteGraph(5000)
	// verifyRandomGraph(graph, "Test_createCompleteGraph")
}

func Test_createCycleGraph(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	graph, _, _ := createCycleGraph(7)
	verifyRandomGraph(graph, "Test_createCycleGraph")
}

func Test_createRandomGraphWithUniformDistribution(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	fmt.Println(getFuncName())
	graph, _, _ := createRandomGraphWithUniformDistribution(500, 10000)
	verifyRandomGraph(graph, getFuncName())
}

func Test_createRandomGraphWithSkewedDistribution(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	graph, _, _ := createRandomGraphWithSkewedDistribution(500, 10000)
	verifyRandomGraph(graph, getFuncName())
}

func Test_createRandomGraphWithAmorDistribution(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	fmt.Println(getFuncName())
	graph, _, _ := createRandomGraphWithAmorDistribution(500, 10000)
	verifyRandomGraph(graph, getFuncName())
}

func Test_coloringGraphWithTheSmallestLastOrdering(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	fmt.Println(getFuncName())
	graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithSkewedDistribution(20, 120)
	// graph, _, _ := createRandomGraphWithSkewedDistribution(20, 120)
	outPutGraph(graph, "createRandomGraphWithSkewedDistribution")
	coloringGraphWithTheSmallestLastOrderingV1(graph, adjacentMatrix, sameCurrentDegreeList)
	verifyColoringConflicts(graph, adjacentMatrix)
}

func Test_coloringGraphWithTheLargestLastOrdering(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	fmt.Println(getFuncName())
	graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithSkewedDistribution(6, 10)
	outPutGraph(graph, "createRandomGraphWithSkewedDistribution")
	coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
}

func Test_coloringGraphWithRandomOrdering(t *testing.T) {
	if TestArr[getFuncName()] != 1 {
		t.Skip()
	}
	fmt.Println(getFuncName())
	graph, adjacentMatrix, _ := createRandomGraphWithSkewedDistribution(6, 10)
	coloringGraphWithRandomOrdering(graph, adjacentMatrix)
}

func Test_tt(t *testing.T) {
	// t.Skip()
	// trace2()
	ex, _ := os.Getwd()
	fmt.Println(ex)
}

func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
	// return frame.Func
}
