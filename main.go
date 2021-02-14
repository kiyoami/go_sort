package main

import (
	"fmt"
	"time"

	"./dataset"
	"./sort"
)

// sort 関数の実行時間を計測する
func clocking(sourceData dataset.Datasets, sort func(dataset.Datasets), caption string) {
	sourceLength := len(sourceData)
	data := make(dataset.Datasets, sourceLength, sourceLength)
	copy(data, sourceData)

	startTime := time.Now()
	sort(data)
	duration := time.Since(startTime)

	// data.disp(fmt.Sprintf("%sソート結果", caption))
	fmt.Printf("%8dμ秒 %sソート\n", duration.Microseconds(), caption)
}

func main() {
	sourceData := make(dataset.Datasets, 20000, 20000)
	sourceData.Create()
	// sourceData.disp("ソート前")

	clocking(sourceData, sort.Buble, "バブル")
	clocking(sourceData, sort.Select, "選択")
	clocking(sourceData, sort.Insert, "挿入")
	clocking(sourceData, sort.Shell, "シェル")
	clocking(sourceData, sort.Quick, "クイック")
}
