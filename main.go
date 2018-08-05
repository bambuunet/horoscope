package main

import (
  "fmt"
)

func main(){
  mjd := getMJD(2018, 12, 22, 12, 12)
  fmt.Printf("%v", mjd)
}

/*
  日時を修正ユリウス日に変換
  @param 日時
  @return 修正ユリウス日
*/
func getMJD(year, month, day, hour, minute int) float64 {
  return 2.2
}


/*
  修正ユリウス日を
  @param 
  @return
*/
