// Generated from c:\Users\sebas\go\bin\Server_OLC2\SwiftLexer.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SwiftLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE", "WHILE", 
			"FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "LET", "BREAK", "CONTINUE", 
			"RETURN", "GUARD", "APPEND", "REMOVELAST", "REMOVE", "COUNT", "ISEMPTY", 
			"FUNC", "STRUCT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", 
			"CORDER", "COMA", "ARROW1", "PUNTO", "MODULO", "WHITESPACE", "COMMENT", 
			"LINE_COMMENT", "ESC_SEQ"
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


	public SwiftLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SwiftLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2=\u019b\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3"+
		"\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3"+
		"\21\3\21\3\21\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\24\3"+
		"\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3"+
		"\25\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3"+
		"\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3"+
		"\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3"+
		"\35\3\35\3\36\6\36\u0124\n\36\r\36\16\36\u0125\3\36\3\36\6\36\u012a\n"+
		"\36\r\36\16\36\u012b\5\36\u012e\n\36\3\37\3\37\7\37\u0132\n\37\f\37\16"+
		"\37\u0135\13\37\3\37\3\37\3 \3 \7 \u013b\n \f \16 \u013e\13 \3!\3!\3!"+
		"\3\"\3\"\3\"\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3\'\3\'\3\'\3(\3(\3(\3)\3)"+
		"\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63"+
		"\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\3\67\38\38\39\39\3:\6:\u017a"+
		"\n:\r:\16:\u017b\3:\3:\3;\3;\3;\3;\7;\u0184\n;\f;\16;\u0187\13;\3;\3;"+
		"\3;\3;\3;\3<\3<\3<\3<\7<\u0192\n<\f<\16<\u0195\13<\3<\3<\3=\3=\3=\3\u0185"+
		"\2>\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35"+
		"\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36"+
		";\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67"+
		"m8o9q:s;u<w=y\2\3\2\t\3\2\62;\3\2$$\4\2C\\c|\6\2\62;C\\aac|\6\2\13\f\17"+
		"\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u01a1\2\3\3\2\2\2\2\5\3"+
		"\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2"+
		"\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3"+
		"\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'"+
		"\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63"+
		"\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2"+
		"?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3"+
		"\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2"+
		"\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2"+
		"e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3"+
		"\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\3{\3\2\2\2\5\177\3\2\2\2\7\u0085"+
		"\3\2\2\2\t\u008a\3\2\2\2\13\u0091\3\2\2\2\r\u0096\3\2\2\2\17\u009c\3\2"+
		"\2\2\21\u00a2\3\2\2\2\23\u00a5\3\2\2\2\25\u00aa\3\2\2\2\27\u00b0\3\2\2"+
		"\2\31\u00b4\3\2\2\2\33\u00bb\3\2\2\2\35\u00c0\3\2\2\2\37\u00c8\3\2\2\2"+
		"!\u00cb\3\2\2\2#\u00cf\3\2\2\2%\u00d3\3\2\2\2\'\u00d9\3\2\2\2)\u00e2\3"+
		"\2\2\2+\u00e9\3\2\2\2-\u00ef\3\2\2\2/\u00f6\3\2\2\2\61\u0101\3\2\2\2\63"+
		"\u0108\3\2\2\2\65\u010e\3\2\2\2\67\u0116\3\2\2\29\u011b\3\2\2\2;\u0123"+
		"\3\2\2\2=\u012f\3\2\2\2?\u0138\3\2\2\2A\u013f\3\2\2\2C\u0142\3\2\2\2E"+
		"\u0145\3\2\2\2G\u0147\3\2\2\2I\u014a\3\2\2\2K\u014d\3\2\2\2M\u014f\3\2"+
		"\2\2O\u0152\3\2\2\2Q\u0155\3\2\2\2S\u0157\3\2\2\2U\u0159\3\2\2\2W\u015b"+
		"\3\2\2\2Y\u015d\3\2\2\2[\u015f\3\2\2\2]\u0161\3\2\2\2_\u0163\3\2\2\2a"+
		"\u0165\3\2\2\2c\u0167\3\2\2\2e\u0169\3\2\2\2g\u016b\3\2\2\2i\u016d\3\2"+
		"\2\2k\u016f\3\2\2\2m\u0171\3\2\2\2o\u0174\3\2\2\2q\u0176\3\2\2\2s\u0179"+
		"\3\2\2\2u\u017f\3\2\2\2w\u018d\3\2\2\2y\u0198\3\2\2\2{|\7K\2\2|}\7p\2"+
		"\2}~\7v\2\2~\4\3\2\2\2\177\u0080\7H\2\2\u0080\u0081\7n\2\2\u0081\u0082"+
		"\7q\2\2\u0082\u0083\7c\2\2\u0083\u0084\7v\2\2\u0084\6\3\2\2\2\u0085\u0086"+
		"\7D\2\2\u0086\u0087\7q\2\2\u0087\u0088\7q\2\2\u0088\u0089\7n\2\2\u0089"+
		"\b\3\2\2\2\u008a\u008b\7U\2\2\u008b\u008c\7v\2\2\u008c\u008d\7t\2\2\u008d"+
		"\u008e\7k\2\2\u008e\u008f\7p\2\2\u008f\u0090\7i\2\2\u0090\n\3\2\2\2\u0091"+
		"\u0092\7v\2\2\u0092\u0093\7t\2\2\u0093\u0094\7w\2\2\u0094\u0095\7g\2\2"+
		"\u0095\f\3\2\2\2\u0096\u0097\7h\2\2\u0097\u0098\7c\2\2\u0098\u0099\7n"+
		"\2\2\u0099\u009a\7u\2\2\u009a\u009b\7g\2\2\u009b\16\3\2\2\2\u009c\u009d"+
		"\7r\2\2\u009d\u009e\7t\2\2\u009e\u009f\7k\2\2\u009f\u00a0\7p\2\2\u00a0"+
		"\u00a1\7v\2\2\u00a1\20\3\2\2\2\u00a2\u00a3\7k\2\2\u00a3\u00a4\7h\2\2\u00a4"+
		"\22\3\2\2\2\u00a5\u00a6\7g\2\2\u00a6\u00a7\7n\2\2\u00a7\u00a8\7u\2\2\u00a8"+
		"\u00a9\7g\2\2\u00a9\24\3\2\2\2\u00aa\u00ab\7y\2\2\u00ab\u00ac\7j\2\2\u00ac"+
		"\u00ad\7k\2\2\u00ad\u00ae\7n\2\2\u00ae\u00af\7g\2\2\u00af\26\3\2\2\2\u00b0"+
		"\u00b1\7h\2\2\u00b1\u00b2\7q\2\2\u00b2\u00b3\7t\2\2\u00b3\30\3\2\2\2\u00b4"+
		"\u00b5\7u\2\2\u00b5\u00b6\7y\2\2\u00b6\u00b7\7k\2\2\u00b7\u00b8\7v\2\2"+
		"\u00b8\u00b9\7e\2\2\u00b9\u00ba\7j\2\2\u00ba\32\3\2\2\2\u00bb\u00bc\7"+
		"e\2\2\u00bc\u00bd\7c\2\2\u00bd\u00be\7u\2\2\u00be\u00bf\7g\2\2\u00bf\34"+
		"\3\2\2\2\u00c0\u00c1\7f\2\2\u00c1\u00c2\7g\2\2\u00c2\u00c3\7h\2\2\u00c3"+
		"\u00c4\7c\2\2\u00c4\u00c5\7w\2\2\u00c5\u00c6\7n\2\2\u00c6\u00c7\7v\2\2"+
		"\u00c7\36\3\2\2\2\u00c8\u00c9\7k\2\2\u00c9\u00ca\7p\2\2\u00ca \3\2\2\2"+
		"\u00cb\u00cc\7x\2\2\u00cc\u00cd\7c\2\2\u00cd\u00ce\7t\2\2\u00ce\"\3\2"+
		"\2\2\u00cf\u00d0\7n\2\2\u00d0\u00d1\7g\2\2\u00d1\u00d2\7v\2\2\u00d2$\3"+
		"\2\2\2\u00d3\u00d4\7d\2\2\u00d4\u00d5\7t\2\2\u00d5\u00d6\7g\2\2\u00d6"+
		"\u00d7\7c\2\2\u00d7\u00d8\7m\2\2\u00d8&\3\2\2\2\u00d9\u00da\7e\2\2\u00da"+
		"\u00db\7q\2\2\u00db\u00dc\7p\2\2\u00dc\u00dd\7v\2\2\u00dd\u00de\7k\2\2"+
		"\u00de\u00df\7p\2\2\u00df\u00e0\7w\2\2\u00e0\u00e1\7g\2\2\u00e1(\3\2\2"+
		"\2\u00e2\u00e3\7t\2\2\u00e3\u00e4\7g\2\2\u00e4\u00e5\7v\2\2\u00e5\u00e6"+
		"\7w\2\2\u00e6\u00e7\7t\2\2\u00e7\u00e8\7p\2\2\u00e8*\3\2\2\2\u00e9\u00ea"+
		"\7i\2\2\u00ea\u00eb\7w\2\2\u00eb\u00ec\7c\2\2\u00ec\u00ed\7t\2\2\u00ed"+
		"\u00ee\7f\2\2\u00ee,\3\2\2\2\u00ef\u00f0\7c\2\2\u00f0\u00f1\7r\2\2\u00f1"+
		"\u00f2\7r\2\2\u00f2\u00f3\7g\2\2\u00f3\u00f4\7p\2\2\u00f4\u00f5\7f\2\2"+
		"\u00f5.\3\2\2\2\u00f6\u00f7\7t\2\2\u00f7\u00f8\7g\2\2\u00f8\u00f9\7o\2"+
		"\2\u00f9\u00fa\7q\2\2\u00fa\u00fb\7x\2\2\u00fb\u00fc\7g\2\2\u00fc\u00fd"+
		"\7N\2\2\u00fd\u00fe\7c\2\2\u00fe\u00ff\7u\2\2\u00ff\u0100\7v\2\2\u0100"+
		"\60\3\2\2\2\u0101\u0102\7t\2\2\u0102\u0103\7g\2\2\u0103\u0104\7o\2\2\u0104"+
		"\u0105\7q\2\2\u0105\u0106\7x\2\2\u0106\u0107\7g\2\2\u0107\62\3\2\2\2\u0108"+
		"\u0109\7e\2\2\u0109\u010a\7q\2\2\u010a\u010b\7w\2\2\u010b\u010c\7p\2\2"+
		"\u010c\u010d\7v\2\2\u010d\64\3\2\2\2\u010e\u010f\7K\2\2\u010f\u0110\7"+
		"u\2\2\u0110\u0111\7G\2\2\u0111\u0112\7o\2\2\u0112\u0113\7r\2\2\u0113\u0114"+
		"\7v\2\2\u0114\u0115\7{\2\2\u0115\66\3\2\2\2\u0116\u0117\7h\2\2\u0117\u0118"+
		"\7w\2\2\u0118\u0119\7p\2\2\u0119\u011a\7e\2\2\u011a8\3\2\2\2\u011b\u011c"+
		"\7u\2\2\u011c\u011d\7v\2\2\u011d\u011e\7t\2\2\u011e\u011f\7w\2\2\u011f"+
		"\u0120\7e\2\2\u0120\u0121\7v\2\2\u0121:\3\2\2\2\u0122\u0124\t\2\2\2\u0123"+
		"\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0123\3\2\2\2\u0125\u0126\3\2"+
		"\2\2\u0126\u012d\3\2\2\2\u0127\u0129\7\60\2\2\u0128\u012a\t\2\2\2\u0129"+
		"\u0128\3\2\2\2\u012a\u012b\3\2\2\2\u012b\u0129\3\2\2\2\u012b\u012c\3\2"+
		"\2\2\u012c\u012e\3\2\2\2\u012d\u0127\3\2\2\2\u012d\u012e\3\2\2\2\u012e"+
		"<\3\2\2\2\u012f\u0133\7$\2\2\u0130\u0132\n\3\2\2\u0131\u0130\3\2\2\2\u0132"+
		"\u0135\3\2\2\2\u0133\u0131\3\2\2\2\u0133\u0134\3\2\2\2\u0134\u0136\3\2"+
		"\2\2\u0135\u0133\3\2\2\2\u0136\u0137\7$\2\2\u0137>\3\2\2\2\u0138\u013c"+
		"\t\4\2\2\u0139\u013b\t\5\2\2\u013a\u0139\3\2\2\2\u013b\u013e\3\2\2\2\u013c"+
		"\u013a\3\2\2\2\u013c\u013d\3\2\2\2\u013d@\3\2\2\2\u013e\u013c\3\2\2\2"+
		"\u013f\u0140\7#\2\2\u0140\u0141\7?\2\2\u0141B\3\2\2\2\u0142\u0143\7?\2"+
		"\2\u0143\u0144\7?\2\2\u0144D\3\2\2\2\u0145\u0146\7#\2\2\u0146F\3\2\2\2"+
		"\u0147\u0148\7~\2\2\u0148\u0149\7~\2\2\u0149H\3\2\2\2\u014a\u014b\7(\2"+
		"\2\u014b\u014c\7(\2\2\u014cJ\3\2\2\2\u014d\u014e\7?\2\2\u014eL\3\2\2\2"+
		"\u014f\u0150\7@\2\2\u0150\u0151\7?\2\2\u0151N\3\2\2\2\u0152\u0153\7>\2"+
		"\2\u0153\u0154\7?\2\2\u0154P\3\2\2\2\u0155\u0156\7@\2\2\u0156R\3\2\2\2"+
		"\u0157\u0158\7>\2\2\u0158T\3\2\2\2\u0159\u015a\7,\2\2\u015aV\3\2\2\2\u015b"+
		"\u015c\7\61\2\2\u015cX\3\2\2\2\u015d\u015e\7-\2\2\u015eZ\3\2\2\2\u015f"+
		"\u0160\7/\2\2\u0160\\\3\2\2\2\u0161\u0162\7*\2\2\u0162^\3\2\2\2\u0163"+
		"\u0164\7+\2\2\u0164`\3\2\2\2\u0165\u0166\7}\2\2\u0166b\3\2\2\2\u0167\u0168"+
		"\7\177\2\2\u0168d\3\2\2\2\u0169\u016a\7<\2\2\u016af\3\2\2\2\u016b\u016c"+
		"\7]\2\2\u016ch\3\2\2\2\u016d\u016e\7_\2\2\u016ej\3\2\2\2\u016f\u0170\7"+
		".\2\2\u0170l\3\2\2\2\u0171\u0172\7/\2\2\u0172\u0173\7@\2\2\u0173n\3\2"+
		"\2\2\u0174\u0175\7\60\2\2\u0175p\3\2\2\2\u0176\u0177\7\'\2\2\u0177r\3"+
		"\2\2\2\u0178\u017a\t\6\2\2\u0179\u0178\3\2\2\2\u017a\u017b\3\2\2\2\u017b"+
		"\u0179\3\2\2\2\u017b\u017c\3\2\2\2\u017c\u017d\3\2\2\2\u017d\u017e\b:"+
		"\2\2\u017et\3\2\2\2\u017f\u0180\7\61\2\2\u0180\u0181\7,\2\2\u0181\u0185"+
		"\3\2\2\2\u0182\u0184\13\2\2\2\u0183\u0182\3\2\2\2\u0184\u0187\3\2\2\2"+
		"\u0185\u0186\3\2\2\2\u0185\u0183\3\2\2\2\u0186\u0188\3\2\2\2\u0187\u0185"+
		"\3\2\2\2\u0188\u0189\7,\2\2\u0189\u018a\7\61\2\2\u018a\u018b\3\2\2\2\u018b"+
		"\u018c\b;\2\2\u018cv\3\2\2\2\u018d\u018e\7\61\2\2\u018e\u018f\7\61\2\2"+
		"\u018f\u0193\3\2\2\2\u0190\u0192\n\7\2\2\u0191\u0190\3\2\2\2\u0192\u0195"+
		"\3\2\2\2\u0193\u0191\3\2\2\2\u0193\u0194\3\2\2\2\u0194\u0196\3\2\2\2\u0195"+
		"\u0193\3\2\2\2\u0196\u0197\b<\2\2\u0197x\3\2\2\2\u0198\u0199\7^\2\2\u0199"+
		"\u019a\t\b\2\2\u019az\3\2\2\2\13\2\u0125\u012b\u012d\u0133\u013c\u017b"+
		"\u0185\u0193\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}