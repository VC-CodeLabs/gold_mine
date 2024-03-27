package main

import (
	"fmt"
	"log"
	"time"
)

type Mine [][]*Node

type Node struct {
	Value int
	Sum   int
}

func main() {
	log.Println("Gold mine!")
	log.Panicln("Use the unit tests to load scenarios!")
}

func goldMine(input [][]int) int {
	start := time.Now()

	mine := parseInput(input)

	result := search(mine)

	fmt.Printf("Execution took %vns\n", time.Since(start).Nanoseconds())

	return result
}

func search(mine Mine) int {
	for col := 0; col < len(mine[0]); col++ {
		for row := 0; row < len(mine); row++ {
			node := mine[row][col]

			top := findTop(mine, row, col)
			left := findLeft(mine, row, col)
			bottom := findBottom(mine, row, col)

			sum := node.Value

			if top != nil {
				topSum := top.Sum + node.Value
				if topSum > sum {
					sum = topSum
				}
			}

			if left != nil {
				leftSum := left.Sum + node.Value
				if leftSum > sum {
					sum = leftSum
				}
			}

			if bottom != nil {
				bottomSum := bottom.Sum + node.Value
				if bottomSum > sum {
					sum = bottomSum
				}
			}

			node.Sum = sum
		}
	}

	sum := 0

	for i := 0; i < len(mine); i++ {
		node := mine[i][len(mine[0])-1]
		if node.Sum > sum {
			sum = node.Sum
		}
	}

	return sum
}

func parseInput(input [][]int) [][]*Node {
	mine := make([][]*Node, len(input))
	for i, row := range input {
		mine[i] = make([]*Node, len(row))

		for j, value := range input[i] {
			mine[i][j] = &Node{
				Value: value,
			}
		}
	}

	return mine
}

func findTop(mine Mine, i int, j int) *Node {
	if i == 0 || j == 0 {
		return nil
	}

	return mine[i-1][j-1]
}

func findLeft(mine Mine, i int, j int) *Node {
	if j == 0 {
		return nil
	}

	return mine[i][j-1]
}

func findBottom(mine Mine, i int, j int) *Node {
	if i == len(mine)-1 || j == 0 {
		return nil
	}

	return mine[i+1][j-1]
}
