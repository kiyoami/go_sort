package dataset

// Copy 新しい領域を割り当てて、データをコピーする
func (sourceData Datasets) Copy() Datasets {
	sourceLength := len(sourceData)
	data := make(Datasets, sourceLength, sourceLength)
	copy(data, sourceData)
	return data
}
