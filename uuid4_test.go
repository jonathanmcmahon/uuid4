package uuid4

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewShouldNotYieldAnyCollisions(t *testing.T) {

	nRuns := 1000

	var uuids [][]byte

	// Generate a bunch of UUIDs
	for i := 0; i < nRuns; i++ {
		u, err := NewBytes()
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

func TestUUIDNewShouldFormatStringCorrectly(t *testing.T) {

	uuid, err := New()
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

func TestUUIDShouldContainCorrectVariantBits(t *testing.T) {

	variantIdx := IdxClkSeqHiRes

	for i := 0; i < 100; i++ {
		uuid, err := NewBytes()
		if err != nil {
			t.Error(err)
		}
		variant := uuid[variantIdx] >> 6
		if variant != 0x2 { // 0b10
			t.Errorf("Variant bits are incorrect: %v", variant)
		}
	}

}

func TestUUIDShouldContainCorrectVersionBits(t *testing.T) {

	versionIdx := IdxTimeHiAndVersion

	for i := 0; i < 100; i++ {
		uuid, err := NewBytes()
		if err != nil {
			t.Error(err)
		}
		version := uuid[versionIdx] >> 4
		if version != 0x4 { // 0b0100
			t.Errorf("Version bits are incorrect: %v", version)
		}
	}

}

func TestUUIDStringNonrandomBitsShouldYieldAppropriateChars(t *testing.T) {

	variantIdx := (IdxClkSeqHiRes * 2) + 3      // compensate for hyphens
	versionIdx := (IdxTimeHiAndVersion * 2) + 2 // compensate for hyphens

	for i := 0; i < 100; i++ {
		uuid, err := New()
		if err != nil {
			t.Error(err)
		}
		validVariantChars := "89ab"
		variant := string(uuid[variantIdx])
		if !strings.Contains(validVariantChars, variant) {
			t.Errorf("Variant character is invalid: %v", variant)
		}
		version := []rune(uuid)[versionIdx]
		if string(version) != "4" {
			t.Errorf("Version character should be 4 but was %v", version)
		}
	}
}

func ExampleNew_output() {
	u, err := New()
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	fmt.Println()

	b, err := NewBytes()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}
