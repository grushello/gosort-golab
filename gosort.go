package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
)

// -----------------------------
// Entry point
// -----------------------------

func main() {
	n := flag.Int("r", -1, "generate N random integers (N >= 10)")
	flag.Parse()

	if *n == -1 {
		log.Fatal("Usage: gosort -r N")
	}

	if err := runRandom(*n); err != nil {
		log.Fatal(err)
	}
}

// -----------------------------
// -r mode logic
// -----------------------------

func runRandom(n int) error {
	if n < 10 {
		return errors.New("N must be >= 10")
	}

	numbers := generateRandomNumbers(n)

	fmt.Println("Original numbers:")
	fmt.Println(numbers)

	chunks := splitIntoChunks(numbers)

	fmt.Println("\nChunks before sorting:")
	printChunks(chunks)

	sortedChunks := sortChunksConcurrently(chunks)

	fmt.Println("\nChunks after sorting:")
	printChunks(sortedChunks)

	result := mergeSortedChunks(sortedChunks)

	fmt.Println("\nFinal sorted result:")
	fmt.Println(result)

	return nil
}

// -----------------------------
// Chunking logic
// -----------------------------

func splitIntoChunks(numbers []int) [][]int {
	n := len(numbers)

	numChunks := int(math.Ceil(math.Sqrt(float64(n))))
	if numChunks < 4 {
		numChunks = 4
	}

	// TODO:
	// Split numbers into numChunks roughly equal slices
	// Return [][]int

	return nil
}

// -----------------------------
// Concurrent sorting
// -----------------------------

func sortChunksConcurrently(chunks [][]int) [][]int {
	// TODO:
	// For each chunk:
	//   - start a goroutine
	//   - sort the chunk (any algorithm allowed)
	// Synchronize and return sorted chunks

	return nil
}

// -----------------------------
// Merge logic
// -----------------------------

func mergeSortedChunks(chunks [][]int) []int {
	// TODO:
	// Merge multiple sorted slices into one sorted slice
	// Do NOT re-sort everything

	return nil
}

// -----------------------------
// Helpers
// -----------------------------

func generateRandomNumbers(n int) []int {
	// TODO:
	// Generate n random integers
	// Range is up to you

	return nil
}

func printChunks(chunks [][]int) {
	for i, c := range chunks {
		fmt.Printf("Chunk %d: %v\n", i, c)
	}
}
