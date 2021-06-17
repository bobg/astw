package astw

import (
	"go/ast"

	"github.com/pkg/errors"
)

func (v *Visitor) VisitCommentGroup(n *ast.CommentGroup, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.CommentGroup; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in CommentGroup (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in CommentGroup (post)")
			}
		}()
	}

	stack2 := append(stack, StackItem{N: n, W: which, I: index})

	for i, comment := range n.List {
		err = v.VisitComment(comment, CommentGroup_List, i, stack2)
		if err != nil {
			return err
		}
	}

	return
}

func (v *Visitor) VisitComment(n *ast.Comment, which Which, index int, stack []StackItem) (err error) {
	if n == nil {
		return nil
	}

	if f := v.Comment; f != nil {
		err = f(n, which, index, stack, true)
		if err != nil {
			return errors.Wrap(err, "in Comment (pre)")
		}
		defer func() {
			if err == nil {
				err = f(n, which, index, stack, false)
				err = errors.Wrap(err, "in Comment (post)")
			}
		}()
	}

	return
}
