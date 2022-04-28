# Evaluate Reverse Polish Notation

Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, and /. Each operand may be an integer or another expression.

Note that division between two integers should truncate toward zero.

It is guaranteed that the given RPN expression is always valid. That means the expression would always evaluate to a result, and there will not be any division by zero operation.

 
## Examples
```
Example 1:

Input: tokens = ["2","1","+","3","*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
Example 2:

Input: tokens = ["4","13","5","/","+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
Example 3:

Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
Output: 22
Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
 

Constraints:

1 <= tokens.length <= 104
tokens[i] is either an operator: "+", "-", "*", or "/", or an integer in the range [-200, 200].
```

# Analysis

## What is [Reverse Polish Notation](http://en.wikipedia.org/wiki/Reverse_Polish_notation).

Reverse Polish Notation 是一種把運算元放在運算子之後的表示式

舉例來說： 3, 4, +  ⇒ 代表  4 + 3

 

## Requests

這邊題目是給定一個 Reverse Polish Notation 的 string list 

要去實作把這個 string list 計算出根據 Reverse Polish Notation 解析出值的方法

## 實作需要解決的問題

1 需要判別輸入的值是運算元還是運算子

2 需要把運算元轉成整數(輸入是字串)

3 題目說明這些運算子是針對整數的運算 也就是說 3/4 =  0

4 決定何時做運算

## 思考

假設給定一個 Reverse Polish Notation 表示式該如何計算數值

可以觀察這裡的運算子 `"+"`, `"-"`, `"*"`, `"/"`  都是二元運算子

代表需要兩個運算元

且運算子一定在運算元之後

符合以下 pattern： 

> **operand_2** **operand_1** **operator**
> 

所以每次當遇到 operator時 才把 兩個operand 做運算

並且注意到 operand 的順序是先遇到的放後面 後進先出

也就是說 3, 4, - ⇒ 4 - 3 而不是 3 - 4

因為後進先出 所以可以透過 stack 來實作存放運算元的 storage

## 實作演算法

使用一個 stack 來儲存所有運算元

step 1: 先檢查輸入 string list tokens 長度是否大於 1

step 2: 如果 長度 == 1 則直接把該值轉為整數回傳

step 3: 如果 長度 > 1

           初始化 index = 0, result = 0 

step 4: 當 index < tokens 長度時 , 檢查 tokens[index] 是否是運算子

step 5: 如果 tokens[index] 是運算子:

            從 stack 中 pop 出兩個運算元

            更新result 為 兩個運算元以運算子做計算  

            把 result 放入 stack 

step 6:  如果 tokens[index] 是運算元:

             把運算子轉成整數放入 stack

step 7: 更新 index = index +1 , 回到 step 4  

step 8: 回傳 result   

## 程式碼

```go
import "strconv"
func evalRPN(tokens []string) int {
    // stack store operand
    if len(tokens) == 1 {
        value, _ := strconv.Atoi(tokens[0])
        return value
    }
    stack := []int{}
    result := 0
    operators := map[string]int{"+": 0, "-": 1, "*": 2, "/": 3}
    for _, input := range tokens {
        if opCode, isOp := operators[input]; isOp {
            topIdx := len(stack) -1
            // pop 2 element
            second := stack[topIdx]
            first := stack[topIdx-1]
            stack = stack[0: topIdx-1]
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
            value, _ := strconv.Atoi(input)
            stack = append(stack, value)
        }
    }
    return result
}
```

## 困難點

1 需要對 [Reverse Polish Notation](http://en.wikipedia.org/wiki/Reverse_Polish_notation) 有理解

2 需要對於程式語言中 string 轉換成 integer 有理解

# Solve Point

- [x]  Understand What is [Reverse Polish Notation](http://en.wikipedia.org/wiki/Reverse_Polish_notation).
- [x]  Figure out what is the request
- [x]  Convert string to Integer