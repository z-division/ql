// Code generated by goyacc -o ql.go -v ql.output -p ql ql.y. DO NOT EDIT.

//line ql.y:2
package parser

import __yyfmt__ "fmt"

//line ql.y:2

//line ql.y:5
type qlSymType struct {
	yys int
	*Token

	ControlFlowExpr
	Expr

	Statements []ControlFlowExpr
	Arguments  []*Argument
}

const LEX_ERROR = 57346
const BLOCK_COMMENT_END = 57347
const COMMENT = 57348
const L_BRACE = 57349
const R_BRACE = 57350
const L_PAREN = 57351
const R_PAREN = 57352
const L_BRACKET = 57353
const R_BRACKET = 57354
const ASSIGN = 57355
const COLON = 57356
const COMMA = 57357
const DOT = 57358
const STAR_STAR = 57359
const AT = 57360
const SEMICOLON = 57361
const NEWLINE = 57362
const LET = 57363
const IF = 57364
const ELSE = 57365
const RETURN = 57366
const OR = 57367
const AND = 57368
const NOT = 57369
const LT = 57370
const GT = 57371
const EQ = 57372
const NE = 57373
const LE = 57374
const GE = 57375
const BITWISE_OR = 57376
const BITWISE_AND = 57377
const XOR = 57378
const L_SHIFT = 57379
const R_SHIFT = 57380
const ADD = 57381
const SUB = 57382
const MUL = 57383
const DIV = 57384
const MOD = 57385
const UNARY = 57386
const IDENT = 57387
const CHAR_LITERAL = 57388
const STRING_LITERAL = 57389
const INT_LITERAL = 57390
const FLOAT_LITERAL = 57391
const BOOL_LITERAL = 57392
const NOOP = 57393

var qlToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"BLOCK_COMMENT_END",
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
	"AT",
	"SEMICOLON",
	"NEWLINE",
	"LET",
	"IF",
	"ELSE",
	"RETURN",
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
	"CHAR_LITERAL",
	"STRING_LITERAL",
	"INT_LITERAL",
	"FLOAT_LITERAL",
	"BOOL_LITERAL",
	"NOOP",
}
var qlStatenames = [...]string{}

const qlEofCode = 1
const qlErrCode = 2
const qlInitialStackSize = 16

//line ql.y:523

//line yacctab:1
var qlExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const qlPrivate = 57344

const qlLast = 244

var qlAct = [...]int{

	13, 27, 14, 46, 47, 48, 10, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 28, 98, 54,
	55, 56, 2, 50, 41, 42, 43, 44, 45, 46,
	47, 48, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 76,
	79, 58, 78, 49, 82, 85, 31, 32, 83, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 92, 28, 57, 28, 44,
	45, 46, 47, 48, 88, 89, 93, 7, 8, 51,
	28, 20, 95, 20, 97, 96, 91, 77, 19, 53,
	101, 100, 7, 8, 15, 20, 52, 16, 90, 94,
	19, 18, 99, 87, 85, 86, 21, 22, 23, 24,
	25, 26, 28, 18, 1, 12, 11, 6, 21, 22,
	23, 24, 25, 26, 9, 84, 81, 80, 3, 17,
	31, 32, 28, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 44, 45, 46, 47, 48, 28,
	85, 0, 19, 40, 41, 42, 43, 44, 45, 46,
	47, 48, 4, 0, 20, 18, 29, 0, 0, 19,
	21, 22, 23, 24, 25, 26, 5, 0, 0, 0,
	0, 0, 18, 30, 0, 0, 0, 21, 22, 23,
	24, 25, 26, 32, 0, 33, 34, 35, 36, 37,
	38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 42, 43, 44,
	45, 46, 47, 48,
}
var qlPact = [...]int{

	83, -1000, -1000, 83, -1000, -1000, 68, -1000, -1000, -1000,
	-1000, -1000, -1000, 31, -1000, 8, 71, 90, 135, 135,
	135, 59, -1000, -1000, -1000, -1000, -1000, -1000, 83, -1000,
	-1000, 135, 135, 135, 135, 135, 135, 135, 135, 135,
	135, 135, 135, 135, 135, 135, 135, 135, 135, 84,
	-1000, 7, 5, 135, -1000, 193, 115, 108, 105, 177,
	193, -27, -27, -27, -27, -27, -27, 128, -12, 200,
	40, 40, -38, -38, -1000, -1000, -1000, 152, 152, -1000,
	98, 81, 31, 52, 10, 59, 83, -1000, -1000, -1000,
	-1000, 135, 69, -5, 104, 31, -1000, -1000, 69, -1000,
	-1000, -1000,
}
var qlPgo = [...]int{

	0, 6, 139, 0, 22, 138, 137, 136, 172, 127,
	1, 2, 126, 125, 124, 186,
}
var qlR1 = [...]int{

	0, 14, 4, 4, 5, 5, 15, 15, 8, 8,
	9, 9, 9, 9, 10, 10, 11, 11, 11, 11,
	11, 11, 12, 13, 13, 1, 1, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 6, 6, 7, 7,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3,
}
var qlR2 = [...]int{

	0, 1, 0, 1, 1, 2, 1, 1, 1, 2,
	1, 1, 1, 1, 3, 5, 3, 4, 5, 6,
	5, 6, 4, 2, 4, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 3, 4, 0, 1, 1, 3,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2,
}
var qlChk = [...]int{

	-1000, -14, -4, -5, -8, -15, -9, 19, 20, 51,
	-1, -12, -13, -3, -11, 21, 24, -2, 40, 27,
	22, 45, 46, 47, 48, 49, 50, -10, 7, -8,
	-15, 25, 26, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 41, 42, 43, 45,
	-1, 18, 16, 9, -3, -3, -3, 18, -4, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, 13, 45, 45,
	-6, -7, -3, -10, 20, 45, 7, 8, -1, -1,
	10, 15, 23, -10, -4, -3, -11, -10, 23, 8,
	-11, -10,
}
var qlDef = [...]int{

	2, -2, 1, 3, 4, 8, 0, 6, 7, 10,
	11, 12, 13, 25, 26, 0, 0, 40, 0, 0,
	0, 27, 28, 29, 30, 31, 32, 33, 2, 5,
	9, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	23, 0, 0, 36, 59, 60, 0, 0, 0, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 0, 0, 34,
	0, 37, 38, 16, 0, 0, 2, 14, 22, 24,
	35, 0, 0, 17, 0, 39, 18, 20, 0, 15,
	19, 21,
}
var qlTok1 = [...]int{

	1,
}
var qlTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
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
//line ql.y:78
		{
			nodes := make([]Node, 0, len(qlDollar[1].Statements))
			for _, node := range qlDollar[1].Statements {
				nodes = append(nodes, node)
			}
			qllex.(*parseContext).setParsed(nodes)
		}
	case 2:
		qlDollar = qlS[qlpt-0 : qlpt+1]
//line ql.y:88
		{
			qlVAL.Statements = nil
		}
	case 3:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:92
		{
			qlVAL.Statements = qlDollar[1].Statements
		}
	case 4:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:98
		{
			if qlDollar[1].ControlFlowExpr != nil {
				qlVAL.Statements = append(qlVAL.Statements, qlDollar[1].ControlFlowExpr)
			}
		}
	case 5:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:103
		{
			if qlDollar[2].ControlFlowExpr != nil {
				qlVAL.Statements = append(qlDollar[1].Statements, qlDollar[2].ControlFlowExpr)
			}
		}
	case 6:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:111
		{
			// do nothing
		}
	case 7:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:114
		{
			// do nothing
		}
	case 8:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:120
		{
			qlVAL.ControlFlowExpr = nil
		}
	case 9:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:123
		{
			qlVAL.ControlFlowExpr = qlDollar[1].ControlFlowExpr
		}
	case 10:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:129
		{
			qlVAL.ControlFlowExpr = &Noop{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 11:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:135
		{
			qlVAL.ControlFlowExpr = &EvalExpr{
				Location:   qlDollar[1].Expr.Loc(),
				Expression: qlDollar[1].Expr,
			}
		}
	case 12:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:141
		{
			qlVAL.ControlFlowExpr = qlDollar[1].ControlFlowExpr
		}
	case 13:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:144
		{
			qlVAL.ControlFlowExpr = qlDollar[1].ControlFlowExpr
		}
	case 14:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:150
		{
			qlVAL.ControlFlowExpr = &ExprBlock{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[3].Token.Loc()),
				LBrace:     qlDollar[1].Token,
				Statements: qlDollar[2].Statements,
				RBrace:     qlDollar[3].Token,
			}
		}
	case 15:
		qlDollar = qlS[qlpt-5 : qlpt+1]
//line ql.y:158
		{
			qlVAL.ControlFlowExpr = &ExprBlock{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[5].Token.Loc()),
				Label:      qlDollar[1].Token,
				At:         qlDollar[2].Token,
				LBrace:     qlDollar[3].Token,
				Statements: qlDollar[4].Statements,
				RBrace:     qlDollar[5].Token,
			}
		}
	case 16:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:171
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[3].ControlFlowExpr.Loc()),
				If:         qlDollar[1].Token,
				Predicate:  qlDollar[2].Expr,
				TrueClause: qlDollar[3].ControlFlowExpr,
			}
		}
	case 17:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:179
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[4].ControlFlowExpr.Loc()),
				If:         qlDollar[1].Token,
				Predicate:  qlDollar[2].Expr,
				TrueClause: qlDollar[4].ControlFlowExpr,
			}
		}
	case 18:
		qlDollar = qlS[qlpt-5 : qlpt+1]
//line ql.y:187
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location:    qlDollar[1].Token.Loc().Merge(qlDollar[5].ControlFlowExpr.Loc()),
				If:          qlDollar[1].Token,
				Predicate:   qlDollar[2].Expr,
				TrueClause:  qlDollar[3].ControlFlowExpr,
				Else:        qlDollar[4].Token,
				FalseClause: qlDollar[5].ControlFlowExpr,
			}
		}
	case 19:
		qlDollar = qlS[qlpt-6 : qlpt+1]
//line ql.y:197
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location:    qlDollar[1].Token.Loc().Merge(qlDollar[6].ControlFlowExpr.Loc()),
				If:          qlDollar[1].Token,
				Predicate:   qlDollar[2].Expr,
				TrueClause:  qlDollar[4].ControlFlowExpr,
				Else:        qlDollar[5].Token,
				FalseClause: qlDollar[6].ControlFlowExpr,
			}
		}
	case 20:
		qlDollar = qlS[qlpt-5 : qlpt+1]
//line ql.y:207
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location:    qlDollar[1].Token.Loc().Merge(qlDollar[5].ControlFlowExpr.Loc()),
				If:          qlDollar[1].Token,
				Predicate:   qlDollar[2].Expr,
				TrueClause:  qlDollar[3].ControlFlowExpr,
				Else:        qlDollar[4].Token,
				FalseClause: qlDollar[5].ControlFlowExpr,
			}
		}
	case 21:
		qlDollar = qlS[qlpt-6 : qlpt+1]
//line ql.y:217
		{
			qlVAL.ControlFlowExpr = &ConditionalExpr{
				Location:    qlDollar[1].Token.Loc().Merge(qlDollar[6].ControlFlowExpr.Loc()),
				If:          qlDollar[1].Token,
				Predicate:   qlDollar[2].Expr,
				TrueClause:  qlDollar[4].ControlFlowExpr,
				Else:        qlDollar[5].Token,
				FalseClause: qlDollar[6].ControlFlowExpr,
			}
		}
	case 22:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:230
		{
			qlVAL.ControlFlowExpr = &AssignExpr{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[4].Expr.Loc()),
				Let:        qlDollar[1].Token,
				Name:       qlDollar[2].Token,
				Assign:     qlDollar[3].Token,
				Expression: qlDollar[4].Expr,
			}
		}
	case 23:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:242
		{
			qlVAL.ControlFlowExpr = &ReturnExpr{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[2].Expr.Loc()),
				Return:     qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
			}
		}
	case 24:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:249
		{
			qlVAL.ControlFlowExpr = &ReturnExpr{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[4].Expr.Loc()),
				Return:     qlDollar[1].Token,
				At:         qlDollar[2].Token,
				Label:      qlDollar[3].Token,
				Expression: qlDollar[4].Expr,
			}
		}
	case 25:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:261
		{
			qlVAL.Expr = qlDollar[1].Expr
		}
	case 26:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:264
		{
			qlVAL.Expr = qlDollar[1].ControlFlowExpr
		}
	case 27:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:271
		{
			qlVAL.Expr = &Identifier{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 28:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:277
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 29:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:283
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 30:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:289
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 31:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:295
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 32:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:301
		{
			qlVAL.Expr = &Literal{
				Location: qlDollar[1].Token.Location,
				Value:    qlDollar[1].Token,
			}
		}
	case 33:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:307
		{
			qlVAL.Expr = qlDollar[1].ControlFlowExpr
		}
	case 34:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:310
		{
			qlVAL.Expr = &Accessor{
				Location:    qlDollar[1].Expr.Loc().Merge(qlDollar[3].Token.Loc()),
				PrimaryExpr: qlDollar[1].Expr,
				Dot:         qlDollar[2].Token,
				Name:        qlDollar[3].Token,
			}
		}
	case 35:
		qlDollar = qlS[qlpt-4 : qlpt+1]
//line ql.y:318
		{
			qlVAL.Expr = &Invocation{
				Location:   qlDollar[1].Expr.Loc().Merge(qlDollar[4].Token.Loc()),
				Expression: qlDollar[1].Expr,
				LParen:     qlDollar[2].Token,
				Arguments:  qlDollar[3].Arguments,
				RParen:     qlDollar[4].Token,
			}
		}
	case 36:
		qlDollar = qlS[qlpt-0 : qlpt+1]
//line ql.y:330
		{ // empty
			qlVAL.Arguments = nil
		}
	case 37:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:333
		{
			qlVAL.Arguments = qlDollar[1].Arguments
		}
	case 38:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:339
		{
			qlVAL.Arguments = []*Argument{
				&Argument{
					Location:   qlDollar[1].Expr.Loc(),
					Expression: qlDollar[1].Expr,
				},
			}
		}
	case 39:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:347
		{
			qlDollar[1].Arguments[len(qlDollar[1].Arguments)-1].Location = qlDollar[1].Arguments[len(qlDollar[1].Arguments)-1].Location.Merge(qlDollar[2].Token.Loc())
			qlDollar[1].Arguments[len(qlDollar[1].Arguments)-1].Comma = qlDollar[2].Token
			qlVAL.Arguments = append(qlDollar[1].Arguments,
				&Argument{
					Location:   qlDollar[3].Expr.Loc(),
					Expression: qlDollar[3].Expr,
				})
		}
	case 40:
		qlDollar = qlS[qlpt-1 : qlpt+1]
//line ql.y:360
		{
			qlVAL.Expr = qlDollar[1].Expr
		}
	case 41:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:363
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 42:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:371
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 43:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:379
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 44:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:387
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 45:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:395
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 46:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:403
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 47:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:411
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 48:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:419
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 49:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:427
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 50:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:435
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 51:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:443
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 52:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:451
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 53:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:459
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 54:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:467
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 55:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:475
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 56:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:483
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 57:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:491
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 58:
		qlDollar = qlS[qlpt-3 : qlpt+1]
//line ql.y:499
		{
			qlVAL.Expr = &BinaryExpr{
				Location: qlDollar[1].Expr.Loc().Merge(qlDollar[3].Expr.Loc()),
				Left:     qlDollar[1].Expr,
				Op:       qlDollar[2].Token,
				Right:    qlDollar[3].Expr,
			}
		}
	case 59:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:507
		{
			qlVAL.Expr = &UnaryExpr{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[2].Expr.Loc()),
				Op:         qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
			}
		}
	case 60:
		qlDollar = qlS[qlpt-2 : qlpt+1]
//line ql.y:514
		{
			qlVAL.Expr = &UnaryExpr{
				Location:   qlDollar[1].Token.Loc().Merge(qlDollar[2].Expr.Loc()),
				Op:         qlDollar[1].Token,
				Expression: qlDollar[2].Expr,
			}
		}
	}
	goto qlstack /* stack new state and value */
}
