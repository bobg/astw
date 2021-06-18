package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) visitFieldList(n *ast.FieldList, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.FieldList; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in FieldList (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in FieldList (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, field := range n.List {
		err = v.visitField(field, FieldList_List, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitField(n *ast.Field, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Field; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in Field (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Field (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitCommentGroup(n.Doc, Field_Doc, 0, stack2)
	if err != nil {
		return err
	}

	for i, ident := range n.Names {
		err = v.visitIdent(ident, Field_Names, i, stack2)
		if err != nil {
			return err
		}
	}

	err = v.visitExpr(n.Type, Field_Type, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitBasicLit(n.Tag, Field_Tag, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitCommentGroup(n.Comment, Field_Comment, 0, stack2)

	return
}
