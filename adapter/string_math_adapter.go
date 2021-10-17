package adapter

import (
	"fmt"
	"strconv"
)

type StringMathAdapter struct{}

// Twice()와 Half()는 이미 구현된 알고리즘을 사용한다.
func (s *StringMathAdapter) Twice(a string) string {
	if v, err := strconv.Atoi(a); err == nil {
		return fmt.Sprintf("Adapter : %d", Twice(v))
	} else {
		return "nil"
	}
}

func (s *StringMathAdapter) Half(a string) string {
	if v, err := strconv.Atoi(a); err == nil {
		return fmt.Sprintf("Adapter : %d", Half(v))
	} else {
		return "nil"
	}
}

func NewStringMathAdapter() MathAdapter {
	return &StringMathAdapter{}
}
