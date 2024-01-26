package algox

import "fmt"

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

// Turn the instance into a list of subsets.
func (p *instance) Subsets() [][]int {
	subsets := [][]int{}
	visited := make(map[*node]bool)

	for column := p.head.right; column != p.head; column = column.right {
		for node := column.down; node != column; node = node.down {
			if !visited[node] {
				subset := []int{}
				subset = append(subset, node.head.index)
				visited[node] = true
				rowNode := node.right
				for {
					if rowNode == node {
						break
					}
					subset = append(subset, rowNode.head.index)
					visited[rowNode] = true
					rowNode = rowNode.right
				}
				subsets = append(subsets, subset)
			}
		}
	}
	return subsets
}

// Create a new exact cover problem instance.
//
// size: the size of the underlying set to cover
// subsets: slices of strictly increasing integers in [0, size)
func New(size int, subsets [][]int) instance {
	if size <= 0 {
		panic("cannot initialize instance with size <= 0")
	}

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
	for index, subset := range subsets {
		if len(subset) == 0 {
			continue
		}

		var last *node
		var first *node
		lastElem := -1

		for _, element := range subset {
			if element >= size || element < 0 {
				panic(fmt.Sprintf("[at index %d] found %d but elements must be in range [0,%d)", index, element, size))
			}

			if lastElem >= element {
				panic(fmt.Sprintf("[at index %d] subset must be increasing", index))
			}
			lastElem = element

			newNode := &node{}

			// rows do not have heads so initialize first
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
