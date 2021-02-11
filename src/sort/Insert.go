package sort

import "../dataset"

// Insert 挿入ソート
func Insert(data dataset.Datasets) {
	size := len(data)
	for sorted := 1; sorted < size; sorted++ {
		pivot := 0 // 挿入位置が見つからないときは、先頭に追加する
		for i := sorted - 1; 0 <= i; i-- {
			if data[i].Weight <= data[sorted].Weight {
				// 挿入位置を発見!!!
				pivot = i + 1
				break
			}
		}
		if pivot < sorted {
			data.Rotate(pivot, sorted, 1)
		}
	}
}
