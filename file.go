package astw

import (
	"go/ast"
	"sort"
)

func (v *Visitor) visitPackage(n *ast.Package, which Which, index int, stack []StackItem) (err error) {
	return v.visitConcreteNode(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.Package; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			var filenames []string
			for filename := range n.Files {
				filenames = append(filenames, filename)
			}
			sort.Strings(filenames)

			for i, filename := range filenames {
				v.Filename = filename
				err = v.visitFile(n.Files[filename], Package_Files, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (v *Visitor) visitFile(n *ast.File, which Which, index int, stack []StackItem) (err error) {
	return v.visitConcreteNode(
		n, which, index, stack,
		func() func(bool, error) error {
			if f := v.File; f != nil {
				return func(pre bool, err error) error { return f(n, which, index, stack, pre, err) }
			}
			return nil
		},
		func(stack2 []StackItem) error {
			err := v.visitCommentGroup(n.Doc, File_Doc, 0, stack2)
			if err != nil {
				return err
			}
			err = v.visitIdent(n.Name, File_Name, 0, stack2)
			if err != nil {
				return err
			}
			for i, decl := range n.Decls {
				err = v.visitAbstractDecl(decl, File_Decls, i, stack2)
				if err != nil {
					return err
				}
			}
			for i, importSpec := range n.Imports {
				err = v.visitImportSpec(importSpec, File_Imports, i, stack2)
				if err != nil {
					return err
				}
			}
			for i, ident := range n.Unresolved {
				err = v.visitIdent(ident, File_Unresolved, i, stack2)
				if err != nil {
					return err
				}
			}
			for i, comment := range n.Comments {
				err = v.visitCommentGroup(comment, File_Comments, i, stack2)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}
