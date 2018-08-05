package main

import (
  "fmt"
  "reflect"
  "math"
)

func main(){
  mjd := getMJD(2018, 3, 9, 10, 58)
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
  semiMajorAxis := 0.3871//いったん水星

  //近日点通過時
  perihelionPassageMJD := getMJD(2018, 3, 10, 10, 58)//いったん水星

  //平均日々運動
  meanMotion := 0.985647365 * math.Pow(semiMajorAxis, -1.5)

  //平均近点離角。
  //近日点通過時からの経過日数に比例し
  //近日点通過時に0度、遠日点通過時に180度となる。
  meanAnomaly := surplusFloat(meanMotion * (mjd - perihelionPassageMJD), 360)

  return meanAnomaly
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
