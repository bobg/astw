package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) VisitBadExpr(n *ast.BadExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BadExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in BadExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in BadExpr (post)")
			}
		}()
	}

	return
}

func (v *Visitor) VisitIdent(n *ast.Ident, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Ident; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Ident (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Ident (post)")
			}
		}()
	}

	return
}

func (v *Visitor) VisitEllipsis(n *ast.Ellipsis, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Ellipsis; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Ellipsis (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Ellipsis (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Elt, Ellipsis_Elt, 0, stack2)
	if err != nil {
		return err
	}

	return
}

func (v *Visitor) VisitBasicLit(n *ast.BasicLit, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BasicLit; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in BasicLit (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in BasicLit (post)")
			}
		}()
	}

	return
}

func (v *Visitor) VisitFuncLit(n *ast.FuncLit, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.FuncLit; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in FuncLit (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in FuncLit (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitFuncType(n.Type, FuncLit_Type, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitBlockStmt(n.Body, FuncLit_Body, 0, stack2)

	return
}

func (v *Visitor) VisitCompositeLit(n *ast.CompositeLit, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CompositeLit; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in CompositeLit (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in CompositeLit (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Type, CompositeLit_Type, 0, stack2)
	if err != nil {
		return err
	}

	for i, expr := range n.Elts {
		err = v.VisitExpr(expr, CompositeLit_Elts, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitParenExpr(n *ast.ParenExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ParenExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in ParenExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in ParenExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, ParenExpr_X, 0, stack2)

	return
}

func (v *Visitor) VisitSelectorExpr(n *ast.SelectorExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SelectorExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in SelectorExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in SelectorExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, SelectorExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitIdent(n.Sel, SelectorExpr_Sel, 0, stack2)

	return
}

func (v *Visitor) VisitIndexExpr(n *ast.IndexExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.IndexExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in IndexExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in IndexExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, IndexExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Index, IndexExpr_Index, 0, stack2)

	return
}

func (v *Visitor) VisitSliceExpr(n *ast.SliceExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SliceExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in SliceExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in SliceExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, SliceExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Low, SliceExpr_Low, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.High, SliceExpr_High, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Max, SliceExpr_Max, 0, stack2)
	if err != nil {
		return err
	}

	return
}

func (v *Visitor) VisitTypeAssertExpr(n *ast.TypeAssertExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.TypeAssertExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in TypeAssertExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in TypeAssertExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, TypeAssertExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Type, TypeAssertExpr_Type, 0, stack2)

	return
}

func (v *Visitor) VisitCallExpr(n *ast.CallExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CallExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in CallExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in CallExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Fun, CallExpr_Fun, 0, stack2)
	if err != nil {
		return err
	}

	for i, expr := range n.Args {
		err = v.VisitExpr(expr, CallExpr_Args, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitStarExpr(n *ast.StarExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.StarExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in StarExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in StarExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, StarExpr_X, 0, stack2)

	return
}

func (v *Visitor) VisitUnaryExpr(n *ast.UnaryExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.UnaryExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in UnaryExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in UnaryExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, UnaryExpr_X, 0, stack2)

	return
}

func (v *Visitor) VisitBinaryExpr(n *ast.BinaryExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BinaryExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in BinaryExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in BinaryExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, BinaryExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Y, BinaryExpr_Y, 0, stack2)

	return
}

func (v *Visitor) VisitKeyValueExpr(n *ast.KeyValueExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.KeyValueExpr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in KeyValueExpr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in KeyValueExpr (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Key, KeyValueExpr_Key, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Value, KeyValueExpr_Value, 0, stack2)

	return
}

func (v *Visitor) VisitArrayType(n *ast.ArrayType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ArrayType; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in ArrayType (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in ArrayType (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Len, ArrayType_Len, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Elt, ArrayType_Elt, 0, stack2)

	return
}

func (v *Visitor) VisitStructType(n *ast.StructType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.StructType; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in StructType (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in StructType (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitFieldList(n.Fields, StructType_Fields, 0, stack2)

	return
}

func (v *Visitor) VisitFuncType(n *ast.FuncType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.FuncType; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in FuncType (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in FuncType (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitFieldList(n.Params, FuncType_Params, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitFieldList(n.Results, FuncType_Results, 0, stack2)

	return
}

func (v *Visitor) VisitInterfaceType(n *ast.InterfaceType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.InterfaceType; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in InterfaceType (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in InterfaceType (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitFieldList(n.Methods, InterfaceType_Methods, 0, stack2)

	return
}

func (v *Visitor) VisitMapType(n *ast.MapType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.MapType; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in MapType (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in MapType (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Key, MapType_Key, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Value, MapType_Value, 0, stack2)

	return
}

func (v *Visitor) VisitChanType(n *ast.ChanType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ChanType; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in ChanType (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in ChanType (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Value, ChanType_Value, 0, stack2)

	return
}
