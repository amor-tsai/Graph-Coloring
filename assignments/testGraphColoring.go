package assignments

import (
	"fmt"
	"sync"
)

func Run2CompleteOrdering(iMinimum int, iMaximum int) {
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		tagName := "completeGraph"
		tagName += "_V6"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
			RunV6coloringGraphWithRandomOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "completeGraph"
		tagName += "_V5"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
			RunV5coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "completeGraph"
		tagName += "_V4"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
			RunV4coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "completeGraph"
		tagName += "_V3"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
			RunV3coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "completeGraph"
		tagName += "_V2"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
			RunV2coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "completeGraph"
		tagName += "_V1"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCompleteGraph(i)
			RunV1coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	wg.Wait()
}

func Run2CycleOrdering(iMinimum int, iMaximum int) {
	var wg sync.WaitGroup
	wg.Add(6)
	go func() {
		tagName := "cycleGraph"
		tagName += "_V6"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCycleGraph(i)
			RunV6coloringGraphWithRandomOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "cycleGraph"
		tagName += "_V5"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCycleGraph(i)
			RunV5coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "cycleGraph"
		tagName += "_V4"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCycleGraph(i)
			RunV4coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "cycleGraph"
		tagName += "_V3"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCycleGraph(i)
			RunV3coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "cycleGraph"
		tagName += "_V2"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCycleGraph(i)
			RunV2coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	go func() {
		tagName := "cycleGraph"
		tagName += "_V1"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createCycleGraph(i)
			RunV1coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()
	wg.Wait()
}

func Run2UniformOrdering(density int, iMinimum int, iMaximum int) {
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

	// go func() {
	// 	tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
	// 	tagName += "V1_"
	// 	for i := iMinimum; i <= iMaximum; i++ {
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
	// 		RunV1coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 	}
	// 	// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 	wg.Done()
	// }()

	go func() {
		tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
		tagName += "V2_"
		for i := iMinimum; i <= iMaximum; i++ {
			graph, adjacentMatrix, sameCurrentDegreeList := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
			RunV2coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// graph,adjacentMatrix,orderList = coloringGraphWithTheLargestLastOrdering(graph,adjacentMatrix,sameCurrentDegreeList)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
		}
		wg.Done()
	}()

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

	// for i := iMinimum; i <= iMaximum; i++ {
	// 	tagName := fmt.Sprintf("randomGraphUniform_density%d_", density)
	// 	c := make(chan int, 6)
	// 	_, adjacentMatrix, _ := createRandomGraphWithUniformDistribution(i, i*(i-1)/2/density)
	// 	// _, adjacentMatrix, _ := createRandomGraphWithUniformDistribution(i, i*(i-1)/100)
	// 	var points []*Point
	// 	for k, v := range adjacentMatrix {
	// 		for k1, v1 := range v {
	// 			if v1 == 1 {
	// 				points = append(points, &Point{k + 1, k1 + 1})
	// 			}
	// 		}
	// 	}

	// 	go func(i int, points []*Point, tagName string) {
	// 		tagName += "V5_"
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
	// 		RunV5coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 		// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 		c <- 1
	// 	}(i, points, tagName)

	// 	go func(i int, points []*Point, tagName string) {
	// 		tagName += "V6_"
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
	// 		RunV6coloringGraphWithRandomOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 		// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 		c <- 2
	// 	}(i, points, tagName)

	// 	go func(i int, points []*Point, tagName string) {
	// 		tagName += "V3_"
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
	// 		RunV3coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 		// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 		c <- 3
	// 	}(i, points, tagName)

	// 	go func(i int, points []*Point, tagName string) {
	// 		tagName += "V4_"
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
	// 		RunV4coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 		// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 		c <- 4
	// 	}(i, points, tagName)

	// 	go func(i int, points []*Point, tagName string) {
	// 		tagName += "V1_"
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
	// 		RunV1coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 		// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 		c <- 5
	// 	}(i, points, tagName)

	// 	go func(i int, points []*Point, tagName string) {
	// 		tagName += "V2_"
	// 		graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
	// 		RunV2coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
	// 		// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
	// 		c <- 6
	// 	}(i, points, tagName)

	// 	for i := 0; i < 6; i++ {
	// 		<-c
	// 	}
	// }
}

func Run2SkewedOrdering(density int, iMinimum int, iMaximum int) {
	for i := iMinimum; i <= iMaximum; i++ {
		tagName := fmt.Sprintf("randomGraphSkewed_density%d_", density)
		c := make(chan int, 6)
		_, adjacentMatrix, _ := createRandomGraphWithSkewedDistribution(i, i*(i-1)/2/density)
		// _, adjacentMatrix, _ := createRandomGraphWithUniformDistribution(i, i*(i-1)/100)
		var points []*Point
		for k, v := range adjacentMatrix {
			for k1, v1 := range v {
				if v1 == 1 {
					points = append(points, &Point{k + 1, k1 + 1})
				}
			}
		}

		go func(i int, points []*Point, tagName string) {
			tagName += "V5_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV5coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 1
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V6_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV6coloringGraphWithRandomOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 2
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V3_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV3coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 3
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V4_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV4coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 4
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V1_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV1coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 5
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V2_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV2coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 6
		}(i, points, tagName)

		for i := 0; i < 6; i++ {
			<-c
		}
	}
}

func Run2AmorOrdering(density int, iMinimum int, iMaximum int) {

	for i := iMinimum; i <= iMaximum; i++ {
		tagName := fmt.Sprintf("randomGraphAmor_density%d_", density)
		c := make(chan int, 6)
		_, adjacentMatrix, _ := createRandomGraphWithAmorDistribution(i, i*(i-1)/2/density)
		// _, adjacentMatrix, _ := createRandomGraphWithUniformDistribution(i, i*(i-1)/100)
		var points []*Point
		for k, v := range adjacentMatrix {
			for k1, v1 := range v {
				if v1 == 1 {
					points = append(points, &Point{k + 1, k1 + 1})
				}
			}
		}

		go func(i int, points []*Point, tagName string) {
			tagName += "V5_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV5coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 1
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V6_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV6coloringGraphWithRandomOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 2
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V3_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV3coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 3
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V4_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV4coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 4
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V1_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV1coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 5
		}(i, points, tagName)

		go func(i int, points []*Point, tagName string) {
			tagName += "V2_"
			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
			RunV2coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
			// Run2startColoring(graph, adjacentMatrix, orderList, tagName)
			c <- 6
		}(i, points, tagName)

		for i := 0; i < 6; i++ {
			<-c
		}
	}
}

func verifyColoringConflicts(verticesList []*Vertex, adjacentMatrix [][]int) {
	for i := 0; i < len(verticesList); i++ {
		for j := 0; j < len(verticesList); j++ {
			if adjacentMatrix[i][j] == 1 && verticesList[i].color == verticesList[j].color {
				panic("color conflicts exist")
			}
		}
	}
}

// func TestRunCompleteAndCycleCurrency(iMinimum int, iMaximum int) {

// 	for i := iMinimum; i <= iMaximum; i++ {
// 		c := make(chan int, 6)
// 		tagName := "completeGraph"
// 		_, adjacentMatrix, _ := createCompleteGraph(i)
// 		var points []*Point
// 		for k, v := range adjacentMatrix {
// 			for k1, v1 := range v {
// 				if v1 == 1 {
// 					points = append(points, &Point{k + 1, k1 + 1})
// 				}
// 			}
// 		}

// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V6"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV6coloringGraphWithRandomOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 1
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V5"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV5coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 2
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V4"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV4coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 3
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V3"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV3coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 4
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V2"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV2coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 5
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V1"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV1coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 6
// 		}(i, points, tagName)

// 		for i := 0; i < 6; i++ {
// 			<-c
// 		}
// 	}

// 	for i := iMinimum; i <= iMaximum; i++ {
// 		c := make(chan int, 6)
// 		tagName := "cycleGraph"
// 		_, adjacentMatrix, _ := createCycleGraph(i)
// 		var points []*Point
// 		for k, v := range adjacentMatrix {
// 			for k1, v1 := range v {
// 				if v1 == 1 {
// 					points = append(points, &Point{k + 1, k1 + 1})
// 				}
// 			}
// 		}

// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V6"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV6coloringGraphWithRandomOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 1
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V5"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV5coloringGraphWithAscendingOrderingofVertexID(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 2
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V4"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV4coloringGraphWithLargestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 3
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V3"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV3coloringGraphWithSmallestOriginalDegreeLast(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 4
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V2"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV2coloringGraphWithTheLargestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 5
// 		}(i, points, tagName)
// 		go func(i int, points []*Point, tagName string) {
// 			tagName += "_V1"
// 			graph, adjacentMatrix, sameCurrentDegreeList := createCustomGraph(i, points)
// 			orderList := RunV1coloringGraphWithTheSmallestLastOrdering(graph, adjacentMatrix, sameCurrentDegreeList, tagName)
// 			Run2startColoring(graph, adjacentMatrix, orderList, tagName)
// 			c <- 6
// 		}(i, points, tagName)

// 		for i := 0; i < 6; i++ {
// 			<-c
// 		}
// 	}
// }
