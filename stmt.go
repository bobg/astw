package astw

import "go/ast"

func (v *Visitor) visitBadStmt(n *ast.BadStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.BadStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		nil,
	)
}

func (v *Visitor) visitDeclStmt(n *ast.DeclStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.DeclStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitAbstractDecl(n.Decl, DeclStmt_Decl, 0, stack2)
		},
	)
}

func (v *Visitor) visitEmptyStmt(n *ast.EmptyStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.EmptyStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		nil,
	)
}

func (v *Visitor) visitLabeledStmt(n *ast.LabeledStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.LabeledStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitIdent(n.Label, LabeledStmt_Label, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractStmt(n.Stmt, LabeledStmt_Stmt, 0, stack2)
		},
	)
}

func (v *Visitor) visitExprStmt(n *ast.ExprStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.ExprStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitAbstractExpr(n.X, ExprStmt_Expr, 0, stack2)
		},
	)
}

func (v *Visitor) visitSendStmt(n *ast.SendStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.SendStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.Chan, SendStmt_Chan, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractExpr(n.Value, SendStmt_Value, 0, stack2)
		},
	)
}

func (v *Visitor) visitIncDecStmt(n *ast.IncDecStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.IncDecStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitAbstractExpr(n.X, IncDecStmt_X, 0, stack2)
		},
	)
}

func (v *Visitor) visitAssignStmt(n *ast.AssignStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.AssignStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			for i, expr := range n.Lhs {
				err := v.visitAbstractExpr(expr, AssignStmt_Lhs, i, stack2)
				if err != nil {
					return err
				}
			}
			for i, expr := range n.Rhs {
				err := v.visitAbstractExpr(expr, AssignStmt_Rhs, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitGoStmt(n *ast.GoStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.GoStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitCallExpr(n.Call, GoStmt_Call, 0, stack2)
		},
	)
}

func (v *Visitor) visitDeferStmt(n *ast.DeferStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.DeferStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitCallExpr(n.Call, DeferStmt_Call, 0, stack2)
		},
	)
}

func (v *Visitor) visitReturnStmt(n *ast.ReturnStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.ReturnStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			for i, expr := range n.Results {
				err := v.visitAbstractExpr(expr, ReturnStmt_Results, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitBranchStmt(n *ast.BranchStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.BranchStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitIdent(n.Label, BranchStmt_Label, 0, stack2)
		},
	)
}

func (v *Visitor) visitBlockStmt(n *ast.BlockStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.BlockStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			for i, stmt := range n.List {
				err := v.visitAbstractStmt(stmt, BlockStmt_List, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitIfStmt(n *ast.IfStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.IfStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractStmt(n.Init, IfStmt_Init, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractExpr(n.Cond, IfStmt_Cond, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitBlockStmt(n.Body, IfStmt_Body, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitAbstractStmt(n.Else, IfStmt_Else, 0, stack2)
		},
	)
}

func (v *Visitor) visitCaseClause(n *ast.CaseClause, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.CaseClause; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			for i, expr := range n.List {
				err := v.visitAbstractExpr(expr, CaseClause_List, i, stack2)
				if err != nil {
					return err
				}
			}
			for i, stmt := range n.Body {
				err := v.visitAbstractStmt(stmt, CaseClause_Body, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitSwitchStmt(n *ast.SwitchStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.SwitchStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractStmt(n.Init, SwitchStmt_Init, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractExpr(n.Tag, SwitchStmt_Tag, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitBlockStmt(n.Body, SwitchStmt_Body, 0, stack2)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

func (v *Visitor) visitTypeSwitchStmt(n *ast.TypeSwitchStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.TypeSwitchStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractStmt(n.Init, TypeSwitchStmt_Init, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractStmt(n.Assign, TypeSwitchStmt_Assign, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitBlockStmt(n.Body, TypeSwitchStmt_Body, 0, stack2)
		},
	)
}

func (v *Visitor) visitCommClause(n *ast.CommClause, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.CommClause; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractStmt(n.Comm, CommClause_Comm, 0, stack2)
			if err != nil {
				return err
			}
			for i, stmt := range n.Body {
				err = v.visitAbstractStmt(stmt, CommClause_Body, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitSelectStmt(n *ast.SelectStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.SelectStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			return v.visitBlockStmt(n.Body, SelectStmt_Body, 0, stack2)
		},
	)
}

func (v *Visitor) visitForStmt(n *ast.ForStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.ForStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractStmt(n.Init, ForStmt_Init, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractExpr(n.Cond, ForStmt_Cond, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractStmt(n.Post, ForStmt_Post, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitBlockStmt(n.Body, ForStmt_Body, 0, stack2)
		},
	)
}

func (v *Visitor) visitRangeStmt(n *ast.RangeStmt, which Which, index int, stack []StackItem) error {
	return v.visitConcreteStmt(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.RangeStmt; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitAbstractExpr(n.Key, RangeStmt_Key, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractExpr(n.Value, RangeStmt_Value, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractExpr(n.X, RangeStmt_X, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitBlockStmt(n.Body, RangeStmt_Body, 0, stack2)
		},
	)
}
