# What each stdlib package costs in a TinyGo binary

Why `gocheck` forbids some stdlib packages and allows others. A package is forbidden **only when it makes the
browser binary measurably heavier** than the webtyp replacement; a ban that saves nothing is a rule nobody can justify.

Reproduce: `benchmark/weight/measure.sh` (TinyGo 0.41.1, `-target wasm -opt z -no-debug`, 2026-11-05).
"Alone" = the package in an empty program. "Next to webtyp/fmt" = the realistic case, an app that already imports `webtyp/fmt`.

| Package | stdlib alone | webtyp alone | Next to webtyp/fmt: stdlib vs webtyp | Verdict |
|---|---|---|---|---|
| `fmt` | 258 KB | 60 KB | — | **forbidden** → `webtyp.com/fmt` |
| `encoding/json` | 333 KB | 56 KB | — | **forbidden** → `webtyp.com/json` |
| `reflect` | +15 KB over empty | — | — | **forbidden** |
| `strconv` | 22 KB | 54 KB | 63.0 KB vs 56.3 KB | **forbidden** → `fmt.Convert` (−7 KB) |
| `strings` (3 functions) | 26 KB | 50 KB | 63.5 KB vs 63.7 KB | allowed: no difference |
| `strings` (the 15 helpers webtyp/fmt duplicates) | 33.0 KB | 61.9 KB | **71.3 KB vs 76.0 KB** | allowed, and **lighter** than the webtyp/fmt copies (−4.7 KB) |
| `errors.New`, `errors.AsType` | +0 KB | — | — | allowed |
| `maps` | +0.7 KB over a plain map | — | — | allowed (using maps at all is a separate rule: skill **wasm**) |

Programs `all_*` and `mall_*` use `Contains`, `HasPrefix`, `HasSuffix`, `Index`, `LastIndex`, `Split`, `TrimSpace`,
`TrimPrefix`, `TrimSuffix`, `ToLower`, `ToUpper`, `Repeat`, `Count`, `ReplaceAll` and `Join`/`JoinSlice` together; `mall_*`
also call `fmt.Sprintf` and `fmt.Err`, as a real app does. Baseline `mall_none` (same app, no string helpers): 61.2 KB.
So in a real app the stdlib `strings` set **adds 10.0 KB** and the webtyp/fmt copies **add 14.8 KB**. TinyGo drops exported
functions nobody calls, so deleting the copies does not make `webtyp/fmt` itself lighter: the saving (−4.8 KB) appears only in
apps that use string helpers.

## Comparing errors: the hidden 27 KB

Comparing two interface values (`err == io.EOF`, `err != ErrX`, and therefore `errors.Is`) makes TinyGo link its
runtime reflection: **+27 KB alone, +24 KB next to webtyp/fmt**. Comparing with `nil` costs nothing.

| Way to check an error | Cost |
|---|---|
| `err != nil` | 0 |
| `e, ok := err.(*NotFound)` | 0 |
| `e, ok := errors.AsType[*NotFound](err)` | 0 |
| `err == io.EOF` / `errors.Is(err, io.EOF)` | +27 KB |

So `gocheck` forbids comparing two interface values (other than with `nil`) and `errors.Is`; check the error's **type**
instead (`errors.AsType`). Programs: `err_*`, `m_wt_eq`, `m_wt_noeq`.

## Speed on the server

Weight is only half: `fmt/benchmark/benchmark_results.md` measures `webtyp/fmt` **3–4× slower** than the stdlib on
the server (e.g. 13 748 vs 3 431 ns/op in string processing), with fewer allocations. Server code pays that price
when the ban applies to it (CODE_GUARD master, D11).
