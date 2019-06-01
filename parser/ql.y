%{
package parser

type strLoc struct {
    str string
    loc Location
}

%}

%union{
    Location

    strLoc

    *AssignExpr
    Expr

    Nodes []Node
}

%token LEX_ERROR
%token ASSIGN L_BRACE R_BRACE L_PAREN R_PAREN SEMICOLON NEWLINE COMMA DOT

%left OR
%left AND
%right NOT
%left LT GT EQ NE LE GE
%left BITWISE_OR
%left BITWISE_AND
%left SHIFT_LEFT SHIFT_RIGHT
%left ADD SUB
%left MUL DIV MOD
%right UNARY

%token <strLoc> IDENTIFIER

// comments that are not directly next to any real tokens
%token <strLoc> COMMENT_GROUP

%token LET

%type <AssignExpr> assignment_expr
%type <Expr> expr unit_expr

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
        $$ = &Identifier {
            Value: $1.str,
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
    }
    | composable_expr AND composable_expr {
    }
    | composable_expr LT composable_expr {
    }
    | composable_expr GT composable_expr {
    }
    | composable_expr EQ composable_expr {
    }
    | composable_expr NE composable_expr {
    }
    | composable_expr LE composable_expr {
    }
    | composable_expr GE composable_expr {
    }
    | composable_expr BITWISE_OR composable_expr {
    }
    | composable_expr BITWISE_AND composable_expr {
    }
    | composable_expr SHIFT_LEFT composable_expr {
    }
    | composable_expr SHIFT_RIGHT composable_expr {
    }
    | composable_expr ADD composable_expr {
    }
    | composable_expr SUB composable_expr {
    }
    | composable_expr MUL composable_expr {
    }
    | composable_expr DIV composable_expr {
    }
    | composable_expr MOD composable_expr {
    }
    | SUB composable_expr %prec UNARY {
    }
    | NOT composable_expr {
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
            Name: $2.str,
            Expression: $4,
        }
    }
    ;

%%
