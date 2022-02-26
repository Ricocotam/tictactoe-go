package tictactoe

import (
	"testing"
)

func TestCheckMoveNegative(t *testing.T) {
	move := new(Pos)
	move.x = -1
	move.y = -1

	checkMove(move)
	if move.isValid {
		t.Errorf("CheckMove should return false and an error and not %t", move.isValid)
	}
}

func TestCheckMoveTooBig(t *testing.T) {
	move := new(Pos)
	move.x = 4
	move.y = 4

	checkMove(move)

	if move.isValid {
		t.Errorf("CheckMove should return false and an error and not %t", move.isValid)
	}
}

func TestAllEqualTrue(t *testing.T) {
	var tests = []string{"", "a", "aaa"}

	for _, test := range tests {
		if !allEqual(test) {
			t.Errorf("allEqual should return true to input %s", test)
		}
	}
}

func TestAllEqualFalse(t *testing.T) {
	tests := []string{"aba", "aab", "ab"}

	for _, test := range tests {
		if allEqual(test) {
			t.Errorf("allEqual should return false to input %s", test)
		}
	}
}

func TestCheckLines(t *testing.T) {
	lines := []string{"+++", "ooo", "xxx"}

	won := checkLines(lines)
	if !won {
		t.Errorf("checkLines returned false instead of true")
	}

	lines = []string{"+x", "oxo", "x"}
	won = checkLines(lines)
	if won {
		t.Errorf("checkLines returned true instead of false")
	}
}

// Rows

func TestCheckRowsEmpty(t *testing.T) {
	var empty_board Grid
	if won := empty_board.checkRows(); won {
		t.Errorf("checkRows returned true instead of false")
	}
}

func TestCheckRowsOneWon(t *testing.T) {
	onerow_won := Grid{
		[3]string{"o", "o", "x"},
		[3]string{"x", "x", "x"},
		[3]string{"o", "x", "o"},
	}

	if won := onerow_won.checkRows(); !won {
		t.Errorf("checkRows returned false instead of true")
	}

	onerow_won = Grid{
		[3]string{"", "", ""},
		[3]string{"x", "x", "x"},
		[3]string{"", "", ""},
	}
	if won := onerow_won.checkRows(); !won {
		t.Errorf("checkRows returned false instead of true")
	}
}

func TestCheckRowsOneNoWin(t *testing.T) {
	onerow_nowin := Grid{
		[3]string{"x", "o", "o"},
		[3]string{"o", "x", "o"},
		[3]string{"o", "o", "x"},
	}

	if won := onerow_nowin.checkRows(); won {
		t.Errorf("checkRows returned true instead of false")
	}

	onerow_nowin = Grid{
		[3]string{"", "", ""},
		[3]string{"o", "x", "o"},
		[3]string{"o", "o", "x"},
	}

	if won := onerow_nowin.checkRows(); won {
		t.Errorf("checkRows returned true instead of false")
	}
}

func TestCheckRowsAllRowWon(t *testing.T) {
	allrow_won := Grid{
		[3]string{"x", "x", "x"},
		[3]string{"x", "x", "x"},
		[3]string{"x", "x", "x"},
	}

	if won := allrow_won.checkRows(); !won {
		t.Errorf("checkRows returned false instead of true")
	}
}
func TestCheckRowsAllRowTie(t *testing.T) {
	allrow_tie := Grid{
		[3]string{"x", "o", "x"},
		[3]string{"x", "o", "x"},
		[3]string{"o", "x", "o"},
	}

	if won := allrow_tie.checkRows(); won {
		t.Errorf("checkRows returned true instead of false")
	}
}

// Columns
func TestCheckColumnsEmpty(t *testing.T) {
	var empty_board Grid
	if won := empty_board.checkColumns(); won {
		t.Errorf("checkColumns returned true instead of false")
	}
}

func TestCheckColumnsOneWon(t *testing.T) {
	onecol_won := Grid{
		[3]string{"o", "o", "x"},
		[3]string{"x", "o", "x"},
		[3]string{"o", "o", "o"},
	}

	if won := onecol_won.checkColumns(); !won {
		t.Errorf("checkColumns returned false instead of true")
	}

	onecol_won = Grid{
		[3]string{"", "o", ""},
		[3]string{"x", "o", ""},
		[3]string{"", "o", ""},
	}
	if won := onecol_won.checkColumns(); !won {
		t.Errorf("checkColumns returned false instead of true")
	}
}

func TestCheckColumnsOneNoWin(t *testing.T) {
	onecol_nowin := Grid{
		[3]string{"x", "o", "o"},
		[3]string{"o", "x", "o"},
		[3]string{"o", "o", "x"},
	}

	if won := onecol_nowin.checkColumns(); won {
		t.Errorf("checkColumns returned true instead of false")
	}

	onecol_nowin = Grid{
		[3]string{"x", "o", ""},
		[3]string{"o", "x", ""},
		[3]string{"o", "o", ""},
	}

	if won := onecol_nowin.checkColumns(); won {
		t.Errorf("checkColumns returned true instead of false")
	}
}

func TestCheckColumnsAllColsWon(t *testing.T) {
	allcol_won := Grid{
		[3]string{"x", "x", "x"},
		[3]string{"x", "x", "x"},
		[3]string{"x", "x", "x"},
	}

	if won := allcol_won.checkColumns(); !won {
		t.Errorf("checkColumns returned false instead of true")
	}
}
func TestCheckColumnsAllColTie(t *testing.T) {
	allcol_tie := Grid{
		[3]string{"x", "o", "x"},
		[3]string{"x", "o", "x"},
		[3]string{"o", "x", "o"},
	}

	if won := allcol_tie.checkColumns(); won {
		t.Errorf("checkColumns returned true instead of false")
	}
}

// Diagonals

func TestCheckDiagonalsEmpty(t *testing.T) {
	var empty_board Grid
	if won := empty_board.checkDiagonals(); won {
		t.Errorf("checkDiagonals returned true instead of false")
	}
}

func TestCheckDiagonalsOneWon(t *testing.T) {
	onediag_won := Grid{
		[3]string{"x", "o", "x"},
		[3]string{"x", "x", "o"},
		[3]string{"o", "x", "x"},
	}

	if won := onediag_won.checkDiagonals(); !won {
		t.Errorf("checkDiagonals returned false instead of true")
	}

	onediag_won = Grid{
		[3]string{"o", "o", "x"},
		[3]string{"x", "x", "o"},
		[3]string{"x", "x", "o"},
	}

	if won := onediag_won.checkDiagonals(); !won {
		t.Errorf("checkDiagonals returned false instead of true")
	}

	onediag_won = Grid{
		[3]string{"x", "", ""},
		[3]string{"", "x", ""},
		[3]string{"", "", "x"},
	}
	if won := onediag_won.checkDiagonals(); !won {
		t.Errorf("checkDiagonals returned false instead of true")
	}
}

func TestCheckDiagonalsOneNoWin(t *testing.T) {
	onediag_nowin := Grid{
		[3]string{"x", "o", "o"},
		[3]string{"o", "o", "x"},
		[3]string{"x", "x", "o"},
	}

	if won := onediag_nowin.checkDiagonals(); won {
		t.Errorf("checkDiagonals returned true instead of false")
	}

	onediag_nowin = Grid{
		[3]string{"o", "", ""},
		[3]string{"", "x", "o"},
		[3]string{"", "o", "x"},
	}

	if won := onediag_nowin.checkDiagonals(); won {
		t.Errorf("checkDiagonals returned true instead of false")
	}
}

func TestCheckDiagonalsAllDiagWon(t *testing.T) {
	alldiag_won := Grid{
		[3]string{"x", "x", "x"},
		[3]string{"x", "x", "x"},
		[3]string{"x", "x", "x"},
	}

	if won := alldiag_won.checkDiagonals(); !won {
		t.Errorf("checkDiagonals returned false instead of true")
	}
}
func TestCheckDiagonalsAllDiagTie(t *testing.T) {
	alldiag_tie := Grid{
		[3]string{"x", "o", "x"},
		[3]string{"x", "o", "x"},
		[3]string{"o", "x", "o"},
	}

	if won := alldiag_tie.checkDiagonals(); won {
		t.Errorf("checkDiagonals returned true instead of false")
	}
}
