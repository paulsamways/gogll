
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"empty/parser/symbols"
)

type Label int

const(
	A10R0 Label = iota
	A10R1
	A10R2
	Name0R0
	Name0R1
	Name1R0
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
	A10R0: {
		symbols.NT_A1, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Name, 
			symbols.T_0,
		}, 
		A10R0, 
	},
	A10R1: {
		symbols.NT_A1, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Name, 
			symbols.T_0,
		}, 
		A10R1, 
	},
	A10R2: {
		symbols.NT_A1, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Name, 
			symbols.T_0,
		}, 
		A10R2, 
	},
	Name0R0: {
		symbols.NT_Name, 0, 0, 
		symbols.Symbols{  
			symbols.T_1,
		}, 
		Name0R0, 
	},
	Name0R1: {
		symbols.NT_Name, 0, 1, 
		symbols.Symbols{  
			symbols.T_1,
		}, 
		Name0R1, 
	},
	Name1R0: {
		symbols.NT_Name, 1, 0, 
		symbols.Symbols{ 
		}, 
		Name1R0, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_A1,0,0 }: A10R0,
	Index{ symbols.NT_A1,0,1 }: A10R1,
	Index{ symbols.NT_A1,0,2 }: A10R2,
	Index{ symbols.NT_Name,0,0 }: Name0R0,
	Index{ symbols.NT_Name,0,1 }: Name0R1,
	Index{ symbols.NT_Name,1,0 }: Name1R0,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_A1:[]Label{ A10R0 },
	symbols.NT_Name:[]Label{ Name0R0,Name1R0 },
}

