// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterForstmt is called when entering the forstmt production.
	EnterForstmt(c *ForstmtContext)

	// EnterSwitchtmt is called when entering the switchtmt production.
	EnterSwitchtmt(c *SwitchtmtContext)

	// EnterCaselist is called when entering the caselist production.
	EnterCaselist(c *CaselistContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterElseiflist is called when entering the elseiflist production.
	EnterElseiflist(c *ElseiflistContext)

	// EnterElseif is called when entering the elseif production.
	EnterElseif(c *ElseifContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterDeclarationstmt is called when entering the declarationstmt production.
	EnterDeclarationstmt(c *DeclarationstmtContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterListParams is called when entering the listParams production.
	EnterListParams(c *ListParamsContext)

	// EnterListArray is called when entering the listArray production.
	EnterListArray(c *ListArrayContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitForstmt is called when exiting the forstmt production.
	ExitForstmt(c *ForstmtContext)

	// ExitSwitchtmt is called when exiting the switchtmt production.
	ExitSwitchtmt(c *SwitchtmtContext)

	// ExitCaselist is called when exiting the caselist production.
	ExitCaselist(c *CaselistContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitElseiflist is called when exiting the elseiflist production.
	ExitElseiflist(c *ElseiflistContext)

	// ExitElseif is called when exiting the elseif production.
	ExitElseif(c *ElseifContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitDeclarationstmt is called when exiting the declarationstmt production.
	ExitDeclarationstmt(c *DeclarationstmtContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitListParams is called when exiting the listParams production.
	ExitListParams(c *ListParamsContext)

	// ExitListArray is called when exiting the listArray production.
	ExitListArray(c *ListArrayContext)
}
