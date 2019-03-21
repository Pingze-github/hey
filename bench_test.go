package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var (
	testUrl = "http://127.0.0.1:10001"
)

func init() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, Hey!")
		})
		err := http.ListenAndServe(":10001", nil)
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Millisecond * 100)
}

func BenchmarkRunNetHTTP(b *testing.B) {
	w := ForTestGetWorker(testUrl, false)
	w.Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.MakeRequestOnly()
	}
}

func BenchmarkRunFastHTTP(b *testing.B) {
	w := ForTestGetWorker(testUrl, true)
	w.Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.MakeRequestOnly()
	}
}

func BenchmarkRunParallelNetHTTP(b *testing.B) {
	w := ForTestGetWorker(testUrl, false)
	w.Init()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.MakeRequestOnly()
		}
	})
}

func BenchmarkRunParallelFastHTTP(b *testing.B) {
	w := ForTestGetWorker(testUrl, true)
	w.Init()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.MakeRequestOnly()
		}
	})
}

func BenchmarkRunParallelNetHTTP128(b *testing.B) {
	w := ForTestGetWorker(testUrl, false)
	w.Init()
	b.ResetTimer()
	b.SetParallelism(128)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.MakeRequestOnly()
		}
	})
}

func BenchmarkRunParallelFastHTTP128(b *testing.B) {
	w := ForTestGetWorker(testUrl, true)
	w.Init()
	b.ResetTimer()
	b.SetParallelism(128)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.MakeRequestOnly()
		}
	})
}
