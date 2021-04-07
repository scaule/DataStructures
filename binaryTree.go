package main

import "fmt"

type BinaryTree struct {
	value     int
	rightLeaf *BinaryTree
	leftLeaf  *BinaryTree
}

func createBinaryTree(firstValue int) *BinaryTree {
	return &BinaryTree{
		value:     firstValue,
		rightLeaf: nil,
		leftLeaf:  nil,
	}
}

func (l *BinaryTree) insert(newValue int, leaf *BinaryTree) bool {
	currentLeaf := l

	if leaf != nil {
		currentLeaf = leaf
	}
	if newValue == currentLeaf.value {
		return true
	}
	if newValue < currentLeaf.value {
		if currentLeaf.leftLeaf == nil {
			currentLeaf.leftLeaf = &BinaryTree{
				value:     newValue,
				rightLeaf: nil,
				leftLeaf:  nil,
			}
			return true
		}
		return currentLeaf.insert(newValue, currentLeaf.leftLeaf)
	} else {
		if currentLeaf.rightLeaf == nil {
			currentLeaf.rightLeaf = &BinaryTree{
				value:     newValue,
				rightLeaf: nil,
				leftLeaf:  nil,
			}
			return true
		}

		return currentLeaf.insert(newValue, currentLeaf.rightLeaf)
	}
}

func (l *BinaryTree) printBinary(leaf *BinaryTree) {
	current := l

	if leaf != nil {
		current = leaf
	}
	fmt.Println(current.value)

	if current.leftLeaf != nil {
		fmt.Println("/")
		current.printBinary(current.leftLeaf)
	}

	if current.rightLeaf != nil {
		fmt.Println("\\")
		current.printBinary(current.rightLeaf)
	}
}

func (l *BinaryTree) search(value int, leaf *BinaryTree) (*BinaryTree){
	currentLeaf := l

	if leaf != nil {
		currentLeaf = leaf
	}

	if currentLeaf.value == value {
		return currentLeaf
	}

	if value < currentLeaf.value {
		if currentLeaf.leftLeaf.value == value {
			return currentLeaf.leftLeaf
		}

		if currentLeaf.leftLeaf == nil {
			return nil
		}

		return currentLeaf.search(value, currentLeaf.leftLeaf)
	} else {
		if currentLeaf.rightLeaf.value == value {
			return currentLeaf.rightLeaf
		}
		if currentLeaf.rightLeaf == nil {
			return nil
		}

		return currentLeaf.search(value, currentLeaf.rightLeaf)
	}
}

