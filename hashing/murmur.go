package hashing

import (
	"encoding/binary"
	"math/bits"
)


const (
    c1 uint64 = 0xcc9e2d51
    c2 uint64 = 0x1b873593
    r1 = 15
    r2 = 13
    m = 5
    n = 0xe6546b64
    murHashSeed uint64 = 0x85ebca6b
)

// MurMurHash32 returns a basic hash (32-bit). Pass seed or 0 for default value.
func Murmur3_32(input string, seed uint64) uint64 {

    // TODO: change this function later to aid in further calculations to uint64

    if seed == 0 {
        seed = murHashSeed
    } 
    hash := seed

    inpBytes := []byte(input)
    inpLen := len(inpBytes)

    offset := 0
    for inpLen - offset >= 8 {
        val := binary.LittleEndian.Uint64(inpBytes[offset : offset + 8])

        val *= c1
        val = rotate64(val, r1)
        val *= c2

        hash ^= val
        hash = rotate64(hash, r2)
        hash = 5 * hash + n

        offset += 8
    }

    for inpLen - offset >= 4 {
        val := binary.LittleEndian.Uint32(inpBytes[offset : offset + 4])

        val *= uint32(c1)
        val = rotate32(val, r1)
        val *= uint32(c2)

        hash ^= uint64(val)
        hash = rotate64(hash, r2)
        hash = 5 * hash + n

        offset += 4
    }

    var tail uint64
    remaining := inpLen - offset
    for i := range remaining {
        tail |= uint64(inpBytes[offset + i]) << (8 * i)
    }

    tail *= c1
    tail = rotate64(tail, r1)
    tail *= c2
    hash ^= tail

    hash ^= uint64(inpLen)
    hash ^= hash >> 33
    hash *= 0xff51afd7ed558ccd
    hash ^= hash >> 33
    hash *= 0xc4ceb9fe1a85ec53
    hash ^= hash >> 33

    return hash
}


func rotate32(x uint32, k int) uint32 {
    return bits.RotateLeft32(x, k)
}

// func rotate64(x uint64, k int) uint64 {
//     return bits.RotateLeft64(x, k)
// }