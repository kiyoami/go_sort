package main

// Buble バブルソート
func Buble(data Datasets) {
	size := len(data)
	for sorted := 0; sorted < size; sorted++ {
		for i := 1; i < size-sorted; i++ {
			if data[i].weight < data[i-1].weight {
				data.swap(i, i-1)
			}
		}
	}
}
