package bubblesort

func BubbleSort(values [] int) {
    flag, valuesLength := true, len(values)
    for flag {
        flag = false
        for i := 1; i < valuesLength; i ++ {
            if values[i - 1] > values[i] {
                values[i], values[i - 1] = values[i - 1], values[i]
                flag = true
            }
        }
    }
}