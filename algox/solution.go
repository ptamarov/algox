package algox

import "fmt"

// Selects the first column with the smallest size.
func (i *instance) selectColumn() *node {
	lightest := i.head.right
	minSize := lightest.size
	current := lightest

	for current != i.head {
		if current.size < minSize {
			minSize = current.size
			lightest = current
		}
		current = current.right
	}
	return lightest
}

func (i *instance) Solve() {
	i.solution = make([]*node, i.head.size)
	i.recursiveSolve(0)
}

func (i *instance) recursiveSolve(k int) {
	// if the list of columns is empty, a solution was found
	if i.head.right == i.head {
		fmt.Println("** found a solution at height ", k, " **")

		for idx := 0; idx < k; idx++ {
			fmt.Print("\t")
			row := i.solution[idx]
			c := row
			for {
				fmt.Print(c.head.index, " ")
				c = c.right
				if c == row {
					break
				}
			}
			fmt.Println()
		}
		return
	}

	head := i.selectColumn() // choose a column with the least number of elements in it
	if head.size == 0 {      // if column is empty, then fail
		return
	} else {
		head.cover() // else cover the column
	}

	// for each row in the column, choose it as a possible part of the solution
	for r := head.down; r != head; r = r.down {
		i.solution[k] = r

		// any other columns colliding with r must be discarded
		for node := r.right; node != r; node = node.right {
			node.head.cover()
		}

		// attempt to solve this subproblem
		i.recursiveSolve(k + 1)

		// undo the covering of columns associated to this row (reverse order)
		for node := r.left; node != r; node = node.left {
			node.head.uncover()
		}
	}
	// uncover this column and move on to another one
	head.uncover()
}
