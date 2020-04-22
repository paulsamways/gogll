
// Package slot is generated by gogll. Do not edit. 
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
package slot

import(
	"bytes"
    "fmt"
    
    "github.com/goccmack/gogll/parser/symbols"
)

type Label int

const(
	Alternate0R0 Label = iota
	Alternate0R1
	Alternate1R0
	Alternate1R1
	Alternates0R0
	Alternates0R1
	Alternates1R0
	Alternates1R1
	Alternates1R2
	Alternates1R3
	GoGLL0R0
	GoGLL0R1
	GoGLL0R2
	NT0R0
	NT0R1
	Package0R0
	Package0R1
	Package0R2
	Rule0R0
	Rule0R1
	Rule0R2
	Rule0R3
	Rule0R4
	Rules0R0
	Rules0R1
	Rules1R0
	Rules1R1
	Rules1R2
	Symbol0R0
	Symbol0R1
	Symbol1R0
	Symbol1R1
	Symbol2R0
	Symbol2R1
	Symbols0R0
	Symbols0R1
	Symbols1R0
	Symbols1R1
	Symbols1R2
	TokID0R0
	TokID0R1
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
    Label 	Label
}

type Index struct {
	NT      symbols.NT
	Alt     int
	Pos     int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
    Alternate0R0: {
        symbols.NT_Alternate, 0, 0, 
        symbols.Symbols{  
            symbols.NT_Symbols,
        }, 
        Alternate0R0, 
    },
    Alternate0R1: {
        symbols.NT_Alternate, 0, 1, 
        symbols.Symbols{  
            symbols.NT_Symbols,
        }, 
        Alternate0R1, 
    },
    Alternate1R0: {
        symbols.NT_Alternate, 1, 0, 
        symbols.Symbols{  
            symbols.T_2,
        }, 
        Alternate1R0, 
    },
    Alternate1R1: {
        symbols.NT_Alternate, 1, 1, 
        symbols.Symbols{  
            symbols.T_2,
        }, 
        Alternate1R1, 
    },
    Alternates0R0: {
        symbols.NT_Alternates, 0, 0, 
        symbols.Symbols{  
            symbols.NT_Alternate,
        }, 
        Alternates0R0, 
    },
    Alternates0R1: {
        symbols.NT_Alternates, 0, 1, 
        symbols.Symbols{  
            symbols.NT_Alternate,
        }, 
        Alternates0R1, 
    },
    Alternates1R0: {
        symbols.NT_Alternates, 1, 0, 
        symbols.Symbols{  
            symbols.NT_Alternate, 
            symbols.T_6, 
            symbols.NT_Alternates,
        }, 
        Alternates1R0, 
    },
    Alternates1R1: {
        symbols.NT_Alternates, 1, 1, 
        symbols.Symbols{  
            symbols.NT_Alternate, 
            symbols.T_6, 
            symbols.NT_Alternates,
        }, 
        Alternates1R1, 
    },
    Alternates1R2: {
        symbols.NT_Alternates, 1, 2, 
        symbols.Symbols{  
            symbols.NT_Alternate, 
            symbols.T_6, 
            symbols.NT_Alternates,
        }, 
        Alternates1R2, 
    },
    Alternates1R3: {
        symbols.NT_Alternates, 1, 3, 
        symbols.Symbols{  
            symbols.NT_Alternate, 
            symbols.T_6, 
            symbols.NT_Alternates,
        }, 
        Alternates1R3, 
    },
    GoGLL0R0: {
        symbols.NT_GoGLL, 0, 0, 
        symbols.Symbols{  
            symbols.NT_Package, 
            symbols.NT_Rules,
        }, 
        GoGLL0R0, 
    },
    GoGLL0R1: {
        symbols.NT_GoGLL, 0, 1, 
        symbols.Symbols{  
            symbols.NT_Package, 
            symbols.NT_Rules,
        }, 
        GoGLL0R1, 
    },
    GoGLL0R2: {
        symbols.NT_GoGLL, 0, 2, 
        symbols.Symbols{  
            symbols.NT_Package, 
            symbols.NT_Rules,
        }, 
        GoGLL0R2, 
    },
    NT0R0: {
        symbols.NT_NT, 0, 0, 
        symbols.Symbols{  
            symbols.T_3,
        }, 
        NT0R0, 
    },
    NT0R1: {
        symbols.NT_NT, 0, 1, 
        symbols.Symbols{  
            symbols.T_3,
        }, 
        NT0R1, 
    },
    Package0R0: {
        symbols.NT_Package, 0, 0, 
        symbols.Symbols{  
            symbols.T_4, 
            symbols.T_5,
        }, 
        Package0R0, 
    },
    Package0R1: {
        symbols.NT_Package, 0, 1, 
        symbols.Symbols{  
            symbols.T_4, 
            symbols.T_5,
        }, 
        Package0R1, 
    },
    Package0R2: {
        symbols.NT_Package, 0, 2, 
        symbols.Symbols{  
            symbols.T_4, 
            symbols.T_5,
        }, 
        Package0R2, 
    },
    Rule0R0: {
        symbols.NT_Rule, 0, 0, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_0, 
            symbols.NT_Alternates, 
            symbols.T_1,
        }, 
        Rule0R0, 
    },
    Rule0R1: {
        symbols.NT_Rule, 0, 1, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_0, 
            symbols.NT_Alternates, 
            symbols.T_1,
        }, 
        Rule0R1, 
    },
    Rule0R2: {
        symbols.NT_Rule, 0, 2, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_0, 
            symbols.NT_Alternates, 
            symbols.T_1,
        }, 
        Rule0R2, 
    },
    Rule0R3: {
        symbols.NT_Rule, 0, 3, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_0, 
            symbols.NT_Alternates, 
            symbols.T_1,
        }, 
        Rule0R3, 
    },
    Rule0R4: {
        symbols.NT_Rule, 0, 4, 
        symbols.Symbols{  
            symbols.NT_NT, 
            symbols.T_0, 
            symbols.NT_Alternates, 
            symbols.T_1,
        }, 
        Rule0R4, 
    },
    Rules0R0: {
        symbols.NT_Rules, 0, 0, 
        symbols.Symbols{  
            symbols.NT_Rule,
        }, 
        Rules0R0, 
    },
    Rules0R1: {
        symbols.NT_Rules, 0, 1, 
        symbols.Symbols{  
            symbols.NT_Rule,
        }, 
        Rules0R1, 
    },
    Rules1R0: {
        symbols.NT_Rules, 1, 0, 
        symbols.Symbols{  
            symbols.NT_Rule, 
            symbols.NT_Rules,
        }, 
        Rules1R0, 
    },
    Rules1R1: {
        symbols.NT_Rules, 1, 1, 
        symbols.Symbols{  
            symbols.NT_Rule, 
            symbols.NT_Rules,
        }, 
        Rules1R1, 
    },
    Rules1R2: {
        symbols.NT_Rules, 1, 2, 
        symbols.Symbols{  
            symbols.NT_Rule, 
            symbols.NT_Rules,
        }, 
        Rules1R2, 
    },
    Symbol0R0: {
        symbols.NT_Symbol, 0, 0, 
        symbols.Symbols{  
            symbols.NT_NT,
        }, 
        Symbol0R0, 
    },
    Symbol0R1: {
        symbols.NT_Symbol, 0, 1, 
        symbols.Symbols{  
            symbols.NT_NT,
        }, 
        Symbol0R1, 
    },
    Symbol1R0: {
        symbols.NT_Symbol, 1, 0, 
        symbols.Symbols{  
            symbols.NT_TokID,
        }, 
        Symbol1R0, 
    },
    Symbol1R1: {
        symbols.NT_Symbol, 1, 1, 
        symbols.Symbols{  
            symbols.NT_TokID,
        }, 
        Symbol1R1, 
    },
    Symbol2R0: {
        symbols.NT_Symbol, 2, 0, 
        symbols.Symbols{  
            symbols.T_5,
        }, 
        Symbol2R0, 
    },
    Symbol2R1: {
        symbols.NT_Symbol, 2, 1, 
        symbols.Symbols{  
            symbols.T_5,
        }, 
        Symbol2R1, 
    },
    Symbols0R0: {
        symbols.NT_Symbols, 0, 0, 
        symbols.Symbols{  
            symbols.NT_Symbol,
        }, 
        Symbols0R0, 
    },
    Symbols0R1: {
        symbols.NT_Symbols, 0, 1, 
        symbols.Symbols{  
            symbols.NT_Symbol,
        }, 
        Symbols0R1, 
    },
    Symbols1R0: {
        symbols.NT_Symbols, 1, 0, 
        symbols.Symbols{  
            symbols.NT_Symbol, 
            symbols.NT_Symbols,
        }, 
        Symbols1R0, 
    },
    Symbols1R1: {
        symbols.NT_Symbols, 1, 1, 
        symbols.Symbols{  
            symbols.NT_Symbol, 
            symbols.NT_Symbols,
        }, 
        Symbols1R1, 
    },
    Symbols1R2: {
        symbols.NT_Symbols, 1, 2, 
        symbols.Symbols{  
            symbols.NT_Symbol, 
            symbols.NT_Symbols,
        }, 
        Symbols1R2, 
    },
    TokID0R0: {
        symbols.NT_TokID, 0, 0, 
        symbols.Symbols{  
            symbols.T_3,
        }, 
        TokID0R0, 
    },
    TokID0R1: {
        symbols.NT_TokID, 0, 1, 
        symbols.Symbols{  
            symbols.T_3,
        }, 
        TokID0R1, 
    },
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_Alternate,0,0 }: Alternate0R0,
	Index{ symbols.NT_Alternate,0,1 }: Alternate0R1,
	Index{ symbols.NT_Alternate,1,0 }: Alternate1R0,
	Index{ symbols.NT_Alternate,1,1 }: Alternate1R1,
	Index{ symbols.NT_Alternates,0,0 }: Alternates0R0,
	Index{ symbols.NT_Alternates,0,1 }: Alternates0R1,
	Index{ symbols.NT_Alternates,1,0 }: Alternates1R0,
	Index{ symbols.NT_Alternates,1,1 }: Alternates1R1,
	Index{ symbols.NT_Alternates,1,2 }: Alternates1R2,
	Index{ symbols.NT_Alternates,1,3 }: Alternates1R3,
	Index{ symbols.NT_GoGLL,0,0 }: GoGLL0R0,
	Index{ symbols.NT_GoGLL,0,1 }: GoGLL0R1,
	Index{ symbols.NT_GoGLL,0,2 }: GoGLL0R2,
	Index{ symbols.NT_NT,0,0 }: NT0R0,
	Index{ symbols.NT_NT,0,1 }: NT0R1,
	Index{ symbols.NT_Package,0,0 }: Package0R0,
	Index{ symbols.NT_Package,0,1 }: Package0R1,
	Index{ symbols.NT_Package,0,2 }: Package0R2,
	Index{ symbols.NT_Rule,0,0 }: Rule0R0,
	Index{ symbols.NT_Rule,0,1 }: Rule0R1,
	Index{ symbols.NT_Rule,0,2 }: Rule0R2,
	Index{ symbols.NT_Rule,0,3 }: Rule0R3,
	Index{ symbols.NT_Rule,0,4 }: Rule0R4,
	Index{ symbols.NT_Rules,0,0 }: Rules0R0,
	Index{ symbols.NT_Rules,0,1 }: Rules0R1,
	Index{ symbols.NT_Rules,1,0 }: Rules1R0,
	Index{ symbols.NT_Rules,1,1 }: Rules1R1,
	Index{ symbols.NT_Rules,1,2 }: Rules1R2,
	Index{ symbols.NT_Symbol,0,0 }: Symbol0R0,
	Index{ symbols.NT_Symbol,0,1 }: Symbol0R1,
	Index{ symbols.NT_Symbol,1,0 }: Symbol1R0,
	Index{ symbols.NT_Symbol,1,1 }: Symbol1R1,
	Index{ symbols.NT_Symbol,2,0 }: Symbol2R0,
	Index{ symbols.NT_Symbol,2,1 }: Symbol2R1,
	Index{ symbols.NT_Symbols,0,0 }: Symbols0R0,
	Index{ symbols.NT_Symbols,0,1 }: Symbols0R1,
	Index{ symbols.NT_Symbols,1,0 }: Symbols1R0,
	Index{ symbols.NT_Symbols,1,1 }: Symbols1R1,
	Index{ symbols.NT_Symbols,1,2 }: Symbols1R2,
	Index{ symbols.NT_TokID,0,0 }: TokID0R0,
	Index{ symbols.NT_TokID,0,1 }: TokID0R1,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_GoGLL:[]Label{ GoGLL0R0 },
	symbols.NT_Package:[]Label{ Package0R0 },
	symbols.NT_Rules:[]Label{ Rules0R0,Rules1R0 },
	symbols.NT_Rule:[]Label{ Rule0R0 },
	symbols.NT_NT:[]Label{ NT0R0 },
	symbols.NT_Alternates:[]Label{ Alternates0R0,Alternates1R0 },
	symbols.NT_Alternate:[]Label{ Alternate0R0,Alternate1R0 },
	symbols.NT_Symbols:[]Label{ Symbols0R0,Symbols1R0 },
	symbols.NT_Symbol:[]Label{ Symbol0R0,Symbol1R0,Symbol2R0 },
	symbols.NT_TokID:[]Label{ TokID0R0 },
}

