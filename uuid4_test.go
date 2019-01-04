package uuid4

import (
	"strings"
	"testing"
)

func TestNewShouldNotYieldAnyCollisions(t *testing.T) {

	nRuns := 1000

	var uuids [][]byte

	// Generate a bunch of UUIDs
	for i := 0; i < nRuns; i++ {
		u, err := New()
		if err != nil {
			t.Error(err)
			return
		}
		uuids = append(uuids, u)
	}

	// Compare each UUID to all the others, ensuring no two are identical
	for i, v := range uuids {
		// We have already compared the first through i-th elements to all
		// other elements, so we can omit those elements from all future iterations
		for _, w := range uuids[i+1:] {
			match := true
			// Compare the bytes in each UUID
			for bIdx, b := range v {
				if b != w[bIdx] {
					match = false
				}
			}
			if match == true {
				t.Errorf("Two matching UUIDs: %v, %v", v, w)
			}
		}
	}
}

func TestUUIDNewStringShouldFormatStringCorrectly(t *testing.T) {

	uuid, err := NewString()
	if err != nil {
		t.Error(err)
	}

	expectedLen := (8 * 4) + 4
	if len(uuid) != expectedLen {
		t.Errorf("UUID length %v did not match expected value %v: %v", len(uuid), expectedLen, uuid)
	}

	for i, char := range uuid {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			if char != rune('-') {
				t.Errorf("Missing hypen (-) at index %v: %v", i, uuid)
			}
		} else {
			if !strings.ContainsRune(hex, char) {
				t.Errorf("Non-hex character %v in uuid at index %v: %v", char, i, uuid)
			}
		}
	}
}
