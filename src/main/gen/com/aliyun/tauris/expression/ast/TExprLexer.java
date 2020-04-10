// Generated from TExprLexer.g4 by ANTLR 4.6

    package com.aliyun.tauris.expression.ast;

import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TExprLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.6", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Integer=1, Float=2, Boolean=3, String=4, AND=5, OR=6, NOT=7, IS=8, IN=9, 
		TRUE=10, FALSE=11, GT=12, GE=13, LT=14, LE=15, EQ=16, NE=17, MATCH=18, 
		LPAREN=19, RPAREN=20, INT=21, FLOAT=22, IDENTIFIER=23, VARIABLE=24, PLUS=25, 
		MINUS=26, MUL=27, DIV=28, MOD=29, POINT=30, E=31, LSHIFT=32, RSHIFT=33, 
		RSHIFT3=34, BAND=35, BEOR=36, BIOR=37, NL=38, DIGIT=39, COMMA=40, LBRACKET=41, 
		RBRACKET=42, WS=43, SLASH=44, REGEX=45;
	public static final int REG = 1;
	public static String[] modeNames = {
		"DEFAULT_MODE", "REG"
	};

	public static final String[] ruleNames = {
		"Integer", "Float", "Boolean", "String", "STRING", "STRING_ESCAPE_SEQ", 
		"ESC", "NAME", "AND", "OR", "NOT", "IS", "IN", "TRUE", "FALSE", "GT", 
		"GE", "LT", "LE", "EQ", "NE", "MATCH", "LPAREN", "RPAREN", "INT", "FLOAT", 
		"IDENTIFIER", "VARIABLE", "PLUS", "MINUS", "MUL", "DIV", "MOD", "POINT", 
		"E", "LSHIFT", "RSHIFT", "RSHIFT3", "BAND", "BEOR", "BIOR", "NL", "DIGIT", 
		"COMMA", "LBRACKET", "RBRACKET", "WS", "SLASH", "REGEX"
	};

	private static final String[] _LITERAL_NAMES = {
		null, null, null, null, null, "'&&'", "'||'", "'not'", "'is'", "'in'", 
		"'true'", "'false'", "'>'", "'>='", "'<'", "'<='", "'=='", "'!='", "'=~'", 
		"'('", "')'", null, null, null, null, "'+'", "'-'", "'*'", null, "'%'", 
		"'.'", null, "'<<'", "'>>'", "'>>>'", "'&'", "'^'", "'|'", "'\n'", null, 
		"','", "'['", "']'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "Integer", "Float", "Boolean", "String", "AND", "OR", "NOT", "IS", 
		"IN", "TRUE", "FALSE", "GT", "GE", "LT", "LE", "EQ", "NE", "MATCH", "LPAREN", 
		"RPAREN", "INT", "FLOAT", "IDENTIFIER", "VARIABLE", "PLUS", "MINUS", "MUL", 
		"DIV", "MOD", "POINT", "E", "LSHIFT", "RSHIFT", "RSHIFT3", "BAND", "BEOR", 
		"BIOR", "NL", "DIGIT", "COMMA", "LBRACKET", "RBRACKET", "WS", "SLASH", 
		"REGEX"
	};
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


	public TExprLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "TExprLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u0430\ud6d1\u8206\uad2d\u4417\uaef1\u8d80\uaadd\2/\u013f\b\1\b\1\4"+
		"\2\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n"+
		"\4\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\3\2\5\2h\n\2\3"+
		"\2\3\2\3\3\5\3m\n\3\3\3\3\3\3\3\3\3\3\4\3\4\5\4u\n\4\3\5\3\5\3\6\3\6\3"+
		"\6\7\6|\n\6\f\6\16\6\177\13\6\3\6\3\6\3\6\3\6\7\6\u0085\n\6\f\6\16\6\u0088"+
		"\13\6\3\6\5\6\u008b\n\6\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\6\t\u0095\n\t"+
		"\r\t\16\t\u0096\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r"+
		"\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\26"+
		"\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\31\3\31\3\32\5\32\u00ce"+
		"\n\32\3\32\6\32\u00d1\n\32\r\32\16\32\u00d2\3\33\5\33\u00d6\n\33\3\33"+
		"\6\33\u00d9\n\33\r\33\16\33\u00da\3\33\3\33\6\33\u00df\n\33\r\33\16\33"+
		"\u00e0\5\33\u00e3\n\33\3\34\3\34\7\34\u00e7\n\34\f\34\16\34\u00ea\13\34"+
		"\3\35\3\35\3\35\7\35\u00ef\n\35\f\35\16\35\u00f2\13\35\3\35\3\35\3\35"+
		"\7\35\u00f7\n\35\f\35\16\35\u00fa\13\35\5\35\u00fc\n\35\3\36\3\36\3\37"+
		"\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3%\3&\3&\3&\3\'\3\'\3\'\3"+
		"\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\6\60\u0127\n\60"+
		"\r\60\16\60\u0128\3\60\3\60\3\61\3\61\3\62\7\62\u0130\n\62\f\62\16\62"+
		"\u0133\13\62\3\62\3\62\3\62\6\62\u0138\n\62\r\62\16\62\u0139\3\62\3\62"+
		"\3\62\3\62\2\2\63\4\3\6\4\b\5\n\6\f\2\16\2\20\2\22\2\24\7\26\b\30\t\32"+
		"\n\34\13\36\f \r\"\16$\17&\20(\21*\22,\23.\24\60\25\62\26\64\27\66\30"+
		"8\31:\32<\33>\34@\35B\36D\37F H!J\"L#N$P%R&T\'V(X)Z*\\+^,`-b.d/\4\2\3"+
		"\17\6\2\f\f\17\17))^^\6\2\f\f\17\17$$^^\n\2$$\61\61^^ddhhppttvv\4\2C\\"+
		"c|\b\2\60\60\62;C\\^^aac|\3\2\62;\5\2C\\aac|\6\2\62;C\\aac|\7\2\60\60"+
		"\62;C\\aac|\4\2GGgg\5\2\13\f\16\17\"\"\3\2\"\"\7\2\f\f\17\17))\61\61^"+
		"^\u0150\2\4\3\2\2\2\2\6\3\2\2\2\2\b\3\2\2\2\2\n\3\2\2\2\2\24\3\2\2\2\2"+
		"\26\3\2\2\2\2\30\3\2\2\2\2\32\3\2\2\2\2\34\3\2\2\2\2\36\3\2\2\2\2 \3\2"+
		"\2\2\2\"\3\2\2\2\2$\3\2\2\2\2&\3\2\2\2\2(\3\2\2\2\2*\3\2\2\2\2,\3\2\2"+
		"\2\2.\3\2\2\2\2\60\3\2\2\2\2\62\3\2\2\2\2\64\3\2\2\2\2\66\3\2\2\2\28\3"+
		"\2\2\2\2:\3\2\2\2\2<\3\2\2\2\2>\3\2\2\2\2@\3\2\2\2\2B\3\2\2\2\2D\3\2\2"+
		"\2\2F\3\2\2\2\2H\3\2\2\2\2J\3\2\2\2\2L\3\2\2\2\2N\3\2\2\2\2P\3\2\2\2\2"+
		"R\3\2\2\2\2T\3\2\2\2\2V\3\2\2\2\2X\3\2\2\2\2Z\3\2\2\2\2\\\3\2\2\2\2^\3"+
		"\2\2\2\2`\3\2\2\2\3b\3\2\2\2\3d\3\2\2\2\4g\3\2\2\2\6l\3\2\2\2\bt\3\2\2"+
		"\2\nv\3\2\2\2\f\u008a\3\2\2\2\16\u008c\3\2\2\2\20\u008f\3\2\2\2\22\u0092"+
		"\3\2\2\2\24\u0098\3\2\2\2\26\u009b\3\2\2\2\30\u009e\3\2\2\2\32\u00a2\3"+
		"\2\2\2\34\u00a5\3\2\2\2\36\u00a8\3\2\2\2 \u00ad\3\2\2\2\"\u00b3\3\2\2"+
		"\2$\u00b5\3\2\2\2&\u00b8\3\2\2\2(\u00ba\3\2\2\2*\u00bd\3\2\2\2,\u00c0"+
		"\3\2\2\2.\u00c3\3\2\2\2\60\u00c8\3\2\2\2\62\u00ca\3\2\2\2\64\u00cd\3\2"+
		"\2\2\66\u00d5\3\2\2\28\u00e4\3\2\2\2:\u00fb\3\2\2\2<\u00fd\3\2\2\2>\u00ff"+
		"\3\2\2\2@\u0101\3\2\2\2B\u0103\3\2\2\2D\u0105\3\2\2\2F\u0107\3\2\2\2H"+
		"\u0109\3\2\2\2J\u010b\3\2\2\2L\u010e\3\2\2\2N\u0111\3\2\2\2P\u0115\3\2"+
		"\2\2R\u0117\3\2\2\2T\u0119\3\2\2\2V\u011b\3\2\2\2X\u011d\3\2\2\2Z\u011f"+
		"\3\2\2\2\\\u0121\3\2\2\2^\u0123\3\2\2\2`\u0126\3\2\2\2b\u012c\3\2\2\2"+
		"d\u0131\3\2\2\2fh\7/\2\2gf\3\2\2\2gh\3\2\2\2hi\3\2\2\2ij\5\64\32\2j\5"+
		"\3\2\2\2km\7/\2\2lk\3\2\2\2lm\3\2\2\2mn\3\2\2\2no\5\64\32\2op\7\60\2\2"+
		"pq\5\64\32\2q\7\3\2\2\2ru\5\36\17\2su\5 \20\2tr\3\2\2\2ts\3\2\2\2u\t\3"+
		"\2\2\2vw\5\f\6\2w\13\3\2\2\2x}\7)\2\2y|\5\16\7\2z|\n\2\2\2{y\3\2\2\2{"+
		"z\3\2\2\2|\177\3\2\2\2}{\3\2\2\2}~\3\2\2\2~\u0080\3\2\2\2\177}\3\2\2\2"+
		"\u0080\u008b\7)\2\2\u0081\u0086\7$\2\2\u0082\u0085\5\16\7\2\u0083\u0085"+
		"\n\3\2\2\u0084\u0082\3\2\2\2\u0084\u0083\3\2\2\2\u0085\u0088\3\2\2\2\u0086"+
		"\u0084\3\2\2\2\u0086\u0087\3\2\2\2\u0087\u0089\3\2\2\2\u0088\u0086\3\2"+
		"\2\2\u0089\u008b\7$\2\2\u008ax\3\2\2\2\u008a\u0081\3\2\2\2\u008b\r\3\2"+
		"\2\2\u008c\u008d\7^\2\2\u008d\u008e\13\2\2\2\u008e\17\3\2\2\2\u008f\u0090"+
		"\7^\2\2\u0090\u0091\t\4\2\2\u0091\21\3\2\2\2\u0092\u0094\t\5\2\2\u0093"+
		"\u0095\t\6\2\2\u0094\u0093\3\2\2\2\u0095\u0096\3\2\2\2\u0096\u0094\3\2"+
		"\2\2\u0096\u0097\3\2\2\2\u0097\23\3\2\2\2\u0098\u0099\7(\2\2\u0099\u009a"+
		"\7(\2\2\u009a\25\3\2\2\2\u009b\u009c\7~\2\2\u009c\u009d\7~\2\2\u009d\27"+
		"\3\2\2\2\u009e\u009f\7p\2\2\u009f\u00a0\7q\2\2\u00a0\u00a1\7v\2\2\u00a1"+
		"\31\3\2\2\2\u00a2\u00a3\7k\2\2\u00a3\u00a4\7u\2\2\u00a4\33\3\2\2\2\u00a5"+
		"\u00a6\7k\2\2\u00a6\u00a7\7p\2\2\u00a7\35\3\2\2\2\u00a8\u00a9\7v\2\2\u00a9"+
		"\u00aa\7t\2\2\u00aa\u00ab\7w\2\2\u00ab\u00ac\7g\2\2\u00ac\37\3\2\2\2\u00ad"+
		"\u00ae\7h\2\2\u00ae\u00af\7c\2\2\u00af\u00b0\7n\2\2\u00b0\u00b1\7u\2\2"+
		"\u00b1\u00b2\7g\2\2\u00b2!\3\2\2\2\u00b3\u00b4\7@\2\2\u00b4#\3\2\2\2\u00b5"+
		"\u00b6\7@\2\2\u00b6\u00b7\7?\2\2\u00b7%\3\2\2\2\u00b8\u00b9\7>\2\2\u00b9"+
		"\'\3\2\2\2\u00ba\u00bb\7>\2\2\u00bb\u00bc\7?\2\2\u00bc)\3\2\2\2\u00bd"+
		"\u00be\7?\2\2\u00be\u00bf\7?\2\2\u00bf+\3\2\2\2\u00c0\u00c1\7#\2\2\u00c1"+
		"\u00c2\7?\2\2\u00c2-\3\2\2\2\u00c3\u00c4\7?\2\2\u00c4\u00c5\7\u0080\2"+
		"\2\u00c5\u00c6\3\2\2\2\u00c6\u00c7\b\27\2\2\u00c7/\3\2\2\2\u00c8\u00c9"+
		"\7*\2\2\u00c9\61\3\2\2\2\u00ca\u00cb\7+\2\2\u00cb\63\3\2\2\2\u00cc\u00ce"+
		"\7/\2\2\u00cd\u00cc\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce\u00d0\3\2\2\2\u00cf"+
		"\u00d1\t\7\2\2\u00d0\u00cf\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\u00d0\3\2"+
		"\2\2\u00d2\u00d3\3\2\2\2\u00d3\65\3\2\2\2\u00d4\u00d6\7/\2\2\u00d5\u00d4"+
		"\3\2\2\2\u00d5\u00d6\3\2\2\2\u00d6\u00d8\3\2\2\2\u00d7\u00d9\t\7\2\2\u00d8"+
		"\u00d7\3\2\2\2\u00d9\u00da\3\2\2\2\u00da\u00d8\3\2\2\2\u00da\u00db\3\2"+
		"\2\2\u00db\u00e2\3\2\2\2\u00dc\u00de\7\60\2\2\u00dd\u00df\t\7\2\2\u00de"+
		"\u00dd\3\2\2\2\u00df\u00e0\3\2\2\2\u00e0\u00de\3\2\2\2\u00e0\u00e1\3\2"+
		"\2\2\u00e1\u00e3\3\2\2\2\u00e2\u00dc\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3"+
		"\67\3\2\2\2\u00e4\u00e8\t\b\2\2\u00e5\u00e7\t\t\2\2\u00e6\u00e5\3\2\2"+
		"\2\u00e7\u00ea\3\2\2\2\u00e8\u00e6\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e99"+
		"\3\2\2\2\u00ea\u00e8\3\2\2\2\u00eb\u00ec\7&\2\2\u00ec\u00f0\t\b\2\2\u00ed"+
		"\u00ef\t\n\2\2\u00ee\u00ed\3\2\2\2\u00ef\u00f2\3\2\2\2\u00f0\u00ee\3\2"+
		"\2\2\u00f0\u00f1\3\2\2\2\u00f1\u00fc\3\2\2\2\u00f2\u00f0\3\2\2\2\u00f3"+
		"\u00f4\7B\2\2\u00f4\u00f8\t\b\2\2\u00f5\u00f7\t\n\2\2\u00f6\u00f5\3\2"+
		"\2\2\u00f7\u00fa\3\2\2\2\u00f8\u00f6\3\2\2\2\u00f8\u00f9\3\2\2\2\u00f9"+
		"\u00fc\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fb\u00eb\3\2\2\2\u00fb\u00f3\3\2"+
		"\2\2\u00fc;\3\2\2\2\u00fd\u00fe\7-\2\2\u00fe=\3\2\2\2\u00ff\u0100\7/\2"+
		"\2\u0100?\3\2\2\2\u0101\u0102\7,\2\2\u0102A\3\2\2\2\u0103\u0104\7\61\2"+
		"\2\u0104C\3\2\2\2\u0105\u0106\7\'\2\2\u0106E\3\2\2\2\u0107\u0108\7\60"+
		"\2\2\u0108G\3\2\2\2\u0109\u010a\t\13\2\2\u010aI\3\2\2\2\u010b\u010c\7"+
		">\2\2\u010c\u010d\7>\2\2\u010dK\3\2\2\2\u010e\u010f\7@\2\2\u010f\u0110"+
		"\7@\2\2\u0110M\3\2\2\2\u0111\u0112\7@\2\2\u0112\u0113\7@\2\2\u0113\u0114"+
		"\7@\2\2\u0114O\3\2\2\2\u0115\u0116\7(\2\2\u0116Q\3\2\2\2\u0117\u0118\7"+
		"`\2\2\u0118S\3\2\2\2\u0119\u011a\7~\2\2\u011aU\3\2\2\2\u011b\u011c\7\f"+
		"\2\2\u011cW\3\2\2\2\u011d\u011e\4\62;\2\u011eY\3\2\2\2\u011f\u0120\7."+
		"\2\2\u0120[\3\2\2\2\u0121\u0122\7]\2\2\u0122]\3\2\2\2\u0123\u0124\7_\2"+
		"\2\u0124_\3\2\2\2\u0125\u0127\t\f\2\2\u0126\u0125\3\2\2\2\u0127\u0128"+
		"\3\2\2\2\u0128\u0126\3\2\2\2\u0128\u0129\3\2\2\2\u0129\u012a\3\2\2\2\u012a"+
		"\u012b\b\60\3\2\u012ba\3\2\2\2\u012c\u012d\7\61\2\2\u012dc\3\2\2\2\u012e"+
		"\u0130\t\r\2\2\u012f\u012e\3\2\2\2\u0130\u0133\3\2\2\2\u0131\u012f\3\2"+
		"\2\2\u0131\u0132\3\2\2\2\u0132\u0134\3\2\2\2\u0133\u0131\3\2\2\2\u0134"+
		"\u0137\5b\61\2\u0135\u0138\5\16\7\2\u0136\u0138\n\16\2\2\u0137\u0135\3"+
		"\2\2\2\u0137\u0136\3\2\2\2\u0138\u0139\3\2\2\2\u0139\u0137\3\2\2\2\u0139"+
		"\u013a\3\2\2\2\u013a\u013b\3\2\2\2\u013b\u013c\5b\61\2\u013c\u013d\3\2"+
		"\2\2\u013d\u013e\b\62\4\2\u013ee\3\2\2\2\33\2\3glt{}\u0084\u0086\u008a"+
		"\u0096\u00cd\u00d2\u00d5\u00da\u00e0\u00e2\u00e8\u00f0\u00f8\u00fb\u0128"+
		"\u0131\u0137\u0139\5\7\3\2\b\2\2\6\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}