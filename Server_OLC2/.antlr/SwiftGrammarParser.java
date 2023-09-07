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
		COUNT=25, ISEMPTY=26, FUNC=27, STRUCT=28, NUMBER=29, STRING=30, ID=31, 
		DIF=32, IG_IG=33, NOT=34, OR=35, AND=36, IG=37, MAY_IG=38, MEN_IG=39, 
		MAYOR=40, MENOR=41, MUL=42, DIV=43, ADD=44, SUB=45, PARIZQ=46, PARDER=47, 
		LLAVEIZQ=48, LLAVEDER=49, D_PTS=50, CORIZQ=51, CORDER=52, COMA=53, ARROW1=54, 
		PUNTO=55, MODULO=56, WHITESPACE=57, COMMENT=58, LINE_COMMENT=59;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_function = 3, RULE_listParamsFunc = 4, 
		RULE_structCreation = 5, RULE_listStructDec = 6, RULE_removetmt = 7, RULE_removelastmt = 8, 
		RULE_appendtmt = 9, RULE_guardtmt = 10, RULE_breaktmt = 11, RULE_continuetmt = 12, 
		RULE_printstmt = 13, RULE_forstmt = 14, RULE_switchtmt = 15, RULE_caselist = 16, 
		RULE_case = 17, RULE_ifstmt = 18, RULE_elseiflist = 19, RULE_elseif = 20, 
		RULE_whilestmt = 21, RULE_declarationstmt = 22, RULE_asignacion = 23, 
		RULE_types = 24, RULE_expr = 25, RULE_listParams = 26, RULE_listArray = 27, 
		RULE_callFunction = 28, RULE_listParamsCall = 29, RULE_listStructExp = 30;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "function", "listParamsFunc", "structCreation", 
			"listStructDec", "removetmt", "removelastmt", "appendtmt", "guardtmt", 
			"breaktmt", "continuetmt", "printstmt", "forstmt", "switchtmt", "caselist", 
			"case", "ifstmt", "elseiflist", "elseif", "whilestmt", "declarationstmt", 
			"asignacion", "types", "expr", "listParams", "listArray", "callFunction", 
			"listParamsCall", "listStructExp"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'true'", "'false'", 
			"'print'", "'if'", "'else'", "'while'", "'for'", "'switch'", "'case'", 
			"'default'", "'in'", "'var'", "'let'", "'break'", "'continue'", "'return'", 
			"'guard'", "'append'", "'removeLast'", "'remove'", "'count'", "'IsEmpty'", 
			"'func'", "'struct'", null, null, null, "'!='", "'=='", "'!'", "'||'", 
			"'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", 
			"'('", "')'", "'{'", "'}'", "':'", "'['", "']'", "','", "'->'", "'.'", 
			"'%'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE", 
			"WHILE", "FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "LET", "BREAK", 
			"CONTINUE", "RETURN", "GUARD", "APPEND", "REMOVELAST", "REMOVE", "COUNT", 
			"ISEMPTY", "FUNC", "STRUCT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", 
			"NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", 
			"CORIZQ", "CORDER", "COMA", "ARROW1", "PUNTO", "MODULO", "WHITESPACE", 
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
			setState(62);
			((SContext)_localctx).block = block();
			setState(63);
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
			setState(67); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(66);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(69); 
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
		public StructCreationContext structCreation;
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
		public StructCreationContext structCreation() {
			return getRuleContext(StructCreationContext.class,0);
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
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(73);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(76);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ifinst 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(79);
				((InstructionContext)_localctx).declarationstmt = declarationstmt();
				 _localctx.inst = ((InstructionContext)_localctx).declarationstmt.dec 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(82);
				((InstructionContext)_localctx).asignacion = asignacion();
				 _localctx.inst = ((InstructionContext)_localctx).asignacion.asign
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(85);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				 _localctx.inst = ((InstructionContext)_localctx).whilestmt.whiles 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(88);
				((InstructionContext)_localctx).forstmt = forstmt();
				 _localctx.inst = ((InstructionContext)_localctx).forstmt.fors 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(91);
				((InstructionContext)_localctx).switchtmt = switchtmt();
				 _localctx.inst = ((InstructionContext)_localctx).switchtmt.swtch 
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(94);
				((InstructionContext)_localctx).breaktmt = breaktmt();
				 _localctx.inst = ((InstructionContext)_localctx).breaktmt.break 
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(97);
				((InstructionContext)_localctx).continuetmt = continuetmt();
				 _localctx.inst = ((InstructionContext)_localctx).continuetmt.continue 
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(100);
				((InstructionContext)_localctx).guardtmt = guardtmt();
				 _localctx.inst = ((InstructionContext)_localctx).guardtmt.guard 
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(103);
				((InstructionContext)_localctx).appendtmt = appendtmt();
				 _localctx.inst = ((InstructionContext)_localctx).appendtmt.append
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(106);
				((InstructionContext)_localctx).removelastmt = removelastmt();
				 _localctx.inst = ((InstructionContext)_localctx).removelastmt.removelast 
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(109);
				((InstructionContext)_localctx).removetmt = removetmt();
				 _localctx.inst = ((InstructionContext)_localctx).removetmt.remove 
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(112);
				((InstructionContext)_localctx).function = function();
				 _localctx.inst = ((InstructionContext)_localctx).function.fun 
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(115);
				((InstructionContext)_localctx).structCreation = structCreation();
				 _localctx.inst = ((InstructionContext)_localctx).structCreation.dec 
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
		enterRule(_localctx, 6, RULE_function);
		try {
			setState(142);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				((FunctionContext)_localctx).FUNC = match(FUNC);
				setState(121);
				((FunctionContext)_localctx).ID = match(ID);
				setState(122);
				match(PARIZQ);
				setState(123);
				((FunctionContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(124);
				match(PARDER);
				setState(125);
				match(LLAVEIZQ);
				setState(126);
				((FunctionContext)_localctx).block = block();
				setState(127);
				match(LLAVEDER);

				    _localctx.fun = instructions.NewFunction((((FunctionContext)_localctx).FUNC!=null?((FunctionContext)_localctx).FUNC.getLine():0), (((FunctionContext)_localctx).FUNC!=null?((FunctionContext)_localctx).FUNC.getCharPositionInLine():0), (((FunctionContext)_localctx).ID!=null?((FunctionContext)_localctx).ID.getText():null), ((FunctionContext)_localctx).listParamsFunc.lpf, environment.NULL, ((FunctionContext)_localctx).block.blk)

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(130);
				((FunctionContext)_localctx).FUNC = match(FUNC);
				setState(131);
				((FunctionContext)_localctx).ID = match(ID);
				setState(132);
				match(PARIZQ);
				setState(133);
				((FunctionContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(134);
				match(PARDER);
				setState(135);
				match(ARROW1);
				setState(136);
				((FunctionContext)_localctx).types = types();
				setState(137);
				match(LLAVEIZQ);
				setState(138);
				((FunctionContext)_localctx).block = block();
				setState(139);
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
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_listParamsFunc, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(145);
				((ListParamsFuncContext)_localctx).ID = match(ID);
				setState(146);
				match(D_PTS);
				setState(147);
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
			setState(162);
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
					setState(153);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(154);
					match(COMA);
					setState(155);
					((ListParamsFuncContext)_localctx).ID = match(ID);
					setState(156);
					match(D_PTS);
					setState(157);
					((ListParamsFuncContext)_localctx).types = types();

					              var arr []interface{}
					              newParam := instructions.NewParamsDeclaration((((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getLine():0), (((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getCharPositionInLine():0), (((ListParamsFuncContext)_localctx).ID!=null?((ListParamsFuncContext)_localctx).ID.getText():null), ((ListParamsFuncContext)_localctx).types.ty)
					              arr = append(((ListParamsFuncContext)_localctx).list.lpf, newParam)
					              _localctx.lpf = arr
					              
					}
					} 
				}
				setState(164);
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
		enterRule(_localctx, 10, RULE_structCreation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			((StructCreationContext)_localctx).STRUCT = match(STRUCT);
			setState(166);
			((StructCreationContext)_localctx).ID = match(ID);
			setState(167);
			match(LLAVEIZQ);
			setState(168);
			((StructCreationContext)_localctx).listStructDec = listStructDec(0);
			setState(169);
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
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_listStructDec, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(173);
				match(VAR);
				setState(174);
				((ListStructDecContext)_localctx).ID = match(ID);
				setState(175);
				match(D_PTS);
				setState(176);
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
			setState(192);
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
					setState(182);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(183);
					match(COMA);
					setState(184);
					match(VAR);
					setState(185);
					((ListStructDecContext)_localctx).ID = match(ID);
					setState(186);
					match(D_PTS);
					setState(187);
					((ListStructDecContext)_localctx).types = types();

					                                                          var arr []interface{}
					                                                          newParams := environment.NewStructType((((ListStructDecContext)_localctx).ID!=null?((ListStructDecContext)_localctx).ID.getText():null), ((ListStructDecContext)_localctx).types.ty)
					                                                          arr = append(((ListStructDecContext)_localctx).list.l, newParams)
					                                                          _localctx.l = arr
					                                                      
					}
					} 
				}
				setState(194);
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
		enterRule(_localctx, 14, RULE_removetmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(195);
			((RemovetmtContext)_localctx).ID = match(ID);
			setState(196);
			match(PUNTO);
			setState(197);
			match(REMOVE);
			setState(198);
			match(PARIZQ);
			setState(199);
			((RemovetmtContext)_localctx).expr = expr(0);
			setState(200);
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
		enterRule(_localctx, 16, RULE_removelastmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			((RemovelastmtContext)_localctx).ID = match(ID);
			setState(204);
			match(PUNTO);
			setState(205);
			match(REMOVELAST);
			setState(206);
			match(PARIZQ);
			setState(207);
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
		enterRule(_localctx, 18, RULE_appendtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(210);
			((AppendtmtContext)_localctx).ID = match(ID);
			setState(211);
			match(PUNTO);
			setState(212);
			match(APPEND);
			setState(213);
			match(PARIZQ);
			setState(214);
			((AppendtmtContext)_localctx).expr = expr(0);
			setState(215);
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
		enterRule(_localctx, 20, RULE_guardtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(218);
			((GuardtmtContext)_localctx).GUARD = match(GUARD);
			setState(219);
			((GuardtmtContext)_localctx).expr = expr(0);
			setState(220);
			match(ELSE);
			setState(221);
			match(LLAVEIZQ);
			setState(222);
			((GuardtmtContext)_localctx).block = block();
			setState(223);
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
		enterRule(_localctx, 22, RULE_breaktmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
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
		enterRule(_localctx, 24, RULE_continuetmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(229);
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
		enterRule(_localctx, 26, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(232);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(233);
			match(PARIZQ);
			setState(234);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(235);
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
		enterRule(_localctx, 28, RULE_forstmt);
		int _la;
		try {
			setState(260);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(238);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(239);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(240);
				match(IN);
				setState(241);
				((ForstmtContext)_localctx).e1 = expr(0);
				setState(242);
				match(PUNTO);
				setState(243);
				match(PUNTO);
				setState(244);
				match(PUNTO);
				setState(245);
				((ForstmtContext)_localctx).e2 = expr(0);
				setState(246);
				match(LLAVEIZQ);
				setState(247);
				((ForstmtContext)_localctx).block = block();
				setState(248);
				match(LLAVEDER);
				 _localctx.fors = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).e1.e,((ForstmtContext)_localctx).e2.e,"nil", ((ForstmtContext)_localctx).block.blk) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(251);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(252);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(253);
				match(IN);
				setState(254);
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
				setState(255);
				match(LLAVEIZQ);
				setState(256);
				((ForstmtContext)_localctx).block = block();
				setState(257);
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
		enterRule(_localctx, 30, RULE_switchtmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			((SwitchtmtContext)_localctx).SWITCH = match(SWITCH);
			setState(263);
			((SwitchtmtContext)_localctx).expr = expr(0);
			setState(264);
			match(LLAVEIZQ);
			setState(265);
			((SwitchtmtContext)_localctx).caselist = caselist(0);
			setState(266);
			match(DEFAULT);
			setState(267);
			match(D_PTS);
			setState(268);
			((SwitchtmtContext)_localctx).block = block();
			setState(269);
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
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_caselist, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(273);
			((CaselistContext)_localctx).case = case();

			                            _localctx.cases= []interface{}{}
			                            _localctx.cases = append(_localctx.cases, ((CaselistContext)_localctx).case.swtch)
			                        
			}
			_ctx.stop = _input.LT(-1);
			setState(282);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
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
					setState(276);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(277);
					((CaselistContext)_localctx).case = case();
					 var arr3 []interface{}
					                                  arr3= append(((CaselistContext)_localctx).listcases.cases, ((CaselistContext)_localctx).case.swtch)
					                                  _localctx.cases= arr3
					                                  
					}
					} 
				}
				setState(284);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
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
		enterRule(_localctx, 34, RULE_case);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			((CaseContext)_localctx).CASE = match(CASE);
			setState(286);
			((CaseContext)_localctx).expr = expr(0);
			setState(287);
			match(D_PTS);
			setState(288);
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
		enterRule(_localctx, 36, RULE_ifstmt);
		try {
			setState(329);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(291);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(292);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(293);
				match(LLAVEIZQ);
				setState(294);
				((IfstmtContext)_localctx).block = block();
				setState(295);
				match(LLAVEDER);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, nil, nil) 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(298);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(299);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(300);
				match(LLAVEIZQ);
				setState(301);
				((IfstmtContext)_localctx).b1 = block();
				setState(302);
				match(LLAVEDER);
				setState(303);
				match(ELSE);
				setState(304);
				match(LLAVEIZQ);
				setState(305);
				((IfstmtContext)_localctx).b2 = block();
				setState(306);
				match(LLAVEDER);
				_localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).b1.blk, nil, ((IfstmtContext)_localctx).b2.blk) 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(309);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(310);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(311);
				match(LLAVEIZQ);
				setState(312);
				((IfstmtContext)_localctx).block = block();
				setState(313);
				match(LLAVEDER);
				setState(314);
				((IfstmtContext)_localctx).elseiflist = elseiflist(0);
				 _localctx.ifinst = instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).block.blk, ((IfstmtContext)_localctx).elseiflist.else_if_list, nil) 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(317);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(318);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(319);
				match(LLAVEIZQ);
				setState(320);
				((IfstmtContext)_localctx).b1 = block();
				setState(321);
				match(LLAVEDER);
				setState(322);
				((IfstmtContext)_localctx).elseiflist = elseiflist(0);
				setState(323);
				match(ELSE);
				setState(324);
				match(LLAVEIZQ);
				setState(325);
				((IfstmtContext)_localctx).b2 = block();
				setState(326);
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
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_elseiflist, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(332);
			((ElseiflistContext)_localctx).elseif = elseif();

			                _localctx.else_if_list = []interface{}{}
			                _localctx.else_if_list = append(_localctx.else_if_list, ((ElseiflistContext)_localctx).elseif.ifinst)

			            
			}
			_ctx.stop = _input.LT(-1);
			setState(341);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
					setState(335);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(336);
					((ElseiflistContext)_localctx).elseif = elseif();
					   var arr2 []interface{}
					                                  arr2 = append(((ElseiflistContext)_localctx).listif.else_if_list, ((ElseiflistContext)_localctx).elseif.ifinst)
					                                  _localctx.else_if_list = arr2
					                              
					}
					} 
				}
				setState(343);
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
		enterRule(_localctx, 40, RULE_elseif);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(344);
			match(ELSE);
			setState(345);
			((ElseifContext)_localctx).IF = match(IF);
			setState(346);
			((ElseifContext)_localctx).expr = expr(0);
			setState(347);
			match(LLAVEIZQ);
			setState(348);
			((ElseifContext)_localctx).block = block();
			setState(349);
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
		enterRule(_localctx, 42, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(352);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(353);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(354);
			match(LLAVEIZQ);
			setState(355);
			((WhilestmtContext)_localctx).block = block();
			setState(356);
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
		enterRule(_localctx, 44, RULE_declarationstmt);
		try {
			setState(387);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(359);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(360);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(361);
				match(D_PTS);
				setState(362);
				((DeclarationstmtContext)_localctx).types = types();
				setState(363);
				match(IG);
				setState(364);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, false);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(367);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(368);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(369);
				match(IG);
				setState(370);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.NULL, ((DeclarationstmtContext)_localctx).expr.e, false);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(373);
				((DeclarationstmtContext)_localctx).LET = match(LET);
				setState(374);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(375);
				match(D_PTS);
				setState(376);
				((DeclarationstmtContext)_localctx).types = types();
				setState(377);
				match(IG);
				setState(378);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				((DeclarationstmtContext)_localctx).dec =  instructions.NewDeclaration((((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getLine():0), (((DeclarationstmtContext)_localctx).LET!=null?((DeclarationstmtContext)_localctx).LET.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), ((DeclarationstmtContext)_localctx).types.ty, ((DeclarationstmtContext)_localctx).expr.e, true);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(381);
				((DeclarationstmtContext)_localctx).VAR = match(VAR);
				setState(382);
				((DeclarationstmtContext)_localctx).ID = match(ID);
				setState(383);
				match(IG);
				setState(384);
				((DeclarationstmtContext)_localctx).expr = expr(0);
				 _localctx.dec = instructions.NewDeclaration((((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getLine():0), (((DeclarationstmtContext)_localctx).VAR!=null?((DeclarationstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarationstmtContext)_localctx).ID!=null?((DeclarationstmtContext)_localctx).ID.getText():null), environment.STRUCT , ((DeclarationstmtContext)_localctx).expr.e, false) 
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
		enterRule(_localctx, 46, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(389);
			((AsignacionContext)_localctx).ID = match(ID);
			setState(390);
			match(IG);
			setState(391);
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
		enterRule(_localctx, 48, RULE_types);
		try {
			setState(410);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(394);
				match(INT);
				 _localctx.ty = environment.INTEGER 
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(396);
				match(FLOAT);
				 _localctx.ty = environment.FLOAT 
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(398);
				match(STR);
				 _localctx.ty = environment.STRING 
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(400);
				match(BOOL);
				 _localctx.ty = environment.BOOLEAN 
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(402);
				match(CORIZQ);
				setState(403);
				types();
				setState(404);
				match(CORDER);
				 _localctx.ty = environment.ARRAY 
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(407);
				match(CORIZQ);
				setState(408);
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
		public TerminalNode COMA() { return getToken(SwiftGrammarParser.COMA, 0); }
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
		int _startState = 50;
		enterRecursionRule(_localctx, 50, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(481);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				setState(413);
				((ExprContext)_localctx).op = match(SUB);
				setState(414);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(23);
				 _localctx.e = expressions.NewOperation(((ExprContext)_localctx).op.GetLine(), ((ExprContext)_localctx).op.GetColumn(), ((ExprContext)_localctx).right.e, "neg", ((ExprContext)_localctx).right.e) 
				}
				break;
			case 2:
				{
				setState(417);
				((ExprContext)_localctx).op = match(NOT);
				setState(418);
				((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
				 _localctx.e = expressions.NewOperation(((ExprContext)_localctx).op.GetLine(), ((ExprContext)_localctx).op.GetColumn(),((ExprContext)_localctx).right.e , (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
				}
				break;
			case 3:
				{
				setState(421);
				((ExprContext)_localctx).callFunction = callFunction();
				 _localctx.e = ((ExprContext)_localctx).callFunction.cf 
				}
				break;
			case 4:
				{
				setState(424);
				((ExprContext)_localctx).ID = match(ID);
				setState(425);
				match(PARIZQ);
				setState(426);
				((ExprContext)_localctx).listStructExp = listStructExp(0);
				setState(427);
				match(PARDER);
				 _localctx.e = expressions.NewStructExp((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null), ((ExprContext)_localctx).listStructExp.l ) 
				}
				break;
			case 5:
				{
				setState(430);
				match(PARIZQ);
				setState(431);
				((ExprContext)_localctx).expr = expr(0);
				setState(432);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 6:
				{
				setState(435);
				((ExprContext)_localctx).list = listArray(0);
				 _localctx.e = ((ExprContext)_localctx).list.p
				}
				break;
			case 7:
				{
				setState(438);
				((ExprContext)_localctx).CORIZQ = match(CORIZQ);
				setState(439);
				((ExprContext)_localctx).listParams = listParams(0);
				setState(440);
				match(CORDER);
				 _localctx.e = expressions.NewArray((((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getLine():0), (((ExprContext)_localctx).CORIZQ!=null?((ExprContext)_localctx).CORIZQ.getCharPositionInLine():0), ((ExprContext)_localctx).listParams.l) 
				}
				break;
			case 8:
				{
				setState(443);
				((ExprContext)_localctx).INT = match(INT);
				setState(444);
				match(PARIZQ);
				setState(445);
				((ExprContext)_localctx).expr = expr(0);
				setState(446);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).INT.GetLine(),((ExprContext)_localctx).INT.GetColumn(),((ExprContext)_localctx).expr.e, environment.INTEGER) 
				}
				break;
			case 9:
				{
				setState(449);
				((ExprContext)_localctx).STR = match(STR);
				setState(450);
				match(PARIZQ);
				setState(451);
				((ExprContext)_localctx).expr = expr(0);
				setState(452);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).STR.GetLine(),((ExprContext)_localctx).STR.GetColumn(),((ExprContext)_localctx).expr.e, environment.STRING) 
				}
				break;
			case 10:
				{
				setState(455);
				((ExprContext)_localctx).FLOAT = match(FLOAT);
				setState(456);
				match(PARIZQ);
				setState(457);
				((ExprContext)_localctx).expr = expr(0);
				setState(458);
				match(PARDER);
				 _localctx.e = expressions.NewConversion(((ExprContext)_localctx).FLOAT.GetLine(),((ExprContext)_localctx).FLOAT.GetColumn(),((ExprContext)_localctx).expr.e, environment.FLOAT) 
				}
				break;
			case 11:
				{
				setState(461);
				((ExprContext)_localctx).ID = match(ID);
				setState(462);
				match(PUNTO);
				setState(463);
				match(COUNT);
				setState(464);
				match(PARIZQ);
				setState(465);
				match(PARDER);
				 _localctx.e = expressions.NewCount((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 12:
				{
				setState(467);
				((ExprContext)_localctx).ID = match(ID);
				setState(468);
				match(PUNTO);
				setState(469);
				match(ISEMPTY);
				setState(470);
				match(PARIZQ);
				setState(471);
				match(PARDER);
				 _localctx.e = expressions.NewIsEmpty((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null)) 
				}
				break;
			case 13:
				{
				setState(473);
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
				setState(475);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 15:
				{
				setState(477);
				((ExprContext)_localctx).TRU = match(TRU);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getLine():0), (((ExprContext)_localctx).TRU!=null?((ExprContext)_localctx).TRU.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 16:
				{
				setState(479);
				((ExprContext)_localctx).FAL = match(FAL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getLine():0), (((ExprContext)_localctx).FAL!=null?((ExprContext)_localctx).FAL.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(520);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(518);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(483);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(484);
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
						setState(485);
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
						setState(488);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(489);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ADD) | (1L << SUB) | (1L << COMA))) != 0)) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(490);
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
						setState(493);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(494);
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
						setState(495);
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
						setState(498);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(499);
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
						setState(500);
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
						setState(503);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(504);
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
						setState(505);
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
						setState(508);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(509);
						((ExprContext)_localctx).op = match(AND);
						setState(510);
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
						setState(513);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(514);
						((ExprContext)_localctx).op = match(OR);
						setState(515);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(522);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
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
		int _startState = 52;
		enterRecursionRule(_localctx, 52, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(524);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(534);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
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
					setState(527);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(528);
					match(COMA);
					setState(529);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(536);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
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
		int _startState = 54;
		enterRecursionRule(_localctx, 54, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(538);
			((ListArrayContext)_localctx).ID = match(ID);
			 _localctx.p = expressions.NewCallVar((((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getCharPositionInLine():0), (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))
			}
			_ctx.stop = _input.LT(-1);
			setState(553);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(551);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
					case 1:
						{
						_localctx = new ListArrayContext(_parentctx, _parentState);
						_localctx.list = _prevctx;
						_localctx.list = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_listArray);
						setState(541);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(542);
						match(CORIZQ);
						setState(543);
						((ListArrayContext)_localctx).expr = expr(0);
						setState(544);
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
						setState(547);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(548);
						match(PUNTO);
						setState(549);
						((ListArrayContext)_localctx).ID = match(ID);
						 _localctx.p = expressions.NewStructAccess((((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetLine(), (((ListArrayContext)_localctx).list!=null?(((ListArrayContext)_localctx).list.start):null).GetColumn(), ((ListArrayContext)_localctx).list.p, (((ListArrayContext)_localctx).ID!=null?((ListArrayContext)_localctx).ID.getText():null))  
						}
						break;
					}
					} 
				}
				setState(555);
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
		enterRule(_localctx, 56, RULE_callFunction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(556);
			((CallFunctionContext)_localctx).ID = match(ID);
			setState(557);
			match(PARIZQ);
			setState(558);
			((CallFunctionContext)_localctx).listParamsCall = listParamsCall(0);
			setState(559);
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
		int _startState = 58;
		enterRecursionRule(_localctx, 58, RULE_listParamsCall, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(567);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				{
				setState(563);
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
			setState(576);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,20,_ctx);
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
					setState(569);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(570);
					match(COMA);
					setState(571);
					((ListParamsCallContext)_localctx).expr = expr(0);

					                                              var arr []interface{}
					                                              arr = append(((ListParamsCallContext)_localctx).list.l, ((ListParamsCallContext)_localctx).expr.e)
					                                              _localctx.l = arr
					                                          
					}
					} 
				}
				setState(578);
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
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_listStructExp, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(586);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				{
				setState(580);
				((ListStructExpContext)_localctx).ID = match(ID);
				setState(581);
				match(D_PTS);
				setState(582);
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
			setState(597);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
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
					setState(588);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(589);
					match(COMA);
					setState(590);
					((ListStructExpContext)_localctx).ID = match(ID);
					setState(591);
					match(D_PTS);
					setState(592);
					((ListStructExpContext)_localctx).expr = expr(0);

					                                                      var arr []interface{}
					                                                      StrExp := environment.NewStructContent((((ListStructExpContext)_localctx).ID!=null?((ListStructExpContext)_localctx).ID.getText():null), ((ListStructExpContext)_localctx).expr.e)
					                                                      arr = append(((ListStructExpContext)_localctx).list.l, StrExp)
					                                                      _localctx.l = arr
					                                                  
					}
					} 
				}
				setState(599);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return listParamsFunc_sempred((ListParamsFuncContext)_localctx, predIndex);
		case 6:
			return listStructDec_sempred((ListStructDecContext)_localctx, predIndex);
		case 16:
			return caselist_sempred((CaselistContext)_localctx, predIndex);
		case 19:
			return elseiflist_sempred((ElseiflistContext)_localctx, predIndex);
		case 25:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 26:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 27:
			return listArray_sempred((ListArrayContext)_localctx, predIndex);
		case 29:
			return listParamsCall_sempred((ListParamsCallContext)_localctx, predIndex);
		case 30:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3=\u025b\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \3\2"+
		"\3\2\3\2\3\2\3\3\6\3F\n\3\r\3\16\3G\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\5\4y\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5\u0091\n\5\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\5\6\u009a\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\7\6\u00a3\n\6"+
		"\f\6\16\6\u00a6\13\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\5\b\u00b7\n\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\7\b\u00c1\n\b"+
		"\f\b\16\b\u00c4\13\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u0107\n\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\7\22\u011b\n\22\f\22\16\22\u011e\13\22\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u014c\n\24"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\7\25\u0156\n\25\f\25\16\25\u0159"+
		"\13\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\5\30\u0186\n\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u019d"+
		"\n\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\5\33\u01e4\n\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\7\33\u0209\n\33\f\33"+
		"\16\33\u020c\13\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\7\34\u0217"+
		"\n\34\f\34\16\34\u021a\13\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3"+
		"\35\3\35\3\35\3\35\3\35\3\35\7\35\u022a\n\35\f\35\16\35\u022d\13\35\3"+
		"\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\5\37\u023a\n\37"+
		"\3\37\3\37\3\37\3\37\3\37\7\37\u0241\n\37\f\37\16\37\u0244\13\37\3 \3"+
		" \3 \3 \3 \3 \3 \5 \u024d\n \3 \3 \3 \3 \3 \3 \3 \7 \u0256\n \f \16 \u0259"+
		"\13 \3 \2\13\n\16\"(\64\668<>!\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36"+
		" \"$&(*,.\60\62\64\668:<>\2\b\3\2 !\4\2,-::\4\2./\67\67\4\2((**\4\2))"+
		"++\3\2\"#\2\u027a\2@\3\2\2\2\4E\3\2\2\2\6x\3\2\2\2\b\u0090\3\2\2\2\n\u0099"+
		"\3\2\2\2\f\u00a7\3\2\2\2\16\u00b6\3\2\2\2\20\u00c5\3\2\2\2\22\u00cd\3"+
		"\2\2\2\24\u00d4\3\2\2\2\26\u00dc\3\2\2\2\30\u00e4\3\2\2\2\32\u00e7\3\2"+
		"\2\2\34\u00ea\3\2\2\2\36\u0106\3\2\2\2 \u0108\3\2\2\2\"\u0112\3\2\2\2"+
		"$\u011f\3\2\2\2&\u014b\3\2\2\2(\u014d\3\2\2\2*\u015a\3\2\2\2,\u0162\3"+
		"\2\2\2.\u0185\3\2\2\2\60\u0187\3\2\2\2\62\u019c\3\2\2\2\64\u01e3\3\2\2"+
		"\2\66\u020d\3\2\2\28\u021b\3\2\2\2:\u022e\3\2\2\2<\u0239\3\2\2\2>\u024c"+
		"\3\2\2\2@A\5\4\3\2AB\7\2\2\3BC\b\2\1\2C\3\3\2\2\2DF\5\6\4\2ED\3\2\2\2"+
		"FG\3\2\2\2GE\3\2\2\2GH\3\2\2\2HI\3\2\2\2IJ\b\3\1\2J\5\3\2\2\2KL\5\34\17"+
		"\2LM\b\4\1\2My\3\2\2\2NO\5&\24\2OP\b\4\1\2Py\3\2\2\2QR\5.\30\2RS\b\4\1"+
		"\2Sy\3\2\2\2TU\5\60\31\2UV\b\4\1\2Vy\3\2\2\2WX\5,\27\2XY\b\4\1\2Yy\3\2"+
		"\2\2Z[\5\36\20\2[\\\b\4\1\2\\y\3\2\2\2]^\5 \21\2^_\b\4\1\2_y\3\2\2\2`"+
		"a\5\30\r\2ab\b\4\1\2by\3\2\2\2cd\5\32\16\2de\b\4\1\2ey\3\2\2\2fg\5\26"+
		"\f\2gh\b\4\1\2hy\3\2\2\2ij\5\24\13\2jk\b\4\1\2ky\3\2\2\2lm\5\22\n\2mn"+
		"\b\4\1\2ny\3\2\2\2op\5\20\t\2pq\b\4\1\2qy\3\2\2\2rs\5\b\5\2st\b\4\1\2"+
		"ty\3\2\2\2uv\5\f\7\2vw\b\4\1\2wy\3\2\2\2xK\3\2\2\2xN\3\2\2\2xQ\3\2\2\2"+
		"xT\3\2\2\2xW\3\2\2\2xZ\3\2\2\2x]\3\2\2\2x`\3\2\2\2xc\3\2\2\2xf\3\2\2\2"+
		"xi\3\2\2\2xl\3\2\2\2xo\3\2\2\2xr\3\2\2\2xu\3\2\2\2y\7\3\2\2\2z{\7\35\2"+
		"\2{|\7!\2\2|}\7\60\2\2}~\5\n\6\2~\177\7\61\2\2\177\u0080\7\62\2\2\u0080"+
		"\u0081\5\4\3\2\u0081\u0082\7\63\2\2\u0082\u0083\b\5\1\2\u0083\u0091\3"+
		"\2\2\2\u0084\u0085\7\35\2\2\u0085\u0086\7!\2\2\u0086\u0087\7\60\2\2\u0087"+
		"\u0088\5\n\6\2\u0088\u0089\7\61\2\2\u0089\u008a\78\2\2\u008a\u008b\5\62"+
		"\32\2\u008b\u008c\7\62\2\2\u008c\u008d\5\4\3\2\u008d\u008e\7\63\2\2\u008e"+
		"\u008f\b\5\1\2\u008f\u0091\3\2\2\2\u0090z\3\2\2\2\u0090\u0084\3\2\2\2"+
		"\u0091\t\3\2\2\2\u0092\u0093\b\6\1\2\u0093\u0094\7!\2\2\u0094\u0095\7"+
		"\64\2\2\u0095\u0096\5\62\32\2\u0096\u0097\b\6\1\2\u0097\u009a\3\2\2\2"+
		"\u0098\u009a\b\6\1\2\u0099\u0092\3\2\2\2\u0099\u0098\3\2\2\2\u009a\u00a4"+
		"\3\2\2\2\u009b\u009c\f\5\2\2\u009c\u009d\7\67\2\2\u009d\u009e\7!\2\2\u009e"+
		"\u009f\7\64\2\2\u009f\u00a0\5\62\32\2\u00a0\u00a1\b\6\1\2\u00a1\u00a3"+
		"\3\2\2\2\u00a2\u009b\3\2\2\2\u00a3\u00a6\3\2\2\2\u00a4\u00a2\3\2\2\2\u00a4"+
		"\u00a5\3\2\2\2\u00a5\13\3\2\2\2\u00a6\u00a4\3\2\2\2\u00a7\u00a8\7\36\2"+
		"\2\u00a8\u00a9\7!\2\2\u00a9\u00aa\7\62\2\2\u00aa\u00ab\5\16\b\2\u00ab"+
		"\u00ac\7\63\2\2\u00ac\u00ad\b\7\1\2\u00ad\r\3\2\2\2\u00ae\u00af\b\b\1"+
		"\2\u00af\u00b0\7\22\2\2\u00b0\u00b1\7!\2\2\u00b1\u00b2\7\64\2\2\u00b2"+
		"\u00b3\5\62\32\2\u00b3\u00b4\b\b\1\2\u00b4\u00b7\3\2\2\2\u00b5\u00b7\b"+
		"\b\1\2\u00b6\u00ae\3\2\2\2\u00b6\u00b5\3\2\2\2\u00b7\u00c2\3\2\2\2\u00b8"+
		"\u00b9\f\5\2\2\u00b9\u00ba\7\67\2\2\u00ba\u00bb\7\22\2\2\u00bb\u00bc\7"+
		"!\2\2\u00bc\u00bd\7\64\2\2\u00bd\u00be\5\62\32\2\u00be\u00bf\b\b\1\2\u00bf"+
		"\u00c1\3\2\2\2\u00c0\u00b8\3\2\2\2\u00c1\u00c4\3\2\2\2\u00c2\u00c0\3\2"+
		"\2\2\u00c2\u00c3\3\2\2\2\u00c3\17\3\2\2\2\u00c4\u00c2\3\2\2\2\u00c5\u00c6"+
		"\7!\2\2\u00c6\u00c7\79\2\2\u00c7\u00c8\7\32\2\2\u00c8\u00c9\7\60\2\2\u00c9"+
		"\u00ca\5\64\33\2\u00ca\u00cb\7\61\2\2\u00cb\u00cc\b\t\1\2\u00cc\21\3\2"+
		"\2\2\u00cd\u00ce\7!\2\2\u00ce\u00cf\79\2\2\u00cf\u00d0\7\31\2\2\u00d0"+
		"\u00d1\7\60\2\2\u00d1\u00d2\7\61\2\2\u00d2\u00d3\b\n\1\2\u00d3\23\3\2"+
		"\2\2\u00d4\u00d5\7!\2\2\u00d5\u00d6\79\2\2\u00d6\u00d7\7\30\2\2\u00d7"+
		"\u00d8\7\60\2\2\u00d8\u00d9\5\64\33\2\u00d9\u00da\7\61\2\2\u00da\u00db"+
		"\b\13\1\2\u00db\25\3\2\2\2\u00dc\u00dd\7\27\2\2\u00dd\u00de\5\64\33\2"+
		"\u00de\u00df\7\13\2\2\u00df\u00e0\7\62\2\2\u00e0\u00e1\5\4\3\2\u00e1\u00e2"+
		"\7\63\2\2\u00e2\u00e3\b\f\1\2\u00e3\27\3\2\2\2\u00e4\u00e5\7\24\2\2\u00e5"+
		"\u00e6\b\r\1\2\u00e6\31\3\2\2\2\u00e7\u00e8\7\25\2\2\u00e8\u00e9\b\16"+
		"\1\2\u00e9\33\3\2\2\2\u00ea\u00eb\7\t\2\2\u00eb\u00ec\7\60\2\2\u00ec\u00ed"+
		"\5\64\33\2\u00ed\u00ee\7\61\2\2\u00ee\u00ef\b\17\1\2\u00ef\35\3\2\2\2"+
		"\u00f0\u00f1\7\r\2\2\u00f1\u00f2\7!\2\2\u00f2\u00f3\7\21\2\2\u00f3\u00f4"+
		"\5\64\33\2\u00f4\u00f5\79\2\2\u00f5\u00f6\79\2\2\u00f6\u00f7\79\2\2\u00f7"+
		"\u00f8\5\64\33\2\u00f8\u00f9\7\62\2\2\u00f9\u00fa\5\4\3\2\u00fa\u00fb"+
		"\7\63\2\2\u00fb\u00fc\b\20\1\2\u00fc\u0107\3\2\2\2\u00fd\u00fe\7\r\2\2"+
		"\u00fe\u00ff\7!\2\2\u00ff\u0100\7\21\2\2\u0100\u0101\t\2\2\2\u0101\u0102"+
		"\7\62\2\2\u0102\u0103\5\4\3\2\u0103\u0104\7\63\2\2\u0104\u0105\b\20\1"+
		"\2\u0105\u0107\3\2\2\2\u0106\u00f0\3\2\2\2\u0106\u00fd\3\2\2\2\u0107\37"+
		"\3\2\2\2\u0108\u0109\7\16\2\2\u0109\u010a\5\64\33\2\u010a\u010b\7\62\2"+
		"\2\u010b\u010c\5\"\22\2\u010c\u010d\7\20\2\2\u010d\u010e\7\64\2\2\u010e"+
		"\u010f\5\4\3\2\u010f\u0110\7\63\2\2\u0110\u0111\b\21\1\2\u0111!\3\2\2"+
		"\2\u0112\u0113\b\22\1\2\u0113\u0114\5$\23\2\u0114\u0115\b\22\1\2\u0115"+
		"\u011c\3\2\2\2\u0116\u0117\f\4\2\2\u0117\u0118\5$\23\2\u0118\u0119\b\22"+
		"\1\2\u0119\u011b\3\2\2\2\u011a\u0116\3\2\2\2\u011b\u011e\3\2\2\2\u011c"+
		"\u011a\3\2\2\2\u011c\u011d\3\2\2\2\u011d#\3\2\2\2\u011e\u011c\3\2\2\2"+
		"\u011f\u0120\7\17\2\2\u0120\u0121\5\64\33\2\u0121\u0122\7\64\2\2\u0122"+
		"\u0123\5\4\3\2\u0123\u0124\b\23\1\2\u0124%\3\2\2\2\u0125\u0126\7\n\2\2"+
		"\u0126\u0127\5\64\33\2\u0127\u0128\7\62\2\2\u0128\u0129\5\4\3\2\u0129"+
		"\u012a\7\63\2\2\u012a\u012b\b\24\1\2\u012b\u014c\3\2\2\2\u012c\u012d\7"+
		"\n\2\2\u012d\u012e\5\64\33\2\u012e\u012f\7\62\2\2\u012f\u0130\5\4\3\2"+
		"\u0130\u0131\7\63\2\2\u0131\u0132\7\13\2\2\u0132\u0133\7\62\2\2\u0133"+
		"\u0134\5\4\3\2\u0134\u0135\7\63\2\2\u0135\u0136\b\24\1\2\u0136\u014c\3"+
		"\2\2\2\u0137\u0138\7\n\2\2\u0138\u0139\5\64\33\2\u0139\u013a\7\62\2\2"+
		"\u013a\u013b\5\4\3\2\u013b\u013c\7\63\2\2\u013c\u013d\5(\25\2\u013d\u013e"+
		"\b\24\1\2\u013e\u014c\3\2\2\2\u013f\u0140\7\n\2\2\u0140\u0141\5\64\33"+
		"\2\u0141\u0142\7\62\2\2\u0142\u0143\5\4\3\2\u0143\u0144\7\63\2\2\u0144"+
		"\u0145\5(\25\2\u0145\u0146\7\13\2\2\u0146\u0147\7\62\2\2\u0147\u0148\5"+
		"\4\3\2\u0148\u0149\7\63\2\2\u0149\u014a\b\24\1\2\u014a\u014c\3\2\2\2\u014b"+
		"\u0125\3\2\2\2\u014b\u012c\3\2\2\2\u014b\u0137\3\2\2\2\u014b\u013f\3\2"+
		"\2\2\u014c\'\3\2\2\2\u014d\u014e\b\25\1\2\u014e\u014f\5*\26\2\u014f\u0150"+
		"\b\25\1\2\u0150\u0157\3\2\2\2\u0151\u0152\f\4\2\2\u0152\u0153\5*\26\2"+
		"\u0153\u0154\b\25\1\2\u0154\u0156\3\2\2\2\u0155\u0151\3\2\2\2\u0156\u0159"+
		"\3\2\2\2\u0157\u0155\3\2\2\2\u0157\u0158\3\2\2\2\u0158)\3\2\2\2\u0159"+
		"\u0157\3\2\2\2\u015a\u015b\7\13\2\2\u015b\u015c\7\n\2\2\u015c\u015d\5"+
		"\64\33\2\u015d\u015e\7\62\2\2\u015e\u015f\5\4\3\2\u015f\u0160\7\63\2\2"+
		"\u0160\u0161\b\26\1\2\u0161+\3\2\2\2\u0162\u0163\7\f\2\2\u0163\u0164\5"+
		"\64\33\2\u0164\u0165\7\62\2\2\u0165\u0166\5\4\3\2\u0166\u0167\7\63\2\2"+
		"\u0167\u0168\b\27\1\2\u0168-\3\2\2\2\u0169\u016a\7\22\2\2\u016a\u016b"+
		"\7!\2\2\u016b\u016c\7\64\2\2\u016c\u016d\5\62\32\2\u016d\u016e\7\'\2\2"+
		"\u016e\u016f\5\64\33\2\u016f\u0170\b\30\1\2\u0170\u0186\3\2\2\2\u0171"+
		"\u0172\7\22\2\2\u0172\u0173\7!\2\2\u0173\u0174\7\'\2\2\u0174\u0175\5\64"+
		"\33\2\u0175\u0176\b\30\1\2\u0176\u0186\3\2\2\2\u0177\u0178\7\23\2\2\u0178"+
		"\u0179\7!\2\2\u0179\u017a\7\64\2\2\u017a\u017b\5\62\32\2\u017b\u017c\7"+
		"\'\2\2\u017c\u017d\5\64\33\2\u017d\u017e\b\30\1\2\u017e\u0186\3\2\2\2"+
		"\u017f\u0180\7\22\2\2\u0180\u0181\7!\2\2\u0181\u0182\7\'\2\2\u0182\u0183"+
		"\5\64\33\2\u0183\u0184\b\30\1\2\u0184\u0186\3\2\2\2\u0185\u0169\3\2\2"+
		"\2\u0185\u0171\3\2\2\2\u0185\u0177\3\2\2\2\u0185\u017f\3\2\2\2\u0186/"+
		"\3\2\2\2\u0187\u0188\7!\2\2\u0188\u0189\7\'\2\2\u0189\u018a\5\64\33\2"+
		"\u018a\u018b\b\31\1\2\u018b\61\3\2\2\2\u018c\u018d\7\3\2\2\u018d\u019d"+
		"\b\32\1\2\u018e\u018f\7\4\2\2\u018f\u019d\b\32\1\2\u0190\u0191\7\6\2\2"+
		"\u0191\u019d\b\32\1\2\u0192\u0193\7\5\2\2\u0193\u019d\b\32\1\2\u0194\u0195"+
		"\7\65\2\2\u0195\u0196\5\62\32\2\u0196\u0197\7\66\2\2\u0197\u0198\b\32"+
		"\1\2\u0198\u019d\3\2\2\2\u0199\u019a\7\65\2\2\u019a\u019b\7\66\2\2\u019b"+
		"\u019d\b\32\1\2\u019c\u018c\3\2\2\2\u019c\u018e\3\2\2\2\u019c\u0190\3"+
		"\2\2\2\u019c\u0192\3\2\2\2\u019c\u0194\3\2\2\2\u019c\u0199\3\2\2\2\u019d"+
		"\63\3\2\2\2\u019e\u019f\b\33\1\2\u019f\u01a0\7/\2\2\u01a0\u01a1\5\64\33"+
		"\31\u01a1\u01a2\b\33\1\2\u01a2\u01e4\3\2\2\2\u01a3\u01a4\7$\2\2\u01a4"+
		"\u01a5\5\64\33\21\u01a5\u01a6\b\33\1\2\u01a6\u01e4\3\2\2\2\u01a7\u01a8"+
		"\5:\36\2\u01a8\u01a9\b\33\1\2\u01a9\u01e4\3\2\2\2\u01aa\u01ab\7!\2\2\u01ab"+
		"\u01ac\7\60\2\2\u01ac\u01ad\5> \2\u01ad\u01ae\7\61\2\2\u01ae\u01af\b\33"+
		"\1\2\u01af\u01e4\3\2\2\2\u01b0\u01b1\7\60\2\2\u01b1\u01b2\5\64\33\2\u01b2"+
		"\u01b3\7\61\2\2\u01b3\u01b4\b\33\1\2\u01b4\u01e4\3\2\2\2\u01b5\u01b6\5"+
		"8\35\2\u01b6\u01b7\b\33\1\2\u01b7\u01e4\3\2\2\2\u01b8\u01b9\7\65\2\2\u01b9"+
		"\u01ba\5\66\34\2\u01ba\u01bb\7\66\2\2\u01bb\u01bc\b\33\1\2\u01bc\u01e4"+
		"\3\2\2\2\u01bd\u01be\7\3\2\2\u01be\u01bf\7\60\2\2\u01bf\u01c0\5\64\33"+
		"\2\u01c0\u01c1\7\61\2\2\u01c1\u01c2\b\33\1\2\u01c2\u01e4\3\2\2\2\u01c3"+
		"\u01c4\7\6\2\2\u01c4\u01c5\7\60\2\2\u01c5\u01c6\5\64\33\2\u01c6\u01c7"+
		"\7\61\2\2\u01c7\u01c8\b\33\1\2\u01c8\u01e4\3\2\2\2\u01c9\u01ca\7\4\2\2"+
		"\u01ca\u01cb\7\60\2\2\u01cb\u01cc\5\64\33\2\u01cc\u01cd\7\61\2\2\u01cd"+
		"\u01ce\b\33\1\2\u01ce\u01e4\3\2\2\2\u01cf\u01d0\7!\2\2\u01d0\u01d1\79"+
		"\2\2\u01d1\u01d2\7\33\2\2\u01d2\u01d3\7\60\2\2\u01d3\u01d4\7\61\2\2\u01d4"+
		"\u01e4\b\33\1\2\u01d5\u01d6\7!\2\2\u01d6\u01d7\79\2\2\u01d7\u01d8\7\34"+
		"\2\2\u01d8\u01d9\7\60\2\2\u01d9\u01da\7\61\2\2\u01da\u01e4\b\33\1\2\u01db"+
		"\u01dc\7\37\2\2\u01dc\u01e4\b\33\1\2\u01dd\u01de\7 \2\2\u01de\u01e4\b"+
		"\33\1\2\u01df\u01e0\7\7\2\2\u01e0\u01e4\b\33\1\2\u01e1\u01e2\7\b\2\2\u01e2"+
		"\u01e4\b\33\1\2\u01e3\u019e\3\2\2\2\u01e3\u01a3\3\2\2\2\u01e3\u01a7\3"+
		"\2\2\2\u01e3\u01aa\3\2\2\2\u01e3\u01b0\3\2\2\2\u01e3\u01b5\3\2\2\2\u01e3"+
		"\u01b8\3\2\2\2\u01e3\u01bd\3\2\2\2\u01e3\u01c3\3\2\2\2\u01e3\u01c9\3\2"+
		"\2\2\u01e3\u01cf\3\2\2\2\u01e3\u01d5\3\2\2\2\u01e3\u01db\3\2\2\2\u01e3"+
		"\u01dd\3\2\2\2\u01e3\u01df\3\2\2\2\u01e3\u01e1\3\2\2\2\u01e4\u020a\3\2"+
		"\2\2\u01e5\u01e6\f\30\2\2\u01e6\u01e7\t\3\2\2\u01e7\u01e8\5\64\33\31\u01e8"+
		"\u01e9\b\33\1\2\u01e9\u0209\3\2\2\2\u01ea\u01eb\f\27\2\2\u01eb\u01ec\t"+
		"\4\2\2\u01ec\u01ed\5\64\33\30\u01ed\u01ee\b\33\1\2\u01ee\u0209\3\2\2\2"+
		"\u01ef\u01f0\f\26\2\2\u01f0\u01f1\t\5\2\2\u01f1\u01f2\5\64\33\27\u01f2"+
		"\u01f3\b\33\1\2\u01f3\u0209\3\2\2\2\u01f4\u01f5\f\25\2\2\u01f5\u01f6\t"+
		"\6\2\2\u01f6\u01f7\5\64\33\26\u01f7\u01f8\b\33\1\2\u01f8\u0209\3\2\2\2"+
		"\u01f9\u01fa\f\24\2\2\u01fa\u01fb\t\7\2\2\u01fb\u01fc\5\64\33\25\u01fc"+
		"\u01fd\b\33\1\2\u01fd\u0209\3\2\2\2\u01fe\u01ff\f\23\2\2\u01ff\u0200\7"+
		"&\2\2\u0200\u0201\5\64\33\24\u0201\u0202\b\33\1\2\u0202\u0209\3\2\2\2"+
		"\u0203\u0204\f\22\2\2\u0204\u0205\7%\2\2\u0205\u0206\5\64\33\23\u0206"+
		"\u0207\b\33\1\2\u0207\u0209\3\2\2\2\u0208\u01e5\3\2\2\2\u0208\u01ea\3"+
		"\2\2\2\u0208\u01ef\3\2\2\2\u0208\u01f4\3\2\2\2\u0208\u01f9\3\2\2\2\u0208"+
		"\u01fe\3\2\2\2\u0208\u0203\3\2\2\2\u0209\u020c\3\2\2\2\u020a\u0208\3\2"+
		"\2\2\u020a\u020b\3\2\2\2\u020b\65\3\2\2\2\u020c\u020a\3\2\2\2\u020d\u020e"+
		"\b\34\1\2\u020e\u020f\5\64\33\2\u020f\u0210\b\34\1\2\u0210\u0218\3\2\2"+
		"\2\u0211\u0212\f\4\2\2\u0212\u0213\7\67\2\2\u0213\u0214\5\64\33\2\u0214"+
		"\u0215\b\34\1\2\u0215\u0217\3\2\2\2\u0216\u0211\3\2\2\2\u0217\u021a\3"+
		"\2\2\2\u0218\u0216\3\2\2\2\u0218\u0219\3\2\2\2\u0219\67\3\2\2\2\u021a"+
		"\u0218\3\2\2\2\u021b\u021c\b\35\1\2\u021c\u021d\7!\2\2\u021d\u021e\b\35"+
		"\1\2\u021e\u022b\3\2\2\2\u021f\u0220\f\5\2\2\u0220\u0221\7\65\2\2\u0221"+
		"\u0222\5\64\33\2\u0222\u0223\7\66\2\2\u0223\u0224\b\35\1\2\u0224\u022a"+
		"\3\2\2\2\u0225\u0226\f\4\2\2\u0226\u0227\79\2\2\u0227\u0228\7!\2\2\u0228"+
		"\u022a\b\35\1\2\u0229\u021f\3\2\2\2\u0229\u0225\3\2\2\2\u022a\u022d\3"+
		"\2\2\2\u022b\u0229\3\2\2\2\u022b\u022c\3\2\2\2\u022c9\3\2\2\2\u022d\u022b"+
		"\3\2\2\2\u022e\u022f\7!\2\2\u022f\u0230\7\60\2\2\u0230\u0231\5<\37\2\u0231"+
		"\u0232\7\61\2\2\u0232\u0233\b\36\1\2\u0233;\3\2\2\2\u0234\u0235\b\37\1"+
		"\2\u0235\u0236\5\64\33\2\u0236\u0237\b\37\1\2\u0237\u023a\3\2\2\2\u0238"+
		"\u023a\b\37\1\2\u0239\u0234\3\2\2\2\u0239\u0238\3\2\2\2\u023a\u0242\3"+
		"\2\2\2\u023b\u023c\f\5\2\2\u023c\u023d\7\67\2\2\u023d\u023e\5\64\33\2"+
		"\u023e\u023f\b\37\1\2\u023f\u0241\3\2\2\2\u0240\u023b\3\2\2\2\u0241\u0244"+
		"\3\2\2\2\u0242\u0240\3\2\2\2\u0242\u0243\3\2\2\2\u0243=\3\2\2\2\u0244"+
		"\u0242\3\2\2\2\u0245\u0246\b \1\2\u0246\u0247\7!\2\2\u0247\u0248\7\64"+
		"\2\2\u0248\u0249\5\64\33\2\u0249\u024a\b \1\2\u024a\u024d\3\2\2\2\u024b"+
		"\u024d\b \1\2\u024c\u0245\3\2\2\2\u024c\u024b\3\2\2\2\u024d\u0257\3\2"+
		"\2\2\u024e\u024f\f\5\2\2\u024f\u0250\7\67\2\2\u0250\u0251\7!\2\2\u0251"+
		"\u0252\7\64\2\2\u0252\u0253\5\64\33\2\u0253\u0254\b \1\2\u0254\u0256\3"+
		"\2\2\2\u0255\u024e\3\2\2\2\u0256\u0259\3\2\2\2\u0257\u0255\3\2\2\2\u0257"+
		"\u0258\3\2\2\2\u0258?\3\2\2\2\u0259\u0257\3\2\2\2\31Gx\u0090\u0099\u00a4"+
		"\u00b6\u00c2\u0106\u011c\u014b\u0157\u0185\u019c\u01e3\u0208\u020a\u0218"+
		"\u0229\u022b\u0239\u0242\u024c\u0257";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}