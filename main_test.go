package main

import (
  "testing"
  "fmt"
  //"math"
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

  xyz := getHeliocentricXYZ(Mercury)

  fmt.Print("TestGetHeliocentricXYZ: \n")
  fmt.Print(xyz)
  fmt.Print("\n\n")
}

func TestGetGeocentricAngle(t *testing.T){
  xyz := XYZ{
    x: 1,
    y: -1,
    z: 0,
  }

  EarthXYZ = XYZ{
    x: 0,
    y: 0,
    z: 0,
  }

  angle := getGeocentricAngle(xyz)

  fmt.Print("TestGetGeocentricAngle: \n")
  fmt.Print(angle)
  fmt.Print("\n\n")
}

func TestGetEccentricAnomaly(t *testing.T){
  M := 3.14159
  e := 0.09341233
  E := getEccentricAnomaly(M, e)

  fmt.Print("TestGetEccentricAnomaly: \n")
  fmt.Print(E)
  fmt.Print("\n\n")
}