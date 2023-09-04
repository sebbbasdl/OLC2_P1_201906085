// Generated from c:\Users\sebas\go\bin\Server_OLC2\SwiftGrammar.g4 by ANTLR 4.9.2

    import "Server2/interfaces"
    import "Server2/environment"
    import "Server2/expressions"
    import "Server2/instructions"
    import "strings"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT=1, FLOAT=2, BOOL=3, STR=4, TRU=5, FAL=6, PRINT=7, IF=8, ELSE=9, WHILE=10, 
		FOR=11, SWITCH=12, CASE=13, DEFAULT=14, IN=15, VAR=16, LET=17, BREAK=18, 
		CONTINUE=19, RETURN=20, GUARD=21, APPEND=22, REMOVELAST=23, REMOVE=24, 
		COUNT=25, ISEMPTY=26, NUMBER=27, STRING=28, ID=29, DIF=30, IG_IG=31, NOT=32, 
		OR=33, AND=34, IG=35, MAY_IG=36, MEN_IG=37, MAYOR=38, MENOR=39, MUL=40, 
		DIV=41, ADD=42, SUB=43, PARIZQ=44, PARDER=45, LLAVEIZQ=46, LLAVEDER=47, 
		D_PTS=48, CORIZQ=49, CORDER=50, COMA=51, PUNTO=52, WHITESPACE=53, COMMENT=54, 
		LINE_COMMENT=55;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_removetmt = 3, 
		RULE_removelastmt = 4, RULE_appendtmt = 5, RULE_guardtmt = 6, RULE_breaktmt = 7, 
		RULE_continuetmt = 8, RULE_printstmt = 9, RULE_forstmt = 10, RULE_switchtmt = 11, 
		RULE_caselist = 12, RULE_case = 13, RULE_ifstmt = 14, RULE_elseiflist = 15, 
		RULE_elseif = 16, RULE_whilestmt = 17, RULE_declarationstmt = 18, RULE_asignacion = 19, 
		RULE_types = 20, RULE_expr = 21, RULE_listParams = 22, RULE_listArray = 23;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "removetmt", "removelastmt", "appendtmt", 
			"guardtmt", "breaktmt", "continuetmt", "printstmt", "forstmt", "switchtmt", 
			"caselist", "case", "ifstmt", "elseiflist", "elseif", "whilestmt", "declarationstmt", 
			"asignacion", "types", "expr", "listParams", "listArray"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'true'", "'false'", 
			"'print'", "'if'", "'else'", "'while'", "'for'", "'switch'", "'case'", 
			"'default'", "'in'", "'var'", "'let'", "'break'", "'continue'", "'return'", 
			"'guard'", "'append'", "'removeLast'", "'remove'", "'count'", "'IsEmpty'", 
			null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", 
			"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", 
			"'}'", "':'", "'['", "']'", "','", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE", 
			"WHILE", "FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "LET", "BREAK", 
			"CONTINUE", "RETURN", "GUARD", "APPEND", "REMOVELAST", "REMOVE", "COUNT", 
			"ISEMPTY", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", 
			"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", 
			"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", "CORDER", 
			"COMA", "PUNTO", "WHITESPACE", "COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SwiftGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SwiftGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SContext extends ParserRuleContext {
		public []interface{} code;
		public BlockContext block;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SwiftGrammarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			((SContext)_localctx).block = block();
			setState(49);
			match(EOF);
			   
			        _localctx.code = ((SContext)_localctx).block.blk
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public []interface{} blk;
		public InstructionContext instruction;
		public List<InstructionContext> ins = new ArrayList<InstructionContext>();
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);

		    _localctx.blk = []interface{}{}
		    var listInt []IInstructionContext
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(53); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(52);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(55); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

			        listInt = localctx.(*BlockContext).GetIns()
			        for _, e := range listInt {
			            _localctx.blk = append(_localctx.blk, e.GetInst())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstructionContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public PrintstmtContext printstmt;
		public IfstmtContext ifstmt;
		public DeclarationstmtContext declarationstmt;
		public AsignacionContext asignacion;
		public WhilestmtContext whilestmt;
		public ForstmtContext forstmt;
		public SwitchtmtContext switchtmt;
		public BreaktmtContext breaktmt;
		public ContinuetmtContext continuetmt;
		public GuardtmtContext guardtmt;
		public AppendtmtContext appendtmt;
		public RemovelastmtContext removelastmt;
		public RemovetmtContext removetmt;
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public DeclarationstmtContext declarationstmt() {
			return getRuleContext(DeclarationstmtContext.class,0);
		}
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public WhilestmtContext whilestmt() {
			return getRuleContext(WhilestmtContext.class,0);
		}
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
		}
		public SwitchtmtContext switchtmt() {
			return getRuleContext(SwitchtmtContext.class,0);
		}
		public BreaktmtContext breaktmt() {
			return getRuleContext(BreaktmtContext.class,0);
		}
		public ContinuetmtContext continuetmt() {
			return getRuleContext(ContinuetmtContext.class,0);
		}
		public GuardtmtContext guardtmt() {
			return getRuleContext(GuardtmtContext.class,0);
		}
		public AppendtmtContext appendtmt() {
			return getRuleContext(AppendtmtContext.class,0);
		}
		public RemovelastmtContext removelastmt() {
			return getRuleContext(RemovelastmtContext.class,0);
		}
		public RemovetmtContext removetmt() {
			return getRuleContext(RemovetmtContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(98);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(59);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(62);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(65);
				((InstructionContext)_localctx).declarationstmt = declarationstmt();
				 _localctx.inst = ((InstructionContext)_localctx).declarationstmt.dec 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(68);
				((InstructionContext)_localctx).asignacion = asignacion();
				 _localctx.inst = ((InstructionContext)_localctx).asignacion.asign
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(71);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whiles 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(74);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.fors 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(77);
				((InstructionContext)_localctx).switchtmt = switchtmt();
				 _localctx.inst = ((InstructionContext)_localctx).switchtmt.swtch 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(80);
				((InstructionContext)_localctx).breaktmt = breaktmt();
				 _localctx.inst = ((InstructionContext)_localctx).breaktmt.break 
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(83);
				((InstructionContext)_localctx).continuetmt = continuetmt();
				 _localctx.inst = ((InstructionContext)_localctx).continuetmt.continue 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(86);
				((InstructionContext)_localctx).guardtmt = guardtmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardtmt.guard 
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(89);
				((InstructionContext)_localctx).appendtmt = appendtmt();
				 _localctx.inst = ((InstructionContext)_localctx).appendtmt.append
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(92);
				((InstructionContext)_localctx).removelastmt = removelastmt();
				 _localctx.inst = ((InstructionContext)_localctx).removelastmt.removelast 
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(95);
				((InstructionContext)_localctx).removetmt = removetmt();
				 _localctx.inst = ((InstructionContext)_localctx).removetmt.remove 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RemovetmtContext extends ParserRuleContext {
		public interfaces.Instruction remove;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode REMOVE() { return getToken(SwiftGrammarParser.REMOVE, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public RemovetmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_removetmt; }
	}

	public final RemovetmtContext removetmt() throws RecognitionException {
		RemovetmtContext _localctx = new RemovetmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_removetmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			((RemovetmtContext)_localctx).ID = match(ID);
			setState(101);
			match(PUNTO);
			setState(102);
			match(REMOVE);
			setState(103);
			match(PARIZQ);
			setState(104);
			((RemovetmtContext)_localctx).expr = expr(0);
			setState(105);
			match(PARDER);
			 _localctx.remove = instructions.NewRemove((((RemovetmtContext)_localctx).ID!=null?((RemovetmtContext)_localctx).ID.getLine():0), (((RemovetmtContext)_localctx).ID!=null?((RemovetmtContext)_localctx).ID.getCharPositionInLine():0), (((RemovetmtContext)_localctx).ID!=null?((RemovetmtContext)_localctx).ID.getText():null), ((RemovetmtContext)_localctx).expr.e) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RemovelastmtContext extends ParserRuleContext {
		public interfaces.Instruction removelast;
		public Token ID;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode REMOVELAST() { return getToken(SwiftGrammarParser.REMOVELAST, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public RemovelastmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_removelastmt; }
	}

	public final RemovelastmtContext removelastmt() throws RecognitionException {
		RemovelastmtContext _localctx = new RemovelastmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_removelastmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			((RemovelastmtContext)_localctx).ID = match(ID);
			setState(109);
			match(PUNTO);
			setState(110);
			match(REMOVELAST);
			setState(111);
			match(PARIZQ);
			setState(112);
			match(PARDER);
			 _localctx.removelast = instructions.NewRemoveLast((((RemovelastmtContext)_localctx).ID!=null?((RemovelastmtContext)_localctx).ID.getLine():0), (((RemovelastmtContext)_localctx).ID!=null?((RemovelastmtContext)_localctx).ID.getCharPositionInLine():0), (((RemovelastmtContext)_localctx).ID!=null?((RemovelastmtContext)_localctx).ID.getText():null)) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AppendtmtContext extends ParserRuleContext {
		public interfaces.Instruction append;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode APPEND() { return getToken(SwiftGrammarParser.APPEND, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public AppendtmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_appendtmt; }
	}

	public final AppendtmtContext appendtmt() throws RecognitionException {
		AppendtmtContext _localctx = new AppendtmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_appendtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			((AppendtmtContext)_localctx).ID = match(ID);
			setState(116);
			match(PUNTO);
			setState(117);
			match(APPEND);
			setState(118);
			match(PARIZQ);
			setState(119);
			((AppendtmtContext)_localctx).expr = expr(0);
			setState(120);
			match(PARDER);
			 _localctx.append = instructions.NewAppend((((AppendtmtContext)_localctx).ID!=null?((AppendtmtContext)_localctx).ID.getLine():0), (((AppendtmtContext)_localctx).ID!=null?((AppendtmtContext)_localctx).ID.getCharPositionInLine():0), (((AppendtmtContext)_localctx).ID!=null?((AppendtmtContext)_localctx).ID.getText():null), ((AppendtmtContext)_localctx).expr.e) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GuardtmtContext extends ParserRuleContext {
		public interfaces.Instruction guard;
		public Token GUARD;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode GUARD() { return getToken(SwiftGrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public GuardtmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardtmt; }
	}

	public final GuardtmtContext guardtmt() throws RecognitionException {
		GuardtmtContext _localctx = new GuardtmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_guardtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			((GuardtmtContext)_localctx).GUARD = match(GUARD);
			setState(124);
			((GuardtmtContext)_localctx).expr = expr(0);
			setState(125);
			match(ELSE);
			setState(126);
			match(LLAVEIZQ);
			setState(127);
			((GuardtmtContext)_localctx).block = block();
			setState(128);
			match(LLAVEDER);
			_localctx.guard = instructions.NewGuard( (((GuardtmtContext)_localctx).GUARD!=null?((GuardtmtContext)_localctx).GUARD.getLine():0), (((GuardtmtContext)_localctx).GUARD!=null?((GuardtmtContext)_localctx).GUARD.getCharPositionInLine():0), ((GuardtmtContext)_localctx).expr.e, ((GuardtmtContext)_localctx).block.blk )
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BreaktmtContext extends ParserRuleContext {
		public interfaces.Instruction break;
		public Token BREAK;
		public TerminalNode BREAK() { return getToken(SwiftGrammarParser.BREAK, 0); }
		public BreaktmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breaktmt; }
	}

	public final BreaktmtContext breaktmt() throws RecognitionException {
		BreaktmtContext _localctx = new BreaktmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_breaktmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			((BreaktmtContext)_localctx).BREAK = match(BREAK);
			 _localctx.break = instructions.NewBreak((((BreaktmtContext)_localctx).BREAK!=null?((BreaktmtContext)_localctx).BREAK.getLine():0), (((BreaktmtContext)_localctx).BREAK!=null?((BreaktmtContext)_localctx).BREAK.getCharPositionInLine():0))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ContinuetmtContext extends ParserRuleContext {
		public interfaces.Instruction continue;
		public Token CONTINUE;
		public TerminalNode CONTINUE() { return getToken(SwiftGrammarParser.CONTINUE, 0); }
		public ContinuetmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continuetmt; }
	}

	public final ContinuetmtContext continuetmt() throws RecognitionException {
		ContinuetmtContext _localctx = new ContinuetmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_continuetmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			((ContinuetmtContext)_localctx).CONTINUE = match(CONTINUE);
			 _localctx.continue = instructions.NewContinue((((ContinuetmtContext)_localctx).CONTINUE!=null?((ContinuetmtContext)_localctx).CONTINUE.getLine():0), (((ContinuetmtContext)_localctx).CONTINUE!=null?((ContinuetmtContext)_localctx).CONTINUE.getCharPositionInLine():0))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrintstmtContext extends ParserRuleContext {
		public interfaces.Instruction prnt;
		public Token PRINT;
		public ExprContext expr;
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(138);
			match(PARIZQ);
			setState(139);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(140);
			match(PARDER);
			 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr.e)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ForstmtContext extends ParserRuleContext {
		public interfaces.Instruction fors;
		public Token FOR;
		public Token ID;
		public ExprContext e1;
		public ExprContext e2;
		public BlockContext block;
		public Token ope;
		public TerminalNode FOR() { return getToken(SwiftGrammarParser.FOR, 0); }
		public List<TerminalNode> ID() { return getTokens(SwiftGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(SwiftGrammarParser.ID, i);
		}
		public TerminalNode IN() { return getToken(SwiftGrammarParser.IN, 0); }
		public List<TerminalNode> PUNTO() { return getTokens(SwiftGrammarParser.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(SwiftGrammarParser.PUNTO, i);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_forstmt);
		int _la;
		try {
			setState(165);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(143);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(144);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(145);
				match(IN);
				setState(146);
				((ForstmtContext)_localctx).e1 = expr(0);
				setState(147);
				match(PUNTO);
				setState(148);
				match(PUNTO);
				setState(149);
				match(PUNTO);
				setState(150);
				((ForstmtContext)_localctx).e2 = expr(0);
				setState(151);
				match(LLAVEIZQ);
				setState(152);
				((ForstmtContext)_localctx).block = block();
				setState(153);
				match(LLAVEDER);
				 _localctx.fors = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).e1.e,((ForstmtContext)_localctx).e2.e,"nil", ((ForstmtContext)_localctx).block.blk) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(156);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(157);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(158);
				match(IN);
				setState(159);
				((ForstmtContext)_localctx).ope = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==STRING || _la==ID) ) {
					((ForstmtContext)_localctx).ope = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(160);
				match(LLAVEIZQ);
				setState(161);
				((ForstmtContext)_localctx).block = block();
				setState(162);
				match(LLAVEDER);
				 _localctx.fors = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), nil,nil, (((ForstmtContext)_localctx).ope!=null?((ForstmtContext)_localctx).ope.getText():null) ,((ForstmtContext)_localctx).block.blk ) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SwitchtmtContext extends ParserRuleContext {
		public interfaces.Instruction swtch;
		public Token SWITCH;
		public ExprContext expr;
		public CaselistContext caselist;
		public BlockContext block;
		public TerminalNode SWITCH() { return getToken(SwiftGrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public CaselistContext caselist() {
			return getRuleContext(CaselistContext.class,0);
		}
		public TerminalNode DEFAULT() { return getToken(SwiftGrammarParser.DEFAULT, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public SwitchtmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchtmt; }
	}

	public final SwitchtmtContext switchtmt() throws RecognitionException {
		SwitchtmtContext _localctx = new SwitchtmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_switchtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			((SwitchtmtContext)_localctx).SWITCH = match(SWITCH);
			setState(168);
			((SwitchtmtContext)_localctx).expr = expr(0);
			setState(169);
			match(LLAVEIZQ);
			setState(170);
			((SwitchtmtContext)_localctx).caselist = caselist(0);
			setState(171);
			match(DEFAULT);
			setState(172);
			match(D_PTS);
			setState(173);
			((SwitchtmtContext)_localctx).block = block();
			setState(174);
			match(LLAVEDER);
			 _localctx.swtch = instructions.NewSwitch( (((SwitchtmtContext)_localctx).SWITCH!=null?((SwitchtmtContext)_localctx).SWITCH.getLine():0), (((SwitchtmtContext)_localctx).SWITCH!=null?((SwitchtmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchtmtContext)_localctx).expr.e,((SwitchtmtContext)_localctx).caselist.cases, ((SwitchtmtContext)_localctx).block.blk ) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CaselistContext extends ParserRuleContext {
		public []interface{} cases;
		public CaselistContext listcases;
		public CaseContext case;
		public CaseContext case() {
			return getRuleContext(CaseContext.class,0);
		}
		public CaselistContext caselist() {
			return getRuleContext(CaselistContext.class,0);
		}
		public CaselistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caselist; }
	}

	public final CaselistContext caselist() throws RecognitionException {
		return caselist(0);
	}

	private CaselistContext caselist(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CaselistContext _localctx = new CaselistContext(_ctx, _parentState);
		CaselistContext _prevctx = _localctx;
		int _startState = 24;
		enterRecursionRule(_localctx, 24, RULE_caselist, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(178);
			((CaselistContext)_localctx).case = case();

			                            _localctx.cases= []interface{}{}
			                            _localctx.cases = append(_localctx.cases, ((CaselistContext)_localctx).case.swtch)
			                        
			}
			_ctx.stop = _input.LT(-1);
			setState(187);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new CaselistContext(_parentctx, _parentState);
					_localctx.listcases = _prevctx;
					_localctx.listcases = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_caselist);
					setState(181);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(182);
					((CaselistContext)_localctx).case = case();
					 var arr3 []interface{}
					                                  arr3= append(((CaselistContext)_localctx).listcases.cases, ((CaselistContext)_localctx).case.swtch)
					                                  _localctx.cases= arr3
					                                  
					}
					} 
				}
				setState(189);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class CaseContext extends ParserRuleContext {
		public interfaces.Instruction swtch;
		public Token CASE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode CASE() { return getToken(SwiftGrammarParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public CaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_case; }
	}

	public final CaseContext case() throws RecognitionException {
		CaseContext _localctx = new CaseContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_case);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(190);
			((CaseContext)_localctx).CASE = match(CASE);
			setState(191);
			((CaseContext)_localctx).expr = expr(0);
			setState(192);
			match(D_PTS);
			setState(193);
			((CaseContext)_localctx).block = block();
			 _localctx.swtch = instructions.NewCase( (((CaseContext)_localctx).CASE!=null?((CaseContext)_localctx).CASE.getLine():0), (((CaseContext)_localctx).CASE!=null?((CaseContext)_localctx).CASE.getCharPositionInLine():0),((CaseContext)_localctx).expr.e, ((CaseContext)_localctx).block.blk ) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfstmtContext extends ParserRuleContext {
		public interfaces.Instruction ifinst;
		public Token IF;
		public ExprContext expr;
		public BlockContext block;
		public BlockContext b1;
		public BlockContext b2;
		public ElseiflistContext elseiflist;
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<TerminalNode> LLAVEIZQ() { return getTokens(SwiftGrammarParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(SwiftGrammarParser.LLAVEIZQ, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(SwiftGrammarParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(SwiftGrammarParser.LLAVEDER, i);
		}
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public ElseiflistContext elseiflist() {
			return getRuleContext(ElseiflistContext.class,0);
		}
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_ifstmt);
		try {
			setState(234);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(196);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(197);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(198);
				match(LLAVEIZQ);
				setState(199);
				((IfstmtContext)_localctx).block = block();
				setState(200);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(203);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(204);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(205);
				match(LLAVEIZQ);
				setState(206);
				((IfstmtContext)_localctx).b1 = block();
				setState(207);
				match(LLAVEDER);
				setState(208);
				match(ELSE);
				setState(209);
				match(LLAVEIZQ);
				setState(210);
				((IfstmtContext)_localctx).b2 = block();
				setState(211);
				match(LLAVEDER);
				_localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).b1.blk, nil, ((IfstmtContext)_localctx).b2.blk) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(214);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(215);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(216);
				match(LLAVEIZQ);
				setState(217);
				((IfstmtContext)_localctx).block = block();
				setState(218);
				match(LLAVEDER);
				setState(219);
				((IfstmtContext)_localctx).elseiflist = elseiflist(0);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, ((IfstmtContext)_localctx).elseiflist.else_if_list, nil) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(222);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(223);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(224);
				match(LLAVEIZQ);
				setState(225);
				((IfstmtContext)_localctx).b1 = block();
				setState(226);
				match(LLAVEDER);
				setState(227);
				((IfstmtContext)_localctx).elseiflist = elseiflist(0);
				setState(228);
				match(ELSE);
				setState(229);
				match(LLAVEIZQ);
				setState(230);
				((IfstmtContext)_localctx).b2 = block();
				setState(231);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).b1.blk, ((IfstmtContext)_localctx).elseiflist.else_if_list, ((IfstmtContext)_localctx).b2.blk) 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ElseiflistContext extends ParserRuleContext {
		public []interface{} else_if_list;
		public ElseiflistContext listif;
		public ElseifContext elseif;
		public ElseifContext elseif() {
			return getRuleContext(ElseifContext.class,0);
		}
		public ElseiflistContext elseiflist() {
			return getRuleContext(ElseiflistContext.class,0);
		}
		public ElseiflistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseiflist; }
	}

	public final ElseiflistContext elseiflist() throws RecognitionException {
		return elseiflist(0);
	}

	private ElseiflistContext elseiflist(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ElseiflistContext _localctx = new ElseiflistContext(_ctx, _parentState);
		ElseiflistContext _prevctx = _localctx;
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_elseiflist, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(237);
			((ElseiflistContext)_localctx).elseif = elseif();

			                _localctx.else_if_list = []interface{}{}
			                _localctx.else_if_list = append(_localctx.else_if_list, ((ElseiflistContext)_localctx).elseif.ifinst)

			            
			}
			_ctx.stop = _input.LT(-1);
			setState(246);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ElseiflistContext(_parentctx, _parentState);
					_localctx.listif = _prevctx;
					_localctx.listif = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_elseiflist);
					setState(240);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(241);
					((ElseiflistContext)_localctx).elseif = elseif();
					   var arr2 []interface{}
					                                  arr2 = append(((ElseiflistContext)_localctx).listif.else_if_list, ((ElseiflistContext)_localctx).elseif.ifinst)
					                                  _localctx.else_if_list = arr2
					                              
					}
					} 
				}
				setState(248);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ElseifContext extends ParserRuleContext {
		public interfaces.Instruction ifinst;
		public Token IF;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode ELSE() { return getToken(SwiftGrammarParser.ELSE, 0); }
		public TerminalNode IF() { return getToken(SwiftGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public ElseifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseif; }
	}

	public final ElseifContext elseif() throws RecognitionException {
		ElseifContext _localctx = new ElseifContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_elseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
			match(ELSE);
			setState(250);
			((ElseifContext)_localctx).IF = match(IF);
			setState(251);
			((ElseifContext)_localctx).expr = expr(0);
			setState(252);
			match(LLAVEIZQ);
			setState(253);
			((ElseifContext)_localctx).block = block();
			setState(254);
			match(LLAVEDER);
			 _localctx.ifinst = instructions.NewIf((((ElseifContext)_localctx).IF!=null?((ElseifContext)_localctx).IF.getLine():0), (((ElseifContext)_localctx).IF!=null?((ElseifContext)_localctx).IF.getCharPositionInLine():0), ((ElseifContext)_localctx).expr.e, ((ElseifContext)_localctx).block.blk, nil, nil) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WhilestmtContext extends ParserRuleContext {
		public interfaces.Instruction whiles;
		public Token WHILE;
		public ExprContext expr;
		public BlockContext block;
		public TerminalNode WHILE() { return getToken(SwiftGrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public WhilestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilestmt; }
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(257);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(258);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(259);
			match(LLAVEIZQ);
			setState(260);
			((WhilestmtContext)_localctx).block = block();
			setState(261);
			match(LLAVEDER);
			 _localctx.whiles =instructions.NewWhiles((((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getLine():0), (((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getCharPositionInLine():0), ((WhilestmtContext)_localctx).expr.e, ((WhilestmtContext)_localctx).block.blk) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarationstmtContext extends ParserRuleContext {
		public interfaces.Instruction dec;
		public Token VAR;
		public Token ID;
		public TypesContext types;
		public ExprContext expr;
		public Token LET;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LET() { return getToken(SwiftGrammarParser.LET, 0); }
		public DeclarationstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarationstmt; }
	}

	public final DeclarationstmtContext declarationstmt() throws RecognitionException {
		DeclarationstmtContext _localctx = new DeclarationstmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_declarationstmt);
		try {
			setState(280);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
				enterOuterAlt(_localctx, 1);
				{
				setState(264);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(265);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(266);
				match(D_PTS);
				setState(267);
				((DeclarationstmtContext)_localctx).types = types();
				setState(268);
				match(IG);
				setState(269);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, false);
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(272);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(273);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(274);
				match(D_PTS);
				setState(275);
				((DeclarationstmtContext)_localctx).types = types();
				setState(276);
				match(IG);
				setState(277);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, true);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AsignacionContext extends ParserRuleContext {
		public interfaces.Instruction asign;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(282);
			((AsignacionContext)_localctx).ID = match(ID);
			setState(283);
			match(IG);
			setState(284);
			((AsignacionContext)_localctx).expr = expr(0);
			 _localctx.asign = instructions.NewAsignacion((((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getLine():0) ,(((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getCharPositionInLine():0) ,(((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getText():null), ((AsignacionContext)_localctx).expr.e) 
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypesContext extends ParserRuleContext {
		public environment.TipoExpresion ty;
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode STR() { return getToken(SwiftGrammarParser.STR, 0); }
		public TerminalNode BOOL() { return getToken(SwiftGrammarParser.BOOL, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_types);
		try {
			setState(303);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(287);
				match(INT);
				 _localctx.ty = environment.INTEGER 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(289);
				match(FLOAT);
				 _localctx.ty = environment.FLOAT 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(291);
				match(STR);
				 _localctx.ty = environment.STRING 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(293);
				match(BOOL);
				 _localctx.ty = environment.BOOLEAN 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(295);
				match(CORIZQ);
				setState(296);
				types();
				setState(297);
				match(CORDER);
				 _localctx.ty = environment.ARRAY 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(300);
				match(CORIZQ);
				setState(301);
				match(CORDER);
				 _localctx.ty = environment.ARRAY 
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext left;
		public ExprContext expr;
		public ListArrayContext list;
		public Token ID;
		public Token CORIZQ;
		public ListParamsContext listParams;
		public Token NUMBER;
		public Token STRING;
		public Token TRU;
		public Token FAL;
		public Token op;
		public ExprContext right;
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode STR() { return getToken(SwiftGrammarParser.STR, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public TerminalNode ISEMPTY() { return getToken(SwiftGrammarParser.ISEMPTY, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public TerminalNode MAY_IG() { return getToken(SwiftGrammarParser.MAY_IG, 0); }
		public TerminalNode MAYOR() { return getToken(SwiftGrammarParser.MAYOR, 0); }
		public TerminalNode MEN_IG() { return getToken(SwiftGrammarParser.MEN_IG, 0); }
		public TerminalNode MENOR() { return getToken(SwiftGrammarParser.MENOR, 0); }
		public TerminalNode IG_IG() { return getToken(SwiftGrammarParser.IG_IG, 0); }
		public TerminalNode DIF() { return getToken(SwiftGrammarParser.DIF, 0); }
		public TerminalNode AND() { return getToken(SwiftGrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(SwiftGrammarParser.OR, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 42;
		enterRecursionRule(_localctx, 42, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(357);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(306);
				match(PARIZQ);
				setState(307);
				((ExprContext)_localctx).expr = expr(0);
				setState(308);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 2:
				{
				setState(311);
				((ExprContext)_localctx).list = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).list.p
				}
				break;
			case 3:
				{
				setState(314);
				match(INT);
				setState(315);
				match(PARIZQ);
				setState(316);
				((ExprContext)_localctx).expr = expr(0);
				setState(317);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).expr.e, environment.INTEGER) 
				}
				break;
			case 4:
				{
				setState(320);
				match(STR);
				setState(321);
				match(PARIZQ);
				setState(322);
				((ExprContext)_localctx).expr = expr(0);
				setState(323);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).expr.e, environment.STRING) 
				}
				break;
			case 5:
				{
				setState(326);
				match(FLOAT);
				setState(327);
				match(PARIZQ);
				setState(328);
				((ExprContext)_localctx).expr = expr(0);
				setState(329);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).expr.e, environment.FLOAT) 
				}
				break;
			case 6:
				{
				setState(332);
				((ExprContext)_localctx).ID = match(ID);
				setState(333);
				match(PUNTO);
				setState(334);
				match(COUNT);
				setState(335);
				match(PARIZQ);
				setState(336);
				match(PARDER);
				 _localctx.e = expressions.NewCount((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 7:
				{
				setState(338);
				((ExprContext)_localctx).ID = match(ID);
				setState(339);
				match(PUNTO);
				setState(340);
				match(ISEMPTY);
				setState(341);
				match(PARIZQ);
				setState(342);
				match(PARDER);
				 _localctx.e = expressions.NewIsEmpty((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 8:
				{
				setState(344);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(345);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(346);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case 9:
				{
				setState(349);
				((ExprContext)_localctx).NUMBER = match(NUMBER);

				        if (strings.Contains((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null),".")){
				            num,err := strconv.ParseFloat((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null), 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.FLOAT)
				        }else{
				            num,err := strconv.Atoi((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null))
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.INTEGER)
				        }
				    
				}
				break;
			case 10:
				{
				setState(351);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 11:
				{
				setState(353);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 12:
				{
				setState(355);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(396);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(394);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(359);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(360);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MUL || _la==DIV) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(361);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(364);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(365);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(366);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(369);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(370);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAY_IG || _la==MAYOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(371);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(374);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(375);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MEN_IG || _la==MENOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(376);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(379);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(380);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIF || _la==IG_IG) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(381);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(384);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(385);
						((ExprContext)_localctx).op = match(AND);
						setState(386);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(389);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(390);
						((ExprContext)_localctx).op = match(OR);
						setState(391);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(398);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ListParamsContext extends ParserRuleContext {
		public []interface{} l;
		public ListParamsContext list;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public ListParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParams; }
	}

	public final ListParamsContext listParams() throws RecognitionException {
		return listParams(0);
	}

	private ListParamsContext listParams(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsContext _localctx = new ListParamsContext(_ctx, _parentState);
		ListParamsContext _prevctx = _localctx;
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(400);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(410);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParams);
					setState(403);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(404);
					match(COMA);
					setState(405);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(412);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ListArrayContext extends ParserRuleContext {
		public interfaces.Expression p;
		public ListArrayContext list;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public ListArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listArray; }
	}

	public final ListArrayContext listArray() throws RecognitionException {
		return listArray(0);
	}

	private ListArrayContext listArray(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListArrayContext _localctx = new ListArrayContext(_ctx, _parentState);
		ListArrayContext _prevctx = _localctx;
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(414);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVar((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(425);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListArrayContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listArray);
					setState(417);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(418);
					match(CORIZQ);
					setState(419);
					((ListArrayContext)_localctx).expr = expr(0);
					setState(420);
					match(CORDER);
					 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
					}
					} 
				}
				setState(427);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 12:
			return caselist_sempred((CaselistContext)_localctx, predIndex);
		case 15:
			return elseiflist_sempred((ElseiflistContext)_localctx, predIndex);
		case 21:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 22:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 23:
			return listArray_sempred((ListArrayContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean caselist_sempred(CaselistContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean elseiflist_sempred(ElseiflistContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 19);
		case 3:
			return precpred(_ctx, 18);
		case 4:
			return precpred(_ctx, 17);
		case 5:
			return precpred(_ctx, 16);
		case 6:
			return precpred(_ctx, 15);
		case 7:
			return precpred(_ctx, 14);
		case 8:
			return precpred(_ctx, 13);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 9:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 10:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\39\u01af\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\3\2\3\2\3\2\3\2\3\3\6\38\n\3\r\3\16\39\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4e"+
		"\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t"+
		"\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00a8"+
		"\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\7\16\u00bc\n\16\f\16\16\16\u00bf\13\16\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20"+
		"\u00ed\n\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\7\21\u00f7\n\21\f"+
		"\21\16\21\u00fa\13\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23"+
		"\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u011b\n\24\3\25\3\25\3\25\3\25"+
		"\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\5\26\u0132\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\5\27\u0168\n\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\7\27\u018d"+
		"\n\27\f\27\16\27\u0190\13\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\30\7\30\u019b\n\30\f\30\16\30\u019e\13\30\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\31\3\31\7\31\u01aa\n\31\f\31\16\31\u01ad\13\31\3\31\2"+
		"\7\32 ,.\60\32\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\2\b"+
		"\3\2\36\37\3\2*+\3\2,-\4\2&&((\4\2\'\'))\3\2 !\2\u01c3\2\62\3\2\2\2\4"+
		"\67\3\2\2\2\6d\3\2\2\2\bf\3\2\2\2\nn\3\2\2\2\fu\3\2\2\2\16}\3\2\2\2\20"+
		"\u0085\3\2\2\2\22\u0088\3\2\2\2\24\u008b\3\2\2\2\26\u00a7\3\2\2\2\30\u00a9"+
		"\3\2\2\2\32\u00b3\3\2\2\2\34\u00c0\3\2\2\2\36\u00ec\3\2\2\2 \u00ee\3\2"+
		"\2\2\"\u00fb\3\2\2\2$\u0103\3\2\2\2&\u011a\3\2\2\2(\u011c\3\2\2\2*\u0131"+
		"\3\2\2\2,\u0167\3\2\2\2.\u0191\3\2\2\2\60\u019f\3\2\2\2\62\63\5\4\3\2"+
		"\63\64\7\2\2\3\64\65\b\2\1\2\65\3\3\2\2\2\668\5\6\4\2\67\66\3\2\2\289"+
		"\3\2\2\29\67\3\2\2\29:\3\2\2\2:;\3\2\2\2;<\b\3\1\2<\5\3\2\2\2=>\5\24\13"+
		"\2>?\b\4\1\2?e\3\2\2\2@A\5\36\20\2AB\b\4\1\2Be\3\2\2\2CD\5&\24\2DE\b\4"+
		"\1\2Ee\3\2\2\2FG\5(\25\2GH\b\4\1\2He\3\2\2\2IJ\5$\23\2JK\b\4\1\2Ke\3\2"+
		"\2\2LM\5\26\f\2MN\b\4\1\2Ne\3\2\2\2OP\5\30\r\2PQ\b\4\1\2Qe\3\2\2\2RS\5"+
		"\20\t\2ST\b\4\1\2Te\3\2\2\2UV\5\22\n\2VW\b\4\1\2We\3\2\2\2XY\5\16\b\2"+
		"YZ\b\4\1\2Ze\3\2\2\2[\\\5\f\7\2\\]\b\4\1\2]e\3\2\2\2^_\5\n\6\2_`\b\4\1"+
		"\2`e\3\2\2\2ab\5\b\5\2bc\b\4\1\2ce\3\2\2\2d=\3\2\2\2d@\3\2\2\2dC\3\2\2"+
		"\2dF\3\2\2\2dI\3\2\2\2dL\3\2\2\2dO\3\2\2\2dR\3\2\2\2dU\3\2\2\2dX\3\2\2"+
		"\2d[\3\2\2\2d^\3\2\2\2da\3\2\2\2e\7\3\2\2\2fg\7\37\2\2gh\7\66\2\2hi\7"+
		"\32\2\2ij\7.\2\2jk\5,\27\2kl\7/\2\2lm\b\5\1\2m\t\3\2\2\2no\7\37\2\2op"+
		"\7\66\2\2pq\7\31\2\2qr\7.\2\2rs\7/\2\2st\b\6\1\2t\13\3\2\2\2uv\7\37\2"+
		"\2vw\7\66\2\2wx\7\30\2\2xy\7.\2\2yz\5,\27\2z{\7/\2\2{|\b\7\1\2|\r\3\2"+
		"\2\2}~\7\27\2\2~\177\5,\27\2\177\u0080\7\13\2\2\u0080\u0081\7\60\2\2\u0081"+
		"\u0082\5\4\3\2\u0082\u0083\7\61\2\2\u0083\u0084\b\b\1\2\u0084\17\3\2\2"+
		"\2\u0085\u0086\7\24\2\2\u0086\u0087\b\t\1\2\u0087\21\3\2\2\2\u0088\u0089"+
		"\7\25\2\2\u0089\u008a\b\n\1\2\u008a\23\3\2\2\2\u008b\u008c\7\t\2\2\u008c"+
		"\u008d\7.\2\2\u008d\u008e\5,\27\2\u008e\u008f\7/\2\2\u008f\u0090\b\13"+
		"\1\2\u0090\25\3\2\2\2\u0091\u0092\7\r\2\2\u0092\u0093\7\37\2\2\u0093\u0094"+
		"\7\21\2\2\u0094\u0095\5,\27\2\u0095\u0096\7\66\2\2\u0096\u0097\7\66\2"+
		"\2\u0097\u0098\7\66\2\2\u0098\u0099\5,\27\2\u0099\u009a\7\60\2\2\u009a"+
		"\u009b\5\4\3\2\u009b\u009c\7\61\2\2\u009c\u009d\b\f\1\2\u009d\u00a8\3"+
		"\2\2\2\u009e\u009f\7\r\2\2\u009f\u00a0\7\37\2\2\u00a0\u00a1\7\21\2\2\u00a1"+
		"\u00a2\t\2\2\2\u00a2\u00a3\7\60\2\2\u00a3\u00a4\5\4\3\2\u00a4\u00a5\7"+
		"\61\2\2\u00a5\u00a6\b\f\1\2\u00a6\u00a8\3\2\2\2\u00a7\u0091\3\2\2\2\u00a7"+
		"\u009e\3\2\2\2\u00a8\27\3\2\2\2\u00a9\u00aa\7\16\2\2\u00aa\u00ab\5,\27"+
		"\2\u00ab\u00ac\7\60\2\2\u00ac\u00ad\5\32\16\2\u00ad\u00ae\7\20\2\2\u00ae"+
		"\u00af\7\62\2\2\u00af\u00b0\5\4\3\2\u00b0\u00b1\7\61\2\2\u00b1\u00b2\b"+
		"\r\1\2\u00b2\31\3\2\2\2\u00b3\u00b4\b\16\1\2\u00b4\u00b5\5\34\17\2\u00b5"+
		"\u00b6\b\16\1\2\u00b6\u00bd\3\2\2\2\u00b7\u00b8\f\4\2\2\u00b8\u00b9\5"+
		"\34\17\2\u00b9\u00ba\b\16\1\2\u00ba\u00bc\3\2\2\2\u00bb\u00b7\3\2\2\2"+
		"\u00bc\u00bf\3\2\2\2\u00bd\u00bb\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\33"+
		"\3\2\2\2\u00bf\u00bd\3\2\2\2\u00c0\u00c1\7\17\2\2\u00c1\u00c2\5,\27\2"+
		"\u00c2\u00c3\7\62\2\2\u00c3\u00c4\5\4\3\2\u00c4\u00c5\b\17\1\2\u00c5\35"+
		"\3\2\2\2\u00c6\u00c7\7\n\2\2\u00c7\u00c8\5,\27\2\u00c8\u00c9\7\60\2\2"+
		"\u00c9\u00ca\5\4\3\2\u00ca\u00cb\7\61\2\2\u00cb\u00cc\b\20\1\2\u00cc\u00ed"+
		"\3\2\2\2\u00cd\u00ce\7\n\2\2\u00ce\u00cf\5,\27\2\u00cf\u00d0\7\60\2\2"+
		"\u00d0\u00d1\5\4\3\2\u00d1\u00d2\7\61\2\2\u00d2\u00d3\7\13\2\2\u00d3\u00d4"+
		"\7\60\2\2\u00d4\u00d5\5\4\3\2\u00d5\u00d6\7\61\2\2\u00d6\u00d7\b\20\1"+
		"\2\u00d7\u00ed\3\2\2\2\u00d8\u00d9\7\n\2\2\u00d9\u00da\5,\27\2\u00da\u00db"+
		"\7\60\2\2\u00db\u00dc\5\4\3\2\u00dc\u00dd\7\61\2\2\u00dd\u00de\5 \21\2"+
		"\u00de\u00df\b\20\1\2\u00df\u00ed\3\2\2\2\u00e0\u00e1\7\n\2\2\u00e1\u00e2"+
		"\5,\27\2\u00e2\u00e3\7\60\2\2\u00e3\u00e4\5\4\3\2\u00e4\u00e5\7\61\2\2"+
		"\u00e5\u00e6\5 \21\2\u00e6\u00e7\7\13\2\2\u00e7\u00e8\7\60\2\2\u00e8\u00e9"+
		"\5\4\3\2\u00e9\u00ea\7\61\2\2\u00ea\u00eb\b\20\1\2\u00eb\u00ed\3\2\2\2"+
		"\u00ec\u00c6\3\2\2\2\u00ec\u00cd\3\2\2\2\u00ec\u00d8\3\2\2\2\u00ec\u00e0"+
		"\3\2\2\2\u00ed\37\3\2\2\2\u00ee\u00ef\b\21\1\2\u00ef\u00f0\5\"\22\2\u00f0"+
		"\u00f1\b\21\1\2\u00f1\u00f8\3\2\2\2\u00f2\u00f3\f\4\2\2\u00f3\u00f4\5"+
		"\"\22\2\u00f4\u00f5\b\21\1\2\u00f5\u00f7\3\2\2\2\u00f6\u00f2\3\2\2\2\u00f7"+
		"\u00fa\3\2\2\2\u00f8\u00f6\3\2\2\2\u00f8\u00f9\3\2\2\2\u00f9!\3\2\2\2"+
		"\u00fa\u00f8\3\2\2\2\u00fb\u00fc\7\13\2\2\u00fc\u00fd\7\n\2\2\u00fd\u00fe"+
		"\5,\27\2\u00fe\u00ff\7\60\2\2\u00ff\u0100\5\4\3\2\u0100\u0101\7\61\2\2"+
		"\u0101\u0102\b\22\1\2\u0102#\3\2\2\2\u0103\u0104\7\f\2\2\u0104\u0105\5"+
		",\27\2\u0105\u0106\7\60\2\2\u0106\u0107\5\4\3\2\u0107\u0108\7\61\2\2\u0108"+
		"\u0109\b\23\1\2\u0109%\3\2\2\2\u010a\u010b\7\22\2\2\u010b\u010c\7\37\2"+
		"\2\u010c\u010d\7\62\2\2\u010d\u010e\5*\26\2\u010e\u010f\7%\2\2\u010f\u0110"+
		"\5,\27\2\u0110\u0111\b\24\1\2\u0111\u011b\3\2\2\2\u0112\u0113\7\23\2\2"+
		"\u0113\u0114\7\37\2\2\u0114\u0115\7\62\2\2\u0115\u0116\5*\26\2\u0116\u0117"+
		"\7%\2\2\u0117\u0118\5,\27\2\u0118\u0119\b\24\1\2\u0119\u011b\3\2\2\2\u011a"+
		"\u010a\3\2\2\2\u011a\u0112\3\2\2\2\u011b\'\3\2\2\2\u011c\u011d\7\37\2"+
		"\2\u011d\u011e\7%\2\2\u011e\u011f\5,\27\2\u011f\u0120\b\25\1\2\u0120)"+
		"\3\2\2\2\u0121\u0122\7\3\2\2\u0122\u0132\b\26\1\2\u0123\u0124\7\4\2\2"+
		"\u0124\u0132\b\26\1\2\u0125\u0126\7\6\2\2\u0126\u0132\b\26\1\2\u0127\u0128"+
		"\7\5\2\2\u0128\u0132\b\26\1\2\u0129\u012a\7\63\2\2\u012a\u012b\5*\26\2"+
		"\u012b\u012c\7\64\2\2\u012c\u012d\b\26\1\2\u012d\u0132\3\2\2\2\u012e\u012f"+
		"\7\63\2\2\u012f\u0130\7\64\2\2\u0130\u0132\b\26\1\2\u0131\u0121\3\2\2"+
		"\2\u0131\u0123\3\2\2\2\u0131\u0125\3\2\2\2\u0131\u0127\3\2\2\2\u0131\u0129"+
		"\3\2\2\2\u0131\u012e\3\2\2\2\u0132+\3\2\2\2\u0133\u0134\b\27\1\2\u0134"+
		"\u0135\7.\2\2\u0135\u0136\5,\27\2\u0136\u0137\7/\2\2\u0137\u0138\b\27"+
		"\1\2\u0138\u0168\3\2\2\2\u0139\u013a\5\60\31\2\u013a\u013b\b\27\1\2\u013b"+
		"\u0168\3\2\2\2\u013c\u013d\7\3\2\2\u013d\u013e\7.\2\2\u013e\u013f\5,\27"+
		"\2\u013f\u0140\7/\2\2\u0140\u0141\b\27\1\2\u0141\u0168\3\2\2\2\u0142\u0143"+
		"\7\6\2\2\u0143\u0144\7.\2\2\u0144\u0145\5,\27\2\u0145\u0146\7/\2\2\u0146"+
		"\u0147\b\27\1\2\u0147\u0168\3\2\2\2\u0148\u0149\7\4\2\2\u0149\u014a\7"+
		".\2\2\u014a\u014b\5,\27\2\u014b\u014c\7/\2\2\u014c\u014d\b\27\1\2\u014d"+
		"\u0168\3\2\2\2\u014e\u014f\7\37\2\2\u014f\u0150\7\66\2\2\u0150\u0151\7"+
		"\33\2\2\u0151\u0152\7.\2\2\u0152\u0153\7/\2\2\u0153\u0168\b\27\1\2\u0154"+
		"\u0155\7\37\2\2\u0155\u0156\7\66\2\2\u0156\u0157\7\34\2\2\u0157\u0158"+
		"\7.\2\2\u0158\u0159\7/\2\2\u0159\u0168\b\27\1\2\u015a\u015b\7\63\2\2\u015b"+
		"\u015c\5.\30\2\u015c\u015d\7\64\2\2\u015d\u015e\b\27\1\2\u015e\u0168\3"+
		"\2\2\2\u015f\u0160\7\35\2\2\u0160\u0168\b\27\1\2\u0161\u0162\7\36\2\2"+
		"\u0162\u0168\b\27\1\2\u0163\u0164\7\7\2\2\u0164\u0168\b\27\1\2\u0165\u0166"+
		"\7\b\2\2\u0166\u0168\b\27\1\2\u0167\u0133\3\2\2\2\u0167\u0139\3\2\2\2"+
		"\u0167\u013c\3\2\2\2\u0167\u0142\3\2\2\2\u0167\u0148\3\2\2\2\u0167\u014e"+
		"\3\2\2\2\u0167\u0154\3\2\2\2\u0167\u015a\3\2\2\2\u0167\u015f\3\2\2\2\u0167"+
		"\u0161\3\2\2\2\u0167\u0163\3\2\2\2\u0167\u0165\3\2\2\2\u0168\u018e\3\2"+
		"\2\2\u0169\u016a\f\25\2\2\u016a\u016b\t\3\2\2\u016b\u016c\5,\27\26\u016c"+
		"\u016d\b\27\1\2\u016d\u018d\3\2\2\2\u016e\u016f\f\24\2\2\u016f\u0170\t"+
		"\4\2\2\u0170\u0171\5,\27\25\u0171\u0172\b\27\1\2\u0172\u018d\3\2\2\2\u0173"+
		"\u0174\f\23\2\2\u0174\u0175\t\5\2\2\u0175\u0176\5,\27\24\u0176\u0177\b"+
		"\27\1\2\u0177\u018d\3\2\2\2\u0178\u0179\f\22\2\2\u0179\u017a\t\6\2\2\u017a"+
		"\u017b\5,\27\23\u017b\u017c\b\27\1\2\u017c\u018d\3\2\2\2\u017d\u017e\f"+
		"\21\2\2\u017e\u017f\t\7\2\2\u017f\u0180\5,\27\22\u0180\u0181\b\27\1\2"+
		"\u0181\u018d\3\2\2\2\u0182\u0183\f\20\2\2\u0183\u0184\7$\2\2\u0184\u0185"+
		"\5,\27\21\u0185\u0186\b\27\1\2\u0186\u018d\3\2\2\2\u0187\u0188\f\17\2"+
		"\2\u0188\u0189\7#\2\2\u0189\u018a\5,\27\20\u018a\u018b\b\27\1\2\u018b"+
		"\u018d\3\2\2\2\u018c\u0169\3\2\2\2\u018c\u016e\3\2\2\2\u018c\u0173\3\2"+
		"\2\2\u018c\u0178\3\2\2\2\u018c\u017d\3\2\2\2\u018c\u0182\3\2\2\2\u018c"+
		"\u0187\3\2\2\2\u018d\u0190\3\2\2\2\u018e\u018c\3\2\2\2\u018e\u018f\3\2"+
		"\2\2\u018f-\3\2\2\2\u0190\u018e\3\2\2\2\u0191\u0192\b\30\1\2\u0192\u0193"+
		"\5,\27\2\u0193\u0194\b\30\1\2\u0194\u019c\3\2\2\2\u0195\u0196\f\4\2\2"+
		"\u0196\u0197\7\65\2\2\u0197\u0198\5,\27\2\u0198\u0199\b\30\1\2\u0199\u019b"+
		"\3\2\2\2\u019a\u0195\3\2\2\2\u019b\u019e\3\2\2\2\u019c\u019a\3\2\2\2\u019c"+
		"\u019d\3\2\2\2\u019d/\3\2\2\2\u019e\u019c\3\2\2\2\u019f\u01a0\b\31\1\2"+
		"\u01a0\u01a1\7\37\2\2\u01a1\u01a2\b\31\1\2\u01a2\u01ab\3\2\2\2\u01a3\u01a4"+
		"\f\4\2\2\u01a4\u01a5\7\63\2\2\u01a5\u01a6\5,\27\2\u01a6\u01a7\7\64\2\2"+
		"\u01a7\u01a8\b\31\1\2\u01a8\u01aa\3\2\2\2\u01a9\u01a3\3\2\2\2\u01aa\u01ad"+
		"\3\2\2\2\u01ab\u01a9\3\2\2\2\u01ab\u01ac\3\2\2\2\u01ac\61\3\2\2\2\u01ad"+
		"\u01ab\3\2\2\2\179d\u00a7\u00bd\u00ec\u00f8\u011a\u0131\u0167\u018c\u018e"+
		"\u019c\u01ab";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}