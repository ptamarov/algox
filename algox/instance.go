package algox

type instance struct {
	head     *node
	solution []*node // keep track of rows in a partial solution
}

// Return the number of elements in the underlying set.
func (p *instance) Size() int {
	return p.head.size
}

// Return the number of subsets in the collection.
func (p *instance) Parts() int {
	return p.head.index
}

// Return a list of pointers to the rows in a solution.
func (p *instance) Solution() []*node {
	return p.solution
}

// Create a new exact cover problem instance.
//
// size: the size of the underlying set to cover
// subsets: slices of strictly increasing integers in [0, size)
func New(size int, subsets [][]int) instance {
	var p instance

	// keep track of the last inserted node in each column
	lastInserted := make(map[int]*node, size)

	// keep a pointer to the column heads
	columnHeads := make(map[int]*node, size)

	// the head of the instance
	numSubsets := len(subsets)
	iHead := &node{size: size, index: numSubsets}
	p.head = iHead

	// create column list at iHead
	currentColumnHead := iHead
	for i := 0; i < size; i++ {
		newColumnNode := &node{index: i}

		// book-keeping
		lastInserted[i] = newColumnNode
		columnHeads[i] = newColumnNode

		// link to previous
		currentColumnHead.right = newColumnNode
		newColumnNode.left = currentColumnHead
		newColumnNode.head = iHead

		// continue
		currentColumnHead = newColumnNode
	}
	// close up the list
	currentColumnHead.right = iHead
	iHead.left = currentColumnHead

	// now populate using subsets
	for _, subset := range subsets {
		var last *node
		var first *node

		for _, element := range subset {
			newNode := &node{}

			// rows do not have heads so initialize first done in circular list and do nothing
			if first == nil {
				first = newNode
			} else {
				last.right = newNode
				newNode.left = last
			}

			// move on to this node
			last = newNode

			// link current node vertically to the last pointer
			lastInserted[element].down = last
			last.up = lastInserted[element]
			last.head = columnHeads[element]

			// update column size
			columnHeads[element].size++

			// update last vertically inserted
			lastInserted[element] = last
		}

		// close up this row
		first.left = last
		last.right = first
	}

	// close up each column
	for i := 0; i < size; i++ {
		lastInserted[i].down = columnHeads[i]
		columnHeads[i].up = lastInserted[i]
	}

	return p
}
