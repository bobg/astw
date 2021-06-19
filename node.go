package astw

import (
	"fmt"
	"go/ast"
	"sort"

	"github.com/pkg/errors"
)

func (v *Visitor) visitPackage(n *ast.Package, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Package; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in Package (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Package (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	var filenames []string
	for filename := range n.Files {
		filenames = append(filenames, filename)
	}
	sort.Strings(filenames)

	for i, filename := range filenames {
		v.Filename = filename
		err = v.visitFile(n.Files[filename], Package_Files, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitFile(n *ast.File, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.File; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in File (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in File (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitCommentGroup(n.Doc, File_Doc, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitIdent(n.Name, File_Name, 0, stack2)
	if err != nil {
		return err
	}

	for i, decl := range n.Decls {
		err = v.visitDecl(decl, File_Decls, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, importSpec := range n.Imports {
		err = v.visitImportSpec(importSpec, File_Imports, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, ident := range n.Unresolved {
		err = v.visitIdent(ident, File_Unresolved, i, stack2)
		if err != nil {
			return err
		}
	}

	for i, comment := range n.Comments {
		err = v.visitCommentGroup(comment, File_Comments, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitNode(n ast.Node, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Node; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in Node (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Node (post)")
		}()
	}

	switch n := n.(type) {
	case ast.Expr:
		err = v.visitExpr(n, which, index, stack)
	case ast.Stmt:
		err = v.visitStmt(n, which, index, stack)
	case ast.Decl:
		err = v.visitDecl(n, which, index, stack)
	case ast.Spec:
		err = v.visitSpec(n, which, index, stack)

	case *ast.File:
		err = v.visitFile(n, which, index, stack)

	default:
		return fmt.Errorf("unknown node type %T", n)
	}

	return
}

func (v *Visitor) visitExpr(n ast.Expr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Expr; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in Expr (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Expr (post)")
		}()
	}

	switch n := n.(type) {
	case *ast.BadExpr:
		err = v.visitBadExpr(n, which, index, stack)
	case *ast.Ident:
		err = v.visitIdent(n, which, index, stack)
	case *ast.Ellipsis:
		err = v.visitEllipsis(n, which, index, stack)
	case *ast.BasicLit:
		err = v.visitBasicLit(n, which, index, stack)
	case *ast.FuncLit:
		err = v.visitFuncLit(n, which, index, stack)
	case *ast.CompositeLit:
		err = v.visitCompositeLit(n, which, index, stack)
	case *ast.ParenExpr:
		err = v.visitParenExpr(n, which, index, stack)
	case *ast.SelectorExpr:
		err = v.visitSelectorExpr(n, which, index, stack)
	case *ast.IndexExpr:
		err = v.visitIndexExpr(n, which, index, stack)
	case *ast.SliceExpr:
		err = v.visitSliceExpr(n, which, index, stack)
	case *ast.TypeAssertExpr:
		err = v.visitTypeAssertExpr(n, which, index, stack)
	case *ast.CallExpr:
		err = v.visitCallExpr(n, which, index, stack)
	case *ast.StarExpr:
		err = v.visitStarExpr(n, which, index, stack)
	case *ast.UnaryExpr:
		err = v.visitUnaryExpr(n, which, index, stack)
	case *ast.BinaryExpr:
		err = v.visitBinaryExpr(n, which, index, stack)
	case *ast.KeyValueExpr:
		err = v.visitKeyValueExpr(n, which, index, stack)
	case *ast.ArrayType:
		err = v.visitArrayType(n, which, index, stack)
	case *ast.StructType:
		err = v.visitStructType(n, which, index, stack)
	case *ast.FuncType:
		err = v.visitFuncType(n, which, index, stack)
	case *ast.InterfaceType:
		err = v.visitInterfaceType(n, which, index, stack)
	case *ast.MapType:
		err = v.visitMapType(n, which, index, stack)
	case *ast.ChanType:
		err = v.visitChanType(n, which, index, stack)
	default:
		return fmt.Errorf("unknown expr type %T", n)
	}

	return
}

func (v *Visitor) visitStmt(n ast.Stmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Stmt; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in Stmt (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Stmt (post)")
		}()
	}

	switch n := n.(type) {
	case *ast.BadStmt:
		err = v.visitBadStmt(n, which, index, stack)
	case *ast.DeclStmt:
		err = v.visitDeclStmt(n, which, index, stack)
	case *ast.EmptyStmt:
		err = v.visitEmptyStmt(n, which, index, stack)
	case *ast.LabeledStmt:
		err = v.visitLabeledStmt(n, which, index, stack)
	case *ast.ExprStmt:
		err = v.visitExprStmt(n, which, index, stack)
	case *ast.SendStmt:
		err = v.visitSendStmt(n, which, index, stack)
	case *ast.IncDecStmt:
		err = v.visitIncDecStmt(n, which, index, stack)
	case *ast.AssignStmt:
		err = v.visitAssignStmt(n, which, index, stack)
	case *ast.GoStmt:
		err = v.visitGoStmt(n, which, index, stack)
	case *ast.DeferStmt:
		err = v.visitDeferStmt(n, which, index, stack)
	case *ast.ReturnStmt:
		err = v.visitReturnStmt(n, which, index, stack)
	case *ast.BranchStmt:
		err = v.visitBranchStmt(n, which, index, stack)
	case *ast.BlockStmt:
		err = v.visitBlockStmt(n, which, index, stack)
	case *ast.IfStmt:
		err = v.visitIfStmt(n, which, index, stack)
	case *ast.CaseClause:
		err = v.visitCaseClause(n, which, index, stack)
	case *ast.SwitchStmt:
		err = v.visitSwitchStmt(n, which, index, stack)
	case *ast.TypeSwitchStmt:
		err = v.visitTypeSwitchStmt(n, which, index, stack)
	case *ast.CommClause:
		err = v.visitCommClause(n, which, index, stack)
	case *ast.SelectStmt:
		err = v.visitSelectStmt(n, which, index, stack)
	case *ast.ForStmt:
		err = v.visitForStmt(n, which, index, stack)
	case *ast.RangeStmt:
		err = v.visitRangeStmt(n, which, index, stack)
	default:
		return fmt.Errorf("unknown stmt type %T", n)
	}

	return
}

func (v *Visitor) visitDecl(n ast.Decl, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Decl; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in Decl (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Decl (post)")
		}()
	}

	switch n := n.(type) {
	case *ast.BadDecl:
		err = v.visitBadDecl(n, which, index, stack)
	case *ast.GenDecl:
		err = v.visitGenDecl(n, which, index, stack)
	case *ast.FuncDecl:
		err = v.visitFuncDecl(n, which, index, stack)
	default:
		return fmt.Errorf("unknown decl type %T", n)
	}

	return
}

func (v *Visitor) visitSpec(n ast.Spec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Spec; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in Spec (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Spec (post)")
		}()
	}

	switch n := n.(type) {
	case *ast.ImportSpec:
		err = v.visitImportSpec(n, which, index, stack)
	case *ast.ValueSpec:
		err = v.visitValueSpec(n, which, index, stack)
	case *ast.TypeSpec:
		err = v.visitTypeSpec(n, which, index, stack)
	default:
		return fmt.Errorf("unknown spec type %T", n)
	}

	return
}
