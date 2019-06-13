// Code generated by goyacc -o ql.go -v ql.output -p ql ql.y. DO NOT EDIT.

//line ql.y:2
package parser

import __yyfmt__ "fmt"

//line ql.y:2

//line ql.y:5
type qlSymType struct {
	yys int
	*Token

	*AssignExpr
	Expr

	Nodes []Node
}

const LEX_ERROR = 57346
const COMMENT = 57347
const L_BRACE = 57348
const R_BRACE = 57349
const L_PAREN = 57350
const R_PAREN = 57351
const L_BRACKET = 57352
const R_BRACKET = 57353
const ASSIGN = 57354
const COLON = 57355
const COMMA = 57356
const DOT = 57357
const STAR_STAR = 57358
const SEMICOLON = 57359
const NEWLINE = 57360
const LET = 57361
const OR = 57362
const AND = 57363
const NOT = 57364
const LT = 57365
const GT = 57366
const EQ = 57367
const NE = 57368
const LE = 57369
const GE = 57370
const BITWISE_OR = 57371
const BITWISE_AND = 57372
const XOR = 57373
const L_SHIFT = 57374
const R_SHIFT = 57375
const ADD = 57376
const SUB = 57377
const MUL = 57378
const DIV = 57379
const MOD = 57380
const UNARY = 57381
const IDENT = 57382
const CHAR = 57383
const STRING = 57384
const INT = 57385
const FLOAT = 57386
const BOOL = 57387
const NOOP = 57388
const TERMINATOR = 57389

var qlToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"COMMENT",
	"L_BRACE",
	"R_BRACE",
	"L_PAREN",
	"R_PAREN",
	"L_BRACKET",
	"R_BRACKET",
	"ASSIGN",
	"COLON",
	"COMMA",
	"DOT",
	"STAR_STAR",
	"SEMICOLON",
	"NEWLINE",
	"LET",
	"OR",
	"AND",
	"NOT",
	"LT",
	"GT",
	"EQ",
	"NE",
	"LE",
	"GE",
	"BITWISE_OR",
	"BITWISE_AND",
	"XOR",
	"L_SHIFT",
	"R_SHIFT",
	"ADD",
	"SUB",
	"MUL",
	"DIV",
	"MOD",
	"UNARY",
	"IDENT",
	"CHAR",
	"STRING",
	"INT",
	"FLOAT",
	"BOOL",
	"NOOP",
	"TERMINATOR",
}
var qlStatenames = [...]string{}

const qlEofCode = 1
const qlErrCode = 2
const qlInitialStackSize = 16

//line ql.y:470

//line yacctab:1
var qlExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const qlPrivate = 57344

const qlLast = 174

var qlAct = [...]int{

	14, 29, 30, 31, 32, 33, 34, 35, 36, 37,
	38, 39, 40, 41, 42, 43, 44, 47, 48, 78,
	10, 7, 9, 71, 50, 11, 49, 80, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 75, 74, 5, 4,
	38, 39, 40, 41, 42, 43, 44, 27, 28, 12,
	29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
	39, 40, 41, 42, 43, 44, 77, 79, 76, 27,
	28, 81, 29, 30, 31, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 43, 44, 26, 73,
	25, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 42, 43, 44, 17, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 46, 3, 52, 16, 8, 51,
	72, 45, 18, 19, 20, 21, 22, 23, 28, 13,
	29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
	39, 40, 41, 42, 43, 44, 37, 38, 39, 40,
	41, 42, 43, 44, 40, 41, 42, 43, 44, 24,
	1, 2, 15, 6,
}
var qlPact = [...]int{

	2, -1000, 2, -1000, -1000, -25, -27, -15, -1000, -1000,
	-1000, 47, 92, -1000, 59, 116, 92, 92, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 92, 92, 92, 92, 92,
	92, 92, 92, 92, 92, 92, 92, 92, 92, 92,
	92, 92, 92, 92, 92, -17, 92, -1000, -22, 37,
	71, 92, -28, 117, -22, 72, 72, 72, 72, 72,
	72, 85, 125, 18, 130, 130, 75, 75, -1000, -1000,
	-1000, -1000, 68, 13, 59, -1000, -1000, -1000, -1000, -1000,
	92, 59,
}
var qlPgo = [...]int{

	0, 173, 126, 172, 0, 125, 171, 170, 169, 130,
	24, 129, 99,
}
var qlR1 = [...]int{

	0, 7, 6, 6, 5, 5, 5, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 8,
	10, 10, 11, 9, 9, 12, 12, 1,
}
var qlR2 = [...]int{

	0, 1, 1, 2, 1, 2, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 4, 3, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 3,
	1, 2, 2, 0, 1, 1, 3, 4,
}
var qlChk = [...]int{

	-1000, -7, -6, -5, 47, 46, -1, 19, -5, 47,
	47, 40, 12, -2, -4, -3, 35, 22, 40, 41,
	42, 43, 44, 45, -8, 8, 6, 20, 21, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 15, 8, -4, -4, -4,
	-10, -11, -2, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, 40, -9, -12, -4, 9, 7, -10, 47, 9,
	14, -4,
}
var qlDef = [...]int{

	0, -2, 1, 2, 4, 0, 0, 0, 3, 5,
	6, 0, 0, 47, 7, 18, 0, 0, 8, 9,
	10, 11, 12, 13, 14, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 43, 37, 38, 0,
	0, 40, 0, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 15, 0, 44, 45, 17, 39, 41, 42, 16,
	0, 46,
}
var qlTok1 = [...]int{

	1,
}
var qlTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47,
}
var qlTok3 = [...]int{
	0,
}

var qlErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	qlDebug        = 0
	qlErrorVerbose = false
)

type qlLexer interface {
	Lex(lval *qlSymType) int
	Error(s string)
}

type qlParser interface {
	Parse(qlLexer) int
	Lookahead() int
}

type qlParserImpl struct {
	lval  qlSymType
	stack [qlInitialStackSize]qlSymType
	char  int
}

func (p *qlParserImpl) Lookahead() int {
	return p.char
}

func qlNewParser() qlParser {
	return &qlParserImpl{}
}

const qlFlag = -1000

func qlTokname(c int) string {
	if c >= 1 && c-1 < len(qlToknames) {
		if qlToknames[c-1] != "" {
			return qlToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func qlStatname(s int) string {
	if s >= 0 && s < len(qlStatenames) {
		if qlStatenames[s] != "" {
			return qlStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func qlErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !qlErrorVerbose {
		return "syntax error"
	}

	for _, e := range qlErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + qlTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := qlPact[state]
	for tok := TOKSTART; tok-1 < len(qlToknames); tok++ {
		if n := base + tok; n >= 0 && n < qlLast && qlChk[qlAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if qlDef[state] == -2 {
		i := 0
		for qlExca[i] != -1 || qlExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; qlExca[i] >= 0; i += 2 {
			tok := qlExca[i]
			if tok < TOKSTART || qlExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if qlExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += qlTokname(tok)
	}
	return res
}

func qllex1(lex qlLexer, lval *qlSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = qlTok1[0]
		goto out
	}
	if char < len(qlTok1) {
		token = qlTok1[char]
		goto out
	}
	if char >= qlPrivate {
		if char < qlPrivate+len(qlTok2) {
			token = qlTok2[char-qlPrivate]
			goto out
		}
	}
	for i := 0; i < len(qlTok3); i += 2 {
		token = qlTok3[i+0]
		if token == char {
			token = qlTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = qlTok2[1] /* unknown char */
	}
	if qlDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", qlTokname(token), uint(char))
	}
	return char, token
}

func qlParse(qllex qlLexer) int {
	return qlNewParser().Parse(qllex)
}

func (qlrcvr *qlParserImpl) Parse(qllex qlLexer) int {
	var qln int
	var qlVAL qlSymType
	var qlDollar []qlSymType
	_ = qlDollar // silence set and not used
	qlS := qlrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	qlstate := 0
	qlrcvr.char = -1
	qltoken := -1 // qlrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		qlstate = -1
		qlrcvr.char = -1
		qltoken = -1
	}()
	qlp := -1
	goto qlstack

ret0:
	return 0

ret1:
	return 1

qlstack:
	/* put a state and value onto the stack */
	if qlDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", qlTokname(qltoken), qlStatname(qlstate))
	}

	qlp++
	if qlp >= len(qlS) {
		nyys := make([]qlSymType, len(qlS)*2)
		copy(nyys, qlS)
		qlS = nyys
	}
	qlS[qlp] = qlVAL
	qlS[qlp].yys = qlstate

qlnewstate:
	qln = qlPact[qlstate]
	if qln <= qlFlag {
		goto qldefault /* simple state */
	}
	if qlrcvr.char < 0 {
		qlrcvr.char, qltoken = qllex1(qllex, &qlrcvr.lval)
	}
	qln += qltoken
	if qln < 0 || qln >= qlLast {
		goto qldefault
	}
	qln = qlAct[qln]
	if qlChk[qln] == qltoken { /* valid shift */
		qlrcvr.char = -1
		qltoken = -1
		qlVAL = qlrcvr.lval
		qlstate = qln
		if Errflag > 0 {
			Errflag--
		}
		goto qlstack
	}

qldefault:
	/* default state action */
	qln = qlDef[qlstate]
	if qln == -2 {
		if qlrcvr.char < 0 {
			qlrcvr.char, qltoken = qllex1(qllex, &qlrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if qlExca[xi+0] == -1 && qlExca[xi+1] == qlstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			qln = qlExca[xi+0]
			if qln < 0 || qln == qltoken {
				break
			}
		}
		qln = qlExca[xi+1]
		if qln < 0 {
			goto ret0
		}
	}
	if qln == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			qllex.Error(qlErrorMessage(qlstate, qltoken))
			Nerrs++
			if qlDebug >= 1 {
				__yyfmt__.Printf("%s", qlStatname(qlstate))
				__yyfmt__.Printf(" saw %s\n", qlTokname(qltoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for qlp >= 0 {
				qln = qlPact[qlS[qlp].yys] + qlErrCode
				if qln >= 0 && qln < qlLast {
					qlstate = qlAct[qln] /* simulate a shift of "error" */
					if qlChk[qlstate] == qlErrCode {
						goto qlstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if qlDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", qlS[qlp].yys)
				}
				qlp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if qlDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", qlTokname(qltoken))
			}
			if qltoken == qlEofCode {
				goto ret1
			}
			qlrcvr.char = -1
			qltoken = -1
			goto qlnewstate /* try again in the same state */
		}
	}

	/* reduction by production qln */
	if qlDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", qln, qlStatname(qlstate))
	}

	qlnt := qln
	qlpt := qlp
	_ = qlpt // guard against "declared and not used"

	qlp -= qlR2[qln]
	// qlp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if qlp+1 >= len(qlS) {
		nyys := make([]qlSymType, len(qlS)*2)
		copy(nyys, qlS)
		qlS = nyys
	}
	qlVAL = qlS[qlp+1]

	/* consult goto table to find next state */
	qln = qlR1[qln]
	qlg := qlPgo[qln]
	qlj := qlg + qlS[qlp].yys + 1

	if qlj >= qlLast {
		qlstate = qlAct[qlg]
	} else {
		qlstate = qlAct[qlj]
		if qlChk[qlstate] != -qln {
			qlstate = qlAct[qlg]
		}
	}
	// dummy call; replaced with literal code
	switch qlnt {

	case 1:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:71
		{
			qllex.(*parseContext).setParsed(qlDollar[1].Nodes)
		}
	case 2:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:78
		{
			qlVAL.Nodes = []Node{qlDollar[1].Expr}
		}
	case 3:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:81
		{
			if qlDollar[2].Expr != nil {
				qlVAL.Nodes = append(qlDollar[1].Nodes, qlDollar[2].Expr)
			}
		}
	case 4:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:90
		{
			qlVAL.Expr = nil
		}
	case 5:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:93
		{
			qlVAL.Expr = &Noop{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 6:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:99
		{
			qlVAL.Expr = qlDollar[1].AssignExpr
		}
	case 7:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:105
		{
		}
	case 8:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:111
		{
			qlVAL.Expr = &Identifier{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 9:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:117
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 10:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:123
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 11:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:129
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 12:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:135
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 13:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:141
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 14:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:147
		{
		}
	case 15:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:149
		{
			qlVAL.Expr = &Accessor{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Token.Loc().End,
				},
				PrimaryExpr: qlDollar[1].Expr,
				Dot:         qlDollar[2].Token,
				Name:        qlDollar[3].Token,
			}
		}
	case 16:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:161
		{
		}
	case 17:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:163
		{
			qlVAL.Expr = &EvalOrderExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[3].Token.Loc().End,
				},
				LParen:     qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
				RParen:     qlDollar[3].Token,
			}
		}
	case 18:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:179
		{
			qlVAL.Expr = qlDollar[1].Expr
		}
	case 19:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:182
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 20:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:194
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 21:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:206
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 22:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:218
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 23:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:230
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 24:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:242
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 25:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:254
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 26:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:266
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 27:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:278
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 28:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:290
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 29:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:302
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 30:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:314
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 31:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:326
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 32:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:338
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 33:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:350
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 34:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:362
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 35:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:374
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 36:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:386
		{
			qlVAL.Expr = &BinaryExpr{
				Location: Location{
					Filename: qlDollar[1].Expr.Loc().Filename,
					Start:    qlDollar[1].Expr.Loc().Start,
					End:      qlDollar[3].Expr.Loc().End,
				},
				Left:  qlDollar[1].Expr,
				Op:    qlDollar[2].Token,
				Right: qlDollar[3].Expr,
			}
		}
	case 37:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:398
		{
			qlVAL.Expr = &UnaryExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[2].Expr.Loc().End,
				},
				Op:         qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
			}
		}
	case 38:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:409
		{
			qlVAL.Expr = &UnaryExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[2].Expr.Loc().End,
				},
				Op:         qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
			}
		}
	case 39:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:423
		{
		}
	case 40:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:429
		{
		}
	case 41:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:431
		{
		}
	case 42:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:436
		{
		}
	case 43:
		qlDollar = qlS[qlpt-0 : qlpt+1]
//line ql.y:441
		{ // empty
		}
	case 44:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:443
		{
		}
	case 45:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:448
		{
		}
	case 46:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:450
		{
		}
	case 47:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:455
		{
			qlVAL.AssignExpr = &AssignExpr{
				Location: Location{
					Filename: qlDollar[1].Token.Loc().Filename,
					Start:    qlDollar[1].Token.Loc().Start,
					End:      qlDollar[4].Expr.Loc().End,
				},
				Let:        qlDollar[1].Token,
				Name:       qlDollar[2].Token,
				Assign:     qlDollar[3].Token,
				Expression: qlDollar[4].Expr,
			}
		}
	}
	goto qlstack /* stack new state and value */
}
