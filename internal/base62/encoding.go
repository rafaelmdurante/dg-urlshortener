package base62

import (
	"strings"
)

type Encoding struct {
	dictionary string
	base       int
}

const (
	base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// NewEncoding returns a new Encoding defined by the given dictionary
func newEncoding(encodingMap string) *Encoding {
	e := new(Encoding)

	e.dictionary = encodingMap
	e.base = len(encodingMap)

	return e
}

var StdEncoding = newEncoding(base62)

func (e *Encoding) Encode(decimal int) string {
	// although there will not be index 0 in this application, this must be valid encoding option
	if decimal == 0 {
		return string(e.dictionary[0])
	}

	var encoded strings.Builder
	for decimal > 0 {
		// prepend the result of base62[decimal % base] to the result
		encoded.WriteString(string(e.dictionary[decimal%e.base]))
		// update the dividend to be quotient of the division
		decimal /= e.base
	}

	return reverse(encoded.String())
}

// Decode takes a Base62 string and decodes it to integer
func (e *Encoding) Decode(encoded string) int {
	var id int
	for _, c := range encoded {
		id = id*e.base + strings.IndexRune(e.dictionary, c)
	}

	return id
}

// reverse takes a string and returns it reversed
func reverse(s string) string {
	r := []rune(s)
	l := len(r)

	for start, end := 0, l-1; start < l/2; start, end = start+1, end-1 {
		r[start], r[end] = r[end], r[start]
	}

	return string(r)
}
