package converter

import "testing"

func TestCToK(t *testing.T) {
	var c Celsius = 5
	k := CToK(c)
	if k != Kelvin(c)+KelvinConst {
		t.Errorf("Expected 278.15, got %s", k)
	}
}

func TestMToF(t *testing.T) {
	var m Meter = 1
	f := MToF(m)
	if f != Feet(3.280839895013123) {
		t.Errorf("Expected 3.280839895013123, got %s", f)
	}
}

func TestKgToP(t *testing.T) {
	var kg Kilogram = 1
	p := KgToP(kg)
	if p != Pound(2.20462262185) {
		t.Errorf("Expected 2.20462262185, got %s", p)
	}
}
