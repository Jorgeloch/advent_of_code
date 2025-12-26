package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// uses the pair (i, n/i) algorithm to compute factors of a number in O(sqrt(n))
func compute_factors(n uint64) []uint64 {
	upper_limit := uint64(math.Ceil(math.Sqrt(float64(n))))
	factors := make([]uint64, 0)
	for i := uint64(1); i <= upper_limit; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			result := n / i
			if result != i && result != n {
				factors = append(factors, result)
			}
		}
	}
	slices.Sort(factors)
	return factors
}

// generates a pattern based on a prefix, that will be repeated at least 2 times
func generate_pattern(prefix uint64, n_digits_prefix uint64, n_digits uint64) uint64 {
	factor := n_digits / n_digits_prefix

	// ensures prefix will be repeated at least 2 times
	if factor == 1 {
		return 0
	}

	// generate all the patterns
	result := uint64(0)
	for i := range factor {
		result += prefix * uint64(math.Pow10(int(i*n_digits_prefix)))
	}
	return result
}

// computes the sum of all patterns that generates an invalid id in a range [start, end]
func compute_values(start uint64, end uint64) uint64 {
	n_digits_start := uint64(math.Floor(math.Log10(float64(start))) + 1)
	n_digits_end := uint64(math.Floor(math.Log10(float64(end))) + 1)

	// this hash map is used to store the patterns that have been computed
	// avoiding counting the same pattern multiple times
	memo := make(map[uint64]bool)

	counter := uint64(0)

	// iterate over the number of digits range
	for i := n_digits_start; i <= n_digits_end; i++ {
		// get the number of factors for that number of digits
		factors := compute_factors(i)

		// generate all the patterns with the given number of digits
		for _, n_digits := range factors {
			// TODO: implement a smarter way to generate the first prefix, avoiding generating
			// a lot of patterns out of the range
			first_prefix := 1 * uint64(math.Pow10(int(n_digits-1)))
			for j := first_prefix; j < first_prefix*10; j++ {
				pattern := generate_pattern(j, n_digits, i)
				if memo[pattern] {
					continue
				}
				if pattern <= end && pattern >= start {
					counter += pattern
					memo[pattern] = true
				}
			}
		}
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
