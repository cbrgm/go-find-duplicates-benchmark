package main

import (
	"strconv"
	"testing"
)

const TESTDATA_ELEMENTS = 1000000

func genStringSlice(size int) []string {
	testData := make([]string, size)
	for i := range testData {
		testData[i] = strconv.Itoa(i)
	}
	return testData
}

func BenchmarkFindDuplicatesTypeToBool(b *testing.B) {
	strings := genStringSlice(TESTDATA_ELEMENTS)
	// Add some duplicates to the slice
	strings = append(strings, "0", "1", "2", "3", "4", "5", "6", "7", "8", "9")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findDuplicatesTypeToBool(strings)
	}
}

func BenchmarkFindDuplicatesTypeToInterface(b *testing.B) {
	strings := genStringSlice(TESTDATA_ELEMENTS)
	// Add some duplicates to the slice
	strings = append(strings, "0", "1", "2", "3", "4", "5", "6", "7", "8", "9")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findDuplicatesTypeToInterface(strings)
	}
}

func BenchmarkFindDuplicatesTypeToStruct(b *testing.B) {
	strings := genStringSlice(TESTDATA_ELEMENTS)
	// Add some duplicates to the slice
	strings = append(strings, "0", "1", "2", "3", "4", "5", "6", "7", "8", "9")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findDuplicatesTypeToStruct(strings)
	}
}
