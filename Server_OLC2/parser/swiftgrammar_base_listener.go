// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseSwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type BaseSwiftGrammarListener struct{}

var _ SwiftGrammarListener = &BaseSwiftGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSwiftGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSwiftGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSwiftGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSwiftGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseSwiftGrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseSwiftGrammarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSwiftGrammarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSwiftGrammarListener) ExitBlock(ctx *BlockContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseSwiftGrammarListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseSwiftGrammarListener) ExitInstruction(ctx *InstructionContext) {}

// EnterReturntmt is called when production returntmt is entered.
func (s *BaseSwiftGrammarListener) EnterReturntmt(ctx *ReturntmtContext) {}

// ExitReturntmt is called when production returntmt is exited.
func (s *BaseSwiftGrammarListener) ExitReturntmt(ctx *ReturntmtContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseSwiftGrammarListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseSwiftGrammarListener) ExitFunction(ctx *FunctionContext) {}

// EnterListParamsFunc is called when production listParamsFunc is entered.
func (s *BaseSwiftGrammarListener) EnterListParamsFunc(ctx *ListParamsFuncContext) {}

// ExitListParamsFunc is called when production listParamsFunc is exited.
func (s *BaseSwiftGrammarListener) ExitListParamsFunc(ctx *ListParamsFuncContext) {}

// EnterStructCreation is called when production structCreation is entered.
func (s *BaseSwiftGrammarListener) EnterStructCreation(ctx *StructCreationContext) {}

// ExitStructCreation is called when production structCreation is exited.
func (s *BaseSwiftGrammarListener) ExitStructCreation(ctx *StructCreationContext) {}

// EnterListStructDec is called when production listStructDec is entered.
func (s *BaseSwiftGrammarListener) EnterListStructDec(ctx *ListStructDecContext) {}

// ExitListStructDec is called when production listStructDec is exited.
func (s *BaseSwiftGrammarListener) ExitListStructDec(ctx *ListStructDecContext) {}

// EnterRemovetmt is called when production removetmt is entered.
func (s *BaseSwiftGrammarListener) EnterRemovetmt(ctx *RemovetmtContext) {}

// ExitRemovetmt is called when production removetmt is exited.
func (s *BaseSwiftGrammarListener) ExitRemovetmt(ctx *RemovetmtContext) {}

// EnterRemovelastmt is called when production removelastmt is entered.
func (s *BaseSwiftGrammarListener) EnterRemovelastmt(ctx *RemovelastmtContext) {}

// ExitRemovelastmt is called when production removelastmt is exited.
func (s *BaseSwiftGrammarListener) ExitRemovelastmt(ctx *RemovelastmtContext) {}

// EnterAppendtmt is called when production appendtmt is entered.
func (s *BaseSwiftGrammarListener) EnterAppendtmt(ctx *AppendtmtContext) {}

// ExitAppendtmt is called when production appendtmt is exited.
func (s *BaseSwiftGrammarListener) ExitAppendtmt(ctx *AppendtmtContext) {}

// EnterGuardtmt is called when production guardtmt is entered.
func (s *BaseSwiftGrammarListener) EnterGuardtmt(ctx *GuardtmtContext) {}

// ExitGuardtmt is called when production guardtmt is exited.
func (s *BaseSwiftGrammarListener) ExitGuardtmt(ctx *GuardtmtContext) {}

// EnterBreaktmt is called when production breaktmt is entered.
func (s *BaseSwiftGrammarListener) EnterBreaktmt(ctx *BreaktmtContext) {}

// ExitBreaktmt is called when production breaktmt is exited.
func (s *BaseSwiftGrammarListener) ExitBreaktmt(ctx *BreaktmtContext) {}

// EnterContinuetmt is called when production continuetmt is entered.
func (s *BaseSwiftGrammarListener) EnterContinuetmt(ctx *ContinuetmtContext) {}

// ExitContinuetmt is called when production continuetmt is exited.
func (s *BaseSwiftGrammarListener) ExitContinuetmt(ctx *ContinuetmtContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseSwiftGrammarListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseSwiftGrammarListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterForstmt is called when production forstmt is entered.
func (s *BaseSwiftGrammarListener) EnterForstmt(ctx *ForstmtContext) {}

// ExitForstmt is called when production forstmt is exited.
func (s *BaseSwiftGrammarListener) ExitForstmt(ctx *ForstmtContext) {}

// EnterSwitchtmt is called when production switchtmt is entered.
func (s *BaseSwiftGrammarListener) EnterSwitchtmt(ctx *SwitchtmtContext) {}

// ExitSwitchtmt is called when production switchtmt is exited.
func (s *BaseSwiftGrammarListener) ExitSwitchtmt(ctx *SwitchtmtContext) {}

// EnterCaselist is called when production caselist is entered.
func (s *BaseSwiftGrammarListener) EnterCaselist(ctx *CaselistContext) {}

// ExitCaselist is called when production caselist is exited.
func (s *BaseSwiftGrammarListener) ExitCaselist(ctx *CaselistContext) {}

// EnterCase is called when production case is entered.
func (s *BaseSwiftGrammarListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production case is exited.
func (s *BaseSwiftGrammarListener) ExitCase(ctx *CaseContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseSwiftGrammarListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseSwiftGrammarListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterElseiflist is called when production elseiflist is entered.
func (s *BaseSwiftGrammarListener) EnterElseiflist(ctx *ElseiflistContext) {}

// ExitElseiflist is called when production elseiflist is exited.
func (s *BaseSwiftGrammarListener) ExitElseiflist(ctx *ElseiflistContext) {}

// EnterElseif is called when production elseif is entered.
func (s *BaseSwiftGrammarListener) EnterElseif(ctx *ElseifContext) {}

// ExitElseif is called when production elseif is exited.
func (s *BaseSwiftGrammarListener) ExitElseif(ctx *ElseifContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseSwiftGrammarListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseSwiftGrammarListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterDeclarationstmt is called when production declarationstmt is entered.
func (s *BaseSwiftGrammarListener) EnterDeclarationstmt(ctx *DeclarationstmtContext) {}

// ExitDeclarationstmt is called when production declarationstmt is exited.
func (s *BaseSwiftGrammarListener) ExitDeclarationstmt(ctx *DeclarationstmtContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseSwiftGrammarListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseSwiftGrammarListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseSwiftGrammarListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseSwiftGrammarListener) ExitTypes(ctx *TypesContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSwiftGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSwiftGrammarListener) ExitExpr(ctx *ExprContext) {}

// EnterListParams is called when production listParams is entered.
func (s *BaseSwiftGrammarListener) EnterListParams(ctx *ListParamsContext) {}

// ExitListParams is called when production listParams is exited.
func (s *BaseSwiftGrammarListener) ExitListParams(ctx *ListParamsContext) {}

// EnterListArray is called when production listArray is entered.
func (s *BaseSwiftGrammarListener) EnterListArray(ctx *ListArrayContext) {}

// ExitListArray is called when production listArray is exited.
func (s *BaseSwiftGrammarListener) ExitListArray(ctx *ListArrayContext) {}

// EnterCallFunctionInst is called when production callFunctionInst is entered.
func (s *BaseSwiftGrammarListener) EnterCallFunctionInst(ctx *CallFunctionInstContext) {}

// ExitCallFunctionInst is called when production callFunctionInst is exited.
func (s *BaseSwiftGrammarListener) ExitCallFunctionInst(ctx *CallFunctionInstContext) {}

// EnterCallFunction is called when production callFunction is entered.
func (s *BaseSwiftGrammarListener) EnterCallFunction(ctx *CallFunctionContext) {}

// ExitCallFunction is called when production callFunction is exited.
func (s *BaseSwiftGrammarListener) ExitCallFunction(ctx *CallFunctionContext) {}

// EnterListParamsCall is called when production listParamsCall is entered.
func (s *BaseSwiftGrammarListener) EnterListParamsCall(ctx *ListParamsCallContext) {}

// ExitListParamsCall is called when production listParamsCall is exited.
func (s *BaseSwiftGrammarListener) ExitListParamsCall(ctx *ListParamsCallContext) {}

// EnterListStructExp is called when production listStructExp is entered.
func (s *BaseSwiftGrammarListener) EnterListStructExp(ctx *ListStructExpContext) {}

// ExitListStructExp is called when production listStructExp is exited.
func (s *BaseSwiftGrammarListener) ExitListStructExp(ctx *ListStructExpContext) {}
