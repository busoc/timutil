package timutil

import (
  "time"
  "fmt"
)

func ExampleGPSTime() {
  t := time.Date(2019, 2, 21, 8, 36, 12, 0, time.UTC)
  g := GPSTime(t)
  fmt.Println(g.Format("2006-01-02 15:04:05"))
  // Output: 2019-02-21 08:36:30
}
