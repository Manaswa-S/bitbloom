package hashing

import (
	"encoding/binary"
	"math/bits"
)


const (
    c1 uint32 = 0xcc9e2d51
    c2 uint32 = 0x1b873593
    r1 = 15
    r2 = 13
    m = 5
    n = 0xe6546b64
    murHashSeed uint32 = 0x85ebca6b
)

// MurMurHash32 returns a basic hash (32-bit). Pass seed or 0 for default value.
func Murmur3_32(input string, seed uint32) uint32 {

    if seed == 0 {
        seed = murHashSeed
    } 
    hash := seed

    inpBytes := []byte(input)
    inpLen := len(inpBytes)

    offset := 0
    for inpLen - offset >= 4 {
        val := binary.LittleEndian.Uint32(inpBytes[offset : offset + 4])

        val *= c1
        val = rotate32(val, r1)
        val *= c2

        hash ^= val
        hash = rotate32(hash, r2)
        hash = 5 * hash + n

        offset += 4
    }

    var tail uint32
    remaining := inpLen - offset
    for i := range remaining {
        tail |= uint32(inpBytes[offset + i]) << (8 * i)
    }

    tail *= c1
    tail = rotate32(tail, r1)
    tail *= c2
    hash ^= tail

    hash ^= uint32(inpLen)
    hash ^= hash >> 16
    hash *= 0x85ebca6b
    hash ^= hash >> 13
    hash *= 0xc2b2ae35
    hash ^= hash >> 16

    return hash
}


func rotate32(x uint32, k int) uint32 {
    return bits.RotateLeft32(x, k)
}