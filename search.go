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
