package main

import (
	"fmt"
	"reflect"
	"slices"
	"time"
)

// ALEK- here's the default grid to update w/ your test(s)
var GOLD_GRID = [][]int{
	{1, 3, 3},
	{2, 1, 4},
	{0, 6, 4},
}

// if enabled, RUN_TESTS skips the default GOLD_GRID above
// and instead runs various tests w/ validation
const RUN_TESTS = false

// OPTIONAL- find the path(s) to max gold
const FIND_PATHS = false

// OPTIONAL- pretty-print the paths to gold
const USE_ANSI = false

func dig(mine [][]int) GoldMap {
	s := time.Now()
	var rows = len(mine)
	if rows == 0 {
		fmt.Println("Mine is void (nil or {})")
		return GoldMap{0, nil}
	}
	var lr = rows - 1
	var cols = len(mine[0])
	if cols == 0 {
		fmt.Println("Mine has no content ({{}})")
		return GoldMap{0, nil}
	}
	var lc = cols - 1

	var nodes [][]node = make([][]node, rows)

	for r := 0; r < rows; r++ {
		nodes[r] = make([]node, cols)
		for c := 0; c < cols; c++ {

			// nodes[r][c] = new(node)

			if c < lc {
				if r > 0 {
					nodes[r][c].up = mine[r-1][c+1]
				} else {
					nodes[r][c].up = -1 // can't step up at top edge
				}

				nodes[r][c].rt = mine[r][c+1] // step right is always safe

				if r < lr {
					nodes[r][c].dn = mine[r+1][c+1]
				} else {
					nodes[r][c].dn = -1 // can't step down at bottom edge
				}
			} else {
				nodes[r][c].up = 0
				nodes[r][c].rt = 0
				nodes[r][c].dn = 0
			}
		}
	}

	// dumpNodes(mine,nodes);

	for c := cols; c > 0; {
		c--
		for r := rows; r > 0; {
			r--
			if c < lc {
				if r > 0 {
					nodes[r][c].up +=
						0 + // ( c == 0 ? mine[r][c] : 0 )
							max(max(nodes[r-1][c+1].up, nodes[r-1][c+1].rt), nodes[r-1][c+1].dn)
				} else {
					nodes[r][c].up += 0
				}

				nodes[r][c].rt +=
					0 + // ( c == 0 ? mine[r][c] : 0 )
						max(max(nodes[r][c+1].up, nodes[r][c+1].rt), nodes[r][c+1].dn)

				if r < lr {
					nodes[r][c].dn +=
						0 + // ( c == 0 ? mine[r][c] : 0 )
							max(max(nodes[r+1][c+1].up, nodes[r+1][c+1].rt), nodes[r+1][c+1].dn)
				} else {
					nodes[r][c].dn += 0
				}

			} else {
				nodes[r][c].up = 0
				nodes[r][c].rt = 0
				nodes[r][c].dn = 0
			}

			/*
			   System.out.println( "r=" + r + " c=" + c + ":");

			   dumpNodes(mine,nodes);
			*/
		}
	}

	/*
	   for( int r = 0; r < rows; r++ ) {
	       for( int c = 0; c < cols; c++ ) {
	           if( r > 0 )
	               nodes[r][c].up += mine[r][c];

	           nodes[r][c].rt += mine[r][c];

	           if( r < lr)
	               nodes[r][c].dn += mine[r][c];
	       }
	   }
	*/

	// dumpNodes(mine, nodes)

	var maxGold = 0
	var paths [][]coord // = make([][]coord, 0)
	for r := 0; r < rows; r++ {
		var maxStep = mine[r][0] + max(max(nodes[r][0].up, nodes[r][0].rt), nodes[r][0].dn)
		if maxStep > 0 {
			if maxStep > maxGold {
				maxGold = maxStep
				if FIND_PATHS {
					paths = make([][]coord, 1) // paths.clear()
					paths[0] = make([]coord, 1)
					paths[0][0].r = r
					paths[0][0].c = 0
					// paths.get(0).add(new coord(r,0));
				}
			} else if maxStep == maxGold {
				if FIND_PATHS {
					var newPath = make([]coord, 1)
					newPath[0].r = r
					newPath[0].c = 0
					paths = append(paths, newPath)
					// paths.add(new ArrayList<>());
					// paths.getLast().add(new coord(r,0));
				}
			}
		}
	}

	/*
		for p := 0; p < len(paths); p++ {
			fmt.Println("Path #", p+1, " [ ", paths[p][0].r, ", ", paths[p][0].c, " ]")
		}
	*/

	// dumpNodes(mine,nodes);

	var maxPaths = 0
	if FIND_PATHS {
		if maxGold > 0 {
			maxPaths = 0
			for p := 0; p < len(paths); p++ {
				var path = paths[p]
				var sum = 0
				for c := 0; c < cols; c++ {
					var coords = path[c]
					if c < len(path) {
						sum += mine[coords.r][coords.c]
					}
					if c+1 < cols && c+1 == len(path) {
						var nOde = nodes[coords.r][coords.c]
						var maxStep = max(max(nOde.up, nOde.rt), nOde.dn)
						var up = maxStep == nOde.up
						var rt = maxStep == nOde.rt
						var dn = maxStep == nOde.dn
						var subPaths = 0
						if up {
							subPaths++
						}
						if rt {
							subPaths++
						}
						if dn {
							subPaths++
						}
						if subPaths > 1 {
							if up {
								if rt {
									var newPath = make([]coord, len(path))
									copy(newPath, path)
									newPath = append(newPath, coord{coords.r, coords.c + 1})
									paths = append(paths, newPath)
								}

								if dn {
									var newPath = make([]coord, len(path))
									copy(newPath, path)
									newPath = append(newPath, coord{coords.r + 1, coords.c + 1})
									paths = append(paths, newPath)
								}

								path = append(path, coord{coords.r - 1, coords.c + 1})
							} else if rt {
								// dn:
								var newPath = make([]coord, len(path))
								copy(newPath, path)
								newPath = append(newPath, coord{coords.r + 1, coords.c + 1})
								paths = append(paths, newPath)

								// rt
								path = append(path, coord{coords.r, coords.c + 1})

							}
						} else {
							if up {
								path = append(path, coord{coords.r - 1, coords.c + 1})
							} else if rt {
								path = append(path, coord{coords.r, coords.c + 1})
							} else { // dn
								path = append(path, coord{coords.r + 1, coords.c + 1})
							}
						}
					}
				}
				paths[p] = path
				// assert(sum == maxGold)
			}
		}
	}

	maxPaths = len(paths)

	if FIND_PATHS {
		for p := 0; p < maxPaths; p++ {
			fmt.Print("Path #", p+1, ":")
			var path = paths[p]
			for s := 0; s < len(path); s++ {
				var rc = path[s]
				fmt.Print(" [", rc.r, ", ", rc.c, "]")
			}
			fmt.Println()
			for r := 0; r < rows; r++ {
				for c := 0; c < cols; c++ {
					var value = mine[r][c]
					if slices.IndexFunc(path, func(tc coord) bool { return tc.r == r && tc.c == c }) != -1 {
						fmt.Printf("%s%5d%s, ", ANSI_GOLD_PREFIX(), value, ANSI_GOLD_SUFFIX())
					} else {
						fmt.Printf("%5d, ", value)
					}
				}
				fmt.Println()
			}
			fmt.Println()

		}
	}

	d := time.Since(s)

	if maxGold == 0 {
		fmt.Print("The mine is devoid of gold??")
	} else {
		fmt.Print("Max gold ", ANSI_GOLD_PREFIX(), maxGold, ANSI_GOLD_SUFFIX())

		if FIND_PATHS {
			fmt.Print(" in ", maxPaths, " path(s).")
		}
	}

	fmt.Printf(" completed in %s", d)

	fmt.Println()

	// report our findings
	return GoldMap{maxGold, paths}

}

type node struct {
	up int // up-and-right
	rt int // straight-right
	dn int // down-and-right
}

type coord struct {
	r int // row
	c int // column
}

// the GoldMap is used to report the results of our search
type GoldMap struct {
	maxGold int
	paths   [][]coord
}

func dumpNodes(mine [][]int, nodes [][]node) {
	fmt.Println()
	var rows = len(mine)
	var cols = len(mine[0])
	var lc = cols - 1
	for r := 0; r < rows; r++ {
		fmt.Print("{ ")
		for c := 0; c < cols; c++ {
			// var dir = "|"
			var delim = ", "
			var max = max(max(nodes[r][c].up, nodes[r][c].rt), nodes[r][c].dn)
			var up = max == nodes[r][c].up
			var rt = max == nodes[r][c].rt
			var dn = max == nodes[r][c].dn
			if c < lc {

			} else {
				delim = ""
				up = false
				rt = false
				dn = false
			}
			fmt.Printf("%3d [%3d %3d %3d: %5d ", mine[r][c], nodes[r][c].up, nodes[r][c].rt, nodes[r][c].dn, max)
			if up {
				fmt.Printf("/")
			} else {
				fmt.Print(" ")
			}
			if rt {
				fmt.Printf("-")
			} else {
				fmt.Print(" ")
			}
			if dn {
				fmt.Printf("\\")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("]")
			fmt.Printf("%s", delim)
		}
		fmt.Println(" }")
	}
	fmt.Println()

}

func main() {

	if RUN_TESTS {
		runTests()
	} else {
		fmt.Println("Using built-in GOLD_GRID")
		dig(GOLD_GRID)
	}
}

func ANSI_GOLD_PREFIX() string {
	if USE_ANSI {
		return "\u001b[1m\u001b[103m\u001b[91m"
	} else {
		return ""
	}

}
func ANSI_GOLD_SUFFIX() string {

	if USE_ANSI {
		return "\u001b[0m"
	} else {
		return ""
	}
}

func runTests() {
	fmt.Println("Running tests...")
	//// my examples
	const TEST_EXTRA_EXAMPLES = true
	if TEST_EXTRA_EXAMPLES {
		testAllOnes()
		testMultiPath()
		testJustOneNugget()
	}

	//// examples from README
	const TEST_EXAMPLES = true
	if TEST_EXAMPLES {
		testExample1()
		testExample2()
		testExample3()
	}

	//// edge cases
	const TEST_EDGE_CASES = true
	if TEST_EDGE_CASES {
		testNilMine()
		testEmptyMine()
		testEmptyMine2()
		testOneDMineNoNuggets()
		testOneDMineOneNugget()
		testFlatShallowMine()
		testTallSkinnyMine()
	}

	//// manufacture max (by dims & total gold) mine
	const TEST_MAX_MINE = true
	if TEST_MAX_MINE {
		testMaxMine()
	}
}

func testAllOnes() {
	mineAllOnes := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	dig(mineAllOnes)
}

func testMultiPath() {

	mineSample := [][]int{
		{0, 0, 0, 10},
		{0, 0, 0, 9},
		{0, 0, 0, 8},
		{1, 1, 1, 8},
	}

	dig(mineSample)
}

func testJustOneNugget() {
	justOne := [][]int{
		{0, 0, 1},
		{0, 0, 0},
		{0, 0, 0},
	}

	dig(justOne)
}

func testExample1() {
	var example1 = [][]int{
		{1, 3, 3},
		{2, 1, 4},
		{0, 6, 4},
	}

	dig(example1)

}

func testExample2() {
	fmt.Println()
	fmt.Println("Testing example 2...")

	var example2 = [][]int{
		{1, 3, 1, 5},
		{2, 2, 4, 1},
		{5, 0, 2, 3},
		{0, 6, 1, 2}}

	var goldMap = dig(example2)

	assert(goldMap.maxGold == 16, fmt.Sprintf("maxGold: Expected %d Actual %d", 16, goldMap.maxGold))

	if FIND_PATHS {
		assert(len(goldMap.paths) == 2, fmt.Sprintf("#paths: Expected %d Actual %d", 2, len(goldMap.paths)))

		var EXPECTED_PATHS = [][]coord{
			{{2, 0}, {1, 1}, {1, 2}, {0, 3}},
			{{2, 0}, {3, 1}, {2, 2}, {2, 3}},
		}
		assert(reflect.DeepEqual(goldMap.paths, EXPECTED_PATHS), fmt.Sprintf("paths mismatch"))
	}

	fmt.Println("...example 2 Passed.")
}

func testExample3() {
	fmt.Println()
	fmt.Println("Testing example 3...")

	var example3 = [][]int{
		{10, 33, 13, 15},
		{22, 21, 04, 1},
		{5, 0, 2, 3},
		{0, 6, 14, 2}}

	var goldMap = dig(example3)

	assert(goldMap.maxGold == 83, fmt.Sprintf("maxGold: Expected %d Actual %d", 83, goldMap.maxGold))

	if FIND_PATHS {
		assert(len(goldMap.paths) == 1, fmt.Sprintf("#paths: Expected %d Actual %d", 1, len(goldMap.paths)))

		var EXPECTED_PATHS = [][]coord{
			{{1, 0}, {0, 1}, {0, 2}, {0, 3}},
		}
		assert(reflect.DeepEqual(goldMap.paths, EXPECTED_PATHS), fmt.Sprintf("paths mismatch"))
	}

	fmt.Println("...example 3 Passed.")
}

func testNilMine() {

	fmt.Println()
	fmt.Println("Testing nil mine...")

	var nilMine [][]int

	var goldMap = dig(nilMine)
	assert(goldMap.maxGold == 0, fmt.Sprintf("maxGold: Expected %d Actual %d", 0, goldMap.maxGold))
	assert(goldMap.paths == nil, fmt.Sprintf("paths != nil"))
}

func testEmptyMine() {
	fmt.Println()
	fmt.Println("Testing empty mine {}...")

	var emptyMine = [][]int{}

	var goldMap = dig(emptyMine)
	assert(goldMap.maxGold == 0, fmt.Sprintf("maxGold: Expected %d Actual %d", 0, goldMap.maxGold))
	assert(goldMap.paths == nil, fmt.Sprintf("paths != nil"))

}

func testEmptyMine2() {
	fmt.Println()
	fmt.Println("Testing empty mine {{}}...")

	var emptyMine = [][]int{{}}

	var goldMap = dig(emptyMine)
	assert(goldMap.maxGold == 0, fmt.Sprintf("maxGold: Expected %d Actual %d", 0, goldMap.maxGold))
	assert(goldMap.paths == nil, fmt.Sprintf("paths != nil"))
}

func testOneDMineNoNuggets() {
	fmt.Println()
	fmt.Println("Testing one-d mine no nuggets {{0}}...")

	var oneDMineNoNuggets = [][]int{{0}}

	var goldMap = dig(oneDMineNoNuggets)
	assert(goldMap.maxGold == 0, fmt.Sprintf("maxGold: Expected %d Actual %d", 0, goldMap.maxGold))

}

func testOneDMineOneNugget() {
	fmt.Println()
	fmt.Println("Testing one-d mine one nugget {{1}}...")

	var oneDMineNoNuggets = [][]int{{1}}

	var goldMap = dig(oneDMineNoNuggets)
	assert(goldMap.maxGold == 1, fmt.Sprintf("maxGold: Expected %d Actual %d", 1, goldMap.maxGold))

}

func testFlatShallowMine() {
	fmt.Println()
	fmt.Println("Testing shallow mine 1x3...")

	var oneDMineNoNuggets = [][]int{{1, 2, 3}}

	var goldMap = dig(oneDMineNoNuggets)

	assert(goldMap.maxGold == 6, fmt.Sprintf("maxGold: Expected %d Actual %d", 6, goldMap.maxGold))

}

func testTallSkinnyMine() {
	fmt.Println()
	fmt.Println("Testing skinny mine 3x1...")

	var oneDMineNoNuggets = [][]int{{1}, {2}, {3}}

	var goldMap = dig(oneDMineNoNuggets)

	assert(goldMap.maxGold == 3, fmt.Sprintf("maxGold: Expected %d Actual %d", 3, goldMap.maxGold))

}

func testMaxMine() {
	fmt.Println()
	fmt.Println("Testing maximum mine (by size)...")

	const MAX_ROWS = 1000
	const MAX_COLS = 1000
	const MAX_GOLD = 9872

	var goldLeft = MAX_GOLD

	fmt.Print("Building max mine...")
	var maxMine = make([][]int, MAX_ROWS)
	for r := 0; r < MAX_ROWS; r++ {
		maxMine[r] = make([]int, MAX_COLS)
		for c := 0; c < MAX_COLS; c++ {
			if goldLeft > 0 && (r == c || /* rectangular: */ (r == MAX_ROWS-1 && c > r)) {
				if (r == MAX_ROWS-1 || /* rectangular: */ r == c) && c == MAX_COLS-1 {
					maxMine[r][c] = goldLeft
					goldLeft = 0
				} else {
					maxMine[r][c] = 1
					goldLeft--
				}
			} else {
				maxMine[r][c] = 0
			}

		}
	}
	fmt.Println("completed.")

	var goldMap = dig(maxMine)

	assert(goldMap.maxGold == MAX_GOLD-goldLeft, fmt.Sprintf("maxGold: Expected %d Actual %d", MAX_GOLD-goldLeft, goldMap.maxGold))

	if FIND_PATHS {
		assert(len(goldMap.paths) == 1, fmt.Sprintf("#paths: Expected %d Actual %d", 1, len(goldMap.paths)))
	}

}

// ///////////////////////////////////////////
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
