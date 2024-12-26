package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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
	for p.lex.ch != '\000' {
		p.advance()
		if p.current.token_type == TOKEN_INSTRUCTION {
			fmt.Print(p.current, "\n")
			m, err := p.parse_multiply()
			if err != nil {
				continue
			}
			instructs = append(instructs, m)
		}
	}
	return instructs
}
func (p *parser) parse_multiply() (multiply, error) {
	m := multiply{}
	p.advance()
	fmt.Print(p.current, "\n")
	if p.current.token_type != TOKEN_OPENING_PAREN {
		return m, fmt.Errorf("illegal")
	}
	p.advance()
	fmt.Print(p.current, "\n")
	if p.current.token_type != TOKEN_NUMBER {
		return m, fmt.Errorf("illegal")
	}
	n, _ := strconv.ParseInt(p.current.literal, 10, 32)
	p.advance()
	fmt.Print(p.current, "\n")
	m.num1 = int(n)
	if p.current.token_type != TOKEN_COMMA {
		return m, fmt.Errorf("illegal")
	}

	p.advance()
	fmt.Print(p.current, "\n")
	if p.current.token_type != TOKEN_NUMBER {
		return m, fmt.Errorf("illegal")
	}
	n, _ = strconv.ParseInt(p.current.literal, 10, 32)
	m.num2 = int(n)
	p.advance()
	fmt.Print(p.current, "\n")
	if p.current.token_type != TOKEN_CLOSING_PAREN {
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
		if literal == "mul" {
			tok.token_type = TOKEN_INSTRUCTION
		} else {
			tok.token_type = TOKEN_ILLEGAL
		}
		tok.literal = literal
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
