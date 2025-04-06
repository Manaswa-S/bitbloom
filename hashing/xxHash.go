package hashing

import (
	"encoding/binary"
	"math/bits"
)
 
var (
	xxHashPrime1 uint64 = 11400714785074694791
	xxHashPrime2 uint64 = 14029467366897019727
	xxHashPrime3 uint64 =  1609587929392839161
	xxHashPrime4 uint64 =  9650029242287828579
	xxHashPrime5 uint64 = 2870177450012600261

	xxHashSeed uint64 = 0x9E3779B185EBCA87
)

// XXHash returns a hash using the XXH64. Pass seed or 0 for default value (golden ratio).
func XXHash64(input string, seed uint64) uint64 {

	if seed == 0 {
		seed = xxHashSeed
	}

	xxHashV1 := seed + xxHashPrime1 + xxHashPrime2
	xxHashV2 := seed + xxHashPrime2
	xxHashV3 := seed
	xxHashV4 := seed - xxHashPrime1
	hash := seed

	inpBytes := []byte(input)
	lenCount := len(inpBytes)
	
	offset := 0

	for lenCount - offset >= 32 {

		v1 := binary.LittleEndian.Uint64(inpBytes[offset + 0: offset + 8])
		xxHashV1 = round(xxHashV1, v1)


		v2 := binary.LittleEndian.Uint64(inpBytes[offset + 8: offset + 16])
		xxHashV2 = round(xxHashV2, v2)


		v3 := binary.LittleEndian.Uint64(inpBytes[offset + 16: offset + 24])
		xxHashV3 = round(xxHashV3, v3)


		v4 := binary.LittleEndian.Uint64(inpBytes[offset + 24: offset + 32])
		xxHashV4 = round(xxHashV4, v4)

		hash = rotate64(xxHashV1, 1) + rotate64(xxHashV2, 7) + rotate64(xxHashV3, 12) + rotate64(xxHashV4, 18)

		offset += 32
	}

	for lenCount - offset >= 8 {
		hash ^= binary.LittleEndian.Uint64(inpBytes[(offset) : (offset + 8)]) * xxHashPrime2
		hash = rotate64(hash, 27) * xxHashPrime1 + xxHashPrime4
		offset += 8
	}

	for lenCount - offset >= 4 {
		val := binary.LittleEndian.Uint32(inpBytes[(offset) : (offset + 4)])
		hash ^= uint64(val) * xxHashPrime1
		hash = rotate64(hash, 23) * xxHashPrime2 + xxHashPrime3
		offset += 4
	}

	for lenCount - offset >= 1 {
		val := (inpBytes[offset])
		hash ^= uint64(val) * xxHashPrime5
		hash = rotate64(hash, 11) * xxHashPrime1
		offset += 1
	}

	hash ^= hash >> 33
	hash *= xxHashPrime2
	hash ^= hash >> 29
	hash *= xxHashPrime3
	hash ^= hash >> 32

	return hash
}

func rotate64(x uint64, k int) uint64 {
	return bits.RotateLeft64(x, k)
}

func round(acc, input uint64) uint64 {
	acc += input * xxHashPrime1
	acc = bits.RotateLeft64(acc, 31)
	acc *= xxHashPrime1
	return acc
}