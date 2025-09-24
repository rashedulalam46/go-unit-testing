Go Unit Testing in Go

This repository demonstrates foundational unit testing practices in Go (Golang). It includes examples of basic arithmetic operations with a Calculator struct, showcasing how to implement and test methods using Go's built-in testing package.

📂 Project Structure

```
go-unit-testing/
├── calculator.go          # Calculator struct and methods
├── calculator_test.go     # Unit tests for Calculator methods
├── main.go                # Entry point to run the application
├── go.mod                 # Go module definition
└── README.md              # Project documentation
```
🧮 Calculator Implementation

The Calculator struct provides basic arithmetic operations:

```
type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
    return a + b
}

func (c *Calculator) Subtract(a, b int) int {
    return a - b
}

func (c *Calculator) Multiply(a, b int) int {
    return a * b
}

func (c *Calculator) Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

## 🧪 Unit Tests

Unit tests are located in calculator_test.go. They cover various scenarios for each arithmetic method, ensuring correctness and handling edge cases.

## 🚀 Running the Application

To execute the application:

1. Clone the repository:
```
git clone https://github.com/rashedulalam46/go-unit-testing.git
cd go-unit-testing
```
2.  Run the application:

```
   go run main.go
```

## 🧪 Running Tests

To run the unit tests:

```
go test -v
```
For a more detailed output:

```
go test -v ./...

```

## 📋 Example Output

```
=== RUN   TestCalculator
=== RUN   TestCalculator/Add
=== RUN   TestCalculator/Subtract
=== RUN   TestCalculator/Multiply
=== RUN   TestCalculator/Divide
--- PASS: TestCalculator (0.00s)
    --- PASS: TestCalculator/Add (0.00s)
    --- PASS: TestCalculator/Subtract (0.00s)
    --- PASS: TestCalculator/Multiply (0.00s)
    --- PASS: TestCalculator/Divide (0.00s)
PASS
ok  	go-unit-testing	0.001s
```

## 🔧 Dependencies

Go 1.18 or higher

No external libraries are required; uses Go's standard testing package.

## 📝 License

This project is open-source and available under the MIT License.





