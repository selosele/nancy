package utils

import "strings"

/* 문자열 매개변수가 비어 있는지 확인한다. */
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

/* 문자열 매개변수가 존재하는지 확인한다. */
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

/* 문자열 첫 번째 매개변수가 없으면 두 번째 매개변수를 반환한다. */
func DefaultString(s string, d string) string {
	if IsNotBlank(s) {
		return s
	}
	return d
}
