package tempconv

import (
	"fmt"
	"testing"
)

func TestCToF(t *testing.T) {
	for _, test := range []struct {
		desc  string
		c     Celsius
		wantF Fahrenheit
	}{
		{
			desc:  "Zero Celsius",
			c:     0,
			wantF: 32,
		},
		{
			desc:  "Positive Celsius",
			c:     100,
			wantF: 212,
		},
		{
			desc:  "Negative Celsius",
			c:     -100,
			wantF: -148,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			f := CToF(test.c)
			if f != test.wantF {
				t.Fatalf("%s: got F %s, want %s", test.desc, f, test.wantF)
			}
		})
	}
}

func TestFToC(t *testing.T) {
	for _, test := range []struct {
		desc  string
		f     Fahrenheit
		wantC Celsius
	}{
		{
			desc:  "Zero Fahrenheit",
			f:     0,
			wantC: -17.77777777777778,
		},
		{
			desc:  "Positive Fahrenheit",
			f:     212,
			wantC: 100,
		},
		{
			desc:  "Negative Fahrenheit",
			f:     -148,
			wantC: -100,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			c := FToC(test.f)
			if c != test.wantC {
				t.Fatalf("%s: got C %s, want %s", test.desc, c, test.wantC)
			}
		})
	}
}

func TestCToK(t *testing.T) {
	for _, test := range []struct {
		desc  string
		c     Celsius
		wantK Kelvin
	}{
		{
			desc:  "Zero Celsius",
			c:     0,
			wantK: 273.15,
		},
		{
			desc:  "Positive Celsius",
			c:     100,
			wantK: 373.15,
		},
		{
			desc:  "Negative Celsius",
			c:     -273.15,
			wantK: 0.0,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			k := CToK(test.c)
			if k != test.wantK {
				t.Fatalf("%s: got K %s, want %s", test.desc, k, test.wantK)
			}
		})
	}
}

func TestKToC(t *testing.T) {
	for _, test := range []struct {
		desc  string
		k     Kelvin
		wantC Celsius
	}{
		{
			desc:  "Zero Kelvin",
			k:     0.0,
			wantC: -273.15,
		},

		{
			desc:  "Positive Kelvin",
			k:     273.15,
			wantC: 0,
		},
		{
			desc:  "Negative Kelvin",
			k:     -226.85,
			wantC: -500,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			c := KToC(test.k)
			if c != test.wantC || c.String() != fmt.Sprintf("%g*C", test.wantC) {
				t.Fatalf("%s: got C %s, want %s", test.desc, c, test.wantC)
			}
		})
	}
}

func TestKToF(t *testing.T) {
	for _, test := range []struct {
		desc  string
		k     Kelvin
		wantF Fahrenheit
	}{
		{
			desc:  "Zero Kelvin",
			k:     0.0,
			wantF: -459.66999999999996,
		},

		{
			desc:  "Positive Kelvin",
			k:     100.15,
			wantF: -279.4,
		},
		{
			desc:  "Negative Kelvin",
			k:     -273.15,
			wantF: -951.3399999999999,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			f := KToF(test.k)
			if f != test.wantF || f.String() != fmt.Sprintf("%g*F", test.wantF) {
				t.Fatalf("%s: got F %s, want %s", test.desc, f, test.wantF)
			}
		})
	}
}

func TestFToK(t *testing.T) {
	for _, test := range []struct {
		desc  string
		f     Fahrenheit
		wantK Kelvin
	}{
		{
			desc:  "Zero Fahrenheit",
			f:     0,
			wantK: 255.3722222222222,
		},
		{
			desc:  "Positive Fahrenheit",
			f:     255.372,
			wantK: 397.24555555555554,
		},
		{
			desc:  "Negative Fahrenheit",
			f:     -459.67,
			wantK: 0,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			k := FToK(test.f)
			if k != test.wantK || k.String() != fmt.Sprintf("%g*K", test.wantK) {
				t.Fatalf("%s: got K %s, want %s", test.desc, k, test.wantK)
			}
		})
	}
}
