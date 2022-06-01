package bitmask

type BitMask interface {
	AddMask(index int) bool
	DelMask(index int)
	HasMask(index int) bool
}

// If max exceeds 64,000, it is recommended to use unlimited BitMask instead
func CreateLimitedBitMask(max int) BitMask {
	if max <= 0 {
		return nil
	}
	if max <= bitsPerUnit {
		return &BitMaskU64{
			max:  max,
			data: 0,
		}
	} else {
		return &BitMaskU64s{
			max:  max,
			data: make([]uint64, calcBitArraySize(max)),
		}
	}
}

func CreateUnlimitedBitMask() BitMask {
	return &BitMaskMap{
		data: make(map[int]uint64),
	}
}
