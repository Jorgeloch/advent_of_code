package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func compute_values(raw_start string, raw_end string) uint64 {
	// get the length of the start and end strings
	n_digits_start := uint64(math.Floor(float64(len(raw_start)) / 2))
	n_digits_end := uint64(math.Ceil(float64(len(raw_end)) / 2))

	// get the prefixes (first halves of the start and end strings)
	start_prefix, err := strconv.ParseUint(raw_start[:n_digits_start], 10, 64)
	if err != nil {
		log.Fatal("error parsing start prefix")
	}
	end_prefix, err := strconv.ParseUint(raw_end[:n_digits_end], 10, 64)
	if err != nil {
		log.Fatal("error parsing end prefix")
	}

	// parse the start and end values (will be important on the loop)
	start, err := strconv.ParseUint(raw_start, 10, 64)
	if err != nil {
		log.Fatal("error parsing start value")
	}
	end, err := strconv.ParseUint(raw_end, 10, 64)
	if err != nil {
		log.Fatal("error parsing end value")
	}

	counter := uint64(0)

	// generate all possible values between start and end
	// and check if they match the condidion of beeing a invalid id
	for ; start_prefix <= end_prefix; start_prefix++ {
		n_digits := int(math.Floor(math.Log10(float64(start_prefix)))) + 1
		invalid_value := start_prefix * (uint64(math.Pow10(n_digits) + 1))
		if invalid_value > end {
			break
		}
		if invalid_value < end && invalid_value > start {
			counter += invalid_value
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

		start := values[0]
		end := values[1]

		counter += compute_values(start, end)
	}

	log.Println("Invalid IDs:", counter)

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading file")
	}
}
