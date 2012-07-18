package ftree

import (
	"fmt"
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
			ret := 0.0
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

func NewRandomFTree(depth int32) *FTree {
	randOp := "add"
	randArgs := make([]*FTree, rand.Uint32() % 10)
	randConst := 0.0
	randVar := "x"
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