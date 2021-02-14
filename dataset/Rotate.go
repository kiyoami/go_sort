package dataset

// Rotate 値をずらす
func (data Datasets) Rotate(first int, last int, step int) {
	save := data[last]
	for i := last; first < i; i -= step {
		data[i] = data[i-1]
	}
	data[first] = save
}
