package mysort

import "../dataset"

// Quick クイックソート
func Quick(sourceData dataset.Datasets) dataset.Datasets {
	data := sourceData.Copy()

	type rng struct {
		first int
		last  int
	}
	spMax := 100
	stack := make([]rng, spMax, spMax)
	sp := 0

	stack[sp] = rng{first: 0, last: len(data)}
	sp++
	for 0 < sp {
		sp--
		pop := stack[sp]
		pivot := pop.first + partition(data[pop.first:pop.last]) // 分割

		if pop.first < pivot {
			if spMax <= sp {
				stack = append(stack, rng{first: pop.first, last: pivot})
				spMax++
			} else {
				stack[sp] = rng{first: pop.first, last: pivot}
			}
			sp++
		}
		if pivot+1 < pop.last {
			if spMax <= sp {
				stack = append(stack, rng{first: pivot, last: pop.last})
				spMax++
			} else {
				stack[sp] = rng{first: pivot, last: pop.last}
			}
			sp++
		}
	}

	Insert(data) // 途中で止めた続きは挿入ソート
	return data
}

func partition(data dataset.Datasets) int {
	last := len(data) - 1
	if last < 1 {
		return last
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

	return i
}

func referenceBetter(data dataset.Datasets) {
	last := len(data) - 1
	mid := int(last / 2)

	weightFirst := data[0].Weight
	weightMid := data[mid].Weight
	weightLast := data[last].Weight

	if (weightFirst <= weightMid && weightMid < weightLast) || (weightLast < weightMid && weightMid <= weightFirst) {
		data.Swap(0, mid)
	} else if (weightFirst < weightLast && weightLast <= weightMid) || (weightMid <= weightLast && weightLast < weightFirst) {
		data.Swap(0, last)
	}
}
