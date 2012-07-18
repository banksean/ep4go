package ftree

import (
	"testing"
)


func Test_Eval(t *testing.T) {
	f := NewFTree("add", Const(1), Const(2))
	if v, err := f.Eval(); v != 3 {
		t.Error("1 + 2 = 3: %v", err)
	}
}

func Test_Eval_Vars(t *testing.T) {
	f := NewFTree("add", Var("x"), Var("y"), Const(2))
	SetVar("x", 1)
	SetVar("y", 2)
	if v, err := f.Eval(); v != 5 {
		t.Error("1 + 2 + 2 = 5: %v", err)
	}
}

func Test_Eval_Nested(t *testing.T) {
	f := NewFTree("add", Var("x"), Var("y"), NewFTree("add", Var("x"), Var("y"), Const(2)))
	SetVar("x", 1)
	SetVar("y", 2)
	if v, err := f.Eval(); v != 8 {
		t.Error("1 + 2 + 1 + 2 + 2  = 8: %v, %v", v, err)
	}
}
