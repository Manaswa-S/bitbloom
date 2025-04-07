# BitBloom

A lightweight, thread-safe, and customizable <a href="https://en.wikipedia.org/wiki/Bloom_filter" target="_blank">Bloom Filter</a> library in Go — built for speed, simplicity, and scale.

---
##### ! Under Progress
---

## Features

- Efficient probabilistic set membership testing
- Plug-and-play with custom hash functions
- Thread-safe for concurrent use
- Bit-level control & visualization utilities
- Persist/Load from file or JSON

---
> _Note: BitBloom was built as a foundational exercise. Expect bigger, more intricate systems soon._
---

## Install

```bash
go get github.com/Manaswa-S/bitbloom/bitbloom
```

## Usage
```go
import "github.com/Manaswa-S/bitbloom/bitbloom"

// create a new instance
bloom := bitbloom.NewBitBloom(20000, 4, nil, nil)
// add an element
bloom.Add("golang")
// check for an element
exists := bloom.Contains("golang") // true (probably)
```

## Persistence
```go
// save bit array
bloom.SaveToFile("filter.bloom")
// load bit array
bloom.LoadFromFile("filter.bloom")
```
## Debugging & Stats
```go
// print internal bit array
bloom.PrintBloom()
// ones count in the bit array
fmt.Println(bloom.OnesCount())
```
## Roadmap
- Dynamic Resizing
- Bit level performance optimizations
- Stats
- Custom seeding

------
### DETAILED EXPLAINATION

##### Hashing

- The bloom filter uses a simple `[]byte` slice to simulate a large bit array.
- Instead of using `k` different hash functions, BitBloom uses **2 base hashes** and derives multiple positions using:

```go
for i := 0; i < rotations; i++ {
  index{i} = (hash1 + (i * hash2)) % bloomSize
}
```
- It uses custom implementation of **XxHash 64-bit** and **MurMur Hash 64-bit**.
  
##### Configuration

- You can customize the behavior of BitBloom using the `NewBitBloom()` constructor.

- BitBloom supports filter size and index rotation customatizations and **plugging in custom hash functions** as long as they match this signature:

```go
func(input string, seed uint64) uint64
// also Hash seeds are to be provided separately for them, if required.
```
##### Saving

- Saving a writes the bloom array as it is to a file, and therefore might incur corrupt files sometimes. (this will be resolved in further revisions by implementing a checksum)


---

##### License

[MIT License](./LICENSE)


