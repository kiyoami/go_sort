package sort

import "../dataset"

// Quick クイックソート
func Quick(data dataset.Datasets) {
	partition(data[0:len(data)])
	// Insert(data) // 途中で止めた続きは挿入ソート
}

func partition(data dataset.Datasets) {
	last := len(data) - 1
	// if last < 10 { // 途中で止める
	// 	return
	// }
	if last < 1 {
		return
	}

	// referenceBetter(data) // 最悪ケース回避の細工

	reference := data[0].Weight
	i := 1 // 0 は reference なので分割の対象外。分割後に間に置く
	j := last
	for i < j {
		// 先頭からスキャンして reference より重い人を探す
		for i < j && data[i].Weight < reference {
			i++
		}

		// 末尾からスキャンして reference より軽い人を探す
		for i <= j && reference <= data[j].Weight {
			j--
		}

		if i < j && data[j].Weight < data[i].Weight {
			data.Swap(i, j)
		}
	}
	if data[j].Weight < reference {
		data.Swap(0, j) // reference を分割したスライスの間に置く
	}

	if 0 < i {
		partition(data[:i-1])
	}
	if i < last {
		partition(data[i:])
	}
}

func referenceBetter(data dataset.Datasets) {
	last := len(data) - 1
	mid := int(last / 2)

	weightFirst := data[0].Weight
	weightMid := data[mid].Weight
	weightLast := data[last].Weight

	if (weightFirst <= weightMid && weightMid < weightLast) || (weightLast < weightMid && weightMid <= weightFirst) {
		data.Swap(0, mid)
	} else if (weightMid <= weightFirst && weightFirst < weightLast) || (weightLast < weightFirst && weightFirst <= weightMid) {
		data.Swap(0, last)
	}
}
