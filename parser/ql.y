%{
package parser
%}

%union{
    *Token

    ControlFlowExpr
    Expr

    Statements []ControlFlowExpr
}

//
// Tokens generated by raw tokenizer
//

%token LEX_ERROR BLOCK_COMMENT_END

%token <Token> COMMENT
%token <Token> L_BRACE R_BRACE L_PAREN R_PAREN L_BRACKET R_BRACKET
%token <Token> ASSIGN COLON  COMMA DOT STAR_STAR

// Terminators
%token <Token> SEMICOLON NEWLINE

// Keywords
%token <Token> LET
%token <Token> IF ELSE
%token <Token> RETURN

%left <Token> OR
%left <Token> AND
%right <Token> NOT
%left <Token> LT GT EQ NE LE GE
%left <Token> BITWISE_OR
%left <Token> BITWISE_AND
%left <Token> XOR
%left <Token> L_SHIFT R_SHIFT
%left <Token> ADD SUB
%left <Token> MUL DIV MOD
%right UNARY

%token <Token> IDENT CHAR STRING INT FLOAT BOOL

//
// Tokens generated by comment processor
//

// For attaching comments at the end of file not associated to any non-comment
// tokens.
%token <Token> NOOP

//
// Tokens generated by terminator processor
//

//
// Types generated by parser
//

%type <Expr> expr unit_expr composable_expr


%type <Statements> statement_list nonempty_statement_list

%type <ControlFlowExpr> statement control_flow_expr
%type <ControlFlowExpr> expr_block conditional_expr assignment_expr return_expr

%%

// TODO(patrick): actual declarations
start:
    statement_list {
        nodes := make([]Node, 0, len($1))
        for _, node := range $1 {
            nodes = append(nodes, node)
        }
        qllex.(*parseContext).setParsed(nodes)
    }
    ;

statement_list:
    /* empty */ {
        $$ = nil
    }
    |
    nonempty_statement_list {
        $$ = $1
    }
    ;

nonempty_statement_list:
    statement {
        if $1 != nil {
            $$ = append($$, $1)
        }
    }
    | nonempty_statement_list statement {
        if $2 != nil {
            $$ = append($1, $2)
        }
    }
    ;

terminator:
    SEMICOLON
    | NEWLINE
    ;

statement:
    terminator {
        $$ = nil
    }
    | control_flow_expr terminator {
    }
    ;

control_flow_expr:
    NOOP {
        $$ = &Noop{
            Location: $1.Location,
            Value: $1,
        }
    }
    | expr_block {
        $$ = $1
    }
    | conditional_expr {
        $$ = $1
    }
    | assignment_expr {
        $$ = $1
    }
    | return_expr {
        $$ = $1
    }
    ;

// TODO(patrick): label
expr_block:
    L_BRACE statement_list R_BRACE {
        $$ = &ExprBlock{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            LBrace: $1,
            Statements: $2,
            RBrace: $3,
        }
    }
    ;

conditional_expr:
    IF composable_expr expr_block {
        $$ = &ConditionalExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            If: $1,
            Predicate: $2,
            TrueClause: $3,
        }
    }
    | IF composable_expr NEWLINE expr_block {
        $$ = &ConditionalExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $4.Loc().End,
            },
            If: $1,
            Predicate: $2,
            TrueClause: $4,
        }
    }
    | IF composable_expr expr_block ELSE conditional_expr {
        $$ = &ConditionalExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $5.Loc().End,
            },
            If: $1,
            Predicate: $2,
            TrueClause: $3,
            Else: $4,
            FalseClause: $5,
        }
    }
    | IF composable_expr NEWLINE expr_block ELSE conditional_expr {
        $$ = &ConditionalExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $6.Loc().End,
            },
            If: $1,
            Predicate: $2,
            TrueClause: $4,
            Else: $5,
            FalseClause: $6,
        }
    }
    | IF composable_expr expr_block ELSE expr_block {
        $$ = &ConditionalExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $5.Loc().End,
            },
            If: $1,
            Predicate: $2,
            TrueClause: $3,
            Else: $4,
            FalseClause: $5,
        }
    }
    | IF composable_expr NEWLINE expr_block ELSE expr_block {
        $$ = &ConditionalExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $6.Loc().End,
            },
            If: $1,
            Predicate: $2,
            TrueClause: $4,
            Else: $5,
            FalseClause: $6,
        }
    }
    ;

assignment_expr:
    LET IDENT ASSIGN expr {
        $$ = &AssignExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $4.Loc().End,
            },
            Let: $1,
            Name: $2,
            Assign: $3,
            Expression: $4,
        }
    }
    ;

// TODO(patrick): label
return_expr:
    RETURN expr {
        $$ = &ReturnExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $2.Loc().End,
            },
            Return: $1,
            Expression: $2,
        }
    }
    ;

expr:
    composable_expr {
        $$ = $1
    }
    | conditional_expr {
        $$ = $1
    }
    ;

// TODO(patrick): tuples.  maybe list slicing
unit_expr:
    IDENT {
        $$ = &Identifier{
            Location: $1.Location,
            Value: $1,
        }
    }
    | CHAR {
        $$ = &Literal{
            Location: $1.Location,
            Value: $1,
        }
    }
    | STRING {
        $$ = &Literal{
            Location: $1.Location,
            Value: $1,
        }
    }
    | INT {
        $$ = &Literal{
            Location: $1.Location,
            Value: $1,
        }
    }
    | FLOAT {
        $$ = &Literal{
            Location: $1.Location,
            Value: $1,
        }
    }
    | BOOL {
        $$ = &Literal{
            Location: $1.Location,
            Value: $1,
        }
    }
    | expr_block {
        $$ = $1
    }
    | unit_expr DOT IDENT {
        $$ = &Accessor{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            PrimaryExpr: $1,
            Dot: $2,
            Name: $3,
        }
    }
    | unit_expr L_PAREN argument_list R_PAREN {
    }
    | L_PAREN composable_expr R_PAREN {
        $$ = &EvalOrderExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            LParen: $1,
            Expression: $2,
            RParen: $3,
        }
    }
    ;

// TODO(patrick): maybe in/like expr.  conditional expr
composable_expr:
    unit_expr {
        $$ = $1
    }
    | composable_expr OR composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr AND composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr LT composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr GT composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr EQ composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr NE composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr LE composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr GE composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr BITWISE_OR composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr BITWISE_AND composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr XOR composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr L_SHIFT composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr R_SHIFT composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr ADD composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr SUB composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr MUL composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr DIV composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr MOD composable_expr {
        $$ = &BinaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $3.Loc().End,
            },
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | SUB composable_expr %prec UNARY {
        $$ = &UnaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $2.Loc().End,
            },
            Op: $1,
            Expression: $2,
        }
    }
    | NOT composable_expr {
        $$ = &UnaryExpr{
            Location: Location{
                Filename: $1.Loc().Filename,
                Start: $1.Loc().Start,
                End: $2.Loc().End,
            },
            Op: $1,
            Expression: $2,
        }
    }
    ;

argument_list:
    { // empty
    }
    | nonempty_argument_list {
    }
    ;

nonempty_argument_list:
    composable_expr {
    }
    | nonempty_argument_list COMMA composable_expr {
    }
    ;

%%
