package utils

import (
	"strconv"
	"strings"
)

// determine whether to include
func ContainsString(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// compare version
func VersionCompare(version1, version2 string) int {
	a := strings.Split(version1, ".")
	b := strings.Split(version2, ".")
	flag := 1
	if len(a) > len(b) {
		a, b = b, a
		flag = -1
	}
	for i := range a {
		x, _ := strconv.Atoi(a[i])
		y, _ := strconv.Atoi(b[i])
		if x < y {
			return -1 * flag
		} else if x > y {
			return 1 * flag
		}
	}
	for _, v := range b[len(a):] {
		y, _ := strconv.Atoi(v)
		if y > 0 {
			return -1 * flag
		}
	}
	return 0
}

// HasSuffixes check string has suffixes in string array.
func HasSuffixes(str string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(str, suffix) {
			return true
		}
	}
	return false
}
