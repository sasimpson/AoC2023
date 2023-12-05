package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

var (
	ErrNoDigitFound = errors.New("no number found")
)

func main() {
	sum := 0
	f, err := os.Open("cmd/day1/data1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		digit, err := decodeLine(scanner.Text())
		if err != nil {
			if errors.Is(err, ErrNoDigitFound) {
				continue
			}
			log.Fatal(err)
		}
		sum = sum + digit
	}

	fmt.Printf("the decoded value is %d\n", sum)
}

func decodeLine(line string) (int, error) {

	forward, err := getNumber(line)
	if err != nil {
		return 0, fmt.Errorf("%w in forward", err)
	}

	bLine := []byte(line)
	slices.Reverse(bLine)
	revLine := string(bLine)

	backward, err := getNumber(revLine)
	if err != nil {
		return 0, fmt.Errorf("%w in reverse", err)
	}

	return strconv.Atoi(fmt.Sprintf("%d%d", forward, backward))
}

func getNumber(data string) (int, error) {
	for _, c := range data {
		if c >= 48 && c <= 58 {
			return int(c) - 48, nil
		}
	}
	return 0, ErrNoDigitFound
}
