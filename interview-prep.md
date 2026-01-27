# Interview Prep - LeetCode Problems by Pattern

**Interview Date**: January 27, 2026  
**Goal**: Brush up on 11 core patterns with strategic problem solving

---

## PHASE 1: CRITICAL FOUNDATION (Easy - 10-15 min each)
**Goal**: Build confidence with basics | **Recommended Time**: 1 hour

- [ ] **LeetCode 1** - Two Sum
  - Pattern: Hash Maps & Sets
  - Link: https://leetcode.com/problems/two-sum/
  - Key Concept: Use a hash map to store seen values for O(1) lookup
  
- [ ] **LeetCode 704** - Binary Search
  - Pattern: Binary Search
  - Link: https://leetcode.com/problems/binary-search/
  - Key Concept: Classic binary search template
  
- [ ] **LeetCode 125** - Valid Palindrome
  - Pattern: Two Pointers
  - Link: https://leetcode.com/problems/valid-palindrome/
  - Key Concept: Two pointer approach from both ends
  
- [ ] **LeetCode 104** - Maximum Depth of Binary Tree
  - Pattern: DFS & BFS
  - Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/
  - Key Concept: Simple recursive DFS traversal

---

## PHASE 2: HIGH PRIORITY (Medium - 20-30 min each)
**Goal**: Master patterns most likely to appear | **Recommended Time**: 1.5 hours

- [ ] **LeetCode 3** - Longest Substring Without Repeating Characters
  - Pattern: Sliding Window
  - Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
  - Key Concept: Variable window with hash map for character tracking
  
- [ ] **LeetCode 424** - Longest Repeating Character Replacement
  - Pattern: Sliding Window
  - Link: https://leetcode.com/problems/longest-repeating-character-replacement/
  - Key Concept: Advanced variable window technique
  
- [ ] **LeetCode 238** - Product of Array Except Self
  - Pattern: Prefix Sum
  - Link: https://leetcode.com/problems/product-of-array-except-self/
  - Key Concept: Prefix and suffix product approach
  
- [ ] **LeetCode 200** - Number of Islands
  - Pattern: DFS & BFS
  - Link: https://leetcode.com/problems/number-of-islands/
  - Key Concept: Graph traversal to count connected components
  
- [ ] **LeetCode 33** - Search in Rotated Sorted Array
  - Pattern: Binary Search
  - Link: https://leetcode.com/problems/search-in-rotated-sorted-array/
  - Key Concept: Modified binary search with rotation
  
- [ ] **LeetCode 322** - Coin Change
  - Pattern: Dynamic Programming
  - Link: https://leetcode.com/problems/coin-change/
  - Key Concept: Classic DP problem - recurrence relation: dp[i] = min(dp[i - coin] + 1)
  
- [ ] **LeetCode 46** - Permutations
  - Pattern: Backtracking
  - Link: https://leetcode.com/problems/permutations/
  - Key Concept: Fundamental backtracking template with used array
  
- [ ] **LeetCode 739** - Daily Temperatures
  - Pattern: Monotonic Stack
  - Link: https://leetcode.com/problems/daily-temperatures/
  - Key Concept: Decreasing monotonic stack to find next greater element

---

## PHASE 3: INTERMEDIATE (Medium-Hard - 20-30 min each)
**Goal**: Deepen understanding | **Recommended Time**: 1 hour (if you have time)

- [ ] **LeetCode 11** - Container With Most Water
  - Pattern: Two Pointers
  - Link: https://leetcode.com/problems/container-with-most-water/
  - Key Concept: Greedy two-pointer approach
  
- [ ] **LeetCode 133** - Clone Graph
  - Pattern: DFS & BFS
  - Link: https://leetcode.com/problems/clone-graph/
  - Key Concept: Complex graph traversal with state management
  
- [ ] **LeetCode 131** - Palindrome Partitioning
  - Pattern: Backtracking
  - Link: https://leetcode.com/problems/palindrome-partitioning/
  - Key Concept: Backtracking with substring validation
  
- [ ] **LeetCode 208** - Implement Trie (Prefix Tree)
  - Pattern: Trie
  - Link: https://leetcode.com/problems/implement-trie-prefix-tree/
  - Key Concept: Essential trie structure and operations
  
- [ ] **LeetCode 39** - Combination Sum
  - Pattern: Backtracking
  - Link: https://leetcode.com/problems/combination-sum/
  - Key Concept: Backtracking with repetition allowed

---

## PHASE 4: BONUS (Hard - Only if extra time)
**Goal**: Polish | **Recommended Time**: 30 min

- [ ] **LeetCode 215** - Kth Largest Element in an Array
  - Pattern: Heaps
  - Link: https://leetcode.com/problems/kth-largest-element-in-an-array/
  - Key Concept: Min-heap of size K
  
- [ ] **LeetCode 84** - Largest Rectangle in Histogram
  - Pattern: Monotonic Stack
  - Link: https://leetcode.com/problems/largest-rectangle-in-histogram/
  - Key Concept: Complex monotonic stack application
  
- [ ] **LeetCode 72** - Edit Distance
  - Pattern: Dynamic Programming
  - Link: https://leetcode.com/problems/edit-distance/
  - Key Concept: 2D DP with recurrence relation

---

## SUGGESTED TIME BREAKDOWN FOR TOMORROW

| Time | Activity | Duration |
|------|----------|----------|
| Start | PHASE 1: Easy Problems | 1 hour |
| +1h | PHASE 2: Core Medium Problems | 1.5 hours |
| +2.5h | PHASE 3: Intermediate (if time permits) | 1 hour |
| +3.5h | Review Weak Areas | 30 min |
| +4h | PHASE 4: Bonus Problems (if extra time) | 30 min |

---

## QUICK REFERENCE CHECKLIST

### Before Coding Any Problem:
- [ ] Ask: Is the input sorted?
- [ ] Ask: What are the constraints? (array size, integer range)
- [ ] Clarify: What should I return on edge cases?
- [ ] Think: Is there a pattern I recognize?

### Pattern Identification Quick Tips:
- **Contiguous subarray/substring?** → Sliding Window
- **Sorted array + search?** → Binary Search
- **Need to find duplicates/exists?** → Hash Map/Set
- **Tree/Graph traversal?** → DFS/BFS
- **Optimal substructure?** → Dynamic Programming
- **Need all combinations/permutations?** → Backtracking
- **Find next greater/smaller?** → Monotonic Stack
- **Word prefix matching?** → Trie
- **Range sum queries?** → Prefix Sum
- **Find Kth largest?** → Heap

### Last-Minute Memory Aids:

**Two Pointers**: Sorted → two ends → move inward  
**Sliding Window**: Contiguous → expand/contract window  
**Hash Maps**: Fast lookup → O(1) access  
**DFS**: Deep first → recursion/stack → exhaustive  
**BFS**: Level by level → queue → shortest path  
**Binary Search**: Divide by 2 each time → O(log n)  
**DP**: Overlapping subproblems → memoize results  
**Backtracking**: Explore all → undo choices → prune  
**Monotonic Stack**: Maintain order → pop when violated  
**Trie**: Character by character → prefix sharing  
**Heaps**: Min/max access → O(1) extraction  

---

## NOTES SECTION
Use this space to write quick insights as you solve:

```
Problem: ________________
Pattern: ________________
Difficulty: Easy / Medium / Hard
Time Taken: ________________
Key Insight: ________________
Mistakes/Learnings: ________________
```

---

## FINAL REMINDERS FOR THE INTERVIEW

✅ **DO:**
- Read problem statement carefully (2-3 min)
- Ask clarifying questions
- Think out loud about approach
- Walk through examples
- Code cleanly with variable names that make sense
- Test with edge cases before submitting

❌ **DON'T:**
- Jump into coding immediately
- Over-optimize too early
- Forget to handle null/empty cases
- Write code you can't explain
- Panic if you get stuck on first attempt

**Remember**: Interviewers care more about your problem-solving approach than perfect code. Communication is key!

---

**Last Update**: January 26, 2026  
**Good Luck!** 🚀
