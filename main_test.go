package main

import (
  "testing"
  "fmt"
)

func TestGetXYZ(t *testing.T){
  mjd := getMJD(2009, 12, 31, 11, 59)
  xyz := getXYZ(mjd, mercury)
  fmt.Print(xyz)
  fmt.Print("\n")
}