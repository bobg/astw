package astw

import "go/ast"

func (v *Visitor) visitCommentGroup(n *ast.CommentGroup, which Which, index int, stack []StackItem) error {
	return v.visitConcreteNode(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.CommentGroup; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			for i, comment := range n.List {
				err := v.visitComment(comment, CommentGroup_List, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitComment(n *ast.Comment, which Which, index int, stack []StackItem) error {
	return v.visitConcreteNode(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.Comment; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		nil,
	)
}
