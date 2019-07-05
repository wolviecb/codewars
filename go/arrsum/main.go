package arrsum

//Shorter uglier
// func TwoSum(numbers []int, target int) [2]int {
// 	for i := range numbers {
// 		for j := range numbers {
// 			if i != j {
// 				if s := numbers[i] + numbers[j]; s == target {
// 					return [2]int{i, j}
// 				}
// 			}
// 		}
// 	}
// 	return [2]int{}
// }

//TwoSum sums
func TwoSum(numbers []int, target int) [2]int {
	r := [2]int{}
	s := 0
	for s != target {
		for i := range numbers {
			for j := range numbers {
				if i != j {
					if s = numbers[i] + numbers[j]; s == target {
						r = [2]int{i, j}
						break
					}
				}
			}
			if s == target {
				break
			}
		}
	}
	return r
}
