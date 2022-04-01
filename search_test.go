package Search

import (
	"container/heap"
	"container/list"
	"log"
	"testing"
)

func init() {
	log.Print("> Search: DFS/BFS")
}

type e2812 struct {
	priority int
	cord     []int
}
type pq2812 []e2812

func (p pq2812) Len() int           { return len(p) }
func (p pq2812) Less(i, j int) bool { return p[i].priority > p[j].priority }
func (p pq2812) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pq2812) Push(x any)        { *p = append(*p, x.(e2812)) }
func (p *pq2812) Pop() any {
	v := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return v
}

// 2812m Find the Safest Path in a Grid
func Test2812(t *testing.T) {
	Dijkstra := func(grid [][]int) int {
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

		// Multi-source BFS
		Dir := []int{0, 1, 0, -1, 0}
		for Q.Len() > 0 {
			cord := Q.Remove(Q.Front()).([]int)
			for i := range Dir[:4] {
				r, c := cord[0]+Dir[i], cord[1]+Dir[i+1]
				if r >= 0 && c >= 0 && Rows > r && Cols > c && grid[r][c] == -1 {
					grid[r][c] = 1 + grid[cord[0]][cord[1]]
					Q.PushBack([]int{r, c})
				}
			}
		}

		type e = e2812
		pq := pq2812{}

		heap.Push(&pq, e{grid[0][0], []int{0, 0}})
		grid[0][0] = -1

		dirs := []int{0, 1, 0, -1, 0}
		for pq.Len() > 0 {
			log.Print("PQ -> ", pq)
			entry := heap.Pop(&pq).(e)

			factor, cord := entry.priority, entry.cord
			if cord[0] == Rows-1 && cord[1] == Cols-1 {
				return factor
			}

			for i := range dirs[:4] {
				r, c := cord[0]+dirs[i], cord[1]+dirs[i+1]
				if r >= 0 && c >= 0 && Rows > r && Cols > c && grid[r][c] != -1 {
					heap.Push(&pq, e{min(factor, grid[r][c]), []int{r, c}})
					grid[r][c] = -1
				}
			}
		}

		return 0
	}

	for _, f := range []func([][]int) int{maximumSafenessFactor, Dijkstra} {
		log.Print("2 ?= ", f([][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}))
		log.Print("0 ?= ", f([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
		log.Print("2 ?= ", f([][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}))
	}
}
