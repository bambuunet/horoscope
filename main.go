package main

import (
  "fmt"
  "math"
)

var MJD float64
var EarthXYZ XYZ

func main(){
  /*datetime := Datetime{
    year: 2008,
    month: 1,
    day: 1,
    hour: 0,
    minute: 0,
  }*/
  datetime := Datetime{
    year: 1906,
    month: 3,
    day: 20,
    hour: 22,
    minute: 30,
  }

  setMJD(datetime)
  setEarthXYZ()

  fmt.Printf("%v\n\n", MJD)
  fmt.Printf("%v\n\n", EarthXYZ)
  fmt.Printf("%v\n\n\n", math.Atan2(EarthXYZ.y, EarthXYZ.x) * 180 / math.Pi + 180)
  

  MercuryXYZ := getHeliocentricXYZ(Mercury)
  MercuryAngle := getGeocentricAngle(MercuryXYZ)
  fmt.Printf("%v\n", MercuryXYZ)
  fmt.Printf("%v\n\n\n", MercuryAngle)

  VenusXYZ := getHeliocentricXYZ(Venus)
  VenusAngle := getGeocentricAngle(VenusXYZ)
  fmt.Printf("%v\n", VenusXYZ)
  fmt.Printf("%v\n\n\n", VenusAngle)


  MarsXYZ := getHeliocentricXYZ(Mars)
  MarsAngle := getGeocentricAngle(MarsXYZ)
  fmt.Printf("%v\n", MarsXYZ)
  fmt.Printf("%v\n\n\n", MarsAngle)

  NeptuneXYZ := getHeliocentricXYZ(Neptune)
  NeptuneAngle := getGeocentricAngle(NeptuneXYZ)
  fmt.Printf("%v\n", NeptuneXYZ)
  fmt.Printf("%v\n\n\n", NeptuneAngle)
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
  EarthXYZ = getHeliocentricXYZ(Earth)
}

/*
  日心座標から地心黄経を取得
  @param 修正ユリウス日、惑星座標
  @return 座標
*/
func getGeocentricAngle(planetXYZ XYZ) float64{
  x := planetXYZ.x - EarthXYZ.x
  y := planetXYZ.y - EarthXYZ.y
  angle := math.Atan2(y, x) * 180 / math.Pi
  if angle < 0{
    angle += 360
  }
  return angle
}

/*
  修正ユリウス日から日心座標を取得
  @param 修正ユリウス日、惑星諸情報
  @return 座標
*/
func getHeliocentricXYZ(planet Planet) XYZ{
  //軌道長半径a
  a := planet.semiMajorAxis

  //近日点通過時M
  /*perihelionPassageMJD := getMJD(
    Datetime{
      year: planet.perihelionPassageMJD.year,
      month: planet.perihelionPassageMJD.month,
      day: planet.perihelionPassageMJD.day,
      hour: planet.perihelionPassageMJD.hour,
      minute: planet.perihelionPassageMJD.minute,
    },
  )*/

  //平均日々運動(度)
  //meanMotion := 0.985647365 * math.Pow(semiMajorAxis, -1.5)
  meanMotion := 360 / 365.242189 / planet.orbitalPeriod
  //meanMotion := 360 / 365.25 / planet.orbitalPeriod

  //近日点引数ω
  ω := planet.perihelionArgument * math.Pi / 180

  //離心率e
  e := planet.eccentricity

  //平均近点離角。
  //近日点通過時からの経過日数に比例し
  //近日点通過時に0度、遠日点通過時に180度となる。
  meanAnomaly := (meanMotion * (MJD - getMJD(planet.epoch)) + planet.meanAnomalyAtEpoch) * math.Pi / 180
  fmt.Print(meanAnomaly / math.Pi * 180,"\n")

  //離心近点角E
  E := getEccentricAnomaly(meanAnomaly - ω, e)
  fmt.Print(E * 180 / math.Pi,"\n")

  //昇交点黄経Ω
  Ω := planet.longitudeOfAscendingNode * math.Pi / 180

  //軌道傾斜角i
  i := planet.inclination * math.Pi / 180

  //計算
  Ex := a * (math.Cos(E) - e)
  Ey := a * math.Sqrt(1 - e * e) * math.Sin(E)
  cosi := math.Cos(i)
  sini := math.Sin(i)
  cosΩ := math.Cos(Ω)
  sinΩ := math.Sin(Ω)
  cosω := math.Cos(ω - Ω)
  sinω := math.Sin(ω - Ω)

  x := Ex * (cosΩ * cosω - sinΩ * cosi * sinω) - Ey * (cosΩ * sinω + sinΩ * cosi * cosω)
  y := Ex * (sinΩ * cosω + cosΩ * cosi * sinω) - Ey * (sinΩ * sinω - cosΩ * cosi * cosω)
  z := Ex * (sini * sinω) + Ey * (sini * cosω)

  return XYZ{
    x: x,
    y: y,
    z: z,
  }


  //行列計算
  /*matrix1 := [][]float64{
    {math.Cos(Ω), - math.Sin(Ω), 0},
    {math.Sin(Ω), math.Cos(Ω), 0},
    {0, 0, 1},
  }
  matrix2 := [][]float64{
    {1, 0},
    {0, math.Cos(i)},
    {0, math.Sin(i)},
  }
  matrix3 := [][]float64{
    {math.Cos(ω - Ω), - math.Sin(ω - Ω)},
    {math.Sin(ω - Ω), math.Cos(ω - Ω)},
  }
  matrix4 := [][]float64{
    {a * math.Cos(E) - a * e},
    {a * math.Sqrt(1 - math.Pow(e, 2)) * math.Sin(E)},
  }


  fmt.Print(matrix1,"\n")
  fmt.Print(matrix2,"\n")
  fmt.Print(matrix3,"\n")
  fmt.Print(matrix4,"\n\n")

  xyz := dotMatrix(matrix3, matrix4)
  fmt.Print(xyz," 1\n")
  xyz = dotMatrix(matrix2, xyz)
  fmt.Print(xyz," 2\n")
  xyz = dotMatrix(matrix1, xyz)

  fmt.Print(math.Atan2(xyz[1][0], xyz[0][0]) * 180 / math.Pi, "\n")

  return XYZ{
    x: xyz[0][0],
    y: xyz[1][0],
    z: xyz[2][0],
  }*/
}

/*
  ケプラー方程式を解く
  @param 平均近点角(Mean Anomaly)、軌道離心率(Eccentricity)
  @return 離心近点角(Eccentric Anomaly)
*/
func getEccentricAnomaly(M float64, e float64) float64 {
  E0 := M
  count := 0
  for {
    deltaE := (M - E0 + e * math.Sin(E0)) / (1 - e * math.Cos(E0))
    E := E0 + deltaE
    E0 = E

    count += 1
    if deltaE < 0.0000000001 || count > 100{
      return surplusFloat(E, 2 * math.Pi)
    }
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


