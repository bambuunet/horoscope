package main

type Planet struct{
  semiMajorAxis float64
  epoch Datetime
  meanAnomalyAtEpoch float64 //M0
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
  epoch: Datetime{
    year: 2008,
    month: 1,
    day: 1,
    hour: 12,
    minute: 0,
  },
  meanAnomalyAtEpoch: 328.1305,
  orbitalPeriod: 0.2408467, //P
  perihelionArgument: 77.5806, //ω
  eccentricity: 0.20563069, //e
  longitudeOfAscendingNode: 48.4257, //Ω
  inclination: 7.0051, //i
}

var Venus = Planet{
  semiMajorAxis: 0.72333199, //a
  epoch: Datetime{
    year: 2008,
    month: 1,
    day: 1,
    hour: 12,
    minute: 0,
  },
  meanAnomalyAtEpoch: 182.7158,
  orbitalPeriod: 0.615207, //P
  perihelionArgument: 131.6758, //ω
  eccentricity: 0.00677323, //e
  longitudeOfAscendingNode: 76.7520, //Ω
  inclination: 3.39471, //i
}

var Earth = Planet{
  semiMajorAxis: 1, //a
  epoch: Datetime{
    year: 2016,
    month: 3,
    day: 20,
    hour: 4,
    minute: 30,
  },
  meanAnomalyAtEpoch: 180 - 2.6,
  orbitalPeriod: 1, //P
  perihelionArgument: 102.992, //ω
  eccentricity: 0.01671022, //e
  longitudeOfAscendingNode: 174.838, //Ω
  inclination:  0.002, //i
}

/*var Earth = Planet{
  semiMajorAxis: 1, //a
  epoch: Datetime{
    year: 2000,
    month: 1,
    day: 1,
    hour: 12,
    minute: 0,
  },
  meanAnomalyAtEpoch: 100.45,
  orbitalPeriod: 1, //P
  perihelionArgument: 102.992, //ω
  eccentricity: 0.01671022, //e
  longitudeOfAscendingNode: 174.838, //Ω
  inclination:  0.002, //i
}*/

var Mars = Planet{
  semiMajorAxis: 1.52368, //a
  epoch: Datetime{
    year: 2008,
    month: 1,
    day: 1,
    hour: 12,
    minute: 0,
  },
  meanAnomalyAtEpoch: 86.5067,
  orbitalPeriod: 1.880866, //P
  perihelionArgument: 336.2075, //ω
  eccentricity: 0.09341233, //e
  longitudeOfAscendingNode: 49.6198, //Ω
  inclination: 1.8497, //i
}


var Neptune = Planet{
  semiMajorAxis: 30.06896348, //a
  epoch: Datetime{
    year: 2000,
    month: 1,
    day: 1,
    hour: 12,
    minute: 0,
  },
  meanAnomalyAtEpoch: 304.88003,
  orbitalPeriod: 164.79, //P
  perihelionArgument: 44.97135, //ω
  eccentricity:   0.00858587, //e
  longitudeOfAscendingNode: 131.72169, //Ω
  inclination:  1.769, //i
}