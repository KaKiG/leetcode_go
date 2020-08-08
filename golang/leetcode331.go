package main

import "strings"

func main() {

}

// using recursion, to validate a tree, we can validate its left child and its right child

func isValidSerialization(preorder string) bool {
	str := strings.Split(preorder, ",")
	i := 0
	return isValid(str, &i) && (i == len(str))
}

func isValid(str []string, i *int) bool {
	// if index moves to the last node and still no return
	// it means the last node should have some children, but there is nothing ,so it is invalid
	// return false
	if *i == len(str) {
		return false
	}

	// if meeting "#", it means the end of a substree, by adding one to the index, we move to the another subtree
	if str[*i] == "#" {
		*i += 1
		return true
	}
	// here adding 1 means meeting moves to its left substree
	*i += 1
	// here left side means result from left subtree, and if left subtree goes to the end
	// it will automatically move to the right subtree, since when meeting "#", we move one step further
	return isValid(str, i) && isValid(str, i)
}
