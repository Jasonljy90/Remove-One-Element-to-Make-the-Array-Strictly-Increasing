package main

func canBeIncreasing(nums []int) bool {
	// Stores the count of numbers that
	// are needed to be removed
	count := 0

	// Store the index of the element
	// that needs to be removed
	index := -1

	// Traverse the range [1, N - 1]
	for i := 1; i < len(nums); i++ {

		// If arr[i-1] is greater than
		// or equal to arr[i]
		if nums[i-1] >= nums[i] {

			// Increment the count by 1
			count++

			// Update index
			index = i
		}
	}

	// If count is greater than one
	if count > 1 {
		return false
	}

	// If no element is removed
	if count == 0 {
		return true
	}

	// If only the last or the
	// first element is removed
	if index == len(nums)-1 || index == 1 {
		return true
	}

	// If a[index] is removed
	if nums[index-1] < nums[index+1] {
		return true
	}

	// If a[index - 1] is removed
	if nums[index-2] < nums[index] {
		return true
	}

	return false
}
