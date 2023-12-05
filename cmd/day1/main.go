package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var (
	ErrNoDigitFound = errors.New("no number found")
)

var numbers = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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

	fmt.Printf("the decoded value for part one is [%d]\n", sum)

	sum = 0
	f2, err := os.Open("cmd/day1/data2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	scanner = bufio.NewScanner(f2)
	for scanner.Scan() {
		digit, err := decodeLine(fixLine(scanner.Text()))
		if err != nil {
			if errors.Is(err, ErrNoDigitFound) {
				continue
			}
			log.Fatal(err)
		}
		sum = sum + digit
	}

	fmt.Printf("the decoded value for part two is [%d]\n", sum)
}

func fixLine(line string) string {
	var regs []*regexp.Regexp
	for _, num := range numbers {
		re := regexp.MustCompile(num)
		regs = append(regs, re)
	}

	for i, re := range regs {
		line = re.ReplaceAllString(line, fmt.Sprintf("%d", i))
	}

	return line
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
