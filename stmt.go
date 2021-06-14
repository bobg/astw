package goast

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) VisitBadStmt(n *ast.BadStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BadStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in BadStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in BadStmt (post)")
			}
		}()
	}

	return
}

func (v *Visitor) VisitDeclStmt(n *ast.DeclStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.DeclStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in DeclStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in DeclStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitDecl(n.Decl, DeclStmt_Decl, 0, stack2)

	return
}

func (v *Visitor) VisitEmptyStmt(n *ast.EmptyStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.EmptyStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in EmptyStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in EmptyStmt (post)")
			}
		}()
	}

	return
}

func (v *Visitor) VisitLabeledStmt(n *ast.LabeledStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.LabeledStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in LabeledStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in LabeledStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitIdent(n.Label, LabeledStmt_Label, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitStmt(n.Stmt, LabeledStmt_Stmt, 0, stack2)

	return
}

func (v *Visitor) VisitExprStmt(n *ast.ExprStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ExprStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in ExprStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in ExprStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, ExprStmt_Expr, 0, stack2)

	return
}

func (v *Visitor) VisitSendStmt(n *ast.SendStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SendStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in SendStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in SendStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Chan, SendStmt_Chan, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Value, SendStmt_Value, 0, stack2)

	return
}

func (v *Visitor) VisitIncDecStmt(n *ast.IncDecStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.IncDecStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in IncDecStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in IncDecStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.X, IncDecStmt_X, 0, stack2)

	return
}

func (v *Visitor) VisitAssignStmt(n *ast.AssignStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.AssignStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in AssignStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in AssignStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, expr := range n.Lhs {
		err = v.VisitExpr(expr, AssignStmt_Lhs, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, expr := range n.Rhs {
		err = v.VisitExpr(expr, AssignStmt_Rhs, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitGoStmt(n *ast.GoStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.GoStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in GoStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in GoStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitCallExpr(n.Call, GoStmt_Call, 0, stack2)

	return
}

func (v *Visitor) VisitDeferStmt(n *ast.DeferStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.DeferStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in DeferStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in DeferStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitCallExpr(n.Call, DeferStmt_Call, 0, stack2)

	return
}

func (v *Visitor) VisitReturnStmt(n *ast.ReturnStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ReturnStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in ReturnStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in ReturnStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, expr := range n.Results {
		err = v.VisitExpr(expr, ReturnStmt_Results, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitBranchStmt(n *ast.BranchStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BranchStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in BranchStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in BranchStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitIdent(n.Label, BranchStmt_Label, 0, stack2)

	return
}

func (v *Visitor) VisitBlockStmt(n *ast.BlockStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BlockStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in BlockStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in BlockStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, stmt := range n.List {
		err = v.VisitStmt(stmt, BlockStmt_List, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitIfStmt(n *ast.IfStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.IfStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in IfStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in IfStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitStmt(n.Init, IfStmt_Init, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Cond, IfStmt_Cond, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitBlockStmt(n.Body, IfStmt_Body, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitStmt(n.Else, IfStmt_Else, 0, stack2)

	return
}

func (v *Visitor) VisitCaseClause(n *ast.CaseClause, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CaseClause; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in CaseClause (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in CaseClause (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, expr := range n.List {
		err = v.VisitExpr(expr, CaseClause_List, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, stmt := range n.Body {
		err = v.VisitStmt(stmt, CaseClause_Body, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitSwitchStmt(n *ast.SwitchStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SwitchStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in SwitchStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in SwitchStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitStmt(n.Init, SwitchStmt_Init, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Tag, SwitchStmt_Tag, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitBlockStmt(n.Body, SwitchStmt_Body, 0, stack2)
	if err != nil {
		return err
	}

	return
}

func (v *Visitor) VisitTypeSwitchStmt(n *ast.TypeSwitchStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.TypeSwitchStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in TypeSwitchStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in TypeSwitchStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitStmt(n.Init, TypeSwitchStmt_Init, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitStmt(n.Assign, TypeSwitchStmt_Assign, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitBlockStmt(n.Body, TypeSwitchStmt_Body, 0, stack2)

	return
}

func (v *Visitor) VisitCommClause(n *ast.CommClause, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CommClause; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in CommClause (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in CommClause (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitStmt(n.Comm, CommClause_Comm, 0, stack2)
	if err != nil {
		return err
	}

	for i, stmt := range n.Body {
		err = v.VisitStmt(stmt, CommClause_Body, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitSelectStmt(n *ast.SelectStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SelectStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in SelectStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in SelectStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitBlockStmt(n.Body, SelectStmt_Body, 0, stack2)

	return
}

func (v *Visitor) VisitForStmt(n *ast.ForStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ForStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in ForStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in ForStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitStmt(n.Init, ForStmt_Init, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Cond, ForStmt_Cond, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitStmt(n.Post, ForStmt_Post, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitBlockStmt(n.Body, ForStmt_Body, 0, stack2)

	return
}

func (v *Visitor) VisitRangeStmt(n *ast.RangeStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.RangeStmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in RangeStmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in RangeStmt (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitExpr(n.Key, RangeStmt_Key, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Value, RangeStmt_Value, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.X, RangeStmt_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitBlockStmt(n.Body, RangeStmt_Body, 0, stack2)

	return
}
