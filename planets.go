package main

type Planet struct{
  semiMajorAxis float64
  perihelionPassageMJD Datetime
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

var Mercury = Planet{
  semiMajorAxis: 0.3871, //a
  perihelionPassageMJD: Datetime{
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

var Earth = Planet{
  semiMajorAxis: 1, //a
  perihelionPassageMJD: Datetime{
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

var Mars = Planet{
  semiMajorAxis: 1.52368, //a
  perihelionPassageMJD: Datetime{
    year: 2018,
    month: 9,
    day: 16,
    hour: 12,
    minute: 53,
  },
  orbitalPeriod: 1.880866, //P
  perihelionArgument: 336.2075, //ω
  eccentricity: 0.09341233, //e
  longitudeOfAscendingNode: 49.6198, //Ω
  inclination: 1.8497, //i
}

var Neptune = Planet{
  semiMajorAxis: 30.06896348, //a
  perihelionPassageMJD: Datetime{
    year: 2051,
    month: 1,
    day: 1,
    hour: 0,
    minute: 0,
  },
  orbitalPeriod: 164.79, //P
  perihelionArgument: 44.97135, //ω
  eccentricity:   0.00858587, //e
  longitudeOfAscendingNode: 131.72169, //Ω
  inclination:  1.769, //i
}