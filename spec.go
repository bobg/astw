package astw

import "go/ast"

func (v *Visitor) visitImportSpec(n *ast.ImportSpec, which Which, index int, stack []StackItem) error {
	return v.visitConcreteSpec(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.ImportSpec; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitCommentGroup(n.Doc, ImportSpec_Doc, 0, stack2)
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
			return v.visitCommentGroup(n.Comment, ImportSpec_Comment, 0, stack2)
		},
	)
}

func (v *Visitor) visitValueSpec(n *ast.ValueSpec, which Which, index int, stack []StackItem) error {
	return v.visitConcreteSpec(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.ValueSpec; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitCommentGroup(n.Doc, ValueSpec_Doc, 0, stack2)
			if err != nil {
				return err
			}
			for i, ident := range n.Names {
				err = v.visitIdent(ident, ValueSpec_Names, i, stack2)
				if err != nil {
					return err
				}
			}
			err = v.visitAbstractExpr(n.Type, ValueSpec_Type, 0, stack2)
			if err != nil {
				return err
			}
			for i, expr := range n.Values {
				err = v.visitAbstractExpr(expr, ValueSpec_Values, i, stack2)
				if err != nil {
					return err
				}
			}
			return v.visitCommentGroup(n.Comment, ValueSpec_Comment, 0, stack2)
		},
	)
}

func (v *Visitor) visitTypeSpec(n *ast.TypeSpec, which Which, index int, stack []StackItem) error {
	return v.visitConcreteSpec(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.TypeSpec; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitCommentGroup(n.Doc, TypeSpec_Doc, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitIdent(n.Name, TypeSpec_Name, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitAbstractExpr(n.Type, TypeSpec_Type, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitCommentGroup(n.Comment, TypeSpec_Comment, 0, stack2)
		},
	)
}
