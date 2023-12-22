package base62

import (
	"strings"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const base = len(base62Chars)

// Encode takes a decimal number and encodes it to Base 62, returning EMPTY string if invalid
func Encode(decimal int) string {
	// although there will not be index 0 in this application, this must be valid encoding option
	if decimal == 0 {
		return string(base62Chars[0])
	}

	var encoded strings.Builder
	for decimal > 0 {
		// prepend the result of base62Chars[decimal % base] to the result
		encoded.WriteString(string(base62Chars[decimal%base]))
		// update the dividend to be quotient of the division
		decimal /= base
	}

	return reverse(encoded.String())
}

// Decode takes a Base62 string and decodes it to integer
func Decode(encoded string) int {
	var id int
	for _, c := range encoded {
		id = id*base + strings.IndexRune(base62Chars, c)
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
