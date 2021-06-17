package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) VisitBadDecl(n *ast.BadDecl, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.BadDecl; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in BadDecl (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in BadDecl (post)")
			}
		}()
	}

	return
}

func (v *Visitor) VisitGenDecl(n *ast.GenDecl, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.GenDecl; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in GenDecl (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in GenDecl (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitCommentGroup(n.Doc, GenDecl_Doc, 0, stack2)
	if err != nil {
		return err
	}

	for i, spec := range n.Specs {
		err = v.VisitSpec(spec, GenDecl_Specs, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitFuncDecl(n *ast.FuncDecl, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.FuncDecl; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in FuncDecl (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in FuncDecl (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitCommentGroup(n.Doc, FuncDecl_Doc, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitFieldList(n.Recv, FuncDecl_Recv, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitIdent(n.Name, FuncDecl_Name, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitFuncType(n.Type, FuncDecl_Type, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitBlockStmt(n.Body, FuncDecl_Body, 0, stack2)

	return
}
