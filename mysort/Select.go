package mysort

import "../dataset"

// Select 選択ソート
func Select(sourceData dataset.Datasets) dataset.Datasets {
	data := sourceData.Copy()

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
	return data
}
