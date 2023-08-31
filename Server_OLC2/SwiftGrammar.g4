grammar SwiftGrammar; 
// import SwiftLexer; 

options {
  tokenVocab = SwiftLexer;
}

@header {
    import "Server2/interfaces"
    import "Server2/environment"
    import "Server2/expressions"
    import "Server2/instructions"
    import "strings"
}


s returns [[]interface{} code]
: block EOF
    {   
        $code = $block.blk
    }
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }

: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;

instruction returns [interfaces.Instruction inst]
: printstmt { $inst = $printstmt.prnt}
| ifstmt { $inst = $ifstmt.ifinst }
| declarationstmt { $inst = $declarationstmt.dec }
| asignacion { $inst = $asignacion.asign}
| whilestmt { $inst = $whilestmt.whiles }
| forstmt { $inst = $forstmt.fors }
| switchtmt { $inst = $switchtmt.swtch }
| breaktmt { $inst = $breaktmt.break }
| continuetmt{ $inst = $continuetmt.continue }
| guardtmt { $inst = $guardtmt.guard }
;

guardtmt returns [interfaces.Instruction guard]
: GUARD expr ELSE LLAVEIZQ block LLAVEDER {$guard = instructions.NewGuard( $GUARD.line, $GUARD.pos, $expr.e, $block.blk )}

;
breaktmt returns [interfaces.Instruction break]
: BREAK { $break = instructions.NewBreak($BREAK.line, $BREAK.pos)}
;

continuetmt returns [interfaces.Instruction continue]
: CONTINUE { $continue = instructions.NewContinue($CONTINUE.line, $CONTINUE.pos)}
;

printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ expr PARDER { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$expr.e)}
;



forstmt returns [interfaces.Instruction fors]
: FOR ID IN e1 = expr PUNTO PUNTO PUNTO e2=expr LLAVEIZQ block LLAVEDER { $fors = instructions.NewFor($FOR.line, $FOR.pos, $ID.text, $e1.e,$e2.e,"nil", $block.blk) }
| FOR ID IN ope = (STRING|ID) LLAVEIZQ block LLAVEDER { $fors = instructions.NewFor($FOR.line, $FOR.pos, $ID.text, nil,nil, $ope.text ,$block.blk ) }
;

switchtmt returns [interfaces.Instruction swtch]
: SWITCH expr LLAVEIZQ  caselist  DEFAULT D_PTS block LLAVEDER{ $swtch = instructions.NewSwitch( $SWITCH.line, $SWITCH.pos, $expr.e,$caselist.cases, $block.blk ) }
;

caselist returns [[]interface{} cases]
: listcases= caselist case { var arr3 []interface{}
                        arr3= append($listcases.cases, $case.swtch)
                        $cases= arr3
                        }
| case                  {
                            $cases= []interface{}{}
                            $cases = append($cases, $case.swtch)
                        }
;

case returns [interfaces.Instruction swtch]
: CASE expr D_PTS block { $swtch = instructions.NewCase( $CASE.line, $CASE.pos,$expr.e, $block.blk ) }
;

ifstmt returns [interfaces.Instruction ifinst]
: IF expr LLAVEIZQ block LLAVEDER { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, nil, nil) }
| IF expr LLAVEIZQ b1=block LLAVEDER ELSE LLAVEIZQ b2=block LLAVEDER{$ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $b1.blk, nil, $b2.blk) }
| IF expr LLAVEIZQ block LLAVEDER elseiflist { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, $elseiflist.else_if_list, nil) }
| IF expr LLAVEIZQ b1=block LLAVEDER elseiflist ELSE LLAVEIZQ b2=block LLAVEDER { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $b1.blk, $elseiflist.else_if_list, $b2.blk) }
;

elseiflist returns [[]interface{} else_if_list]
: listif=elseiflist elseif {   var arr2 []interface{}
                        arr2 = append($listif.else_if_list, $elseif.ifinst)
                        $else_if_list = arr2
                    }
| elseif    {
                $else_if_list = []interface{}{}
                $else_if_list = append($else_if_list, $elseif.ifinst)

            }
;


elseif returns [interfaces.Instruction ifinst]
: ELSE IF expr LLAVEIZQ block LLAVEDER { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, nil, nil) }
;


whilestmt returns [interfaces.Instruction whiles]
: WHILE expr LLAVEIZQ block LLAVEDER { $whiles =instructions.NewWhiles($WHILE.line, $WHILE.pos, $expr.e, $block.blk) }
;

declarationstmt returns [interfaces.Instruction dec]
: VAR ID D_PTS types IG expr
    {
        $dec = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, $types.ty, $expr.e);
    }
;



asignacion returns [interfaces.Instruction asign]
: ID IG expr { $asign = instructions.NewAsignacion($ID.line ,$ID.pos ,$ID.text, $expr.e) }
;



types returns[environment.TipoExpresion ty]
: INT { $ty = environment.INTEGER }
| FLOAT { $ty = environment.FLOAT }
| STR { $ty = environment.STRING }
| BOOL { $ty = environment.BOOLEAN }
| CORIZQ types CORDER { $ty = environment.ARRAY }
| CORIZQ  CORDER { $ty = environment.ARRAY }
;

expr returns [interfaces.Expression e]
: left=expr op=(MUL|DIV) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| PARIZQ expr PARDER { $e = $expr.e }
| list=listArray { $e = $list.p}
| INT PARIZQ expr PARDER { $e = expressions.NewConversion($expr.e, environment.INTEGER) }
| STR PARIZQ expr PARDER { $e = expressions.NewConversion($expr.e, environment.STRING) }
| FLOAT PARIZQ expr PARDER { $e = expressions.NewConversion($expr.e, environment.FLOAT) }
| CORIZQ listParams CORDER { $e = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, $listParams.l) }
| NUMBER                             
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        $e = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.STRING)
    }                        
| TRU { $e = expressions.NewPrimitive($TRU.line, $TRU.pos, true, environment.BOOLEAN) }
| FAL { $e = expressions.NewPrimitive($FAL.line, $FAL.pos, false, environment.BOOLEAN) }
;

listParams returns[[]interface{} l]
: list=listParams COMA expr {
                                var arr []interface{}
                                arr = append($list.l, $expr.e)
                                $l = arr
                            }   
| expr {
            $l = []interface{}{}
            $l = append($l, $expr.e)
        }
;



listArray returns[interfaces.Expression p]
: list = listArray CORIZQ expr CORDER { $p = expressions.NewArrayAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $expr.e) }
| ID { $p = expressions.NewCallVar($ID.line, $ID.pos, $ID.text)}
;