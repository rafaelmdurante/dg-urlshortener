package base62

import (
	"fmt"
	"testing"
)

var tests = []struct {
	decimal int
	base62  string
}{
	{
		decimal: 0,
		base62:  "0",
	},
	{
		decimal: 1, // first possible id
		base62:  "1",
	},
	{
		decimal: 10,
		base62:  "A",
	},
	{
		decimal: 61,
		base62:  "z",
	},
	{
		decimal: 62,
		base62:  "10",
	},
	{
		decimal: 273,
		base62:  "4P",
	},
	{
		decimal: 11_157,
		base62:  "2tx",
	},
	{
		decimal: 18_840_318,
		base62:  "1H3E6",
	},
	{
		decimal: 387_591_234,
		base62:  "QEIBW",
	},
	{
		decimal: 56_800_235_583, // last possible ID
		base62:  "zzzzzz",
	},
}

func TestEncode(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("encode %d", tt.decimal), func(t *testing.T) {
			if got := StdEncoding.Encode(tt.decimal); got != tt.base62 {
				t.Errorf("base62.Encode(%d) = '%s', want '%s'",
					tt.decimal, got, tt.base62)
			}
		})
	}

	// negative amount should return empty string
	n := -1
	t.Run(fmt.Sprintf("encode negative"), func(t *testing.T) {
		if got := StdEncoding.Encode(n); got != "" {
			t.Errorf("base62.Encode(%d) = '%s', want '%s'",
				n, got, "")
		}
	})
}

func TestDecode(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("decode %s", tt.base62), func(t *testing.T) {
			if got := StdEncoding.Decode(tt.base62); got != tt.decimal {
				t.Errorf("base62.Decode(%s) = '%d', want '%d'",
					tt.base62, got, tt.decimal)
			}
		})
	}

	// empty string should return 0
	s := ""
	t.Run(fmt.Sprintf("decode empty string"), func(t *testing.T) {
		if got := StdEncoding.Decode(s); got != 0 {
			t.Errorf("base62.Decode(%s) = '%d', want '%d'",
				s, got, 0)
		}
	})
}

func TestEncode64(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("encode %d", tt.decimal), func(t *testing.T) {
			if got := Base64Encoding.Encode(tt.decimal); got != tt.base62 {
				t.Errorf("base64.Encode(%d) = '%s', want '%s'",
					tt.decimal, got, tt.base62)
			}
		})
	}

	// negative amount should return empty string
	n := -1
	t.Run(fmt.Sprintf("encode negative"), func(t *testing.T) {
		if got := Base64Encoding.Encode(n); got != "" {
			t.Errorf("base62.Encode(%d) = '%s', want '%s'",
				n, got, "")
		}
	})
}
