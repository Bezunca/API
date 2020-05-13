package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func CleanString(s string) string {
	if s != "" {
		space := regexp.MustCompile(`\s+`)
		return space.ReplaceAllString(strings.ReplaceAll(strings.TrimSpace(s), "\n", ""), " ")
	} else {
		return ""
	}
}

func StringToInt(s string) int {
	sInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	} else {
		return sInt
	}
}

func StringToFloat32(s string) float32 {
	s = strings.ReplaceAll(s, ",", ".")
	sFloat, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	} else {
		return float32(sFloat)
	}
}
