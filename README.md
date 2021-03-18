# Cost of embedding Go structs with methods

The surprising part about embedding Go structs with methods in other
structs is that every such methods is copied into object files,
causing binary bloat. So if you have a struct with `M` methods that
`N` other structs embed, you have `N * M` cost in symbol table size.

## Motivation

This came up investigationg slow `go build` times. Causation:

- `go build` is slow on a small program

- because `go tool link` is spending a lot of time

- because one of the packages `P` being linked is bloated with a lot
  (1M) of symbols

- because `P` has hundreds of structs embedding other structs from
  package `Q` that have a lot of methods, and every such struct/method
  combination has a symbol in the `P` binary package

## To repro

Check out the example with a.go defining a struct with `A1` method,
and `b.go` defining a struct `B` that embeds `A`.


```
$ ./repro.sh

         U gist.github.com/t0yv0/ef2cffcc5c3308563be40a0885d0a66c/a.A.A1
    33ab T gist.github.com/t0yv0/ef2cffcc5c3308563be40a0885d0a66c/b.(*B).A1
    3463 T gist.github.com/t0yv0/ef2cffcc5c3308563be40a0885d0a66c/b.B.A1
```


The script `repro.sh` hacks into the Go compilation pass to find the
`_pkg_.a` package archive corresponding to the b package, and calls
`go tool nm` on it. Notice that there is an entry for `b.(*B).A1()`,
although package `b` does not define `A1`.
