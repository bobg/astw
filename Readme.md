# Astw - Enhanced abstract-syntax-tree walker for Go

In the Go standard library is the package `go/ast`,
which defines functions and types for understanding parsed Go code.
This includes the function `Walk`,
which calls a visitor’s `Visit` method on each `Node` in a syntax tree.

For many applications,
the visitor will need more context than just the node it’s currently visiting.
So implementations will typically do such things as maintaining a stack of nodes being visited,
and distinguishing between the recursive visits of one subnode versus another.

This package, `astw`,
provides an enhanced `Walk` function with a `Visitor` based on callbacks,
one for each syntax-tree node type.
It tracks the extra details that syntax-tree walkers typically need.

Each callback for a given node is called twice:
once before visiting its children
(the “pre” visit)
and once after
(the “post” visit).
This is true even for nodes that have no children.

Each callback receives as arguments:
- The node being visited;
- An enumerated constant describing which child of its parent this is;
- A slice index, in case this node is in a slice of its parent’s children;
- The stack of nodes above this node in the syntax tree;
- A boolean telling whether this is a “pre” (true) visit or a “post” (false) visit;
- The error, if any, produced by visiting the node’s children (always `nil` during “pre” visits).

Callbacks are never invoked on `nil` nodes.

In a “post” visit,
when the received error is non-`nil`,
a callback may decide whether and how to propagate the error to the caller.
A node callback should typically begin with

```go
if err != nil {
  return err
}
```

unless it wants to ignore, decorate, or otherwise alter errors in its subtree.

Nodes in a Go syntax tree have concrete types like `*ast.IfStmt` and `*ast.BinaryExpr`.
Each concrete type has its own callback in `Visitor`:
`Visitor.IfStmt`, `Visitor.BinaryExpr`, and so on.

Most node types also implement one of these abstract interfaces:
`ast.Expr`, `ast.Stmt`, `ast.Decl`, and `ast.Spec`.
`Visitor` has callbacks for these types too.
If the concrete-type callback for a given node is set,
then that’s used when the node is visited.
On the other hand,
if it’s not set but the abstract-type callback is,
then _that_ callback is used when the node is visited.

For example,
if your program sets `Visitor.BinaryExpr` and `Visitor.Expr`,
then your `BinaryExpr` callback will be called for every `*ast.BinaryExpr` node you visit,
and your `Expr` callback will be called for every other `ast.Expr` node you visit that’s not a `*ast.BinaryExpr`.

All syntax node types implement the interface `ast.Node`.
`Visitor.Node` is a callback for this catch-all type.
If you define a `Node` callback then it is called for every node that doesn’t have a more-specific callback.

## Detailed example

Consider this Go program fragment:

```go
if x == 7 {
  y++
  return z
}
```

After parsing, this is represented by a `*ast.IfStmt`.

Now imagine this `IfStmt` is passed to this package’s `Walk` function,
together with a `Visitor` with suitable callbacks defined.
Here is the sequence of events that will occur,
assuming all the relevant callbacks in the `Visitor`, `v`, are non-`nil`.
(Callbacks that are `nil` are simply skipped.)

1. v.IfStmt( _the IfStmt_, `Top`, 0, nil, true, nil )

This is the “pre” visit of the top `IfStmt` node.
Because this node was reached directly via the `Walk` function,
there is no information about its parent.
So the `Which` value is `Top` and the stack is empty.

2. v.BinaryExpr( _the x==7 node_, `IfStmt_Cond`, 0, [ _the IfStmt_ ], true, nil )

Now the `IfStmt`’s condition subnode is visited,
via its abstract `Expr` type.

The `Which` value,
`IfStmt_Cond`,
tells which child of the `IfStmt` this is.
(It’s the `Cond` field of the `IfStmt` type.)

Its parent, the `IfStmt` itself, is in the stack passed to `v.Expr`.

Note also that the `IfStmt` type includes an optional `Init` sub-statement,
but this `IfStmt` doesn’t use one,
so that callback is skipped.

3. v.Ident( _the x node_, `BinaryExpr_X`, 0, [ _the IfStmt_, _the x==7 node_ ], true, nil )

The `x` part of `x==7` is visited.

4. v.Ident( _the x node_, `BinaryExpr_X`, 0, [ _the IfStmt_, _the x==7 node_ ], false, err )

A `*Ident` has no children,
so now it’s time for the “post” visit of the same node on the way out of this subtree.

The value of `err` in step 4 is whatever error was returned in step 3.

5. v.BasicLit( _the 7 node_, `BinaryExpr_Y`, 0, [ _the IfStmt_, _the x==7 node_ ], true, nil )
6. v.BasicLit( _the 7 node_, `BinaryExpr_Y`, 0, [ _the IfStmt_, _the x==7 node_ ], false, err )

The `7` is pre-visited and post-visited.
It’s a “basic literal.”

7. v.BinaryExpr( _the x==7 node_, `IfStmt_Cond`, 0, [ _the IfStmt_ ], false, err )

Continuing to unwind the call stack, the `x==7` node is now post-visited.

8. v.BlockStmt( _the { ... } node_, `IfStmt_Body`, 0, [ _the IfStmt_ ], true, nil )

The next child of the `IfStmt`, the body, is now pre-visited.

9. v.IncDecStmt( _the y++ node_, `BlockStmt_List`, 0, [ _the IfStmt_, _the { ... } node_ ], true, nil )

A `BlockStmt` contains a field,
`List`,
whose value is a slice of `ast.Stmt`.
So now we visit each statement in the `BlockStmt`’s list,
starting with the `y++` statement.

The value of the `index` parameter, 0,
which is irrelevant for child nodes that aren’t part of a slice,
now tells us that this is the first element in the `BlockStmt`’s list.

10. v.Ident( _the y node_, `IncDecStmt_X`, 0, [ _the IfStmt_, _the { ... } node_, _the y++ node_ ], true, nil )
11. v.Ident( _the y node_, `IncDecStmt_X`, 0, [ _the IfStmt_, _the { ... } node_, _the y++ node_ ], false, err )

We descend into and then out of the sole child of the `y++` node.

12. v.IncDecStmt( _the y++ node_, `BlockStmt_List`, 0, [ _the IfStmt_, _the { ... } node_ ], false, err )

Post-visiting the `y++` node.

13. v.ReturnStmt( _the return z node_, `BlockStmt_List`, 1, [ _the IfStmt_, _the { ... } node_ ], true, nil )

We now visit the second child of the `BlockStmt`:
the `return z` node.

The value of the `index` parameter, 1,
tells us that this is the second element in the `BlockStmt`’s list.

14. v.Ident( _the z node_, `ReturnStmt_Results`, 0, [ _the IfStmt_, _the { ... } node_, _the return z node_ ], true, nil )
15. v.Ident( _the z node_, `ReturnStmt_Results`, 0, [ _the IfStmt_, _the { ... } node_, _the return z node_ ], false, err )

Descending into and out of the sole child of the `return z` node.

16. v.ReturnStmt( _the return z node_, `BlockStmt_List`, 1, [ _the IfStmt_, _the { ... } node_ ], false, err )
17. v.BlockStmt( _the { ... } node_, `IfStmt_Body`, 0, [ _the IfStmt_ ], false, err )
18. v.IfStmt( _the IfStmt_, `Top`, 0, nil, false, err )

Post-visiting everything on the way out of the tree,
all the way back to the top.
