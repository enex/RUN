// Copyright (c) 2014, Rob Thornton
// All rights reserved.
// This source code is governed by a Simplied BSD-License. Please see the
// LICENSE included in this distribution for a copy of the full license
// or, if one is not included, you may also find a copy at
// http://opensource.org/licenses/BSD-2-Clause

package token

import "fmt"

type Pos uint

var NoPos = Pos(0)

// Valid returns true if p does not equal NoPos
func (p Pos) Valid() bool {
	return p != NoPos
}

//Represents the absolute position, usefull for error messages
type Position struct {
	Filename string
	Col, Row int
}

func (p Position) String() string {
	if p.Filename == "" {
		return fmt.Sprintf("%d:%d", p.Row, p.Col)
	}
	return fmt.Sprintf("%s:%d:%d", p.Filename, p.Row, p.Col)
}

func (p Position) Pos() Position {
	return p
}
