package main

type Planet struct{
  semiMajorAxis float64
  MJD Datetime
  orbitalPeriod float64
  perihelionArgument float64
}

type Datetime struct{
  year int
  month int
  day int
  hour int
  minute int
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
  orbitalPeriod: 0.2408467,
  perihelionArgument: 77.5806,
}

