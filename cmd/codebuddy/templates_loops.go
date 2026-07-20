package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ==================== DYNAMIC LIST LOOP ====================

func dynamicListLoop(reader *bufio.Reader) string {
	fmt.Print("\n? What's in your list (e.g., names, numbers, servers, files): ")
	content := readInput(reader)

	fmt.Print("? Your list variable name: ")
	listName := readInput(reader)

	fmt.Print("? What do you call each item: ")
	itemName := readInput(reader)

	fmt.Println("? What do you want to do with each item?")
	fmt.Println("  [1] Print it")
	fmt.Println("  [2] Print with index number")
	fmt.Println("  [3] Check a condition and act")
	fmt.Println("  [4] Transform it (change each item)")
	fmt.Println("  [5] Collect items that match a condition")
	fmt.Println("  [6] Count items matching a condition")
	fmt.Print("\n  Select: ")
	action := readInput(reader)

	switch action {
	case "1":
		return fmt.Sprintf("# Print all %s\nfor %s in %s:\n    print(%s)", content, itemName, listName, itemName)
	case "2":
		return fmt.Sprintf("# Print all %s with index\nfor index, %s in enumerate(%s):\n    print(f\"{index + 1}. {%s}\")", content, itemName, listName, itemName)
	case "3":
		fmt.Printf("? What condition to check (use '%s' as variable, e.g., %s > 10): ", itemName, itemName)
		condition := readInput(reader)
		fmt.Print("? What to do when TRUE: ")
		trueAction := readInput(reader)
		fmt.Print("? What to do when FALSE (or 'skip' to do nothing): ")
		falseAction := readInput(reader)

		if strings.ToLower(falseAction) == "skip" || falseAction == "" {
			return fmt.Sprintf("# Check %s in %s\nfor %s in %s:\n    if %s:\n        %s", content, listName, itemName, listName, condition, trueAction)
		}
		return fmt.Sprintf("# Check %s in %s\nfor %s in %s:\n    if %s:\n        %s\n    else:\n        %s", content, listName, itemName, listName, condition, trueAction, falseAction)
	case "4":
		fmt.Printf("? How to transform each %s (e.g., %s.upper(), %s * 2): ", itemName, itemName, itemName)
		transform := readInput(reader)
		fmt.Print("? Store results in what variable: ")
		resultVar := readInput(reader)
		return fmt.Sprintf("# Transform %s\n%s = []\nfor %s in %s:\n    %s.append(%s)\nprint(%s)", content, resultVar, itemName, listName, resultVar, transform, resultVar)
	case "5":
		fmt.Printf("? Condition to match (e.g., %s > 100, 'active' in %s): ", itemName, itemName)
		condition := readInput(reader)
		fmt.Print("? Store matches in what variable: ")
		resultVar := readInput(reader)
		return fmt.Sprintf("# Filter %s where %s\n%s = [%s for %s in %s if %s]\nprint(f\"Found {len(%s)} matches:\")\nfor item in %s:\n    print(f\"  {item}\")", content, condition, resultVar, itemName, itemName, listName, condition, resultVar, resultVar)
	case "6":
		fmt.Printf("? Condition to count (e.g., %s > 50): ", itemName)
		condition := readInput(reader)
		return fmt.Sprintf("# Count %s where %s\ncount = sum(1 for %s in %s if %s)\nprint(f\"Count: {count}\")", content, condition, itemName, listName, condition)
	default:
		return fmt.Sprintf("for %s in %s:\n    print(%s)", itemName, listName, itemName)
	}
}

// ==================== DYNAMIC RANGE LOOP ====================

func dynamicRangeLoop(reader *bufio.Reader) string {
	fmt.Print("\n? Start number: ")
	start := readInput(reader)

	fmt.Print("? End number: ")
	end := readInput(reader)

	fmt.Print("? Step (default 1, or e.g., 2 for every other number): ")
	step := readInput(reader)
	if step == "" || step == "1" {
		step = ""
	}

	fmt.Print("? Variable name: ")
	varName := readInput(reader)

	fmt.Println("? What to do inside?")
	fmt.Println("  [1] Print the number")
	fmt.Println("  [2] Calculate something (square, cube, etc.)")
	fmt.Println("  [3] Accumulate (sum, product, etc.)")
	fmt.Println("  [4] Check condition (only even, only odd, divisible by X)")
	fmt.Println("  [5] Build a list from the numbers")
	fmt.Print("\n  Select: ")
	action := readInput(reader)

	rangeExpr := fmt.Sprintf("range(%s, %s)", start, end)
	if step != "" {
		rangeExpr = fmt.Sprintf("range(%s, %s, %s)", start, end, step)
	}

	switch action {
	case "1":
		return fmt.Sprintf("for %s in %s:\n    print(%s)", varName, rangeExpr, varName)
	case "2":
		fmt.Printf("? Expression to calculate (use '%s', e.g., %s ** 2): ", varName, varName)
		expr := readInput(reader)
		return fmt.Sprintf("# Calculate for each number\nfor %s in %s:\n    result = %s\n    print(f\"{%s} → {result}\")", varName, rangeExpr, expr, varName)
	case "3":
		fmt.Print("? What to accumulate (sum/product/max/min): ")
		accType := readInput(reader)
		fmt.Print("? Starting value (e.g., 0 for sum, 1 for product): ")
		startVal := readInput(reader)

		operator := "+="
		if strings.Contains(strings.ToLower(accType), "product") {
			operator = "*="
		}
		return fmt.Sprintf("# %s of numbers from %s to %s\ntotal = %s\nfor %s in %s:\n    total %s %s\nprint(f\"%s: {total}\")", accType, start, end, startVal, varName, rangeExpr, operator, varName, accType)
	case "4":
		fmt.Printf("? Condition (e.g., %s %% 2 == 0 for even): ", varName)
		condition := readInput(reader)
		return fmt.Sprintf("# Filter numbers where %s\nfor %s in %s:\n    if %s:\n        print(%s)", condition, varName, rangeExpr, condition, varName)
	case "5":
		fmt.Printf("? Expression for each item (e.g., %s * 10): ", varName)
		expr := readInput(reader)
		fmt.Print("? List variable name: ")
		listVar := readInput(reader)
		return fmt.Sprintf("# Build list from range\n%s = [%s for %s in %s]\nprint(%s)", listVar, expr, varName, rangeExpr, listVar)
	default:
		return fmt.Sprintf("for %s in %s:\n    print(%s)", varName, rangeExpr, varName)
	}
}

// ==================== DYNAMIC DICT LOOP ====================

func dynamicDictLoop(reader *bufio.Reader) string {
	fmt.Print("\n? What does your dictionary store (e.g., student scores, server configs, user data): ")
	content := readInput(reader)

	fmt.Print("? Dictionary variable name: ")
	dictName := readInput(reader)

	fmt.Print("? What do you call each key: ")
	keyName := readInput(reader)

	fmt.Print("? What do you call each value: ")
	valueName := readInput(reader)

	fmt.Println("? What do you want to do?")
	fmt.Println("  [1] Print all key-value pairs")
	fmt.Println("  [2] Filter by condition on value")
	fmt.Println("  [3] Filter by condition on key")
	fmt.Println("  [4] Find min/max value")
	fmt.Println("  [5] Transform values into a new dict")
	fmt.Println("  [6] Group/categorize items")
	fmt.Print("\n  Select: ")
	action := readInput(reader)

	switch action {
	case "1":
		return fmt.Sprintf("# Print all %s\nfor %s, %s in %s.items():\n    print(f\"{%s}: {%s}\")", content, keyName, valueName, dictName, keyName, valueName)
	case "2":
		fmt.Printf("? Condition on value (e.g., %s > 50, %s == 'active'): ", valueName, valueName)
		condition := readInput(reader)
		return fmt.Sprintf("# Filter %s where %s\nfor %s, %s in %s.items():\n    if %s:\n        print(f\"{%s}: {%s}\")", content, condition, keyName, valueName, dictName, condition, keyName, valueName)
	case "3":
		fmt.Printf("? Condition on key (e.g., %s.startswith('admin')): ", keyName)
		condition := readInput(reader)
		return fmt.Sprintf("# Filter %s where %s\nfor %s, %s in %s.items():\n    if %s:\n        print(f\"{%s}: {%s}\")", content, condition, keyName, valueName, dictName, condition, keyName, valueName)
	case "4":
		fmt.Println("  [1] Maximum value")
		fmt.Println("  [2] Minimum value")
		fmt.Print("  Select: ")
		minmax := readInput(reader)
		fn := "max"
		if minmax == "2" {
			fn = "min"
		}
		return fmt.Sprintf("# Find %s in %s\n%s_key = %s(%s, key=%s.get)\nprint(f\"%s: {%s_key} → {%s[%s_key]}\")", fn, content, fn, fn, dictName, dictName, strings.Title(fn), fn, dictName, fn)
	case "5":
		fmt.Printf("? How to transform each value (e.g., %s * 2, %s.upper()): ", valueName, valueName)
		transform := readInput(reader)
		return fmt.Sprintf("# Transform %s values\nnew_dict = {%s: %s for %s, %s in %s.items()}\nprint(new_dict)", content, keyName, transform, keyName, valueName, dictName)
	case "6":
		fmt.Print("? How to categorize (e.g., 'pass' if score >= 50 else 'fail'): ")
		category := readInput(reader)
		return fmt.Sprintf("# Categorize %s\nfrom collections import defaultdict\ngroups = defaultdict(list)\nfor %s, %s in %s.items():\n    category = %s\n    groups[category].append(%s)\n\nfor cat, items in groups.items():\n    print(f\"{cat}: {items}\")", content, keyName, valueName, dictName, category, keyName)
	default:
		return fmt.Sprintf("for %s, %s in %s.items():\n    print(f\"{%s}: {%s}\")", keyName, valueName, dictName, keyName, valueName)
	}
}

// ==================== DYNAMIC FILE LOOP ====================

func dynamicFileLoop(reader *bufio.Reader) string {
	fmt.Print("\n? File path (e.g., \"data.txt\", \"logs/app.log\"): ")
	filePath := readInput(reader)

	fmt.Print("? What do you call each line: ")
	lineName := readInput(reader)

	fmt.Println("? What do you want to do with the file?")
	fmt.Println("  [1] Print all lines")
	fmt.Println("  [2] Search for a keyword")
	fmt.Println("  [3] Count lines matching a pattern")
	fmt.Println("  [4] Parse CSV-style data (split by delimiter)")
	fmt.Println("  [5] Get specific line numbers")
	fmt.Print("\n  Select: ")
	action := readInput(reader)

	switch action {
	case "1":
		return fmt.Sprintf("# Read and print all lines\nwith open(%s, \"r\") as file:\n    for %s in file:\n        print(%s.strip())", filePath, lineName, lineName)
	case "2":
		fmt.Print("? Keyword to search: ")
		keyword := readInput(reader)
		return fmt.Sprintf("# Search for '%s' in file\nwith open(%s, \"r\") as file:\n    for line_num, %s in enumerate(file, 1):\n        if \"%s\" in %s:\n            print(f\"Line {line_num}: {%s.strip()}\")", keyword, filePath, lineName, keyword, lineName, lineName)
	case "3":
		fmt.Print("? Pattern to count (e.g., ERROR, WARNING): ")
		pattern := readInput(reader)
		return fmt.Sprintf("# Count lines containing '%s'\ncount = 0\nwith open(%s, \"r\") as file:\n    for %s in file:\n        if \"%s\" in %s:\n            count += 1\nprint(f\"Found {count} lines with '%s'\")", pattern, filePath, lineName, pattern, lineName, pattern)
	case "4":
		fmt.Print("? Delimiter (e.g., ',' for CSV, '\\t' for tab, '|' for pipe): ")
		delimiter := readInput(reader)
		fmt.Print("? How many columns: ")
		cols := readInput(reader)
		return fmt.Sprintf("# Parse %s as delimited file ('%s' separator, %s columns)\nwith open(%s, \"r\") as file:\n    for %s in file:\n        parts = %s.strip().split(\"%s\")\n        if len(parts) >= %s:\n            print(parts)", filePath, delimiter, cols, filePath, lineName, lineName, delimiter, cols)
	case "5":
		fmt.Print("? Which lines (e.g., first 10, last 5, specific like 5-15): ")
		lineSpec := readInput(reader)
		if strings.Contains(lineSpec, "first") {
			return fmt.Sprintf("# Read first N lines\nwith open(%s, \"r\") as file:\n    for i, %s in enumerate(file):\n        if i >= 10:  # adjust number\n            break\n        print(%s.strip())", filePath, lineName, lineName)
		}
		return fmt.Sprintf("# Read specific lines (%s)\nwith open(%s, \"r\") as file:\n    lines = file.readlines()\n    # Adjust range as needed\n    for %s in lines[4:15]:  # lines 5-15\n        print(%s.strip())", lineSpec, filePath, lineName, lineName)
	default:
		return fmt.Sprintf("with open(%s, \"r\") as file:\n    for %s in file:\n        print(%s.strip())", filePath, lineName, lineName)
	}
}

// ==================== DYNAMIC NESTED LOOP ====================

func dynamicNestedLoop(reader *bufio.Reader) string {
	fmt.Println("\n? What kind of nested operation?")
	fmt.Println("  [1] Loop through a 2D list/matrix")
	fmt.Println("  [2] Compare all pairs in a list")
	fmt.Println("  [3] Generate combinations from two lists")
	fmt.Println("  [4] Multiplication table")
	fmt.Println("  [5] Pattern printing (triangle, square)")
	fmt.Print("\n  Select: ")
	choice := readInput(reader)

	switch choice {
	case "1":
		fmt.Print("? Matrix variable name: ")
		matrix := readInput(reader)
		fmt.Println("? What to do?")
		fmt.Println("  [1] Print grid")
		fmt.Println("  [2] Find max/min")
		fmt.Println("  [3] Sum all elements")
		fmt.Print("  Select: ")
		subChoice := readInput(reader)
		switch subChoice {
		case "2":
			return fmt.Sprintf("# Find max in 2D matrix\n# Time: O(rows * cols), Space: O(1)\nmax_val = float('-inf')\nfor row in %s:\n    for item in row:\n        max_val = max(max_val, item)\nprint(f\"Max: {max_val}\")", matrix)
		case "3":
			return fmt.Sprintf("# Sum all elements in 2D matrix\n# Time: O(rows * cols), Space: O(1)\ntotal = sum(item for row in %s for item in row)\nprint(f\"Total: {total}\")", matrix)
		default:
			return fmt.Sprintf("# Print 2D matrix as grid\nfor row in %s:\n    print(\" \".join(str(x) for x in row))", matrix)
		}
	case "2":
		fmt.Print("? List variable name: ")
		listName := readInput(reader)
		fmt.Print("? What to check between pairs (e.g., find duplicates, find pairs summing to X): ")
		desc := readInput(reader)
		return fmt.Sprintf("# %s — compare all pairs\n# Time: O(n²), Space: O(1)\n# Note: Use hash map approach for O(n) if possible\nfor i in range(len(%s)):\n    for j in range(i + 1, len(%s)):\n        # Check pair: %s[i] and %s[j]\n        print(f\"Pair: {%s[i]}, {%s[j]}\")", desc, listName, listName, listName, listName, listName, listName)
	case "3":
		fmt.Print("? First list name: ")
		list1 := readInput(reader)
		fmt.Print("? Second list name: ")
		list2 := readInput(reader)
		fmt.Print("? What to do with each combination: ")
		action := readInput(reader)
		return fmt.Sprintf("# All combinations from %s and %s\n# Time: O(n * m), Space: O(1)\nfor a in %s:\n    for b in %s:\n        %s", list1, list2, list1, list2, action)
	case "4":
		fmt.Print("? Up to what number: ")
		num := readInput(reader)
		return fmt.Sprintf("# Multiplication table up to %s\nfor i in range(1, %s + 1):\n    for j in range(1, %s + 1):\n        print(f\"{i*j:4}\", end=\"\")\n    print()", num, num, num)
	case "5":
		fmt.Print("? Number of rows: ")
		rows := readInput(reader)
		fmt.Println("? Pattern type:")
		fmt.Println("  [1] Right triangle (*)")
		fmt.Println("  [2] Pyramid (centered)")
		fmt.Println("  [3] Number triangle")
		fmt.Print("  Select: ")
		patChoice := readInput(reader)
		switch patChoice {
		case "2":
			return fmt.Sprintf("# Pyramid pattern\nn = %s\nfor i in range(1, n + 1):\n    spaces = \" \" * (n - i)\n    stars = \"*\" * (2 * i - 1)\n    print(spaces + stars)", rows)
		case "3":
			return fmt.Sprintf("# Number triangle\nfor i in range(1, %s + 1):\n    for j in range(1, i + 1):\n        print(j, end=\" \")\n    print()", rows)
		default:
			return fmt.Sprintf("# Right triangle\nfor i in range(1, %s + 1):\n    print(\"*\" * i)", rows)
		}
	default:
		return "for i in range(3):\n    for j in range(3):\n        print(f\"({i}, {j})\")"
	}
}

// ==================== DYNAMIC LIST COMPREHENSION ====================

func dynamicListComprehension(reader *bufio.Reader) string {
	fmt.Print("\n? Source list/range variable: ")
	source := readInput(reader)

	fmt.Print("? Item variable name: ")
	item := readInput(reader)

	fmt.Println("? What do you want to create?")
	fmt.Println("  [1] Transform each item (map)")
	fmt.Println("  [2] Filter items (keep only matching)")
	fmt.Println("  [3] Transform + filter (both)")
	fmt.Println("  [4] Create a dictionary")
	fmt.Println("  [5] Flatten a nested list")
	fmt.Print("\n  Select: ")
	choice := readInput(reader)

	switch choice {
	case "1":
		fmt.Printf("? How to transform (e.g., %s.upper(), %s * 2, len(%s)): ", item, item, item)
		transform := readInput(reader)
		fmt.Print("? Result variable name: ")
		result := readInput(reader)
		return fmt.Sprintf("# Transform: %s\n# Time: O(n), Space: O(n)\n%s = [%s for %s in %s]\nprint(%s)", transform, result, transform, item, source, result)
	case "2":
		fmt.Printf("? Condition to keep (e.g., %s > 10, len(%s) > 3): ", item, item)
		condition := readInput(reader)
		fmt.Print("? Result variable name: ")
		result := readInput(reader)
		return fmt.Sprintf("# Filter: keep where %s\n# Time: O(n), Space: O(k) matches\n%s = [%s for %s in %s if %s]\nprint(f\"{len(%s)} items matched\")\nprint(%s)", condition, result, item, item, source, condition, result, result)
	case "3":
		fmt.Printf("? Condition (e.g., %s > 0): ", item)
		condition := readInput(reader)
		fmt.Printf("? Transform (e.g., %s ** 2): ", item)
		transform := readInput(reader)
		fmt.Print("? Result variable name: ")
		result := readInput(reader)
		return fmt.Sprintf("# Filter + Transform\n# Time: O(n), Space: O(k) matches\n%s = [%s for %s in %s if %s]\nprint(%s)", result, transform, item, source, condition, result)
	case "4":
		fmt.Printf("? Key expression (e.g., %s, %s.lower()): ", item, item)
		keyExpr := readInput(reader)
		fmt.Printf("? Value expression (e.g., len(%s), %s * 2): ", item, item)
		valExpr := readInput(reader)
		fmt.Print("? Result variable name: ")
		result := readInput(reader)
		return fmt.Sprintf("# Create dictionary\n# Time: O(n), Space: O(n)\n%s = {%s: %s for %s in %s}\nprint(%s)", result, keyExpr, valExpr, item, source, result)
	case "5":
		return fmt.Sprintf("# Flatten nested list\n# Time: O(n*m), Space: O(n*m)\nflat = [%s for sublist in %s for %s in sublist]\nprint(flat)", item, source, item)
	default:
		return fmt.Sprintf("result = [%s for %s in %s]", item, item, source)
	}
}
