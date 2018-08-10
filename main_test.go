package main

import (
  "testing"
  "fmt"
  "math"
)

func TestGetHeliocentricXYZ(t *testing.T){
  MJD = getMJD(
    Datetime{
      year: 2009,
      month: 12,
      day: 31,
      hour: 11,
      minute: 59,
    },
  )

  xyz := getHeliocentricXYZ(mercury)
  fmt.Print("TestGetHeliocentricXYZ: ")
  fmt.Print(xyz)
  fmt.Print("\n")
}

func TestGetGeocentricAngle(t *testing.T){
  xyz := XYZ{
    x: 1,
    y: 1,
    z: 0,
  }
  earthXYZ = XYZ{
    x: 0,
    y: 0,
    z: 0,
  }
  fmt.Print(math.Atan2(1, 1))
  angle := getGeocentricAngle(xyz)
  fmt.Print("TestGetGeocentricAngle: ")
  fmt.Print(angle)
  fmt.Print("\n")
}