lexer grammar SwiftLexer;



// --------------- Tokens
// types
INT:    'Int';
FLOAT:  'Float';
BOOL:   'Bool';
STR:    'String';

// reserved words
TRU:    'true';
FAL:    'false';
PRINT:  'print';
IF:     'if';
ELSE:   'else';
WHILE:  'while';
FOR:    'for';
SWITCH:   'switch';
CASE:      'case';
DEFAULT:   'default';
IN:     'in';
VAR:    'var';
LET: 'let';
BREAK: 'break';
CONTINUE: 'continue';
RETURN :  'return';
GUARD : 'guard';
APPEND: 'append';
REMOVELAST: 'removeLast';
REMOVE: 'remove';
COUNT: 'count';
ISEMPTY: 'isEmpty';
AT: 'at';
FUNC: 'func';
STRUCT: 'struct';



// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"'~["]*'"';
ID: ([a-zA-Z])[a-zA-Z0-9_]*;

// symbols

DIF:            '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:             '=';
MAY_IG:         '>=';
MEN_IG:         '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
D_PTS:          ':';
CORIZQ:         '[';
CORDER:         ']';
COMA:           ',';
ARROW1:         '->';
PUNTO:          '.';
MODULO:         '%';
GUIONB:         '_';


// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;



