package dataset

// Swap 値の交換
func (data Datasets) Swap(i int, j int) {
	save := data[i]
	data[i] = data[j]
	data[j] = save
}
