package main

import "fmt"

// ==================== DSA SORT HELPERS ====================

func generateQuickSortDynamic(listName, dataType string) string {
	return fmt.Sprintf(`# Quick Sort — your %s
# Time: O(n log n) average, O(n²) worst
# Space: O(log n) recursion stack

def quick_sort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quick_sort(left) + middle + quick_sort(right)

%s = quick_sort(%s)
print(%s)`, dataType, listName, listName, listName)
}

func generateMergeSortDynamic(listName, dataType string) string {
	return fmt.Sprintf(`# Merge Sort — your %s
# Time: O(n log n) always (guaranteed)
# Space: O(n) — needs extra arrays
# Stable: Yes

def merge_sort(arr):
    if len(arr) <= 1:
        return arr
    mid = len(arr) // 2
    left = merge_sort(arr[:mid])
    right = merge_sort(arr[mid:])
    return merge(left, right)

def merge(left, right):
    result = []
    i = j = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result

%s = merge_sort(%s)
print(%s)`, dataType, listName, listName, listName)
}

func generateBubbleSortDynamic(listName, dataType string) string {
	return fmt.Sprintf(`# Bubble Sort — your %s
# Time: O(n²) — SLOW for large lists! Use only for learning.
# Space: O(1) — in-place

def bubble_sort(arr):
    n = len(arr)
    for i in range(n):
        swapped = False
        for j in range(0, n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                swapped = True
        if not swapped:
            break
    return arr

%s = bubble_sort(%s)
print(%s)

# ⚠️ For production, use: %s.sort() — Python's built-in is MUCH faster`, dataType, listName, listName, listName, listName)
}

// ==================== DSA PATTERN HELPERS ====================

func generateBinarySearch() string {
	return `def binary_search(arr, target):
    """
    Binary Search (iterative)
    Time:  O(log n)
    Space: O(1)
    Requires: sorted array
    """
    low, high = 0, len(arr) - 1

    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1

    return -1  # not found`
}

func generateTwoPointer() string {
	return `def two_pointer_pair(arr, target):
    """
    Two Pointer — find pair summing to target in sorted array
    Time:  O(n)
    Space: O(1)
    """
    left, right = 0, len(arr) - 1

    while left < right:
        current = arr[left] + arr[right]
        if current == target:
            return (left, right)
        elif current < target:
            left += 1
        else:
            right -= 1

    return None`
}

func generateSlidingWindow() string {
	return `def max_sum_window(arr, k):
    """
    Sliding Window — max sum subarray of size k
    Time:  O(n)
    Space: O(1)
    """
    window = sum(arr[:k])
    best = window

    for i in range(k, len(arr)):
        window += arr[i] - arr[i - k]
        best = max(best, window)

    return best`
}

func generateHashMapLookup() string {
	return `def find_duplicates(arr):
    """
    Hash Map — find duplicates in O(n)
    Time:  O(n)
    Space: O(n)
    """
    seen = set()
    duplicates = []

    for item in arr:
        if item in seen:
            duplicates.append(item)
        seen.add(item)

    return duplicates`
}

func generateTwoSum() string {
	return `def two_sum(nums, target):
    """
    Two Sum — find indices summing to target
    Time:  O(n)
    Space: O(n)
    """
    seen = {}
    for i, num in enumerate(nums):
        complement = target - num
        if complement in seen:
            return [seen[complement], i]
        seen[num] = i
    return []`
}

func generateReverseLinkedList() string {
	return `class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def reverse_linked_list(head):
    """
    Reverse linked list in-place
    Time:  O(n)
    Space: O(1)
    """
    prev = None
    current = head

    while current:
        next_node = current.next
        current.next = prev
        prev = current
        current = next_node

    return prev`
}

func generateBFS() string {
	return `from collections import deque

def bfs(graph, start):
    """
    Breadth-First Search — level by level
    Time:  O(V + E)
    Space: O(V)
    Use: shortest path (unweighted), level-order
    """
    visited = set([start])
    queue = deque([start])
    result = []

    while queue:
        node = queue.popleft()
        result.append(node)

        for neighbor in graph[node]:
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append(neighbor)

    return result

# Usage:
# graph = {'A': ['B', 'C'], 'B': ['D'], 'C': ['E'], 'D': [], 'E': []}
# bfs(graph, 'A')`
}

func generateDFS() string {
	return `def dfs(graph, start, visited=None):
    """
    Depth-First Search — go deep first
    Time:  O(V + E)
    Space: O(V)
    Use: cycle detection, path finding, topological sort
    """
    if visited is None:
        visited = set()

    visited.add(start)
    result = [start]

    for neighbor in graph[start]:
        if neighbor not in visited:
            result.extend(dfs(graph, neighbor, visited))

    return result

# Iterative (no stack overflow for large graphs):
def dfs_iterative(graph, start):
    visited = set()
    stack = [start]
    result = []

    while stack:
        node = stack.pop()
        if node not in visited:
            visited.add(node)
            result.append(node)
            stack.extend(reversed(graph[node]))

    return result`
}

func generateDPFibonacci() string {
	return `def fibonacci(n):
    """
    Fibonacci — optimized DP
    Time:  O(n)
    Space: O(1) — only stores last two values!
    """
    if n <= 1:
        return n

    prev2, prev1 = 0, 1
    for _ in range(2, n + 1):
        current = prev1 + prev2
        prev2 = prev1
        prev1 = current

    return prev1

# fibonacci(50) = 12586269025 — instant!`
}

func generateValidParentheses() string {
	return `def is_valid(s):
    """
    Valid Parentheses — check balanced brackets
    Time:  O(n)
    Space: O(n)
    """
    stack = []
    mapping = {')': '(', '}': '{', ']': '['}

    for char in s:
        if char in mapping:
            if not stack or stack[-1] != mapping[char]:
                return False
            stack.pop()
        else:
            stack.append(char)

    return len(stack) == 0

# is_valid("([{}])")  → True
# is_valid("(]")      → False`
}

func generateCoinChange() string {
	return `def coin_change(coins, amount):
    """
    Coin Change — minimum coins to make amount
    Time:  O(amount * len(coins))
    Space: O(amount)
    LeetCode #322
    """
    dp = [float('inf')] * (amount + 1)
    dp[0] = 0

    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i:
                dp[i] = min(dp[i], dp[i - coin] + 1)

    return dp[amount] if dp[amount] != float('inf') else -1

# Usage:
# coin_change([1, 5, 10, 25], 30)  → 2 (25 + 5)`
}

func generateLCS() string {
	return `def longest_common_subsequence(text1, text2):
    """
    Longest Common Subsequence
    Time:  O(m * n)
    Space: O(m * n), can be optimized to O(min(m, n))
    LeetCode #1143
    """
    m, n = len(text1), len(text2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if text1[i - 1] == text2[j - 1]:
                dp[i][j] = dp[i - 1][j - 1] + 1
            else:
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

    return dp[m][n]

# Usage:
# longest_common_subsequence("abcde", "ace")  → 3 ("ace")`
}
