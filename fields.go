package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) VisitFieldList(n *ast.FieldList, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.FieldList; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in FieldList (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in FieldList (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, field := range n.List {
		err = v.VisitField(field, FieldList_List, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitField(n *ast.Field, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Field; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Field (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Field (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, ident := range n.Names {
		err = v.VisitIdent(ident, Field_Names, i, stack2)
		if err != nil {
			return err
		}
	}

	err = v.VisitExpr(n.Type, Field_Type, 0, stack2)

	return
}
