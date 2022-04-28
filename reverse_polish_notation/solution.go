package reverse_polish_notation

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 1 {
		value, _ := strconv.Atoi(tokens[0])
		return value
	}
	// use stack to parse
	operators := map[string]int{"+": 0, "-": 1, "*": 2, "/": 3}
	stack := []int{}
	result := 0
	for _, value := range tokens {
		if opCode, isOp := operators[value]; isOp {
			topIdx := len(stack) - 1
			second := stack[topIdx]
			first := stack[topIdx-1]
			stack = stack[0 : topIdx-1]
			switch opCode {
			case 0:
				result = first + second
			case 1:
				result = first - second
			case 2:
				result = first * second
			default:
				result = first / second
			}
			stack = append(stack, result)
		} else {
			val, _ := strconv.Atoi(value)
			stack = append(stack, val)
		}
	}
	return result
}
