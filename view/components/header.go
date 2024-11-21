package components

import "fmt"

func Header(title string) {
	fmt.Printf("=============================== %v ===============================\n", title)
	fmt.Println()
}
