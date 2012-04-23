package roman

import "testing"

var tcs = []struct {
	n   int
	rom string
}{
	{1990, "MCMXC"},
	{2008, "MMVIII"},
	{1666, "MDCLXVI"},
	{3888888, "M̅M̅M̅D̅C̅C̅C̅L̅X̅X̅X̅V̅I̅I̅I̅DCCCLXXXVIII"},
}

func TestParse(t *testing.T) {
	for _, tc := range tcs {
		switch n, err := Parse(tc.rom); {
		case err != nil:
			t.Errorf("parse %s: %v", tc.rom, err)
		case n != tc.n:
			t.Errorf("parse %s: expected %d, got %d.", tc.rom, tc.n, n)
		default:
			t.Log(tc.rom, "==", n)
		}
	}
}

func TestFormat(t *testing.T) {
	for _, tc := range tcs {
		switch rom, ok := Format(tc.n); {
		case !ok:
			t.Errorf(`format %d: expected %s, got "not representable"`,
				tc.n, tc.rom)
		case rom != tc.rom:
			t.Errorf("format %d: expected %s, got %s", tc.n, tc.rom, rom)
		default:
			t.Log(tc.n, "==", rom)
		}
	}
}
