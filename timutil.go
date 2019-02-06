package timutil

import (
  "time"
)

var dates = []time.Time {
  time.Date(1981, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(1982, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(1983, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(1985, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(1987, 12, 31, 0, 0, 0, 0, time.UTC),
  time.Date(1989, 12, 31, 0, 0, 0, 0, time.UTC),
  time.Date(1990, 12, 31, 0, 0, 0, 0, time.UTC),
  time.Date(1992, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(1993, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(1994, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(1995, 12, 31, 0, 0, 0, 0, time.UTC),
  time.Date(1997, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(1998, 12, 31, 0, 0, 0, 0, time.UTC),
  time.Date(2005, 12, 31, 0, 0, 0, 0, time.UTC),
  time.Date(2008, 12, 31, 0, 0, 0, 0, time.UTC),
  time.Date(2012, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(2015, 6, 30, 0, 0, 0, 0, time.UTC),
  time.Date(2016, 12, 31, 0, 0, 0, 0, time.UTC),
}

const FormatDate = "2006-01-02T15:04:05.000"

var (
  UNIX = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
  GPS  = time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
  delta = GPS.Sub(UNIX)
)

func Split5(t time.Time) (uint32, uint8) {
	s, n := float64(t.Unix()), float64(t.UnixNano())/1000000.0
	m := (n - (s * 1000)) / 1000 * 255

	return uint32(s), uint8(m)
}

func Split6(t time.Time) (uint32, uint8) {
  return 0, 0
}

func Join5(coarse uint32, fine uint8) time.Time {
	t := time.Unix(int64(coarse), 0).UTC()

	fs := float64(fine) / 256.0 * 1000.0
	ms := time.Duration(fs) * time.Millisecond

	return utcTime(t.Add(ms))
}

func Join6(coarse uint32, fine uint8) time.Time {
	t := time.Unix(int64(coarse), 0).UTC()

	fs := float64(fine) / 65536.0 * 1000.0
	ms := time.Duration(fs) * time.Millisecond
	return utcTime(t.Add(ms))
}

func gpsTime(t time.Time) time.Time {
  return t.UTC().Add(-delta-leap(t))
}

func utcTime(t time.Time) time.Time {
  return t.Add(delta+leap(t)).UTC()
}

func leap(t time.Time) time.Duration {
  var d time.Duration
  for i := 0; i < len(dates) && t.Before(dates[i]); i++ {
    if t.After(dates[i]) {
      d += time.Second
    }
  }
  return d
}
