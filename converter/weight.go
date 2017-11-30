package converter

import "fmt"

type Kilogram float64
type Pound float64

const PoundConst Pound = 2.20462262185

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}
