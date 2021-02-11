package sort

import "../dataset"

// Select 選択ソート
func Select(data dataset.Datasets) {
	size := len(data)
	for sorted := 0; sorted < size; sorted++ {
		pivot := sorted
		for i := sorted + 1; i < size; i++ {
			if data[i].Weight < data[pivot].Weight {
				pivot = i
			}
		}
		if pivot != sorted {
			data.Swap(pivot, sorted)
		}
	}
}
