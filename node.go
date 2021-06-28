package astw

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
)

func (v *Visitor) visitAbstractNode(n ast.Node, which Which, index int, stack []StackItem) (err error) {
	switch n := n.(type) {
	case ast.Expr:
		err = v.visitAbstractExpr(n, which, index, stack)
	case ast.Stmt:
		err = v.visitAbstractStmt(n, which, index, stack)
	case ast.Decl:
		err = v.visitAbstractDecl(n, which, index, stack)
	case ast.Spec:
		err = v.visitAbstractSpec(n, which, index, stack)

	case *ast.File:
		err = v.visitFile(n, which, index, stack)

	default:
		return fmt.Errorf("unknown node type %T", n)
	}

	return
}

func (v *Visitor) visitAbstractExpr(n ast.Expr, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
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

func (v *Visitor) visitAbstractStmt(n ast.Stmt, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
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

func (v *Visitor) visitAbstractDecl(n ast.Decl, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
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

func (v *Visitor) visitAbstractSpec(n ast.Spec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
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

func (v *Visitor) visitConcreteNode(n ast.Node, which Which, index int, stack []StackItem, cbfunc func() func(bool, error) error, childfunc func([]StackItem) error) error {
	return v.visitConcreteNodeHelper(n, which, index, stack, cbfunc(), childfunc)
}

// ErrSkip is an error that a pre-visit callback can return to cause Walk to skip its children.
// The post-visit of the same callback is also skipped.
// Unlike other errors, it is not propagated up the call stack
// (i.e., the post-visit of the parent callback will receive a value of nil for its err argument).
var ErrSkip = errors.New("skip")

func (v *Visitor) visitConcreteNodeHelper(n ast.Node, which Which, index int, stack []StackItem, cb func(bool, error) error, childfunc func([]StackItem) error) error {
	// The following line is brought to you by the dreaded typed nil
	// (`if n == nil` is not good enough).
	if reflect.ValueOf(n).IsNil() {
		return nil
	}
	if cb == nil {
		if f := v.Node; f != nil {
			cb = func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
		}
	}
	if cb != nil {
		err := cb(true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return err
		}
	}
	var err error
	if childfunc != nil {
		err = childfunc(append(stack, StackItem{N: n, W: which, I: index}))
	}
	if cb != nil {
		return cb(false, err)
	}
	return err
}

func (v *Visitor) visitConcreteDecl(n ast.Decl, which Which, index int, stack []StackItem, cbfunc func() func(bool, error) error, childfunc func([]StackItem) error) error {
	if n == nil {
		return nil
	}
	cb := cbfunc()
	if cb == nil {
		if f := v.Decl; f != nil {
			cb = func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
		}
	}
	return v.visitConcreteNodeHelper(n, which, index, stack, cb, childfunc)
}

func (v *Visitor) visitConcreteExpr(n ast.Expr, which Which, index int, stack []StackItem, cbfunc func() func(bool, error) error, childfunc func([]StackItem) error) error {
	if n == nil {
		return nil
	}
	cb := cbfunc()
	if cb == nil {
		if f := v.Expr; f != nil {
			cb = func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
		}
	}
	return v.visitConcreteNodeHelper(n, which, index, stack, cb, childfunc)
}

func (v *Visitor) visitConcreteSpec(n ast.Spec, which Which, index int, stack []StackItem, cbfunc func() func(bool, error) error, childfunc func([]StackItem) error) error {
	if n == nil {
		return nil
	}
	cb := cbfunc()
	if cb == nil {
		if f := v.Spec; f != nil {
			cb = func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
		}
	}
	return v.visitConcreteNodeHelper(n, which, index, stack, cb, childfunc)
}

func (v *Visitor) visitConcreteStmt(n ast.Stmt, which Which, index int, stack []StackItem, cbfunc func() func(bool, error) error, childfunc func([]StackItem) error) error {
	if n == nil {
		return nil
	}
	cb := cbfunc()
	if cb == nil {
		if f := v.Stmt; f != nil {
			cb = func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
		}
	}
	return v.visitConcreteNodeHelper(n, which, index, stack, cb, childfunc)
}
