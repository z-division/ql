%{
package parser
%}

// NOTE: We'll take advantage of the fact that goyacc generates qlSymType as
// a struct rather than an union, and populate debugging information such as
// Location into the struct.
%union{
    loc Location

    strVal string
}

%token LEX_ERROR
%token <strVal> SELECT FROM WHERE

%%

statement_list : statement /* empty */
    | statement_list statement
    ;

statement : query_statement
    ;

query_statement : SELECT FROM WHERE
    ;

%%

