package main

import (
  "fmt"
  "math"
)

func main(){
  mjd := getMJD(2009, 12, 31, 11, 59)
  xyz := getXYZ(mjd)
  fmt.Printf("%v\n", mjd)
  fmt.Printf("%v\n", xyz)
}

/*
  日時を修正ユリウス日に変換
  @param 日時
  @return 修正ユリウス日
*/
func getMJD(year, month, day, hour, minute int) float64 {
  //1月,2月は前年の13月,14月とする
  if month == 1 || month == 2 {
    year -= 1
    month += 12
  }

  result := math.Floor(365.25 * float64(year))
  result += math.Floor(float64(year) / 400)
  result -= math.Floor(float64(year) / 100)
  result += math.Floor(30.59 * (float64(month) - 2))
  result += float64(day)
  result += float64(hour) / 24
  result += float64(minute) / 1440
  result -= 678912
  return result
}

/*
  修正ユリウス日から座標を取得
  @param 修正ユリウス日
  @return 座標
*/
func getXYZ(mjd float64) float64{
  //軌道長半径
  //semiMajorAxis := 0.3871//いったん水星

  //近日点通過時M
  perihelionPassageMJD := getMJD(2018, 3, 10, 10, 58)//いったん水星

  //平均日々運動
  //meanMotion := 0.985647365 * math.Pow(semiMajorAxis, -1.5)
  meanMotion := 360 / 365.24219 / 0.2408467

  //近日点引数ω
  perihelionArgument := 77.5806

  //平均近点離角。
  //近日点通過時からの経過日数に比例し
  //近日点通過時に0度、遠日点通過時に180度となる。
  meanAnomaly := surplusFloat(meanMotion * (mjd - perihelionPassageMJD), 360)

  //日心黄経
  longitudeOfHeliocentric := surplusFloat(meanAnomaly + perihelionArgument, 360)

  //離心率e
  eccentricity := 0.20563069

  //昇交点黄経Ω
  longitudeOfAscendingNode := 48.4257

  //軌道傾斜角i
  inclination := 7.0051

  //行列計算
  matrix1 := [3][3]{
    {math.Cos(longitudeOfAscendingNode), math.Sin(longitudeOfAscendingNode), 0},
    {math.Sin(longitudeOfAscendingNode), math.Cos(longitudeOfAscendingNode), 0},
    {0, 0, 1},
  }
  matrix2 := [3][2]{
    {1, 0},
    {0, math.Cos(inclination)},
    {0, math.Sin(inclination)},
  }
  matrix3 := [2][2]{
    {math.Cos(perihelionArgument), - math.Sin(perihelionArgument)},
    {math.Sin(perihelionArgument), math.Cos(perihelionArgument)},
  }
  matrix4 := [2][1]{
    {},
    {},
  }

  return 
}



