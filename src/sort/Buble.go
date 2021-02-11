package sort

import "../dataset"

// Buble バブルソート
func Buble(data dataset.Datasets) {
	size := len(data)
	for sorted := 0; sorted < size; sorted++ {
		for i := 1; i < size-sorted; i++ {
			if data[i].Weight < data[i-1].Weight {
				data.Swap(i, i-1)
			}
		}
	}
}
