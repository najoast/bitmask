package bitmask

import (
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Errorf("assert failed")
	}
}

func test(t *testing.T, bm BitMask, max, maxBitIndex int) {
	if bm == nil {
		assert(t, max == 0)
		return
	}
	for i := 0; i < maxBitIndex; i++ {
		bm.AddMask(i)
	}
	// assert(t, bm.(*BitMaskU64).data == (1<<maxBitIndex)-1)

	for i := 0; i < max; i++ {
		if i < maxBitIndex {
			assert(t, bm.HasMask(i) == true)
			bm.DelMask(i)
		} else {
			assert(t, bm.HasMask(i) == false)
		}
	}

	// assert(t, bm.(*BitMaskU64).data == 0)
	for i := 0; i < maxBitIndex; i++ {
		assert(t, bm.HasMask(i) == false)
	}
}

func testLimited(t *testing.T, max, maxBitIndex int) {
	bm := CreateLimitedBitMask(max)
	test(t, bm, max, maxBitIndex)
}

func testUnlimited(t *testing.T, max, maxBitIndex int) {
	bm := CreateUnlimitedBitMask()
	test(t, bm, max, maxBitIndex)
}

// -----------------------------------------------------

func TestU64(t *testing.T) {
	testLimited(t, 64, 64)
	testLimited(t, 50, 30)
	testLimited(t, 1, 1)
	testLimited(t, 0, 0)
}

func TestU64s(t *testing.T) {
	testLimited(t, 64000, 64000)
	testLimited(t, 64000, 128)
	testLimited(t, 1024, 512)
	testLimited(t, 129, 128)
	testLimited(t, 65, 65)
	testLimited(t, 65, 10)
}

func TestUnlimited(t *testing.T) {
	testUnlimited(t, 9999999, 9999999)
	testUnlimited(t, 9999999, 100)
	testUnlimited(t, 64000, 64000)
	testUnlimited(t, 64000, 128)
	testUnlimited(t, 1024, 512)
	testUnlimited(t, 129, 128)
	testUnlimited(t, 65, 65)
	testUnlimited(t, 65, 10)
	testUnlimited(t, 20, 10)
}
