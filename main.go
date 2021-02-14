package main

import (
	"fmt"
	"time"

	"./dataset"
	"./mysort"
)

// sort 関数の実行時間を計測する
func clocking(sourceData dataset.Datasets, sort func(dataset.Datasets) dataset.Datasets, caption string) {
	startTime := time.Now()
	result := sort(sourceData)
	duration := time.Since(startTime)

	// result.Disp(fmt.Sprintf("%sソート結果", caption))
	fmt.Printf("%8dμ秒 %sソート\n", duration.Microseconds(), caption)
}

func main() {
	sourceData := make(dataset.Datasets, 5, 5)
	sourceData.Create()
	// sourceData.Disp("ソート前")

	clocking(sourceData, mysort.Buble, "バブル")
	clocking(sourceData, mysort.Select, "選択")
	clocking(sourceData, mysort.Insert, "挿入")
	clocking(sourceData, mysort.Shell, "シェル")
	clocking(sourceData, mysort.Quick, "クイック")
}
