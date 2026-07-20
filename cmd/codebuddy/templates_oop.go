package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ==================== OOP COMMAND ====================

func newOOPCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oop",
		Short: "🏗️ Object-Oriented Programming patterns",
	}

	cmd.AddCommand(
		newInheritanceCmd(),
		newPolymorphismCmd(),
		newEncapsulationCmd(),
		newAbstractCmd(),
		newDesignPatternCmd(),
	)

	return cmd
}

func newInheritanceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "inheritance",
		Short: "Generate class with inheritance",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🏗️ Inheritance Generator")
			fmt.Println("─────────────────────────")

			fmt.Print("\n? Parent class name: ")
			parent := readInput(reader)

			fmt.Print("? What does the parent represent: ")
			parentDesc := readInput(reader)

			fmt.Print("? Parent attributes (comma-separated): ")
			parentAttrs := readInput(reader)

			fmt.Print("? Parent methods (comma-separated, e.g., speak, move): ")
			parentMethods := readInput(reader)

			fmt.Print("? Child class name: ")
			child := readInput(reader)

			fmt.Print("? Extra attributes for child: ")
			childAttrs := readInput(reader)

			fmt.Print("? Does child override any parent method? (yes/no): ")
			overrides := readInput(reader)

			// Build parent
			pAttrs := strings.Split(parentAttrs, ",")
			var pInitParams []string
			var pInitBody []string
			for _, a := range pAttrs {
				a = strings.TrimSpace(a)
				if a == "" {
					continue
				}
				pInitParams = append(pInitParams, a)
				pInitBody = append(pInitBody, fmt.Sprintf("        self.%s = %s", a, a))
			}

			pMethods := strings.Split(parentMethods, ",")
			var methodCode string
			for _, m := range pMethods {
				m = strings.TrimSpace(m)
				if m == "" {
					continue
				}
				methodCode += fmt.Sprintf("\n\n    def %s(self):\n        \"\"\"Base %s behavior.\"\"\"\n        pass  # TODO: implement", m, m)
			}

			// Build child
			cAttrs := strings.Split(childAttrs, ",")
			allParams := append(pInitParams, []string{}...)
			var cInitBody []string
			for _, a := range cAttrs {
				a = strings.TrimSpace(a)
				if a == "" {
					continue
				}
				allParams = append(allParams, a)
				cInitBody = append(cInitBody, fmt.Sprintf("        self.%s = %s", a, a))
			}

			var overrideCode string
			if strings.ToLower(overrides) == "yes" {
				fmt.Print("? Which method to override: ")
				overMethod := readInput(reader)
				fmt.Print("? What should the child version do: ")
				overDesc := readInput(reader)
				overrideCode = fmt.Sprintf("\n\n    def %s(self):\n        \"\"\"%s\"\"\"\n        # TODO: %s\n        pass", overMethod, overDesc, overDesc)
			}

			code := fmt.Sprintf(`class %s:
    """%s"""

    def __init__(self, %s):
%s%s


class %s(%s):
    """Inherits from %s."""

    def __init__(self, %s):
        super().__init__(%s)
%s%s

# Usage:
# obj = %s(%s)`,
				parent, parentDesc,
				strings.Join(pInitParams, ", "),
				strings.Join(pInitBody, "\n"),
				methodCode,
				child, parent, parent,
				strings.Join(allParams, ", "),
				strings.Join(pInitParams, ", "),
				strings.Join(cInitBody, "\n"),
				overrideCode,
				child, strings.Join(allParams, ", "),
			)

			printGenerated(code)
		},
	}
}

func newPolymorphismCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "polymorphism",
		Short: "Generate polymorphism example",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🏗️ Polymorphism Generator")
			fmt.Println("──────────────────────────")

			fmt.Print("\n? Base class name (e.g., Shape, Animal, Vehicle): ")
			base := readInput(reader)

			fmt.Print("? Common method name (e.g., area, speak, start): ")
			method := readInput(reader)

			fmt.Print("? Child classes (comma-separated, e.g., Circle, Square, Triangle): ")
			children := readInput(reader)

			childList := strings.Split(children, ",")

			var classCode string
			classCode = fmt.Sprintf(`from abc import ABC, abstractmethod


class %s(ABC):
    """Abstract base — forces all children to implement %s()"""

    @abstractmethod
    def %s(self):
        pass

    def __repr__(self):
        return f"{self.__class__.__name__}"
`, base, method, method)

			for _, c := range childList {
				c = strings.TrimSpace(c)
				if c == "" {
					continue
				}
				classCode += fmt.Sprintf(`

class %s(%s):
    def %s(self):
        # TODO: implement %s for %s
        pass
`, c, base, method, method, c)
			}

			classCode += fmt.Sprintf(`

# Polymorphism in action:
def process_all(items: list[%s]):
    """Same method call, different behavior per type."""
    for item in items:
        print(f"{item}: {item.%s()}")

# Usage:
# shapes = [%s]
# process_all(shapes)`, base, method, strings.Join(childList, "(), ")+"()")

			printGenerated(classCode)
		},
	}
}

func newEncapsulationCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "encapsulation",
		Short: "Generate class with private attributes and properties",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🏗️ Encapsulation Generator")
			fmt.Println("───────────────────────────")

			fmt.Print("\n? Class name: ")
			className := readInput(reader)

			fmt.Print("? Private attributes (comma-separated, e.g., balance, password): ")
			attrs := readInput(reader)

			fmt.Print("? Which should have getter (read access)? (comma-separated or 'all'): ")
			getters := readInput(reader)

			fmt.Print("? Which should have setter (write access with validation)? (comma-separated or 'none'): ")
			setters := readInput(reader)

			attrList := strings.Split(attrs, ",")
			getterList := strings.Split(getters, ",")
			setterList := strings.Split(setters, ",")

			var initBody []string
			var propCode string
			var initParams []string

			for _, a := range attrList {
				a = strings.TrimSpace(a)
				if a == "" {
					continue
				}
				initParams = append(initParams, a)
				initBody = append(initBody, fmt.Sprintf("        self._%s = %s  # private", a, a))

				// Check if getter needed
				hasGetter := strings.ToLower(getters) == "all"
				if !hasGetter {
					for _, g := range getterList {
						if strings.TrimSpace(g) == a {
							hasGetter = true
							break
						}
					}
				}

				if hasGetter {
					propCode += fmt.Sprintf(`
    @property
    def %s(self):
        """Read-only access to %s."""
        return self._%s
`, a, a, a)
				}

				// Check if setter needed
				hasSetter := false
				for _, s := range setterList {
					if strings.TrimSpace(s) == a {
						hasSetter = true
						break
					}
				}

				if hasSetter {
					propCode += fmt.Sprintf(`
    @%s.setter
    def %s(self, value):
        """Validated write access to %s."""
        # TODO: add validation
        if value is None:
            raise ValueError("%s cannot be None")
        self._%s = value
`, a, a, a, a, a)
				}
			}

			code := fmt.Sprintf(`class %s:
    """Encapsulated class — private attributes with controlled access."""

    def __init__(self, %s):
%s
%s
# Usage:
# obj = %s(%s)
# print(obj.balance)   ← uses @property (getter)
# obj.balance = 100    ← uses @setter (with validation)
# print(obj._balance)  ← WORKS but BAD PRACTICE (bypasses validation)`,
				className,
				strings.Join(initParams, ", "),
				strings.Join(initBody, "\n"),
				propCode,
				className,
				strings.Join(initParams, ", "),
			)

			printGenerated(code)
		},
	}
}

func newAbstractCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "abstract",
		Short: "Generate abstract base class (ABC)",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🏗️ Abstract Class Generator")
			fmt.Println("────────────────────────────")

			fmt.Print("\n? Abstract class name (e.g., Database, PaymentGateway, Logger): ")
			className := readInput(reader)

			fmt.Print("? Abstract methods (comma-separated, e.g., connect, query, close): ")
			methods := readInput(reader)

			fmt.Print("? One concrete implementation name (e.g., PostgresDB, StripePayment): ")
			implName := readInput(reader)

			methodList := strings.Split(methods, ",")

			var abstractMethods string
			var implMethods string

			for _, m := range methodList {
				m = strings.TrimSpace(m)
				if m == "" {
					continue
				}
				abstractMethods += fmt.Sprintf(`
    @abstractmethod
    def %s(self):
        """Must be implemented by subclass."""
        pass
`, m)
				implMethods += fmt.Sprintf(`
    def %s(self):
        """Concrete implementation of %s."""
        # TODO: implement
        pass
`, m, m)
			}

			code := fmt.Sprintf(`from abc import ABC, abstractmethod


class %s(ABC):
    """
    Abstract base class — defines the interface.
    Cannot be instantiated directly.
    All subclasses MUST implement the abstract methods.
    """
%s

class %s(%s):
    """%s implementation of %s."""
%s

# Usage:
# db = %s()    ← ERROR! Cannot instantiate abstract class
# db = %s()   ← Works! All abstract methods implemented
# db.%s()`,
				className, abstractMethods,
				implName, className, implName, className,
				implMethods,
				className, implName,
				strings.TrimSpace(methodList[0]),
			)

			printGenerated(code)
		},
	}
}

func newDesignPatternCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "design-pattern",
		Short: "Common OOP design patterns",
		Run: func(cmd *cobra.Command, args []string) {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("\n🏗️ Design Pattern Generator")
			fmt.Println("────────────────────────────")

			fmt.Println("\n? Which pattern?")
			fmt.Println("  [1] Singleton (one instance only)")
			fmt.Println("  [2] Factory (create objects without specifying class)")
			fmt.Println("  [3] Observer (event system / pub-sub)")
			fmt.Println("  [4] Strategy (swap algorithms at runtime)")
			fmt.Println("  [5] Decorator pattern (add behavior dynamically)")
			fmt.Print("\n  Select: ")
			choice := readInput(reader)

			var code string

			switch choice {
			case "1":
				fmt.Print("\n? Class name for singleton: ")
				name := readInput(reader)
				code = fmt.Sprintf(`class %s:
    """
    Singleton — only ONE instance ever exists.
    Use for: DB connections, config managers, loggers.
    """
    _instance = None

    def __new__(cls, *args, **kwargs):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
            cls._instance._initialized = False
        return cls._instance

    def __init__(self):
        if self._initialized:
            return
        self._initialized = True
        # TODO: initialize your singleton

# Usage:
# a = %s()
# b = %s()
# assert a is b  # True — same object!`, name, name, name)
			case "2":
				fmt.Print("\n? Base product name (e.g., Notification, Payment, Database): ")
				base := readInput(reader)
				fmt.Print("? Product types (comma-separated, e.g., Email, SMS, Push): ")
				types := readInput(reader)

				typeList := strings.Split(types, ",")
				var classes string
				var factoryCases string
				for _, t := range typeList {
					t = strings.TrimSpace(t)
					if t == "" {
						continue
					}
					classes += fmt.Sprintf(`
class %s%s(%s):
    def execute(self):
        return "%s %s sent"
`, t, base, base, t, strings.ToLower(base))
					factoryCases += fmt.Sprintf(`        "%s": %s%s,
`, strings.ToLower(t), t, base)
				}

				code = fmt.Sprintf(`from abc import ABC, abstractmethod


class %s(ABC):
    @abstractmethod
    def execute(self):
        pass
%s

class %sFactory:
    """Factory — creates the right %s without you knowing the class."""

    _registry = {
%s    }

    @classmethod
    def create(cls, type_name: str) -> %s:
        klass = cls._registry.get(type_name.lower())
        if not klass:
            raise ValueError(f"Unknown type: {type_name}")
        return klass()

# Usage:
# notif = %sFactory.create("email")
# notif.execute()`, base, classes, base, strings.ToLower(base), factoryCases, base, base)
			case "3":
				code = `class EventSystem:
    """
    Observer / Pub-Sub — components react to events without tight coupling.
    Use for: UI updates, notifications, logging, webhooks.
    """

    def __init__(self):
        self._listeners = {}  # event_name → [callbacks]

    def on(self, event: str, callback):
        """Subscribe to an event."""
        if event not in self._listeners:
            self._listeners[event] = []
        self._listeners[event].append(callback)

    def off(self, event: str, callback):
        """Unsubscribe from an event."""
        if event in self._listeners:
            self._listeners[event].remove(callback)

    def emit(self, event: str, *args, **kwargs):
        """Fire an event — all subscribers get notified."""
        for callback in self._listeners.get(event, []):
            callback(*args, **kwargs)

# Usage:
events = EventSystem()

def on_user_signup(user):
    print(f"Welcome email sent to {user}")

def on_user_signup_log(user):
    print(f"[LOG] New user: {user}")

events.on("signup", on_user_signup)
events.on("signup", on_user_signup_log)
events.emit("signup", "chiru@example.com")
# Both handlers fire!`
			case "4":
				fmt.Print("\n? Context name (e.g., Sorter, Compressor, PaymentProcessor): ")
				context := readInput(reader)
				fmt.Print("? Strategy variations (comma-separated, e.g., Quick, Merge, Bubble): ")
				strategies := readInput(reader)

				stratList := strings.Split(strategies, ",")
				var stratClasses string
				for _, s := range stratList {
					s = strings.TrimSpace(s)
					if s == "" {
						continue
					}
					stratClasses += fmt.Sprintf(`
class %sStrategy(Strategy):
    def execute(self, data):
        # TODO: implement %s strategy
        return data
`, s, s)
				}

				code = fmt.Sprintf(`from abc import ABC, abstractmethod


class Strategy(ABC):
    @abstractmethod
    def execute(self, data):
        pass
%s

class %s:
    """Context — uses a strategy that can be swapped at runtime."""

    def __init__(self, strategy: Strategy):
        self._strategy = strategy

    def set_strategy(self, strategy: Strategy):
        """Change strategy dynamically."""
        self._strategy = strategy

    def process(self, data):
        return self._strategy.execute(data)

# Usage:
# processor = %s(%sStrategy())
# result = processor.process(data)
# processor.set_strategy(%sStrategy())  # swap!
# result = processor.process(data)`, stratClasses, context, context, strings.TrimSpace(stratList[0]), strings.TrimSpace(stratList[1]))
			case "5":
				code = `from functools import wraps


def log_calls(func):
    """Decorator: logs every call with args and return value."""
    @wraps(func)
    def wrapper(*args, **kwargs):
        print(f"→ {func.__name__}({args}, {kwargs})")
        result = func(*args, **kwargs)
        print(f"← {func.__name__} returned: {result}")
        return result
    return wrapper


def retry(max_attempts=3):
    """Decorator: retries function on failure."""
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            for attempt in range(1, max_attempts + 1):
                try:
                    return func(*args, **kwargs)
                except Exception as e:
                    if attempt == max_attempts:
                        raise
                    print(f"Attempt {attempt} failed: {e}")
        return wrapper
    return decorator


def cache_result(func):
    """Decorator: memoize results (manual lru_cache)."""
    cache = {}
    @wraps(func)
    def wrapper(*args):
        if args in cache:
            return cache[args]
        result = func(*args)
        cache[args] = result
        return result
    return wrapper


# Usage:
@log_calls
@retry(max_attempts=3)
def call_api(url):
    # TODO: make API call
    return {"status": "ok"}

@cache_result
def expensive_calc(n):
    return n ** 2 + n ** 3`
			default:
				fmt.Println("❌ Invalid selection")
				return
			}

			printGenerated(code)
		},
	}
}
