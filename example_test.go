package astw_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"

	"github.com/bobg/astw"
)

func Example() {
	// In this example,
	// a Go file is parsed from the standard input.
	// Its functions are then scanned for parameters that are used only for their method calls.
	// These are candidates for being declared with (possibly smaller) interface types,
	// to improve composability and simplify testing.
	fset := token.NewFileSet()
	tree, err := parser.ParseFile(fset, "", os.Stdin, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Callbacks in the *astw.Visitor below are methods on this object,
	// which allows extra state
	// (namely, the fset object)
	// to be carried through the syntax-tree walk.
	mv := &myVisitor{fset: fset}

	// This visitor inspects both top-level (named) function declarations
	// and inline (anonymous) function literals.
	v := &astw.Visitor{
		FuncDecl: mv.onFuncDecl,
		FuncLit:  mv.onFuncLit,
	}

	err = astw.Walk(v, tree)
	if err != nil {
		log.Fatal(err)
	}
}

type myVisitor struct {
	fset *token.FileSet
}

func (mv *myVisitor) onFuncDecl(funcDecl *ast.FuncDecl, which astw.Which, index int, stack []astw.StackItem, pre bool, err error) error {
	if err != nil {
		return err
	}
	if !pre {
		return nil
	}
	return mv.walkFunc(funcDecl.Type, funcDecl.Body)
}

func (mv *myVisitor) onFuncLit(funcLit *ast.FuncLit, which astw.Which, index int, stack []astw.StackItem, pre bool, err error) error {
	if err != nil {
		return err
	}
	if !pre {
		return nil
	}
	return mv.walkFunc(funcLit.Type, funcLit.Body)
}

func (mv *myVisitor) walkFunc(typ *ast.FuncType, body *ast.BlockStmt) error {
	for _, field := range typ.Params.List {
		for _, ident := range field.Names {
			// Start a separate, new walk of the function body
			// looking for uses of this parameter.

			var (
				methodsOnly = true
				methodNames = make(map[string]bool)
			)

			iv := &identVisitor{
				paramIdent:  ident,
				methodsOnly: &methodsOnly,
				methodNames: methodNames,
			}

			v := &astw.Visitor{Ident: iv.onIdent}
			err := astw.Walk(v, body)
			if err != nil {
				return err
			}

			if methodsOnly && len(methodNames) > 0 {
				var sorted []string
				for method := range methodNames {
					sorted = append(sorted, method)
				}
				sort.Strings(sorted)
				fmt.Printf("%s: Parameter %s is used only for these methods: %v\n", mv.fset.PositionFor(typ.Pos(), false), ident.Name, sorted)
			}
		}
	}
	return nil
}

type identVisitor struct {
	paramIdent  *ast.Ident
	methodsOnly *bool
	methodNames map[string]bool
}

func (iv *identVisitor) onIdent(ident *ast.Ident, which astw.Which, index int, stack []astw.StackItem, pre bool, err error) error {
	if err != nil {
		return err
	}
	if !pre {
		return nil
	}
	if !*iv.methodsOnly {
		// If we've already determined that the parameter is used for other than it's methods, return early.
		return nil
	}
	if ident.Name != iv.paramIdent.Name {
		// Identifier must match the parameter we're searching for.
		return nil
	}
	if ident.Obj.Pos() != iv.paramIdent.Obj.Pos() {
		// Check not only that ident and paramIdent have the same name,
		// but that they refer to the same thing.
		return nil
	}

	// The following logic checks that the identifier X appears in an expression of the form X.method(...)
	// (i.e., it's the left-hand side of a SelectorExpr, which is the left-hand part of a CallExpr).
	// Anything else sets methodsOnly to false.

	if len(stack) < 2 {
		*iv.methodsOnly = false
		return nil
	}
	parent := stack[len(stack)-1]
	parentNode := parent.N
	sel, ok := parentNode.(*ast.SelectorExpr)
	if !ok {
		*iv.methodsOnly = false
		return nil
	}
	if which != astw.SelectorExpr_X {
		*iv.methodsOnly = false
		return nil
	}
	grandparentNode := stack[len(stack)-2].N
	if _, ok := grandparentNode.(*ast.CallExpr); !ok {
		*iv.methodsOnly = false
		return nil
	}
	if parent.W != astw.CallExpr_Fun {
		*iv.methodsOnly = false
		return nil
	}

	iv.methodNames[sel.Sel.Name] = true
	return nil
}
