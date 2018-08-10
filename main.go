package main

import (
  "fmt"
  "math"
)

var MJD float64
var earthXYZ XYZ

func main(){
  datetime := Datetime{
    year: 2018,
    month: 12,
    day: 31,
    hour: 11,
    minute: 59,
  }

  setMJD(datetime)
  setEarthXYZ()

  mercuryXYZ := getHeliocentricXYZ(mercury)

  fmt.Printf("%v\n", MJD)
  fmt.Printf("%v\n", mercuryXYZ)
}

/*
  指定の修正ユリウス日をセット
  @param 年月日
*/
func setMJD(dt Datetime) {
  MJD = getMJD(dt)
}

/*
  地球の日心座標をセット
  @param 修正ユリウス日
*/
func setEarthXYZ() {
  earthXYZ = getHeliocentricXYZ(earth)
}

/*
  日心座標から地心黄経を取得
  @param 修正ユリウス日、惑星座標
  @return 座標
*/
func getGeocentricAngle(planetXYZ XYZ) float64{
  x := planetXYZ.x - earthXYZ.x
  y := planetXYZ.y - earthXYZ.y
  return math.Atan2(y, x)
}

/*
  修正ユリウス日から日心座標を取得
  @param 修正ユリウス日、惑星諸情報
  @return 座標
*/
func getHeliocentricXYZ(planet Planet) XYZ{
  //軌道長半径
  semiMajorAxis := planet.semiMajorAxis//いったん水星

  //近日点通過時M
  perihelionPassageMJD := getMJD(
    Datetime{
      year: planet.MJD.year,
      month: planet.MJD.month,
      day: planet.MJD.day,
      hour: planet.MJD.hour,
      minute: planet.MJD.minute,
    },
  )

  //平均日々運動
  //meanMotion := 0.985647365 * math.Pow(semiMajorAxis, -1.5)
  meanMotion := 360 / 365.242189 / planet.orbitalPeriod

  //近日点引数ω
  perihelionArgument := planet.perihelionArgument

  //平均近点離角。
  //近日点通過時からの経過日数に比例し
  //近日点通過時に0度、遠日点通過時に180度となる。
  meanAnomaly := surplusFloat(meanMotion * (MJD - perihelionPassageMJD), 360)

  //日心黄経
  longitudeOfHeliocentric := surplusFloat(meanAnomaly + perihelionArgument, 360)

  //離心率e
  eccentricity := planet.eccentricity

  //昇交点黄経Ω
  longitudeOfAscendingNode := planet.longitudeOfAscendingNode

  //軌道傾斜角i
  inclination := planet.inclination

  //行列計算
  matrix1 := [][]float64{
    {math.Cos(longitudeOfAscendingNode), math.Sin(longitudeOfAscendingNode), 0},
    {math.Sin(longitudeOfAscendingNode), math.Cos(longitudeOfAscendingNode), 0},
    {0, 0, 1},
  }
  matrix2 := [][]float64{
    {1, 0},
    {0, math.Cos(inclination)},
    {0, math.Sin(inclination)},
  }
  matrix3 := [][]float64{
    {math.Cos(perihelionArgument), - math.Sin(perihelionArgument)},
    {math.Sin(perihelionArgument), math.Cos(perihelionArgument)},
  }
  matrix4 := [][]float64{
    {semiMajorAxis * math.Sqrt(1 - math.Pow(eccentricity, 2)) * math.Cos(longitudeOfHeliocentric) - semiMajorAxis * eccentricity},
    {semiMajorAxis * math.Sqrt(1 - math.Pow(eccentricity, 2)) * math.Sin(longitudeOfHeliocentric)},
  }
  xyz := dotMatrix(matrix1, matrix2)
  xyz = dotMatrix(xyz, matrix3)
  xyz = dotMatrix(xyz, matrix4)

  return XYZ{
    x: xyz[0][0],
    y: xyz[1][0],
    z: xyz[2][0],
  }
}

/*
  日時を修正ユリウス日に変換
  @param 日時
  @return 修正ユリウス日
*/
func getMJD(dt Datetime) float64 {
  //1月,2月は前年の13月,14月とする
  if dt.month == 1 || dt.month == 2 {
    dt.year -= 1
    dt.month += 12
  }

  result := math.Floor(365.25 * float64(dt.year))
  result += math.Floor(float64(dt.year) / 400)
  result -= math.Floor(float64(dt.year) / 100)
  result += math.Floor(30.59 * (float64(dt.month) - 2))
  result += float64(dt.day)
  result += float64(dt.hour) / 24
  result += float64(dt.minute) / 1440
  result -= 678912
  return result
}


