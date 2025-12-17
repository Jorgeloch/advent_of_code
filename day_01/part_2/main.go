package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

const DIAL_SIZE = 100

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

		password = password + (distance / DIAL_SIZE) // get the number of full rotations
		distance = mod(distance, DIAL_SIZE)          // get the remaining distance after full rotations

		// add the verification for the new position to see if it reaches zero at some point of the last rotation
		switch direction {
		case 'R':
			{
				if distance >= (DIAL_SIZE - dial_position) {
					password++
				}
				dial_position = mod(dial_position+distance, DIAL_SIZE)
			}
		case 'L':
			{
				// dial position != 0 grants that there will not be any doubled counts
				if dial_position != 0 && distance >= dial_position {
					password++
				}
				dial_position = mod(dial_position-distance, DIAL_SIZE)
			}
		}
	}

	log.Println(password)

	if err := scanner.Err(); err != nil {
		log.Fatal("error reading file")
	}
}
