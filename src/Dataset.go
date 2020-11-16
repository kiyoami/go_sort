package main

import (
	"fmt"
	"math/rand"
)

// Dataset 型定義
type Dataset struct {
	name   string
	weight float64
}

// Datasets Datasetのスライス型
type Datasets []Dataset

// create 乱数でデータを初期化
func (data Datasets) create() {
	max := len(data)
	for i := 0; i < max; i++ {
		data[i] = Dataset{fmt.Sprintf("user%08d", i), rand.Float64()*50.0 + 50.0}
	}
}

// 値確認用
func (data Datasets) disp(caption string) {
	fmt.Println(caption)
	for i := 0; i < len(data); i++ {
		fmt.Printf("%4d %-10s %7.3fKg\n", i, data[i].name, data[i].weight)
	}
	fmt.Println("-------")
}

// swap 値の交換
func (data Datasets) swap(i int, j int) {
	save := data[i]
	data[i] = data[j]
	data[j] = save
}

// rotate 値をずらす
func (data Datasets) rotate(first int, last int, step int) {
	save := data[last]
	for i := last; first < i; i -= step {
		data[i] = data[i-1]
	}
	data[first] = save
}
