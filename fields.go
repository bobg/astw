package astw

import "go/ast"

func (v *Visitor) visitFieldList(n *ast.FieldList, which Which, index int, stack []StackItem) error {
	return v.visitConcreteNode(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.FieldList; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			for i, field := range n.List {
				err := v.visitField(field, FieldList_List, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitField(n *ast.Field, which Which, index int, stack []StackItem) error {
	return v.visitConcreteNode(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.Field; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitCommentGroup(n.Doc, Field_Doc, 0, stack2)
			if err != nil {
				return err
			}
			for i, ident := range n.Names {
				err = v.visitIdent(ident, Field_Names, i, stack2)
				if err != nil {
					return err
				}
			}
			err = v.visitAbstractExpr(n.Type, Field_Type, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitBasicLit(n.Tag, Field_Tag, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitCommentGroup(n.Comment, Field_Comment, 0, stack2)
		},
	)
}
