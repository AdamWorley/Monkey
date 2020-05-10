package token

// Type is the string representation of a token
type Type string

// Token is a unit that represent a token type and it's string literal
type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string] Type{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) Type{
	if tok, ok := keywords[ident]; ok{
		return tok
	}

	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers

	IDENT = "IDENT" //add, x, y...
	INT   = "INT"

	// Operators

	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"

	EQ = "=="
	NOT_EQ = "!="

	// Delimiters

	COMMA     = ","
	SEMICOLON = ";"
	PERIOD    = "."

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE	 = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)
