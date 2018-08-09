package main

import (
  "reflect"
  "math"
  //"fmt"
)

/*
  行列のドット積を取得
  @param 行列1、行列2
  @return 行列1と行列2のドット積の配列
*/
func dotMatrix(matrix1 [][]float64, matrix2 [][]float64) [][]float64{
  result := make([][]float64, len(matrix1))
  for i1 := 0; i1 < len(matrix1); i1++ {
    result[i1] = make([]float64, len(matrix2[0]))
    for i2 := 0; i2 < len(matrix2[0]); i2++ {
      for i3 := 0; i3 < len(matrix1[0]); i3++ {
        result[i1][i2] += matrix1[i1][i3] * matrix2[i3][i2]
      }
    }
 }
 return result
}

/*
  float型の剰余を取得
  @param 割られる数、割る数
  @return 剰余
*/
func surplusFloat(dividend float64, divisor interface{}) float64{

  r := reflect.ValueOf(divisor)
  var divisor2 float64
  switch r.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
      divisor2 = float64(r.Int())
    default:
      divisor2 = r.Float()
  }

  return dividend - math.Floor(dividend / divisor2) * divisor2
}