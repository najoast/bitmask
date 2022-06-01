package bitmask

type BitMaskMap struct {
	data map[int]uint64
}

func (bm *BitMaskMap) AddMask(index int) bool {
	i, bitIndex := calcBitArrayIndex(index)
	bm.data[i] = AddMask(bm.data[i], bitIndex)
	return true
}

func (bm *BitMaskMap) DelMask(index int) {
	i, bitIndex := calcBitArrayIndex(index)
	bm.data[i] = DelMask(bm.data[i], bitIndex)
}

func (bm *BitMaskMap) HasMask(index int) bool {
	i, bitIndex := calcBitArrayIndex(index)
	return HasMask(bm.data[i], bitIndex)
}
