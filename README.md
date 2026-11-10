# smaz-pure-safe-go

> **100% Pure Safe Go port of the legendary Smaz compression algorithm, synthesized automatically from C99 by the DecodeIQ CodeTruth Engine (`--target go`).**
> **Zero CGo. Zero `unsafe`. Zero heap allocations per operation (`0 allocs/op`). 203+ MB/s throughput.**

[![Go Report Card](https://goreportcard.com/badge/github.com/coboliqcom/smaz-pure-safe-go)](https://goreportcard.com/report/github.com/coboliqcom/smaz-pure-safe-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Parity: 100% Verified](https://img.shields.io/badge/Parity-100%25%20Verified-brightgreen.svg)]()

---

## ⚡ Executive Summary

Enterprise migrations from legacy C/C++ often stall when facing binary dependencies, native wrappers, or unsafe pointer gymnastics. Typical automated converters either:
1. Wrap native code in **CGo**, crippling cross-compilation, inflating memory footprints, and imposing 150ns invocation overhead per call.
2. Rely on **`unsafe.Pointer`**, retaining memory safety vulnerabilities and buffer overflow vectors in production Go services.

**`smaz-pure-safe-go` proves that deterministic, automated AST-level lifting eliminates this trade-off.**

The code in this repository was lifted directly from Salvatore Sanfilippo's original [`smaz.c`](https://github.com/antirez/smaz) into idiomatic, zero-allocation Go using **DecodeIQ's CodeTruth Engine**.

---

## 🚀 Performance Benchmarks

Measured on dual-socket enterprise silicon (`Intel Xeon E5-2680 v2 @ 2.80GHz`, Windows / Go 1.24):

```text
goos: windows
goarch: amd64
pkg: smaz
cpu: Intel(R) Xeon(R) CPU E5-2680 v2 @ 2.80GHz

BenchmarkSmazCompress-20       462268       2624 ns/op     23.25 MB/s       0 B/op       0 allocs/op
BenchmarkSmazDecompress-20    4028990        299.2 ns/op  203.89 MB/s       0 B/op       0 allocs/op
PASS
ok      smaz    3.631s
```

### Key Architectural Invariants:
* **0 B/op & 0 allocs/op:** Zero runtime heap allocations. Zero garbage collector (GC) pauses.
* **Pure Idiomatic Go:** Slices and byte windows replace C raw pointer arithmetic (`slot += slot[0] + 2`).
* **Sub-Microsecond Latency:** Decompress operates in **~299 nanoseconds** at **203 MB/s**.

---

## 📦 Usage

```go
package main

import (
	"fmt"
	"github.com/coboliqcom/smaz-pure-safe-go"
)

func main() {
	src := []byte("the quick brown fox jumps over the lazy dog http://google.com")
	
	// Pre-allocate buffer or slice view
	compressed := make([]byte, len(src)*2+128)
	cLen := smaz.SmazCompress(src, len(src), compressed, len(compressed))
	
	decompressed := make([]byte, len(src)*4+128)
	dLen := smaz.SmazDecompress(compressed[:cLen], cLen, decompressed, len(decompressed))
	
	fmt.Printf("Original:     %d bytes\n", len(src))
	fmt.Printf("Compressed:   %d bytes (%.1f%% ratio)\n", cLen, float64(cLen)/float64(len(src))*100)
	fmt.Printf("Decompressed: %s\n", string(decompressed[:dLen]))
}
```

---

## 🧪 Verification & Testing

Run the test suite and benchmarks locally:

```bash
git clone https://github.com/coboliqcom/smaz-pure-safe-go.git
cd smaz-pure-safe-go
go test -v -bench=. -benchmem
```

---

## 🛡️ License

MIT License. See [LICENSE](LICENSE) for details.
Original Smaz C algorithm Copyright (c) 2006-2009 Salvatore Sanfilippo.
Synthesized and verified by [CobolIQ / DecodeIQ](https://coboliq.com).
