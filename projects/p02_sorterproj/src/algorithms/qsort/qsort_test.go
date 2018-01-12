package qsort

import "testing"

func TestQSort1(t *testing.T) {
    values := []int {5, 5, 4, 4, 4, 3, 3, 2, 1}
    v := make([]int ,len(values))
    copy(v, values)
    QSort(values)
    t.Errorf("before: ", values)
    t.Errorf(" after: ", v)
}