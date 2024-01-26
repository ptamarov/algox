package algox

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
