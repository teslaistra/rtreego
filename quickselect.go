package rtreego

//implementation by https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/205552/C%2B%2B-quick-select-O(n)-run-time

func partitionInt(arr []int64, l int64, r int64) int64 {
	x := arr[r]
	i := l

	for j := l; j <= r-1; j++ {
		if arr[j] <= x {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}
func Kth_smallestInt(arr []int64, l, r, k int64) int64 {
	if k < 0 || k > r-l+1 {
		return -1
	}

	//partition the array (similar to quicksort)
	index := partitionInt(arr, l, r)

	//if position same as k return rank-k element
	if index-l == k-1 {
		return arr[index]
	}

	//if rank-k element is less, look in left sub-array
	if index-l > k-1 {
		return Kth_smallestInt(arr, l, index-1, k)
	}
	return Kth_smallestInt(arr, index+1, r, k-index+l-1)
}
func GetNSmallestInt(nums []int64, n int64) []int64 {
	var arr []int64
	minMax := Kth_smallestInt(nums, 0, int64(len(nums)-1), n)
	for i := range nums {
		if nums[i] <= minMax {
			arr = append(arr, nums[i])

		}
		i++
	}
	return arr
}
func FindKthSmallest(nums []int64, k int64) int64 {

	n := len(nums)
	arr := make([]int64, n)

	for i := 0; i < n; i = i {
		arr[i] = nums[i]
		i++
	}
	result := Kth_smallestInt(arr, 0, int64(n-1), int64(n)-k+1)
	return result
}

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
func KthSmallest(arr []SearchObject, l, r, k int64) SearchObject {
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
		return KthSmallest(arr, l, index-1, k)
	}
	return KthSmallest(arr, index+1, r, k-index+l-1)
}
func GetNSmallest(nums []SearchObject, n int64) []Spatial {
	var arr []Spatial
	kMin := KthSmallest(nums, 0, int64(len(nums)-1), n)
	for i := range nums {
		if nums[i].Distance <= kMin.Distance {
			arr = append(arr, nums[i].Object)
		}
		i++
	}
	return arr
}
