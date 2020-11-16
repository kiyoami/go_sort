package main

// Select 選択ソート
func Select(data Datasets) {
	size := len(data)
	for sorted := 0; sorted < size; sorted++ {
		pivot := sorted
		for i := sorted + 1; i < size; i++ {
			if data[i].weight < data[pivot].weight {
				pivot = i
			}
		}
		if pivot != sorted {
			data.swap(pivot, sorted)
		}
	}
}
