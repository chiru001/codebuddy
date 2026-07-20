package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ==================== GENERATOR COMMAND ====================

func newGeneratorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generator",
		Short: "🔄 Generator functions (yield, lazy evaluation)",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🔄 Python Generator Builder")
			fmt.Println("────────────────────────────")

			fmt.Println("\n? What kind of generator?")
			fmt.Println("  [1] Yield items one at a time (memory efficient)")
			fmt.Println("  [2] Infinite generator (counter, IDs, etc.)")
			fmt.Println("  [3] Read large file line by line")
			fmt.Println("  [4] Chunked processing (batch items)")
			fmt.Println("  [5] Pipeline (chain generators)")
			fmt.Print("\n  Select: ")
			choice := readInput(reader)

			var code string

			switch choice {
			case "1":
				fmt.Print("\n? Generator function name: ")
				name := readInput(reader)
				fmt.Print("? What data source (e.g., list, database query, API response): ")
				source := readInput(reader)
				fmt.Print("? What to yield per iteration: ")
				yieldExpr := readInput(reader)
				fmt.Print("? Any filter condition (or 'none'): ")
				condition := readInput(reader)

				if strings.ToLower(condition) == "none" || condition == "" {
					code = fmt.Sprintf(`def %s(%s):
    """
    Generator — yields %s one at a time.
    Memory: O(1) — only one item in memory at a time!
    """
    for item in %s:
        yield %s

# Usage:
# for result in %s(data):
#     print(result)
#
# Or convert to list (careful with large data!):
# all_results = list(%s(data))`, name, source, yieldExpr, source, yieldExpr, name, name)
				} else {
					code = fmt.Sprintf(`def %s(%s):
    """
    Generator — yields %s where %s.
    Memory: O(1) — only one item in memory at a time!
    """
    for item in %s:
        if %s:
            yield %s

# Usage:
# for result in %s(data):
#     print(result)`, name, source, yieldExpr, condition, source, condition, yieldExpr, name)
				}
			case "2":
				fmt.Print("\n? Generator name (e.g., id_generator, counter): ")
				name := readInput(reader)
				fmt.Print("? Start value: ")
				start := readInput(reader)
				fmt.Print("? Step (increment by): ")
				step := readInput(reader)
				code = fmt.Sprintf(`def %s(start=%s, step=%s):
    """
    Infinite generator — produces values forever.
    Memory: O(1) — never stores all values!
    Use next() to get one value at a time.
    """
    current = start
    while True:
        yield current
        current += step

# Usage:
# gen = %s()
# print(next(gen))  # %s
# print(next(gen))  # %s + %s
# print(next(gen))  # keeps going forever...
#
# Or use with limit:
# from itertools import islice
# first_10 = list(islice(%s(), 10))`, name, start, step, name, start, start, step, name)
			case "3":
				fmt.Print("\n? Function name: ")
				name := readInput(reader)
				fmt.Print("? Any processing per line (e.g., strip, split by comma, parse JSON, or 'none'): ")
				processing := readInput(reader)

				processLine := "line.strip()"
				if strings.Contains(processing, "split") {
					processLine = "line.strip().split(\",\")"
				} else if strings.Contains(processing, "json") {
					processLine = "json.loads(line)"
				} else if processing == "none" || processing == "" {
					processLine = "line.strip()"
				} else {
					processLine = processing
				}

				code = fmt.Sprintf(`def %s(filepath):
    """
    Generator — reads file line by line.
    Memory: O(1) — handles files of ANY size (even 100GB)!
    """
    with open(filepath, "r") as f:
        for line in f:
            yield %s

# Usage:
# for line in %s("huge_file.log"):
#     print(line)
#
# Count lines without loading entire file:
# total = sum(1 for _ in %s("huge_file.log"))`, name, processLine, name, name)
			case "4":
				fmt.Print("\n? Function name: ")
				name := readInput(reader)
				fmt.Print("? Default chunk size: ")
				chunkSize := readInput(reader)
				code = fmt.Sprintf(`def %s(data, chunk_size=%s):
    """
    Generator — yields data in chunks/batches.
    Memory: O(chunk_size) — processes large datasets without loading all.
    Use for: batch database inserts, API pagination, parallel processing.
    """
    for i in range(0, len(data), chunk_size):
        yield data[i:i + chunk_size]

# Usage:
# big_list = list(range(10000))
# for batch in %s(big_list, 100):
#     process_batch(batch)  # 100 items at a time
#     print(f"Processed batch of {len(batch)}")`, name, chunkSize, name)
			case "5":
				code = `# Generator Pipeline — chain generators for data processing
# Each step processes ONE item at a time = O(1) memory!

def read_lines(filepath):
    """Stage 1: Read file lazily."""
    with open(filepath) as f:
        for line in f:
            yield line.strip()

def filter_lines(lines, keyword):
    """Stage 2: Keep only matching lines."""
    for line in lines:
        if keyword in line:
            yield line

def transform(lines):
    """Stage 3: Transform each line."""
    for line in lines:
        yield line.upper()

# Usage — chain them together:
# pipeline = transform(filter_lines(read_lines("app.log"), "ERROR"))
# for result in pipeline:
#     print(result)
#
# Memory: O(1) regardless of file size!
# Each item flows through ALL stages before the next item starts.`
			default:
				fmt.Println("❌ Invalid selection")
				return
			}

			printGenerated(code)
		},
	}
}

// ==================== ERROR HANDLING COMMAND ====================

func newErrorHandlingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "error-handling",
		Short: "⚠️ Try/except patterns",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n⚠️ Python Error Handling Generator")
			fmt.Println("───────────────────────────────────")

			fmt.Println("\n? What kind of error handling?")
			fmt.Println("  [1] Basic try/except (catch specific errors)")
			fmt.Println("  [2] Custom exception class")
			fmt.Println("  [3] Retry with backoff")
			fmt.Println("  [4] Context manager (with statement)")
			fmt.Println("  [5] Graceful degradation (fallback values)")
			fmt.Print("\n  Select: ")
			choice := readInput(reader)

			var code string

			switch choice {
			case "1":
				fmt.Print("\n? What operation might fail (e.g., file read, API call, DB query): ")
				operation := readInput(reader)
				fmt.Print("? Expected error types (comma-separated, e.g., FileNotFoundError, ValueError): ")
				errors := readInput(reader)

				errorList := strings.Split(errors, ",")
				var exceptBlocks string
				for _, e := range errorList {
					e = strings.TrimSpace(e)
					if e == "" {
						continue
					}
					exceptBlocks += fmt.Sprintf(`    except %s as e:
        print(f"%s: {e}")
        # TODO: handle %s
`, e, e, e)
				}

				code = fmt.Sprintf(`# Error handling for: %s
try:
    # TODO: %s
    result = None

%sexcept Exception as e:
    # Catch-all for unexpected errors
    print(f"Unexpected error: {e}")
    raise  # re-raise if you can't handle it
else:
    # Runs ONLY if no exception occurred
    print("Success!")
finally:
    # ALWAYS runs (cleanup: close files, connections, etc.)
    pass`, operation, operation, exceptBlocks)
			case "2":
				fmt.Print("\n? Exception class name (e.g., ValidationError, AuthError): ")
				name := readInput(reader)
				fmt.Print("? What it represents: ")
				desc := readInput(reader)
				fmt.Print("? Extra fields besides message (comma-separated, or 'none'): ")
				fields := readInput(reader)

				var extraFields string
				var initParams string
				var initBody string
				if fields != "none" && fields != "" {
					fieldList := strings.Split(fields, ",")
					for _, f := range fieldList {
						f = strings.TrimSpace(f)
						initParams += fmt.Sprintf(", %s=None", f)
						initBody += fmt.Sprintf("        self.%s = %s\n", f, f)
						extraFields += fmt.Sprintf("        if self.%s:\n            parts.append(f\"%s={self.%s}\")\n", f, f, f)
					}
				}

				code = fmt.Sprintf(`class %s(Exception):
    """%s"""

    def __init__(self, message: str%s):
        super().__init__(message)
        self.message = message
%s
    def __str__(self):
        parts = [self.message]
%s        return " | ".join(parts)

# Usage:
# raise %s("Invalid input", field="email")
#
# try:
#     validate(data)
# except %s as e:
#     print(f"Validation failed: {e}")`, name, desc, initParams, initBody, extraFields, name, name)
			case "3":
				fmt.Print("\n? Function name to retry: ")
				funcName := readInput(reader)
				fmt.Print("? Max attempts: ")
				maxAttempts := readInput(reader)
				fmt.Print("? Delay between retries (seconds): ")
				delay := readInput(reader)

				code = fmt.Sprintf(`import time

def %s_with_retry(max_attempts=%s, delay=%s):
    """
    Retry with exponential backoff.
    Attempt 1: wait %ss, Attempt 2: wait %s*2s, Attempt 3: wait %s*4s...
    """
    for attempt in range(1, max_attempts + 1):
        try:
            result = %s()  # TODO: your actual call
            return result
        except Exception as e:
            if attempt == max_attempts:
                print(f"All {max_attempts} attempts failed!")
                raise
            wait = delay * (2 ** (attempt - 1))  # exponential backoff
            print(f"Attempt {attempt} failed: {e}. Retrying in {wait}s...")
            time.sleep(wait)

# Usage:
# result = %s_with_retry()`, funcName, maxAttempts, delay, delay, delay, delay, funcName, funcName)
			case "4":
				fmt.Print("\n? Context manager class name (e.g., DatabaseConnection, Timer): ")
				name := readInput(reader)
				fmt.Print("? What resource to manage: ")
				resource := readInput(reader)

				code = fmt.Sprintf(`class %s:
    """
    Context Manager — auto-cleanup of %s.
    Use with 'with' statement for guaranteed cleanup.
    """

    def __enter__(self):
        """Setup: acquire the resource."""
        print("Acquiring %s...")
        # TODO: open connection / acquire resource
        self.resource = None
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        """Cleanup: ALWAYS runs, even if error occurred inside 'with'."""
        print("Releasing %s...")
        # TODO: close connection / release resource
        if exc_type:
            print(f"Error occurred: {exc_val}")
        return False  # False = re-raise exception, True = suppress it

# Usage:
# with %s() as ctx:
#     # do work with ctx.resource
#     pass
# # cleanup happens automatically here!`, name, resource, resource, resource, name)
			case "5":
				fmt.Print("\n? What operation (e.g., API call, config read, cache lookup): ")
				operation := readInput(reader)
				fmt.Print("? Fallback value if it fails: ")
				fallback := readInput(reader)

				code = fmt.Sprintf(`def safe_%s():
    """
    Graceful degradation — returns fallback instead of crashing.
    Use for: non-critical operations where a default is acceptable.
    """
    try:
        # TODO: %s
        result = None
        return result
    except Exception as e:
        print(f"Warning: %s failed ({e}), using fallback")
        return %s

# Usage:
# value = safe_%s()  # never crashes, always returns something`, operation, operation, operation, fallback, operation)
			default:
				fmt.Println("❌ Invalid selection")
				return
			}

			printGenerated(code)
		},
	}
}

// ==================== FILE I/O COMMAND ====================

func newFileIOCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "file-io",
		Short: "📁 File operations (read, write, CSV, JSON)",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n📁 Python File I/O Generator")
			fmt.Println("─────────────────────────────")

			fmt.Println("\n? What file operation?")
			fmt.Println("  [1] Read a text file")
			fmt.Println("  [2] Write to a text file")
			fmt.Println("  [3] Read/Write JSON")
			fmt.Println("  [4] Read/Write CSV")
			fmt.Println("  [5] Append to a file (log style)")
			fmt.Print("\n  Select: ")
			choice := readInput(reader)

			var code string

			switch choice {
			case "1":
				fmt.Print("\n? File path variable/string: ")
				path := readInput(reader)
				fmt.Println("? How to read:")
				fmt.Println("  [1] Entire file at once")
				fmt.Println("  [2] Line by line (memory efficient)")
				fmt.Println("  [3] First N lines")
				fmt.Print("  Select: ")
				how := readInput(reader)

				switch how {
				case "1":
					code = fmt.Sprintf(`# Read entire file
with open(%s, "r") as f:
    content = f.read()
print(content)`, path)
				case "2":
					code = fmt.Sprintf(`# Read line by line (memory efficient for large files)
with open(%s, "r") as f:
    for line in f:
        print(line.strip())`, path)
				case "3":
					fmt.Print("  ? How many lines: ")
					n := readInput(reader)
					code = fmt.Sprintf(`# Read first %s lines
with open(%s, "r") as f:
    for i, line in enumerate(f):
        if i >= %s:
            break
        print(line.strip())`, n, path, n)
				}
			case "2":
				fmt.Print("\n? File path: ")
				path := readInput(reader)
				fmt.Print("? What data to write (variable name): ")
				data := readInput(reader)
				code = fmt.Sprintf(`# Write to file (overwrites existing content)
with open(%s, "w") as f:
    f.write(%s)
print(f"Written to %s")

# To write multiple lines:
# lines = ["line1", "line2", "line3"]
# with open(%s, "w") as f:
#     f.writelines(line + "\n" for line in lines)`, path, data, path, path)
			case "3":
				fmt.Println("\n? JSON operation:")
				fmt.Println("  [1] Read JSON file → Python dict")
				fmt.Println("  [2] Write Python dict → JSON file")
				fmt.Print("  Select: ")
				jsonOp := readInput(reader)

				fmt.Print("? File path: ")
				path := readInput(reader)

				if jsonOp == "1" {
					code = fmt.Sprintf(`import json

# Read JSON file into Python dict
with open(%s, "r") as f:
    data = json.load(f)

print(type(data))  # dict or list
print(data)`, path)
				} else {
					fmt.Print("? Variable name of your dict/list: ")
					varName := readInput(reader)
					code = fmt.Sprintf(`import json

# Write Python dict/list to JSON file
with open(%s, "w") as f:
    json.dump(%s, f, indent=2)

print(f"Saved to %s")`, path, varName, path)
				}
			case "4":
				fmt.Println("\n? CSV operation:")
				fmt.Println("  [1] Read CSV → list of dicts")
				fmt.Println("  [2] Write list of dicts → CSV")
				fmt.Print("  Select: ")
				csvOp := readInput(reader)

				fmt.Print("? File path: ")
				path := readInput(reader)

				if csvOp == "1" {
					code = fmt.Sprintf(`import csv

# Read CSV as list of dictionaries (header = keys)
with open(%s, "r") as f:
    reader = csv.DictReader(f)
    rows = list(reader)

print(f"Read {len(rows)} rows")
for row in rows[:5]:  # print first 5
    print(row)`, path)
				} else {
					fmt.Print("? Column headers (comma-separated): ")
					headers := readInput(reader)
					fmt.Print("? Data variable name (list of dicts): ")
					varName := readInput(reader)
					headerList := strings.Split(headers, ",")
					var trimmed []string
					for _, h := range headerList {
						trimmed = append(trimmed, fmt.Sprintf("\"%s\"", strings.TrimSpace(h)))
					}
					code = fmt.Sprintf(`import csv

# Write list of dicts to CSV
headers = [%s]

with open(%s, "w", newline="") as f:
    writer = csv.DictWriter(f, fieldnames=headers)
    writer.writeheader()
    writer.writerows(%s)

print(f"Written {len(%s)} rows to %s")`, strings.Join(trimmed, ", "), path, varName, varName, path)
				}
			case "5":
				fmt.Print("\n? Log file path: ")
				path := readInput(reader)
				code = fmt.Sprintf(`from datetime import datetime

def log(message, filepath=%s):
    """Append a timestamped log entry."""
    timestamp = datetime.now().strftime("%%Y-%%m-%%d %%H:%%M:%%S")
    entry = f"[{timestamp}] {message}\n"
    with open(filepath, "a") as f:
        f.write(entry)

# Usage:
# log("Application started")
# log("User logged in: chiru")
# log("ERROR: connection timeout")`, path)
			default:
				fmt.Println("❌ Invalid selection")
				return
			}

			printGenerated(code)
		},
	}
}

// ==================== API COMMAND ====================

func newAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "🌐 API code (requests, FastAPI endpoints, REST)",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🌐 Python API Generator")
			fmt.Println("────────────────────────")

			fmt.Println("\n? What do you need?")
			fmt.Println("  [1] Make HTTP request (call an external API)")
			fmt.Println("  [2] Create FastAPI endpoint (build your own API)")
			fmt.Println("  [3] Create Flask endpoint")
			fmt.Print("\n  Select: ")
			choice := readInput(reader)

			var code string

			switch choice {
			case "1":
				fmt.Println("\n? HTTP method:")
				fmt.Println("  [1] GET")
				fmt.Println("  [2] POST")
				fmt.Println("  [3] PUT")
				fmt.Println("  [4] DELETE")
				fmt.Print("  Select: ")
				method := readInput(reader)

				fmt.Print("? URL (or URL variable name): ")
				url := readInput(reader)

				fmt.Print("? Need authentication? (yes/no): ")
				auth := readInput(reader)

				var headers string
				if strings.ToLower(auth) == "yes" {
					headers = `
headers = {
    "Authorization": f"Bearer {token}",
    "Content-Type": "application/json"
}`
				} else {
					headers = `
headers = {"Content-Type": "application/json"}`
				}

				switch method {
				case "1":
					code = fmt.Sprintf(`import requests
%s

response = requests.get(%s, headers=headers, timeout=30)

if response.status_code == 200:
    data = response.json()
    print(data)
else:
    print(f"Error: {response.status_code} - {response.text}")`, headers, url)
				case "2":
					fmt.Print("? What data to send (variable name, or describe): ")
					payload := readInput(reader)
					code = fmt.Sprintf(`import requests
%s

payload = %s  # TODO: your request body

response = requests.post(%s, json=payload, headers=headers, timeout=30)

if response.status_code in (200, 201):
    data = response.json()
    print("Success:", data)
else:
    print(f"Error: {response.status_code} - {response.text}")`, headers, payload, url)
				case "3":
					fmt.Print("? What data to send: ")
					payload := readInput(reader)
					code = fmt.Sprintf(`import requests
%s

payload = %s

response = requests.put(%s, json=payload, headers=headers, timeout=30)

if response.status_code == 200:
    print("Updated:", response.json())
else:
    print(f"Error: {response.status_code}")`, headers, payload, url)
				case "4":
					code = fmt.Sprintf(`import requests
%s

response = requests.delete(%s, headers=headers, timeout=30)

if response.status_code in (200, 204):
    print("Deleted successfully")
else:
    print(f"Error: {response.status_code}")`, headers, url)
				}
			case "2":
				fmt.Print("\n? Endpoint path (e.g., /users, /items/{id}): ")
				path := readInput(reader)
				fmt.Println("? HTTP method:")
				fmt.Println("  [1] GET")
				fmt.Println("  [2] POST")
				fmt.Print("  Select: ")
				method := readInput(reader)

				if method == "1" {
					code = fmt.Sprintf(`from fastapi import FastAPI, HTTPException

app = FastAPI()

@app.get("%s")
async def get_handler():
    """GET %s"""
    try:
        # TODO: fetch data
        return {"data": [], "message": "success"}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

# Run: uvicorn main:app --reload`, path, path)
				} else {
					fmt.Print("? Request body fields (comma-separated): ")
					fields := readInput(reader)

					fieldList := strings.Split(fields, ",")
					var pydanticFields string
					for _, f := range fieldList {
						f = strings.TrimSpace(f)
						pydanticFields += fmt.Sprintf("    %s: str\n", f)
					}

					code = fmt.Sprintf(`from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

app = FastAPI()

class RequestBody(BaseModel):
%s

@app.post("%s")
async def post_handler(body: RequestBody):
    """POST %s"""
    try:
        # TODO: process body
        return {"message": "created", "data": body.dict()}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

# Run: uvicorn main:app --reload`, pydanticFields, path, path)
				}
			case "3":
				fmt.Print("\n? Endpoint path: ")
				path := readInput(reader)
				fmt.Println("? HTTP method:")
				fmt.Println("  [1] GET")
				fmt.Println("  [2] POST")
				fmt.Print("  Select: ")
				method := readInput(reader)

				if method == "1" {
					code = fmt.Sprintf(`from flask import Flask, jsonify

app = Flask(__name__)

@app.route("%s", methods=["GET"])
def get_handler():
    try:
        # TODO: fetch data
        return jsonify({"data": [], "message": "success"})
    except Exception as e:
        return jsonify({"error": str(e)}), 500

# Run: flask run`, path)
				} else {
					code = fmt.Sprintf(`from flask import Flask, jsonify, request

app = Flask(__name__)

@app.route("%s", methods=["POST"])
def post_handler():
    try:
        body = request.get_json()
        # TODO: process body
        return jsonify({"message": "created", "data": body}), 201
    except Exception as e:
        return jsonify({"error": str(e)}), 500

# Run: flask run`, path)
				}
			default:
				fmt.Println("❌ Invalid selection")
				return
			}

			printGenerated(code)
		},
	}
}
