package utils

import "strings"

/* 문자열 매개변수가 비어 있는지 확인한다. */
func IsBlank(v string) bool {
	return len(strings.TrimSpace(v)) == 0
}

/* 문자열 매개변수가 존재하는지 확인한다. */
func IsNotBlank(v string) bool {
	return !IsBlank(v)
}

/* 문자열 첫 번째 매개변수가 없으면 두 번째 매개변수를 반환한다. */
func DefaultString(v string, d string) string {
	if IsNotBlank(v) {
		return v
	}
	return d
}
