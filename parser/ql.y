%{
package parser
%}

%union{
    *Token

    // scoped control flow expressions
    *ScopeDef
    *ExprBlock
    *ConditionalExpr
    *ForExpr

    ControlFlowExpr
    Expr

    Statements []ControlFlowExpr
    Arguments []*Argument

    *TypeSpec
}

//
// Tokens generated by raw tokenizer
//

%token LEX_ERROR BLOCK_COMMENT_END

%token <Token> COMMENT
%token <Token> L_BRACE R_BRACE L_PAREN R_PAREN L_BRACKET R_BRACKET
%token <Token> ASSIGN COLON COMMA DOT STAR_STAR AT

// Terminators
%token <Token> SEMICOLON NEWLINE

// Keywords
%token <Token> LET
%token <Token> IF ELSE
%token <Token> FOR
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

%token <Token> IDENT
%token <Token> BYTE_LITERAL STRING_LITERAL
%token <Token> INT_LITERAL FLOAT_LITERAL BOOL_LITERAL

%token <Token> BYTE_TYPE FLOAT_TYPE INT_TYPE UINT_TYPE

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

%type <Token> literal conversion_type

%type <Expr> expr unit_expr composable_expr

%type <Statements> statement_list nonempty_statement_list
%type <Arguments> argument_list nonempty_argument_list

%type <ScopeDef> scope_def

%type <ExprBlock> expr_block base_expr_block
%type <ConditionalExpr> conditional_expr base_conditional_expr
%type <ForExpr> for_expr base_for_expr

%type <ControlFlowExpr> statement control_flow_expr
%type <ControlFlowExpr> assignment_expr return_expr

%type <TypeSpec> type_spec

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
    SEMICOLON {
        // do nothing
    }
    | NEWLINE {
        // do nothing
    }
    ;

statement:
    terminator {
        $$ = nil
    }
    | control_flow_expr terminator {
        $$ = $1
    }
    ;

control_flow_expr:
    NOOP {
        $$ = &Noop{
            Location: $1.Location,
            Value: $1,
        }
    }
    | expr {
        $$ = &EvalExpr{
            Location: $1.Loc(),
            Expression: $1,
        }
    }
    | assignment_expr {
        $$ = $1
    }
    | return_expr {
        $$ = $1
    }
    ;

scope_def:
    /* empty */ {
        $$ = nil
    }
    | IDENT AT {
        $$ = &ScopeDef{
            Location: $1.Location.Merge($2.Location),
            Name: $1,
            At: $2,
        }
    }
    ;

type_spec:
    BYTE_TYPE {
        // TODO
    }
    | FLOAT_TYPE {
        // TODO
    }
    | INT_TYPE {
        // TODO
    }
    | UINT_TYPE {
        // TODO
    }
    ;

expr_block:
    scope_def base_expr_block {
        $$ = $2
        if $1 != nil {
            $$.Location = $1.Location.Merge($2.Location)
            $$.ScopeDef = $1
        }
    }
    ;

base_expr_block:
    L_BRACE statement_list R_BRACE {
        $$ = &ExprBlock{
            Location: $1.Loc().Merge($3.Loc()),
            LBrace: $1,
            Statements: $2,
            RBrace: $3,
        }
    }
    ;

conditional_expr:
    scope_def base_conditional_expr {
        $$ = $2
        if $1 != nil {
            $$.Location = $1.Location.Merge($2.Location)
            $$.ScopeDef = $1
        }
    }
    ;

base_conditional_expr:
    IF composable_expr base_expr_block {
        $$ = &ConditionalExpr{
            Location: $1.Loc().Merge($3.Loc()),
            If: $1,
            Predicate: $2,
            TrueClause: $3,
        }
    }
    | IF composable_expr NEWLINE base_expr_block {
        $$ = &ConditionalExpr{
            Location: $1.Loc().Merge($4.Loc()),
            If: $1,
            Predicate: $2,
            TrueClause: $4,
        }
    }
    | IF composable_expr base_expr_block ELSE base_conditional_expr {
        $$ = &ConditionalExpr{
            Location: $1.Loc().Merge($5.Loc()),
            If: $1,
            Predicate: $2,
            TrueClause: $3,
            Else: $4,
            FalseClause: $5,
        }
    }
    | IF composable_expr NEWLINE base_expr_block ELSE base_conditional_expr {
        $$ = &ConditionalExpr{
            Location: $1.Loc().Merge($6.Loc()),
            If: $1,
            Predicate: $2,
            TrueClause: $4,
            Else: $5,
            FalseClause: $6,
        }
    }
    | IF composable_expr base_expr_block ELSE base_expr_block {
        $$ = &ConditionalExpr{
            Location: $1.Loc().Merge($5.Loc()),
            If: $1,
            Predicate: $2,
            TrueClause: $3,
            Else: $4,
            FalseClause: $5,
        }
    }
    | IF composable_expr NEWLINE base_expr_block ELSE base_expr_block {
        $$ = &ConditionalExpr{
            Location: $1.Loc().Merge($6.Loc()),
            If: $1,
            Predicate: $2,
            TrueClause: $4,
            Else: $5,
            FalseClause: $6,
        }
    }
    ;

for_expr:
    scope_def base_for_expr {
        $$ = $2
        if $1 != nil {
            $$.Location = $1.Location.Merge($2.Location)
            $$.ScopeDef = $1
        }
    }
    ;

base_for_expr:
    FOR composable_expr expr_block {
        $$ = &ForExpr{
            Location: $1.Location.Merge($3.Location),
            For: $1,
            Predicate: $2,
            Body: $3,
        }
    }
    | FOR composable_expr NEWLINE expr_block {
        $$ = &ForExpr{
            Location: $1.Location.Merge($4.Location),
            For: $1,
            Predicate: $2,
            Body: $4,
        }
    }
    ;

assignment_expr:
    LET IDENT ASSIGN expr {
        $$ = &AssignExpr{
            Location: $1.Loc().Merge($4.Loc()),
            Let: $1,
            Name: $2,
            Assign: $3,
            Expression: $4,
        }
    }
    |
    LET IDENT type_spec ASSIGN expr {
        $$ = &AssignExpr{
            Location: $1.Loc().Merge($5.Loc()),
            Let: $1,
            Name: $2,
            TypeSpec: $3,
            Assign: $4,
            Expression: $5,
        }
    }
    |
    IDENT ASSIGN expr {
        $$ = &AssignExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Name: $1,
            Assign: $2,
            Expression: $3,
        }
    }
    ;

return_expr:
    RETURN expr {
        $$ = &ReturnExpr{
            Location: $1.Loc().Merge($2.Loc()),
            Return: $1,
            Expression: $2,
        }
    }
    | RETURN AT IDENT expr {
        $$ = &ReturnExpr{
            Location: $1.Loc().Merge($4.Loc()),
            Return: $1,
            At: $2,
            Label: $3,
            Expression: $4,
        }
    }
    ;

expr:
    composable_expr {
        $$ = $1
    }
    ;

literal:
    BYTE_LITERAL {
        $$ = $1
    }
    | STRING_LITERAL {
        $$ = $1
    }
    | INT_LITERAL {
        $$ = $1
    }
    | FLOAT_LITERAL {
        $$ = $1
    }
    | BOOL_LITERAL {
        $$ = $1
    }
    ;

conversion_type:
    BYTE_TYPE {
        $$ = $1
    }
    | INT_TYPE {
        $$ = $1
    }
    | UINT_TYPE {
        $$ = $1
    }
    | FLOAT_TYPE {
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
    | literal {
        $$ = &Literal{
            Location: $1.Location,
            Value: $1,
        }
    }
    | conversion_type L_PAREN composable_expr R_PAREN {
        $$ = &TypeConversion{
            Location: $1.Loc().Merge($4.Loc()),
            Type: $1,
            LParen: $2,
            Expression: $3,
            RParen: $4,
        }
    }
    | expr_block {
        $$ = $1
    }
    | conditional_expr {
        $$ = $1
    }
    | for_expr {
        $$ = $1
    }
    | unit_expr DOT IDENT {
        $$ = &Accessor{
            Location: $1.Loc().Merge($3.Loc()),
            PrimaryExpr: $1,
            Dot: $2,
            Name: $3,
        }
    }
    | unit_expr L_PAREN argument_list R_PAREN {
        $$ = &Invocation{
            Location: $1.Loc().Merge($4.Loc()),
            Expression: $1,
            LParen: $2,
            Arguments: $3,
            RParen: $4,
        }
    }
    ;

argument_list:
    { // empty
        $$ = nil
    }
    | nonempty_argument_list {
        $$ = $1
    }
    ;

nonempty_argument_list:
    composable_expr {
        $$ = []*Argument{
            &Argument{
                Location: $1.Loc(),
                Expression: $1,
            },
        }
    }
    | nonempty_argument_list COMMA composable_expr {
        $1[len($1)-1].Location = $1[len($1)-1].Location.Merge($2.Loc())
        $1[len($1)-1].Comma = $2
        $$ = append($1,
            &Argument{
                Location: $3.Loc(),
                Expression: $3,
            })
    }
    ;

// TODO(patrick): maybe in/like expr
composable_expr:
    unit_expr {
        $$ = $1
    }
    | composable_expr OR composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr AND composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr LT composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr GT composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr EQ composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr NE composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr LE composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr GE composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr BITWISE_OR composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr BITWISE_AND composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr XOR composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr L_SHIFT composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr R_SHIFT composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr ADD composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr SUB composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr MUL composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr DIV composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | composable_expr MOD composable_expr {
        $$ = &BinaryExpr{
            Location: $1.Loc().Merge($3.Loc()),
            Left: $1,
            Op: $2,
            Right: $3,
        }
    }
    | SUB composable_expr %prec UNARY {
        $$ = &UnaryExpr{
            Location: $1.Loc().Merge($2.Loc()),
            Op: $1,
            Expression: $2,
        }
    }
    | NOT composable_expr {
        $$ = &UnaryExpr{
            Location: $1.Loc().Merge($2.Loc()),
            Op: $1,
            Expression: $2,
        }
    }
    ;

%%
