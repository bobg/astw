package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) visitImportSpec(n *ast.ImportSpec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ImportSpec; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in ImportSpec (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in ImportSpec (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitCommentGroup(n.Doc, ImportSpec_Doc, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitIdent(n.Name, ImportSpec_Name, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitBasicLit(n.Path, ImportSpec_Path, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitCommentGroup(n.Comment, ImportSpec_Comment, 0, stack2)

	return
}

func (v *Visitor) visitValueSpec(n *ast.ValueSpec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.ValueSpec; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in ValueSpec (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in ValueSpec (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitCommentGroup(n.Doc, ValueSpec_Doc, 0, stack2)
	if err != nil {
		return err
	}

	for i, ident := range n.Names {
		err = v.visitIdent(ident, ValueSpec_Names, i, stack2)
		if err != nil {
			return err
		}
	}

	err = v.visitExpr(n.Type, ValueSpec_Type, 0, stack2)
	if err != nil {
		return err
	}

	for i, expr := range n.Values {
		err = v.visitExpr(expr, ValueSpec_Values, i, stack2)
		if err != nil {
			return err
		}
	}

	err = v.visitCommentGroup(n.Comment, ValueSpec_Comment, 0, stack2)

	return
}

func (v *Visitor) visitTypeSpec(n *ast.TypeSpec, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.TypeSpec; f != nil {
		err = f(n, which, index, stack, true, nil)
		if errors.Is(err, ErrSkip) {
			return nil
		}
		if err != nil {
			return errors.Wrap(err, "in TypeSpec (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in TypeSpec (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	err = v.visitCommentGroup(n.Doc, TypeSpec_Doc, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitIdent(n.Name, TypeSpec_Name, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitExpr(n.Type, TypeSpec_Type, 0, stack2)
	if err != nil {
		return err
	}

	err = v.visitCommentGroup(n.Comment, TypeSpec_Comment, 0, stack2)

	return
}
