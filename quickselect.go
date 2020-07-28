package rtreego

/*
// Partition using Lomuto partition scheme
func partition(a []int64, left, right, pivotIndex int64) int64 {
	// Pick pivotIndex as pivot from the array
	pivot := a[pivotIndex]

	// Move pivot to end
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// elements less than pivot will be pushed to the left of pIndex
	// elements more than pivot will be pushed to the right of pIndex
	// equal elements can go either way
	pIndex := left

	// each time we finds an element less than or equal to pivot, pIndex
	// is incremented and that element would be placed before the pivot.
	for i := left; i < right; i++ {
		if a[i] <= pivot {
			a[i], a[pIndex] = a[pIndex], a[i]
			pIndex++
		}
	}

	// Move pivot to its final place
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// return pIndex (index of pivot element)
	return pIndex
}

// Returns the k-th smallest element of list within left..right
// (i.e. left <= k <= right). The search space within the array is
// changing for each round - but the list is still the same size.
// Thus, k does not need to be updated with each round.
func Quickselect(A []int64, left, right, k int64) int64 {
	// If the array contains only one element, return that element
	if left == right {
		return A[left]
	}

	// select a pivotIndex between left and right
	pivotIndex := left + (right - left + 1)

	pivotIndex = partition(A, left, right, pivotIndex)

	// The pivot is in its final sorted position
	if k == pivotIndex {
		return A[k]
	}

	// if k is less than the pivot index
	if k < pivotIndex {
		return Quickselect(A, left, pivotIndex-1, k)
	}

	// if k is more than the pivot index else
	return Quickselect(A, pivotIndex+1, right, k)
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
*/

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

//implementation by https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/205552/C%2B%2B-quick-select-O(n)-run-time
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
func Kth_smallest(arr []SearchObject, l, r, k int64) SearchObject {
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
		return Kth_smallest(arr, l, index-1, k)
	}
	return Kth_smallest(arr, index+1, r, k-index+l-1)
}
func GetNSmallest(nums []SearchObject, n int64) []Spatial {
	var arr []Spatial
	kMin := Kth_smallest(nums, 0, int64(len(nums)-1), n)
	for i := range nums {
		if nums[i].Distance <= kMin.Distance {
			arr = append(arr, nums[i].Object)
		}
		i++
	}
	return arr
}
