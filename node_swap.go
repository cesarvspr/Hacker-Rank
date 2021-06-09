package main

import (
	"fmt"
)

/*
 * Complete the 'swapNodes' function below.
 * This is not made by cesarvspr.
 * Cesarvspr wrote the solution again and made some error controls
 * with comments for educational purposes.
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY indexes
 *  2. INTEGER_ARRAY queries
 */

// index, level, k: all one based
func swapNodes(tree []int, index, level, k int) {
	i := (index - 1) * 2
	//THIS CHECKS IS WE ARE IN THE CORRECT LEVEL SO WE CAN SWAP
	//The requirement is: if h (level) is a multiple of k,
	//swap the left and right subtrees of that level
	if level%k == 0 {
		tree[i], tree[i+1] = tree[i+1], tree[i]
		// fmt.Println("SWAPED: ")

	}

	level++
	if tree[i] != -1 {
		swapNodes(tree, tree[i], level, k)
	}

	if tree[i+1] != -1 {
		swapNodes(tree, tree[i+1], level, k)
	}
	// fmt.Println("SWAP END ")
}

func main() {
	var err error
	var t int // swap operations
	var left, right int
	var k int
	var n int
	fmt.Scan(&n)
	// fmt.Println("n is: ", n)
	var tree = []int{} //binary tree
	// fmt.Println("tree is: ", tree)
	for i := 0; i != n; i++ {
		_, err = fmt.Scan(&left, &right)
		if err != nil {
			fmt.Errorf("Error: ", err)
		}
		tree = append(tree, left, right)
	}
	// fmt.Println("new tree is: ", tree)
	// fmt.Println("")
	_, err = fmt.Scan(&t)
	if err != nil {
		fmt.Errorf("Error: ", err)
	}
	// fmt.Println("t is: ", t)

	for i := 0; i != t; i++ {
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Errorf("Error: ", err)
		}
		// fmt.Println("k is: ", k)
		swapNodes(tree, 1, 1, k)
		var results []int
		printTree(tree, 1, &results)
		str := fmt.Sprintf("%v", results)
		fmt.Println(str[1 : len(str)-1])

	}
	_, err = fmt.Scan(&t)
	if err != nil {
		fmt.Errorf("Error: ", err)
	}
}

// index: 1 based
func printTree(tree []int, index int, results *[]int) {
	i := (index - 1) * 2
	if tree[i] != -1 {
		printTree(tree, tree[i], results)
	}

	*results = append(*results, index)
	// fmt.Println("results changed, result now is:", *results)
	if tree[i+1] != -1 {
		printTree(tree, tree[i+1], results)
	}
}
