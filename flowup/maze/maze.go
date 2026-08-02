package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

var dirs = [4]point{
	{-1, 0}, // 上
	{0, -1}, // 左
	{1, 0},  // 下
	{0, 1},  // 右
}

func readMaze(filename string) [][]int {
	file, _ := os.Open(filename)
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range col {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		for _, dir := range dirs {
			next := cur.add(dir)
			// maze at next is 0
			// and steps at next is 0
			// and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 { // 坐标不存在或者撞墙了
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 { // 坐标不存在或者已经走过了
				continue
			}
			if next == start { // 回到原点
				continue
			}
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1

			if next == end {
				return steps
			}

			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("flowup/maze/maze.txt")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
