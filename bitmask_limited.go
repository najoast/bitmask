package bitmask

// -------------------------------------------------

type BitMaskU64 struct {
	max  int
	data uint64
}

func (bm *BitMaskU64) AddMask(index int) bool {
	if index > bm.max {
		return false
	}
	bm.data = AddMask(bm.data, index)
	return true
}

func (bm *BitMaskU64) DelMask(index int) {
	if index > bm.max {
		return
	}
	bm.data = DelMask(bm.data, index)
}

func (bm *BitMaskU64) HasMask(index int) bool {
	if index > bm.max {
		return false
	}
	return HasMask(bm.data, index)
}

// -------------------------------------------------

type BitMaskU64s struct {
	max  int
	data []uint64
}

func (bm *BitMaskU64s) AddMask(index int) bool {
	if index > bm.max {
		return false
	}
	i, bitIndex := calcBitArrayIndex(index)
	bm.data[i] = AddMask(bm.data[i], bitIndex)
	return true
}

func (bm *BitMaskU64s) DelMask(index int) {
	if index > bm.max {
		return
	}
	i, bitIndex := calcBitArrayIndex(index)
	bm.data[i] = DelMask(bm.data[i], bitIndex)
}

func (bm *BitMaskU64s) HasMask(index int) bool {
	if index > bm.max {
		return false
	}
	i, bitIndex := calcBitArrayIndex(index)
	return HasMask(bm.data[i], bitIndex)
}

// -------------------------------------------------
