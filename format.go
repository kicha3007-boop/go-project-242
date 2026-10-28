package code

import "fmt"

// unitBase — во сколько раз соседние единицы измерения отличаются друг от друга.
const unitBase = 1024

// units — единицы измерения по возрастанию.
var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

// formatSize переводит размер в байтах в строку: "123B", а с human — в наибольшую подходящую единицу ("24.0MB").
func formatSize(size int64, human bool) string {
	if !human || size < unitBase {
		return fmt.Sprintf("%dB", size)
	}

	value := float64(size)
	unit := 0

	for value >= unitBase && unit < len(units)-1 {
		value /= unitBase
		unit++
	}

	return fmt.Sprintf("%.1f%s", value, units[unit])
}
