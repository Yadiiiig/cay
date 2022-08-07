package core

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/rivo/tview"
// )

// var table = []struct {
// 	input rune
// }{
// 	{input: 34},
// 	{input: 39},
// 	{input: 40},
// 	{input: 91},
// 	{input: 123},
// 	{input: 100},
// }

// // Benchmarking 2 options on how to deal with autocompletion of special keys
// // Results can be found below

// func BenchmarkKeyStroke(b *testing.B) {
// 	text_view := tview.NewTextView()

// 	for _, v := range table {
// 		b.Run(fmt.Sprintf("key input: %s", string(v.input)), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				KeyStroke(text_view, v.input)
// 			}
// 		})
// 	}
// }

// func BenchmarkKeyStrokeMap(b *testing.B) {
// 	text_view := tview.NewTextView()

// 	for _, v := range table {
// 		b.Run(fmt.Sprintf("key input: %s", string(v.input)), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				KeyStrokeMap(text_view, v.input)
// 			}
// 		})
// 	}
// }

// /*

// tldr;
// 	KeyStrokeMap seems to have faster result on every input,
// 	plus it also looks alot cleaner then KeyStroke.

// ⇒  go test -bench=.
// goos: linux
// goarch: amd64
// pkg: ide/core
// cpu: AMD Ryzen 7 3700X 8-Core Processor
// BenchmarkKeyStroke/key_input:_"-16         	  615500	     77815 ns/op
// BenchmarkKeyStroke/key_input:_'-16         	    7630	    149374 ns/op
// BenchmarkKeyStroke/key_input:_(-16         	    7449	    148860 ns/op
// BenchmarkKeyStroke/key_input:_[-16         	    7779	    151442 ns/op
// BenchmarkKeyStroke/key_input:_{-16         	    9789	    151121 ns/op
// BenchmarkKeyStroke/key_input:_d-16         	    7188	    148038 ns/op

// BenchmarkKeyStrokeMap/key_input:_"-16      	  569817	     72736 ns/op
// BenchmarkKeyStrokeMap/key_input:_'-16      	    8440	    138108 ns/op
// BenchmarkKeyStrokeMap/key_input:_(-16      	    8713	    139956 ns/op
// BenchmarkKeyStrokeMap/key_input:_[-16      	    8066	    144692 ns/op
// BenchmarkKeyStrokeMap/key_input:_{-16      	   10000	    146221 ns/op
// BenchmarkKeyStrokeMap/key_input:_d-16      	    8592	    144650 ns/op
// PASS
// ok  	ide/core	101.774s

// ⇒  go test -bench=.
// goos: linux
// goarch: amd64
// pkg: ide/core
// cpu: AMD Ryzen 7 3700X 8-Core Processor
// BenchmarkKeyStroke/key_input:_"-16         	  654205	     82185 ns/op
// BenchmarkKeyStroke/key_input:_'-16         	    6667	    152901 ns/op
// BenchmarkKeyStroke/key_input:_(-16         	    6326	    160134 ns/op
// BenchmarkKeyStroke/key_input:_[-16         	    6381	    161364 ns/op
// BenchmarkKeyStroke/key_input:_{-16         	    6780	    158359 ns/op
// BenchmarkKeyStroke/key_input:_d-16         	    8556	    159521 ns/op

// BenchmarkKeyStrokeMap/key_input:_"-16      	  609153	     78021 ns/op
// BenchmarkKeyStrokeMap/key_input:_'-16      	    8072	    148390 ns/op
// BenchmarkKeyStrokeMap/key_input:_(-16      	    7533	    148800 ns/op
// BenchmarkKeyStrokeMap/key_input:_[-16      	    7525	    153386 ns/op
// BenchmarkKeyStrokeMap/key_input:_{-16      	    7005	    153056 ns/op
// BenchmarkKeyStrokeMap/key_input:_d-16      	    7628	    149922 ns/op
// PASS
// ok  	ide/core	112.706s
// */
