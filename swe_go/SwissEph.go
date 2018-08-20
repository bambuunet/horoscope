package SwissEph

func swe_calc_ut(tjd_ut float, ipl int, iflag int, []xx float, serr string) int{
  return swe_calc(tjd_ut + SweDate.getDeltaT(tjd_ut), ipl, iflag, xx, serr)
}

func swe_calc() int{
 
}


