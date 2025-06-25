package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"const":  CONST,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"ret":    RET,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"   // 1, 2, 3, ...
	IDENT   = "IDENT" // add, x, foobar
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	BANG    = "!"
	ASTERIX = "*"
	SLASH   = "/"

	LT = "<"
	GT = ">"

	COMA      = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"

	LBRACE = "{"
	RBRACE = "}"
	// keywords
	FUNCTION = "fn"
	LET      = "LET"
	CONST    = "CONST"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	RET      = "RET"
	EQ       = "=="
	NOT_EQ   = "!="

	COMMENT       = "//"
	COMMENT_START = "/*"
	COMMENT_END   = "*/"
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
