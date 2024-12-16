package main

import (
	"errors"
	"fmt"
)

// 제네릭 함수 정의
func PrintValue[T any](value T) {
	fmt.Println(value)
}

// 제네릭 제약조건
// comparable 인터페이스는 비교 연산자(==, !=, >, <, >=, <=)를 사용할 수 있는 타입을 의미합니다.
func FindIndex[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

// 제네릭 제약 조건 정의
type Number interface {
	int | int32 | int64 | float64 | float32
}

func Sum[T Number](numbers []T) T {
	var num T
	for _, v := range numbers {
		num += v
	}
	return num
}

// 제네릭 구조체 정의
type Stack[T any] struct {
	items []T
}

// 구조체 생성자
func NewStack[T any]() *Stack[T] {
	// 슬라이스 초기화
	return &Stack[T]{items: []T{}}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T

	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, nil
}

func main() {
	PrintValue(1)
	PrintValue("Hello, World!")
	PrintValue(3.14)

	fmt.Println(FindIndex([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))

	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
