package bitbloom

import (
	"encoding/json"
	"fmt"
	"math/bits"
	"os"
	"sync"

	"github.com/Manaswa-S/bitbloom/hashing"
)


type BitBloom struct {
	BloomSize uint64 // the total number of bits in the filter
	bloom []byte // the filter array
	mu sync.Mutex
	irotations uint64 // the number of index rotations to perform while index calculation

	// 			func(toHash, seed) hash
	hashFuncs []func(string, uint64) uint64 // only two for now
	hashSeeds []uint64 // the seeds for the hash funcs
}

// NewBitBloom initializes a new bloom filter. 
// bloomSize is the total number of bits in the array and ir is the count of index rotations to perform.
// Pass 0 to use default values.
// 
// 	Default Values : 
//	ir = 4
//	bloomSize = 20_000
func NewBitBloom(bloomSize uint64, ir uint64, hFs []func(string, uint64) uint64, hSs []uint64) *BitBloom {

	// TODO: maybe custom/configurable user given hash functions ?

	var defaultindexRotations uint64 = 4
	var defaultBloomSize uint64 = 20_000

	if ir < 1 {
		ir = defaultindexRotations
	}

	if bloomSize < 1 {
		bloomSize = defaultBloomSize
	}

	if len(hFs) < 2 {
		hFs = append(hFs, hashing.XXHash64)
		hFs = append(hFs, hashing.Murmur3_32)

		hSs = append(hSs, 0)
		hSs = append(hSs, 0)
	}

	return &BitBloom{
		BloomSize: bloomSize,
		bloom: make([]byte, (bloomSize / 8) + 1),

		mu: sync.Mutex{},
		irotations: ir,
		hashFuncs: hFs,
		hashSeeds: hSs,
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// Core Functionality

// Add adds a new elem to the array.
func (bb *BitBloom) Add(elem string) bool {
	indexes := bb.getHashes(elem)

	for _, index := range indexes {
		bb.mu.Lock()
		bb.bloom[index / 8] |= (1 << (index % 8))
		bb.mu.Unlock()
	}
    return true
}

// Contains checks if the given elem is probably present in the array or not at all.
func (bb *BitBloom) Contains(elem string) bool {
	indexes := bb.getHashes(elem)

	for _, index := range indexes {
		if bb.bloom[index / 8] & (1 << (index % 8)) == 0 {
			return false
		}
	}
	return true
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// Helpers

// getHashes gets the indexes determined from the hashes.
func (bb *BitBloom) getHashes(elem string) []uint64 {

	hash1_xh := bb.hashFuncs[0](elem, bb.hashSeeds[0])
	hash2_mh := bb.hashFuncs[1](elem, bb.hashSeeds[1])

	indexes := make([]uint64, bb.irotations)  
	for i := uint64(1); i <= bb.irotations; i++ {
		// TODO: may want to use something better than just (i * hash)
		indexes[i] = (hash1_xh + (i * hash2_mh)) % bb.BloomSize
	}

	return indexes
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// Debug capabilities

// PrintBloom prints the entire bloom array one byte at a time.
func (bb *BitBloom) PrintBloom() {
	// TODO: maybe a better way to visualize this ?
    for i, b := range bb.bloom {
        fmt.Printf("Byte %d: %08b\n", i, b)
    }
}

// BitCount returns the ones count in the bloom array.
func (bb *BitBloom) OnesCount() int {
	count := 0
	for _, b := range bb.bloom {
		count += bits.OnesCount8(b)
	}
	return count
}

// Reset sets all bits of the bloom array to 0.
func (bb *BitBloom) Reset() {
	bb.mu.Lock()
	defer bb.mu.Unlock()
	for i := range bb.bloom {
		bb.bloom[i] = 0
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// Bloom persistence

// SaveToFile saves the core bloom array to a simple file in pure bytes.
func (bb *BitBloom) SaveToFile(filename string) error {
	return os.WriteFile(filename, bb.bloom, 0644)
}

// LoadFromFile loads the pure bytes to the bloom array.
func (bb *BitBloom) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	bb.BloomSize = uint64(len(data))
	// TODO: this is wrong, length errors can occur
	copy(bb.bloom, data)
	return nil
}

// TODO: add better export/import capabilities like gob,json along with metadata, etc
// also we can then add a checksum to the file.

// TODO: maybe add a statistics counter that shows hit rate, etc when used. tracks count globally.

// TODO: add dynamic resizing of the main bloom array.

// TODO: add benchmarks and unit tests











// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// trials

type BitBloomCaps struct {
	BloomSize uint64 // the total number of bits in the filter
	Bloom []byte // the filter array
	IRotations uint64 // the number of index rotations to perform while index calculation
}

func (bb *BitBloom) SaveJSON() {

	data := BitBloomCaps{
		BloomSize: bb.BloomSize,
		Bloom: bb.bloom,
		IRotations: bb.irotations,
	}

	f, err := os.Create("json")
	if err != nil {
		fmt.Println(err)
		return
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	err = enc.Encode(data)
	if err != nil {
		fmt.Println(err)
		return
	}

}	