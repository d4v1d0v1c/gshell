package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"var": VAR,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"   // 1, 2, 3, ...
	IDENT   = "IDENT" // add, x, foobar
	ASSIGN  = "="
	PLUS    = "+"

	COMA      = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"

	LBRACE = "{"
	RBRACE = "}"
	// keywords
	FUNCTION = "fn"
	LET      = "LET"
	VAR      = "VAR"
)

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
