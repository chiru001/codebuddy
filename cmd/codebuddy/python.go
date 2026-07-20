package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var outputFile string

func newPythonCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "python",
		Short: "🐍 Python code generator",
		Long:  `Generate production-ready Python code from your requirements.`,
	}

	cmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Save generated code to a file")

	cmd.AddCommand(
		newForLoopCmd(),
		newFunctionCmd(),
		newClassCmd(),
		newOOPCmd(),
		newGeneratorCmd(),
		newErrorHandlingCmd(),
		newFileIOCmd(),
		newAPICmd(),
		newDSACmd(),
	)

	return cmd
}

// ==================== FOR LOOP ====================

func newForLoopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "for-loop",
		Short: "Generate a for loop based on your needs",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🐍 Python For Loop Generator")
			fmt.Println("─────────────────────────────")

			color.Cyan("\n? What are you looping over?")
			fmt.Println("  [1] A list")
			fmt.Println("  [2] A range of numbers")
			fmt.Println("  [3] A dictionary")
			fmt.Println("  [4] A file")
			fmt.Println("  [5] Nested loop (2D)")
			fmt.Println("  [6] List comprehension (one-liner)")
			fmt.Print("\n  Select: ")

			choice := readInputWithValidation(reader, []string{"1", "2", "3", "4", "5", "6"})

			var code string

			switch choice {
			case "1":
				code = dynamicListLoop(reader)
			case "2":
				code = dynamicRangeLoop(reader)
			case "3":
				code = dynamicDictLoop(reader)
			case "4":
				code = dynamicFileLoop(reader)
			case "5":
				code = dynamicNestedLoop(reader)
			case "6":
				code = dynamicListComprehension(reader)
			}

			printGenerated(code)
		},
	}
}

// ==================== FUNCTION ====================

func newFunctionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "function",
		Short: "Generate a function based on your needs",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🐍 Python Function Generator")
			fmt.Println("─────────────────────────────")

			color.Cyan("\n? Function name: ")
			funcName := readInput(reader)

			color.Cyan("? What does this function do (in plain English): ")
			description := readInput(reader)

			color.Cyan("? Parameters (comma-separated, e.g., name, age, salary): ")
			params := readInput(reader)

			color.Cyan("? Parameter types (comma-separated, e.g., str, int, float): ")
			types := readInput(reader)

			color.Cyan("? Should it return something?")
			fmt.Println("  [1] Yes — return a value")
			fmt.Println("  [2] No — just perform an action (print, save, etc.)")
			fmt.Print("\n  Select: ")
			returnChoice := readInputWithValidation(reader, []string{"1", "2"})

			var returnType string
			var returnStatement string
			if returnChoice == "1" {
				color.Cyan("? What should it return (e.g., result, filtered_list, total): ")
				returnVal := readInput(reader)
				color.Cyan("? Return type (e.g., int, str, list, dict, bool): ")
				returnType = readInput(reader)
				returnStatement = fmt.Sprintf("    return %s", returnVal)
			} else {
				returnType = "None"
				returnStatement = "    pass  # TODO: implement action"
			}

			color.Cyan("? Add error handling?")
			fmt.Println("  [1] Yes")
			fmt.Println("  [2] No")
			fmt.Print("\n  Select: ")
			errorChoice := readInputWithValidation(reader, []string{"1", "2"})

			// Build the function
			paramList := strings.Split(params, ",")
			typeList := strings.Split(types, ",")
			var typedParams []string
			var docParams []string
			for i, p := range paramList {
				p = strings.TrimSpace(p)
				t := "any"
				if i < len(typeList) {
					t = strings.TrimSpace(typeList[i])
				}
				typedParams = append(typedParams, fmt.Sprintf("%s: %s", p, t))
				docParams = append(docParams, fmt.Sprintf("        %s (%s): description", p, t))
			}

			var code string
			if errorChoice == "1" {
				code = fmt.Sprintf(`def %s(%s) -> %s:
    """
    %s

    Args:
%s

    Returns:
        %s: description
    """
    try:
%s
    except TypeError as e:
        print(f"Type error: {e}")
        raise
    except ValueError as e:
        print(f"Value error: {e}")
        raise
    except Exception as e:
        print(f"Unexpected error in %s: {e}")
        raise`,
					funcName,
					strings.Join(typedParams, ", "),
					returnType,
					description,
					strings.Join(docParams, "\n"),
					returnType,
					"    "+returnStatement,
					funcName,
				)
			} else {
				code = fmt.Sprintf(`def %s(%s) -> %s:
    """
    %s

    Args:
%s

    Returns:
        %s: description
    """
%s`,
					funcName,
					strings.Join(typedParams, ", "),
					returnType,
					description,
					strings.Join(docParams, "\n"),
					returnType,
					returnStatement,
				)
			}

			printGenerated(code)
		},
	}
}

// ==================== CLASS ====================

func newClassCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "class",
		Short: "Generate a class based on your needs",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🐍 Python Class Generator")
			fmt.Println("──────────────────────────")

			color.Cyan("\n? Class name: ")
			className := readInput(reader)

			color.Cyan("? What does this class represent (in plain English): ")
			description := readInput(reader)

			color.Cyan("? Attributes (comma-separated, e.g., name, age, email): ")
			attrs := readInput(reader)

			color.Cyan("? Attribute types (comma-separated, e.g., str, int, str): ")
			types := readInput(reader)

			color.Cyan("? Default values (comma-separated, use 'none' for no default, e.g., none, 0, none): ")
			defaults := readInput(reader)

			color.Cyan("? What methods does it need?")
			fmt.Println("  [1] Just __init__ and __repr__")
			fmt.Println("  [2] Add custom methods (I'll describe them)")
			fmt.Println("  [3] Use @dataclass (auto-generates boilerplate)")
			fmt.Print("\n  Select: ")
			methodChoice := readInputWithValidation(reader, []string{"1", "2", "3"})

			attrList := strings.Split(attrs, ",")
			typeList := strings.Split(types, ",")
			defaultList := strings.Split(defaults, ",")

			var code string

			if methodChoice == "3" {
				// Dataclass
				var fields []string
				for i, a := range attrList {
					a = strings.TrimSpace(a)
					t := "str"
					if i < len(typeList) {
						t = strings.TrimSpace(typeList[i])
					}
					d := "none"
					if i < len(defaultList) {
						d = strings.TrimSpace(defaultList[i])
					}
					if d == "none" || d == "" {
						fields = append(fields, fmt.Sprintf("    %s: %s", a, t))
					} else {
						fields = append(fields, fmt.Sprintf("    %s: %s = %s", a, t, d))
					}
				}
				code = fmt.Sprintf(`from dataclasses import dataclass

@dataclass
class %s:
    """%s"""
%s`,
					className,
					description,
					strings.Join(fields, "\n"),
				)
			} else {
				// Regular class
				var initParams []string
				var initBody []string
				var reprFields []string
				for i, a := range attrList {
					a = strings.TrimSpace(a)
					t := "str"
					if i < len(typeList) {
						t = strings.TrimSpace(typeList[i])
					}
					d := "none"
					if i < len(defaultList) {
						d = strings.TrimSpace(defaultList[i])
					}
					if d == "none" || d == "" {
						initParams = append(initParams, fmt.Sprintf("%s: %s", a, t))
					} else {
						initParams = append(initParams, fmt.Sprintf("%s: %s = %s", a, t, d))
					}
					initBody = append(initBody, fmt.Sprintf("        self.%s = %s", a, a))
					reprFields = append(reprFields, fmt.Sprintf("%s={self.%s}", a, a))
				}

				code = fmt.Sprintf(`class %s:
    """%s"""

    def __init__(self, %s):
%s

    def __repr__(self):
        return f"%s(%s)"`,
					className,
					description,
					strings.Join(initParams, ", "),
					strings.Join(initBody, "\n"),
					className,
					strings.Join(reprFields, ", "),
				)

				// Add custom methods
				if methodChoice == "2" {
					color.Cyan("\n? How many custom methods: ")
					countStr := readInput(reader)
					count := 1
					fmt.Sscanf(countStr, "%d", &count)

					for m := 0; m < count; m++ {
						fmt.Printf("\n  --- Method %d ---\n", m+1)
						color.Cyan("  ? Method name: ")
						methName := readInput(reader)
						color.Cyan("  ? Parameters (besides self, comma-separated, or 'none'): ")
						methParams := readInput(reader)
						color.Cyan("  ? What does it do (plain English): ")
						methDesc := readInput(reader)
						color.Cyan("  ? Does it return something? (yes/no): ")
						methReturn := readInput(reader)

						params := "self"
						if methParams != "none" && methParams != "" {
							params = "self, " + methParams
						}

						var body string
						if strings.ToLower(methReturn) == "yes" {
							body = fmt.Sprintf("        # TODO: %s\n        return None", methDesc)
						} else {
							body = fmt.Sprintf("        # TODO: %s\n        pass", methDesc)
						}

						code += fmt.Sprintf("\n\n    def %s(%s):\n        \"\"\"%s\"\"\"\n%s", methName, params, methDesc, body)
					}
				}
			}

			printGenerated(code)
		},
	}
}

// ==================== DSA ====================

func newDSACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dsa",
		Short: "⚡ DSA patterns (optimized for speed & memory)",
		Long:  "Generate optimized Data Structures & Algorithms code — interview ready.",
	}

	cmd.AddCommand(
		newDSASortCmd(),
		newDSASearchCmd(),
		newDSAPatternCmd(),
	)

	return cmd
}

func newDSASortCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sort",
		Short: "Sorting algorithms (customized to your data)",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n⚡ Python Sorting — Customized")
			fmt.Println("──────────────────────────────")

			color.Cyan("\n? What are you sorting (e.g., numbers, strings, objects, dicts): ")
			dataDesc := readInput(reader)

			color.Cyan("? Variable name of your list: ")
			listName := readInput(reader)

			color.Cyan("? Sort by what?")
			fmt.Println("  [1] Natural order (ascending)")
			fmt.Println("  [2] Reverse order (descending)")
			fmt.Println("  [3] By a specific key/attribute (e.g., sort employees by salary)")
			fmt.Println("  [4] Custom comparison")
			fmt.Print("\n  Select: ")
			sortChoice := readInputWithValidation(reader, []string{"1", "2", "3", "4"})

			color.Cyan("? Which algorithm?")
			fmt.Println("  [1] Python built-in (best for most cases) — O(n log n)")
			fmt.Println("  [2] Quick Sort (manual implementation)")
			fmt.Println("  [3] Merge Sort (stable, good for linked data)")
			fmt.Println("  [4] Bubble Sort (educational only)")
			fmt.Print("\n  Select: ")
			algoChoice := readInputWithValidation(reader, []string{"1", "2", "3", "4"})

			var code string

			switch algoChoice {
			case "1":
				switch sortChoice {
				case "1":
					code = fmt.Sprintf("# Sort %s in ascending order\n# Time: O(n log n), Space: O(n)\n%s.sort()\nprint(%s)", dataDesc, listName, listName)
				case "2":
					code = fmt.Sprintf("# Sort %s in descending order\n# Time: O(n log n), Space: O(n)\n%s.sort(reverse=True)\nprint(%s)", dataDesc, listName, listName)
				case "3":
					color.Cyan("? Key/attribute to sort by (e.g., 'salary', 'age', 'name'): ")
					key := readInput(reader)
					if strings.ToLower(dataDesc) == "dicts" || strings.Contains(strings.ToLower(dataDesc), "dict") {
						code = fmt.Sprintf("# Sort %s by '%s'\n# Time: O(n log n), Space: O(n)\n%s.sort(key=lambda item: item[\"%s\"])\nprint(%s)", dataDesc, key, listName, key, listName)
					} else {
						code = fmt.Sprintf("# Sort %s by %s attribute\n# Time: O(n log n), Space: O(n)\n%s.sort(key=lambda item: item.%s)\nprint(%s)", dataDesc, key, listName, key, listName)
					}
				case "4":
					color.Cyan("? Describe your sort rule (e.g., shorter strings first, then alphabetical): ")
					rule := readInput(reader)
					code = fmt.Sprintf("# Custom sort: %s\n# Time: O(n log n), Space: O(n)\n%s.sort(key=lambda item: (len(item), item))  # TODO: adjust key for: %s\nprint(%s)", rule, listName, rule, listName)
				}
			case "2":
				code = generateQuickSortDynamic(listName, dataDesc)
			case "3":
				code = generateMergeSortDynamic(listName, dataDesc)
			case "4":
				code = generateBubbleSortDynamic(listName, dataDesc)
			}

			printGenerated(code)
		},
	}
}

func newDSASearchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "Searching algorithms (customized to your data)",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n⚡ Python Search — Customized")
			fmt.Println("─────────────────────────────")

			color.Cyan("\n? What are you searching in (e.g., list of numbers, list of names, dict): ")
			searchDesc := readInput(reader)
			_ = searchDesc // used for context in comments

			color.Cyan("? Variable name: ")
			varName := readInput(reader)

			color.Cyan("? What are you looking for: ")
			target := readInput(reader)

			color.Cyan("? Is your data sorted?")
			fmt.Println("  [1] Yes — use binary search (O(log n), fastest)")
			fmt.Println("  [2] No — use linear scan or hash (O(n))")
			fmt.Println("  [3] I need to find multiple matches (filter)")
			fmt.Print("\n  Select: ")
			choice := readInputWithValidation(reader, []string{"1", "2", "3"})

			var code string

			switch choice {
			case "1":
				code = fmt.Sprintf(`# Binary Search for %s in %s
# Time: O(log n), Space: O(1)
# Requires: %s must be sorted!

def binary_search(%s, target):
    low, high = 0, len(%s) - 1

    while low <= high:
        mid = (low + high) // 2
        if %s[mid] == target:
            return mid
        elif %s[mid] < target:
            low = mid + 1
        else:
            high = mid - 1

    return -1  # not found

# Usage:
index = binary_search(%s, %s)
if index != -1:
    print(f"Found at index {index}")
else:
    print("Not found")`, target, varName, varName, varName, varName, varName, varName, varName, target)
			case "2":
				code = fmt.Sprintf(`# Search for %s in %s
# Time: O(n), Space: O(1)

# Method 1: Simple check
if %s in %s:
    print(f"Found: %s")
else:
    print("Not found")

# Method 2: Get index
try:
    index = %s.index(%s)
    print(f"Found at index {index}")
except ValueError:
    print("Not found")

# Method 3: Using hash set for repeated lookups (O(1) per lookup after O(n) setup)
lookup_set = set(%s)
if %s in lookup_set:
    print("Found (O(1) lookup)")`, target, varName, target, varName, target, varName, target, varName, target)
			case "3":
				color.Cyan("? Condition to filter by (e.g., item > 50, 'python' in item): ")
				condition := readInput(reader)
				code = fmt.Sprintf(`# Filter %s where %s
# Time: O(n), Space: O(k) where k = number of matches

results = [item for item in %s if %s]
print(f"Found {len(results)} matches:")
for item in results:
    print(f"  {item}")`, varName, condition, varName, condition)
			}

			printGenerated(code)
		},
	}
}

func newDSAPatternCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pattern",
		Short: "Common DSA patterns (interview favorites)",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n⚡ Python DSA Patterns")
			fmt.Println("──────────────────────")

			color.Cyan("\n? Which pattern do you need?")
			fmt.Println("  [1] Two Sum (find pair summing to target)        — O(n)")
			fmt.Println("  [2] Sliding Window (max/min in subarray)         — O(n)")
			fmt.Println("  [3] Two Pointer (pair finding in sorted array)   — O(n)")
			fmt.Println("  [4] Stack (matching brackets, undo operations)   — O(n)")
			fmt.Println("  [5] BFS (shortest path, level traversal)         — O(V+E)")
			fmt.Println("  [6] DFS (explore all paths, cycle detection)     — O(V+E)")
			fmt.Println("  [7] Dynamic Programming (optimize recursion)     — varies")
			fmt.Println("  [8] Reverse Linked List                          — O(n)")
			fmt.Print("\n  Select: ")
			choice := readInputWithValidation(reader, []string{"1", "2", "3", "4", "5", "6", "7", "8"})

			var code string

			switch choice {
			case "1":
				color.Cyan("\n? Your array variable name: ")
				arrName := readInput(reader)
				color.Cyan("? Target sum: ")
				target := readInput(reader)
				code = fmt.Sprintf(`# Two Sum — find indices of two numbers adding to %s
# Time: O(n), Space: O(n)

def two_sum(%s, target):
    seen = {}  # value → index
    for i, num in enumerate(%s):
        complement = target - num
        if complement in seen:
            return [seen[complement], i]
        seen[num] = i
    return []

# Usage:
result = two_sum(%s, %s)
if result:
    print(f"Indices: {result}")
    print(f"Values: {%s[result[0]]} + {%s[result[1]]} = %s")
else:
    print("No pair found")`, target, arrName, arrName, arrName, target, arrName, arrName, target)
			case "2":
				color.Cyan("\n? Your array variable name: ")
				arrName := readInput(reader)
				color.Cyan("? Window size: ")
				windowSize := readInput(reader)
				color.Cyan("? What to find in the window?")
				fmt.Println("  [1] Maximum sum")
				fmt.Println("  [2] Minimum sum")
				fmt.Println("  [3] Average")
				fmt.Print("  Select: ")
				winChoice := readInputWithValidation(reader, []string{"1", "2", "3"})

				operation := "max"
				if winChoice == "2" {
					operation = "min"
				} else if winChoice == "3" {
					operation = "avg"
				}

				if operation == "avg" {
					code = fmt.Sprintf(`# Sliding Window — average of subarrays of size %s
# Time: O(n), Space: O(1)

k = %s
window_sum = sum(%s[:k])
best_avg = window_sum / k

for i in range(k, len(%s)):
    window_sum += %s[i] - %s[i - k]
    current_avg = window_sum / k
    best_avg = max(best_avg, current_avg)

print(f"Best average: {best_avg:.2f}")`, windowSize, windowSize, arrName, arrName, arrName, arrName)
				} else {
					code = fmt.Sprintf(`# Sliding Window — %s sum of subarrays of size %s
# Time: O(n), Space: O(1)

k = %s
window_sum = sum(%s[:k])
best = window_sum

for i in range(k, len(%s)):
    window_sum += %s[i] - %s[i - k]
    best = %s(best, window_sum)

print(f"Best %s sum (window=%s): {best}")`, operation, windowSize, windowSize, arrName, arrName, arrName, arrName, operation, operation, windowSize)
				}
			case "3":
				color.Cyan("\n? Your sorted array name: ")
				arrName := readInput(reader)
				color.Cyan("? Target sum: ")
				target := readInput(reader)
				code = fmt.Sprintf(`# Two Pointer — find pair in sorted array summing to %s
# Time: O(n), Space: O(1)
# Requires: %s must be sorted!

left, right = 0, len(%s) - 1

while left < right:
    current_sum = %s[left] + %s[right]
    if current_sum == %s:
        print(f"Found: {%s[left]} + {%s[right]} = %s")
        break
    elif current_sum < %s:
        left += 1
    else:
        right -= 1
else:
    print("No pair found")`, target, arrName, arrName, arrName, arrName, target, arrName, arrName, target, target)
			case "4":
				code = generateValidParentheses()
			case "5":
				code = generateBFS()
			case "6":
				code = generateDFS()
			case "7":
				color.Cyan("\n? What kind of DP problem?")
				fmt.Println("  [1] Fibonacci / climbing stairs")
				fmt.Println("  [2] Coin change (minimum coins to make amount)")
				fmt.Println("  [3] Longest common subsequence")
				fmt.Print("  Select: ")
				dpChoice := readInputWithValidation(reader, []string{"1", "2", "3"})
				switch dpChoice {
				case "1":
					code = generateDPFibonacci()
				case "2":
					code = generateCoinChange()
				case "3":
					code = generateLCS()
				}
			case "8":
				code = generateReverseLinkedList()
			}

			printGenerated(code)
		},
	}
}

// ==================== HELPERS ====================

func printGenerated(code string) {
	green := color.New(color.FgGreen, color.Bold)
	cyan := color.New(color.FgCyan)

	fmt.Println()
	green.Println("✅ Generated:")
	fmt.Println("─────────────")
	cyan.Printf("\n%s\n", code)

	// Copy to clipboard
	err := clipboard.WriteAll(code)
	if err != nil {
		color.Yellow("\n⚠️  Could not copy to clipboard: %v", err)
	} else {
		green.Println("\n📋 Copied to clipboard!")
	}

	// Save to file if --output flag is set
	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(code+"\n"), 0644)
		if err != nil {
			color.Red("\n❌ Failed to write to file: %v", err)
		} else {
			green.Printf("💾 Saved to %s\n", outputFile)
		}
	}
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readInputWithValidation(reader *bufio.Reader, validOptions []string) string {
	for {
		input := readInput(reader)
		for _, opt := range validOptions {
			if input == opt {
				return input
			}
		}
		color.Red("  ❌ Invalid selection '%s'. Valid options: [%s]", input, strings.Join(validOptions, ", "))
		color.Cyan("  Please try again: ")
	}
}
