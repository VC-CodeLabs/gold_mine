package main

// import other packages
import (
	"fmt"
	"math/rand"
	"sync"
)

var maxGoldSize = 9
var maxY = 0
var maxX = 0
var mine = make([][]int, 0)
var bestPathSum = 0
var mu sync.Mutex

func main() {
	// Alek, replace this mine with the test mine
	mine = createMine(5, 5) // x, y

	maxX = len(mine[0]) - 1
	maxY = len(mine) - 1
	printMine()
	findBestPath()
	fmt.Printf("\nSum of best path: %d", bestPathSum)
}

func createMine(x, y int) [][]int {
	new_mine := make([][]int, y)
	for i := 0; i < y; i++ {
		new_mine[i] = make([]int, x)
		for i2 := 0; i2 < x; i2++ {
			new_mine[i][i2] = rand.Intn(maxGoldSize + 1)
		}
	}

	return new_mine
}

func printMine() {
	for i := 0; i < len(mine); i++ {
		for j := 0; j < len(mine[i]); j++ {
			fmt.Printf("[%d]", mine[i][j])
		}
		fmt.Printf("\n")
	}
}

func findBestPath() {
	var wg sync.WaitGroup
	for y := 0; y < len(mine); y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			findPathFromCell(y, 0, 0)
		}(y)
	}
	wg.Wait()
}

func findPathFromCell(y, x, currSum int) {
	currSum += mine[y][x]
	if x == maxX {
		bestPathSum = max(bestPathSum, currSum)
		return
	}

	for i := -1; i <= 1; i++ {
		if y+i >= 0 && y+i <= maxY {
			findPathFromCell(y+i, x+1, currSum)
		}
	}
}
