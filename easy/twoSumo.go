package easy

// import "fmt"

func TwoSum(nums []int, target int) []int {
	var result []int // ini slice
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j {
				if nums[i]+nums[j] == target {
					result = append(result,i)
					result = append(result,j)
				}
			}
		}
		if len(result)==2{
			break
		}
	}
	return result
}