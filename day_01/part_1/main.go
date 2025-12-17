package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const DIAL_SIZE = 100

func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dial_position := 50
	password := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("error parsing distance")
		}

		switch direction {
		case 'R':
			dial_position = mod(dial_position+distance, DIAL_SIZE)
		case 'L':
			dial_position = mod(dial_position-distance, DIAL_SIZE)
		}

		if dial_position == 0 {
			password++
		}
	}

	log.Println(password)

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading file")
	}
}
