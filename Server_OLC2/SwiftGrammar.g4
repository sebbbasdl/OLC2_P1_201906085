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
| appendtmt { $inst = $appendtmt.append}
| removelastmt { $inst = $removelastmt.removelast }
| removetmt { $inst = $removetmt.remove }
| function { $inst = $function.fun }
| structCreation { $inst = $structCreation.dec }
;


function returns [ interfaces.Instruction fun ]
: FUNC ID PARIZQ listParamsFunc PARDER LLAVEIZQ block LLAVEDER
{
    $fun = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.lpf, environment.NULL, $block.blk)
}
| FUNC ID PARIZQ listParamsFunc PARDER ARROW1 types LLAVEIZQ block LLAVEDER
{
    $fun = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.lpf, $types.ty, $block.blk)
}
;

listParamsFunc returns[[]interface{} lpf]
: list=listParamsFunc COMA ID D_PTS types {
    var arr []interface{}
    newParam := instructions.NewParamsDeclaration($ID.line, $ID.pos, $ID.text, $types.ty)
    arr = append($list.lpf, newParam)
    $lpf = arr
    }
| ID D_PTS types {
    $lpf = []interface{}{}
    newParam := instructions.NewParamsDeclaration($ID.line, $ID.pos, $ID.text, $types.ty)
    $lpf = append($lpf, newParam)
    }
|  { $lpf = []interface{}{} }
;

structCreation returns[interfaces.Instruction dec]
: STRUCT ID LLAVEIZQ listStructDec LLAVEDER { $dec = instructions.NewStruct($STRUCT.line, $STRUCT.pos, $ID.text, $listStructDec.l) }
;

listStructDec returns[[]interface{} l]
: list=listStructDec COMA VAR ID D_PTS types {
                                                var arr []interface{}
                                                newParams := environment.NewStructType($ID.text, $types.ty)
                                                arr = append($list.l, newParams)
                                                $l = arr
                                            }
| VAR ID D_PTS types {
                        var arr []interface{}
                        newParams := environment.NewStructType($ID.text, $types.ty)
                        arr = append(arr, newParams)
                        $l = arr
                    }
|  { $l = []interface{}{} }
;

removetmt returns [interfaces.Instruction remove]
: ID PUNTO REMOVE PARIZQ expr  PARDER { $remove = instructions.NewRemove($ID.line, $ID.pos, $ID.text, $expr.e) }
;

removelastmt returns [interfaces.Instruction removelast]
: ID PUNTO REMOVELAST PARIZQ  PARDER { $removelast = instructions.NewRemoveLast($ID.line, $ID.pos, $ID.text) }
;

appendtmt returns [interfaces.Instruction append]
:ID PUNTO APPEND PARIZQ expr PARDER { $append = instructions.NewAppend($ID.line, $ID.pos, $ID.text, $expr.e) }
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
: VAR ID D_PTS types IG expr {$dec = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, $types.ty, $expr.e, false);}
| VAR ID  IG expr {$dec = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, environment.NULL, $expr.e, false);}
| LET ID D_PTS types IG expr {$dec = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, $expr.e, true);}
| VAR ID IG expr  { $dec = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, environment.STRUCT , $expr.e, false) }
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
: op=SUB right=expr { $e = expressions.NewOperation($op.GetLine(), $op.GetColumn(), $right.e, "neg", $right.e) }
| left=expr op=(MUL|DIV|MODULO) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB|COMA) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| op=NOT right=expr { $e = expressions.NewOperation($op.GetLine(), $op.GetColumn(),$right.e , $op.text, $right.e) }
| callFunction { $e = $callFunction.cf }
| ID PARIZQ listStructExp PARDER { $e = expressions.NewStructExp($ID.line, $ID.pos, $ID.text, $listStructExp.l ) }
| PARIZQ expr PARDER { $e = $expr.e }
| list=listArray { $e = $list.p}
| CORIZQ listParams CORDER { $e = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, $listParams.l) }
| INT PARIZQ expr PARDER { $e = expressions.NewConversion($INT.GetLine(),$INT.GetColumn(),$expr.e, environment.INTEGER) }
| STR PARIZQ expr PARDER { $e = expressions.NewConversion($STR.GetLine(),$STR.GetColumn(),$expr.e, environment.STRING) }
| FLOAT PARIZQ expr PARDER { $e = expressions.NewConversion($FLOAT.GetLine(),$FLOAT.GetColumn(),$expr.e, environment.FLOAT) }
| ID PUNTO COUNT PARIZQ  PARDER { $e = expressions.NewCount($ID.line, $ID.pos, $ID.text) }
| ID PUNTO ISEMPTY PARIZQ  PARDER { $e = expressions.NewIsEmpty($ID.line, $ID.pos, $ID.text) }
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
| list = listArray PUNTO ID { $p = expressions.NewStructAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $ID.text)  }
| ID { $p = expressions.NewCallVar($ID.line, $ID.pos, $ID.text)}
;

callFunction returns[interfaces.Expression cf]
: ID PARIZQ listParamsCall PARDER { $cf = expressions.NewCallExp($ID.line, $ID.pos, $ID.text, $listParamsCall.l) }
;

listParamsCall returns[[]interface{} l]
: list=listParamsCall COMA expr {
                                    var arr []interface{}
                                    arr = append($list.l, $expr.e)
                                    $l = arr
                                }
| expr  {
            $l = []interface{}{}
            $l = append($l, $expr.e)
        }
|   {
        $l = []interface{}{}
    }
;

listStructExp returns[[]interface{} l]
: list=listStructExp COMA ID D_PTS expr {
                                            var arr []interface{}
                                            StrExp := environment.NewStructContent($ID.text, $expr.e)
                                            arr = append($list.l, StrExp)
                                            $l = arr
                                        }
| ID D_PTS expr{
                    var arr []interface{}
                    StrExp := environment.NewStructContent($ID.text, $expr.e)
                    arr = append(arr, StrExp)
                    $l = arr
                }
|   {
        $l = []interface{}{}
    }
;