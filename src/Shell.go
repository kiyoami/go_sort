package main

import (
	"math"
)

// Shell シェルソート
func Shell(data Datasets) {
	size := len(data)
	base := 3
	exponent := math.Floor(math.Log(float64(size)) / math.Log(float64(base)))
	step := int(math.Pow(float64(base), exponent))
	for 0 < step {
		for offset := 0; offset < step; offset++ {
			for sorted := step + offset; sorted < size; sorted += step {
				pivot := offset // 挿入位置が見つからないときは、先頭に追加する
				for i := sorted - step; 0 <= i; i -= step {
					if data[i].weight <= data[sorted].weight {
						// 挿入位置を発見!!!
						pivot = i + step
						break
					}
				}
				if pivot < sorted {
					data.rotate(pivot, sorted, step)
				}
			}
		}
		step /= base
	}
}
