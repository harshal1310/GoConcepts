// utils/mathutils.go
package utils

func Add(a int, b int) int { // Exported
	return a + b
}

var Pi = 3.14 // Exported

func subtract(a int, b int) int { // Not accessible outside
	return a - b
}

var version = "1.0" // Not exported
