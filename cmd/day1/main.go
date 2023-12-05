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
	sum2 := 0
	f, err := os.Open("cmd/day1/data1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		digit, err := decodeLine(line)
		if err != nil {
			if errors.Is(err, ErrNoDigitFound) {
				continue
			}
			log.Fatal(err)
		}
		sum = sum + digit

		digit2, err := decodeLine(fixLine(line))
		if err != nil {
			if errors.Is(err, ErrNoDigitFound) {
				continue
			}
			log.Fatal(err)
		}
		sum2 = sum2 + digit2
	}

	fmt.Printf("the decoded value for part one is [%d]\n", sum)

	fmt.Printf("the decoded value for part two is [%d]\n", sum2)
}

func fixLine(line string) string {
	var regs []*regexp.Regexp
	for _, num := range numbers {
		re := regexp.MustCompile(num)
		regs = append(regs, re)
	}
	for i, re := range regs {
		//fmt.Printf("%s => ", line)
		line = re.ReplaceAllString(line, strconv.Itoa(i))
		//fmt.Printf("%s\n", line)

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
