package main

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/dannoshepard/utils"
)

type Mine [][]*Node

type Node struct {
	Value int
	Sum   int
}

func main() {
	log.Print("Gold mine!")
	raw := utils.ReadStdin()

	data, err := clean(raw)
	if err != nil {
		log.Fatal(err)
	}

	var rawMine [][]int
	if err := json.Unmarshal([]byte(data), &rawMine); err != nil {
		log.Fatal(err)
	}

	output := goldMine(rawMine)
	log.Printf("Most gold: %v\n", output)
}

func goldMine(input [][]int) int {
	mine := parseInput(input)

	return search(mine)
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

func clean(inputString string) (string, error) {
	// Remove spaces and newlines
	inputString = strings.ReplaceAll(inputString, " ", "")
	inputString = strings.ReplaceAll(inputString, "\n", "")

	pattern := `^[\[\]{},:\s\d.+-]*$`
	match, err := regexp.MatchString(pattern, inputString)
	if err != nil {
		return "", errors.New("unable to validate input string")
	}

	if match {
		return inputString, nil
	}

	return "", errors.New("invalid json string provided")
}
