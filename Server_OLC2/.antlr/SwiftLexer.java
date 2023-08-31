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
		FOR=11, SWITCH=12, CASE=13, DEFAULT=14, IN=15, VAR=16, BREAK=17, CONTINUE=18, 
		RETURN=19, GUARD=20, NUMBER=21, STRING=22, ID=23, DIF=24, IG_IG=25, NOT=26, 
		OR=27, AND=28, IG=29, MAY_IG=30, MEN_IG=31, MAYOR=32, MENOR=33, MUL=34, 
		DIV=35, ADD=36, SUB=37, PARIZQ=38, PARDER=39, LLAVEIZQ=40, LLAVEDER=41, 
		D_PTS=42, CORIZQ=43, CORDER=44, COMA=45, PUNTO=46, WHITESPACE=47, COMMENT=48, 
		LINE_COMMENT=49;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE", "WHILE", 
			"FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "BREAK", "CONTINUE", 
			"RETURN", "GUARD", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", 
			"CORDER", "COMA", "PUNTO", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'true'", "'false'", 
			"'print'", "'if'", "'else'", "'while'", "'for'", "'switch'", "'case'", 
			"'default'", "'in'", "'var'", "'break'", "'continue'", "'return'", "'guard'", 
			null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", 
			"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", 
			"'}'", "':'", "'['", "']'", "','", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE", 
			"WHILE", "FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "BREAK", "CONTINUE", 
			"RETURN", "GUARD", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", 
			"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", 
			"CORDER", "COMA", "PUNTO", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\63\u014b\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\3\2"+
		"\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16"+
		"\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\26\6\26\u00d9\n\26\r\26\16\26\u00da\3\26\3\26\6\26\u00df"+
		"\n\26\r\26\16\26\u00e0\5\26\u00e3\n\26\3\27\3\27\7\27\u00e7\n\27\f\27"+
		"\16\27\u00ea\13\27\3\27\3\27\3\30\3\30\7\30\u00f0\n\30\f\30\16\30\u00f3"+
		"\13\30\3\31\3\31\3\31\3\32\3\32\3\32\3\33\3\33\3\34\3\34\3\34\3\35\3\35"+
		"\3\35\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3"+
		"%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\6"+
		"\60\u012a\n\60\r\60\16\60\u012b\3\60\3\60\3\61\3\61\3\61\3\61\7\61\u0134"+
		"\n\61\f\61\16\61\u0137\13\61\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62\3"+
		"\62\7\62\u0142\n\62\f\62\16\62\u0145\13\62\3\62\3\62\3\63\3\63\3\63\3"+
		"\u0135\2\64\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33"+
		"\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67"+
		"\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\2\3\2\t"+
		"\3\2\62;\3\2$$\4\2C\\c|\6\2\62;C\\aac|\6\2\13\f\17\17\"\"^^\4\2\f\f\17"+
		"\17\t\2\"#%%--/\60<<BB]_\2\u0151\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2"+
		"\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2"+
		"\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2"+
		"\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2"+
		"\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2"+
		"\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2"+
		"\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O"+
		"\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2"+
		"\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\3g\3\2\2\2\5k\3\2\2\2"+
		"\7q\3\2\2\2\tv\3\2\2\2\13}\3\2\2\2\r\u0082\3\2\2\2\17\u0088\3\2\2\2\21"+
		"\u008e\3\2\2\2\23\u0091\3\2\2\2\25\u0096\3\2\2\2\27\u009c\3\2\2\2\31\u00a0"+
		"\3\2\2\2\33\u00a7\3\2\2\2\35\u00ac\3\2\2\2\37\u00b4\3\2\2\2!\u00b7\3\2"+
		"\2\2#\u00bb\3\2\2\2%\u00c1\3\2\2\2\'\u00ca\3\2\2\2)\u00d1\3\2\2\2+\u00d8"+
		"\3\2\2\2-\u00e4\3\2\2\2/\u00ed\3\2\2\2\61\u00f4\3\2\2\2\63\u00f7\3\2\2"+
		"\2\65\u00fa\3\2\2\2\67\u00fc\3\2\2\29\u00ff\3\2\2\2;\u0102\3\2\2\2=\u0104"+
		"\3\2\2\2?\u0107\3\2\2\2A\u010a\3\2\2\2C\u010c\3\2\2\2E\u010e\3\2\2\2G"+
		"\u0110\3\2\2\2I\u0112\3\2\2\2K\u0114\3\2\2\2M\u0116\3\2\2\2O\u0118\3\2"+
		"\2\2Q\u011a\3\2\2\2S\u011c\3\2\2\2U\u011e\3\2\2\2W\u0120\3\2\2\2Y\u0122"+
		"\3\2\2\2[\u0124\3\2\2\2]\u0126\3\2\2\2_\u0129\3\2\2\2a\u012f\3\2\2\2c"+
		"\u013d\3\2\2\2e\u0148\3\2\2\2gh\7K\2\2hi\7p\2\2ij\7v\2\2j\4\3\2\2\2kl"+
		"\7H\2\2lm\7n\2\2mn\7q\2\2no\7c\2\2op\7v\2\2p\6\3\2\2\2qr\7D\2\2rs\7q\2"+
		"\2st\7q\2\2tu\7n\2\2u\b\3\2\2\2vw\7U\2\2wx\7v\2\2xy\7t\2\2yz\7k\2\2z{"+
		"\7p\2\2{|\7i\2\2|\n\3\2\2\2}~\7v\2\2~\177\7t\2\2\177\u0080\7w\2\2\u0080"+
		"\u0081\7g\2\2\u0081\f\3\2\2\2\u0082\u0083\7h\2\2\u0083\u0084\7c\2\2\u0084"+
		"\u0085\7n\2\2\u0085\u0086\7u\2\2\u0086\u0087\7g\2\2\u0087\16\3\2\2\2\u0088"+
		"\u0089\7r\2\2\u0089\u008a\7t\2\2\u008a\u008b\7k\2\2\u008b\u008c\7p\2\2"+
		"\u008c\u008d\7v\2\2\u008d\20\3\2\2\2\u008e\u008f\7k\2\2\u008f\u0090\7"+
		"h\2\2\u0090\22\3\2\2\2\u0091\u0092\7g\2\2\u0092\u0093\7n\2\2\u0093\u0094"+
		"\7u\2\2\u0094\u0095\7g\2\2\u0095\24\3\2\2\2\u0096\u0097\7y\2\2\u0097\u0098"+
		"\7j\2\2\u0098\u0099\7k\2\2\u0099\u009a\7n\2\2\u009a\u009b\7g\2\2\u009b"+
		"\26\3\2\2\2\u009c\u009d\7h\2\2\u009d\u009e\7q\2\2\u009e\u009f\7t\2\2\u009f"+
		"\30\3\2\2\2\u00a0\u00a1\7u\2\2\u00a1\u00a2\7y\2\2\u00a2\u00a3\7k\2\2\u00a3"+
		"\u00a4\7v\2\2\u00a4\u00a5\7e\2\2\u00a5\u00a6\7j\2\2\u00a6\32\3\2\2\2\u00a7"+
		"\u00a8\7e\2\2\u00a8\u00a9\7c\2\2\u00a9\u00aa\7u\2\2\u00aa\u00ab\7g\2\2"+
		"\u00ab\34\3\2\2\2\u00ac\u00ad\7f\2\2\u00ad\u00ae\7g\2\2\u00ae\u00af\7"+
		"h\2\2\u00af\u00b0\7c\2\2\u00b0\u00b1\7w\2\2\u00b1\u00b2\7n\2\2\u00b2\u00b3"+
		"\7v\2\2\u00b3\36\3\2\2\2\u00b4\u00b5\7k\2\2\u00b5\u00b6\7p\2\2\u00b6 "+
		"\3\2\2\2\u00b7\u00b8\7x\2\2\u00b8\u00b9\7c\2\2\u00b9\u00ba\7t\2\2\u00ba"+
		"\"\3\2\2\2\u00bb\u00bc\7d\2\2\u00bc\u00bd\7t\2\2\u00bd\u00be\7g\2\2\u00be"+
		"\u00bf\7c\2\2\u00bf\u00c0\7m\2\2\u00c0$\3\2\2\2\u00c1\u00c2\7e\2\2\u00c2"+
		"\u00c3\7q\2\2\u00c3\u00c4\7p\2\2\u00c4\u00c5\7v\2\2\u00c5\u00c6\7k\2\2"+
		"\u00c6\u00c7\7p\2\2\u00c7\u00c8\7w\2\2\u00c8\u00c9\7g\2\2\u00c9&\3\2\2"+
		"\2\u00ca\u00cb\7t\2\2\u00cb\u00cc\7g\2\2\u00cc\u00cd\7v\2\2\u00cd\u00ce"+
		"\7w\2\2\u00ce\u00cf\7t\2\2\u00cf\u00d0\7p\2\2\u00d0(\3\2\2\2\u00d1\u00d2"+
		"\7i\2\2\u00d2\u00d3\7w\2\2\u00d3\u00d4\7c\2\2\u00d4\u00d5\7t\2\2\u00d5"+
		"\u00d6\7f\2\2\u00d6*\3\2\2\2\u00d7\u00d9\t\2\2\2\u00d8\u00d7\3\2\2\2\u00d9"+
		"\u00da\3\2\2\2\u00da\u00d8\3\2\2\2\u00da\u00db\3\2\2\2\u00db\u00e2\3\2"+
		"\2\2\u00dc\u00de\7\60\2\2\u00dd\u00df\t\2\2\2\u00de\u00dd\3\2\2\2\u00df"+
		"\u00e0\3\2\2\2\u00e0\u00de\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00e3\3\2"+
		"\2\2\u00e2\u00dc\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3,\3\2\2\2\u00e4\u00e8"+
		"\7$\2\2\u00e5\u00e7\n\3\2\2\u00e6\u00e5\3\2\2\2\u00e7\u00ea\3\2\2\2\u00e8"+
		"\u00e6\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9\u00eb\3\2\2\2\u00ea\u00e8\3\2"+
		"\2\2\u00eb\u00ec\7$\2\2\u00ec.\3\2\2\2\u00ed\u00f1\t\4\2\2\u00ee\u00f0"+
		"\t\5\2\2\u00ef\u00ee\3\2\2\2\u00f0\u00f3\3\2\2\2\u00f1\u00ef\3\2\2\2\u00f1"+
		"\u00f2\3\2\2\2\u00f2\60\3\2\2\2\u00f3\u00f1\3\2\2\2\u00f4\u00f5\7#\2\2"+
		"\u00f5\u00f6\7?\2\2\u00f6\62\3\2\2\2\u00f7\u00f8\7?\2\2\u00f8\u00f9\7"+
		"?\2\2\u00f9\64\3\2\2\2\u00fa\u00fb\7#\2\2\u00fb\66\3\2\2\2\u00fc\u00fd"+
		"\7~\2\2\u00fd\u00fe\7~\2\2\u00fe8\3\2\2\2\u00ff\u0100\7(\2\2\u0100\u0101"+
		"\7(\2\2\u0101:\3\2\2\2\u0102\u0103\7?\2\2\u0103<\3\2\2\2\u0104\u0105\7"+
		"@\2\2\u0105\u0106\7?\2\2\u0106>\3\2\2\2\u0107\u0108\7>\2\2\u0108\u0109"+
		"\7?\2\2\u0109@\3\2\2\2\u010a\u010b\7@\2\2\u010bB\3\2\2\2\u010c\u010d\7"+
		">\2\2\u010dD\3\2\2\2\u010e\u010f\7,\2\2\u010fF\3\2\2\2\u0110\u0111\7\61"+
		"\2\2\u0111H\3\2\2\2\u0112\u0113\7-\2\2\u0113J\3\2\2\2\u0114\u0115\7/\2"+
		"\2\u0115L\3\2\2\2\u0116\u0117\7*\2\2\u0117N\3\2\2\2\u0118\u0119\7+\2\2"+
		"\u0119P\3\2\2\2\u011a\u011b\7}\2\2\u011bR\3\2\2\2\u011c\u011d\7\177\2"+
		"\2\u011dT\3\2\2\2\u011e\u011f\7<\2\2\u011fV\3\2\2\2\u0120\u0121\7]\2\2"+
		"\u0121X\3\2\2\2\u0122\u0123\7_\2\2\u0123Z\3\2\2\2\u0124\u0125\7.\2\2\u0125"+
		"\\\3\2\2\2\u0126\u0127\7\60\2\2\u0127^\3\2\2\2\u0128\u012a\t\6\2\2\u0129"+
		"\u0128\3\2\2\2\u012a\u012b\3\2\2\2\u012b\u0129\3\2\2\2\u012b\u012c\3\2"+
		"\2\2\u012c\u012d\3\2\2\2\u012d\u012e\b\60\2\2\u012e`\3\2\2\2\u012f\u0130"+
		"\7\61\2\2\u0130\u0131\7,\2\2\u0131\u0135\3\2\2\2\u0132\u0134\13\2\2\2"+
		"\u0133\u0132\3\2\2\2\u0134\u0137\3\2\2\2\u0135\u0136\3\2\2\2\u0135\u0133"+
		"\3\2\2\2\u0136\u0138\3\2\2\2\u0137\u0135\3\2\2\2\u0138\u0139\7,\2\2\u0139"+
		"\u013a\7\61\2\2\u013a\u013b\3\2\2\2\u013b\u013c\b\61\2\2\u013cb\3\2\2"+
		"\2\u013d\u013e\7\61\2\2\u013e\u013f\7\61\2\2\u013f\u0143\3\2\2\2\u0140"+
		"\u0142\n\7\2\2\u0141\u0140\3\2\2\2\u0142\u0145\3\2\2\2\u0143\u0141\3\2"+
		"\2\2\u0143\u0144\3\2\2\2\u0144\u0146\3\2\2\2\u0145\u0143\3\2\2\2\u0146"+
		"\u0147\b\62\2\2\u0147d\3\2\2\2\u0148\u0149\7^\2\2\u0149\u014a\t\b\2\2"+
		"\u014af\3\2\2\2\13\2\u00da\u00e0\u00e2\u00e8\u00f1\u012b\u0135\u0143\3"+
		"\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}