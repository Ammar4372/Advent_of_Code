package main

import (
	"fmt"
	"io"
	"os"
)

const TOKEN_INSTRUCTION = "instruction"
const TOKEN_COMMA = ","
const TOKEN_OPENING_PAREN = "("
const TOKEN_CLOSING_PAREN = ")"
const TOKEN_NUMBER = "number"
const TOKEN_ILLEGAL = "illegal"

type lexer struct {
	input         []byte
	input_length  int
	position      int
	read_position int
	ch            byte
}

type token struct {
	literal    string
	token_type string
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	l := lexerInit(buf)
	fmt.Printf("%s\n", l.input)
	tokens := []token{}
	for {
		tokens = append(tokens, l.nextToken())
		if l.ch == '\000' {
			break
		}
	}
	fmt.Printf("%v\n", tokens)
}
func (l *lexer) next() {
	if l.read_position >= l.input_length {
		l.ch = '\000'
		return
	}
	l.ch = l.input[l.read_position]
	l.position = l.read_position
	l.read_position += 1
}
func lexerInit(input []byte) *lexer {
	l := &lexer{
		input:         input,
		input_length:  len(input),
		position:      0,
		read_position: 0,
		ch:            0,
	}
	l.next()
	return l
}
func (l *lexer) readInt() string {
	position := l.position
	for '0' <= l.ch && '9' >= l.ch {
		l.next()
	}
	return string(l.input[position:l.position])
}
func (l *lexer) readInstruct() string {
	position := l.position
	for 'u' == l.ch || 'l' == l.ch || l.ch == 'm' {
		l.next()
	}
	return string(l.input[position:l.position])
}
func (l *lexer) nextToken() token {
	literal := ""
	tok := token{}
	switch {
	case l.ch == 'm':
		literal = l.readInstruct()
		tok.literal = literal
		tok.token_type = TOKEN_INSTRUCTION
		return tok
	case l.ch >= '0' && l.ch <= '9':
		literal = l.readInt()
		tok.literal = literal
		tok.token_type = TOKEN_NUMBER
		return tok
	case l.ch == '(':
		tok.token_type = TOKEN_OPENING_PAREN
		break
	case l.ch == ')':
		tok.token_type = TOKEN_CLOSING_PAREN
		break
	case l.ch == ',':
		tok.token_type = TOKEN_COMMA
		break
	default:
		tok.literal = string(l.ch)
		tok.token_type = TOKEN_ILLEGAL
		break
	}
	l.next()
	return tok
}
