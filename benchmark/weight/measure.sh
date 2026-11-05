#!/usr/bin/env bash
# Builds every program in this folder with TinyGo (-target wasm -opt z -no-debug)
# and prints its size in bytes. Results are recorded in docs/STDLIB_WEIGHT.md.
set -euo pipefail
cd "$(dirname "$0")"
go mod tidy >/dev/null
out=$(mktemp -d)
for d in */; do
	d=${d%/}
	tinygo build -o "$out/$d.wasm" -target wasm -opt z -no-debug "./$d"
	printf "%-14s %7d\n" "$d" "$(stat -c%s "$out/$d.wasm")"
done
rm -rf "$out"
