package hkid

import (
	"regexp"
	"strconv"
	"strings"
)

func clearString(s string) string {
	newString := strings.Trim(s, " ")
	newString = strings.ToUpper(newString)
	return newString
}

func getRemainder(hkid *hkid) string {
	charSum := getCharSum(hkid.part1)
	hkidSum := calPart2Remainder(hkid.part2, charSum)

	remainder := ""

	switch hkidSum {
	case 11:
		remainder = "0"
	case 10:
		remainder = "A"
	default:
		remainder = strconv.Itoa(hkidSum)
	}

	return remainder
}

func calPart2Remainder(s string, i int) int {
	sum := 0

	for k, v := range s {
		i, _ := strconv.Atoi(string(v))
		sum += (7 - k) * i
	}

	mod := 11
	return mod - ((i + sum) % mod)
}

func getCharSum(part1 string) int {
	charMap := getCharMap()
	if len(part1) == 1 {
		return 324 + charMap[part1]*8
	}

	total := 0
	weight := getWeight()
	for k, v := range part1 {
		total += weight[k] * charMap[string(v)]
	}
	return total
}

func getWeight() map[int]int {
	weight := make(map[int]int)
	weight[0] = 9
	weight[1] = 8
	return weight
}

func getCharMap() map[string]int {
	asciiNum := 65 // Uppercase A
	idCharMap := make(map[string]int)

	for i := 0; i <= 26; i++ {
		idCharMap[string(i+asciiNum)] = 10 + i
	}

	return idCharMap
}

func validatePatten(string string) (*hkid, error) {
	r, _ := regexp.Compile(`^(?P<p1>\D{1,2})(?P<p2>\d{6})\((?P<p3>[\w{1}0-9aA])\)$`)
	match := r.FindStringSubmatch(string)
	if len(match) == 0 {
		return newHkidNil(), newPatterNotMatchError()
	}
	return newHkid(match[1], match[2], match[3]), nil
}
