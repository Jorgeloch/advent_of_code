package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func compute_values(start uint64, end uint64) uint64 {
	n_digits_start := uint64(math.Floor(math.Log10(float64(start))) + 1)
	n_digits := uint64(math.Floor(float64(n_digits_start) / 2))

	prefix := uint64(math.Floor(float64(start) / math.Pow(10, float64(n_digits_start-n_digits))))

	counter := uint64(0)

	for {
		n_digits_prefix := uint64(math.Floor(math.Log10(float64(prefix))) + 1)

		pattern := prefix * (uint64(math.Pow(10, float64(n_digits_prefix))) + 1)

		if pattern > end {
			break
		}

		if pattern >= start && pattern <= end {
			counter += pattern
		}

		prefix++
	}

	return counter
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := uint64(0)

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "-")

		start, err := strconv.ParseUint(values[0], 10, 64)
		if err != nil {
			log.Fatal("error parsing start value")
		}
		end, err := strconv.ParseUint(values[1], 10, 64)
		if err != nil {
			log.Fatal("error parsing end value")
		}

		counter += compute_values(start, end)
	}

	log.Println("Invalid IDs:", counter)

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading file")
	}
}
