
// Generated by gogll. Do not edit.
//
//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package symbols

func IsNonTerminal(symbol string) bool {
	return nonTerminals[symbol]
}

func IsTerminal(symbol string) bool {
	return !nonTerminals[symbol]
}

var nonTerminals = map[string]bool{ 
	"Alternate":true,
	"Alternates":true,
	"CharLiteral":true,
	"EscapedChar":true,
	"GoGLL":true,
	"Head":true,
	"NTChar":true,
	"NTChars":true,
	"NTID":true,
	"NonTerminal":true,
	"Package":true,
	"Rule":true,
	"Rules":true,
	"Sep":true,
	"SepChar":true,
	"SepE":true,
	"String":true,
	"StringChars":true,
	"Symbol":true,
	"Symbols":true,
	"Terminal":true,
}

var reservedWords = map[string]bool{ 
	"*":true,
	"empty":true,
	"\\":true,
	"t":true,
	"\"":true,
	"package":true,
	"space":true,
	"lowcase":true,
	"'":true,
	"n":true,
	"|":true,
	"anyof":true,
	"number":true,
	"letter":true,
	"upcase":true,
	"not":true,
	"r":true,
	":":true,
	";":true,
	"any":true,
}

