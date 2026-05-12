package functions


func Squre(x int) int { return x * x}


func Sum (nums[]int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
} 
