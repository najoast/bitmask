package bitmask

func AddMask(data uint64, index int) uint64 {
	return data | (1 << index)
}

func DelMask(data uint64, index int) uint64 {
	return data & ^(1 << index)
}

func HasMask(data uint64, index int) bool {
	return (data & (1 << index)) != 0
}

func calcBitArrayIndex(index int) (arrIndex, bitIndex int) {
	return index / bitsPerUnit, index % bitsPerUnit
}

func calcBitArraySize(max int) int {
	return (max / bitsPerUnit) + 1
}
