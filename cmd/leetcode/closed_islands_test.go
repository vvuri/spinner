package leetcode

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
)

var maps [][]int
var la, lb, islands int

func findIsland(i int, j int) bool {
	isIsland := true
	//if maps[i][j] > 2 {
	//	return true
	//}

	maps[i][j] = 3 + islands

	if i == 0 || i == la-1 || j == 0 || j == lb-1 {
		isIsland = false
	}

	if j != lb-1 && maps[i][j+1] == 0 {
		if !findIsland(i, j+1) {
			isIsland = false
		}
	}

	if i != la-1 && maps[i+1][j] == 0 {
		if !findIsland(i+1, j) {
			isIsland = false
		}
	}

	if j != 0 && maps[i][j-1] == 0 {
		if !findIsland(i, j-1) {
			isIsland = false
		}
	}

	if i != 0 && maps[i-1][j] == 0 {
		if !findIsland(i-1, j) {
			isIsland = false
		}
	}

	return isIsland
}

func printMap() {
	var st string
	for i := 0; i < la; i++ {
		for j := 0; j < lb; j++ {
			st = st + strconv.Itoa(maps[i][j])
		}
		log.Println(st)
		st = ""
	}
	log.Println()
}

func closedIsland(grid [][]int) int {
	maps = grid
	la, lb = len(maps), len(maps[0])
	islands = 0

	for i := 1; i < la-1; i++ {
		for j := 1; j < lb-1; j++ {
			if maps[i][j] == 0 {
				if findIsland(i, j) {
					islands++
					printMap()
				}
			}
		}
	}

	printMap()

	return islands
}

func TestClosedIslandsOne(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
	}
	assert.Equal(t, closedIsland(grid), 1)
}

func TestClosedIslandsTwo(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}
	assert.Equal(t, closedIsland(grid), 2)
}

func TestClosedIslands19(t *testing.T) {
	grid := [][]int{
		{0, 1, 1, 1, 0},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}
	assert.Equal(t, closedIsland(grid), 1)
}

func TestClosedIslands0(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 1, 0, 0, 0, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 1, 1, 1, 1, 0},
	}
	assert.Equal(t, closedIsland(grid), 0)
}

func TestClosedIslands1(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 1, 1, 1, 1, 0, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 1, 1, 1, 1, 0},
	}
	assert.Equal(t, closedIsland(grid), 1)
}

func TestClosedIslands40(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1},
		{0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0},
		{0, 1, 0, 1, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1},
		{1, 1, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0},
		{0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1},
		{1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0},
		{0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
		{0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0},
	}
	assert.Equal(t, closedIsland(grid), 5)
}
