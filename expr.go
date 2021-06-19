package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) visitBadExpr(n *ast.BadExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BadExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in BadExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in BadExpr (post)")
		}()
	}

	return
}

func (v *Visitor) visitIdent(n *ast.Ident, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Ident; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in Ident (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Ident (post)")
		}()
	}

	return
}

func (v *Visitor) visitEllipsis(n *ast.Ellipsis, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Ellipsis; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in Ellipsis (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Ellipsis (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Elt, Ellipsis_Elt, 0, stack2)
	if err != nil {
		return err
	}

	return
}

func (v *Visitor) visitBasicLit(n *ast.BasicLit, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BasicLit; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in BasicLit (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in BasicLit (post)")
		}()
	}

	return
}

func (v *Visitor) visitFuncLit(n *ast.FuncLit, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.FuncLit; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in FuncLit (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in FuncLit (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitFuncType(n.Type, FuncLit_Type, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitBlockStmt(n.Body, FuncLit_Body, 0, stack2)

	return
}

func (v *Visitor) visitCompositeLit(n *ast.CompositeLit, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CompositeLit; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in CompositeLit (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in CompositeLit (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Type, CompositeLit_Type, 0, stack2)
	if err != nil {
		return err
	}

	for i, expr := range n.Elts {
		err = v.visitExpr(expr, CompositeLit_Elts, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitParenExpr(n *ast.ParenExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ParenExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in ParenExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in ParenExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, ParenExpr_X, 0, stack2)

	return
}

func (v *Visitor) visitSelectorExpr(n *ast.SelectorExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SelectorExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in SelectorExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in SelectorExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, SelectorExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitIdent(n.Sel, SelectorExpr_Sel, 0, stack2)

	return
}

func (v *Visitor) visitIndexExpr(n *ast.IndexExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.IndexExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in IndexExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in IndexExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, IndexExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Index, IndexExpr_Index, 0, stack2)

	return
}

func (v *Visitor) visitSliceExpr(n *ast.SliceExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SliceExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in SliceExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in SliceExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, SliceExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Low, SliceExpr_Low, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.High, SliceExpr_High, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Max, SliceExpr_Max, 0, stack2)
	if err != nil {
		return err
	}

	return
}

func (v *Visitor) visitTypeAssertExpr(n *ast.TypeAssertExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.TypeAssertExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in TypeAssertExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in TypeAssertExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, TypeAssertExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Type, TypeAssertExpr_Type, 0, stack2)

	return
}

func (v *Visitor) visitCallExpr(n *ast.CallExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CallExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in CallExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in CallExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Fun, CallExpr_Fun, 0, stack2)
	if err != nil {
		return err
	}

	for i, expr := range n.Args {
		err = v.visitExpr(expr, CallExpr_Args, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitStarExpr(n *ast.StarExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.StarExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in StarExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in StarExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, StarExpr_X, 0, stack2)

	return
}

func (v *Visitor) visitUnaryExpr(n *ast.UnaryExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.UnaryExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in UnaryExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in UnaryExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, UnaryExpr_X, 0, stack2)

	return
}

func (v *Visitor) visitBinaryExpr(n *ast.BinaryExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BinaryExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in BinaryExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in BinaryExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, BinaryExpr_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Y, BinaryExpr_Y, 0, stack2)

	return
}

func (v *Visitor) visitKeyValueExpr(n *ast.KeyValueExpr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.KeyValueExpr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in KeyValueExpr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in KeyValueExpr (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Key, KeyValueExpr_Key, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Value, KeyValueExpr_Value, 0, stack2)

	return
}

func (v *Visitor) visitArrayType(n *ast.ArrayType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ArrayType; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in ArrayType (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in ArrayType (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Len, ArrayType_Len, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Elt, ArrayType_Elt, 0, stack2)

	return
}

func (v *Visitor) visitStructType(n *ast.StructType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.StructType; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in StructType (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in StructType (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitFieldList(n.Fields, StructType_Fields, 0, stack2)

	return
}

func (v *Visitor) visitFuncType(n *ast.FuncType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.FuncType; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in FuncType (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in FuncType (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitFieldList(n.Params, FuncType_Params, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitFieldList(n.Results, FuncType_Results, 0, stack2)

	return
}

func (v *Visitor) visitInterfaceType(n *ast.InterfaceType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.InterfaceType; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in InterfaceType (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in InterfaceType (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitFieldList(n.Methods, InterfaceType_Methods, 0, stack2)

	return
}

func (v *Visitor) visitMapType(n *ast.MapType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.MapType; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in MapType (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in MapType (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Key, MapType_Key, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Value, MapType_Value, 0, stack2)

	return
}

func (v *Visitor) visitChanType(n *ast.ChanType, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ChanType; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in ChanType (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in ChanType (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Value, ChanType_Value, 0, stack2)

	return
}
