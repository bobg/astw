// Package astw implements an enhanced walker/visitor for Go abstract syntax trees.
package astw

import (
	"go/ast"
)

type Which int

const (
	Top Which = iota

	File_Name
	File_Decls
	File_Imports
	File_Unresolved

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
	Field_Names
	Field_Type

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

	GenDecl_Specs
	FuncDecl_Recv
	FuncDecl_Name
	FuncDecl_Type
	FuncDecl_Body

	ImportSpec_Name
	ValueSpec_Names
	ValueSpec_Type
	ValueSpec_Values
	TypeSpec_Name
	TypeSpec_Type
)

type StackItem struct {
	N ast.Node
	W Which
	I int
}

type Visitor struct {
	Node func(node ast.Node, which Which, index int, stack []StackItem, pre bool) error

	File func(file *ast.File, which Which, index int, stack []StackItem, pre bool) error

	Expr func(expr ast.Expr, which Which, index int, stack []StackItem, pre bool) error
	Stmt func(stmt ast.Stmt, which Which, index int, stack []StackItem, pre bool) error
	Decl func(decl ast.Decl, which Which, index int, stack []StackItem, pre bool) error
	Spec func(spec ast.Spec, which Which, index int, stack []StackItem, pre bool) error

	BadExpr        func(badExpr *ast.BadExpr, which Which, index int, stack []StackItem, pre bool) error
	Ident          func(ident *ast.Ident, which Which, index int, stack []StackItem, pre bool) error
	Ellipsis       func(ellipsis *ast.Ellipsis, which Which, index int, stack []StackItem, pre bool) error
	BasicLit       func(basicLit *ast.BasicLit, which Which, index int, stack []StackItem, pre bool) error
	FuncLit        func(funcLit *ast.FuncLit, which Which, index int, stack []StackItem, pre bool) error
	CompositeLit   func(compositeLit *ast.CompositeLit, which Which, index int, stack []StackItem, pre bool) error
	ParenExpr      func(parenExpr *ast.ParenExpr, which Which, index int, stack []StackItem, pre bool) error
	SelectorExpr   func(selectorExpr *ast.SelectorExpr, which Which, index int, stack []StackItem, pre bool) error
	IndexExpr      func(indexExpr *ast.IndexExpr, which Which, index int, stack []StackItem, pre bool) error
	SliceExpr      func(sliceExpr *ast.SliceExpr, which Which, index int, stack []StackItem, pre bool) error
	TypeAssertExpr func(typeAssertExpr *ast.TypeAssertExpr, which Which, index int, stack []StackItem, pre bool) error
	CallExpr       func(callExpr *ast.CallExpr, which Which, index int, stack []StackItem, pre bool) error
	StarExpr       func(starExpr *ast.StarExpr, which Which, index int, stack []StackItem, pre bool) error
	UnaryExpr      func(unaryExpr *ast.UnaryExpr, which Which, index int, stack []StackItem, pre bool) error
	BinaryExpr     func(binaryExpr *ast.BinaryExpr, which Which, index int, stack []StackItem, pre bool) error
	KeyValueExpr   func(keyValueExpr *ast.KeyValueExpr, which Which, index int, stack []StackItem, pre bool) error
	ArrayType      func(arrayType *ast.ArrayType, which Which, index int, stack []StackItem, pre bool) error
	StructType     func(structType *ast.StructType, which Which, index int, stack []StackItem, pre bool) error
	FuncType       func(funcType *ast.FuncType, which Which, index int, stack []StackItem, pre bool) error
	InterfaceType  func(interfaceType *ast.InterfaceType, which Which, index int, stack []StackItem, pre bool) error
	MapType        func(mapType *ast.MapType, which Which, index int, stack []StackItem, pre bool) error
	ChanType       func(chanType *ast.ChanType, which Which, index int, stack []StackItem, pre bool) error

	FieldList func(fieldList *ast.FieldList, which Which, index int, stack []StackItem, pre bool) error
	Field     func(field *ast.Field, which Which, index int, stack []StackItem, pre bool) error

	BadStmt        func(badStmt *ast.BadStmt, which Which, index int, stack []StackItem, pre bool) error
	DeclStmt       func(declStmt *ast.DeclStmt, which Which, index int, stack []StackItem, pre bool) error
	EmptyStmt      func(emptyStmt *ast.EmptyStmt, which Which, index int, stack []StackItem, pre bool) error
	LabeledStmt    func(labeledStmt *ast.LabeledStmt, which Which, index int, stack []StackItem, pre bool) error
	ExprStmt       func(exprStmt *ast.ExprStmt, which Which, index int, stack []StackItem, pre bool) error
	SendStmt       func(sendStmt *ast.SendStmt, which Which, index int, stack []StackItem, pre bool) error
	IncDecStmt     func(incDecStmt *ast.IncDecStmt, which Which, index int, stack []StackItem, pre bool) error
	AssignStmt     func(assignStmt *ast.AssignStmt, which Which, index int, stack []StackItem, pre bool) error
	GoStmt         func(goStmt *ast.GoStmt, which Which, index int, stack []StackItem, pre bool) error
	DeferStmt      func(deferStmt *ast.DeferStmt, which Which, index int, stack []StackItem, pre bool) error
	ReturnStmt     func(returnStmt *ast.ReturnStmt, which Which, index int, stack []StackItem, pre bool) error
	BranchStmt     func(branchStmt *ast.BranchStmt, which Which, index int, stack []StackItem, pre bool) error
	BlockStmt      func(blockStmt *ast.BlockStmt, which Which, index int, stack []StackItem, pre bool) error
	IfStmt         func(ifStmt *ast.IfStmt, which Which, index int, stack []StackItem, pre bool) error
	CaseClause     func(caseClause *ast.CaseClause, which Which, index int, stack []StackItem, pre bool) error
	SwitchStmt     func(switchStmt *ast.SwitchStmt, which Which, index int, stack []StackItem, pre bool) error
	TypeSwitchStmt func(typeSwitchStmt *ast.TypeSwitchStmt, which Which, index int, stack []StackItem, pre bool) error
	CommClause     func(commClause *ast.CommClause, which Which, index int, stack []StackItem, pre bool) error
	SelectStmt     func(selectStmt *ast.SelectStmt, which Which, index int, stack []StackItem, pre bool) error
	ForStmt        func(forStmt *ast.ForStmt, which Which, index int, stack []StackItem, pre bool) error
	RangeStmt      func(rangeStmt *ast.RangeStmt, which Which, index int, stack []StackItem, pre bool) error

	BadDecl  func(badDecl *ast.BadDecl, which Which, index int, stack []StackItem, pre bool) error
	GenDecl  func(genDecl *ast.GenDecl, which Which, index int, stack []StackItem, pre bool) error
	FuncDecl func(funcDecl *ast.FuncDecl, which Which, index int, stack []StackItem, pre bool) error

	ImportSpec func(importSpec *ast.ImportSpec, which Which, index int, stack []StackItem, pre bool) error
	ValueSpec  func(valueSpec *ast.ValueSpec, which Which, index int, stack []StackItem, pre bool) error
	TypeSpec   func(typeSpec *ast.TypeSpec, which Which, index int, stack []StackItem, pre bool) error
}

func Walk(v *Visitor, n ast.Node) error {
	return v.VisitNode(n, Top, 0, nil)
}
