package SwissEph

const(
  SUNDAY=0
  MONDAY=1
  TUESDAY=2
  WEDNESDAY=3
  THURSDAY=4
  FRIDAY=5
  SATURDAY=6

  SE_JUL_CAL=false
  SE_GREG_CAL=true
  SE_KEEP_DATE=true
  SE_KEEP_JD=false

// for delta t: tidal acceleration in the mean motion of the moon
  /**
  * Tidal acceleration value in the mean motion of the moon of DE403 (-25.8).
  */
  SE_TIDAL_DE403=-25.8
  /**
  * Tidal acceleration value in the mean motion of the moon of DE404 (-25.8).
  */
  SE_TIDAL_DE404=-25.8
  /**
  * Tidal acceleration value in the mean motion of the moon of DE405 (-25.7376).
  */
  SE_TIDAL_DE405=-25.7376
  /**
  * Tidal acceleration value in the mean motion of the moon of DE406 (-25.7376).
  */
  SE_TIDAL_DE406=-25.7376
  /**
  * Tidal acceleration value in the mean motion of the moon of DE200 (-23.8946).
  */
  SE_TIDAL_DE200=-23.8946
  /**
  * Tidal acceleration value in the mean motion of the moon of -26.
  */
  SE_TIDAL_26=-26.0
  /**
  * Default tidal acceleration value in the mean motion of the moon (=SE_TIDAL_DE406).
  * @see #SE_TIDAL_DE406
  */
  SE_TIDAL_DEFAULT=SE_TIDAL_DE406

  /**
  * The Julian day number of 1970 January 1.0. Useful for conversions
  * from or to a Date object.
  * @see #getDate(long)
  */
  JD0=2440587.5          /* 1970 January 1.0 */
)

init_leapseconds_done := false
tid_acc := SE_TIDAL_DEFAULT
init_dt_done := false
var jd float
// JD for the start of the Gregorian calendar system (October 15, 1582):
jdCO := 2299160.5
calType bool
private int year
private int month
private int day
private double hour
private double deltaT
private boolean deltatIsValid=false


type SweDate struct{
  SweDateMethods
}

type SweDateMethods interface{
  getDeltaT()
}

func getDeltaT(tjd_ut) float{
  double sdt = calc_deltaT(tjd, SE_TIDAL_DEFAULT)
  return sdt
}