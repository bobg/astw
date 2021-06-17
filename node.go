package astw

import (
	"fmt"
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) VisitFile(n *ast.File, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.File; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in File (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in File (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitIdent(n.Name, File_Name, 0, stack2)
	if err != nil {
		return err
	}

	for i, decl := range n.Decls {
		err = v.VisitDecl(decl, File_Decls, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, importSpec := range n.Imports {
		err = v.VisitImportSpec(importSpec, File_Imports, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, ident := range n.Unresolved {
		err = v.VisitIdent(ident, File_Unresolved, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitNode(n ast.Node, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Node; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Node (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Node (post)")
			}
		}()
	}

	switch n := n.(type) {
	case ast.Expr:
		err = v.VisitExpr(n, which, index, stack)
	case ast.Stmt:
		err = v.VisitStmt(n, which, index, stack)
	case ast.Decl:
		err = v.VisitDecl(n, which, index, stack)
	case ast.Spec:
		err = v.VisitSpec(n, which, index, stack)

	case *ast.File:
		err = v.VisitFile(n, which, index, stack)

	default:
		return fmt.Errorf("unknown node type %T", n)
	}

	return
}

func (v *Visitor) VisitExpr(n ast.Expr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Expr; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Expr (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Expr (post)")
			}
		}()
	}

	switch n := n.(type) {
	case *ast.BadExpr:
		err = v.VisitBadExpr(n, which, index, stack)
	case *ast.Ident:
		err = v.VisitIdent(n, which, index, stack)
	case *ast.Ellipsis:
		err = v.VisitEllipsis(n, which, index, stack)
	case *ast.BasicLit:
		err = v.VisitBasicLit(n, which, index, stack)
	case *ast.FuncLit:
		err = v.VisitFuncLit(n, which, index, stack)
	case *ast.CompositeLit:
		err = v.VisitCompositeLit(n, which, index, stack)
	case *ast.ParenExpr:
		err = v.VisitParenExpr(n, which, index, stack)
	case *ast.SelectorExpr:
		err = v.VisitSelectorExpr(n, which, index, stack)
	case *ast.IndexExpr:
		err = v.VisitIndexExpr(n, which, index, stack)
	case *ast.SliceExpr:
		err = v.VisitSliceExpr(n, which, index, stack)
	case *ast.TypeAssertExpr:
		err = v.VisitTypeAssertExpr(n, which, index, stack)
	case *ast.CallExpr:
		err = v.VisitCallExpr(n, which, index, stack)
	case *ast.StarExpr:
		err = v.VisitStarExpr(n, which, index, stack)
	case *ast.UnaryExpr:
		err = v.VisitUnaryExpr(n, which, index, stack)
	case *ast.BinaryExpr:
		err = v.VisitBinaryExpr(n, which, index, stack)
	case *ast.KeyValueExpr:
		err = v.VisitKeyValueExpr(n, which, index, stack)
	case *ast.ArrayType:
		err = v.VisitArrayType(n, which, index, stack)
	case *ast.StructType:
		err = v.VisitStructType(n, which, index, stack)
	case *ast.FuncType:
		err = v.VisitFuncType(n, which, index, stack)
	case *ast.InterfaceType:
		err = v.VisitInterfaceType(n, which, index, stack)
	case *ast.MapType:
		err = v.VisitMapType(n, which, index, stack)
	case *ast.ChanType:
		err = v.VisitChanType(n, which, index, stack)
	default:
		return fmt.Errorf("unknown expr type %T", n)
	}

	return
}

func (v *Visitor) VisitStmt(n ast.Stmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Stmt; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Stmt (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Stmt (post)")
			}
		}()
	}

	switch n := n.(type) {
	case *ast.BadStmt:
		err = v.VisitBadStmt(n, which, index, stack)
	case *ast.DeclStmt:
		err = v.VisitDeclStmt(n, which, index, stack)
	case *ast.EmptyStmt:
		err = v.VisitEmptyStmt(n, which, index, stack)
	case *ast.LabeledStmt:
		err = v.VisitLabeledStmt(n, which, index, stack)
	case *ast.ExprStmt:
		err = v.VisitExprStmt(n, which, index, stack)
	case *ast.SendStmt:
		err = v.VisitSendStmt(n, which, index, stack)
	case *ast.IncDecStmt:
		err = v.VisitIncDecStmt(n, which, index, stack)
	case *ast.AssignStmt:
		err = v.VisitAssignStmt(n, which, index, stack)
	case *ast.GoStmt:
		err = v.VisitGoStmt(n, which, index, stack)
	case *ast.DeferStmt:
		err = v.VisitDeferStmt(n, which, index, stack)
	case *ast.ReturnStmt:
		err = v.VisitReturnStmt(n, which, index, stack)
	case *ast.BranchStmt:
		err = v.VisitBranchStmt(n, which, index, stack)
	case *ast.BlockStmt:
		err = v.VisitBlockStmt(n, which, index, stack)
	case *ast.IfStmt:
		err = v.VisitIfStmt(n, which, index, stack)
	case *ast.CaseClause:
		err = v.VisitCaseClause(n, which, index, stack)
	case *ast.SwitchStmt:
		err = v.VisitSwitchStmt(n, which, index, stack)
	case *ast.TypeSwitchStmt:
		err = v.VisitTypeSwitchStmt(n, which, index, stack)
	case *ast.CommClause:
		err = v.VisitCommClause(n, which, index, stack)
	case *ast.SelectStmt:
		err = v.VisitSelectStmt(n, which, index, stack)
	case *ast.ForStmt:
		err = v.VisitForStmt(n, which, index, stack)
	case *ast.RangeStmt:
		err = v.VisitRangeStmt(n, which, index, stack)
	default:
		return fmt.Errorf("unknown stmt type %T", n)
	}

	return
}

func (v *Visitor) VisitDecl(n ast.Decl, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Decl; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Decl (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Decl (post)")
			}
		}()
	}

	switch n := n.(type) {
	case *ast.BadDecl:
		err = v.VisitBadDecl(n, which, index, stack)
	case *ast.GenDecl:
		err = v.VisitGenDecl(n, which, index, stack)
	case *ast.FuncDecl:
		err = v.VisitFuncDecl(n, which, index, stack)
	default:
		return fmt.Errorf("unknown decl type %T", n)
	}

	return
}

func (v *Visitor) VisitSpec(n ast.Spec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Spec; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Spec (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Spec (post)")
			}
		}()
	}

	switch n := n.(type) {
	case *ast.ImportSpec:
		err = v.VisitImportSpec(n, which, index, stack)
	case *ast.ValueSpec:
		err = v.VisitValueSpec(n, which, index, stack)
	case *ast.TypeSpec:
		err = v.VisitTypeSpec(n, which, index, stack)
	default:
		return fmt.Errorf("unknown spec type %T", n)
	}

	return
}
