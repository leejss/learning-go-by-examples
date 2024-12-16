package main

import (
	"bytes"
	"io"
)

type CustomScanner struct {
	reader    io.Reader
	buffer    []byte
	position  int
	tokenSize int
}

func NewCustomScanner(reader io.Reader) *CustomScanner {
	return &CustomScanner{
		reader:    reader,
		buffer:    make([]byte, 4096),
		position:  0,
		tokenSize: 0,
	}
}

func (s *CustomScanner) Scan() bool {
	if s.position >= s.tokenSize {
		n, err := s.reader.Read(s.buffer)

		if err != nil {
			return false
		}

		s.position = 0
		s.tokenSize = n
	}

	return s.tokenSize > 0

}

func (s *CustomScanner) Line() string {
	var line bytes.Buffer

	for s.position < s.tokenSize {
		b := s.buffer[s.position]
		s.position++

		if b == '\n' {
			break
		}

		line.WriteByte(b)
	}

	return line.String()
}
