package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) VisitImportSpec(n *ast.ImportSpec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ImportSpec; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in ImportSpec (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in ImportSpec (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitCommentGroup(n.Doc, ImportSpec_Doc, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitIdent(n.Name, ImportSpec_Name, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitBasicLit(n.Path, ImportSpec_Path, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitCommentGroup(n.Comment, ImportSpec_Comment, 0, stack2)

	return
}

func (v *Visitor) VisitValueSpec(n *ast.ValueSpec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ValueSpec; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in ValueSpec (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in ValueSpec (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitCommentGroup(n.Doc, ValueSpec_Doc, 0, stack2)
	if err != nil {
		return err
	}

	for i, ident := range n.Names {
		err = v.VisitIdent(ident, ValueSpec_Names, i, stack2)
		if err != nil {
			return err
		}
	}

	err = v.VisitExpr(n.Type, ValueSpec_Type, 0, stack2)
	if err != nil {
		return err
	}

	for i, expr := range n.Values {
		err = v.VisitExpr(expr, ValueSpec_Values, i, stack2)
		if err != nil {
			return err
		}
	}

	err = v.VisitCommentGroup(n.Comment, ValueSpec_Comment, 0, stack2)

	return
}

func (v *Visitor) VisitTypeSpec(n *ast.TypeSpec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.TypeSpec; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in TypeSpec (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in TypeSpec (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.VisitCommentGroup(n.Doc, TypeSpec_Doc, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitIdent(n.Name, TypeSpec_Name, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitExpr(n.Type, TypeSpec_Type, 0, stack2)
	if err != nil {
		return err
	}

	err = v.VisitCommentGroup(n.Comment, TypeSpec_Comment, 0, stack2)

	return
}
