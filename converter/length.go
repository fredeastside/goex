package converter

import "fmt"

// Meter is length in meters
type Meter float64

// Feet is length in feets
type Feet float64

const FeetConst Feet = 0.3048

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}
