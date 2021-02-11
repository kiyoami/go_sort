package dataset

import (
	"fmt"
	"math/rand"
)

// Create 乱数でデータを初期化
func (data Datasets) Create() {
	max := len(data)
	for i := 0; i < max; i++ {
		a := float64(i)/float64(max) + 0.5
		data[i] = Dataset{fmt.Sprintf("user%08d", i), rand.Float64()*50.0 + 50.0*a}
	}
}
