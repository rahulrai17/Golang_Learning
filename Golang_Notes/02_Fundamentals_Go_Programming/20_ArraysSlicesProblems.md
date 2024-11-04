## Coding problems for arrays and Slices

Here's a list of coding problems that range from beginner to expert, specifically chosen to deepen your understanding of arrays and slices. Each problem includes a solution approach and explanation.

---

### Beginner Problems

#### 1. **Reverse an Array**

**Problem**: Given an array, reverse the order of elements in place.

- **Input**: `nums = [1, 2, 3, 4, 5]`
- **Output**: `[5, 4, 3, 2, 1]`

**Solution**:

```go
func reverse(nums []int) {
    left, right := 0, len(nums)-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
```

**Explanation**: Swap elements from the start and end moving toward the center.

---

#### 2. **Remove Duplicates from a Sorted Array**

**Problem**: Given a sorted array, remove duplicates in-place and return the new length of the array.

- **Input**: `nums = [1, 1, 2, 3, 3, 4]`
- **Output**: New length is `4`, modified `nums` is `[1, 2, 3, 4, ...]`

**Solution**:

```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    j := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[j] {
            j++
            nums[j] = nums[i]
        }
    }
    return j + 1
}
```

**Explanation**: Use two pointers to overwrite duplicate values and return the count of unique elements.

---

### Intermediate Problems

#### 3. **Rotate Array by K Steps**

**Problem**: Rotate an array to the right by `k` steps.

- **Input**: `nums = [1, 2, 3, 4, 5, 6, 7], k = 3`
- **Output**: `[5, 6, 7, 1, 2, 3, 4]`

**Solution**:

```go
func rotate(nums []int, k int) {
    k %= len(nums)
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

func reverse(nums []int) {
    for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
        nums[left], nums[right] = nums[right], nums[left]
    }
}
```

**Explanation**: Rotate the array by reversing it, then reverse the first `k` elements, and finally reverse the rest.

---

#### 4. **Two Sum**

**Problem**: Find two numbers in an array that add up to a target sum. Return their indices.

- **Input**: `nums = [2, 7, 11, 15], target = 9`
- **Output**: `[0, 1]`

**Solution**:

```go
func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, num := range nums {
        if j, found := seen[target-num]; found {
            return []int{j, i}
        }
        seen[num] = i
    }
    return nil
}
```

**Explanation**: Use a map to track the indices of numbers. Check if the complement exists in the map.

---

### Advanced Problems

#### 5. **Product of Array Except Self**

**Problem**: Return an array where each element is the product of all elements in the input array except itself.

- **Input**: `nums = [1, 2, 3, 4]`
- **Output**: `[24, 12, 8, 6]`

**Solution**:

```go
func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    prefix := 1
    for i := 0; i < n; i++ {
        result[i] = prefix
        prefix *= nums[i]
    }
    suffix := 1
    for i := n-1; i >= 0; i-- {
        result[i] *= suffix
        suffix *= nums[i]
    }
    return result
}
```

**Explanation**: Compute prefix products for each element, then compute suffix products in reverse to avoid division.

---

#### 6. **Find All Subarrays with a Given Sum**

**Problem**: Find all subarrays in an array that sum up to a given target.

- **Input**: `nums = [1, 2, 3, -2, 5], target = 5`
- **Output**: `[[2, 3], [5]]`

**Solution**:

```go
func subarraysWithSum(nums []int, target int) [][]int {
    var result [][]int
    for start := 0; start < len(nums); start++ {
        sum := 0
        for end := start; end < len(nums); end++ {
            sum += nums[end]
            if sum == target {
                result = append(result, nums[start:end+1])
            }
        }
    }
    return result
}
```

**Explanation**: Use two nested loops to calculate sums of subarrays starting from each index.

---

### Expert Problems

#### 7. **Maximum Subarray Sum (Kadane’s Algorithm)**

**Problem**: Find the contiguous subarray with the maximum sum.

- **Input**: `nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]`
- **Output**: `6` (subarray `[4, -1, 2, 1]`)

**Solution**:

```go
func maxSubArray(nums []int) int {
    maxSum := nums[0]
    currentSum := nums[0]
    for i := 1; i < len(nums); i++ {
        currentSum = max(nums[i], currentSum + nums[i])
        maxSum = max(maxSum, currentSum)
    }
    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**Explanation**: Track the maximum sum of subarrays ending at each position.

---

#### 8. **Longest Substring with At Most K Distinct Characters**

**Problem**: Find the length of the longest substring containing at most `k` distinct characters.

- **Input**: `s = "eceba", k = 2`
- **Output**: `3` (substring `"ece"`)

**Solution**:

```go
func lengthOfLongestSubstringKDistinct(s string, k int) int {
    if k == 0 {
        return 0
    }
    charCount := make(map[byte]int)
    left, maxLength := 0, 0
    for right := 0; right < len(s); right++ {
        charCount[s[right]]++
        for len(charCount) > k {
            charCount[s[left]]--
            if charCount[s[left]] == 0 {
                delete(charCount, s[left])
            }
            left++
        }
        maxLength = max(maxLength, right - left + 1)
    }
    return maxLength
}
```

**Explanation**: Use a sliding window and a map to keep track of character counts within the window.

---

These problems and solutions provide a solid foundation in handling arrays and slices across different levels of complexity in Go, helping to build skills from basic operations to efficient, optimized solutions for more advanced algorithms.
