package main

import (
  "testing"
  "fmt"
)

func TestDotMatrix(t *testing.T){
  matrix1 := [][]float64{{1, 10}, {2, 1}, {1, 0}, {1, 1}}
  matrix2 := [][]float64{{1, 2, 0, 1}, {1, 0, 1, 1}}
  result := dotMatrix(matrix1, matrix2)
  fmt.Print(result)
}