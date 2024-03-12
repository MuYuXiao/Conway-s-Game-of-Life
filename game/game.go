package game

import (
	"math/rand"
	"sync"
)

type Cell bool

type Board [][]Cell

func NewBoard(rows, cols int) Board {
	board := make(Board, rows)
	for i := range board {
		board[i] = make([]Cell, cols)
	}
	return board
}

func (b Board) RandBoard() {
	for x, row := range b {
		for y := range row {
			if rand.Float32() >= 0.5 {
				b[x][y] = true
			}
		}
	}
}

func (b Board) Update() {
	newBoard := NewBoard(len(b), len(b[0]))
	var wg sync.WaitGroup
	for x, row := range b {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := range row {
				aliveNeighbors := b.countAliveNeighbors(x, y)
				if b[x][y] {
					if aliveNeighbors == 2 || aliveNeighbors == 3 {
						newBoard[x][y] = true
					}
				} else {
					if aliveNeighbors == 3 {
						newBoard[x][y] = true
					}
				}
			}
		}()
		wg.Wait()
	}
	copy(b, newBoard)
}

func (b Board) countAliveNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newX, newY := x+i, y+j
			if newX >= 0 && newX < len(b) && newY >= 0 && newY < len(b[0]) {
				if b[newX][newY] {
					count++
				}
			}
		}
	}
	return count
}
