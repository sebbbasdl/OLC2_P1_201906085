// Generated from c:\Users\sebas\go\bin\Server2\SwiftLexer.g4 by ANTLR 4.9.2
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
		RETURN=19, NUMBER=20, STRING=21, ID=22, DIF=23, IG_IG=24, NOT=25, OR=26, 
		AND=27, IG=28, MAY_IG=29, MEN_IG=30, MAYOR=31, MENOR=32, MUL=33, DIV=34, 
		ADD=35, SUB=36, PARIZQ=37, PARDER=38, LLAVEIZQ=39, LLAVEDER=40, D_PTS=41, 
		CORIZQ=42, CORDER=43, COMA=44, PUNTO=45, WHITESPACE=46, COMMENT=47, LINE_COMMENT=48;
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
			"RETURN", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", 
			"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", 
			"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "D_PTS", "CORIZQ", "CORDER", 
			"COMA", "PUNTO", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'Int'", "'Float'", "'Bool'", "'String'", "'true'", "'false'", 
			"'print'", "'if'", "'else'", "'while'", "'for'", "'switch'", "'case'", 
			"'default'", "'in'", "'var'", "'break'", "'continue'", "'return'", null, 
			null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", 
			"'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", 
			"':'", "'['", "']'", "','", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "FLOAT", "BOOL", "STR", "TRU", "FAL", "PRINT", "IF", "ELSE", 
			"WHILE", "FOR", "SWITCH", "CASE", "DEFAULT", "IN", "VAR", "BREAK", "CONTINUE", 
			"RETURN", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", 
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\62\u0143\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\3\2\3\2\3\2\3"+
		"\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3"+
		"\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3"+
		"\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\25\6\25\u00d1\n\25\r\25\16\25"+
		"\u00d2\3\25\3\25\6\25\u00d7\n\25\r\25\16\25\u00d8\5\25\u00db\n\25\3\26"+
		"\3\26\7\26\u00df\n\26\f\26\16\26\u00e2\13\26\3\26\3\26\3\27\3\27\7\27"+
		"\u00e8\n\27\f\27\16\27\u00eb\13\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32"+
		"\3\32\3\33\3\33\3\33\3\34\3\34\3\34\3\35\3\35\3\36\3\36\3\36\3\37\3\37"+
		"\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)"+
		"\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\6/\u0122\n/\r/\16/\u0123\3/\3/\3\60"+
		"\3\60\3\60\3\60\7\60\u012c\n\60\f\60\16\60\u012f\13\60\3\60\3\60\3\60"+
		"\3\60\3\60\3\61\3\61\3\61\3\61\7\61\u013a\n\61\f\61\16\61\u013d\13\61"+
		"\3\61\3\61\3\62\3\62\3\62\3\u012d\2\63\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21"+
		"\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30"+
		"/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.["+
		"/]\60_\61a\62c\2\3\2\t\3\2\62;\3\2$$\4\2C\\c|\6\2\62;C\\aac|\6\2\13\f"+
		"\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u0149\2\3\3\2\2\2\2"+
		"\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2"+
		"\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2"+
		"\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2"+
		"\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2"+
		"\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2"+
		"\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2"+
		"K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3"+
		"\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\3e\3\2\2"+
		"\2\5i\3\2\2\2\7o\3\2\2\2\tt\3\2\2\2\13{\3\2\2\2\r\u0080\3\2\2\2\17\u0086"+
		"\3\2\2\2\21\u008c\3\2\2\2\23\u008f\3\2\2\2\25\u0094\3\2\2\2\27\u009a\3"+
		"\2\2\2\31\u009e\3\2\2\2\33\u00a5\3\2\2\2\35\u00aa\3\2\2\2\37\u00b2\3\2"+
		"\2\2!\u00b5\3\2\2\2#\u00b9\3\2\2\2%\u00bf\3\2\2\2\'\u00c8\3\2\2\2)\u00d0"+
		"\3\2\2\2+\u00dc\3\2\2\2-\u00e5\3\2\2\2/\u00ec\3\2\2\2\61\u00ef\3\2\2\2"+
		"\63\u00f2\3\2\2\2\65\u00f4\3\2\2\2\67\u00f7\3\2\2\29\u00fa\3\2\2\2;\u00fc"+
		"\3\2\2\2=\u00ff\3\2\2\2?\u0102\3\2\2\2A\u0104\3\2\2\2C\u0106\3\2\2\2E"+
		"\u0108\3\2\2\2G\u010a\3\2\2\2I\u010c\3\2\2\2K\u010e\3\2\2\2M\u0110\3\2"+
		"\2\2O\u0112\3\2\2\2Q\u0114\3\2\2\2S\u0116\3\2\2\2U\u0118\3\2\2\2W\u011a"+
		"\3\2\2\2Y\u011c\3\2\2\2[\u011e\3\2\2\2]\u0121\3\2\2\2_\u0127\3\2\2\2a"+
		"\u0135\3\2\2\2c\u0140\3\2\2\2ef\7K\2\2fg\7p\2\2gh\7v\2\2h\4\3\2\2\2ij"+
		"\7H\2\2jk\7n\2\2kl\7q\2\2lm\7c\2\2mn\7v\2\2n\6\3\2\2\2op\7D\2\2pq\7q\2"+
		"\2qr\7q\2\2rs\7n\2\2s\b\3\2\2\2tu\7U\2\2uv\7v\2\2vw\7t\2\2wx\7k\2\2xy"+
		"\7p\2\2yz\7i\2\2z\n\3\2\2\2{|\7v\2\2|}\7t\2\2}~\7w\2\2~\177\7g\2\2\177"+
		"\f\3\2\2\2\u0080\u0081\7h\2\2\u0081\u0082\7c\2\2\u0082\u0083\7n\2\2\u0083"+
		"\u0084\7u\2\2\u0084\u0085\7g\2\2\u0085\16\3\2\2\2\u0086\u0087\7r\2\2\u0087"+
		"\u0088\7t\2\2\u0088\u0089\7k\2\2\u0089\u008a\7p\2\2\u008a\u008b\7v\2\2"+
		"\u008b\20\3\2\2\2\u008c\u008d\7k\2\2\u008d\u008e\7h\2\2\u008e\22\3\2\2"+
		"\2\u008f\u0090\7g\2\2\u0090\u0091\7n\2\2\u0091\u0092\7u\2\2\u0092\u0093"+
		"\7g\2\2\u0093\24\3\2\2\2\u0094\u0095\7y\2\2\u0095\u0096\7j\2\2\u0096\u0097"+
		"\7k\2\2\u0097\u0098\7n\2\2\u0098\u0099\7g\2\2\u0099\26\3\2\2\2\u009a\u009b"+
		"\7h\2\2\u009b\u009c\7q\2\2\u009c\u009d\7t\2\2\u009d\30\3\2\2\2\u009e\u009f"+
		"\7u\2\2\u009f\u00a0\7y\2\2\u00a0\u00a1\7k\2\2\u00a1\u00a2\7v\2\2\u00a2"+
		"\u00a3\7e\2\2\u00a3\u00a4\7j\2\2\u00a4\32\3\2\2\2\u00a5\u00a6\7e\2\2\u00a6"+
		"\u00a7\7c\2\2\u00a7\u00a8\7u\2\2\u00a8\u00a9\7g\2\2\u00a9\34\3\2\2\2\u00aa"+
		"\u00ab\7f\2\2\u00ab\u00ac\7g\2\2\u00ac\u00ad\7h\2\2\u00ad\u00ae\7c\2\2"+
		"\u00ae\u00af\7w\2\2\u00af\u00b0\7n\2\2\u00b0\u00b1\7v\2\2\u00b1\36\3\2"+
		"\2\2\u00b2\u00b3\7k\2\2\u00b3\u00b4\7p\2\2\u00b4 \3\2\2\2\u00b5\u00b6"+
		"\7x\2\2\u00b6\u00b7\7c\2\2\u00b7\u00b8\7t\2\2\u00b8\"\3\2\2\2\u00b9\u00ba"+
		"\7d\2\2\u00ba\u00bb\7t\2\2\u00bb\u00bc\7g\2\2\u00bc\u00bd\7c\2\2\u00bd"+
		"\u00be\7m\2\2\u00be$\3\2\2\2\u00bf\u00c0\7e\2\2\u00c0\u00c1\7q\2\2\u00c1"+
		"\u00c2\7p\2\2\u00c2\u00c3\7v\2\2\u00c3\u00c4\7k\2\2\u00c4\u00c5\7p\2\2"+
		"\u00c5\u00c6\7w\2\2\u00c6\u00c7\7g\2\2\u00c7&\3\2\2\2\u00c8\u00c9\7t\2"+
		"\2\u00c9\u00ca\7g\2\2\u00ca\u00cb\7v\2\2\u00cb\u00cc\7w\2\2\u00cc\u00cd"+
		"\7t\2\2\u00cd\u00ce\7p\2\2\u00ce(\3\2\2\2\u00cf\u00d1\t\2\2\2\u00d0\u00cf"+
		"\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\u00d0\3\2\2\2\u00d2\u00d3\3\2\2\2\u00d3"+
		"\u00da\3\2\2\2\u00d4\u00d6\7\60\2\2\u00d5\u00d7\t\2\2\2\u00d6\u00d5\3"+
		"\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9"+
		"\u00db\3\2\2\2\u00da\u00d4\3\2\2\2\u00da\u00db\3\2\2\2\u00db*\3\2\2\2"+
		"\u00dc\u00e0\7$\2\2\u00dd\u00df\n\3\2\2\u00de\u00dd\3\2\2\2\u00df\u00e2"+
		"\3\2\2\2\u00e0\u00de\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00e3\3\2\2\2\u00e2"+
		"\u00e0\3\2\2\2\u00e3\u00e4\7$\2\2\u00e4,\3\2\2\2\u00e5\u00e9\t\4\2\2\u00e6"+
		"\u00e8\t\5\2\2\u00e7\u00e6\3\2\2\2\u00e8\u00eb\3\2\2\2\u00e9\u00e7\3\2"+
		"\2\2\u00e9\u00ea\3\2\2\2\u00ea.\3\2\2\2\u00eb\u00e9\3\2\2\2\u00ec\u00ed"+
		"\7#\2\2\u00ed\u00ee\7?\2\2\u00ee\60\3\2\2\2\u00ef\u00f0\7?\2\2\u00f0\u00f1"+
		"\7?\2\2\u00f1\62\3\2\2\2\u00f2\u00f3\7#\2\2\u00f3\64\3\2\2\2\u00f4\u00f5"+
		"\7~\2\2\u00f5\u00f6\7~\2\2\u00f6\66\3\2\2\2\u00f7\u00f8\7(\2\2\u00f8\u00f9"+
		"\7(\2\2\u00f98\3\2\2\2\u00fa\u00fb\7?\2\2\u00fb:\3\2\2\2\u00fc\u00fd\7"+
		"@\2\2\u00fd\u00fe\7?\2\2\u00fe<\3\2\2\2\u00ff\u0100\7>\2\2\u0100\u0101"+
		"\7?\2\2\u0101>\3\2\2\2\u0102\u0103\7@\2\2\u0103@\3\2\2\2\u0104\u0105\7"+
		">\2\2\u0105B\3\2\2\2\u0106\u0107\7,\2\2\u0107D\3\2\2\2\u0108\u0109\7\61"+
		"\2\2\u0109F\3\2\2\2\u010a\u010b\7-\2\2\u010bH\3\2\2\2\u010c\u010d\7/\2"+
		"\2\u010dJ\3\2\2\2\u010e\u010f\7*\2\2\u010fL\3\2\2\2\u0110\u0111\7+\2\2"+
		"\u0111N\3\2\2\2\u0112\u0113\7}\2\2\u0113P\3\2\2\2\u0114\u0115\7\177\2"+
		"\2\u0115R\3\2\2\2\u0116\u0117\7<\2\2\u0117T\3\2\2\2\u0118\u0119\7]\2\2"+
		"\u0119V\3\2\2\2\u011a\u011b\7_\2\2\u011bX\3\2\2\2\u011c\u011d\7.\2\2\u011d"+
		"Z\3\2\2\2\u011e\u011f\7\60\2\2\u011f\\\3\2\2\2\u0120\u0122\t\6\2\2\u0121"+
		"\u0120\3\2\2\2\u0122\u0123\3\2\2\2\u0123\u0121\3\2\2\2\u0123\u0124\3\2"+
		"\2\2\u0124\u0125\3\2\2\2\u0125\u0126\b/\2\2\u0126^\3\2\2\2\u0127\u0128"+
		"\7\61\2\2\u0128\u0129\7,\2\2\u0129\u012d\3\2\2\2\u012a\u012c\13\2\2\2"+
		"\u012b\u012a\3\2\2\2\u012c\u012f\3\2\2\2\u012d\u012e\3\2\2\2\u012d\u012b"+
		"\3\2\2\2\u012e\u0130\3\2\2\2\u012f\u012d\3\2\2\2\u0130\u0131\7,\2\2\u0131"+
		"\u0132\7\61\2\2\u0132\u0133\3\2\2\2\u0133\u0134\b\60\2\2\u0134`\3\2\2"+
		"\2\u0135\u0136\7\61\2\2\u0136\u0137\7\61\2\2\u0137\u013b\3\2\2\2\u0138"+
		"\u013a\n\7\2\2\u0139\u0138\3\2\2\2\u013a\u013d\3\2\2\2\u013b\u0139\3\2"+
		"\2\2\u013b\u013c\3\2\2\2\u013c\u013e\3\2\2\2\u013d\u013b\3\2\2\2\u013e"+
		"\u013f\b\61\2\2\u013fb\3\2\2\2\u0140\u0141\7^\2\2\u0141\u0142\t\b\2\2"+
		"\u0142d\3\2\2\2\13\2\u00d2\u00d8\u00da\u00e0\u00e9\u0123\u012d\u013b\3"+
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