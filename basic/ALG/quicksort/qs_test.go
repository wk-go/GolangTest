package main

import "testing"
var(
    a = []int{2221,5767,866,65656,65445,6565,6767,8878,97,65,4334,66567,8787,665,335,778678,455,6,7,667,8564,898797,6666}
    a2 []int = a
)
func BenchmarkQuickSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        QuickSort(a, 0, len(a)-1)
    }
}

func BenchmarkQuickSortIterator(b *testing.B) {
    for i := 0; i < b.N; i++ {
        QuickSortIterator(a2, 0, len(a2)-1)
    }
}
