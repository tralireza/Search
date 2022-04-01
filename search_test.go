package Search

import "log"

func init() {
	log.Print("> Search: DFS/BFS")
}

// 2812m Find the Safest Path in a Grid
func Test2812(t *testing.T) {
	log.Print("2 ?= ", maximumSafenessFactor([][]int{{0, 0, 1}, {0, 0, 0}, {0, 0, 0}}))
	log.Print("0 ?= ", maximumSafenessFactor([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
	log.Print("2 ?= ", maximumSafenessFactor([][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {1, 0, 0, 0}}))
}
