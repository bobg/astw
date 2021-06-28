package astw

import "go/ast"

func (v *Visitor) visitBadExpr(n *ast.BadExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.BadExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		nil,
	)
}

func (v *Visitor) visitIdent(n *ast.Ident, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.Ident; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		nil,
	)
}

func (v *Visitor) visitEllipsis(n *ast.Ellipsis, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.Ellipsis; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		nil,
	)
}

func (v *Visitor) visitBasicLit(n *ast.BasicLit, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.BasicLit; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		nil,
	)
}

func (v *Visitor) visitFuncLit(n *ast.FuncLit, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.FuncLit; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitFuncType(n.Type, FuncLit_Type, 0, stack2)
			if err != nil {
				return err
			}

			return v.visitBlockStmt(n.Body, FuncLit_Body, 0, stack2)
		},
	)
}

func (v *Visitor) visitCompositeLit(n *ast.CompositeLit, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.CompositeLit; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.Type, CompositeLit_Type, 0, stack2)
			if err != nil {
				return err
			}
			for i, expr := range n.Elts {
				err = v.visitAbstractExpr(expr, CompositeLit_Elts, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitParenExpr(n *ast.ParenExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.ParenExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitAbstractExpr(n.X, ParenExpr_X, 0, stack2)
		},
	)
}

func (v *Visitor) visitSelectorExpr(n *ast.SelectorExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.SelectorExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.X, SelectorExpr_X, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitIdent(n.Sel, SelectorExpr_Sel, 0, stack2)
		},
	)
}

func (v *Visitor) visitIndexExpr(n *ast.IndexExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.IndexExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.X, IndexExpr_X, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractExpr(n.Index, IndexExpr_Index, 0, stack2)
		},
	)
}

func (v *Visitor) visitSliceExpr(n *ast.SliceExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.SliceExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.X, SliceExpr_X, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractExpr(n.Low, SliceExpr_Low, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractExpr(n.High, SliceExpr_High, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractExpr(n.Max, SliceExpr_Max, 0, stack2)
		},
	)
}

func (v *Visitor) visitTypeAssertExpr(n *ast.TypeAssertExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.TypeAssertExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.X, TypeAssertExpr_X, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractExpr(n.Type, TypeAssertExpr_Type, 0, stack2)
		},
	)
}

func (v *Visitor) visitCallExpr(n *ast.CallExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.CallExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.Fun, CallExpr_Fun, 0, stack2)
			if err != nil {
				return err
			}
			for i, expr := range n.Args {
				err = v.visitAbstractExpr(expr, CallExpr_Args, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitStarExpr(n *ast.StarExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.StarExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitAbstractExpr(n.X, StarExpr_X, 0, stack2)
		},
	)
}

func (v *Visitor) visitUnaryExpr(n *ast.UnaryExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.UnaryExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitAbstractExpr(n.X, UnaryExpr_X, 0, stack2)
		},
	)
}

func (v *Visitor) visitBinaryExpr(n *ast.BinaryExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.BinaryExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.X, BinaryExpr_X, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractExpr(n.Y, BinaryExpr_Y, 0, stack2)
		},
	)
}

func (v *Visitor) visitKeyValueExpr(n *ast.KeyValueExpr, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.KeyValueExpr; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.Key, KeyValueExpr_Key, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractExpr(n.Value, KeyValueExpr_Value, 0, stack2)
		},
	)
}

func (v *Visitor) visitArrayType(n *ast.ArrayType, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.ArrayType; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.Len, ArrayType_Len, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractExpr(n.Elt, ArrayType_Elt, 0, stack2)
		},
	)
}

func (v *Visitor) visitStructType(n *ast.StructType, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.StructType; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitFieldList(n.Fields, StructType_Fields, 0, stack2)
		},
	)
}

func (v *Visitor) visitFuncType(n *ast.FuncType, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.FuncType; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitFieldList(n.Params, FuncType_Params, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitFieldList(n.Results, FuncType_Results, 0, stack2)
		},
	)
}

func (v *Visitor) visitInterfaceType(n *ast.InterfaceType, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.InterfaceType; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitFieldList(n.Methods, InterfaceType_Methods, 0, stack2)
		},
	)
}

func (v *Visitor) visitMapType(n *ast.MapType, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.MapType; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.Key, MapType_Key, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractExpr(n.Value, MapType_Value, 0, stack2)
		},
	)
}

func (v *Visitor) visitChanType(n *ast.ChanType, which Which, index int, stack []StackItem) error {
	return v.visitConcreteExpr(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.ChanType; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitAbstractExpr(n.Value, ChanType_Value, 0, stack2)
		},
	)
}
