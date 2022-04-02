package Search

import (
	"container/list"
	"log"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

// 2812m Find the Safest Path in a Grid
func maximumSafenessFactor(grid [][]int) int {
	Rows, Cols := len(grid), len(grid[0])

	Q := list.New()
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if grid[r][c] == 1 {
				grid[r][c] = 0
				Q.PushBack([]int{r, c})
			} else {
				grid[r][c] = -1
			}
		}
	}

	Dirs := []int{0, 1, 0, -1, 0}
	for Q.Len() > 0 { // multi-source BFS
		cord := Q.Remove(Q.Front()).([]int)
		distance := grid[cord[0]][cord[1]]
		for i := range Dirs[:4] {
			r, c := cord[0]+Dirs[i], cord[1]+Dirs[i+1]
			if r >= 0 && c >= 0 && Rows > r && Cols > c && grid[r][c] == -1 {
				grid[r][c] = distance + 1
				Q.PushBack([]int{r, c})
			}
		}
	}

	for r := range grid {
		log.Print(" |> ", grid[r])
	}

	ValidPath := func(factor int) bool {
		Vis := make([][]bool, Rows)
		for r := range Vis {
			Vis[r] = make([]bool, Cols)
		}

		Q := list.New()
		Q.PushBack([]int{0, 0})
		Vis[0][0] = true

		for Q.Len() > 0 {
			cord := Q.Remove(Q.Front()).([]int)
			if cord[0] == Rows-1 && cord[1] == Cols-1 {
				return true
			}
			for i := range Dirs[:4] {
				r, c := cord[0]+Dirs[i], cord[1]+Dirs[i+1]
				if r >= 0 && c >= 0 && Rows > r && Cols > c && !Vis[r][c] && grid[r][c] >= factor {
					Vis[r][c] = true
					Q.PushBack([]int{r, c})
				}
			}
		}
		return false
	}

	l, r := 0, max(Rows, Cols) // Manhattan distance
	factor := 0
	for l < r {
		m := l + (r-l)>>1
		log.Print(" -> ", m)
		if ValidPath(m) {
			factor = m
			l = m + 1
		} else {
			r = m
		}
	}
	return factor
}

// 1219m Path with Maximum Gold
func getMaximumGold(grid [][]int) int {
	Rows, Cols := len(grid), len(grid[0])

	dirs := []int{0, 1, 0, -1, 0}
	gold := 0
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if grid[r][c] == 0 {
				continue
			}

			// need BackTracking
			Q, V, G := list.New(), list.New(), list.New()
			Q.PushBack([]int{r, c})
			V.PushBack(map[[2]int]bool{})
			G.PushBack(grid[r][c])

			for Q.Len() > 0 {
				log.Print(r, c, " -> ", Q.Len())

				cord := Q.Remove(Q.Front()).([]int)
				Vis := V.Remove(V.Front()).(map[[2]int]bool)
				g := G.Remove(G.Front()).(int)

				Vis[[2]int{cord[0], cord[1]}] = true
				gold = max(g, gold)

				for i := range dirs[:4] {
					x, y := cord[0]+dirs[i], cord[1]+dirs[i+1]
					if x >= 0 && y >= 0 && Rows > x && Cols > y && !Vis[[2]int{x, y}] && grid[x][y] > 0 {
						Q.PushBack([]int{x, y})
						nVis := map[[2]int]bool{}
						for k, v := range Vis {
							nVis[k] = v
						}
						V.PushBack(nVis)
						G.PushBack(g + grid[x][y])
					}
				}
			}
		}
	}
	return gold
}
