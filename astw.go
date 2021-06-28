// Package astw implements an enhanced walker/visitor for Go abstract syntax trees.
package astw

import "go/ast"

// Which describes a specific child of a specific go/ast Node type.
type Which int

// Values for Which.
// The names of these values have the form x_y,
// where x is the base name of a concrete go/ast struct type
// (e.g., ForStmt)
// and y is the name of a field in that struct
// (e.g., Init, Cond, Post, and Body in the case of ForStmt).
const (
	// Top is the Which value that Walk uses for the node at the top of the syntax tree.
	Top Which = iota

	Package_Files

	File_Doc
	File_Name
	File_Decls
	File_Imports
	File_Unresolved
	File_Comments

	CommentGroup_List

	Ellipsis_Elt
	FuncLit_Type
	FuncLit_Body
	CompositeLit_Type
	CompositeLit_Elts
	ParenExpr_X
	SelectorExpr_X
	SelectorExpr_Sel
	IndexExpr_X
	IndexExpr_Index
	SliceExpr_X
	SliceExpr_Low
	SliceExpr_High
	SliceExpr_Max
	TypeAssertExpr_X
	TypeAssertExpr_Type
	CallExpr_Fun
	CallExpr_Args
	StarExpr_X
	UnaryExpr_X
	BinaryExpr_X
	BinaryExpr_Y
	KeyValueExpr_Key
	KeyValueExpr_Value
	ArrayType_Len
	ArrayType_Elt
	StructType_Fields
	FuncType_Params
	FuncType_Results
	InterfaceType_Methods
	MapType_Key
	MapType_Value
	ChanType_Value

	FieldList_List

	Field_Doc
	Field_Names
	Field_Type
	Field_Tag
	Field_Comment

	DeclStmt_Decl
	LabeledStmt_Label
	LabeledStmt_Stmt
	ExprStmt_Expr
	SendStmt_Chan
	SendStmt_Value
	IncDecStmt_X
	AssignStmt_Lhs
	AssignStmt_Rhs
	GoStmt_Call
	DeferStmt_Call
	ReturnStmt_Results
	BranchStmt_Label
	BlockStmt_List
	IfStmt_Init
	IfStmt_Cond
	IfStmt_Body
	IfStmt_Else
	CaseClause_List
	CaseClause_Body
	SwitchStmt_Init
	SwitchStmt_Tag
	SwitchStmt_Body
	TypeSwitchStmt_Init
	TypeSwitchStmt_Assign
	TypeSwitchStmt_Body
	CommClause_Comm
	CommClause_Body
	SelectStmt_Body
	ForStmt_Init
	ForStmt_Cond
	ForStmt_Post
	ForStmt_Body
	RangeStmt_Key
	RangeStmt_Value
	RangeStmt_X
	RangeStmt_Body

	GenDecl_Doc
	GenDecl_Specs

	FuncDecl_Doc
	FuncDecl_Recv
	FuncDecl_Name
	FuncDecl_Type
	FuncDecl_Body

	ImportSpec_Doc
	ImportSpec_Name
	ImportSpec_Path
	ImportSpec_Comment

	ValueSpec_Doc
	ValueSpec_Names
	ValueSpec_Type
	ValueSpec_Values
	ValueSpec_Comment

	TypeSpec_Doc
	TypeSpec_Name
	TypeSpec_Type
	TypeSpec_Comment
)

// StackItem is an item in the stack that callbacks receive.
type StackItem struct {
	// N is a node in the syntax tree.
	N ast.Node

	// W is the Which value for node N.
	W Which

	// I is the index value for node N.
	I int
}

// Visitor is a structure full of callbacks,
// for walking a syntax tree via the Walk function.
//
// Depending on the application,
// it is usually necessary to set values only for a few callbacks.
// A zero Visitor is usable, but won't do anything interesting.
//
// The name of each callback corresponds to a type in the go/ast package.
// That callback is called twice for each node of that type to be encountered in the walk:
// once when descending into the tree, before child nodes are visited
// (the "pre" visit),
// and once when traversing back out of the tree,
// after child nodes are visited
// (the "post" visit).
//
// Each callback takes the following arguments:
//   - the node (of the appropriate type);
//   - a Which value, telling which child of the parent this node is;
//   - an index value, for when this node is one member of a slice in its parent, telling which member it is;
//   - a stack of nodes above this one in the tree, with stack[0] the root of the tree and stack[len(stack)-1] the node's immediate parent;
//   - a boolean telling whether this is the pre visit (true) or the post visit (false);
//   - the error value, if any, produced by visiting this node's children. This is always nil in a "pre" visit.
//
// Each node in a syntax tree has a concrete type like *ast.IfStmt or *ast.BinaryExpr.
// Visitor contains a callback for each possible concrete type.
// Many of these types also implement one of these abstract interfaces:
// ast.Expr, ast.Stmt, ast.Decl, and ast.Spec.
// Visitor additionally contains callbacks for these abstract types.
// If the concrete-type callback for a given node is not set,
// but the abstract-type callback is,
// then that will be the callback for the node.
//
// All node types implement the abstract interface ast.Node.
// Visitor contains a callback for that catch-all type too.
// If it is defined,
// it is used for all nodes that do not have a more-specific callback defined.
//
// Package nodes are handled specially.
// Unique among go/ast Node types,
// a Package node's children exist not as individual fields or in a slice,
// but in a map,
// mapping filenames to *ast.File nodes.
// When File nodes are visited via a Package node,
// they are visited in lexical filename order.
// The index value passed to the File callback reflects this ordering.
// When the File callback is invoked via a Package node parent
// (i.e., the Which value is Package_Files),
// the filename from the map can be found in the Visitor's Filename field.
type Visitor struct {
	// Node is the catch-all callback for any node that has no more-specific callback defined
	// (i.e., a concrete-type callback, or one of Expr, Stmt, Decl, or Spec).
	Node func(node ast.Node, which Which, index int, stack []StackItem, pre bool, err error) error

	// Expr, Stmt, Decl, and Spec are the catch-all callbacks
	// for any ast.Expr, ast.Stmt, ast.Decl, or ast.Spec node (respectively)
	// that has no concrete-type callback defined.

	Expr func(expr ast.Expr, which Which, index int, stack []StackItem, pre bool, err error) error
	Stmt func(stmt ast.Stmt, which Which, index int, stack []StackItem, pre bool, err error) error
	Decl func(decl ast.Decl, which Which, index int, stack []StackItem, pre bool, err error) error
	Spec func(spec ast.Spec, which Which, index int, stack []StackItem, pre bool, err error) error

	// Filename is the name of the current file,
	// when the File callback is invoked as a child of a Package node.
	Filename string

	// Package, File, Comment, CommentGroup, FieldList, and Field
	// are Nodes that are not Exprs, Stmts, Decls, or Specs.

	Package      func(pkg *ast.Package, which Which, index int, stack []StackItem, pre bool, err error) error
	File         func(file *ast.File, which Which, index int, stack []StackItem, pre bool, err error) error
	Comment      func(comment *ast.Comment, which Which, index int, stack []StackItem, pre bool, err error) error
	CommentGroup func(commentGroup *ast.CommentGroup, which Which, index int, stack []StackItem, pre bool, err error) error
	FieldList    func(fieldList *ast.FieldList, which Which, index int, stack []StackItem, pre bool, err error) error
	Field        func(field *ast.Field, which Which, index int, stack []StackItem, pre bool, err error) error

	// Concrete-type callbacks for ast.Expr nodes.

	BadExpr        func(badExpr *ast.BadExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	Ident          func(ident *ast.Ident, which Which, index int, stack []StackItem, pre bool, err error) error
	Ellipsis       func(ellipsis *ast.Ellipsis, which Which, index int, stack []StackItem, pre bool, err error) error
	BasicLit       func(basicLit *ast.BasicLit, which Which, index int, stack []StackItem, pre bool, err error) error
	FuncLit        func(funcLit *ast.FuncLit, which Which, index int, stack []StackItem, pre bool, err error) error
	CompositeLit   func(compositeLit *ast.CompositeLit, which Which, index int, stack []StackItem, pre bool, err error) error
	ParenExpr      func(parenExpr *ast.ParenExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	SelectorExpr   func(selectorExpr *ast.SelectorExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	IndexExpr      func(indexExpr *ast.IndexExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	SliceExpr      func(sliceExpr *ast.SliceExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	TypeAssertExpr func(typeAssertExpr *ast.TypeAssertExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	CallExpr       func(callExpr *ast.CallExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	StarExpr       func(starExpr *ast.StarExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	UnaryExpr      func(unaryExpr *ast.UnaryExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	BinaryExpr     func(binaryExpr *ast.BinaryExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	KeyValueExpr   func(keyValueExpr *ast.KeyValueExpr, which Which, index int, stack []StackItem, pre bool, err error) error
	ArrayType      func(arrayType *ast.ArrayType, which Which, index int, stack []StackItem, pre bool, err error) error
	StructType     func(structType *ast.StructType, which Which, index int, stack []StackItem, pre bool, err error) error
	FuncType       func(funcType *ast.FuncType, which Which, index int, stack []StackItem, pre bool, err error) error
	InterfaceType  func(interfaceType *ast.InterfaceType, which Which, index int, stack []StackItem, pre bool, err error) error
	MapType        func(mapType *ast.MapType, which Which, index int, stack []StackItem, pre bool, err error) error
	ChanType       func(chanType *ast.ChanType, which Which, index int, stack []StackItem, pre bool, err error) error

	// Concrete-type callbacks for ast.Stmt nodes.

	BadStmt        func(badStmt *ast.BadStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	DeclStmt       func(declStmt *ast.DeclStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	EmptyStmt      func(emptyStmt *ast.EmptyStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	LabeledStmt    func(labeledStmt *ast.LabeledStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	ExprStmt       func(exprStmt *ast.ExprStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	SendStmt       func(sendStmt *ast.SendStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	IncDecStmt     func(incDecStmt *ast.IncDecStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	AssignStmt     func(assignStmt *ast.AssignStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	GoStmt         func(goStmt *ast.GoStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	DeferStmt      func(deferStmt *ast.DeferStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	ReturnStmt     func(returnStmt *ast.ReturnStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	BranchStmt     func(branchStmt *ast.BranchStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	BlockStmt      func(blockStmt *ast.BlockStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	IfStmt         func(ifStmt *ast.IfStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	CaseClause     func(caseClause *ast.CaseClause, which Which, index int, stack []StackItem, pre bool, err error) error
	SwitchStmt     func(switchStmt *ast.SwitchStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	TypeSwitchStmt func(typeSwitchStmt *ast.TypeSwitchStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	CommClause     func(commClause *ast.CommClause, which Which, index int, stack []StackItem, pre bool, err error) error
	SelectStmt     func(selectStmt *ast.SelectStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	ForStmt        func(forStmt *ast.ForStmt, which Which, index int, stack []StackItem, pre bool, err error) error
	RangeStmt      func(rangeStmt *ast.RangeStmt, which Which, index int, stack []StackItem, pre bool, err error) error

	// Concrete-type callbacks for ast.Decl nodes.

	BadDecl  func(badDecl *ast.BadDecl, which Which, index int, stack []StackItem, pre bool, err error) error
	GenDecl  func(genDecl *ast.GenDecl, which Which, index int, stack []StackItem, pre bool, err error) error
	FuncDecl func(funcDecl *ast.FuncDecl, which Which, index int, stack []StackItem, pre bool, err error) error

	// Concrete-type callbacks for ast.Spec nodes.

	ImportSpec func(importSpec *ast.ImportSpec, which Which, index int, stack []StackItem, pre bool, err error) error
	ValueSpec  func(valueSpec *ast.ValueSpec, which Which, index int, stack []StackItem, pre bool, err error) error
	TypeSpec   func(typeSpec *ast.TypeSpec, which Which, index int, stack []StackItem, pre bool, err error) error
}

// Walk walks the syntax tree rooted at n using the Visitor v.
func Walk(v *Visitor, n ast.Node) error {
	return v.visitAbstractNode(n, Top, 0, nil)
}
