package main

import (
	"fmt"
	"slices"
)

const FIND_PATHS = true

const USE_ANSI = true

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

func dig(mine [][]int) {
	var rows = len(mine)
	var lr = rows - 1
	var cols = len(mine[0])
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

	var maxGold = -1
	var paths [][]coord = make([][]coord, 0)
	for r := 0; r < rows; r++ {
		var maxStep = mine[r][0] + max(max(nodes[r][0].up, nodes[r][0].rt), nodes[r][0].dn)
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

	if maxGold == 0 {
		fmt.Print("The mine is devoid of gold??")
	} else {
		fmt.Print("Max gold ", ANSI_GOLD_PREFIX(), maxGold, ANSI_GOLD_SUFFIX())

		if FIND_PATHS {
			fmt.Print(" in ", maxPaths, " path(s).")
		}
	}

	fmt.Println()

	if FIND_PATHS {
		for p := 0; p < maxPaths; p++ {
			fmt.Print("Path #", p, ":")
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

	mineAllOnes := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 2},
	}
	dig(mineAllOnes)

	mineSample := [][]int{
		{0, 0, 0, 10},
		{0, 0, 0, 9},
		{0, 0, 0, 8},
		{1, 1, 1, 8},
	}

	dig(mineSample)

	justOne := [][]int{
		{0, 0, 1},
		{0, 0, 0},
		{1, 0, 0},
	}

	dig(justOne)

}
