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
		COUNT=25, ISEMPTY=26, AT=27, FUNC=28, STRUCT=29, NUMBER=30, STRING=31, 
		ID=32, DIF=33, IG_IG=34, NOT=35, OR=36, AND=37, IG=38, MAY_IG=39, MEN_IG=40, 
		MAYOR=41, MENOR=42, MUL=43, DIV=44, ADD=45, SUB=46, PARIZQ=47, PARDER=48, 
		LLAVEIZQ=49, LLAVEDER=50, D_PTS=51, CORIZQ=52, CORDER=53, COMA=54, ARROW1=55, 
		PUNTO=56, MODULO=57, GUIONB=58, WHITESPACE=59, COMMENT=60, LINE_COMMENT=61;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_returntmt = 3, 
		RULE_function = 4, RULE_listParamsFunc = 5, RULE_structCreation = 6, RULE_listStructDec = 7, 
		RULE_removetmt = 8, RULE_removelastmt = 9, RULE_appendtmt = 10, RULE_guardtmt = 11, 
		RULE_breaktmt = 12, RULE_continuetmt = 13, RULE_printstmt = 14, RULE_forstmt = 15, 
		RULE_switchtmt = 16, RULE_caselist = 17, RULE_case = 18, RULE_ifstmt = 19, 
		RULE_elseiflist = 20, RULE_elseif = 21, RULE_whilestmt = 22, RULE_declarationstmt = 23, 
		RULE_asignacion = 24, RULE_types = 25, RULE_expr = 26, RULE_listParams = 27, 
		RULE_listArray = 28, RULE_callFunctionInst = 29, RULE_callFunction = 30, 
		RULE_listParamsCall = 31, RULE_listStructExp = 32;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "returntmt", "function", "listParamsFunc", 
			"structCreation", "listStructDec", "removetmt", "removelastmt", "appendtmt", 
			"guardtmt", "breaktmt", "continuetmt", "printstmt", "forstmt", "switchtmt", 
			"caselist", "case", "ifstmt", "elseiflist", "elseif", "whilestmt", "declarationstmt", 
			"asignacion", "types", "expr", "listParams", "listArray", "callFunctionInst", 
			"callFunction", "listParamsCall", "listStructExp"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'true'", "'false'", 
			"'print'", "'if'", "'else'", "'while'", "'for'", "'switch'", "'case'", 
			"'default'", "'in'", "'var'", "'let'", "'break'", "'continue'", "'return'", 
			"'guard'", "'append'", "'removeLast'", "'remove'", "'count'", "'isEmpty'", 
			"'at'", "'func'", "'struct'", null, null, null, "'!='", "'=='", "'!'", 
			"'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", 
			"'-'", "'('", "')'", "'{'", "'}'", "':'", "'['", "']'", "','", "'->'", 
			"'.'", "'%'", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE", 
			"WHILE", "FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "LET", "BREAK", 
			"CONTINUE", "RETURN", "GUARD", "APPEND", "REMOVELAST", "REMOVE", "COUNT", 
			"ISEMPTY", "AT", "FUNC", "STRUCT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", 
			"CORIZQ", "CORDER", "COMA", "ARROW1", "PUNTO", "MODULO", "GUIONB", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT"
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
			setState(66);
			((SContext)_localctx).block = block();
			setState(67);
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
			setState(71); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(70);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(73); 
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
		public FunctionContext function;
		public CallFunctionInstContext callFunctionInst;
		public StructCreationContext structCreation;
		public ReturntmtContext returntmt;
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
		public FunctionContext function() {
			return getRuleContext(FunctionContext.class,0);
		}
		public CallFunctionInstContext callFunctionInst() {
			return getRuleContext(CallFunctionInstContext.class,0);
		}
		public StructCreationContext structCreation() {
			return getRuleContext(StructCreationContext.class,0);
		}
		public ReturntmtContext returntmt() {
			return getRuleContext(ReturntmtContext.class,0);
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
			setState(128);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(77);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(80);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(83);
				((InstructionContext)_localctx).declarationstmt = declarationstmt();
				 _localctx.inst = ((InstructionContext)_localctx).declarationstmt.dec 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(86);
				((InstructionContext)_localctx).asignacion = asignacion();
				 _localctx.inst = ((InstructionContext)_localctx).asignacion.asign
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(89);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whiles 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(92);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.fors 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(95);
				((InstructionContext)_localctx).switchtmt = switchtmt();
				 _localctx.inst = ((InstructionContext)_localctx).switchtmt.swtch 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(98);
				((InstructionContext)_localctx).breaktmt = breaktmt();
				 _localctx.inst = ((InstructionContext)_localctx).breaktmt.break 
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(101);
				((InstructionContext)_localctx).continuetmt = continuetmt();
				 _localctx.inst = ((InstructionContext)_localctx).continuetmt.continue 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(104);
				((InstructionContext)_localctx).guardtmt = guardtmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardtmt.guard 
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(107);
				((InstructionContext)_localctx).appendtmt = appendtmt();
				 _localctx.inst = ((InstructionContext)_localctx).appendtmt.append
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(110);
				((InstructionContext)_localctx).removelastmt = removelastmt();
				 _localctx.inst = ((InstructionContext)_localctx).removelastmt.removelast 
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(113);
				((InstructionContext)_localctx).removetmt = removetmt();
				 _localctx.inst = ((InstructionContext)_localctx).removetmt.remove 
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(116);
				((InstructionContext)_localctx).function = function();
				 _localctx.inst = ((InstructionContext)_localctx).function.fun 
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(119);
				((InstructionContext)_localctx).callFunctionInst = callFunctionInst();
				 _localctx.inst = ((InstructionContext)_localctx).callFunctionInst.cfi 
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(122);
				((InstructionContext)_localctx).structCreation = structCreation();
				 _localctx.inst = ((InstructionContext)_localctx).structCreation.dec 
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(125);
				((InstructionContext)_localctx).returntmt = returntmt();
				 _localctx.inst = ((InstructionContext)_localctx).returntmt.ret 
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

	public static class ReturntmtContext extends ParserRuleContext {
		public interfaces.Instruction ret;
		public Token RETURN;
		public ExprContext expr;
		public TerminalNode RETURN() { return getToken(SwiftGrammarParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturntmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returntmt; }
	}

	public final ReturntmtContext returntmt() throws RecognitionException {
		ReturntmtContext _localctx = new ReturntmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_returntmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			((ReturntmtContext)_localctx).RETURN = match(RETURN);
			setState(131);
			((ReturntmtContext)_localctx).expr = expr(0);
			  _localctx.ret = instructions.NewReturn((((ReturntmtContext)_localctx).RETURN!=null?((ReturntmtContext)_localctx).RETURN.getLine():0), (((ReturntmtContext)_localctx).RETURN!=null?((ReturntmtContext)_localctx).RETURN.getCharPositionInLine():0), ((ReturntmtContext)_localctx).expr.e)  
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

	public static class FunctionContext extends ParserRuleContext {
		public interfaces.Instruction fun;
		public Token FUNC;
		public Token ID;
		public ListParamsFuncContext listParamsFunc;
		public BlockContext block;
		public TypesContext types;
		public TerminalNode FUNC() { return getToken(SwiftGrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListParamsFuncContext listParamsFunc() {
			return getRuleContext(ListParamsFuncContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public TerminalNode ARROW1() { return getToken(SwiftGrammarParser.ARROW1, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public FunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function; }
	}

	public final FunctionContext function() throws RecognitionException {
		FunctionContext _localctx = new FunctionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_function);
		try {
			setState(156);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(134);
				((FunctionContext)_localctx).FUNC = match(FUNC);
				setState(135);
				((FunctionContext)_localctx).ID = match(ID);
				setState(136);
				match(PARIZQ);
				setState(137);
				((FunctionContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(138);
				match(PARDER);
				setState(139);
				match(LLAVEIZQ);
				setState(140);
				((FunctionContext)_localctx).block = block();
				setState(141);
				match(LLAVEDER);

				    _localctx.fun = instructions.NewFunction((((FunctionContext)_localctx).FUNC!=null?((FunctionContext)_localctx).FUNC.getLine():0), (((FunctionContext)_localctx).FUNC!=null?((FunctionContext)_localctx).FUNC.getCharPositionInLine():0), (((FunctionContext)_localctx).ID!=null?((FunctionContext)_localctx).ID.getText():null), ((FunctionContext)_localctx).listParamsFunc.lpf, environment.NULL, ((FunctionContext)_localctx).block.blk)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(144);
				((FunctionContext)_localctx).FUNC = match(FUNC);
				setState(145);
				((FunctionContext)_localctx).ID = match(ID);
				setState(146);
				match(PARIZQ);
				setState(147);
				((FunctionContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(148);
				match(PARDER);
				setState(149);
				match(ARROW1);
				setState(150);
				((FunctionContext)_localctx).types = types();
				setState(151);
				match(LLAVEIZQ);
				setState(152);
				((FunctionContext)_localctx).block = block();
				setState(153);
				match(LLAVEDER);

				    _localctx.fun = instructions.NewFunction((((FunctionContext)_localctx).FUNC!=null?((FunctionContext)_localctx).FUNC.getLine():0), (((FunctionContext)_localctx).FUNC!=null?((FunctionContext)_localctx).FUNC.getCharPositionInLine():0), (((FunctionContext)_localctx).ID!=null?((FunctionContext)_localctx).ID.getText():null), ((FunctionContext)_localctx).listParamsFunc.lpf, ((FunctionContext)_localctx).types.ty, ((FunctionContext)_localctx).block.blk)

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

	public static class ListParamsFuncContext extends ParserRuleContext {
		public []interface{} lpf;
		public ListParamsFuncContext list;
		public Token ID;
		public TypesContext types;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsFuncContext listParamsFunc() {
			return getRuleContext(ListParamsFuncContext.class,0);
		}
		public ListParamsFuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParamsFunc; }
	}

	public final ListParamsFuncContext listParamsFunc() throws RecognitionException {
		return listParamsFunc(0);
	}

	private ListParamsFuncContext listParamsFunc(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsFuncContext _localctx = new ListParamsFuncContext(_ctx, _parentState);
		ListParamsFuncContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_listParamsFunc, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(159);
				((ListParamsFuncContext)_localctx).ID = match(ID);
				setState(160);
				match(D_PTS);
				setState(161);
				((ListParamsFuncContext)_localctx).types = types();

				    _localctx.lpf = []interface{}{}
				    newParam := instructions.NewParamsDeclaration((((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getLine():0), (((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getCharPositionInLine():0), (((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getText():null), ((ListParamsFuncContext)_localctx).types.ty)
				    _localctx.lpf = append(_localctx.lpf, newParam)
				    
				}
				break;
			case 2:
				{
				 _localctx.lpf = []interface{}{} 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(176);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsFuncContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParamsFunc);
					setState(167);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(168);
					match(COMA);
					setState(169);
					((ListParamsFuncContext)_localctx).ID = match(ID);
					setState(170);
					match(D_PTS);
					setState(171);
					((ListParamsFuncContext)_localctx).types = types();

					              var arr []interface{}
					              newParam := instructions.NewParamsDeclaration((((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getLine():0), (((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getCharPositionInLine():0), (((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getText():null), ((ListParamsFuncContext)_localctx).types.ty)
					              arr = append(((ListParamsFuncContext)_localctx).list.lpf, newParam)
					              _localctx.lpf = arr
					              
					}
					} 
				}
				setState(178);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
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

	public static class StructCreationContext extends ParserRuleContext {
		public interfaces.Instruction dec;
		public Token STRUCT;
		public Token ID;
		public ListStructDecContext listStructDec;
		public TerminalNode STRUCT() { return getToken(SwiftGrammarParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(SwiftGrammarParser.LLAVEIZQ, 0); }
		public ListStructDecContext listStructDec() {
			return getRuleContext(ListStructDecContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(SwiftGrammarParser.LLAVEDER, 0); }
		public StructCreationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structCreation; }
	}

	public final StructCreationContext structCreation() throws RecognitionException {
		StructCreationContext _localctx = new StructCreationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_structCreation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(179);
			((StructCreationContext)_localctx).STRUCT = match(STRUCT);
			setState(180);
			((StructCreationContext)_localctx).ID = match(ID);
			setState(181);
			match(LLAVEIZQ);
			setState(182);
			((StructCreationContext)_localctx).listStructDec = listStructDec(0);
			setState(183);
			match(LLAVEDER);
			 _localctx.dec = instructions.NewStruct((((StructCreationContext)_localctx).STRUCT!=null?((StructCreationContext)_localctx).STRUCT.getLine():0), (((StructCreationContext)_localctx).STRUCT!=null?((StructCreationContext)_localctx).STRUCT.getCharPositionInLine():0), (((StructCreationContext)_localctx).ID!=null?((StructCreationContext)_localctx).ID.getText():null), ((StructCreationContext)_localctx).listStructDec.l) 
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

	public static class ListStructDecContext extends ParserRuleContext {
		public []interface{} l;
		public ListStructDecContext list;
		public Token ID;
		public TypesContext types;
		public TerminalNode VAR() { return getToken(SwiftGrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListStructDecContext listStructDec() {
			return getRuleContext(ListStructDecContext.class,0);
		}
		public ListStructDecContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listStructDec; }
	}

	public final ListStructDecContext listStructDec() throws RecognitionException {
		return listStructDec(0);
	}

	private ListStructDecContext listStructDec(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListStructDecContext _localctx = new ListStructDecContext(_ctx, _parentState);
		ListStructDecContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_listStructDec, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(194);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(187);
				match(VAR);
				setState(188);
				((ListStructDecContext)_localctx).ID = match(ID);
				setState(189);
				match(D_PTS);
				setState(190);
				((ListStructDecContext)_localctx).types = types();

				                        var arr []interface{}
				                        newParams := environment.NewStructType((((ListStructDecContext)_localctx).ID!=null?((ListStructDecContext)_localctx).ID.getText():null), ((ListStructDecContext)_localctx).types.ty)
				                        arr = append(arr, newParams)
				                        _localctx.l = arr
				                    
				}
				break;
			case 2:
				{
				 _localctx.l = []interface{}{} 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(206);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListStructDecContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listStructDec);
					setState(196);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(197);
					match(COMA);
					setState(198);
					match(VAR);
					setState(199);
					((ListStructDecContext)_localctx).ID = match(ID);
					setState(200);
					match(D_PTS);
					setState(201);
					((ListStructDecContext)_localctx).types = types();

					                                                          var arr []interface{}
					                                                          newParams := environment.NewStructType((((ListStructDecContext)_localctx).ID!=null?((ListStructDecContext)_localctx).ID.getText():null), ((ListStructDecContext)_localctx).types.ty)
					                                                          arr = append(((ListStructDecContext)_localctx).list.l, newParams)
					                                                          _localctx.l = arr
					                                                      
					}
					} 
				}
				setState(208);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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

	public static class RemovetmtContext extends ParserRuleContext {
		public interfaces.Instruction remove;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode REMOVE() { return getToken(SwiftGrammarParser.REMOVE, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public TerminalNode AT() { return getToken(SwiftGrammarParser.AT, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
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
		enterRule(_localctx, 16, RULE_removetmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(209);
			((RemovetmtContext)_localctx).ID = match(ID);
			setState(210);
			match(PUNTO);
			setState(211);
			match(REMOVE);
			setState(212);
			match(PARIZQ);
			setState(213);
			match(AT);
			setState(214);
			match(D_PTS);
			setState(215);
			((RemovetmtContext)_localctx).expr = expr(0);
			setState(216);
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
		enterRule(_localctx, 18, RULE_removelastmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(219);
			((RemovelastmtContext)_localctx).ID = match(ID);
			setState(220);
			match(PUNTO);
			setState(221);
			match(REMOVELAST);
			setState(222);
			match(PARIZQ);
			setState(223);
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
		enterRule(_localctx, 20, RULE_appendtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			((AppendtmtContext)_localctx).ID = match(ID);
			setState(227);
			match(PUNTO);
			setState(228);
			match(APPEND);
			setState(229);
			match(PARIZQ);
			setState(230);
			((AppendtmtContext)_localctx).expr = expr(0);
			setState(231);
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
		enterRule(_localctx, 22, RULE_guardtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
			((GuardtmtContext)_localctx).GUARD = match(GUARD);
			setState(235);
			((GuardtmtContext)_localctx).expr = expr(0);
			setState(236);
			match(ELSE);
			setState(237);
			match(LLAVEIZQ);
			setState(238);
			((GuardtmtContext)_localctx).block = block();
			setState(239);
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
		enterRule(_localctx, 24, RULE_breaktmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
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
		enterRule(_localctx, 26, RULE_continuetmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(245);
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
		public ExprContext expr1;
		public ExprContext expr2;
		public TerminalNode PRINT() { return getToken(SwiftGrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_printstmt);
		try {
			setState(262);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(248);
				((PrintstmtContext)_localctx).PRINT = match(PRINT);
				setState(249);
				match(PARIZQ);
				setState(250);
				((PrintstmtContext)_localctx).expr = expr(0);
				setState(251);
				match(PARDER);
				 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr.e, nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(254);
				((PrintstmtContext)_localctx).PRINT = match(PRINT);
				setState(255);
				match(PARIZQ);
				setState(256);
				((PrintstmtContext)_localctx).expr1 = expr(0);
				setState(257);
				match(COMA);
				setState(258);
				((PrintstmtContext)_localctx).expr2 = expr(0);
				setState(259);
				match(PARDER);
				 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr1.e, ((PrintstmtContext)_localctx).expr2.e)
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

	public static class ForstmtContext extends ParserRuleContext {
		public interfaces.Instruction fors;
		public Token FOR;
		public Token ID;
		public ExprContext e1;
		public ExprContext e2;
		public BlockContext block;
		public Token GUIONB;
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
		public TerminalNode GUIONB() { return getToken(SwiftGrammarParser.GUIONB, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_forstmt);
		int _la;
		try {
			setState(299);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(264);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(265);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(266);
				match(IN);
				setState(267);
				((ForstmtContext)_localctx).e1 = expr(0);
				setState(268);
				match(PUNTO);
				setState(269);
				match(PUNTO);
				setState(270);
				match(PUNTO);
				setState(271);
				((ForstmtContext)_localctx).e2 = expr(0);
				setState(272);
				match(LLAVEIZQ);
				setState(273);
				((ForstmtContext)_localctx).block = block();
				setState(274);
				match(LLAVEDER);
				 _localctx.fors = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).e1.e,((ForstmtContext)_localctx).e2.e,"nil", ((ForstmtContext)_localctx).block.blk) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(277);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(278);
				((ForstmtContext)_localctx).GUIONB = match(GUIONB);
				setState(279);
				match(IN);
				setState(280);
				((ForstmtContext)_localctx).e1 = expr(0);
				setState(281);
				match(PUNTO);
				setState(282);
				match(PUNTO);
				setState(283);
				match(PUNTO);
				setState(284);
				((ForstmtContext)_localctx).e2 = expr(0);
				setState(285);
				match(LLAVEIZQ);
				setState(286);
				((ForstmtContext)_localctx).block = block();
				setState(287);
				match(LLAVEDER);
				 _localctx.fors = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).GUIONB!=null?((ForstmtContext)_localctx).GUIONB.getText():null), ((ForstmtContext)_localctx).e1.e,((ForstmtContext)_localctx).e2.e,"nil", ((ForstmtContext)_localctx).block.blk) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(290);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(291);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(292);
				match(IN);
				setState(293);
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
				setState(294);
				match(LLAVEIZQ);
				setState(295);
				((ForstmtContext)_localctx).block = block();
				setState(296);
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
		enterRule(_localctx, 32, RULE_switchtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(301);
			((SwitchtmtContext)_localctx).SWITCH = match(SWITCH);
			setState(302);
			((SwitchtmtContext)_localctx).expr = expr(0);
			setState(303);
			match(LLAVEIZQ);
			setState(304);
			((SwitchtmtContext)_localctx).caselist = caselist(0);
			setState(305);
			match(DEFAULT);
			setState(306);
			match(D_PTS);
			setState(307);
			((SwitchtmtContext)_localctx).block = block();
			setState(308);
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
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_caselist, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(312);
			((CaselistContext)_localctx).case = case();

			                            _localctx.cases= []interface{}{}
			                            _localctx.cases = append(_localctx.cases, ((CaselistContext)_localctx).case.swtch)
			                        
			}
			_ctx.stop = _input.LT(-1);
			setState(321);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
					setState(315);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(316);
					((CaselistContext)_localctx).case = case();
					 var arr3 []interface{}
					                                  arr3= append(((CaselistContext)_localctx).listcases.cases, ((CaselistContext)_localctx).case.swtch)
					                                  _localctx.cases= arr3
					                                  
					}
					} 
				}
				setState(323);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
		enterRule(_localctx, 36, RULE_case);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			((CaseContext)_localctx).CASE = match(CASE);
			setState(325);
			((CaseContext)_localctx).expr = expr(0);
			setState(326);
			match(D_PTS);
			setState(327);
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
		enterRule(_localctx, 38, RULE_ifstmt);
		try {
			setState(368);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(330);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(331);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(332);
				match(LLAVEIZQ);
				setState(333);
				((IfstmtContext)_localctx).block = block();
				setState(334);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(337);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(338);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(339);
				match(LLAVEIZQ);
				setState(340);
				((IfstmtContext)_localctx).b1 = block();
				setState(341);
				match(LLAVEDER);
				setState(342);
				match(ELSE);
				setState(343);
				match(LLAVEIZQ);
				setState(344);
				((IfstmtContext)_localctx).b2 = block();
				setState(345);
				match(LLAVEDER);
				_localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).b1.blk, nil, ((IfstmtContext)_localctx).b2.blk) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(348);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(349);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(350);
				match(LLAVEIZQ);
				setState(351);
				((IfstmtContext)_localctx).block = block();
				setState(352);
				match(LLAVEDER);
				setState(353);
				((IfstmtContext)_localctx).elseiflist = elseiflist(0);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, ((IfstmtContext)_localctx).elseiflist.else_if_list, nil) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(356);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(357);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(358);
				match(LLAVEIZQ);
				setState(359);
				((IfstmtContext)_localctx).b1 = block();
				setState(360);
				match(LLAVEDER);
				setState(361);
				((IfstmtContext)_localctx).elseiflist = elseiflist(0);
				setState(362);
				match(ELSE);
				setState(363);
				match(LLAVEIZQ);
				setState(364);
				((IfstmtContext)_localctx).b2 = block();
				setState(365);
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
		int _startState = 40;
		enterRecursionRule(_localctx, 40, RULE_elseiflist, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(371);
			((ElseiflistContext)_localctx).elseif = elseif();

			                _localctx.else_if_list = []interface{}{}
			                _localctx.else_if_list = append(_localctx.else_if_list, ((ElseiflistContext)_localctx).elseif.ifinst)

			            
			}
			_ctx.stop = _input.LT(-1);
			setState(380);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
					setState(374);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(375);
					((ElseiflistContext)_localctx).elseif = elseif();
					   var arr2 []interface{}
					                                  arr2 = append(((ElseiflistContext)_localctx).listif.else_if_list, ((ElseiflistContext)_localctx).elseif.ifinst)
					                                  _localctx.else_if_list = arr2
					                              
					}
					} 
				}
				setState(382);
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
		enterRule(_localctx, 42, RULE_elseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(383);
			match(ELSE);
			setState(384);
			((ElseifContext)_localctx).IF = match(IF);
			setState(385);
			((ElseifContext)_localctx).expr = expr(0);
			setState(386);
			match(LLAVEIZQ);
			setState(387);
			((ElseifContext)_localctx).block = block();
			setState(388);
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
		enterRule(_localctx, 44, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(391);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(392);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(393);
			match(LLAVEIZQ);
			setState(394);
			((WhilestmtContext)_localctx).block = block();
			setState(395);
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
		enterRule(_localctx, 46, RULE_declarationstmt);
		try {
			setState(438);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(398);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(399);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(400);
				match(D_PTS);
				setState(401);
				((DeclarationstmtContext)_localctx).types = types();
				setState(402);
				match(IG);
				setState(403);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, false);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(406);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(407);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(408);
				match(IG);
				setState(409);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.NULL, ((DeclarationstmtContext)_localctx).expr.e, false);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(412);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(413);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(414);
				match(IG);
				setState(415);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.NULL, ((DeclarationstmtContext)_localctx).expr.e, true);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(418);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(419);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(420);
				match(D_PTS);
				setState(421);
				((DeclarationstmtContext)_localctx).types = types();
				setState(422);
				match(IG);
				setState(423);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, true);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(426);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(427);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(428);
				match(IG);
				setState(429);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 ((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.STRUCT , ((DeclarationstmtContext)_localctx).expr.e, false) ;
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(432);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(433);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(434);
				match(IG);
				setState(435);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 ((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.STRUCT , ((DeclarationstmtContext)_localctx).expr.e, true) ;
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

	public static class AsignacionContext extends ParserRuleContext {
		public interfaces.Instruction asign;
		public Token ID;
		public Token IG;
		public ExprContext expr;
		public Token ADD;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(SwiftGrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_asignacion);
		try {
			setState(451);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(440);
				((AsignacionContext)_localctx).ID = match(ID);
				setState(441);
				((AsignacionContext)_localctx).IG = match(IG);
				setState(442);
				((AsignacionContext)_localctx).expr = expr(0);
				 _localctx.asign = instructions.NewAsignacion((((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getLine():0) ,(((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getCharPositionInLine():0) ,(((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getText():null), ((AsignacionContext)_localctx).expr.e, (((AsignacionContext)_localctx).IG!=null?((AsignacionContext)_localctx).IG.getText():null)) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(445);
				((AsignacionContext)_localctx).ID = match(ID);
				setState(446);
				((AsignacionContext)_localctx).ADD = match(ADD);
				setState(447);
				((AsignacionContext)_localctx).IG = match(IG);
				setState(448);
				((AsignacionContext)_localctx).expr = expr(0);
				 _localctx.asign = instructions.NewAsignacion((((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getLine():0) ,(((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getCharPositionInLine():0) ,(((AsignacionContext)_localctx).ID!=null?((AsignacionContext)_localctx).ID.getText():null), ((AsignacionContext)_localctx).expr.e, (((AsignacionContext)_localctx).ADD!=null?((AsignacionContext)_localctx).ADD.getText():null)+(((AsignacionContext)_localctx).IG!=null?((AsignacionContext)_localctx).IG.getText():null) ) 
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
		enterRule(_localctx, 50, RULE_types);
		try {
			setState(469);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(453);
				match(INT);
				 _localctx.ty = environment.INTEGER 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(455);
				match(FLOAT);
				 _localctx.ty = environment.FLOAT 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(457);
				match(STR);
				 _localctx.ty = environment.STRING 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(459);
				match(BOOL);
				 _localctx.ty = environment.BOOLEAN 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(461);
				match(CORIZQ);
				setState(462);
				types();
				setState(463);
				match(CORDER);
				 _localctx.ty = environment.ARRAY 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(466);
				match(CORIZQ);
				setState(467);
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
		public Token op;
		public ExprContext right;
		public ExprContext expr;
		public CallFunctionContext callFunction;
		public Token ID;
		public ListStructExpContext listStructExp;
		public ListArrayContext list;
		public Token CORIZQ;
		public ListParamsContext listParams;
		public Token INT;
		public Token STR;
		public Token FLOAT;
		public Token NUMBER;
		public Token STRING;
		public Token TRU;
		public Token FAL;
		public TerminalNode SUB() { return getToken(SwiftGrammarParser.SUB, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode NOT() { return getToken(SwiftGrammarParser.NOT, 0); }
		public CallFunctionContext callFunction() {
			return getRuleContext(CallFunctionContext.class,0);
		}
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListStructExpContext listStructExp() {
			return getRuleContext(ListStructExpContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public TerminalNode CORIZQ() { return getToken(SwiftGrammarParser.CORIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(SwiftGrammarParser.CORDER, 0); }
		public TerminalNode INT() { return getToken(SwiftGrammarParser.INT, 0); }
		public TerminalNode STR() { return getToken(SwiftGrammarParser.STR, 0); }
		public TerminalNode FLOAT() { return getToken(SwiftGrammarParser.FLOAT, 0); }
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(SwiftGrammarParser.COUNT, 0); }
		public TerminalNode ISEMPTY() { return getToken(SwiftGrammarParser.ISEMPTY, 0); }
		public TerminalNode NUMBER() { return getToken(SwiftGrammarParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(SwiftGrammarParser.STRING, 0); }
		public TerminalNode TRU() { return getToken(SwiftGrammarParser.TRU, 0); }
		public TerminalNode FAL() { return getToken(SwiftGrammarParser.FAL, 0); }
		public TerminalNode MUL() { return getToken(SwiftGrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(SwiftGrammarParser.DIV, 0); }
		public TerminalNode MODULO() { return getToken(SwiftGrammarParser.MODULO, 0); }
		public TerminalNode ADD() { return getToken(SwiftGrammarParser.ADD, 0); }
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
		int _startState = 52;
		enterRecursionRule(_localctx, 52, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(536);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				{
				setState(472);
				((ExprContext)_localctx).op = match(SUB);
				setState(473);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(23);
				 _localctx.e = expressions.NewOperation(((ExprContext)_localctx).op.GetLine(), ((ExprContext)_localctx).op.GetColumn(), ((ExprContext)_localctx).right.e, "neg", ((ExprContext)_localctx).right.e) 
				}
				break;
			case 2:
				{
				setState(476);
				((ExprContext)_localctx).op = match(NOT);
				setState(477);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
				 _localctx.e = expressions.NewOperation(((ExprContext)_localctx).op.GetLine(), ((ExprContext)_localctx).op.GetColumn(),((ExprContext)_localctx).right.e , (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
				}
				break;
			case 3:
				{
				setState(480);
				((ExprContext)_localctx).callFunction = callFunction();
				 _localctx.e = ((ExprContext)_localctx).callFunction.cf 
				}
				break;
			case 4:
				{
				setState(483);
				((ExprContext)_localctx).ID = match(ID);
				setState(484);
				match(PARIZQ);
				setState(485);
				((ExprContext)_localctx).listStructExp = listStructExp(0);
				setState(486);
				match(PARDER);
				 _localctx.e = expressions.NewStructExp((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null), ((ExprContext)_localctx).listStructExp.l ) 
				}
				break;
			case 5:
				{
				setState(489);
				match(PARIZQ);
				setState(490);
				((ExprContext)_localctx).expr = expr(0);
				setState(491);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 6:
				{
				setState(494);
				((ExprContext)_localctx).list = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).list.p
				}
				break;
			case 7:
				{
				setState(497);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(498);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(499);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case 8:
				{
				setState(502);
				((ExprContext)_localctx).INT = match(INT);
				setState(503);
				match(PARIZQ);
				setState(504);
				((ExprContext)_localctx).expr = expr(0);
				setState(505);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).INT.GetLine(),((ExprContext)_localctx).INT.GetColumn(),((ExprContext)_localctx).expr.e, environment.INTEGER) 
				}
				break;
			case 9:
				{
				setState(508);
				((ExprContext)_localctx).STR = match(STR);
				setState(509);
				match(PARIZQ);
				setState(510);
				((ExprContext)_localctx).expr = expr(0);
				setState(511);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).STR.GetLine(),((ExprContext)_localctx).STR.GetColumn(),((ExprContext)_localctx).expr.e, environment.STRING) 
				}
				break;
			case 10:
				{
				setState(514);
				((ExprContext)_localctx).FLOAT = match(FLOAT);
				setState(515);
				match(PARIZQ);
				setState(516);
				((ExprContext)_localctx).expr = expr(0);
				setState(517);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).FLOAT.GetLine(),((ExprContext)_localctx).FLOAT.GetColumn(),((ExprContext)_localctx).expr.e, environment.FLOAT) 
				}
				break;
			case 11:
				{
				setState(520);
				((ExprContext)_localctx).ID = match(ID);
				setState(521);
				match(PUNTO);
				setState(522);
				match(COUNT);
				 _localctx.e = expressions.NewCount((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 12:
				{
				setState(524);
				((ExprContext)_localctx).ID = match(ID);
				setState(525);
				match(PUNTO);
				setState(526);
				match(ISEMPTY);
				 _localctx.e = expressions.NewIsEmpty((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 13:
				{
				setState(528);
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
			case 14:
				{
				setState(530);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 15:
				{
				setState(532);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 16:
				{
				setState(534);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(575);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(573);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(538);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(539);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MUL) | (1L << DIV) | (1L << MODULO))) != 0)) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(540);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(23);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(543);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(544);
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
						setState(545);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(22);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(548);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(549);
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
						setState(550);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(21);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(553);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(554);
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
						setState(555);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(558);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(559);
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
						setState(560);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(563);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(564);
						((ExprContext)_localctx).op = match(AND);
						setState(565);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(568);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(569);
						((ExprContext)_localctx).op = match(OR);
						setState(570);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(577);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
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
		int _startState = 54;
		enterRecursionRule(_localctx, 54, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(579);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(589);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
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
					setState(582);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(583);
					match(COMA);
					setState(584);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(591);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
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
		public TerminalNode PUNTO() { return getToken(SwiftGrammarParser.PUNTO, 0); }
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
		int _startState = 56;
		enterRecursionRule(_localctx, 56, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(593);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVar((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(608);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(606);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
					case 1:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(596);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(597);
						match(CORIZQ);
						setState(598);
						((ListArrayContext)_localctx).expr = expr(0);
						setState(599);
						match(CORDER);
						 _localctx.p = expressions.NewArrayAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, ((ListArrayContext)_localctx).expr.e) 
						}
						break;
					case 2:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(602);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(603);
						match(PUNTO);
						setState(604);
						((ListArrayContext)_localctx).ID = match(ID);
						 _localctx.p = expressions.NewStructAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))  
						}
						break;
					}
					} 
				}
				setState(610);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
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

	public static class CallFunctionInstContext extends ParserRuleContext {
		public interfaces.Instruction cfi;
		public Token ID;
		public ListParamsContext listParams;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public CallFunctionInstContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callFunctionInst; }
	}

	public final CallFunctionInstContext callFunctionInst() throws RecognitionException {
		CallFunctionInstContext _localctx = new CallFunctionInstContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_callFunctionInst);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(611);
			((CallFunctionInstContext)_localctx).ID = match(ID);
			setState(612);
			match(PARIZQ);
			setState(613);
			((CallFunctionInstContext)_localctx).listParams = listParams(0);
			setState(614);
			match(PARDER);
			 _localctx.cfi = instructions.NewCall_Function((((CallFunctionInstContext)_localctx).ID!=null?((CallFunctionInstContext)_localctx).ID.getLine():0), (((CallFunctionInstContext)_localctx).ID!=null?((CallFunctionInstContext)_localctx).ID.getCharPositionInLine():0), (((CallFunctionInstContext)_localctx).ID!=null?((CallFunctionInstContext)_localctx).ID.getText():null), ((CallFunctionInstContext)_localctx).listParams.l) 
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

	public static class CallFunctionContext extends ParserRuleContext {
		public interfaces.Expression cf;
		public Token ID;
		public ListParamsCallContext listParamsCall;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(SwiftGrammarParser.PARIZQ, 0); }
		public ListParamsCallContext listParamsCall() {
			return getRuleContext(ListParamsCallContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(SwiftGrammarParser.PARDER, 0); }
		public CallFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callFunction; }
	}

	public final CallFunctionContext callFunction() throws RecognitionException {
		CallFunctionContext _localctx = new CallFunctionContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_callFunction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(617);
			((CallFunctionContext)_localctx).ID = match(ID);
			setState(618);
			match(PARIZQ);
			setState(619);
			((CallFunctionContext)_localctx).listParamsCall = listParamsCall(0);
			setState(620);
			match(PARDER);
			 _localctx.cf = expressions.NewCallExp((((CallFunctionContext)_localctx).ID!=null?((CallFunctionContext)_localctx).ID.getLine():0), (((CallFunctionContext)_localctx).ID!=null?((CallFunctionContext)_localctx).ID.getCharPositionInLine():0), (((CallFunctionContext)_localctx).ID!=null?((CallFunctionContext)_localctx).ID.getText():null), ((CallFunctionContext)_localctx).listParamsCall.l) 
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

	public static class ListParamsCallContext extends ParserRuleContext {
		public []interface{} l;
		public ListParamsCallContext list;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListParamsCallContext listParamsCall() {
			return getRuleContext(ListParamsCallContext.class,0);
		}
		public ListParamsCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParamsCall; }
	}

	public final ListParamsCallContext listParamsCall() throws RecognitionException {
		return listParamsCall(0);
	}

	private ListParamsCallContext listParamsCall(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsCallContext _localctx = new ListParamsCallContext(_ctx, _parentState);
		ListParamsCallContext _prevctx = _localctx;
		int _startState = 62;
		enterRecursionRule(_localctx, 62, RULE_listParamsCall, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(628);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				{
				setState(624);
				((ListParamsCallContext)_localctx).expr = expr(0);

				            _localctx.l = []interface{}{}
				            _localctx.l = append(_localctx.l, ((ListParamsCallContext)_localctx).expr.e)
				        
				}
				break;
			case 2:
				{

				        _localctx.l = []interface{}{}
				    
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(637);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsCallContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParamsCall);
					setState(630);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(631);
					match(COMA);
					setState(632);
					((ListParamsCallContext)_localctx).expr = expr(0);

					                                              var arr []interface{}
					                                              arr = append(((ListParamsCallContext)_localctx).list.l, ((ListParamsCallContext)_localctx).expr.e)
					                                              _localctx.l = arr
					                                          
					}
					} 
				}
				setState(639);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
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

	public static class ListStructExpContext extends ParserRuleContext {
		public []interface{} l;
		public ListStructExpContext list;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(SwiftGrammarParser.ID, 0); }
		public TerminalNode D_PTS() { return getToken(SwiftGrammarParser.D_PTS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
		public ListStructExpContext listStructExp() {
			return getRuleContext(ListStructExpContext.class,0);
		}
		public ListStructExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listStructExp; }
	}

	public final ListStructExpContext listStructExp() throws RecognitionException {
		return listStructExp(0);
	}

	private ListStructExpContext listStructExp(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListStructExpContext _localctx = new ListStructExpContext(_ctx, _parentState);
		ListStructExpContext _prevctx = _localctx;
		int _startState = 64;
		enterRecursionRule(_localctx, 64, RULE_listStructExp, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(647);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				{
				setState(641);
				((ListStructExpContext)_localctx).ID = match(ID);
				setState(642);
				match(D_PTS);
				setState(643);
				((ListStructExpContext)_localctx).expr = expr(0);

				                    var arr []interface{}
				                    StrExp := environment.NewStructContent((((ListStructExpContext)_localctx).ID!=null?((ListStructExpContext)_localctx).ID.getText():null), ((ListStructExpContext)_localctx).expr.e)
				                    arr = append(arr, StrExp)
				                    _localctx.l = arr
				                
				}
				break;
			case 2:
				{

				        _localctx.l = []interface{}{}
				    
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(658);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListStructExpContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listStructExp);
					setState(649);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(650);
					match(COMA);
					setState(651);
					((ListStructExpContext)_localctx).ID = match(ID);
					setState(652);
					match(D_PTS);
					setState(653);
					((ListStructExpContext)_localctx).expr = expr(0);

					                                                      var arr []interface{}
					                                                      StrExp := environment.NewStructContent((((ListStructExpContext)_localctx).ID!=null?((ListStructExpContext)_localctx).ID.getText():null), ((ListStructExpContext)_localctx).expr.e)
					                                                      arr = append(((ListStructExpContext)_localctx).list.l, StrExp)
					                                                      _localctx.l = arr
					                                                  
					}
					} 
				}
				setState(660);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
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
		case 5:
			return listParamsFunc_sempred((ListParamsFuncContext)_localctx, predIndex);
		case 7:
			return listStructDec_sempred((ListStructDecContext)_localctx, predIndex);
		case 17:
			return caselist_sempred((CaselistContext)_localctx, predIndex);
		case 20:
			return elseiflist_sempred((ElseiflistContext)_localctx, predIndex);
		case 26:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 27:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 28:
			return listArray_sempred((ListArrayContext)_localctx, predIndex);
		case 31:
			return listParamsCall_sempred((ListParamsCallContext)_localctx, predIndex);
		case 32:
			return listStructExp_sempred((ListStructExpContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listParamsFunc_sempred(ListParamsFuncContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean listStructDec_sempred(ListStructDecContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean caselist_sempred(CaselistContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean elseiflist_sempred(ElseiflistContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 22);
		case 5:
			return precpred(_ctx, 21);
		case 6:
			return precpred(_ctx, 20);
		case 7:
			return precpred(_ctx, 19);
		case 8:
			return precpred(_ctx, 18);
		case 9:
			return precpred(_ctx, 17);
		case 10:
			return precpred(_ctx, 16);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 11:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 12:
			return precpred(_ctx, 3);
		case 13:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listParamsCall_sempred(ListParamsCallContext _localctx, int predIndex) {
		switch (predIndex) {
		case 14:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean listStructExp_sempred(ListStructExpContext _localctx, int predIndex) {
		switch (predIndex) {
		case 15:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3?\u0298\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\3\2\3\2\3\2\3\2\3\3\6\3J\n\3\r\3\16\3K\3\3\3\3\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\u0083\n\4\3\5"+
		"\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u009f\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\5\7\u00a8\n\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\7\7\u00b1\n\7\f\7\16\7\u00b4"+
		"\13\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t"+
		"\u00c5\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\t\u00cf\n\t\f\t\16\t\u00d2"+
		"\13\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\16\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u0109\n\20\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\5\21\u012e\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\7\23\u0142\n\23\f\23"+
		"\16\23\u0145\13\23\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3"+
		"\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3"+
		"\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3"+
		"\25\3\25\3\25\3\25\3\25\3\25\5\25\u0173\n\25\3\26\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\7\26\u017d\n\26\f\26\16\26\u0180\13\26\3\27\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u01b9\n\31\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u01c6\n\32\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\5\33\u01d8\n\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\5\34"+
		"\u021b\n\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\7\34\u0240\n\34\f\34\16"+
		"\34\u0243\13\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\7\35\u024e"+
		"\n\35\f\35\16\35\u0251\13\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3"+
		"\36\3\36\3\36\3\36\3\36\3\36\7\36\u0261\n\36\f\36\16\36\u0264\13\36\3"+
		"\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\5!\u0277"+
		"\n!\3!\3!\3!\3!\3!\7!\u027e\n!\f!\16!\u0281\13!\3\"\3\"\3\"\3\"\3\"\3"+
		"\"\3\"\5\"\u028a\n\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\7\"\u0293\n\"\f\"\16"+
		"\"\u0296\13\"\3\"\2\13\f\20$*\668:@B#\2\4\6\b\n\f\16\20\22\24\26\30\32"+
		"\34\36 \"$&(*,.\60\62\64\668:<>@B\2\b\3\2!\"\4\2-.;;\3\2/\60\4\2))++\4"+
		"\2**,,\3\2#$\2\u02bc\2D\3\2\2\2\4I\3\2\2\2\6\u0082\3\2\2\2\b\u0084\3\2"+
		"\2\2\n\u009e\3\2\2\2\f\u00a7\3\2\2\2\16\u00b5\3\2\2\2\20\u00c4\3\2\2\2"+
		"\22\u00d3\3\2\2\2\24\u00dd\3\2\2\2\26\u00e4\3\2\2\2\30\u00ec\3\2\2\2\32"+
		"\u00f4\3\2\2\2\34\u00f7\3\2\2\2\36\u0108\3\2\2\2 \u012d\3\2\2\2\"\u012f"+
		"\3\2\2\2$\u0139\3\2\2\2&\u0146\3\2\2\2(\u0172\3\2\2\2*\u0174\3\2\2\2,"+
		"\u0181\3\2\2\2.\u0189\3\2\2\2\60\u01b8\3\2\2\2\62\u01c5\3\2\2\2\64\u01d7"+
		"\3\2\2\2\66\u021a\3\2\2\28\u0244\3\2\2\2:\u0252\3\2\2\2<\u0265\3\2\2\2"+
		">\u026b\3\2\2\2@\u0276\3\2\2\2B\u0289\3\2\2\2DE\5\4\3\2EF\7\2\2\3FG\b"+
		"\2\1\2G\3\3\2\2\2HJ\5\6\4\2IH\3\2\2\2JK\3\2\2\2KI\3\2\2\2KL\3\2\2\2LM"+
		"\3\2\2\2MN\b\3\1\2N\5\3\2\2\2OP\5\36\20\2PQ\b\4\1\2Q\u0083\3\2\2\2RS\5"+
		"(\25\2ST\b\4\1\2T\u0083\3\2\2\2UV\5\60\31\2VW\b\4\1\2W\u0083\3\2\2\2X"+
		"Y\5\62\32\2YZ\b\4\1\2Z\u0083\3\2\2\2[\\\5.\30\2\\]\b\4\1\2]\u0083\3\2"+
		"\2\2^_\5 \21\2_`\b\4\1\2`\u0083\3\2\2\2ab\5\"\22\2bc\b\4\1\2c\u0083\3"+
		"\2\2\2de\5\32\16\2ef\b\4\1\2f\u0083\3\2\2\2gh\5\34\17\2hi\b\4\1\2i\u0083"+
		"\3\2\2\2jk\5\30\r\2kl\b\4\1\2l\u0083\3\2\2\2mn\5\26\f\2no\b\4\1\2o\u0083"+
		"\3\2\2\2pq\5\24\13\2qr\b\4\1\2r\u0083\3\2\2\2st\5\22\n\2tu\b\4\1\2u\u0083"+
		"\3\2\2\2vw\5\n\6\2wx\b\4\1\2x\u0083\3\2\2\2yz\5<\37\2z{\b\4\1\2{\u0083"+
		"\3\2\2\2|}\5\16\b\2}~\b\4\1\2~\u0083\3\2\2\2\177\u0080\5\b\5\2\u0080\u0081"+
		"\b\4\1\2\u0081\u0083\3\2\2\2\u0082O\3\2\2\2\u0082R\3\2\2\2\u0082U\3\2"+
		"\2\2\u0082X\3\2\2\2\u0082[\3\2\2\2\u0082^\3\2\2\2\u0082a\3\2\2\2\u0082"+
		"d\3\2\2\2\u0082g\3\2\2\2\u0082j\3\2\2\2\u0082m\3\2\2\2\u0082p\3\2\2\2"+
		"\u0082s\3\2\2\2\u0082v\3\2\2\2\u0082y\3\2\2\2\u0082|\3\2\2\2\u0082\177"+
		"\3\2\2\2\u0083\7\3\2\2\2\u0084\u0085\7\26\2\2\u0085\u0086\5\66\34\2\u0086"+
		"\u0087\b\5\1\2\u0087\t\3\2\2\2\u0088\u0089\7\36\2\2\u0089\u008a\7\"\2"+
		"\2\u008a\u008b\7\61\2\2\u008b\u008c\5\f\7\2\u008c\u008d\7\62\2\2\u008d"+
		"\u008e\7\63\2\2\u008e\u008f\5\4\3\2\u008f\u0090\7\64\2\2\u0090\u0091\b"+
		"\6\1\2\u0091\u009f\3\2\2\2\u0092\u0093\7\36\2\2\u0093\u0094\7\"\2\2\u0094"+
		"\u0095\7\61\2\2\u0095\u0096\5\f\7\2\u0096\u0097\7\62\2\2\u0097\u0098\7"+
		"9\2\2\u0098\u0099\5\64\33\2\u0099\u009a\7\63\2\2\u009a\u009b\5\4\3\2\u009b"+
		"\u009c\7\64\2\2\u009c\u009d\b\6\1\2\u009d\u009f\3\2\2\2\u009e\u0088\3"+
		"\2\2\2\u009e\u0092\3\2\2\2\u009f\13\3\2\2\2\u00a0\u00a1\b\7\1\2\u00a1"+
		"\u00a2\7\"\2\2\u00a2\u00a3\7\65\2\2\u00a3\u00a4\5\64\33\2\u00a4\u00a5"+
		"\b\7\1\2\u00a5\u00a8\3\2\2\2\u00a6\u00a8\b\7\1\2\u00a7\u00a0\3\2\2\2\u00a7"+
		"\u00a6\3\2\2\2\u00a8\u00b2\3\2\2\2\u00a9\u00aa\f\5\2\2\u00aa\u00ab\78"+
		"\2\2\u00ab\u00ac\7\"\2\2\u00ac\u00ad\7\65\2\2\u00ad\u00ae\5\64\33\2\u00ae"+
		"\u00af\b\7\1\2\u00af\u00b1\3\2\2\2\u00b0\u00a9\3\2\2\2\u00b1\u00b4\3\2"+
		"\2\2\u00b2\u00b0\3\2\2\2\u00b2\u00b3\3\2\2\2\u00b3\r\3\2\2\2\u00b4\u00b2"+
		"\3\2\2\2\u00b5\u00b6\7\37\2\2\u00b6\u00b7\7\"\2\2\u00b7\u00b8\7\63\2\2"+
		"\u00b8\u00b9\5\20\t\2\u00b9\u00ba\7\64\2\2\u00ba\u00bb\b\b\1\2\u00bb\17"+
		"\3\2\2\2\u00bc\u00bd\b\t\1\2\u00bd\u00be\7\22\2\2\u00be\u00bf\7\"\2\2"+
		"\u00bf\u00c0\7\65\2\2\u00c0\u00c1\5\64\33\2\u00c1\u00c2\b\t\1\2\u00c2"+
		"\u00c5\3\2\2\2\u00c3\u00c5\b\t\1\2\u00c4\u00bc\3\2\2\2\u00c4\u00c3\3\2"+
		"\2\2\u00c5\u00d0\3\2\2\2\u00c6\u00c7\f\5\2\2\u00c7\u00c8\78\2\2\u00c8"+
		"\u00c9\7\22\2\2\u00c9\u00ca\7\"\2\2\u00ca\u00cb\7\65\2\2\u00cb\u00cc\5"+
		"\64\33\2\u00cc\u00cd\b\t\1\2\u00cd\u00cf\3\2\2\2\u00ce\u00c6\3\2\2\2\u00cf"+
		"\u00d2\3\2\2\2\u00d0\u00ce\3\2\2\2\u00d0\u00d1\3\2\2\2\u00d1\21\3\2\2"+
		"\2\u00d2\u00d0\3\2\2\2\u00d3\u00d4\7\"\2\2\u00d4\u00d5\7:\2\2\u00d5\u00d6"+
		"\7\32\2\2\u00d6\u00d7\7\61\2\2\u00d7\u00d8\7\35\2\2\u00d8\u00d9\7\65\2"+
		"\2\u00d9\u00da\5\66\34\2\u00da\u00db\7\62\2\2\u00db\u00dc\b\n\1\2\u00dc"+
		"\23\3\2\2\2\u00dd\u00de\7\"\2\2\u00de\u00df\7:\2\2\u00df\u00e0\7\31\2"+
		"\2\u00e0\u00e1\7\61\2\2\u00e1\u00e2\7\62\2\2\u00e2\u00e3\b\13\1\2\u00e3"+
		"\25\3\2\2\2\u00e4\u00e5\7\"\2\2\u00e5\u00e6\7:\2\2\u00e6\u00e7\7\30\2"+
		"\2\u00e7\u00e8\7\61\2\2\u00e8\u00e9\5\66\34\2\u00e9\u00ea\7\62\2\2\u00ea"+
		"\u00eb\b\f\1\2\u00eb\27\3\2\2\2\u00ec\u00ed\7\27\2\2\u00ed\u00ee\5\66"+
		"\34\2\u00ee\u00ef\7\13\2\2\u00ef\u00f0\7\63\2\2\u00f0\u00f1\5\4\3\2\u00f1"+
		"\u00f2\7\64\2\2\u00f2\u00f3\b\r\1\2\u00f3\31\3\2\2\2\u00f4\u00f5\7\24"+
		"\2\2\u00f5\u00f6\b\16\1\2\u00f6\33\3\2\2\2\u00f7\u00f8\7\25\2\2\u00f8"+
		"\u00f9\b\17\1\2\u00f9\35\3\2\2\2\u00fa\u00fb\7\t\2\2\u00fb\u00fc\7\61"+
		"\2\2\u00fc\u00fd\5\66\34\2\u00fd\u00fe\7\62\2\2\u00fe\u00ff\b\20\1\2\u00ff"+
		"\u0109\3\2\2\2\u0100\u0101\7\t\2\2\u0101\u0102\7\61\2\2\u0102\u0103\5"+
		"\66\34\2\u0103\u0104\78\2\2\u0104\u0105\5\66\34\2\u0105\u0106\7\62\2\2"+
		"\u0106\u0107\b\20\1\2\u0107\u0109\3\2\2\2\u0108\u00fa\3\2\2\2\u0108\u0100"+
		"\3\2\2\2\u0109\37\3\2\2\2\u010a\u010b\7\r\2\2\u010b\u010c\7\"\2\2\u010c"+
		"\u010d\7\21\2\2\u010d\u010e\5\66\34\2\u010e\u010f\7:\2\2\u010f\u0110\7"+
		":\2\2\u0110\u0111\7:\2\2\u0111\u0112\5\66\34\2\u0112\u0113\7\63\2\2\u0113"+
		"\u0114\5\4\3\2\u0114\u0115\7\64\2\2\u0115\u0116\b\21\1\2\u0116\u012e\3"+
		"\2\2\2\u0117\u0118\7\r\2\2\u0118\u0119\7<\2\2\u0119\u011a\7\21\2\2\u011a"+
		"\u011b\5\66\34\2\u011b\u011c\7:\2\2\u011c\u011d\7:\2\2\u011d\u011e\7:"+
		"\2\2\u011e\u011f\5\66\34\2\u011f\u0120\7\63\2\2\u0120\u0121\5\4\3\2\u0121"+
		"\u0122\7\64\2\2\u0122\u0123\b\21\1\2\u0123\u012e\3\2\2\2\u0124\u0125\7"+
		"\r\2\2\u0125\u0126\7\"\2\2\u0126\u0127\7\21\2\2\u0127\u0128\t\2\2\2\u0128"+
		"\u0129\7\63\2\2\u0129\u012a\5\4\3\2\u012a\u012b\7\64\2\2\u012b\u012c\b"+
		"\21\1\2\u012c\u012e\3\2\2\2\u012d\u010a\3\2\2\2\u012d\u0117\3\2\2\2\u012d"+
		"\u0124\3\2\2\2\u012e!\3\2\2\2\u012f\u0130\7\16\2\2\u0130\u0131\5\66\34"+
		"\2\u0131\u0132\7\63\2\2\u0132\u0133\5$\23\2\u0133\u0134\7\20\2\2\u0134"+
		"\u0135\7\65\2\2\u0135\u0136\5\4\3\2\u0136\u0137\7\64\2\2\u0137\u0138\b"+
		"\22\1\2\u0138#\3\2\2\2\u0139\u013a\b\23\1\2\u013a\u013b\5&\24\2\u013b"+
		"\u013c\b\23\1\2\u013c\u0143\3\2\2\2\u013d\u013e\f\4\2\2\u013e\u013f\5"+
		"&\24\2\u013f\u0140\b\23\1\2\u0140\u0142\3\2\2\2\u0141\u013d\3\2\2\2\u0142"+
		"\u0145\3\2\2\2\u0143\u0141\3\2\2\2\u0143\u0144\3\2\2\2\u0144%\3\2\2\2"+
		"\u0145\u0143\3\2\2\2\u0146\u0147\7\17\2\2\u0147\u0148\5\66\34\2\u0148"+
		"\u0149\7\65\2\2\u0149\u014a\5\4\3\2\u014a\u014b\b\24\1\2\u014b\'\3\2\2"+
		"\2\u014c\u014d\7\n\2\2\u014d\u014e\5\66\34\2\u014e\u014f\7\63\2\2\u014f"+
		"\u0150\5\4\3\2\u0150\u0151\7\64\2\2\u0151\u0152\b\25\1\2\u0152\u0173\3"+
		"\2\2\2\u0153\u0154\7\n\2\2\u0154\u0155\5\66\34\2\u0155\u0156\7\63\2\2"+
		"\u0156\u0157\5\4\3\2\u0157\u0158\7\64\2\2\u0158\u0159\7\13\2\2\u0159\u015a"+
		"\7\63\2\2\u015a\u015b\5\4\3\2\u015b\u015c\7\64\2\2\u015c\u015d\b\25\1"+
		"\2\u015d\u0173\3\2\2\2\u015e\u015f\7\n\2\2\u015f\u0160\5\66\34\2\u0160"+
		"\u0161\7\63\2\2\u0161\u0162\5\4\3\2\u0162\u0163\7\64\2\2\u0163\u0164\5"+
		"*\26\2\u0164\u0165\b\25\1\2\u0165\u0173\3\2\2\2\u0166\u0167\7\n\2\2\u0167"+
		"\u0168\5\66\34\2\u0168\u0169\7\63\2\2\u0169\u016a\5\4\3\2\u016a\u016b"+
		"\7\64\2\2\u016b\u016c\5*\26\2\u016c\u016d\7\13\2\2\u016d\u016e\7\63\2"+
		"\2\u016e\u016f\5\4\3\2\u016f\u0170\7\64\2\2\u0170\u0171\b\25\1\2\u0171"+
		"\u0173\3\2\2\2\u0172\u014c\3\2\2\2\u0172\u0153\3\2\2\2\u0172\u015e\3\2"+
		"\2\2\u0172\u0166\3\2\2\2\u0173)\3\2\2\2\u0174\u0175\b\26\1\2\u0175\u0176"+
		"\5,\27\2\u0176\u0177\b\26\1\2\u0177\u017e\3\2\2\2\u0178\u0179\f\4\2\2"+
		"\u0179\u017a\5,\27\2\u017a\u017b\b\26\1\2\u017b\u017d\3\2\2\2\u017c\u0178"+
		"\3\2\2\2\u017d\u0180\3\2\2\2\u017e\u017c\3\2\2\2\u017e\u017f\3\2\2\2\u017f"+
		"+\3\2\2\2\u0180\u017e\3\2\2\2\u0181\u0182\7\13\2\2\u0182\u0183\7\n\2\2"+
		"\u0183\u0184\5\66\34\2\u0184\u0185\7\63\2\2\u0185\u0186\5\4\3\2\u0186"+
		"\u0187\7\64\2\2\u0187\u0188\b\27\1\2\u0188-\3\2\2\2\u0189\u018a\7\f\2"+
		"\2\u018a\u018b\5\66\34\2\u018b\u018c\7\63\2\2\u018c\u018d\5\4\3\2\u018d"+
		"\u018e\7\64\2\2\u018e\u018f\b\30\1\2\u018f/\3\2\2\2\u0190\u0191\7\22\2"+
		"\2\u0191\u0192\7\"\2\2\u0192\u0193\7\65\2\2\u0193\u0194\5\64\33\2\u0194"+
		"\u0195\7(\2\2\u0195\u0196\5\66\34\2\u0196\u0197\b\31\1\2\u0197\u01b9\3"+
		"\2\2\2\u0198\u0199\7\22\2\2\u0199\u019a\7\"\2\2\u019a\u019b\7(\2\2\u019b"+
		"\u019c\5\66\34\2\u019c\u019d\b\31\1\2\u019d\u01b9\3\2\2\2\u019e\u019f"+
		"\7\23\2\2\u019f\u01a0\7\"\2\2\u01a0\u01a1\7(\2\2\u01a1\u01a2\5\66\34\2"+
		"\u01a2\u01a3\b\31\1\2\u01a3\u01b9\3\2\2\2\u01a4\u01a5\7\23\2\2\u01a5\u01a6"+
		"\7\"\2\2\u01a6\u01a7\7\65\2\2\u01a7\u01a8\5\64\33\2\u01a8\u01a9\7(\2\2"+
		"\u01a9\u01aa\5\66\34\2\u01aa\u01ab\b\31\1\2\u01ab\u01b9\3\2\2\2\u01ac"+
		"\u01ad\7\22\2\2\u01ad\u01ae\7\"\2\2\u01ae\u01af\7(\2\2\u01af\u01b0\5\66"+
		"\34\2\u01b0\u01b1\b\31\1\2\u01b1\u01b9\3\2\2\2\u01b2\u01b3\7\23\2\2\u01b3"+
		"\u01b4\7\"\2\2\u01b4\u01b5\7(\2\2\u01b5\u01b6\5\66\34\2\u01b6\u01b7\b"+
		"\31\1\2\u01b7\u01b9\3\2\2\2\u01b8\u0190\3\2\2\2\u01b8\u0198\3\2\2\2\u01b8"+
		"\u019e\3\2\2\2\u01b8\u01a4\3\2\2\2\u01b8\u01ac\3\2\2\2\u01b8\u01b2\3\2"+
		"\2\2\u01b9\61\3\2\2\2\u01ba\u01bb\7\"\2\2\u01bb\u01bc\7(\2\2\u01bc\u01bd"+
		"\5\66\34\2\u01bd\u01be\b\32\1\2\u01be\u01c6\3\2\2\2\u01bf\u01c0\7\"\2"+
		"\2\u01c0\u01c1\7/\2\2\u01c1\u01c2\7(\2\2\u01c2\u01c3\5\66\34\2\u01c3\u01c4"+
		"\b\32\1\2\u01c4\u01c6\3\2\2\2\u01c5\u01ba\3\2\2\2\u01c5\u01bf\3\2\2\2"+
		"\u01c6\63\3\2\2\2\u01c7\u01c8\7\3\2\2\u01c8\u01d8\b\33\1\2\u01c9\u01ca"+
		"\7\4\2\2\u01ca\u01d8\b\33\1\2\u01cb\u01cc\7\6\2\2\u01cc\u01d8\b\33\1\2"+
		"\u01cd\u01ce\7\5\2\2\u01ce\u01d8\b\33\1\2\u01cf\u01d0\7\66\2\2\u01d0\u01d1"+
		"\5\64\33\2\u01d1\u01d2\7\67\2\2\u01d2\u01d3\b\33\1\2\u01d3\u01d8\3\2\2"+
		"\2\u01d4\u01d5\7\66\2\2\u01d5\u01d6\7\67\2\2\u01d6\u01d8\b\33\1\2\u01d7"+
		"\u01c7\3\2\2\2\u01d7\u01c9\3\2\2\2\u01d7\u01cb\3\2\2\2\u01d7\u01cd\3\2"+
		"\2\2\u01d7\u01cf\3\2\2\2\u01d7\u01d4\3\2\2\2\u01d8\65\3\2\2\2\u01d9\u01da"+
		"\b\34\1\2\u01da\u01db\7\60\2\2\u01db\u01dc\5\66\34\31\u01dc\u01dd\b\34"+
		"\1\2\u01dd\u021b\3\2\2\2\u01de\u01df\7%\2\2\u01df\u01e0\5\66\34\21\u01e0"+
		"\u01e1\b\34\1\2\u01e1\u021b\3\2\2\2\u01e2\u01e3\5> \2\u01e3\u01e4\b\34"+
		"\1\2\u01e4\u021b\3\2\2\2\u01e5\u01e6\7\"\2\2\u01e6\u01e7\7\61\2\2\u01e7"+
		"\u01e8\5B\"\2\u01e8\u01e9\7\62\2\2\u01e9\u01ea\b\34\1\2\u01ea\u021b\3"+
		"\2\2\2\u01eb\u01ec\7\61\2\2\u01ec\u01ed\5\66\34\2\u01ed\u01ee\7\62\2\2"+
		"\u01ee\u01ef\b\34\1\2\u01ef\u021b\3\2\2\2\u01f0\u01f1\5:\36\2\u01f1\u01f2"+
		"\b\34\1\2\u01f2\u021b\3\2\2\2\u01f3\u01f4\7\66\2\2\u01f4\u01f5\58\35\2"+
		"\u01f5\u01f6\7\67\2\2\u01f6\u01f7\b\34\1\2\u01f7\u021b\3\2\2\2\u01f8\u01f9"+
		"\7\3\2\2\u01f9\u01fa\7\61\2\2\u01fa\u01fb\5\66\34\2\u01fb\u01fc\7\62\2"+
		"\2\u01fc\u01fd\b\34\1\2\u01fd\u021b\3\2\2\2\u01fe\u01ff\7\6\2\2\u01ff"+
		"\u0200\7\61\2\2\u0200\u0201\5\66\34\2\u0201\u0202\7\62\2\2\u0202\u0203"+
		"\b\34\1\2\u0203\u021b\3\2\2\2\u0204\u0205\7\4\2\2\u0205\u0206\7\61\2\2"+
		"\u0206\u0207\5\66\34\2\u0207\u0208\7\62\2\2\u0208\u0209\b\34\1\2\u0209"+
		"\u021b\3\2\2\2\u020a\u020b\7\"\2\2\u020b\u020c\7:\2\2\u020c\u020d\7\33"+
		"\2\2\u020d\u021b\b\34\1\2\u020e\u020f\7\"\2\2\u020f\u0210\7:\2\2\u0210"+
		"\u0211\7\34\2\2\u0211\u021b\b\34\1\2\u0212\u0213\7 \2\2\u0213\u021b\b"+
		"\34\1\2\u0214\u0215\7!\2\2\u0215\u021b\b\34\1\2\u0216\u0217\7\7\2\2\u0217"+
		"\u021b\b\34\1\2\u0218\u0219\7\b\2\2\u0219\u021b\b\34\1\2\u021a\u01d9\3"+
		"\2\2\2\u021a\u01de\3\2\2\2\u021a\u01e2\3\2\2\2\u021a\u01e5\3\2\2\2\u021a"+
		"\u01eb\3\2\2\2\u021a\u01f0\3\2\2\2\u021a\u01f3\3\2\2\2\u021a\u01f8\3\2"+
		"\2\2\u021a\u01fe\3\2\2\2\u021a\u0204\3\2\2\2\u021a\u020a\3\2\2\2\u021a"+
		"\u020e\3\2\2\2\u021a\u0212\3\2\2\2\u021a\u0214\3\2\2\2\u021a\u0216\3\2"+
		"\2\2\u021a\u0218\3\2\2\2\u021b\u0241\3\2\2\2\u021c\u021d\f\30\2\2\u021d"+
		"\u021e\t\3\2\2\u021e\u021f\5\66\34\31\u021f\u0220\b\34\1\2\u0220\u0240"+
		"\3\2\2\2\u0221\u0222\f\27\2\2\u0222\u0223\t\4\2\2\u0223\u0224\5\66\34"+
		"\30\u0224\u0225\b\34\1\2\u0225\u0240\3\2\2\2\u0226\u0227\f\26\2\2\u0227"+
		"\u0228\t\5\2\2\u0228\u0229\5\66\34\27\u0229\u022a\b\34\1\2\u022a\u0240"+
		"\3\2\2\2\u022b\u022c\f\25\2\2\u022c\u022d\t\6\2\2\u022d\u022e\5\66\34"+
		"\26\u022e\u022f\b\34\1\2\u022f\u0240\3\2\2\2\u0230\u0231\f\24\2\2\u0231"+
		"\u0232\t\7\2\2\u0232\u0233\5\66\34\25\u0233\u0234\b\34\1\2\u0234\u0240"+
		"\3\2\2\2\u0235\u0236\f\23\2\2\u0236\u0237\7\'\2\2\u0237\u0238\5\66\34"+
		"\24\u0238\u0239\b\34\1\2\u0239\u0240\3\2\2\2\u023a\u023b\f\22\2\2\u023b"+
		"\u023c\7&\2\2\u023c\u023d\5\66\34\23\u023d\u023e\b\34\1\2\u023e\u0240"+
		"\3\2\2\2\u023f\u021c\3\2\2\2\u023f\u0221\3\2\2\2\u023f\u0226\3\2\2\2\u023f"+
		"\u022b\3\2\2\2\u023f\u0230\3\2\2\2\u023f\u0235\3\2\2\2\u023f\u023a\3\2"+
		"\2\2\u0240\u0243\3\2\2\2\u0241\u023f\3\2\2\2\u0241\u0242\3\2\2\2\u0242"+
		"\67\3\2\2\2\u0243\u0241\3\2\2\2\u0244\u0245\b\35\1\2\u0245\u0246\5\66"+
		"\34\2\u0246\u0247\b\35\1\2\u0247\u024f\3\2\2\2\u0248\u0249\f\4\2\2\u0249"+
		"\u024a\78\2\2\u024a\u024b\5\66\34\2\u024b\u024c\b\35\1\2\u024c\u024e\3"+
		"\2\2\2\u024d\u0248\3\2\2\2\u024e\u0251\3\2\2\2\u024f\u024d\3\2\2\2\u024f"+
		"\u0250\3\2\2\2\u02509\3\2\2\2\u0251\u024f\3\2\2\2\u0252\u0253\b\36\1\2"+
		"\u0253\u0254\7\"\2\2\u0254\u0255\b\36\1\2\u0255\u0262\3\2\2\2\u0256\u0257"+
		"\f\5\2\2\u0257\u0258\7\66\2\2\u0258\u0259\5\66\34\2\u0259\u025a\7\67\2"+
		"\2\u025a\u025b\b\36\1\2\u025b\u0261\3\2\2\2\u025c\u025d\f\4\2\2\u025d"+
		"\u025e\7:\2\2\u025e\u025f\7\"\2\2\u025f\u0261\b\36\1\2\u0260\u0256\3\2"+
		"\2\2\u0260\u025c\3\2\2\2\u0261\u0264\3\2\2\2\u0262\u0260\3\2\2\2\u0262"+
		"\u0263\3\2\2\2\u0263;\3\2\2\2\u0264\u0262\3\2\2\2\u0265\u0266\7\"\2\2"+
		"\u0266\u0267\7\61\2\2\u0267\u0268\58\35\2\u0268\u0269\7\62\2\2\u0269\u026a"+
		"\b\37\1\2\u026a=\3\2\2\2\u026b\u026c\7\"\2\2\u026c\u026d\7\61\2\2\u026d"+
		"\u026e\5@!\2\u026e\u026f\7\62\2\2\u026f\u0270\b \1\2\u0270?\3\2\2\2\u0271"+
		"\u0272\b!\1\2\u0272\u0273\5\66\34\2\u0273\u0274\b!\1\2\u0274\u0277\3\2"+
		"\2\2\u0275\u0277\b!\1\2\u0276\u0271\3\2\2\2\u0276\u0275\3\2\2\2\u0277"+
		"\u027f\3\2\2\2\u0278\u0279\f\5\2\2\u0279\u027a\78\2\2\u027a\u027b\5\66"+
		"\34\2\u027b\u027c\b!\1\2\u027c\u027e\3\2\2\2\u027d\u0278\3\2\2\2\u027e"+
		"\u0281\3\2\2\2\u027f\u027d\3\2\2\2\u027f\u0280\3\2\2\2\u0280A\3\2\2\2"+
		"\u0281\u027f\3\2\2\2\u0282\u0283\b\"\1\2\u0283\u0284\7\"\2\2\u0284\u0285"+
		"\7\65\2\2\u0285\u0286\5\66\34\2\u0286\u0287\b\"\1\2\u0287\u028a\3\2\2"+
		"\2\u0288\u028a\b\"\1\2\u0289\u0282\3\2\2\2\u0289\u0288\3\2\2\2\u028a\u0294"+
		"\3\2\2\2\u028b\u028c\f\5\2\2\u028c\u028d\78\2\2\u028d\u028e\7\"\2\2\u028e"+
		"\u028f\7\65\2\2\u028f\u0290\5\66\34\2\u0290\u0291\b\"\1\2\u0291\u0293"+
		"\3\2\2\2\u0292\u028b\3\2\2\2\u0293\u0296\3\2\2\2\u0294\u0292\3\2\2\2\u0294"+
		"\u0295\3\2\2\2\u0295C\3\2\2\2\u0296\u0294\3\2\2\2\33K\u0082\u009e\u00a7"+
		"\u00b2\u00c4\u00d0\u0108\u012d\u0143\u0172\u017e\u01b8\u01c5\u01d7\u021a"+
		"\u023f\u0241\u024f\u0260\u0262\u0276\u027f\u0289\u0294";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}