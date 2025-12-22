package lexer

type Tag int

const (
	TagEof Tag = 0
	TagNum Tag = iota + 256
	TagFloat
	TagLess
	TagLessOrEqual
	TagGreater
	TagGreaterOrEqual
	TagEqual
	TagNotEqual
	TagId
	TagTrue
	TagFalse
)

type Token struct {
	Tag
}

func newToken(tag Tag) Token {
	return Token{tag}
}

type Num struct {
	Token
	Value int
}

func newNum(v int) Num {
	return Num{newToken(TagNum), v}
}

type Float struct {
	Token
	Value float64
}

func newFloat(v float64) Float {
	return Float{newToken(TagFloat), v}
}

type Word struct {
	Token
	Lexeme string
}

func newWord(tag Tag, lexeme string) Word {
	return Word{newToken(tag), lexeme}
}
