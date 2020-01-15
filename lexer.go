package rsql

import (
	"bytes"
	"strings"

	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

// types :
const (
	Expression = iota
	String
	Or
	And
	Numeric
	Text
	Group
	Whitespace
)

type Token struct {
	Type        int
	Value       string
	Lexeme      []byte
	TC          int
	StartLine   int
	StartColumn int
	EndLine     int
	EndColumn   int
}

type defaultTokenLexer struct {
	ids   map[string]int
	lexer *lexmachine.Lexer
}

func newDefaultTokenLexer() *defaultTokenLexer {
	return &defaultTokenLexer{
		ids: map[string]int{
			"whitespace": Whitespace,
			"grouping":   Group,
			"string":     String,
			"text":       Text,
			"numeric":    Numeric,
			"and":        And,
			"or":         Or,
			"operator":   Expression,
		},
	}
}

func (l *defaultTokenLexer) addActions(lexer *lexmachine.Lexer) {
	b := new(bytes.Buffer)
	b.WriteByte('(')
	for k := range operators {
		b.WriteString(escapeStr(k))
		b.WriteByte('|')
	}
	b.Truncate(b.Len() - 1)
	b.WriteByte(')')

	lexer.Add([]byte(`\s`), l.token("whitespace"))
	lexer.Add([]byte(`\(|\)`), l.token("grouping"))
	lexer.Add([]byte(`\"(\\.|[^\"])*\"`), l.token("string"))
	lexer.Add([]byte(`\'(\\.|[^\'])*\'`), l.token("string"))
	lexer.Add([]byte(`(\,|or)`), l.token("or"))
	lexer.Add([]byte(`(\;|and)`), l.token("and"))
	lexer.Add([]byte(`(\-)?([0-9]*\.[0-9]+|[0-9]+)`), l.token("numeric"))
	lexer.Add([]byte(`[a-zA-Z0-9\_\.\%]+`), l.token("text"))
	lexer.Add(b.Bytes(), l.token("operator"))
	l.lexer = lexer
}

func (l *defaultTokenLexer) token(name string) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return &Token{
			Type:        l.ids[name],
			Value:       string(m.Bytes),
			Lexeme:      m.Bytes,
			TC:          m.TC,
			StartLine:   m.StartLine,
			StartColumn: m.StartColumn,
			EndLine:     m.EndLine,
			EndColumn:   m.EndColumn,
		}, nil
	}
}

func escapeStr(str string) string {
	length := len(str)
	blr := new(strings.Builder)
	for i := 0; i < length; i++ {
		if (str[i] >= 'a' && str[i] <= 'z') ||
			(str[i] >= 'A' && str[i] <= 'Z') ||
			(str[i] >= '0' && str[i] <= '9') {
			blr.WriteByte(str[i])
		} else {
			blr.WriteByte('\\')
			blr.WriteByte(str[i])
		}
	}
	return blr.String()
}
