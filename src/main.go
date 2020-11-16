package main

import (
	"fmt"
	"time"
)

// sort 関数の実行時間を計測する
func clocking(sourceData Datasets, sort func(Datasets), caption string) {
	sourceLength := len(sourceData)
	data := make(Datasets, sourceLength, sourceLength)
	copy(data, sourceData)

	startTime := time.Now()
	sort(data)
	duration := time.Since(startTime)

	// data.disp(fmt.Sprintf("%sソート結果", caption))
	fmt.Printf("%8dμ秒 %sソート\n", duration.Microseconds(), caption)
}

func main() {
	sourceData := make(Datasets, 10000, 10000)
	sourceData.create()
	// sourceData.disp("ソート前")

	clocking(sourceData, Buble, "バブル")
	clocking(sourceData, Select, "選択")
	clocking(sourceData, Insert, "挿入")
	clocking(sourceData, Shell, "シェル")
	clocking(sourceData, Quick, "クイック")
}
