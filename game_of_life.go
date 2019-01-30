package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func randomizeGrid(grid [][]int) [][]int {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2)
		}
	}
	return grid
}

func cleanScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func prettyPrint(grid [][]int) {
	cleanScreen()
	var buffer bytes.Buffer
	for i := range grid {
		for j := range grid[i] {
			buffer.WriteString(strconv.Itoa(grid[i][j]) + " ")
		}
		buffer.WriteString("\n")
	}
	fmt.Printf("\033[2K\r%s", buffer.String())
	time.Sleep(10e7 * time.Nanosecond)
}

func getAliveCells(grid [][]int, i int, j int) int {
	var aliveCells int
	var xTemp int
	var yTemp int
	//fmt.Println("Start")
	//fmt.Printf("x: %d y: %d\n", i, j)
	for x := i - 1; x < i+2; x++ {
		for y := j - 1; y < j+2; y++ {
			if x == i && y == j {
				continue
			}

			switch {
			case x < 0:
				xTemp = len(grid) - 1
			case x == len(grid):
				xTemp = 0
			default:
				xTemp = x
			}

			switch {
			case y < 0:
				yTemp = len(grid) - 1
			case y == len(grid):
				yTemp = 0
			default:
				yTemp = y
			}

			//fmt.Printf("x: %d y: %d\n", xTemp, yTemp)
			if grid[xTemp][yTemp] == 1 {
				aliveCells++
			}
		}
	}
	return aliveCells
}

func nextStage(grid [][]int) [][]int {
	buffer := make([][]int, len(grid))
	for i := range buffer {
		buffer[i] = make([]int, len(grid))
	}
	for i := range grid {
		for j := range grid[i] {
			aliveCells := getAliveCells(grid, i, j)
			//fmt.Println(aliveCells)
			switch {
			case aliveCells == 3:
				buffer[i][j] = 1
			case aliveCells == 2:
				buffer[i][j] = grid[i][j]
			case aliveCells > 3:
				buffer[i][j] = 0
			case aliveCells < 2:
				buffer[i][j] = 0
			}

		}
	}
	return buffer
}

func glider(grid [][]int) [][]int {
	lenght := len(grid) - 1
	grid[lenght-1][lenght] = 1
	grid[lenght-2][lenght-1] = 1
	grid[lenght][lenght-2] = 1
	grid[lenght-1][lenght-2] = 1
	grid[lenght-2][lenght-2] = 1

	return grid
}

func main() {
	var height = 60
	var width = 60
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	grid = randomizeGrid(grid)
	//grid = glider(grid)
	prettyPrint(grid)
	for i := 0; i < 10000; i++ {
		grid = nextStage(grid)
		prettyPrint(grid)
	}
	fmt.Println("Done")
	for i := 10; i >= 0; i-- {
		fmt.Printf("\033[2K\r%d", i)
		time.Sleep(1 * time.Second)
	}
}
