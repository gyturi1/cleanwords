package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

var hun_double_letters map[string]bool = map[string]bool{
	"cs": true,
	"dz": true,
	"gy": true,
	"ly": true,
	"ny": true,
	"sz": true,
	"ty": true,
	"zs": true,
}

var hun_tripple_letter = "dzs"

var hun_single_letters [35]rune = [35]rune{
	'a', 'á', 'b', 'c', 'd', 'e', 'é',
	'f', 'g', 'h', 'i', 'í', 'j', 'k',
	'l', 'm', 'n', 'o', 'ó', 'ö', 'ő',
	'p', 'q', 'r', 's', 't', 'u', 'ú',
	'ü', 'ű', 'v', 'w', 'x', 'y', 'z',
}

const input = "all.txt"
const output = "possibleFiveLetterWords.txt"

var enc = charmap.ISO8859_2

func main() {
	sc, f := readFile(input)
	defer f.Close()

	out, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	for sc.Scan() {
		s := string(sc.Bytes())
		if isValidWord(s) && isFiveLetterWord(s) {
			if _, err := out.WriteString(fmt.Sprintln(strings.ToLower(s))); err != nil {
				panic(err)
			}
		}
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}
}

func readFile(filename string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	r := transform.NewReader(f, enc.NewDecoder())

	return bufio.NewScanner(r), f
}

func isFiveLetterWord(s string) bool {
	return isAnyFive(possibleLetterCounts(s))
}

func possibleLetterCounts(s string) []int {
	return countLetters([]rune(s), 0)
}

func countLetters(rs []rune, count int) []int {
	if len(rs) == 0 {
		return []int{count}
	}
	firstSingle := countLetters(rs[1:], count+1)

	var firstDouble []int
	if isDoubleLetter(rs[:2]) {
		firstDouble = countLetters(rs[2:], count+1)
	}

	var firstTripple []int
	if isTrippleLetter(rs[:3]) {
		firstTripple = countLetters(rs[3:], count+1)
	}

	var ret []int
	ret = append(ret, firstSingle...)
	ret = append(ret, firstDouble...)
	ret = append(ret, firstTripple...)
	return ret
}

func isTrippleLetter(l []rune) bool {
	return string(l) == hun_tripple_letter
}

func isDoubleLetter(l []rune) bool {
	_, ok := hun_double_letters[string(l)]
	return ok
}

func isAnyFive(ps []int) bool {
	for _, p := range ps {
		if p == 5 {
			return true
		}
	}
	return false
}

func isValidWord(s string) bool {
	i := utf8.RuneCountInString(s)
	if i > 4 && i < 11 {
		c := 0
		for _, l := range s {
			if isValidLetter(l) {
				c++
			}
		}
		if c == i {
			return true
		}
	}
	return false
}

func isValidLetter(l rune) bool {
	for _, r := range hun_single_letters {
		if l == r {
			return true
		}
	}
	return false
}
