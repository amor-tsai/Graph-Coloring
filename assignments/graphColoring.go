package assignments

import "fmt"

var TestOutput int = 0

func countSameCurrentDegreeList(sameCurrentDegreeList []*Vertex) int {
	count := 0
	for i := 0; i < len(sameCurrentDegreeList); i++ {
		currNode := sameCurrentDegreeList[i]
		for currNode != nil {
			count++
			currNode = currNode.next
		}
	}
	return count
}

//Smallest Last Vertex Ordering
func coloringGraphWithTheSmallestLastOrdering(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex) ([]*Vertex, [][]int, []int) {
	oLength := len(verticesList)
	orderList := make([]int, oLength)
	count := oLength - 1
	i := 0
	for count >= 0 {
		// fmt.Printf("count is %d the %d loop\n", count, i)
		if sameCurrentDegreeList[i] != nil {
			//pick a vetex to be removed
			pickVertexIndex := sameCurrentDegreeList[i].value - 1
			//push on stack
			orderList[count] = verticesList[pickVertexIndex].value
			//remove from dll
			if nil != sameCurrentDegreeList[i].next {
				sameCurrentDegreeList[i] = sameCurrentDegreeList[i].next
				sameCurrentDegreeList[i].last = nil
			} else {
				sameCurrentDegreeList[i] = nil
			}
			//record when degree deleted
			verticesList[pickVertexIndex].markDelete()
			//decrement degrees from its adjacent vertices
			currNode := verticesList[pickVertexIndex].head
			for currNode != nil {
				verticesList[currNode.value-1].degree--
				currNode = currNode.next
			}
			//move adjacent vertices in the dll
			count--
			sameCurrentDegreeList = initSDgreeList(verticesList) //this is rebuilding the same degree list
			i = 0
		} else {
			i++
		}
	}
	return verticesList, adjacentMatrix, orderList
}

//Largest Last Vertex Ordering
func coloringGraphWithTheLargestLastOrdering(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex) ([]*Vertex, [][]int, []int) {
	oLength := len(verticesList)
	orderList := make([]int, oLength)
	count := oLength - 1
	for i := oLength - 1; i >= 0; i-- {
		if sameCurrentDegreeList[i] != nil {
			//pick a vetex to be removed
			pickVertexIndex := sameCurrentDegreeList[i].value - 1
			//push on stack
			// reservedList[count] = verticesList[pickVertexIndex]
			orderList[count] = verticesList[pickVertexIndex].value
			//remove from dll
			if nil != sameCurrentDegreeList[i].next {
				sameCurrentDegreeList[i] = sameCurrentDegreeList[i].next
				sameCurrentDegreeList[i].last = nil
			} else {
				sameCurrentDegreeList[i] = nil
			}
			//record when degree deleted
			// reservedList[count-1].markDelete()
			verticesList[pickVertexIndex].markDelete()
			// reservedList[count-1].degreeWhenDelete = reservedList[count-1].degree
			currNode := verticesList[pickVertexIndex].head
			for currNode != nil {
				//decrease the degree of its adjacent vertices
				verticesList[currNode.value-1].degree--
				currNode = currNode.next
			}
			//move adjacent vertices in the dll
			count--
			if count < 0 {
				break
			}
			sameCurrentDegreeList = initSDgreeList(verticesList)
			i = oLength
		}
	}
	//graph coloring part
	//TODO need to optimize the loop to reduce the times it looped
	// startColoring(verticesList, adjacentMatrix, orderList)
	return verticesList, adjacentMatrix, orderList
}

//Smallest Original Degree Last
func coloringGraphWithSmallestOriginalDegreeLast(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex) ([]*Vertex, [][]int, []int) {
	oLength := len(verticesList)
	orderList := make([]int, oLength)
	count := oLength - 1
	for i := 0; i < oLength; i++ {
		currNode := sameCurrentDegreeList[i]
		for currNode != nil {
			orderList[count] = currNode.value
			count--
			currNode = currNode.next
		}
	}
	// startColoring(verticesList, adjacentMatrix, orderList)
	return verticesList, adjacentMatrix, orderList
}

//largest Original Degree Last
func coloringGraphWithLargestOriginalDegreeLast(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex) ([]*Vertex, [][]int, []int) {
	oLength := len(verticesList)
	orderList := make([]int, oLength)
	count := oLength - 1
	for i := oLength - 1; i >= 0; i-- {
		currNode := sameCurrentDegreeList[i]
		for currNode != nil {
			orderList[count] = currNode.value
			count--
			currNode = currNode.next
		}
	}
	return verticesList, adjacentMatrix, orderList
	// startColoring(verticesList, adjacentMatrix, orderList)
}

//ascending order of vertex ID
func coloringGraphWithAscendingOrderingofVertexID(verticesList []*Vertex, adjacentMatrix [][]int) ([]*Vertex, [][]int, []int) {
	oLength := len(verticesList)
	orderList := make([]int, oLength)
	for i := 0; i < oLength; i++ {
		orderList[i] = verticesList[i].value
	}
	// startColoring(verticesList, adjacentMatrix, orderList)
	return verticesList, adjacentMatrix, orderList
}

//uniform random ordering
func coloringGraphWithRandomOrdering(verticesList []*Vertex, adjacentMatrix [][]int) ([]*Vertex, [][]int, []int) {
	oLength := len(verticesList)
	orderList := randomNumberSequenceGenerator(oLength, oLength, Uniform)
	return verticesList, adjacentMatrix, orderList
}

//given sequence, start coloring through its order
func startColoring(verticesList []*Vertex, adjacentMatrix [][]int, orderList []int) {
	oLength := len(verticesList)
	for i := 0; i < oLength; i++ {
		currNode := verticesList[orderList[i]-1].head
		colorSlot := make([]int, oLength) //to guarantee every vertex has different color
		for currNode != nil {
			if 0 != verticesList[currNode.value-1].color {
				colorSlot[verticesList[currNode.value-1].color-1] = 1
			}
			currNode = currNode.next
		}
		for k := 0; k < oLength; k++ {
			if colorSlot[k] != 1 {
				verticesList[orderList[i]-1].color = k + 1
				break
			}
		}
	}
	if TestOutput == 1 {
		maxColor := 0
		maxDegreeWhenDeleted := 0
		terminalCliqueSize := 0
		for key, v := range orderList {
			if maxColor < verticesList[v-1].color {
				maxColor = verticesList[v-1].color
			}
			if maxDegreeWhenDeleted < verticesList[v-1].degreeWhenDelete {
				maxDegreeWhenDeleted = verticesList[v-1].degreeWhenDelete
			}
			if key == verticesList[v-1].degreeWhenDelete && verticesList[v-1].degreeWhenDelete-1 == terminalCliqueSize {
				terminalCliqueSize++
			}
			fmt.Println("vertex " + fmt.Sprint(verticesList[v-1].value) + " orginal degree " + fmt.Sprint(verticesList[v-1].oDegree) + " degree when deleted " + fmt.Sprint(verticesList[v-1].degreeWhenDelete) + " color " + fmt.Sprint(verticesList[v-1].color))
		}
		fmt.Printf("number of color used: %d \n", maxColor)
		fmt.Printf("the maximum degree when deleted: %d\n", maxDegreeWhenDeleted)
		fmt.Printf("the size of terminal clique: %d\n", terminalCliqueSize+1)
	}
}

func markColoringInfo(verticesList []*Vertex, orderList []int, algo int, timeUsed int, fileName string) {

	maxColor := 0
	maxDegree := 0
	// maxDegreeWhenDeleted := 0
	// terminalCliqueSize := 0
	totalDegree := 0
	v := len(orderList)
	for _, v := range orderList {
		if maxColor < verticesList[v-1].color {
			maxColor = verticesList[v-1].color
		}
		if maxDegree < verticesList[v-1].oDegree {
			maxDegree = verticesList[v-1].oDegree
		}
		totalDegree += verticesList[v-1].oDegree
	}
	writeFile(fileName, fmt.Sprintf("%d\t%d\t%d\t%d\t%d\t%d\t%d\n", v, totalDegree/2, totalDegree/v, maxDegree, algo, maxColor, timeUsed))
}

//given sequence, start coloring through its order
func startColoringV1(verticesList []*Vertex, adjacentMatrix [][]int, orderList []int) {
	oLength := len(verticesList)
	defaultColor := 1
	corloredV := make([]int, oLength)
	for i := 0; i < oLength; i++ {
		if corloredV[orderList[i]-1] == 1 {
			continue
		}
		verticesList[orderList[i]-1].color = defaultColor
		corloredV[orderList[i]-1] = 1
		for k := 0; k < oLength; k++ {
			if k != orderList[i]-1 {
				if adjacentMatrix[orderList[i]-1][k] != 1 {
					verticesList[i].color = defaultColor
					corloredV[i] = 1
				}
			}
		}
		defaultColor++
	}
	if TestOutput == 1 {
		for _, v := range orderList {
			fmt.Println("vertex " + fmt.Sprint(verticesList[v-1].value) + " orginal degree " + fmt.Sprint(verticesList[v-1].oDegree) + " degree when deleted " + fmt.Sprint(verticesList[v-1].degreeWhenDelete) + " color " + fmt.Sprint(verticesList[v-1].color))
		}
	}
}

func coloringGraphWithTheSmallestLastOrderingSlow(verticesList []*Vertex, adjacentMatrix [][]int) {
	var minDegree, removeVertexIndex int
	defaultColor := 1
	oLength := len(verticesList)
	orderList := make([]int, oLength)
	reservedList, _ := initGraph(oLength)
	for i := 0; i < oLength; i++ {
		verticesListLength := len(verticesList)
		for j := 0; j < verticesListLength; j++ {
			if j == 0 {
				minDegree = verticesList[j].degree
				removeVertexIndex = j
			} else {
				if minDegree > verticesList[j].degree {
					minDegree = verticesList[j].degree
					removeVertexIndex = j
				}
			}
		}
		fmt.Println("removeVertex " + fmt.Sprint(verticesList[removeVertexIndex].value))
		orderList[i] = verticesList[removeVertexIndex].value
		reservedList[i] = verticesList[removeVertexIndex]
		//remove vertex
		for k := 0; k < len(verticesList); k++ {
			if k == removeVertexIndex {
				continue
			}
			if adjacentMatrix[verticesList[k].value-1][verticesList[removeVertexIndex].value-1] == 1 {
				verticesList[k].degreeDecrement()
			}
		}
		verticesList[removeVertexIndex] = verticesList[verticesListLength-1]
		verticesList = verticesList[:verticesListLength-1]
		// verifyRandomGraph(verticesList)
	}
	for i := 0; i < oLength; i++ {
		if i == 0 {
			reservedList[i].color = defaultColor
			continue
		} else {
			takeColor := 0
			for j := 0; j < i; j++ {
				if adjacentMatrix[reservedList[i].value-1][reservedList[j].value-1] != 1 {
					if takeColor == 0 {
						takeColor = reservedList[j].color
					} else {
						continue
					}
				} else {
					if takeColor != reservedList[j].color && takeColor != 0 {
						continue
					} else {
						takeColor = reservedList[j].color + 1
					}
				}
			}
			reservedList[i].color = takeColor
		}
	}
	for _, v := range reservedList {
		fmt.Println("vertex " + fmt.Sprint(v.value) + " orginal degree " + fmt.Sprint(v.oDegree) + " degree when deleted " + fmt.Sprint(v.degree) + " color " + fmt.Sprint(v.color))
	}
	// outPutGraph(verticesList,"coloringGraphWithTheSmallestLastOrdering")
}

//Smallest Last Vertex Ordering
//TODO not stable,need to fixed
func coloringGraphWithTheSmallestLastOrderingV1(verticesList []*Vertex, adjacentMatrix [][]int, sameCurrentDegreeList []*Vertex) ([]*Vertex, [][]int, []int) {
	oLength := len(verticesList)
	orderList := make([]int, oLength)
	count := oLength - 1
	i := 0
	for count >= 0 {
		if sameCurrentDegreeList[i] != nil {
			//pick a vetex to be removed
			pickVertexIndex := sameCurrentDegreeList[i].value - 1
			//push on stack
			// reservedList[count] = verticesList[pickVertexIndex]
			orderList[count] = verticesList[pickVertexIndex].value
			//remove from dll
			if nil != sameCurrentDegreeList[i].next {
				sameCurrentDegreeList[i] = sameCurrentDegreeList[i].next
				sameCurrentDegreeList[i].last = nil
			} else {
				sameCurrentDegreeList[i] = nil
			}
			//record when degree deleted
			verticesList[pickVertexIndex].markDelete()
			//decrement degrees from its adjacent vertices
			currNode := verticesList[pickVertexIndex].head
			for currNode != nil {
				if !verticesList[currNode.value-1].deleted {
					//remove vertex from sameCurrentDegreeList
					sameCurrentDegreeListRemoveVertex(sameCurrentDegreeList, verticesList[currNode.value-1])
					verticesList[currNode.value-1].degree--
					//add vertex from sameCurrentDegreeList
					sameCurrentDegreeListAddVertex(sameCurrentDegreeList, verticesList[currNode.value-1])
				}
				currNode = currNode.next
				// fmt.Printf("currNode loop count:%d i:%d \n", count, i)
			}
			//move adjacent vertices in the dll
			count--
			i = 0
			// fmt.Printf("scdListLen:%d\n", countSameCurrentDegreeList(sameCurrentDegreeList))
		} else {
			i++
		}
	}
	// fmt.Println(orderList)
	// startColoring(verticesList, adjacentMatrix, orderList)
	return verticesList, adjacentMatrix, orderList
}
