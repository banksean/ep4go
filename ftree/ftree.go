package ftree

import (
  "fmt"
//	"log"
	"math"
	"math/rand"
)

var varVals = map[string]float64{}

type FTree struct {
	op string
	args []*FTree
	constVal float64
	varName string
}

func (f *FTree) String() string {
  s := fmt.Sprintf("op: %q, constVal: %f, varName: %q\n\t", f.op, f.constVal, f.varName)
  s += fmt.Sprintf("%v", f.args)
//  for _, arg := range f.args {
//    s += arg.String()
//  }
  return s
}

func (f *FTree) Eval() (float64, error) {
	switch f.op {
		case "const":
			return f.constVal, nil
		case "var":
			return varVals[f.varName], nil
		case "add":
			ret := 0.0
			for _, arg := range f.args {
				v, err := arg.Eval()
				if err != nil {
					return 0.0, err
				}
				ret = ret + v
			}
			return ret, nil
		case "sub":
			ret := 0.0
			for _, arg := range f.args {
				v, err := arg.Eval()
				if err != nil {
					return 0.0, err
				}
				ret = ret - v
			}
			return ret, nil
		case "mul":
			ret := 1.0
			for _, arg := range f.args {
				v, err := arg.Eval()
				if err != nil {
					return 0.0, err
				}
				ret = ret * v
			}
			return ret, nil
	}
	return 0.0, fmt.Errorf("Unknown op: %q", f.op)
}

func (f *FTree) EvalSigmoid() (uint8, error) {
	v, err := f.Eval()
	return uint8(Sigmoid(v) * 256), err
}

func randOp() string {
  ops := [...]string{"add", "sub", "mul", "const", "var"}
  return ops[rand.Int() % len(ops)]
}

func NewRandomFTree(depth int32) *FTree {
  randConst := rand.Float64()
  randVar := "x"
  if rand.Float64() > 0.5 {
    randVar = "y"
  }
	randOp := randOp()
	randArgs := make([]*FTree, rand.Uint32() % 5)
  if depth > 0 {
    for i, _ := range randArgs {
      randArgs[i] = NewRandomFTree(depth-1)
    }
  }

	return &FTree{randOp, randArgs, randConst, randVar}
}

func NewFTree(op string, args ...*FTree) *FTree {
	return &FTree{op, args, 0.0, ""}
}

func Const(v float64) *FTree {
	return &FTree{"const", nil, v, ""}
}

func Var(name string) *FTree {
	return &FTree{"var", nil, 0.0, name}
}

func SetVar(name string, val float64) {
	varVals[name] = val
}

func Sigmoid(t float64) float64 {
	return 1.0 / (1 + math.Exp(t))
}