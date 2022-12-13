package assignments

import (
	"fmt"
	"sync"
	"time"
)

func RunProjectMain() {

	// points := []*Point{
	// 	{x: 1, y: 5},
	// 	{x: 1, y: 3},
	// 	{x: 1, y: 2},
	// 	{x: 2, y: 6},
	// 	{x: 2, y: 4},
	// 	{x: 3, y: 5},
	// 	{x: 3, y: 4},
	// 	{x: 3, y: 6},
	// 	{x: 4, y: 6},
	// 	{x: 5, y: 6},
	// 	{x: 5, y: 1},
	// 	{x: 3, y: 1},
	// 	{x: 2, y: 1},
	// 	{x: 6, y: 2},
	// 	{x: 4, y: 2},
	// 	{x: 5, y: 3},
	// 	{x: 4, y: 3},
	// 	{x: 6, y: 3},
	// 	{x: 6, y: 4},
	// 	{x: 6, y: 5},
	// }
	// // graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithSkewedDistribution(10, 20)
	// graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(6, points)
	// outPutGraph(graph, "createCustomGraph")
	// _, _, orderList := coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
	// startColoring(graph, adjacentMatrix, orderList)

	// runCompleteAndCycleGraphGenerator()
	// runCompleteAndCycleGraphGenerator()
	// runRandomGraphGeneratorUniform()
	// runRandomGraphGeneratorSkewed()
	// runRandomGraphGeneratorAmor()

	// points := []*Point{
	// 	{x: 1, y: 2},
	// 	{x: 1, y: 4},
	// 	{x: 1, y: 7},
	// 	{x: 1, y: 10},
	// 	{x: 2, y: 1},
	// 	{x: 2, y: 3},
	// 	{x: 2, y: 4},
	// 	{x: 2, y: 5},
	// 	{x: 2, y: 6},
	// 	{x: 2, y: 7},
	// 	{x: 2, y: 10},
	// 	{x: 3, y: 2},
	// 	{x: 3, y: 7},
	// 	{x: 4, y: 1},
	// 	{x: 4, y: 2},
	// 	{x: 4, y: 7},
	// 	{x: 4, y: 8},
	// 	{x: 4, y: 10},
	// 	{x: 5, y: 2},
	// 	{x: 6, y: 2},
	// 	{x: 6, y: 7},
	// 	{x: 6, y: 10},
	// 	{x: 7, y: 1},
	// 	{x: 7, y: 2},
	// 	{x: 7, y: 3},
	// 	{x: 7, y: 4},
	// 	{x: 7, y: 6},
	// 	{x: 7, y: 8},
	// 	{x: 7, y: 9},
	// 	{x: 7, y: 10},
	// 	{x: 8, y: 4},
	// 	{x: 8, y: 7},
	// 	{x: 8, y: 10},
	// 	{x: 9, y: 7},
	// 	{x: 10, y: 1},
	// 	{x: 10, y: 2},
	// 	{x: 10, y: 4},
	// 	{x: 10, y: 6},
	// 	{x: 10, y: 7},
	// 	{x: 10, y: 8},
	// }

	// graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(5, points)
	// graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithSkewedDistribution(20, 90)
	// graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithSkewedDistribution(5, 5)
	// graph, _, _ := createRandomGraphWithSkewedDistribution(20, 120)
	// outPutGraph(graph, "createRandomGraphWithSkewedDistribution")
	// coloringGraphWithTheSmallestLastOrderingV1(graph, adjacentMatrix, sameCurrentDegreeList)
	// verifyColoringConflicts(graph, adjacentMatrix)

	// for i := 0; i < 10000; i++ {
	// 	// fmt.Printf("the %d times :\n", i)
	// 	// graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(20, 150)
	// 	// outPutGraph(graph, "createRandomGraphWithSkewedDistribution")
	// 	// for j := 0; j < 10; j++ {
	// 	// 	for k := 0; k < 10; k++ {
	// 	// 		if adjacentMatrix[j][k] == 1 {
	// 	// 			fmt.Printf("{x:%d,y:%d},\n", j+1, k+1)
	// 	// 		}
	// 	// 	}
	// 	// }
	// 	// coloringGraphWithTheSmallestLastOrderingV1(graph, adjacentMatrix, sameCurrentDegreeList)
	// 	// verifyColoringConflicts(graph, adjacentMatrix)
	// }

	// graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithSkewedDistribution(6, 10)
	// outPutGraph(graph, "createRandomGraphWithSkewedDistribution")

	// TestRunCompleteAndCycleCurrency(100, 1000)

	// TestRunCompleteAndCycleCurrency(10000, 10000)

	// for i := 10; i < 1000; i++ {
	// 	tagName := "Run2coloringGraphWithTheSmallestLastOrdering"
	// 	graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
	// 	Run2coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// }

	// for i := 10; i < 1000; i++ {
	// 	tagName := "Run2coloringGraphWithTheSmallestLastOrdering"
	// 	graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
	// 	Run2coloringGraphWithTheSmallestLastOrderingV1(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// }

	// graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(6)
	// outPutGraph(graph, "createCompleteGraph")
	// _, _, orderList := coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
	// startColoring(graph, adjacentMatrix, orderList)
	// verifyColoringConflicts(graph, adjacentMatrix)

	// Run2CompleteOrdering(100, 2500)
	// Run2CycleOrdering(100, 1000)
	// Run2UniformColoring(2, 100, 2500)
	// Run2CompleteColoring(100, 2500)
	// Run2UniformOrdering(2, 100, 2000)
	// Run2UniformOrdering(25, 100, 1000)
	// Run2UniformOrdering(2, 100, 1000)
	// Run2SkewedOrdering(50, 100, 1000)
	// Run2SkewedOrdering(25, 100, 1000)
	// Run2SkewedOrdering(2, 100, 1000)
	// Run2AmorOrdering(50, 100, 1000)
	// Run2AmorOrdering(25, 100, 1000)
	// Run2AmorOrdering(2, 100, 1000)

	// for i := 100; i <= 1000; i += 100 {
	// 	compareDiffColoringGraph(i, 1, 1)
	// 	compareDiffColoringGraph(i, 2, 1)

	// 	compareDiffColoringGraph(i, 3, 50)
	// 	compareDiffColoringGraph(i, 4, 50)
	// 	compareDiffColoringGraph(i, 5, 50)

	// 	compareDiffColoringGraph(i, 3, 4)
	// 	compareDiffColoringGraph(i, 4, 4)
	// 	compareDiffColoringGraph(i, 5, 4)
	// }
	var wg sync.WaitGroup
	wg.Add(6)

	go func() {
		compareDiffColoringGraph(8000, 1, 1)
		wg.Done()
	}()

	go func() {
		compareDiffColoringGraph(10000, 1, 1)
		wg.Done()
	}()

	go func() {
		compareDiffColoringGraph(8000, 2, 1)
		wg.Done()
	}()

	go func() {
		compareDiffColoringGraph(10000, 2, 1)
		wg.Done()
	}()

	go func() {
		compareDiffColoringGraph(10000, 3, 20)
		wg.Done()
	}()

	go func() {
		compareDiffColoringGraph(8000, 3, 20)
		wg.Done()
	}()

	// for i := 5000; i <= 10000; i += 1000 {
	// 	compareDiffColoringGraph(i, 1, 1)
	// 	compareDiffColoringGraph(i, 2, 1)

	// 	compareDiffColoringGraph(i, 3, 4)
	// 	compareDiffColoringGraph(i, 4, 4)
	// 	compareDiffColoringGraph(i, 5, 4)

	// }
	wg.Wait()

}

func RunV1coloringGraphWithTheSmallestLastOrdering(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex, tagName string) []int {
	start1 := time.Now()
	_, _, orderList := coloringGraphWithTheSmallestLastOrdering(verticesList, adjacentMatrix, sameCurrentDegreeList)
	end1 := time.Since(start1).Microseconds()
	writeFile(fmt.Sprintf("%s_%s", tagName, getFuncName()), fmt.Sprintf("%d\t%d\t%d\n", len(verticesList), getEdgesFromGraph(verticesList), end1))
	return orderList
}

func Run2coloringGraphWithTheSmallestLastOrderingV1(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex, tagName string) []int {
	start1 := time.Now()
	_, _, orderList := coloringGraphWithTheSmallestLastOrderingV1(verticesList, adjacentMatrix, sameCurrentDegreeList)
	end1 := time.Since(start1).Microseconds()
	writeFile(fmt.Sprintf("%s_%s", tagName, getFuncName()), fmt.Sprintf("%d\t%d\t%d\n", len(verticesList), getEdgesFromGraph(verticesList), end1))
	return orderList
}

func RunV2coloringGraphWithTheLargestLastOrdering(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex, tagName string) []int {
	start1 := time.Now()
	_, _, orderList := coloringGraphWithTheLargestLastOrdering(verticesList, adjacentMatrix, sameCurrentDegreeList)
	end1 := time.Since(start1).Microseconds()
	writeFile(fmt.Sprintf("%s_%s", tagName, getFuncName()), fmt.Sprintf("%d\t%d\t%d\n", len(verticesList), getEdgesFromGraph(verticesList), end1))
	return orderList
}

func RunV3coloringGraphWithSmallestOriginalDegreeLast(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex, tagName string) []int {
	start1 := time.Now()
	_, _, orderList := coloringGraphWithSmallestOriginalDegreeLast(verticesList, adjacentMatrix, sameCurrentDegreeList)
	end1 := time.Since(start1).Microseconds()
	writeFile(fmt.Sprintf("%s_%s", tagName, getFuncName()), fmt.Sprintf("%d\t%d\t%d\n", len(verticesList), getEdgesFromGraph(verticesList), end1))
	return orderList
}

func RunV4coloringGraphWithLargestOriginalDegreeLast(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex, tagName string) []int {
	start1 := time.Now()
	_, _, orderList := coloringGraphWithLargestOriginalDegreeLast(verticesList, adjacentMatrix, sameCurrentDegreeList)
	end1 := time.Since(start1).Microseconds()
	writeFile(fmt.Sprintf("%s_%s", tagName, getFuncName()), fmt.Sprintf("%d\t%d\t%d\n", len(verticesList), getEdgesFromGraph(verticesList), end1))
	return orderList
}

func RunV5coloringGraphWithAscendingOrderingofVertexID(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex, tagName string) []int {
	start1 := time.Now()
	_, _, orderList := coloringGraphWithAscendingOrderingofVertexID(verticesList, adjacentMatrix)
	end1 := time.Since(start1).Microseconds()
	writeFile(fmt.Sprintf("%s_%s", tagName, getFuncName()), fmt.Sprintf("%d\t%d\t%d\n", len(verticesList), getEdgesFromGraph(verticesList), end1))
	return orderList
}

func RunV6coloringGraphWithRandomOrdering(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex, tagName string) []int {
	start1 := time.Now()
	_, _, orderList := coloringGraphWithRandomOrdering(verticesList, adjacentMatrix)
	end1 := time.Since(start1).Microseconds()
	writeFile(fmt.Sprintf("%s_%s", tagName, getFuncName()), fmt.Sprintf("%d\t%d\t%d\n", len(verticesList), getEdgesFromGraph(verticesList), end1))
	return orderList
}

func Run2startColoring(verticesList []*Vertex, adjacentMatrix [][]int, orderList []int, tagName string) {
	start1 := time.Now()
	startColoring(verticesList, adjacentMatrix, orderList)
	end1 := time.Since(start1).Microseconds()
	writeFile(fmt.Sprintf("%s_%s", tagName, getFuncName()), fmt.Sprintf("%d\t%d\t%d\n", len(verticesList), getEdgesFromGraph(verticesList), end1))
}

func getEdgesFromGraph(verticesList []*Vertex) int {
	count := 0
	for _, v := range verticesList {
		count += v.oDegree
	}
	return count / 2
}

func runCompleteAndCycleGraphGenerator() {
	from := 100
	end := 10000
	chs := make([]chan int, 2)
	chs[0] = make(chan int)
	chs[1] = make(chan int)
	go Run2createCompleteGraph(from, end, chs[0])
	go Run2createCycleGraph(from, end, chs[1])
	for _, ch := range chs {
		<-ch
	}
	fmt.Println("done")
}

func runRandomGraphGeneratorUniform() {
	from := 100
	end := 3000
	chs := make([]chan int, 4)
	chs[0] = make(chan int)
	chs[1] = make(chan int)
	chs[2] = make(chan int)
	chs[3] = make(chan int)
	go Run2createRandomGraphWithUniformDistribution(from, end, 1, chs[0])
	go Run2createRandomGraphWithUniformDistribution(from, end, 2, chs[1])
	go Run2createRandomGraphWithUniformDistribution(from, end, 4, chs[2])
	go Run2createRandomGraphWithUniformDistribution(from, end, 50, chs[3])
	for _, ch := range chs {
		<-ch
	}
	fmt.Println("done")
}

func runRandomGraphGeneratorSkewed() {
	from := 100
	end := 3000
	chs := make([]chan int, 4)
	chs[0] = make(chan int)
	chs[1] = make(chan int)
	chs[2] = make(chan int)
	chs[3] = make(chan int)
	go Run2createRandomGraphWithSkewedDistribution(from, end, 1, chs[0])
	go Run2createRandomGraphWithSkewedDistribution(from, end, 2, chs[1])
	go Run2createRandomGraphWithSkewedDistribution(from, end, 4, chs[2])
	go Run2createRandomGraphWithSkewedDistribution(from, end, 50, chs[3])
	for _, ch := range chs {
		<-ch
	}
	fmt.Println("done")
}

func runRandomGraphGeneratorAmor() {
	from := 100
	end := 3000
	chs := make([]chan int, 4)
	chs[0] = make(chan int)
	chs[1] = make(chan int)
	chs[2] = make(chan int)
	chs[3] = make(chan int)
	go Run2createRandomGraphWithAmorDistribution(from, end, 1, chs[0])
	go Run2createRandomGraphWithAmorDistribution(from, end, 2, chs[1])
	go Run2createRandomGraphWithAmorDistribution(from, end, 4, chs[2])
	go Run2createRandomGraphWithAmorDistribution(from, end, 50, chs[3])
	for _, ch := range chs {
		<-ch
	}
	fmt.Println("done")
}

func Run2createCompleteGraph(from int, end int, ch chan int) {
	for from <= end {
		start1 := time.Now()
		createCompleteGraph(from)
		end1 := time.Since(start1).Microseconds()
		writeFile(fmt.Sprintf("BenchRun2createCompleteGraph"), fmt.Sprintf("%d\t%d\n", from, end1))
		from++
	}
	ch <- 1
}

func Run2createCycleGraph(from int, end int, ch chan int) {
	for from <= end {
		start1 := time.Now()
		createCycleGraph(from)
		end1 := time.Since(start1).Microseconds()
		writeFile(fmt.Sprintf("BenchRun2createCycleGraph"), fmt.Sprintf("%d\t%d\n", from, end1))
		from++
	}
	ch <- 1
}

func Run2createRandomGraphWithUniformDistribution(from int, end int, density int, ch chan int) {
	for from <= end {
		start1 := time.Now()
		createRandomGraphWithUniformDistribution(from, from*(from-1)/(2*density))
		end1 := time.Since(start1).Microseconds()
		writeFile(fmt.Sprintf("BenchRun2createRandomGraphWithUniformDistribution_den%d", density), fmt.Sprintf("%d\t%d\n", from, end1))
		from++
	}
	ch <- 1
}

func Run2createRandomGraphWithSkewedDistribution(from int, end int, density int, ch chan int) {
	for from <= end {
		start1 := time.Now()
		createRandomGraphWithAmorDistribution(from, from*(from-1)/(2*density))
		end1 := time.Since(start1).Microseconds()
		writeFile(fmt.Sprintf("BenchRun2createRandomGraphWithSkewedDistribution_den%d", density), fmt.Sprintf("%d\t%d\n", from, end1))
		from++
	}
	ch <- 1
}

func Run2createRandomGraphWithAmorDistribution(from int, end int, density int, ch chan int) {
	for from <= end {
		start1 := time.Now()
		createRandomGraphWithAmorDistribution(from, from*(from-1)/(2*density))
		end1 := time.Since(start1).Microseconds()
		writeFile(fmt.Sprintf("BenchRun2createRandomGraphWithAmorDistribution_den%d", density), fmt.Sprintf("%d\t%d\n", from, end1))
		from++
	}
	ch <- 1
}
