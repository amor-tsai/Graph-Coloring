package assignments

var Uniform int = 1
var Skewed int = 2
var Amor int = 3

/**
* V = Number of vertices. (MAX = 10,000)
* E = Number of conflicts between pairs of vertices for random graphs. (MAX = 2,000,000)
* G = COMPLETE-1 | CYCLE-2 | RANDOM-3 (with DIST below)
* DIST = UNIFORM | SKEWED | YOURS
*
 */

/**
create complete graph
*/
func createCompleteGraph(v int) ([]*Vertex, [][]int, []*Vertex) {
	if v < 2 {
		panic("v should greater than 1")
	}
	//adjacentMatrix is used to easily find out whether two elements are adjacent
	graph, adjacentMatrix := initGraph(v)
	for i := 1; i <= v; i++ {
		for j := 1; j <= v; j++ {
			if i != j {
				graph[i-1].addAdjacentVertex(j)
				adjacentMatrix[i-1][j-1] = 1
			}
		}
	}
	sameCurrentDegreeList := initSDgreeList(graph)
	return graph, adjacentMatrix, sameCurrentDegreeList
}

/**
create cycle graph
*/
func createCycleGraph(v int) ([]*Vertex, [][]int, []*Vertex) {
	if v < 2 {
		panic("v should greater than 1")
	}
	graph, adjacentMatrix := initGraph(v)
	for i := 1; i <= v; i++ {
		left := i%v + 1
		right := i - 1
		if right == 0 {
			right = v
		}
		graph[i-1].addAdjacentVertex(left)
		adjacentMatrix[i-1][left-1] = 1
		if right != left {
			graph[i-1].addAdjacentVertex(right)
			adjacentMatrix[i-1][right-1] = 1
		}
	}
	sameCurrentDegreeList := initSDgreeList(graph)
	return graph, adjacentMatrix, sameCurrentDegreeList
}

/**
create random graph with uniform distribution
*/
func createRandomGraphWithUniformDistribution(v int, e int) ([]*Vertex, [][]int, []*Vertex) {
	if v < 2 {
		panic("v should greater than 1")
	}
	maxLength := v * (v - 1) / 2
	if e > maxLength {
		panic("e can't be greater than maximum edge")
	}
	graph, adjacentMatrix := initGraph(v)
	points := make([]*Point, maxLength)
	count := 0
	for i := 0; i < v; i++ {
		for j := i + 1; j < v; j++ {
			// fmt.Printf("count is %d \n", count)
			points[count] = &Point{x: i, y: j}
			count++
		}
	}
	sequence := randomNumberSequenceGenerator(e, maxLength, Uniform)
	for i := 0; i < len(sequence); i++ {
		graph[points[sequence[i]-1].x].addAdjacentVertex(points[sequence[i]-1].y + 1)
		graph[points[sequence[i]-1].y].addAdjacentVertex(points[sequence[i]-1].x + 1)
		adjacentMatrix[points[sequence[i]-1].x][points[sequence[i]-1].y] = 1
		adjacentMatrix[points[sequence[i]-1].y][points[sequence[i]-1].x] = 1
	}
	sameCurrentDegreeList := initSDgreeList(graph)
	return graph, adjacentMatrix, sameCurrentDegreeList
}

/**
create random graph with skewed distribution
*/
func createRandomGraphWithSkewedDistribution(v int, e int) ([]*Vertex, [][]int, []*Vertex) {
	if v < 2 {
		panic("v should greater than 1")
	}
	maxLength := v * (v - 1) / 2
	if e > maxLength {
		panic("e can't be greater than maximum edge")
	}
	graph, adjacentMatrix := initGraph(v)
	points := make([]*Point, maxLength)
	count := 0
	for i := 0; i < v; i++ {
		for j := i + 1; j < v; j++ {
			points[count] = &Point{x: i, y: j}
			count++
		}
	}
	if e*3 > maxLength {
		sequence := randomNumberSequenceGenerator(e, maxLength, Skewed)
		for i := 0; i < len(sequence); i++ {
			graph[points[sequence[i]-1].x].addAdjacentVertex(points[sequence[i]-1].y + 1)
			graph[points[sequence[i]-1].y].addAdjacentVertex(points[sequence[i]-1].x + 1)
			adjacentMatrix[points[sequence[i]-1].x][points[sequence[i]-1].y] = 1
			adjacentMatrix[points[sequence[i]-1].y][points[sequence[i]-1].x] = 1
		}
	} else {
		icount := 0
		for i := 0; i < e; i++ {
			index := skewedRandomNumber(maxLength) - 1
			if adjacentMatrix[points[index].x][points[index].y] != 1 {
				graph[points[index].x].addAdjacentVertex(points[index].y + 1)
				graph[points[index].y].addAdjacentVertex(points[index].x + 1)
				adjacentMatrix[points[index].x][points[index].y] = 1
				adjacentMatrix[points[index].y][points[index].x] = 1
			} else {
				i--
			}
			icount++
		}
	}
	sameCurrentDegreeList := initSDgreeList(graph)
	return graph, adjacentMatrix, sameCurrentDegreeList
}

/**
create random graph with tiered distribution
*/
func createRandomGraphWithAmorDistribution(v int, e int) ([]*Vertex, [][]int, []*Vertex) {
	if v < 2 {
		panic("v should greater than 1")
	}
	maxLength := v * (v - 1) / 2
	if e > maxLength {
		panic("e can't be greater than maximum edge")
	}
	graph, adjacentMatrix := initGraph(v)
	points := make([]*Point, maxLength)
	count := 0
	for i := 0; i < v; i++ {
		for j := i + 1; j < v; j++ {
			points[count] = &Point{x: i, y: j}
			count++

		}
	}
	if e*3 > maxLength {
		sequence := randomNumberSequenceGenerator(e, maxLength, Amor)
		for i := 0; i < len(sequence); i++ {
			graph[points[sequence[i]-1].x].addAdjacentVertex(points[sequence[i]-1].y + 1)
			graph[points[sequence[i]-1].y].addAdjacentVertex(points[sequence[i]-1].x + 1)
			adjacentMatrix[points[sequence[i]-1].x][points[sequence[i]-1].y] = 1
			adjacentMatrix[points[sequence[i]-1].y][points[sequence[i]-1].x] = 1
		}
	} else {
		icount := 0
		for i := 0; i < e; i++ {
			index := amorRandomNumber(maxLength) - 1
			if adjacentMatrix[points[index].x][points[index].y] != 1 {
				graph[points[index].x].addAdjacentVertex(points[index].y + 1)
				graph[points[index].y].addAdjacentVertex(points[index].x + 1)
				adjacentMatrix[points[index].x][points[index].y] = 1
				adjacentMatrix[points[index].y][points[index].x] = 1
			} else {
				i--
			}
			icount++
		}
		// fmt.Printf("\nin fact, it runs %d times\n", icount)
	}
	sameCurrentDegreeList := initSDgreeList(graph)
	return graph, adjacentMatrix, sameCurrentDegreeList
}

//For points, it needs to include both Point{x,y} and Point{y,x}
func createCustomGraph(v int, points []*Point) ([]*Vertex, [][]int, []*Vertex) {
	if v < 2 {
		panic("v should greater than 1")
	}
	graph, adjacentMatrix := initGraph(v)
	for i := 0; i < len(points); i++ {
		graph[points[i].x-1].addAdjacentVertex(points[i].y)
		// graph[points[i].y-1].addAdjacentVertex(points[i].x)
		adjacentMatrix[points[i].x-1][points[i].y-1] = 1
		// adjacentMatrix[points[i].y-1][points[i].x-1] = 1
	}
	// verifyRandomGraph(graph)
	sameCurrentDegreeList := initSDgreeList(graph)
	// outPutGraph(graph, "createRandomGraphWithUniformDistribution")
	return graph, adjacentMatrix, sameCurrentDegreeList
}
