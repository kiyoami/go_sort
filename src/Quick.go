package main

// Quick クイックソート
func Quick(data Datasets) {
	partition(data[0:len(data)])
	// Insert(data)
}

func referenceBetter(data Datasets) {
	last := len(data) - 1
	mid := int(last / 2)

	weightFirst := data[0].weight
	weightMid := data[mid].weight
	weightLast := data[last].weight

	if (weightFirst <= weightMid && weightMid < weightLast) || (weightLast < weightMid && weightMid <= weightFirst) {
		data.swap(0, mid)
	} else if (weightMid <= weightFirst && weightFirst < weightLast) || (weightLast < weightFirst && weightFirst <= weightMid) {
		data.swap(0, last)
	}
}

func partition(data Datasets) {
	last := len(data) - 1
	if last < 1 {
		return
	}

	// referenceBetter(data) // 最悪ケース回避の細工

	reference := data[0].weight
	i := 1 // 0 は reference なので分割の対象外。分割後に間に置く
	j := last
	for i < j {
		// 先頭からスキャンして reference より重い人を探す
		for i < j && data[i].weight < reference {
			i++
		}

		// 末尾からスキャンして reference より軽い人を探す
		for i <= j && reference <= data[j].weight {
			j--
		}

		if i < j && data[j].weight < data[i].weight {
			data.swap(i, j)
		}
	}
	if data[j].weight < reference {
		data.swap(0, j) // reference を分割したスライスの間に置く
	}

	if 0 < i {
		partition(data[:i-1])
	}
	if i < last {
		partition(data[i:])
	}
}
