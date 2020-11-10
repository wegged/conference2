package main

import (
	"reflect"
	"testing"
)

func BenchmarkPar(b *testing.B) {
	content := readUrlToString("http://www.engadget.com")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		letterFrequencyConcurrent(content)
	}
}


func BenchmarkSeq(b *testing.B) {
	content := readUrlToString("http://www.engadget.com")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		letterFrequencySequential(content)
	}
}

func TestParEqSeq(t *testing.T) {
	content := readUrlToString("http://www.engadget.com")
	expected := letterFrequencySequential(content)
	parResult := letterFrequencyConcurrent(content)
	if !reflect.DeepEqual(expected, parResult) {
		t.Fail()
	}
}