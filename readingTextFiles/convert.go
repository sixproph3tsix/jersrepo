package main
import (
	"fmt"
	"strconv"
	)

func main() {
	GetBool("TRUE")
	GetInt("2")
	GetFloat64("3.14")
}


func GetBool(str1 string) bool {
	b, err := strconv.ParseBool(str1)
	if err != nil {
		fmt.Println("Error", err)
		return false
	}
	fmt.Printf("string was: %v, bool is: %v\n", str1, b)
	return b
}

func GetInt(str2 string) int {
	i, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("error", err)
		return 0
	}
	fmt.Printf("string was: %v, int is: %v\n", str2, i)
	return i
}

func GetFloat64(str3 string) float64 {
	f, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println("Error", err)
		return 0.0
	}
	fmt.Printf("String was: %v,\t Float64 is: %v", str3, f)
	return f
}