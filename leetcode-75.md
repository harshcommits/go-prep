# LeetCode 75: Core Patterns & Principles (Golang Edition)

## 1. Two Pointers: 

Used for searching pairs or reversing data in linear time.
Time: O(n) | Space: O(1)

```go
func reverseString(s []byte) {
    left, right := 0, len(s)-1
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}
```

## 2. Sliding Window

Ideal for contiguous subarrays or substrings.
Time: O(n) | Space: O(1) or O(k)

```go
func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i := 0; i < k; i++ { sum += nums[i] }
    maxSum := sum
    for i := k; i < len(nums); i++ {
        sum += nums[i] - nums[i-k] // Slide the window
        if sum > maxSum { maxSum = sum }
    }
    return float64(maxSum) / float64(k)
}
```

## 3. Prefix Sum
Pre-calculating cumulative totals to answer range queries quickly.

```go
func pivotIndex(nums []int) int {
    totalSum, leftSum := 0, 0
    for _, x := range nums { totalSum += x }
    for i, x := range nums {
        if leftSum == totalSum - leftSum - x { return i }
        leftSum += x
    }
    return -1
}
```

## 4. Hash Maps & Sets
Used for O(1) lookups and frequency counting.

```go
func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    for _, num := range nums {
        if seen[num] { return true }
        seen[num] = true
    }
    return false
}
```

## 5. DFS & BFS (Trees/Graphs)
Traversing nodes depth-first (recursion/stack) or breadth-first (queue).

```go
// DFS Example (Max Depth of Binary Tree)
func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    left := maxDepth(root.Left)
    right := maxDepth(root.Right)
    if left > right { return left + 1 }
    return right + 1
}
```

## 6. Heaps (Priority Queues)
Finding the K-th largest/smallest element efficiently.Time: O(n \log k) | Space: O(k)

```go
// Using container/heap for a Min-Heap
type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

## 7. Monotonic Stack
Maintaining a stack in increasing or decreasing order to find the "next greater element."

```go
func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stack := []int{} // stores indices
    for i, t := range temperatures {
        for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
            idx := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res[idx] = i - idx
        }
        stack = append(stack, i)
    }
    return res
}
```

## 8. Binary Search
Efficiently searching a sorted space.
Time: O(\log n)

```go
func search(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        mid := low + (high-low)/2
        if nums[mid] == target { return mid }
        if nums[mid] < target { low = mid + 1 } else { high = mid - 1 }
    }
    return -1
}
```

## 9. Dynamic Programming (DP)
Solving sub-problems and storing results to avoid redundant work.

```go
func climbStairs(n int) int {
    if n <= 2 { return n }
    dp := make([]int, n+1)
    dp[1], dp[2] = 1, 2
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```

## 10. Backtracking
Systematically exploring all possible configurations (e.g., permutations).Go

```go
func backtrack(res *[][]int, temp []int, nums []int, used []bool) {
    if len(temp) == len(nums) {
        c := make([]int, len(temp))
        copy(c, temp)
        *res = append(*res, c)
        return
    }
    for i := 0; i < len(nums); i++ {
        if used[i] { continue }
        used[i] = true
        backtrack(res, append(temp, nums[i]), nums, used)
        used[i] = false // undo the choice
    }
}
```