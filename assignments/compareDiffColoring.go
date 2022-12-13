package assignments

import (
	"fmt"
	"sync"
	"time"
)

func Run2CompleteColoring(iMinimum int, iMaximum int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		tagName := fmt.Sprintf("Run2CompleteColoring")
		// tagName += "V1_"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
			_, _, orderList := coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	wg.Wait()
}

func Run2UniformColoring(density int, iMinimum int, iMaximum int) {
	var wg sync.WaitGroup
	wg.Add(1)

	// _, adjacentMatrix, _ := createRandomGraphWithUniformDistribution(3000,)
	// // _, adjacentMatrix, _ := createRandomGraphWithUniformDistribution(i, i*(i-1)/100)
	// var points []*Point
	// for k, v := range adjacentMatrix {
	// 	for k1, v1 := range v {
	// 		if v1 == 1 {
	// 			points = append(points, &Point{k + 1, k1 + 1})
	// 		}
	// 	}
	// }

	go func() {
		tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
		tagName += "V1_"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
			_, _, orderList := coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()

	// go func() {
	// 	tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
	// 	tagName += "V2_"
	// 	for i := iMinimum; i <= iMaximum; i++ {
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
	// 		_, _, orderList := coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
	// 		Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
	// 	tagName += "V3_"
	// 	for i := iMinimum; i <= iMaximum; i++ {
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
	// 		RunV3coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 	}
	// 	// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 	wg.Done()
	// }()

	// go func() {
	// 	tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
	// 	tagName += "V4_"
	// 	for i := iMinimum; i <= iMaximum; i++ {
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
	// 		RunV4coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 	}
	// 	// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 	wg.Done()
	// }()

	// go func() {
	// 	tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
	// 	tagName += "V5_"
	// 	for i := iMinimum; i <= iMaximum; i++ {
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
	// 		RunV5coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 	}
	// 	// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 	wg.Done()
	// }()

	// go func() {
	// 	tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
	// 	tagName += "V6_"
	// 	for i := iMinimum; i <= iMaximum; i++ {
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
	// 		RunV6coloringGraphWithRandomOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 	}
	// 	// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 	wg.Done()
	// }()

	wg.Wait()
}

//graphType 1 complete graph 2 cycle graph 3 random uniform graph 4 random skewed graph 5 random amor graph
func compareDiffColoringGraph(v int, graphType int, density int) {
	var adjacentMatrix [][]int
	var tagName string
	switch graphType {
	case 1:
		tagName = "completeGraph"
		_, adjacentMatrix, _ = createCompleteGraph(v)
	case 2:
		tagName = "cycleGraph"
		_, adjacentMatrix, _ = createCycleGraph(v)
	case 3:
		tagName = "uniformGraph"
		_, adjacentMatrix, _ = createRandomGraphWithUniformDistribution(v, v*(v-1)/2/density)
	case 4:
		tagName = "skewedGraph"
		_, adjacentMatrix, _ = createRandomGraphWithSkewedDistribution(v, v*(v-1)/2/density)
	case 5:
		tagName = "amorGraph"
		_, adjacentMatrix, _ = createRandomGraphWithAmorDistribution(v, v*(v-1)/2/density)
	default:
		panic("input type is illegal")
	}

	var points []*Point
	for k, v := range adjacentMatrix {
		for k1, v1 := range v {
			if v1 == 1 {
				points = append(points, &Point{k + 1, k1 + 1})
			}
		}
	}
	var wg sync.WaitGroup
	wg.Add(6)

	func() {
		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(v, points)
		start1 := time.Now()
		_, _, orderList := coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
		startColoring(graph, adjacentMatrix, orderList)
		end1 := time.Since(start1).Microseconds()
		go func() {
			markColoringInfo(graph, orderList, 1, int(end1), tagName)
			wg.Done()
		}()
	}()

	func() {
		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(v, points)
		start1 := time.Now()
		_, _, orderList := coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList)
		startColoring(graph, adjacentMatrix, orderList)
		end1 := time.Since(start1).Microseconds()
		go func() {
			markColoringInfo(graph, orderList, 2, int(end1), tagName)
			wg.Done()
		}()
	}()

	func() {
		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(v, points)
		start1 := time.Now()
		_, _, orderList := coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList)
		startColoring(graph, adjacentMatrix, orderList)
		end1 := time.Since(start1).Microseconds()
		go func() {
			markColoringInfo(graph, orderList, 3, int(end1), tagName)
			wg.Done()
		}()
	}()

	func() {
		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(v, points)
		start1 := time.Now()
		_, _, orderList := coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList)
		startColoring(graph, adjacentMatrix, orderList)
		end1 := time.Since(start1).Microseconds()
		go func() {
			markColoringInfo(graph, orderList, 4, int(end1), tagName)
			wg.Done()
		}()
	}()

	func() {
		graph, adjacentMatrix, _ := createCustomGraph(v, points)
		start1 := time.Now()
		_, _, orderList := coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix)
		startColoring(graph, adjacentMatrix, orderList)
		end1 := time.Since(start1).Microseconds()
		go func() {
			markColoringInfo(graph, orderList, 5, int(end1), tagName)
			wg.Done()
		}()
	}()

	func() {
		graph, adjacentMatrix, _ := createCustomGraph(v, points)
		start1 := time.Now()
		_, _, orderList := coloringGraphWithRandomOrdering(graph, adjacentMatrix)
		startColoring(graph, adjacentMatrix, orderList)
		end1 := time.Since(start1).Microseconds()
		go func() {
			markColoringInfo(graph, orderList, 6, int(end1), tagName)
			wg.Done()
		}()
	}()
	wg.Wait()
}

