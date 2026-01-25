# LeetCode 75: Core Patterns & Principles (Golang Edition)

## 1. Two Pointers

**Core Idea**: Use two indices starting from different positions (usually start/end or slow/fast) to solve problems in a single pass.

**When to Use**:
- Arrays/strings that are **sorted** or **need pairing**
- Problems asking for "find a pair that satisfies condition"
- Reversing sequences, removing duplicates in-place
- Container with Most Water, Valid Palindrome

**Why It Works**: By moving pointers toward each other or at different speeds, you avoid nested loops. If the array is sorted, moving one pointer inward gives you information about which direction to continue.

**Key Insight**: The sorting allows you to make informed decisions—if looking for a target sum and current sum is too small, move the smaller pointer right (increase sum). If too large, move the larger pointer left (decrease sum).

**Common Pitfalls**: Forgetting to check sorted constraint, using when array order matters, off-by-one errors with pointer movement.

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

**Core Idea**: Maintain a window of fixed or variable size that "slides" across the array/string. Instead of recalculating values for each window, update incrementally by removing the leftmost element and adding the rightmost element.

**When to Use**:
- Problems with **contiguous subarrays or substrings**
- "Maximum/minimum sum of subarray of size k"
- "Longest substring without repeating characters"
- "Find all anagrams in string"
- Requires solving in O(n) rather than O(n²)

**Why It Works**: Brute force would check every subarray (O(n²)). Sliding window avoids redundant computation by only removing one element and adding one element per iteration.

**Key Insight**: 
- **Fixed window**: Calculate first window, then slide by removing left and adding right
- **Variable window**: Use two pointers (left/right) to expand/contract based on a condition

**Algorithm Steps**:
1. Create window of appropriate size
2. Calculate result for first window
3. Slide: remove from left, add from right
4. Update max/min as you go

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

**Core Idea**: Pre-compute cumulative sums from the start of the array. This allows you to answer "sum of elements from index i to j" in O(1) time instead of recalculating.

**When to Use**:
- "Find sum of subarray from index i to j" (queried multiple times)
- "Pivot index" problems, range sum queries
- Problems involving "left sum equals right sum"
- 2D matrix sum problems

**Why It Works**: Without prefix sums, calculating range sum is O(n). With preprocessing, it becomes O(1): `sum[i:j] = prefix[j+1] - prefix[i]`

**Key Insight**: Build an array where `prefix[i]` = sum of all elements from index 0 to i-1. Then any range sum is just the difference of two prefix values.

**Example Logic**:
- `totalSum` = sum of entire array
- For each index, check if `leftSum == rightSum`
- `rightSum = totalSum - leftSum - current_element`

**Tradeoff**: Requires O(n) extra space for the prefix array (can be optimized for single-pass problems).

Time: O(n) preprocessing, O(1) per query | Space: O(n)

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

**Core Idea**: Use hash tables (maps/dicts) to store values for O(1) average-case lookup, or sets to track existence.

**When to Use**:
- "Check if element exists" (avoid O(n) search)
- "Count frequency of elements"
- "Find duplicates"
- "Group elements by some property"
- "Anagram/character mapping problems"
- "Two Sum" style problems (find if X exists)

**Why It Works**: Hash tables provide O(1) average-case access, whereas arrays require O(n) linear search. This trades memory for speed.

**Two Main Patterns**:
1. **Set for existence**: Track what you've seen to avoid duplicates or find missing numbers
2. **Map for counting**: Store frequency of elements, useful for finding duplicates, majorities, or validating constraints

**Key Insight**: When you need fast lookups or to avoid revisiting elements, hash tables are the default choice. They're fundamental for reducing O(n²) problems to O(n).

**Common Use Cases**:
- Valid anagram: count char frequencies and compare
- Contains duplicate: track seen numbers
- Two sum: for each num, check if (target - num) exists

**Pitfalls**: Hash collisions (rare in modern implementations), using when order matters, forgetting map lookup is O(1) average not guaranteed.

Time: O(n) | Space: O(n)

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

**Core Idea**: Systematically explore all nodes in a tree/graph. **DFS** goes deep first (using recursion or stack), **BFS** goes level by level (using queue).

**When to Use DFS**:
- Tree/graph traversal with recursion
- "All paths from root to leaf"
- "Detect cycles" in directed graphs
- "Number of connected components"
- Problems naturally solved by thinking recursively
- Need to explore all possibilities (backtracking)

**When to Use BFS**:
- "Shortest path" in unweighted graph
- "Level-order traversal" of tree
- "Minimum distance"
- Breadth-first exploration needed
- Problems requiring exploring one "level" before moving to next

**DFS Logic**:
1. Visit current node
2. Recursively visit all unvisited children
3. Backtrack and process children results

**BFS Logic**:
1. Start with initial node(s) in queue
2. Dequeue node, process it, enqueue unvisited neighbors
3. Repeat until queue empty

**Key Insight**: DFS naturally explores deeply (good for exhaustive search), BFS explores broadly (good for finding shortest paths). DFS uses implicit stack (recursion), BFS needs explicit queue.

**Space Complexity**:
- DFS: O(height) for balanced tree, O(n) for skewed tree (call stack)
- BFS: O(width) of tree, worst case O(n)

Time: O(V + E) for V vertices and E edges | Space: O(V)

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

**Core Idea**: A complete binary tree where each parent is smaller (min-heap) or larger (max-heap) than its children. Allows efficient extraction of the minimum/maximum element.

**When to Use**:
- "Find K-th largest/smallest element"
- "Merge K sorted lists"
- "Top K frequent elements"
- "Median of data stream"
- Need to repeatedly access the min/max efficiently
- Problems requiring dynamic min/max tracking

**Why It Works**: 
- Extracting min/max: O(1)
- Inserting element: O(log n)
- Better than sorting repeatedly (O(n log n)) when you only need top K

**Key Insight**: A **min-heap of size K** storing K largest elements keeps the smallest of the K largest at the root. To find K-th largest, build min-heap of top K, answer is the root. This avoids storing all n elements.

**Min-Heap vs Max-Heap**:
- **Min-heap**: Root is smallest. Use `Less(i,j) = h[i] < h[j]`
- **Max-heap**: Root is largest. Use `Less(i,j) = h[i] > h[j]` (flip comparison)

**Algorithm for K-th Largest**:
1. Add all elements to min-heap, keep size ≤ K
2. If heap size > K, remove minimum
3. Root of heap is K-th largest

**Complexity**:
- Building heap: O(n log k) by inserting n elements with max size k
- Space: O(k) for storing only K elements

Time: O(n log k) | Space: O(k)

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

**Core Idea**: Maintain a stack that stores indices (or values) in strictly increasing or decreasing order. When you encounter an element that breaks the order, pop elements and record the relationship.

**When to Use**:
- "Next greater/smaller element"
- "Daily temperatures" style problems
- "Largest rectangle in histogram"
- "Trapping rain water"
- Problems where you need to find closest greater/smaller element to the right/left

**Why It Works**: Naive approach checks each element against all elements to the right (O(n²)). Monotonic stack efficiently tracks candidates in one pass.

**Key Insight**: As you iterate left-to-right, the stack represents potential "next greater" candidates. When you find an element greater than stack top, you've found the next greater element for all popped elements.

**Algorithm Pattern**:
1. For each new element:
   - While stack not empty AND current element > stack.top():
     - Pop from stack (you've found their next greater)
     - Record the relationship
   - Push current element to stack
2. Remaining stack elements have no greater element to their right

**Increasing vs Decreasing Stack**:
- **Increasing stack**: Maintains indices in increasing order of their values. Use for "next greater element"
- **Decreasing stack**: Maintains indices in decreasing order. Use for "previous greater element"

**Why Store Indices?**: Storing indices (not just values) lets you calculate the distance/gap to the next greater element.

**Space Insight**: Stack never grows beyond n, and each element is pushed/popped once, making it O(n) time total despite nested loops.

Time: O(n) | Space: O(n)

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

**Core Idea**: Eliminate half of the remaining search space with each comparison by checking the middle element. Only works on **sorted data**.

**When to Use**:
- Searching in **sorted array**
- Finding "first/last occurrence"
- Finding "peak element"
- Problems asking for O(log n) complexity
- Search space that can be divided into left/right halves
- "Rotated sorted array"

**Why It Works**: Each comparison eliminates half the candidates. Reduces O(n) linear search to O(log n).

**Key Insight**: You're not just searching for exact matches. Binary search on **answer space**: can I achieve this value? If yes, try higher; if no, try lower. This pattern solves many problems not obviously about searching.

**Algorithm**:
1. Set `low` to start, `high` to end
2. While `low <= high`:
   - Calculate `mid = low + (high - low) / 2` (avoids overflow)
   - If `nums[mid] == target`, return mid
   - If `nums[mid] < target`, move right: `low = mid + 1`
   - Else, move left: `high = mid - 1`
3. If not found, return -1

**Critical Details**:
- Use `low + (high - low) / 2` instead of `(low + high) / 2` to prevent integer overflow
- Understand `low <= high` vs `low < high` for different problem types
- For first occurrence: when found, continue searching left
- For last occurrence: when found, continue searching right

**Common Pitfall**: Not verifying array is sorted, or forgetting to handle edge cases (single element, target at boundaries).

Time: O(log n) | Space: O(1)

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

**Core Idea**: Break problem into overlapping subproblems. Solve each subproblem once and store the result (memoization). Reuse results instead of recalculating.

**When to Use**:
- Problems with **overlapping subproblems** (naive recursion recalculates same subproblems)
- Problems with **optimal substructure** (optimal solution depends on optimal solutions to subproblems)
- "Climb stairs" (ways to reach step n = ways to reach step n-1 + ways to reach step n-2)
- "Coin change", "Knapsack", "Longest increasing subsequence"
- Problems asking for "max/min ways/count"

**Why It Works**: Without DP, recursive solution recomputes same values exponentially. Storing results reduces exponential time to polynomial.

**Two Approaches**:

1. **Top-Down (Memoization)**:
   - Start from the problem and recursively break it down
   - Cache results as you compute them
   - Natural but uses recursion stack

2. **Bottom-Up (Tabulation)**:
   - Start from base cases
   - Build up to the final answer iteratively
   - More efficient, avoids recursion overhead

**Key Steps to Solve DP Problem**:
1. Define what `dp[i]` means (e.g., "ways to reach step i")
2. Find **recurrence relation**: How does `dp[i]` depend on previous states?
3. Identify **base cases**: `dp[0]`, `dp[1]`, etc.
4. Build solution from base cases upward (or use memoization from top down)

**Example Analysis (Climb Stairs)**:
- To reach step n, you either came from step n-1 or n-2
- So `dp[n] = dp[n-1] + dp[n-2]`
- Base cases: 1 way to reach step 1, 2 ways to reach step 2

**Space Optimization**: Many DP solutions use O(n) space, but can be optimized to O(1) if you only need recent states (like climb stairs only needs last two values).

Time: O(n) | Space: O(n) or O(1) with optimization

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

**Core Idea**: Recursively explore all possible solutions by making a choice, recursing into subproblems, then undoing the choice (backtracking) to try other options.

**When to Use**:
- "Find all permutations/combinations"
- "Solve N-Queens problem"
- "Word search in grid"
- "Sudoku solver"
- Problems with multiple solution paths that need exhaustive exploration
- "Generate all subsets"
- Constraint satisfaction problems

**Why It Works**: Some problems require exploring all possibilities. Backtracking avoids exponential blowup by intelligently pruning branches that can't lead to valid solutions.

**Key Insight**: The **undo step** is critical. After exploring a branch, you must restore the state so that other branches start from a clean slate. This is what distinguishes backtracking from simple recursion.

**Algorithm Pattern**:
```
backtrack(path, choices):
  if goal_reached(path):
    add path to results
    return
  
  for each choice in choices:
    if is_valid(choice):
      make_choice(choice)  // Add to path
      backtrack(path, remaining_choices)
      undo_choice(choice)  // Remove from path
```

**Template Structure**:
1. **Base case**: If path is complete and valid, add to results
2. **Recursive case**: Try each possible next choice
3. **Backtrack**: Undo the choice before trying the next one

**Permutations vs Combinations**:
- **Permutations**: All orderings matter. Mark elements as used to ensure each used exactly once
- **Combinations**: Only selection matters, not order. Only proceed with remaining choices

**Pruning**: The true power of backtracking comes from early termination. If partial solution violates constraint, don't continue exploring that branch.

**Example (Permutations)**:
- `used` array tracks which numbers already in current permutation
- When permutation size equals n, we have a complete solution
- Undo (`used[i] = false`) lets us try different number at same position

**Complexity**: Usually exponential O(N! for permutations, 2^N for subsets) because you explore all possibilities, but pruning can significantly reduce actual runtime.

Time: O(N! × N) for permutations | Space: O(N) for recursion depth + path storage

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

## 11. Trie (Prefix Tree)

**Core Idea**: A tree-based data structure where each node represents a character. Paths from root to nodes spell out strings. Efficient for problems involving string prefixes and exact word matching.

**When to Use**:
- "Autocomplete" functionality
- "Spell checker"
- "Word search in grid"
- "Longest common prefix"
- "Implement dictionary"
- Problems involving **prefix matching** or **dictionary lookups**
- "Replace Words" (find if word starts with any dictionary word)
- "Stream of characters" validation

**Why It Works**: Unlike hash maps which require full string lookup, Tries leverage **prefix sharing**. Many words share common beginnings, so storing them in a Trie saves space and enables fast prefix searches.

**Key Structure**:
- **TrieNode**: Contains a map/array of children (one per character) and a boolean flag indicating if it's an end-of-word
- **Root node**: Empty, serves as entry point
- **Each path from root to leaf**: Represents a stored string

**Key Insight**: The power of Tries is that all words with the same prefix share the same path down the tree. This enables:
- O(L) search where L is word length (not O(n) like hash map)
- Prefix-based operations without scanning entire dictionary
- Space efficiency when many words share prefixes

**Advantages Over Hash Maps**:
- Hash maps: O(1) lookup for exact word, but O(n) for prefix matching
- Tries: O(L) for both exact match and prefix matching
- Better for autocomplete and prefix-based problems
- Trade-off: More space per character (26 pointers for lowercase letters)

**Common Operations**:
1. **Insert**: Traverse/create path for each character, mark last node as end-of-word
2. **Search**: Traverse path, return true only if we complete all characters AND last node is marked as end-of-word
3. **StartsWith**: Traverse path, return true if we complete all characters (don't care about end-of-word flag)
4. **Delete**: Remove node if it has no children and isn't end-of-word (recurse back)

**Space Complexity**: O(ALPHABET_SIZE × N × M) where N is number of words, M is average word length, ALPHABET_SIZE is 26 for lowercase letters. With prefix sharing, actual space is often much less than storing all strings separately.

Time: O(L) for insert/search/startsWith where L is word length | Space: O(ALPHABET_SIZE × N × M)

```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEndOfWord bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if _, exists := node.children[ch]; !exists {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if n, exists := node.children[ch]; exists {
            node = n
        } else {
            return false
        }
    }
    return node.isEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        if n, exists := node.children[ch]; exists {
            node = n
        } else {
            return false
        }
    }
    return true
}
```