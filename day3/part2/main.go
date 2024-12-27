package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const MULTIPLY = "instruction"
const COMMA = ","
const OPENING_PAREN = "("
const CLOSING_PAREN = ")"
const NUMBER = "number"
const DO = "do"
const DONT = "don't"
const ILLEGAL = "illegal"

type lexer struct {
	input         []byte
	input_length  int
	position      int
	read_position int
	ch            byte
}

type parser struct {
	lex     *lexer
	current token
}

type multiply struct {
	num1 int
	num2 int
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
	p := parser{
		lex: l,
	}
	instructions := p.parse_instructions()
	sum := 0
	for _, v := range instructions {
		sum += (v.num1 * v.num2)
	}
	fmt.Printf("Sum: %d\n", sum)
}

func (p *parser) parse_instructions() []multiply {
	instructs := []multiply{}
	do := true
	for p.lex.ch != '\000' {
		p.advance()
		if p.current.token_type == DO {
			if err := p.parse_do_or_dont(); err == nil {
				do = true
			}
		}
		if p.current.token_type == DONT {
			if err := p.parse_do_or_dont(); err == nil {
				do = false
			}
		}
		if p.current.token_type == MULTIPLY && do {
			m, err := p.parse_multiply()
			if err != nil {
				continue
			}
			instructs = append(instructs, m)
		}
	}
	return instructs
}
func (p *parser) parse_do_or_dont() error {
	p.advance()
	if p.current.token_type != OPENING_PAREN {
		return fmt.Errorf("illegal")
	}
	p.advance()
	if p.current.token_type != CLOSING_PAREN {
		return fmt.Errorf("illegal")
	}
	return nil
}
func (p *parser) parse_multiply() (multiply, error) {
	m := multiply{}
	p.advance()
	if p.current.token_type != OPENING_PAREN {
		return m, fmt.Errorf("illegal")
	}
	p.advance()
	if p.current.token_type != NUMBER {
		return m, fmt.Errorf("illegal")
	}
	n, _ := strconv.ParseInt(p.current.literal, 10, 32)
	p.advance()
	m.num1 = int(n)
	if p.current.token_type != COMMA {
		return m, fmt.Errorf("illegal")
	}

	p.advance()
	if p.current.token_type != NUMBER {
		return m, fmt.Errorf("illegal")
	}
	n, _ = strconv.ParseInt(p.current.literal, 10, 32)
	m.num2 = int(n)
	p.advance()
	if p.current.token_type != CLOSING_PAREN {
		return m, fmt.Errorf("illegal")
	}
	return m, nil
}

func (p *parser) advance() {
	p.current = p.lex.nextToken()
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

func (l *lexer) parseInt() string {
	position := l.position
	for '0' <= l.ch && '9' >= l.ch {
		l.next()
	}
	return string(l.input[position:l.position])
}

func (l *lexer) parseMulInstruction() string {
	position := l.position
	for l.ch == 'u' || l.ch == 'l' || l.ch == 'm' {
		l.next()
	}
	return string(l.input[position:l.position])
}
func (l *lexer) parseDoInstructions() string {
	position := l.position
	for l.ch == 'd' || l.ch == 'o' || l.ch == 'n' || l.ch == '\'' || l.ch == 't' {
		l.next()
	}
	return string(l.input[position:l.position])
}

func (l *lexer) nextToken() token {
	literal := ""
	tok := token{}
	switch {
	case l.ch == 'm':
		literal = l.parseMulInstruction()
		if literal == "mul" {
			tok.token_type = MULTIPLY
		} else {
			tok.token_type = ILLEGAL
		}
		return tok
	case l.ch == 'd':
		literal = l.parseDoInstructions()
		if literal == "do" {
			tok.token_type = DO
		} else if literal == "don't" {
			tok.token_type = DONT
		} else {
			tok.token_type = ILLEGAL
		}
		return tok
	case l.ch >= '0' && l.ch <= '9':
		literal = l.parseInt()
		tok.literal = literal
		tok.token_type = NUMBER
		return tok
	case l.ch == '(':
		tok.token_type = OPENING_PAREN
		break
	case l.ch == ')':
		tok.token_type = CLOSING_PAREN
		break
	case l.ch == ',':
		tok.token_type = COMMA
		break
	default:
		tok.literal = string(l.ch)
		tok.token_type = ILLEGAL
		break
	}
	l.next()
	return tok
}
