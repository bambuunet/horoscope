package main

type Planet struct{
  semiMajorAxis float64
  MJD Datetime
  orbitalPeriod float64 //P
  perihelionArgument float64 //ω
  eccentricity float64 //e
  longitudeOfAscendingNode float64 //Ω
  inclination float64 //i
}

type Datetime struct{
  year int
  month int
  day int
  hour int
  minute int
}

type XYZ struct{
  x float64
  y float64
  z float64
}

var mercury = Planet{
  semiMajorAxis: 0.3871,
  MJD: Datetime{
    year: 2018,
    month: 3,
    day: 10,
    hour: 10,
    minute: 58,
  },
  orbitalPeriod: 0.2408467, //P
  perihelionArgument: 77.5806, //ω
  eccentricity: 0.20563069, //e
  longitudeOfAscendingNode: 48.4257, //Ω
  inclination: 7.0051, //i
}

var earth = Planet{
  semiMajorAxis: 1,
  MJD: Datetime{
    year: 2018,
    month: 1,
    day: 3,
    hour: 5,
    minute: 35,
  },
  orbitalPeriod: 1, //P
  perihelionArgument: 102.992, //ω
  eccentricity: 0.01671022, //e
  longitudeOfAscendingNode: 174.838, //Ω
  inclination:  0.002, //i
}