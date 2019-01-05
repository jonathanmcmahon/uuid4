// Package uuid4 provides generation of version 4 Universally Unique IDentifiers (UUIDs), compliant with RFC 4122 (https://tools.ietf.org/html/rfc4122).
package uuid4

import (
	"crypto/rand"
	"fmt"
	"strings"
)

// NewString returns a new version 4 UUID in string representation.
//
// Example: "48c51d77-69be-a444-c12f-984f2db2c269"
//
// source: https://tools.ietf.org/html/rfc4122#section-3
func NewString() (string, error) {

	uuid, err := New()
	if err != nil {
		return "", err
	}

	var dst strings.Builder
	for i, v := range uuid {
		// Insert hyphens at appropriate indices
		if i == 4 || i == 6 || i == 8 || i == 10 {
			dst.WriteString(sep)
		}
		// Encode MSB
		dst.WriteString(string(hex[v>>4]))
		// Encode LSB
		dst.WriteString(string(hex[v&0xF]))
	}

	return dst.String(), nil
}

// New returns a new version 4 UUID as a byte slice.
//
// source: https://tools.ietf.org/html/rfc4122#section-4.4
func New() ([]byte, error) {

	uuid := make([]byte, uuidBytes)

	// First we set all the bits to random values.
	_, err := rand.Read(uuid)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	// Then we apply two bitwise logical ops that, when applied in the correct
	// sequence, will as a whole accomplish the following:

	// a) set the two most significant bits (bits 6 and 7) of the
	// clock_seq_hi_and_reserved to zero and one, respectively;
	uuid[IdxClkSeqHiRes] &= 0xbf
	uuid[IdxClkSeqHiRes] |= 0x80

	// b) set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number from
	// Section 4.1.3 (the 4-bit version number for v4 is: 0100)
	uuid[IdxTimeHiAndVersion] &= 0x0f
	uuid[IdxTimeHiAndVersion] |= 0x40

	return uuid, nil
}

// Byte indices for the location of fields within the uuid, and their byte length.
//
//    0                   1                   2                   3
//    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                          time_low                             |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |       time_mid                |         time_hi_and_version   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |clk_seq_hi_res |  clk_seq_low  |         node (0-1)            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                         node (2-5)                            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// source: https://tools.ietf.org/html/rfc4122#section-4.1.2
const (
	IdxTimeLow   = 0
	BytesTimeLow = 4

	IdxTimeMid   = 4
	BytesTimeMid = 2

	IdxTimeHiAndVersion   = 6
	BytesTimeHiAndVersion = 2

	IdxClkSeqHiRes   = 8
	BytesClkSeqHiRes = 1

	IdxClkSeqLow   = 9
	BytesClkSeqLow = 1

	IdxNode   = 10
	BytesNode = 6
)
const uuidBytes = 16 // uuid is 128 bits
const hex = "0123456789abcdef"
const sep = "-"
