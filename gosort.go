//Viacheslav Pototskyi 241ADB183
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// -----------------------------
// Entry point
// -----------------------------

func main() {
	r := flag.Int("r", -1, "generate N random integers (N >= 10)")
	i := flag.String("i", "", "input file with integers")
	d := flag.String("d", "", "directory containing .txt files")
	flag.Parse()

	switch {
	case *r != -1:
		if err := runRandom(*r); err != nil {
			log.Fatal(err)
		}
	case *i != "":
		if err := runFile(*i); err != nil {
			log.Fatal(err)
		}
	case *d != "":
		if err := runDirectory(*d); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Usage: gosort -r N | -i input.txt | -d directory")
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

	fmt.Println("Initial numbers:")
	fmt.Println(numbers)

	chunks := splitIntoChunks(numbers)

	fmt.Println("\nChunks before sorting:")
	printChunks(chunks)

	sortedChunks := sortChunksConcurrently(chunks)

	fmt.Println("\nChunks after sorting:")
	printChunks(sortedChunks)

	result := mergeSortedChunks(sortedChunks)

	fmt.Println("\nFinal result:")
	fmt.Println(result)

	return nil
}

// File and directory logic

func runFile(filename string) error {
	numbers, err := readNumbersFromFile(filename)
	if err != nil {
		return err
	}

	if len(numbers) < 10 {
		return errors.New("file must contain at least 10 integers")
	}

	runAndPrint(numbers)
	return nil
}

func runDirectory(dir string) error {
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		return errors.New("invalid directory")
	}

	outDir := dir + "_sorted_Viacheslav_Pototskyi_241ADB183"
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".txt" {
			continue
		}

		inputPath := filepath.Join(dir, f.Name())
		numbers, err := readNumbersFromFile(inputPath)
		if err != nil {
			return err
		}
		if len(numbers) < 10 {
			return fmt.Errorf("%s has fewer than 10 numbers", f.Name())
		}

		chunks := splitIntoChunks(numbers)
		sortedChunks := sortChunksConcurrently(chunks)
		result := mergeSortedChunks(sortedChunks)

		outputPath := filepath.Join(outDir, f.Name())
		if err := writeNumbersToFile(outputPath, result); err != nil {
			return err
		}
	}

	return nil
}
// common printing logic

func runAndPrint(numbers []int) {
	fmt.Println("Initial numbers:")
	fmt.Println(numbers)

	chunks := splitIntoChunks(numbers)

	fmt.Println("\nChunks before sorting:")
	printChunks(chunks)

	sortedChunks := sortChunksConcurrently(chunks)

	fmt.Println("\nChunks after sorting:")
	printChunks(sortedChunks)

	result := mergeSortedChunks(sortedChunks)

	fmt.Println("\nFinal result:")
	fmt.Println(result)
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

	baseSize := n / numChunks
	remainder := n % numChunks

	chunks := make([][]int, 0, numChunks)

	start := 0
	for i := 0; i < numChunks; i++ {
		size := baseSize
		if i < remainder {
			size++
		}
		end := start + size
		if start >= n {
			break
		}
		if end > n {
			end = n
		}
		chunks = append(chunks, numbers[start:end])
		start = end
	}

	return chunks
}

// -----------------------------
// Concurrent sorting
// -----------------------------

func sortChunksConcurrently(chunks [][]int) [][]int {
	var wg sync.WaitGroup
	wg.Add(len(chunks))

	for i := range chunks {
		go func(idx int) {
			defer wg.Done()
			sort.Ints(chunks[idx])
		}(i)
	}

	wg.Wait()
	return chunks
}

// -----------------------------
// Merge logic
// -----------------------------

func mergeSortedChunks(chunks [][]int) []int {
	totalSize := 0
	for _, c := range chunks {
		totalSize += len(c)
	}

	result := make([]int, 0, totalSize)
	indices := make([]int, len(chunks))

	for len(result) < totalSize {
		minChunk := -1
		minValue := 0

		for i, c := range chunks {
			if indices[i] >= len(c) {
				continue
			}
			val := c[indices[i]]
			if minChunk == -1 || val < minValue {
				minValue = val
				minChunk = i
			}
		}

		result = append(result, minValue)
		indices[minChunk]++
	}

	return result
}

// -----------------------------
// Helpers
// -----------------------------

func readNumbersFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %s", line)
		}
		numbers = append(numbers, val)
	}

	return numbers, scanner.Err()
}

func writeNumbersToFile(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, n := range numbers {
		fmt.Fprintln(w, n)
	}
	return w.Flush()
}


func generateRandomNumbers(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}
	return nums
}

func printChunks(chunks [][]int) {
	for i, c := range chunks {
		fmt.Printf("Chunk %d: %v\n", i, c)
	}
}
