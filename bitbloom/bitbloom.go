package bitbloom

import (
	"bitbloom/hashing"
	"fmt"
)


type BitBloom struct {
	BloomSize int64 // the total number of bits in the filter
	bloomLen int64 // the length of the filter (bytes)
	bloom []byte // the filter array
	HashFuns []func(string) uint64
}

func NewBitBloom(bloomSize int64, hashFuns []func(string) uint64) *BitBloom {
	bloomLen := bloomSize / 8 + 1
	return &BitBloom{
		BloomSize: bloomSize,
		bloom: make([]byte, bloomLen),
		HashFuns: hashFuns,
	}
}


func (bb *BitBloom) AddElem(elem string) bool {
	// 1) Get hashes 

	// 1.a) can use custom hashes too, TODO:

	xhash := hashing.XXHash64(elem, 0)
	murhash := hashing.Murmur3_32(elem, 0)

	fmt.Printf("XHS : %d : MH : %d\n", xhash, murhash)

    return true
}

// func findElem(elem string) bool {
//     bitIndex := FNV1A(elem) % uint64(bloommodsize)
//     byteIndex := bitIndex / 8
//     bitPos := bitIndex % 8
    
//     fmt.Printf("%d : %d\n", byteIndex, bitPos)


//     if bloom[byteIndex] & (1 << bitPos) != 0 {
//         return true
//     }

//     return false
// }

// func printBloom() {
//     for i, b := range bloom {
//         fmt.Printf("Byte %d: %08b\n", i, b)
//     }
// }
