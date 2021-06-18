package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) visitBadStmt(n *ast.BadStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BadStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in BadStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in BadStmt (post)")
		}()
	}

	return
}

func (v *Visitor) visitDeclStmt(n *ast.DeclStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.DeclStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in DeclStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in DeclStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitDecl(n.Decl, DeclStmt_Decl, 0, stack2)

	return
}

func (v *Visitor) visitEmptyStmt(n *ast.EmptyStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.EmptyStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in EmptyStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in EmptyStmt (post)")
		}()
	}

	return
}

func (v *Visitor) visitLabeledStmt(n *ast.LabeledStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.LabeledStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in LabeledStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in LabeledStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitIdent(n.Label, LabeledStmt_Label, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitStmt(n.Stmt, LabeledStmt_Stmt, 0, stack2)

	return
}

func (v *Visitor) visitExprStmt(n *ast.ExprStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ExprStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in ExprStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in ExprStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, ExprStmt_Expr, 0, stack2)

	return
}

func (v *Visitor) visitSendStmt(n *ast.SendStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SendStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in SendStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in SendStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Chan, SendStmt_Chan, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Value, SendStmt_Value, 0, stack2)

	return
}

func (v *Visitor) visitIncDecStmt(n *ast.IncDecStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.IncDecStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in IncDecStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in IncDecStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.X, IncDecStmt_X, 0, stack2)

	return
}

func (v *Visitor) visitAssignStmt(n *ast.AssignStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.AssignStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in AssignStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in AssignStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, expr := range n.Lhs {
		err = v.visitExpr(expr, AssignStmt_Lhs, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, expr := range n.Rhs {
		err = v.visitExpr(expr, AssignStmt_Rhs, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitGoStmt(n *ast.GoStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.GoStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in GoStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in GoStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitCallExpr(n.Call, GoStmt_Call, 0, stack2)

	return
}

func (v *Visitor) visitDeferStmt(n *ast.DeferStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.DeferStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in DeferStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in DeferStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitCallExpr(n.Call, DeferStmt_Call, 0, stack2)

	return
}

func (v *Visitor) visitReturnStmt(n *ast.ReturnStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ReturnStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in ReturnStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in ReturnStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, expr := range n.Results {
		err = v.visitExpr(expr, ReturnStmt_Results, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitBranchStmt(n *ast.BranchStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BranchStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in BranchStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in BranchStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitIdent(n.Label, BranchStmt_Label, 0, stack2)

	return
}

func (v *Visitor) visitBlockStmt(n *ast.BlockStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BlockStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in BlockStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in BlockStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, stmt := range n.List {
		err = v.visitStmt(stmt, BlockStmt_List, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitIfStmt(n *ast.IfStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.IfStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in IfStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in IfStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitStmt(n.Init, IfStmt_Init, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Cond, IfStmt_Cond, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitBlockStmt(n.Body, IfStmt_Body, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitStmt(n.Else, IfStmt_Else, 0, stack2)

	return
}

func (v *Visitor) visitCaseClause(n *ast.CaseClause, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CaseClause; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in CaseClause (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in CaseClause (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, expr := range n.List {
		err = v.visitExpr(expr, CaseClause_List, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, stmt := range n.Body {
		err = v.visitStmt(stmt, CaseClause_Body, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitSwitchStmt(n *ast.SwitchStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SwitchStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in SwitchStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in SwitchStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitStmt(n.Init, SwitchStmt_Init, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Tag, SwitchStmt_Tag, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitBlockStmt(n.Body, SwitchStmt_Body, 0, stack2)
	if err != nil {
		return err
	}

	return
}

func (v *Visitor) visitTypeSwitchStmt(n *ast.TypeSwitchStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.TypeSwitchStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in TypeSwitchStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in TypeSwitchStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitStmt(n.Init, TypeSwitchStmt_Init, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitStmt(n.Assign, TypeSwitchStmt_Assign, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitBlockStmt(n.Body, TypeSwitchStmt_Body, 0, stack2)

	return
}

func (v *Visitor) visitCommClause(n *ast.CommClause, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CommClause; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in CommClause (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in CommClause (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitStmt(n.Comm, CommClause_Comm, 0, stack2)
	if err != nil {
		return err
	}

	for i, stmt := range n.Body {
		err = v.visitStmt(stmt, CommClause_Body, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitSelectStmt(n *ast.SelectStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.SelectStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in SelectStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in SelectStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitBlockStmt(n.Body, SelectStmt_Body, 0, stack2)

	return
}

func (v *Visitor) visitForStmt(n *ast.ForStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ForStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in ForStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in ForStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitStmt(n.Init, ForStmt_Init, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Cond, ForStmt_Cond, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitStmt(n.Post, ForStmt_Post, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitBlockStmt(n.Body, ForStmt_Body, 0, stack2)

	return
}

func (v *Visitor) visitRangeStmt(n *ast.RangeStmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.RangeStmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in RangeStmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in RangeStmt (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitExpr(n.Key, RangeStmt_Key, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Value, RangeStmt_Value, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.X, RangeStmt_X, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitBlockStmt(n.Body, RangeStmt_Body, 0, stack2)

	return
}
