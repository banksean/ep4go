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

func Test_Eval_Mul(t *testing.T) {
	f := NewFTree("mul", Const(2))
	if v, err := f.Eval(); v != 2 {
		t.Error("1 * 2  = 2: %v, %v", v, err)
	}

	f = NewFTree("mul", Const(2), Var("x"))
	SetVar("x", 4)
	if v, err := f.Eval(); v != 8 {
		t.Error("2 * 4  = 8: %v, %v", v, err)
	}
}
/*
func Test_Recur(t *testing.T) {
  t.Log("depth 4")
  Recur(4)
}
*/
func Test_RandomTree(t *testing.T) {
  t.Log("New randomTree")

  f := NewRandomFTree(4)
  t.Log("Tree: %v", f)

/*
	v, err := f.Eval()
	if v != 8 {
		t.Error("2 * 4  = 8: %v, %v", v, err)
	}
  fmt.Printf("Tree: %v", v)
*/
}
