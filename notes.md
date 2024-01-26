_Goal_: understand the _exact cover problem_ and Knuth's solution using _Algorithm X_. 

Implement through doubly linked lists and apply it to solve Sudoku and other similar problems. 

Links and references:

- https://www.ocf.berkeley.edu/~jchu/publicportal/sudoku/sudoku.paper.html
- https://stackoverflow.com/questions/1518335/the-dancing-links-algorithm-an-explanation-that-is-less-explanatory-but-more-o
- https://code.google.com/archive/p/narorumo/wikis/SudokuDLX.wiki
- https://www.ocf.berkeley.edu/~jchu/publicportal/sudoku/0011047.pdf
- https://en.wikipedia.org/wiki/Exact_cover
- https://shivankaul.com/blog/sudoku

## 1. Exact cover

Let $X$ be a finite set and let $\mathcal S$ be a collection of subsets of $X$. We say that $\mathcal S^*\subseteq \mathcal S$
is an _exact cover_ of $X$ if for each $x\in X$ there exists a unique $S_x \in \mathcal S^*$ such that $x\in S_x$. The 
_exact cover problem_ is the following:

> Given a finite set $X$ and a collection of subsets $\mathcal S$, find an exact cover $\mathcal S^*\subseteq X$ or
determine if none exists.

By labelling the elements of $X$ by integers $\{0,\ldots,n-1\}$, we can store each $S\subseteq X$ as a binary 
vector of length $n$, and an alternative formulation is then:

> Given a binary matrix with $k$ rows and $n$ columns, determine if there exists a subset of rows $\mathcal V$
such that for all distinct $v, w\in \mathcal V$ we have $v\cdot w = 0$, and such that $\sum_{v\in\mathcal V} v = \mathbf 1$,
the vector of ones. 

In other words, we want the rows to have no $1$s in common, and we want all columns to have at least one $1$ in one of the rows.
Clearly, the condition $\sum_{v\in\mathcal V} v = \mathbf 1$ is necessary but not sufficient, and it can thus be 
quickly checked before trying to find a solution. Moreover, we can also assume that each $v$ is not the zero vector. 

> *Observation*. It is clear that this decision problem belongs to the $\mathsf{NP}$ class. It is in fact  $\mathsf{NP}$-complete.

## 2. Algorithm X

The following algorithm provides a solution to the exact cover problem, it is due to Knuth. The input of the algorithm
is the $k\times n$ matrix $A$ representing the $k$ subsets in $\mathcal S^*$ of $\{0, \ldots, n-1\}$ and a partial solution
$P$ to the problem, consisting of some row indices. 

The algorithm  is recursive, and halts when a valid solution is recursively found. Thus, we allow the matrix $A$ to be empty (representing the trivial solution: the empty partition is an exact cover of the empty set). 

We assume that given such a non-empty 
matrix $A$, we have a deterministic procedure `GetColumn(A)` that selects a column of $A$. Given a column index $c$, we also 
assume we have a non-deterministic procedure `GetRow(A,c)` that returns an index $r$ such that $A_{r,c} = 1$ or fails otherwise.
We then follow these steps:

$\mathsf{ExactCover}(A, P)$

0. If $A$ is empty, terminate successfully and return $P$.
1. Let $c = \mathsf{GetColumn}(A)$. If column $c$ is zero, terminate _unsuccessfully_.
2. Let $r = \mathsf{GetRow}(A,c)$.
3. Add row index $r$ to $P$ to get $P'$.
4. For each $j$ such that $A_{r,j} = 1$ {
    for each $i$ such that $A_{i,j} = 1$ { 
        delete row $i$ }
    delete column $j$ } 
5. Obtain a new matrix $A'$ 
6. Recursively run $\mathsf{ExactCover}(A', P')$.

In this way, the algorithm is really just follows a trial-and-error approach, in which it builds a
tree, going one level downwards each time it chooses a new row.

## 3. Knuth's dancing links

An apparent issue with storing a collection of subsets as a binary matrix is that we will usually
deal with very sparse matrices. At the same time, we are tasked with the issue of deleting and adding 
back columns and rows in the back-tracking procedure. Knuth's paper highlights an idea of Hitotumatu and
Noshita, that if we have a simple linked list node storing a binary value

```go
type data struct {
    // some data here
}

type node struct {
    left  *node
    right *node
    data  data
}
```

then just like

```go 
func (x *node) Delete() {
    x.right.left = x.left
    x.left.right = x.right
}
```

removes a node from a list, the operation

```go
func (x *node) Restore() {
    x.right.left = x 
    x.left.right = x
}
```

returns it to the list, provided that we have not destroyed the pointers stored by $x$.