package mysort

import (
	"fmt"
	"math/rand"
	"testing"

	"dataset"
)

const (
	dataMax int = 1000
	ok      int = -1
)

func dataRandom() dataset.Datasets {
	data := make(dataset.Datasets, dataMax, dataMax)
	for i := 0; i < len(data); i++ {
		data[i] = dataset.Dataset{
			Name:   fmt.Sprintf("data%05d", i),
			Weight: float64(rand.Int31n(10000)) / 10.0,
		}
	}
	return data
}

func dataSeq() dataset.Datasets {
	data := make(dataset.Datasets, dataMax, dataMax)
	for i := 0; i < len(data); i++ {
		data[i] = dataset.Dataset{
			Name:   fmt.Sprintf("data%05d", i),
			Weight: float64(10000-i) / 10.0,
		}
	}
	return data
}

func checkOrder(sorted dataset.Datasets) int {
	var beforeValue float64
	for i, current := range sorted {
		if beforeValue > current.Weight {
			return i
		}
		beforeValue = current.Weight
	}
	return ok
}

func TestBuble(t *testing.T) {
	sourceData := dataRandom()

	sorted := Buble(sourceData)

	result := checkOrder(sorted)
	if result != ok {
		t.Error("バブルソート", result, sorted)
	}
}

func TestSelect(t *testing.T) {
	sourceData := dataRandom()

	sorted := Select(sourceData)

	result := checkOrder(sorted)
	if result != ok {
		t.Error("選択ソート", result, sorted)
	}
}

func TestInsert(t *testing.T) {
	sourceData := dataRandom()

	sorted := Insert(sourceData)

	result := checkOrder(sorted)
	if result != ok {
		t.Error("挿入ソート", result, sorted)
	}
}

func TestShell(t *testing.T) {
	sourceData := dataRandom()

	sorted := Shell(sourceData)

	result := checkOrder(sorted)
	if result != ok {
		t.Error("シェルソート", result, sorted)
	}
}

func TestQuick(t *testing.T) {
	sourceData := dataRandom()

	sorted := Quick(sourceData)

	result := checkOrder(sorted)
	if result != ok {
		t.Error("クイックソート", result)
	}
}

func TestQuick2(t *testing.T) {
	sourceData := dataSeq()

	sorted := Quick(sourceData)

	result := checkOrder(sorted)
	if result != ok {
		t.Error("クイックソート 最悪", result)
	}
}

func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sourceData := dataRandom()
		Quick(sourceData)
	}
}

func BenchmarkQuick2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sourceData := dataSeq()
		Quick(sourceData)
	}
}
