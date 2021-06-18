package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) visitCommentGroup(n *ast.CommentGroup, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CommentGroup; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in CommentGroup (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in CommentGroup (post)")
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, comment := range n.List {
		err = v.visitComment(comment, CommentGroup_List, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) visitComment(n *ast.Comment, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Comment; f != nil {
		err = f(n, which, index, stack, true, nil)
		if err != nil {
			return errors.Wrap(err, "in Comment (pre)")
		}
		defer func() {
			err = f(n, which, index, stack, false, err)
			err = errors.Wrap(err, "in Comment (post)")
		}()
	}

	return
}
