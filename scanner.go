package main

import (
	"io"
	"strings"
)

type Scanner struct {
	src        []byte
	srcIndex   int
	tokens     []Token
	tokenIndex int
	fullScaned bool
}

func (s *Scanner) Init() {
	s.tokens = []Token{}
	s.tokenIndex = -1
	s.srcIndex = -1
}

// TODO move to parser?
func (s *Scanner) peek() (Token, int) {
	index := s.srcIndex
	tok, pos := s.next()
	s.srcIndex = index
	return tok, pos
}

// For comment
func (s *Scanner) nextLine() Token {
	text := ""
	for {
		if p, err := s.PeepCh(); err == io.EOF || p == "\n" {
			break
		}
		ch, _ := s.nextCh()
		text += ch
	}
	return Token{text, CommentType}
}

func (s *Scanner) next() (Token, int) {
	pos := s.srcIndex
	ch, err := s.nextCh()
	if err != nil && err == io.EOF {
		s.fullScaned = true
		return Token{"", EOFType}, pos
	}

	for IsSpace(ch) || ch == "\n" {
		ch, err = s.nextCh()
		if err != nil && err != io.EOF {
			panic(err)
		}
	}

	text := ""
	isNum := false
	switch Kind(ch) {
	case LETTER:
		for ch != Space && err != io.EOF {
			text += ch
			ch, err = s.nextCh()
		}
	case DIGIT:
		isNum = true
		for ch != Space && err != io.EOF {
			text += ch
			if Kind(ch) == LETTER {
				panic("Invalid variable name: " + text)
			}
			ch, err = s.nextCh()
		}
	case LBRACE:
		return ToToken(ch, isNum), pos
	case RBRACE:
		return ToToken(ch, isNum), pos
	default: // Operator
		for ch != Space && err != io.EOF {
			text += ch
			if p, _ := s.PeepCh(); Kind(p) != OTHER {
				break
			}
			ch, err = s.nextCh()
		}
	}

	if err != nil && err == io.EOF {
		// if err and not EOF, increase err count
		s.fullScaned = true
	}

	return ToToken(text, isNum), pos
}

func (s *Scanner) nextCh() (string, error) {
	s.srcIndex += 1
	if s.srcIndex >= len(s.src) {
		return "", io.EOF
	}
	return string(s.src[s.srcIndex]), nil
}

func (s *Scanner) PeepCh() (string, error) {
	if s.srcIndex+1 >= len(s.src) {
		return "", io.EOF
	}
	return string(s.src[s.srcIndex+1]), nil
}

func (s *Scanner) undoCh() {
	if 0 <= s.srcIndex-1 {
		s.srcIndex -= 1
	}
}

func (s *Scanner) push(token Token) {
	s.tokens = append(s.tokens, token)
}

func ToToken(token string, num bool) Token {
	if num {
		if strings.Contains(token, ".") {
			return Token{token, DoubleType}
		} else {
			return Token{token, IntType}
		}
	}

	kind := KeywordType(token)
	return Token{token, kind}
}

func IsSpace(ch string) bool {
	if ch == Space {
		return true
	}
	return false
}