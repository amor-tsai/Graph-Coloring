package assignments

import (
	"fmt"
)

var Test = 0

type SubNode struct {
	value int
	index int
	next  *SubNode
	last  *SubNode
}

type SDegree struct {
	degree      int
	nodesNumber int
	next        *Vertex
	// head        *Vertex
	// tail        *Vertex
	// list        []*Vertex
}

type Vertex struct {
	value            int
	degree           int
	oDegree          int
	degreeWhenDelete int
	color            int
	deleted          bool
	next             *Vertex
	last             *Vertex
	head             *SubNode //the first adjacent node
	tail             *SubNode // the last adjacent node
}

func (node *SubNode) append(data int) *SubNode {
	if node.value == data {
		panic("something got wrong, you have a repeat node")
	}
	subNode := initSubNode(data)
	subNode.last = node
	subNode.index = node.index + 1
	node.next = subNode
	return subNode
}

func (node *Vertex) append(subNode *Vertex) *Vertex {
	if node == nil {
		node = subNode
	} else {
		node.next = subNode
		subNode.last = node
	}
	return node
}

// func (sdegree *SDegree) append(vertex *Vertex) *SDegree {
// 	sdegree.next = vertex
// }

// func (sdegree *SDegree) remove(vertex *Vertex) *SDegree {

// }

func (vertex *Vertex) addAdjacentVertex(data int) {
	if nil != vertex.head {
		vertex.tail = vertex.tail.append(data)
	} else {
		vertex.head = initSubNode(data)
		vertex.tail = vertex.head
	}
	vertex.degreeIncrement()
}

// func (vertex *Vertex) addAdjacentVertex(node *Vertex) {
// 	if vertex.next != nil {
// 		vertex.next.last = node
// 		node.next = vertex.next
// 		node.last = vertex
// 		vertex.next = node
// 	} else {
// 		vertex.next = node
// 		node.last = vertex
// 	}
// 	vertex.degreeIncrement()
// }

func (vertex *Vertex) degreeIncrement() {
	vertex.degree++
	vertex.oDegree++
}

func (vertex *Vertex) degreeDecrement() {
	vertex.degree--
}

func (vertex *Vertex) markDelete() {
	vertex.deleted = true
	vertex.degreeWhenDelete = vertex.degree
}

// return true if the vertex is marked deleted
func (vertex *Vertex) isDelete() bool {
	if vertex.deleted {
		return true
	} else {
		return false
	}
}

func sameCurrentDegreeListRemoveVertex(sameCurrentDegreeList []*Vertex, vertexRemoved *Vertex) []*Vertex {
	currNode := sameCurrentDegreeList[vertexRemoved.degree]
	for currNode != nil {
		if currNode.value == vertexRemoved.value {
			if nil == currNode.last {
				sameCurrentDegreeList[vertexRemoved.degree] = currNode.next
				if nil != sameCurrentDegreeList[vertexRemoved.degree] {
					sameCurrentDegreeList[vertexRemoved.degree].last = nil
				}
			} else {
				currNode.last.next = currNode.next
				if currNode.next != nil {
					currNode.next.last = currNode.last
				}

			}
			currNode = nil
			break
		}
		currNode = currNode.next
	}
	return sameCurrentDegreeList
}

// func sameCurrentDegreeListRemoveVertex(sameCurrentDegreeList []*Vertex, vertexRemoved *Vertex) []*Vertex {
// 	currNode := sameCurrentDegreeList[vertexRemoved.degree]
// 	for currNode != nil {
// 		if currNode.value == vertexRemoved.value {
// 			if nil == currNode.next {
// 				if vertexRemoved.last == nil {
// 					sameCurrentDegreeList[vertexRemoved.degree] = nil
// 				} else {
// 					currNode.last.next = nil
// 					currNode.last = nil
// 				}
// 			} else {
// 				if currNode.last == nil {
// 					sameCurrentDegreeList[currNode.degree] = currNode
// 					sameCurrentDegreeList[currNode.degree].last = nil
// 				} else {
// 					currNode.last.next = currNode.next
// 					currNode.next.last = currNode.last
// 				}
// 			}
// 			break
// 		}
// 		currNode = currNode.next
// 	}
// 	return sameCurrentDegreeList
// }

func sameCurrentDegreeListAddVertex(sameCurrentDegreeList []*Vertex, vertexAdd *Vertex) []*Vertex {
	// fmt.Printf("verterxAdd.degree:%d\n", vertexAdd.degree)
	vertexAdd.next = nil
	vertexAdd.last = nil
	currNode := sameCurrentDegreeList[vertexAdd.degree]
	if nil == currNode {
		sameCurrentDegreeList[vertexAdd.degree] = vertexAdd
		sameCurrentDegreeList[vertexAdd.degree].next = nil
		sameCurrentDegreeList[vertexAdd.degree].last = nil
	} else {
		vertexAdd.next = sameCurrentDegreeList[vertexAdd.degree]
		sameCurrentDegreeList[vertexAdd.degree].last = vertexAdd
		sameCurrentDegreeList[vertexAdd.degree] = vertexAdd
	}
	return sameCurrentDegreeList
}

// func sameCurrentDegreeListAddVertex(sameCurrentDegreeList []*Vertex, vertexAdd *Vertex) []*Vertex {
// 	// fmt.Printf("verterxAdd.degree:%d\n", vertexAdd.degree)
// 	currNode := sameCurrentDegreeList[vertexAdd.degree]
// 	if nil == currNode {
// 		sameCurrentDegreeList[vertexAdd.degree] = vertexAdd
// 		vertexAdd.next = nil
// 		vertexAdd.last = nil
// 	} else {
// 		vertexAdd.next = sameCurrentDegreeList[vertexAdd.degree]
// 		sameCurrentDegreeList[vertexAdd.degree].last = vertexAdd
// 		sameCurrentDegreeList[vertexAdd.degree] = vertexAdd
// 	}
// 	return sameCurrentDegreeList
// }

func initSDgreeList(vertexList []*Vertex) []*Vertex {
	length := len(vertexList)
	sdegreeList := make([]*Vertex, length)
	for i := 0; i < length; i++ {
		if vertexList[i].isDelete() {
			continue
		}
		if sdegreeList[vertexList[i].degree] == nil {
			sdegreeList[vertexList[i].degree] = vertexList[i]
		} else {
			vertexList[i].next = sdegreeList[vertexList[i].degree]
			sdegreeList[vertexList[i].degree].last = vertexList[i]
			sdegreeList[vertexList[i].degree] = vertexList[i]
			// sdegreeList[vertexList[i].degree].next = vertexList[i]
			// vertexList[i].last = sdegreeList[vertexList[i].degree]
		}
	}
	return sdegreeList
}

func initVertex(data int) *Vertex {
	vertex := &Vertex{value: data}
	vertex.degree = 0
	vertex.oDegree = 0
	vertex.color = 0
	vertex.degreeWhenDelete = 0
	vertex.last = nil
	vertex.next = nil
	vertex.head = nil
	vertex.deleted = false
	return vertex
}

func initSubNode(data int) *SubNode {
	subNode := &SubNode{value: data}
	subNode.index = 0
	subNode.last = nil
	subNode.next = nil
	return subNode
}

func initGraph(v int) ([]*Vertex, [][]int) {
	graph := make([]*Vertex, v)
	adjacentMatrix := make([][]int, v)
	for i := 0; i < v; i++ {
		graph[i] = initVertex(i + 1)
		adjacentMatrix[i] = make([]int, v)
	}
	return graph, adjacentMatrix
}

// func removeVertex(verticesList []*Vertex, removeVertexIndex int) []*Vertex {
// 	removeVertex := verticesList[removeVertexIndex]
// 	verticesListLength := len(verticesList)
// 	if verticesList[removeVertexIndex].degree > 0 {
// 		for i := 0; i < verticesListLength; i++ {
// 			if i != removeVertexIndex {
// 				currentNode := verticesList[i].head
// 				for currentNode != nil {
// 					if currentNode.value == removeVertex.value {
// 						if currentNode.last == nil && currentNode.next != nil {
// 							verticesList[i].head = currentNode.next
// 							currentNode.next.last = nil
// 						} else if currentNode.next == nil && currentNode.last != nil {
// 							currentNode.last.next = nil
// 							verticesList[i].tail = currentNode.last
// 						} else if currentNode.next == nil && currentNode.last == nil {
// 							verticesList[i].head = nil
// 							verticesList[i].tail = nil
// 						} else {
// 							currentNode.last.next = currentNode.next
// 							currentNode.next.last = currentNode.last
// 						}
// 						verticesList[i].degreeDecrement()
// 						break
// 					}
// 					currentNode = currentNode.next
// 				}
// 			}
// 		}
// 	}
// 	verticesList[removeVertexIndex] = verticesList[verticesListLength-1]
// 	verticesList = verticesList[:verticesListLength-1]
// 	return verticesList
// }

func outPutGraph(verticesList []*Vertex, fileName string) {
	verticesNum := len(verticesList)
	var outPutVerticesNumber string
	var outPutVertex []string
	var outPutAdjacencyList []string
	location := 0
	indexOfOutputVertex := 1
	indexOfOutputAdjacencyList := verticesNum + 1
	for key, vertex := range verticesList {
		if 0 == key {
			location += verticesNum + 1
		} else {
			location += verticesList[key-1].degree
		}
		if verticesList[key].degree != 0 {
			outPutVertex = append(outPutVertex, fmt.Sprint(location)+" #"+fmt.Sprint(indexOfOutputVertex)+
				"th value=starting location for vertex "+fmt.Sprint(vertex.value)+"'s edges \n")
		} else {
			outPutVertex = append(outPutVertex, "-1 #"+fmt.Sprint(indexOfOutputVertex)+
				"th value=starting location for vertex "+fmt.Sprint(vertex.value)+"'s edges \n")
		}
		currentNode := vertex.head
		for currentNode != nil {
			outPutAdjacencyList = append(outPutAdjacencyList, fmt.Sprint(currentNode.value)+" #"+fmt.Sprint(indexOfOutputAdjacencyList)+"th value=Vertex "+
				fmt.Sprint(vertex.value)+" is adjacent to Vertex "+fmt.Sprint(currentNode.value)+" \n")
			currentNode = currentNode.next
			indexOfOutputAdjacencyList++
		}
		indexOfOutputVertex++
	}
	outPutVerticesNumber = fmt.Sprint(verticesNum) + " #0th value=Number of vertices \n"

	// fmt.Print(outPutVerticesNumber)
	writeFile(fileName, outPutVerticesNumber)
	for _, value := range outPutVertex {
		// fmt.Print(value)
		if Test == 1 {
			fmt.Print(value)
		} else {
			writeFile(fileName, value)
		}
	}
	for _, value := range outPutAdjacencyList {
		// fmt.Print(value)
		if Test == 1 {
			fmt.Print(value)
		} else {
			writeFile(fileName, value)
		}
	}
	if Test == 1 {
		fmt.Println(outPutVertex)
		fmt.Println(outPutAdjacencyList)
	} else {
		writeFile(fileName, "\n")
	}

	// fmt.Println(outPutVertex)
	// fmt.Println(outPutAdjacencyList)
}

func verifyRandomGraph(verticesList []*Vertex, fileName string) {
	l := len(verticesList)
	if l > 0 {
		edgeNumber := 0
		var ajacentMatrix [5000][5000]int
		for _, vertex := range verticesList {
			currentNode := vertex.head
			for currentNode != nil {
				ajacentMatrix[vertex.value-1][currentNode.value-1] = 1
				currentNode = currentNode.next
			}
			edgeNumber += vertex.degree
			fmt.Println("vertex " + fmt.Sprint(vertex.value) + " has " + fmt.Sprint(vertex.degree) + " edges")
			writeFile(fileName+"_vertexId_degree", fmt.Sprintf("%d \t %d\n", vertex.value, vertex.degree))
		}
		for key1, arr1 := range ajacentMatrix {
			for key2 := range arr1 {
				if ajacentMatrix[key2][key1] != ajacentMatrix[key1][key2] {
					panic("fail to verify at key " + fmt.Sprint(key2) + " and key " + fmt.Sprint(key1))
				}
			}
		}
		fmt.Println("edge number " + fmt.Sprint(edgeNumber/2))
		fmt.Println("vertices number " + fmt.Sprint(len(verticesList)))
	}
}
