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
var calType bool
var year int
var month int
var day int
var hour float
var deltaT float
deltatIsValid := false


type SweDate struct{
  SweDateMethods
}

type SweDateMethods interface{
  getDeltaT()
}

func getDeltaT(tjd_ut) float{
  return calc_deltaT(tjd, SE_TIDAL_DEFAULT)
}


////////////////////////////////////////////////////////////////////////////
/// deltaT:
////////////////////////////////////////////////////////////////////////////
/* DeltaT = Ephemeris Time - Universal Time, in days.
 * 
 * 1620 - today + a couple of years:
 * ---------------------------------
 * The tabulated values of deltaT, in hundredths of a second,
 * were taken from The Astronomical Almanac 1997, page K8.  The program
 * adjusts for a value of secular tidal acceleration ndot = -25.7376.
 * arcsec per century squared, the value used in JPL's DE403 ephemeris.
 * ELP2000 (and DE200) used the value -23.8946.
 * To change ndot, one can
 * either redefine SE_TIDAL_DEFAULT in swephexp.h
 * or use the routine swe_set_tid_acc() before calling Swiss 
 * Ephemeris.
 * Bessel's interpolation formula is implemented to obtain fourth 
 * order interpolated values at intermediate times.
 *
 * -1000 - 1620:
 * ---------------------------------
 * For dates between -500 and 1600, the table given by Morrison &
 * Stephenson (2004; p. 332) is used, with linear interpolation.
 * This table is based on an assumed value of ndot = -26.
 * The program adjusts for ndot = -25.7376.
 * For 1600 - 1620, a linear interpolation between the last value
 * of the latter and the first value of the former table is made.
 *
 * before -1000:
 * ---------------------------------
 * For times before -1100, a formula of Morrison & Stephenson (2004)
 * (p. 332) is used:
 * dt = 32 * t * t - 20 sec, where t is centuries from 1820 AD.
 * For -1100 to -1000, a transition from this formula to the Stephenson
 * table has been implemented in order to avoid a jump.
 *
 * future:
 * ---------------------------------
 * For the time after the last tabulated value, we use the formula
 * of Stephenson (1997; p. 507), with a modification that avoids a jump
 * at the end of the tabulated period. A linear term is added that
 * makes a slow transition from the table to the formula over a period
 * of 100 years. (Need not be updated, when table will be enlarged.)
 *
 * References:
 *
 * Stephenson, F. R., and L. V. Morrison, "Long-term changes
 * in the rotation of the Earth: 700 B.C. to A.D. 1980,"
 * Philosophical Transactions of the Royal Society of London
 * Series A 313, 47-70 (1984)
 *
 * Borkowski, K. M., "ELP2000-85 and the Dynamical Time
 * - Universal Time relation," Astronomy and Astrophysics
 * 205, L8-L10 (1988)
 * Borkowski's formula is derived from partly doubtful eclipses 
 * going back to 2137 BC and uses lunar position based on tidal 
 * coefficient of -23.9 arcsec/cy^2.
 *
 * Chapront-Touze, Michelle, and Jean Chapront, _Lunar Tables
 * and Programs from 4000 B.C. to A.D. 8000_, Willmann-Bell 1991
 * Their table agrees with the one here, but the entries are
 * rounded to the nearest whole second.
 *
 * Stephenson, F. R., and M. A. Houlden, _Atlas of Historical
 * Eclipse Maps_, Cambridge U. Press (1986)
 *
 * Stephenson, F.R. & Morrison, L.V., "Long-Term Fluctuations in 
 * the Earth's Rotation: 700 BC to AD 1990", Philosophical 
 * Transactions of the Royal Society of London, 
 * Ser. A, 351 (1995), 165-202. 
 *
 * Stephenson, F. Richard, _Historical Eclipses and Earth's 
 * Rotation_, Cambridge U. Press (1997)
 * 
 * Morrison, L. V., and F.R. Stephenson, "Historical Values of the Earth's
 * Clock Error DT and the Calculation of Eclipses", JHA xxxv (2004),
 * pp.327-336
 *
 * Table from AA for 1620 through today
 * Note, Stephenson and Morrison's table starts at the year 1630.
 * The Chapronts' table does not agree with the Almanac prior to 1630.
 * The actual accuracy decreases rapidly prior to 1780.
 *
 * Jean Meeus, Astronomical Algorithms, 2nd edition, 1998.
 * 
 * For a comprehensive collection of publications and formulae, see:
 * http://www.phys.uu.nl/~vgent/astro/deltatime.htm
 * 
 * For future values of delta t, the following data from the 
 * Earth Orientation Department of the US Naval Observatory can be used:
 * (TAI-UTC) from: ftp://maia.usno.navy.mil/ser7/tai-utc.dat
 * (UT1-UTC) from: ftp://maia.usno.navy.mil/ser7/finals.all
 * file description in: ftp://maia.usno.navy.mil/ser7/readme.finals
 * Delta T = TAI-UT1 + 32.184 sec = (TAI-UTC) - (UT1-UTC) + 32.184 sec
 *
 * Also, there is the following file:
 * http://maia.usno.navy.mil/ser7/deltat.data, but it is about 3 months
 * behind (on 3 feb 2009)
 *
 * Last update of table dt[]: Dieter Koch, 3 feb 2009.
 * ATTENTION: Whenever updating this table, do not forget to adjust
 * the macros TABEND and TABSIZ !
 */

const(
  TABSTART = 1620 
  TABEND = 2014
  TABSIZ = TABEND - TABSTART + 1

  /* we make the table greater for additional values read from external file */
  TABSIZ_SPACE = TABSIZ + 100
)

[]dt := []int {
/* 1620.0 thru 1659.0 */
12400, 11900, 11500, 11000, 10600, 10200, 9800, 9500, 9100, 8800,
8500, 8200, 7900, 7700, 7400, 7200, 7000, 6700, 6500, 6300,
6200, 6000, 5800, 5700, 5500, 5400, 5300, 5100, 5000, 4900,
4800, 4700, 4600, 4500, 4400, 4300, 4200, 4100, 4000, 3800,
/* 1660.0 thru 1699.0 */
3700, 3600, 3500, 3400, 3300, 3200, 3100, 3000, 2800, 2700,
2600, 2500, 2400, 2300, 2200, 2100, 2000, 1900, 1800, 1700,
1600, 1500, 1400, 1400, 1300, 1200, 1200, 1100, 1100, 1000,
1000, 1000, 900, 900, 900, 900, 900, 900, 900, 900,
/* 1700.0 thru 1739.0 */
900, 900, 900, 900, 900, 900, 900, 900, 1000, 1000,
1000, 1000, 1000, 1000, 1000, 1000, 1000, 1100, 1100, 1100,
1100, 1100, 1100, 1100, 1100, 1100, 1100, 1100, 1100, 1100,
1100, 1100, 1100, 1100, 1200, 1200, 1200, 1200, 1200, 1200,
/* 1740.0 thru 1779.0 */
1200, 1200, 1200, 1200, 1300, 1300, 1300, 1300, 1300, 1300,
1300, 1400, 1400, 1400, 1400, 1400, 1400, 1400, 1500, 1500,
1500, 1500, 1500, 1500, 1500, 1600, 1600, 1600, 1600, 1600,
1600, 1600, 1600, 1600, 1600, 1700, 1700, 1700, 1700, 1700,
/* 1780.0 thru 1799.0 */
1700, 1700, 1700, 1700, 1700, 1700, 1700, 1700, 1700, 1700,
1700, 1700, 1600, 1600, 1600, 1600, 1500, 1500, 1400, 1400,
/* 1800.0 thru 1819.0 */
1370, 1340, 1310, 1290, 1270, 1260, 1250, 1250, 1250, 1250,
1250, 1250, 1250, 1250, 1250, 1250, 1250, 1240, 1230, 1220,
/* 1820.0 thru 1859.0 */
1200, 1170, 1140, 1110, 1060, 1020, 960, 910, 860, 800,
750, 700, 660, 630, 600, 580, 570, 560, 560, 560,
570, 580, 590, 610, 620, 630, 650, 660, 680, 690,
710, 720, 730, 740, 750, 760, 770, 770, 780, 780,
/* 1860.0 thru 1899.0 */
788, 782, 754, 697, 640, 602, 541, 410, 292, 182,
161, 10, -102, -128, -269, -324, -364, -454, -471, -511,
-540, -542, -520, -546, -546, -579, -563, -564, -580, -566,
-587, -601, -619, -664, -644, -647, -609, -576, -466, -374,
/* 1900.0 thru 1939.0 */
-272, -154, -2, 124, 264, 386, 537, 614, 775, 913,
1046, 1153, 1336, 1465, 1601, 1720, 1824, 1906, 2025, 2095,
2116, 2225, 2241, 2303, 2349, 2362, 2386, 2449, 2434, 2408,
2402, 2400, 2387, 2395, 2386, 2393, 2373, 2392, 2396, 2402,
/* 1940.0 thru 1979.0 */
 2433, 2483, 2530, 2570, 2624, 2677, 2728, 2778, 2825, 2871,
 2915, 2957, 2997, 3036, 3072, 3107, 3135, 3168, 3218, 3268,
 3315, 3359, 3400, 3447, 3503, 3573, 3654, 3743, 3829, 3920,
 4018, 4117, 4223, 4337, 4449, 4548, 4646, 4752, 4853, 4959,
/* 1980.0 thru 1999.0 */
 5054, 5138, 5217, 5296, 5379, 5434, 5487, 5532, 5582, 5630,
 5686, 5757, 5831, 5912, 5998, 6078, 6163, 6230, 6297, 6347,
/* 2000.0 thru 2009.0 */
 6383, 6409, 6430, 6447, 6457, 6469, 6485, 6515, 6546, 6578,
/* Extrapolated values, 2010 - 2014 */
 6607, 6660, 6700, 6750, 6800,
// JAVA ONLY: add 100 empty elements, see constant TABSIZ_SPACE above!
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
}

const(
  TAB2_SIZ    = 27
  TAB2_START  = -1000
  TAB2_END    = 1600
  TAB2_STEP   = 100
  LTERM_EQUATION_YSTART = 1820
  LTERM_EQUATION_COEFF = 32
)

/* Table for -1000 through 1600, from Morrison & Stephenson (2004).  */
[]dt2 := []int {
/*-1000  -900  -800  -700  -600  -500  -400  -300  -200  -100*/
  25400,23700,22000,21000,19040,17190,15530,14080,12790,11640,
/*    0   100   200   300   400   500   600   700   800   900*/
  10580, 9600, 8640, 7680, 6700, 5710, 4740, 3810, 2960, 2200,
/* 1000  1100  1200  1300  1400  1500  1600,                 */
   1570, 1090,  740,  490,  320,  200,  120,
};

func calc_deltaT(tjd float, tid_acc float) float{
  ans := 0.0
  ans2 := 0.0
  ans3 := 0.0
  p := 0.0
  B := 0.0
  B2 := 0.0
  Y := 0.0
  Ygreg := 0.0
  dd := 0.0 // To remove Java warning of "maybe" not initialized
  var []d [6]int
  var i int
  var iy  int
  var k int
  /* read additional values from swedelta.txt */
  tabsiz := init_dt()
  tabend := TABSTART + tabsiz - 1
  Y = 2000.0 + (tjd - SwephData.J2000)/365.25;
  Ygreg = 2000.0 + (tjd - SwephData.J2000)/365.2425;
  /* before -500:
   * formula by Stephenson (1997; p. 508) but adjusted to fit the starting
   * point of table dt2 (Stephenson 1997). */
  if( Y < TAB2_START ) {
    B = (Y - LTERM_EQUATION_YSTART) * 0.01;
    ans = -20 + LTERM_EQUATION_COEFF * B * B;
    ans = adjust_for_tidacc(tid_acc, ans, Y);
    /* transition from formula to table over 100 years */
    if (Y >= TAB2_START - 100) {
      /* starting value of table dt2: */
      ans2 = adjust_for_tidacc(tid_acc, dt2[0], TAB2_START);
      /* value of formula at epoch TAB2_START */
      B = (TAB2_START - LTERM_EQUATION_YSTART) * 0.01;
      ans3 = -20 + LTERM_EQUATION_COEFF * B * B;
      ans3 = adjust_for_tidacc(tid_acc, ans3, Y);
      dd = ans3 - ans2;
      B = (Y - (TAB2_START - 100)) * 0.01;
      /* fit to starting point of table dt2. */
      ans = ans - dd * B;
    }
  }
  /* between -500 and 1600:
   * linear interpolation between values of table dt2 (Stephenson 1997) */
  if (Y >= TAB2_START && Y < TAB2_END) {
    double Yjul = 2000 + (tjd - 2451557.5) / 365.25;
    p = SMath.floor(Yjul);
    iy = (int) ((p - TAB2_START) / TAB2_STEP);
    dd = (Yjul - (TAB2_START + TAB2_STEP * iy)) / TAB2_STEP;
    ans = dt2[iy] + (dt2[iy+1] - dt2[iy]) * dd;
    /* correction for tidal acceleration used by our ephemeris */
    ans = adjust_for_tidacc(tid_acc, ans, Y);
  }
  /* between 1600 and 1620:
   * linear interpolation between 
   * end of table dt2 and start of table dt */
  if (Y >= TAB2_END && Y < TABSTART) { 
    B = TABSTART - TAB2_END;
    iy = (TAB2_END - TAB2_START) / TAB2_STEP;
    dd = (Y - TAB2_END) / B;
    ans = dt2[iy] + dd * (dt[0] / 100.0 - dt2[iy]);
    ans = adjust_for_tidacc(tid_acc, ans, Y);
  }
  /* 1620 - today + a few years (tabend):
   * Besselian interpolation from tabulated values in table dt.
   * See AA page K11.
   */
  if (Y >= TABSTART && Y <= tabend) {
    /* Index into the table.
     */
    p = SMath.floor(Y);
    iy = (int) (p - TABSTART);
    /* Zeroth order estimate is value at start of year
     */
    ans = dt[iy];
    k = iy + 1;
    if( k >= tabsiz )
      return deltatIsDone(ans, Y, B, tid_acc, tabsiz, tabend); /* No data, can't go on. */
    /* The fraction of tabulation interval
     */
    p = Y - p;
    /* First order interpolated value
     */
    ans += p*(dt[k] - dt[iy]);
    if( (iy-1 < 0) || (iy+2 >= tabsiz) )
      return deltatIsDone(ans, Y, B, tid_acc, tabsiz, tabend); /* can't do second differences */
    /* Make table of first differences
     */
    k = iy - 2;
    for( i=0; i<5; i++ ) {
      if( (k < 0) || (k+1 >= tabsiz) ) 
        d[i] = 0;
      else
        d[i] = dt[k+1] - dt[k];
      k += 1;
    }
    /* Compute second differences
     */
    for( i=0; i<4; i++ )
      d[i] = d[i+1] - d[i];
    B = 0.25*p*(p-1.0);
    ans += B*(d[1] + d[2]);
//    printf( "B %.4lf, ans %.4lf\n", B, ans );
    if( iy+2 >= tabsiz )
      return deltatIsDone(ans, Y, B, tid_acc, tabsiz, tabend);
    /* Compute third differences
     */
    for( i=0; i<3; i++ )
      d[i] = d[i+1] - d[i];
    B = 2.0*B/3.0;
    ans += (p-0.5)*B*d[1];
//    printf( "B %.4lf, ans %.4lf\n", B*(p-0.5), ans );
    if( (iy-2 < 0) || (iy+3 > tabsiz) )
      return deltatIsDone(ans, Y, B, tid_acc, tabsiz, tabend);
    /* Compute fourth differences
     */
    for( i=0; i<2; i++ )
      d[i] = d[i+1] - d[i];
    B = 0.125*B*(p+1.0)*(p-2.0);
    ans += B*(d[0] + d[1]);
//    printf( "B %.4lf, ans %.4lf\n", B, ans );
  }

  return deltatIsDone(ans, Y, B, tid_acc, tabsiz, tabend);
}

func deltatIsDone(double ans, double Y, double B, double tid_acc, int tabsiz, int tabend) float {
// //#ifdef TRACE0
//     // Trace.level++; Don't increment here, as the calling method calc_deltat() does not decrement on return!
//     Trace.log("SweDate.deltatIsDone(double, double, double, double, int, int)");
// //#endif /* TRACE0 */
  double ans2, ans3, B2, dd;
  if (Y >= TABSTART && Y <= tabend) {
    ans *= 0.01;
    ans = adjust_for_tidacc(tid_acc, ans, Y);
  }
  /* today - :
   * Formula Stephenson (1997; p. 507),
   * with modification to avoid jump at end of AA table,
   * similar to what Meeus 1998 had suggested.
   * Slow transition within 100 years.
   */
  if (Y > tabend) {
    B = 0.01 * (Y - 1820);
    ans = -20 + 31 * B * B;
    /* slow transition from tabulated values to Stephenson formula: */
    if (Y <= tabend+100) {
      B2 = 0.01 * (tabend - 1820);
      ans2 = -20 + 31 * B2 * B2;
      ans3 = dt[tabsiz-1] * 0.01;
      dd = (ans2 - ans3);
      ans += dd * (Y - (tabend + 100)) * 0.01;
    }
  }
  return ans / 86400.0;
}


/* Read delta t values from external file.
 * record structure: year(whitespace)delta_t in 0.01 sec.
 */
func init_dt() int {
  FilePtr fp := nil
  var year int
  var tab_index int
  var tabsiz int
  var i int
  var s string
  if (!init_dt_done) {
    init_dt_done = true
    /* no error message if file is missing */

    if ((fp = sw.swi_fopen(-1, "swe_deltat.txt", sw.swed.ephepath, null)) == null &&
        (fp = sw.swi_fopen(-1, "sedeltat.txt", sw.swed.ephepath, null)) == null) {
      return TABSIZ;  // I think, I could skip this one...
    }

    while ((s=fp.readLine()) != null) {
      s.trim();
      if (s.length() == 0 || s.charAt(0) == '#') {
        continue;
      }
      year = SwissLib.atoi(s);
      tab_index = year - TABSTART;
      /* table space is limited. no error msg, if exceeded */
      if (tab_index >= TABSIZ_SPACE)
        continue;
      if (s.length() > 4) {
        s = s.substring(4).trim();
      }
      dt[tab_index] = (short)(SwissLib.atof(s) * 100 + 0.5);
    }

    fp.close()
  }
  /* find table size */
  tabsiz = 2001 - TABSTART + 1;
  for (i = tabsiz - 1; i < TABSIZ_SPACE; i++) {
    if (dt[i] == 0)
      break;
    else
      tabsiz++;
  }
  tabsiz--
  return tabsiz
}
