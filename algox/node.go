package algox

// a toric node
type node struct {
	up    *node
	down  *node
	left  *node
	right *node
	head  *node
	size  int
	index int
}

// Remove the node from its vertical list.
func (n *node) vremove() {
	n.up.down = n.down
	n.down.up = n.up
}

// Restores the node to its vertical list.
func (n *node) vrestore() {
	n.down.up = n
	n.up.down = n
}

// Remove the node from its horizontal list.
func (n *node) hremove() {
	n.right.left = n.left
	n.left.right = n.right
}

// Restores the node to its horizontal list.
func (n *node) hrestore() {
	n.right.left = n
	n.left.right = n
}

// Cover a column.
func (head *node) cover() {
	head.hremove() // hide column from its own list
	for node := head.down; node != head; node = node.down {
		// hide all rows in column from their own columns
		for mode := node.right; mode != node; mode = mode.right {
			mode.vremove()
			mode.head.size--
		}
	}
}

// Uncover a column.
func (head *node) uncover() {
	for node := head.up; node != head; node = node.up {
		// hide all rows in column from their own columns
		for mode := node.left; mode != node; mode = mode.left {
			mode.vrestore()
			mode.head.size++
		}
	}
	head.hrestore() // hide column from its own list
}
