package dataset

import "fmt"

// Disp 値確認用
func (data Datasets) Disp(caption string) {
	fmt.Println(caption)
	for i := 0; i < len(data); i++ {
		fmt.Printf("%4d %-10s %7.3fKg\n", i, data[i].Name, data[i].Weight)
	}
	fmt.Println("-------")
}
