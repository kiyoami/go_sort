package main

// Insert 挿入ソート
func Insert(data Datasets) {
	size := len(data)
	for sorted := 1; sorted < size; sorted++ {
		pivot := 0 // 挿入位置が見つからないときは、先頭に追加する
		for i := sorted - 1; 0 <= i; i-- {
			if data[i].weight <= data[sorted].weight {
				// 挿入位置を発見!!!
				pivot = i + 1
				break
			}
		}
		if pivot < sorted {
			data.rotate(pivot, sorted, 1)
		}
	}
}
