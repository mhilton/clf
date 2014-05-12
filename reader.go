package clf

import (
	"bufio"
	"io"
)

type Reader struct {
	br  *bufio.Reader
	err error
}

// NewReader returns a new Reader
func NewReader(r io.Reader) *Reader {
	return &Reader{br: bufio.NewReader(r)}
}

// Read reads the next log from the input.
func (r *Reader) Read() (Log, error) {
	var l Log
	var err error

	l.Raw, err = r.br.ReadString('\n')
	if err != nil {
		return l, err
	}

	scan(&l)

	return l, nil
}

type state int

const (
	start state = iota
	plain
	quote
	bracket
)

func scan(l *Log) {
	var state state
	var f int

	for i, c := range l.Raw {
		switch state {
		case start:
			switch c {
			case ' ':
				continue
			case '"':
				f = i + 1
				state = quote
			case '[':
				f = i + 1
				state = bracket
			default:
				f = i
				state = plain
			}

		case plain:
			switch c {
			case ' ', '\n':
				l.Fields = append(l.Fields, l.Raw[f:i])
				state = start
			default:
				continue
			}

		case quote:
			switch c {
			case '"', '\n':
				l.Fields = append(l.Fields, l.Raw[f:i])
				state = start
			default:
				continue
			}

		case bracket:
			switch c {
			case ']', '\n':
				l.Fields = append(l.Fields, l.Raw[f:i])
				state = start
			default:
				continue
			}
		}
	}

	if state != start {
		l.Fields = append(l.Fields, l.Raw[f:len(l.Raw)])
	}
}
