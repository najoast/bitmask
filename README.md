# bitmask
A bitmask golang implementation.

# Usage
```go
// Limited bitmask, more efficiently but has upper limit.
// If max exceeds 64,000, it is recommended to use unlimited BitMask instead.
bm := bitmask.CreateLimitedBitMask(64000)
bm.AddMask(1) // bm.HasMask(1) == true
bm.DelMask(1) // bm.HasMask(1) == false

// Unlimit bitmask
bm := bitmask.CreateUnlimitedBitMask()
for i:=0; i<9999999999; i++ {
    bm.AddMask(i)
}
```