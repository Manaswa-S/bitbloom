package hashing

const (
    FNV1AOffest uint64 = 14695981039346656037
    FNV1APrime64 uint64 = 1099511628211
)

// FNV1A returns a basic FNV 1a hash (64-bit) of given input string.
func FNV1A(input string) uint64 {
	
    hash := FNV1AOffest

    for i := range len(input) {
        hash ^= uint64(input[i])
        hash *= FNV1APrime64
    }

    return hash
}


