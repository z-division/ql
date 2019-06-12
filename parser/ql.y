%{
package parser
%}

%union{
    *Token

    *AssignExpr
    Expr

    Nodes []Node
}

%token LEX_ERROR
%token <Token> L_BRACE R_BRACE L_PAREN R_PAREN L_BRACKET R_BRACKET
%token <Token> ASSIGN COLON SEMICOLON NEWLINE COMMA DOT STAR_STAR

%token <Token> LET

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

%token <Token> IDENTIFIER CHARACTER STRING

// comments that are not directly next to any real tokens
%token <Token> COMMENT_GROUP

%type <AssignExpr> assignment_expr
%type <Expr> expr unit_expr composable_expr

// TODO(patrick): actual declarations
%type <Expr> declaration
%type <Nodes> declarations

%%

start:
    declarations {
        qllex.(*parseContext).setParsed($1)
    }
    ;

// TODO(patrick): handle newlines correctly, maybe by the tokenizer?
declarations:
    declaration {
        $$ = []Node{$1}
    }
    | declarations SEMICOLON declaration {
        $$ = append($1, $3)
    }
    | declarations NEWLINE declaration {
        $$ = append($1, $3)
    }
    ;

// TODO(patrick): actual declaration
declaration:
    assignment_expr {
        $$ = $1
    }
    | COMMENT_GROUP {
    }
    ;

expr:
    composable_expr {
    }
    ;

// TODO(patrick): literals / tuples.  maybe list slicing
unit_expr:
    IDENTIFIER {
        $$ = &Identifier{
            Location: $1.Location,
            Value: $1,
        }
    }
    | CHARACTER {
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
    | expr_block {
    }
    | nested_identifier {
    }
    | unit_expr L_PAREN argument_list R_PAREN {
    }
    | L_PAREN composable_expr R_PAREN {
    }
    ;

nested_identifier:
    unit_expr DOT IDENTIFIER {
    }
    ;

// TODO(patrick): maybe in/like expr.  conditional expr
composable_expr:
    unit_expr {
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

expr_block:
    L_BRACE expr_or_comment_list R_BRACE {
    }
    ;

// TODO(patrick): handle newlines correctly, maybe by the tokenizer?
expr_or_comment_list:
    expr_or_comment {
    }
    | expr_or_comment SEMICOLON expr_or_comment_list {
    }
    | expr_or_comment NEWLINE expr_or_comment_list {
    }
    ;

expr_or_comment:
    expr {
    }
    | COMMENT_GROUP {
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

assignment_expr:
    LET IDENTIFIER ASSIGN expr {
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

%%
