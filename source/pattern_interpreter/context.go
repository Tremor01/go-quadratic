package pattern_interpreter

/*
| := or
* := one or more repeats for token
[TOKEN] := optional token

NAN := n a n
INF := i n f

NOT_ZERO := 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
ZERO := 0
DIGIT := NOT_ZERO | ZERO

SIGNS := - | +
POINT := .

EXP := e | E
EXP_TOKEN := EXP [SIGNS] DIGIT*

N := NOT_ZERO [DIGIT*]
Z := [SIGNS] ZERO | N 
R := Z POINT DIGIT* | Z

R_SHORT1 := [SIGNS] POINT DIGIT*
R_SHORT2 := Z POINT

NUMBER := R | SHORT_R1 | SHORT_R2 [EXP_TOKEN]

*/


var NOT_ZERO = []rune {'1', '2', '3', '4', '5', '6', '7', '8', '9'} 
var ZERO = []rune {'0'} 

var SIGNS = []rune {'+', '-'} 
var POINT = []rune {'.'}
var EXP = []rune {'e', 'E'}

var NAN = []rune {'n', 'a', 'n'}
var INF = []rune {'i', 'n', 'f'}

type Context struct {
	row []rune
	pos int
	str string
}

func NewContext(row string) *Context {
	return &Context{
		row: []rune(row), 
		pos: 0,
		str: row,
	}
}

func (c Context) Read() rune {
	return c.row[c.pos]
}

func (c Context) IsEnd() bool {
	return c.pos == len(c.row)
}

func (c *Context) Next() {
	c.pos++
}

func (c Context) GetRow() string {
	return c.str
}

func (c *Context) ResetPos() {
	c.pos = 0
}