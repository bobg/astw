package astw

import "go/ast"

func (v *Visitor) visitBadDecl(n *ast.BadDecl, which Which, index int, stack []StackItem) error {
	return v.visitConcreteDecl(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.BadDecl; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		nil,
	)
}

func (v *Visitor) visitGenDecl(n *ast.GenDecl, which Which, index int, stack []StackItem) error {
	return v.visitConcreteDecl(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.GenDecl; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitCommentGroup(n.Doc, GenDecl_Doc, 0, stack2)
			if err != nil {
				return err
			}
			for i, spec := range n.Specs {
				err = v.visitAbstractSpec(spec, GenDecl_Specs, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitFuncDecl(n *ast.FuncDecl, which Which, index int, stack []StackItem) error {
	return v.visitConcreteDecl(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.FuncDecl; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitCommentGroup(n.Doc, FuncDecl_Doc, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitFieldList(n.Recv, FuncDecl_Recv, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitIdent(n.Name, FuncDecl_Name, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitFuncType(n.Type, FuncDecl_Type, 0, stack2)
			if err != nil {
				return err
			}
			return v.visitBlockStmt(n.Body, FuncDecl_Body, 0, stack2)
		},
	)
}
