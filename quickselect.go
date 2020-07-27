package rtreego

import "fmt"

func partition(arr []SearchObject, l int64, r int64) int64 {
	x := arr[r].Distance
	i := l

	for j := l; j <= r-1; j++ {
		if arr[j].Distance <= x {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}

func kth_smallest(arr []SearchObject, l, r, k int64) SearchObject {
	if k < 0 || k > r-l+1 {
		return SearchObject{}
	}

	//partition the array (similar to quicksort)
	index := partition(arr, l, r)

	//if position same as k return rank-k element
	if index-l == k-1 {
		return arr[index]
	}

	//if rank-k element is less, look in left sub-array
	if index-l > k-1 {
		return kth_smallest(arr, l, index-1, k)
	}
	return kth_smallest(arr, index+1, r, k-index+l-1)
}

//implementation by https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/205552/C%2B%2B-quick-select-O(n)-run-time
func FindKthSmallest(nums []SearchObject, k int64) SearchObject {

	n := len(nums)
	arr := make([]SearchObject, n)

	for i := 0; i < n; i++ {
		arr[i] = nums[i]
	}
	result := kth_smallest(arr, 0, int64(n-1), int64(int64(n)-k+1))
	return result
}

func GetNSmallest(nums []SearchObject, n int64) []Spatial {
	fmt.Println(len(nums))
	var arr []Spatial
	minMax := FindKthSmallest(nums, n)
	for i := range nums {
		if nums[i].Distance <= minMax.Distance {
			arr = append(arr, nums[i].Object)

		}
		i++
	}
	return arr
}
