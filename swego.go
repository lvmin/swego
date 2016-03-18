// Package swego defines an interface for interfacing with the Swiss Ephemeris.
package swego

// CalcFlags represents the flags argument of swe_calc and swe_calc_ut in a
// stateless way.
type CalcFlags struct {
	Flags   int32
	TopoLoc TopoLoc
	SidMode SidMode

	// FileNameJPL represents the argument to swe_set_jpl_file.
	FileNameJPL string
}

// TopoLoc represents the arguments to swe_set_topo.
type TopoLoc struct {
	Lat  float64
	Long float64
	Alt  float64
}

// AyanamsaExFlags represents the flags argument of swe_get_ayanamsa_ex and
// swe_get_ayanamsa_ex_ut in a stateless way.
type AyanamsaExFlags struct {
	Flags   int32
	SidMode SidMode
}

// HousesExFlags represents the flags argument of swe_houses_ex in a stateless
// way.
type HousesExFlags struct {
	Flags   int32
	SidMode SidMode
}

// SidMode represents the arguments of swe_set_sid_mode.
type SidMode struct {
	Mode   int32
	T0     float64
	AyanT0 float64
}

// CalType represents the calendar type used in julian date conversion.
type CalType int

// Calendar types.
const (
	Julian    CalType = 0
	Gregorian CalType = 1
)

// Interface defines a standardized way for interfacing with the Swiss
// Ephemeris library from Go.
type Interface interface {
	// Version returns the version of the Swiss Ephemeris.
	Version() string

	// SetPath sets the ephemeris data path.
	SetPath(ephepath string)
	// Close closes the Swiss Ephemeris library.
	Close()

	// Calc calculates the position and optionally the speed of planet pl at
	// Julian Date (in Ephemeris Time) et with calculation flags fl.
	Calc(et float64, pl int, fl CalcFlags) (xx [6]float64, cfl int, err error)
	// CalcUT calculates the position and optionally the speed of planet pl at
	// Julian Date (in Universal Time) ut with calculation flags fl. Within the C
	// library swe_deltat is called to convert Universal Time to Ephemeris Time.
	CalcUT(ut float64, pl int, fl CalcFlags) (xx [6]float64, cfl int, err error)

	// PlanetName returns the name of planet pl.
	PlanetName(pl int) string

	// GetAyanamsa returns the ayanamsa for Julian Date (in Ephemeris Time) et.
	GetAyanamsa(et float64) float64
	// GetAyanamsaUT returns the ayanamsa for Julian Date (in Universal Time) ut.
	GetAyanamsaUT(ut float64) float64
	// GetAyanamsaEx returns the ayanamsa for Julian Date (in Ephemeris Time) et.
	// It is equal to GetAyanamsa but uses the ΔT consistent with the ephemeris
	// passed in fl.Flags.
	GetAyanamsaEx(et float64, fl AyanamsaExFlags) (float64, error)
	// GetAyanamsaExUT returns the ayanamsa for Julian Date (in Universal Time) ut.
	// It is equal to GetAyanamsaUT but uses the ΔT consistent with the ephemeris
	// passed in fl.Flags.
	GetAyanamsaExUT(ut float64, fl AyanamsaExFlags) (float64, error)
	// GetAyanamsaName returns the name of sidmode.
	GetAyanamsaName(sidmode int32) string

	// JulDay returns the corresponding Julian Date for the given date. Calendar
	// type ct is used to clearify the year y, Julian or Gregorian.
	JulDay(y, m, d int, h float64, ct CalType) float64
	// RevJul returns the corresponding calendar date for the given Julian Date.
	// Calendar type ct is used to clearify the year y, Julian or Gregorian.
	RevJul(jd float64, ct CalType) (y, m, d int, h float64)
	// UTCToJD returns the corresponding Julian Date in Ephemeris and Universal
	// Time for the given date and accounts for leap seconds in the conversion.
	// Calendar type ct is used to clearify the year y, Julian or Gregorian.
	UTCToJD(y, m, d int, h float64, ct CalType) (et, ut float64, err error)
	// JdETToUTC returns the corresponding calendar date for the given Julian
	// Date in Ephemeris Time and accounts for leap seconds in the conversion.
	// Calendar type ct is used to clearify the year y, Julian or Gregorian.
	JdETToUTC(et float64, ct CalType) (y, m, d, h, i int, s float64)
	// JdETToUTC returns the corresponding calendar date for the given Julian
	// Date in Universal Time and accounts for leap seconds in the conversion.
	// Calendar type ct is used to clearify the year y, Julian or Gregorian.
	JdUT1ToUTC(ut1 float64, ct CalType) (y, m, d, h, i int, s float64)

	Houses(ut, geolat, geolon float64, hsys int) ([]float64, [10]float64)
	HousesEx(ut float64, fl HousesExFlags, geolat, geolon float64, hsys int) ([]float64, [10]float64)
	HousesArmc(armc, geolat, eps float64, hsys int) ([]float64, [10]float64)
	HousePos(armc, geolat, eps float64, hsys int, xpin [2]float64) (float64, error)
	HouseName(hsys int) string

	// DeltaT returns the ΔT for the Julian Date jd.
	DeltaT(jd float64) float64
	// DeltaTEx returns the ΔT for the Julian Date jd.
	DeltaTEx(jd float64, fl int32) (float64, error)

	TimeEqu(jd float64) (float64, error)
	LMTToLAT(jdLMT, geolon float64) (float64, error)
	LATToLMT(jdLAT, geolon float64) (float64, error)

	SidTime0(ut, eps, nut float64) float64
	SidTime(ut float64) float64
}