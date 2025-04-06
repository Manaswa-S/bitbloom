package main

import (
	"bitbloom/bitbloom"
	"bitbloom/hashing"
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

var bloomSize int64 = 200_000

func main() {
	// bloomMenu()
	// testMurMur()
	// testxxHash()
	testDistribution()
}


func bloomMenu() {
    fmt.Printf("----- Bloom Filter Implementation In Go ----- \n")
    
	bb := bitbloom.NewBitBloom(bloomSize, nil)

    menuStart:

    fmt.Printf( 
    "\n1) Add one element. \n" + 
    "2) Search for an element. \n" + 
    "3) Add entire array of elements. \n" + 
    "4) Print entire bloom array. \n" +
    "5) MurMur Demo \n" +
    "0) Exit. \n" + 
    " > Choice : ")

    choice := 9
    _, err := fmt.Scan(&choice)
    if err != nil {
        fmt.Printf("Something went wrong : %v", err)
        return
    }

    switch choice {
    case 1:
        // Add one element 
        var elem string
        fmt.Print("Enter new element : ")
        _, err := fmt.Scanln(&elem)
        if err != nil {
            fmt.Printf("Something went wrong : %v\n", err)
            break
        }

        done := bb.AddElem(elem)
        if !done {
            fmt.Println("Something went wrong : could not add new element.")
            break
        }
        fmt.Println("Added new element.")

    case 2:
        // Search for an element.
        fmt.Print("Enter element to search : ")
        reader := bufio.NewReader(os.Stdin)
        toSearch, err := reader.ReadString('\n')
        if err != nil {
            fmt.Printf("Something went wrong : %v\n", err)
            break
        }
        fmt.Printf("\nSearching for : %s\n", toSearch)

        // found := findElem(toSearch)
        // if found {
        //     fmt.Println("Element probably exists.")
        // } else {
        //     fmt.Println("Element does not exist.")
        // }

    case 3:
        // Add entire slice.
        elems := WordsGenerator()
        for _, elem := range elems {
            done := bb.AddElem(elem)
            if !done {
                fmt.Println("Something went wrong : could not add new element.")
            }
        }

    case 4:
        // Print entire bloom array as bits
        // printBloom()
    case 5:
        // MurMur("")
    case 0:
        fmt.Println("Exiting ...")
        return
    default:
        fmt.Println("Wrong choice. Try again.")
    }

    goto menuStart
}












// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// TESTS

func WordsGenerator() (words []string) {
    // file, err := os.OpenFile("./tests/sentences.txt", 0644, os.FileMode(os.O_RDONLY))

    file, err := os.OpenFile("./tests/movie_list.txt", 0644, os.FileMode(os.O_RDONLY))
    if err != nil {
        fmt.Println(err)
        return
    }

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    if scanner.Err() != nil {
        fmt.Println(scanner.Err())
        return
    }

	return
}


func testMurMur() {

	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// counter := 0

	rotationsCount := 100

	strings := WordsGenerator()
	stringsLen := len(strings)

	// hashSet := make(map[uint32]string)
	// collisions := 0
	
	start := time.Now()
	for i := 0; i < rotationsCount; i++ {
		// wg.Add(1)
		// go func() {
			for _, str := range strings {
				_ = hashing.Murmur3_32(str, 0)

				// if ex, ok := hashSet[got]; ok {
				// 	if ex != str {
				// 		collisions++
				// 		fmt.Printf("%s === %s\n", str, ex)
				// 	}
				// } else {
				// 	hashSet[got] = str
				// }

				
				// mu.Lock()
				// counter++
				// mu.Unlock()
			}
			// wg.Done()
		// }()
	}

	// wg.Wait()

	ttnano := time.Since(start).Nanoseconds()

	totalstrLen := 0
	for _, str := range strings {
		totalstrLen += len(str)
	}
	avgstrLen := (totalstrLen / stringsLen)

	processedCount := rotationsCount * stringsLen
	
	ttperhash := ttnano / int64(processedCount)

	// collisionRate := float32(collisions) / float32(processedCount)

	fmt.Printf("Time Taken in nanoseconds : %d\n", ttnano)
	fmt.Printf("Average string Len : %d\n", avgstrLen)
	fmt.Printf("Processed : %d\n", processedCount)
	fmt.Printf("TT per Hash MurMur3 in nanoseconds: %d\n", ttperhash)

	// fmt.Printf("Collisions : %d\n", collisions)
	// fmt.Printf("Collision Rate : %f\n", collisionRate)
}

func testxxHash() {

	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// counter := 0

	rotationsCount := 10000

	strings := WordsGenerator()
	stringsLen := len(strings)

	// hashSet := make(map[uint32]string)
	// collisions := 0
	
	start := time.Now()
	for i := 0; i < rotationsCount; i++ {
		// wg.Add(1)
		// go func() {
			for _, str := range strings {
				_ = hashing.XXHash64(str, 0)

				// if ex, ok := hashSet[got]; ok {
				// 	if ex != str {
				// 		collisions++
				// 		fmt.Printf("%s === %s\n", str, ex)
				// 	}
				// } else {
				// 	hashSet[got] = str
				// }

				
				// mu.Lock()
				// counter++
				// mu.Unlock()
			}
			// wg.Done()
		// }()
	}

	// wg.Wait()

	ttnano := time.Since(start).Nanoseconds()

	totalstrLen := 0
	for _, str := range strings {
		totalstrLen += len(str)
	}
	avgstrLen := (totalstrLen / stringsLen)

	processedCount := rotationsCount * stringsLen
	
	ttperhash := ttnano / int64(processedCount)

	// collisionRate := float32(collisions) / float32(processedCount)

	fmt.Printf("Time Taken in nanoseconds : %d\n", ttnano)
	fmt.Printf("Average string Len : %d\n", avgstrLen)
	fmt.Printf("Processed : %d\n", processedCount)
	fmt.Printf("TT per Hash xxHash in nanoseconds: %d\n", ttperhash)

	// fmt.Printf("Collisions : %d\n", collisions)
	// fmt.Printf("Collision Rate : %f\n", collisionRate)
}


func testDistribution() {

	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// counter := 0

	rotationsCount := 1

	strings := WordsGenerator()
	stringsLen := len(strings)

	hashSet := make(map[uint64]string)
	collisions := 0

	buckets := 20_000
	counts := make([]int, buckets)
	
	start := time.Now()
	for i := 0; i < rotationsCount; i++ {
		for _, str := range strings {
			xxhash := hashing.XXHash64(str, 0)
			if ex, ok := hashSet[xxhash]; ok {
				if ex != str {
					collisions++
				}
			} else {
				hashSet[xxhash] = str
				counts[(xxhash * uint64(buckets)) >> 64]++
			}
			// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

			// murhash := hashing.Murmur3_32(str, 0)
			// if ex, ok := hashSet[murhash]; ok {
			// 	if ex != str {
			// 		collisions++
			// 	}
			// } else {
			// 	hashSet[murhash] = str
			// 	counts[murhash % uint32(buckets)]++
			// }
		}
	}

	ttnano := time.Since(start).Nanoseconds()

	totalstrLen := 0
	for _, str := range strings {
		totalstrLen += len(str)
	}
	avgstrLen := (totalstrLen / stringsLen)

	processedCount := rotationsCount * stringsLen
	
	ttperhash := ttnano / int64(processedCount)

	collisionRate := float32(collisions) / float32(processedCount)

	fmt.Printf("Time Taken in nanoseconds : %d\n", ttnano)
	fmt.Printf("Average string Len : %d\n", avgstrLen)
	fmt.Printf("Processed : %d\n", processedCount)
	fmt.Printf("TT per Hash xxHash in nanoseconds: %d\n", ttperhash)

	fmt.Printf("Collisions : %d\n", collisions)
	fmt.Printf("Collision Rate : %f\n", collisionRate)


	// Now analyze variance
	total := processedCount
	expected := float64(total) / float64(buckets)
	var sumSq float64
	for _, count := range counts {
		diff := float64(count) - expected
		sumSq += diff * diff
	}

	stddev := math.Sqrt(sumSq / float64(buckets))
	fmt.Printf("Expected per bucket: %.2f\n", expected)
	fmt.Printf("Standard deviation: %.2f\n", stddev)
	fmt.Printf("Spread score (stddev/expected): %.5f\n", stddev/expected)
}