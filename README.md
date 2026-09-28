# invariant

`invariant` provides small assertions for programmer errors. They are enabled by default.

## Why invariants?

An invariant is a condition that must hold when the code is correct. A failed invariant points to a programmer error. This library was inspired by [TigerBeetle's TigerStyle](https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md), especially its guidance on assertions.

```go
invariant.True(account.Balance >= 0, "balance is negative")

switch kind {
case debit, credit:
default:
	invariant.Unreachable("unknown transaction kind")
}
```

Disable assertion panics with `-tags=invariant_disabled`. `True` and `Unreachable` then do nothing, and `Enabled` is false. Their arguments are still evaluated, so guard an expensive or side-effectful expression when it must be skipped:

```go
if invariant.Enabled {
	invariant.True(tree.IsConsistent(), "tree is inconsistent")
}
```

Failed assertions panic with an `invariant.Failure` value.
