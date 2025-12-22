package lexer

type Lexer interface {
	Next() string
	Line() string
	Pos() int
}
