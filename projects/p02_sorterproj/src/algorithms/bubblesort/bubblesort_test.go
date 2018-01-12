package bubblesort

import "testing"

func TestBubbleSort(t *testing.T) {
    v1 := [] int {5, 5, 4, 4, 4, 3, 3, 2, 1}
    v2 := make([] int, len(v1))
    copy(v2, v1)
    BubbleSort(v1)
    t.Errorf("before: ", v2)
    t.Errorf(" after: ", v1)
}