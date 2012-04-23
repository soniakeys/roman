// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package roman

import (
	"errors"
	"fmt"
)

var m = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// error message prefix
const rp = "roman.Parse: "

// unicode overbar combining character
const occ = 0x305

func Parse(rom string) (n int, err error) {
	if rom == "" {
		return 0, errors.New(rp + "Empty string")
	}
	is := []rune(rom) // easier to convert string up front
	var c0 rune       // c0: roman character last read
	var cv0 int       // cv0: value of cv

	// the key to the algorithm is to process digits from right to left
	for i := len(is) - 1; i >= 0; i-- {
		// read roman digit
		c := is[i]
		k := c == occ
		if k {
			if i == 0 {
				return 0, errors.New(rp +
					"Overbar combining character invalid at position 0")
			}
			i--
			c = is[i]
		}
		cv := m[c]
		if cv == 0 {
			if c == occ {
				return 0, fmt.Errorf(rp+
					"Overbar combining character invalid at position %d", i)
			} else {
				return 0, fmt.Errorf(rp+
					"Character unrecognized as Roman digit: %c", c)
			}
		}
		if k {
			c = -c // convention indicating overbar
			cv *= 1000
		}

		// handle cases of new, same, subtractive, changed, in that order.
		switch {
		default: // case 4: digit change
			fallthrough
		case c0 == 0: // no previous digit
			c0 = c
			cv0 = cv
		case c == c0: // same digit
		case cv*5 == cv0 || cv*10 == cv0: // subtractive
			// handle next digit as new.
			// a subtractive digit doesn't count as a previous digit.
			c0 = 0
			n -= cv  // subtract...
			continue // ...instead of adding
		}
		n += cv // add, in all cases except subtractive
	}
	return
}
