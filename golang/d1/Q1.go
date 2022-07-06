package main

import "fmt"

type Matrix struct {
	rowCount, colsCount int
	mx                  [][]int
}

func (M *Matrix) getRows() int {
	return M.rowCount
}

func (M *Matrix) getCols() int {
	return M.colsCount
}

func (M *Matrix) set(x, y, val int) {
	if x >= 0 && x < M.getRows() && y >= 0 && y < M.getCols() {
		M.mx[x][y] = val
	}
}

func (M *Matrix) add(T *Matrix) {
	for i := 0; i < M.getRows(); i++ {
		for j := 0; j < M.getCols(); j++ {
			M.mx[i][j] += T.mx[i][j]
		}
	}
}

func (M *Matrix) print() {
	fmt.Print("{")
	for i := 0; i < M.getRows(); i++ {
		fmt.Print("{")
		for j := 0; j < M.getCols(); j++ {
			fmt.Print(M.mx[i][j])
			if j != M.getCols()-1 {
				fmt.Print(", ")
			}
		}

		fmt.Print("}")
		if i != M.getRows()-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("}")
}

func NewMatrix(rowsCount, colsCount int) *Matrix {
	m := new(Matrix)
	m.rowCount = rowsCount
	m.colsCount = colsCount
	m.mx = make([][]int, rowsCount)
	for i := 0; i < rowsCount; i++ {
		m.mx[i] = make([]int, colsCount)
	}
	return m
}

func main() {
	m1 := NewMatrix(4, 4)
	for i := 0; i < m1.getRows(); i++ {
		for j := 0; j < m1.getCols(); j++ {
			m1.set(i, j, i+j)
		}
	}
	m2 := NewMatrix(4, 4)
	for i := 0; i < m2.getRows(); i++ {
		for j := 0; j < m2.getCols(); j++ {
			m2.set(i, j, i+j)
		}
	}
	m1.add(m2)
	m2.print()
	m1.print()
}
