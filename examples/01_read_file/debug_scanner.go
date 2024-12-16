package main

import (
	"fmt"
	"strings"
)

// 디버그를 위해 스캐너의 현재 상태를 출력하는 기능을 추가한 스캐너
type DebugScanner struct {
	reader    *strings.Reader // strings.Reader로 변경하여 디버깅을 쉽게 함
	buffer    []byte
	position  int
	tokenSize int
}

func NewDebugScanner(content string) *DebugScanner {
	return &DebugScanner{
		reader:    strings.NewReader(content),
		buffer:    make([]byte, 10), // 디버깅을 위해 버퍼 크기를 10으로 줄임
		position:  0,
		tokenSize: 0,
	}
}

func (s *DebugScanner) printState(action string) {
	fmt.Printf("\n=== %s ===\n", action)
	fmt.Printf("버퍼 내용: %v (문자로: %s)\n", s.buffer[:s.tokenSize], string(s.buffer[:s.tokenSize]))
	fmt.Printf("현재 위치: %d\n", s.position)
	fmt.Printf("토큰 크기: %d\n", s.tokenSize)
	fmt.Println("==============")
}

func (s *DebugScanner) Scan() bool {
	if s.position >= s.tokenSize {
		fmt.Println("\n📖 버퍼가 비었거나 끝에 도달했습니다. 새로운 데이터를 읽어옵니다...")
		n, err := s.reader.Read(s.buffer)
		if err != nil {
			fmt.Println("📕 더 이상 읽을 데이터가 없습니다.")
			return false
		}
		s.position = 0
		s.tokenSize = n
		s.printState("새로운 데이터 읽기 완료")
	}
	return s.tokenSize > 0
}

func (s *DebugScanner) Line() string {
	var line strings.Builder
	fmt.Println("\n🔍 한 줄 읽기 시작...")

	for s.position < s.tokenSize {
		b := s.buffer[s.position]
		fmt.Printf("현재 읽는 문자: %c (위치: %d)\n", b, s.position)
		s.position++

		if b == '\n' {
			fmt.Println("👀 줄바꿈 문자를 발견했습니다!")
			break
		}
		line.WriteByte(b)
	}

	result := line.String()
	fmt.Printf("📝 읽은 결과: %s\n", result)
	return result
}

// func main() {
// 	// 테스트할 텍스트 준비
// 	content := "Hello\nWorld"
// 	fmt.Printf("📌 테스트할 텍스트: %q\n", content)

// 	scanner := NewDebugScanner(content)

// 	fmt.Println("\n🎬 스캔 시작...")
// 	lineNumber := 1

// 	for scanner.Scan() {
// 		fmt.Printf("\n=== 📄 %d번째 줄 읽기 ===\n", lineNumber)
// 		line := scanner.Line()
// 		fmt.Printf("=== ✅ %d번째 줄 결과: %q ===\n", lineNumber, line)
// 		lineNumber++
// 	}

// 	fmt.Println("\n🏁 스캔 완료")
// }
