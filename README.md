# 2. Progression

Write a program that finds the `n`th term of the progression

<img src="https://render.githubusercontent.com/render/math?math=\begin{cases}a_0 = 0\\a_1 = 1\\a_i = 3a_{i-2}%2b2a_{i-1}\end{cases}">
   
1. Write the program in the directory `progression`.
1. Implement the logic in a function declared as `func nthTerm(n uint) int`
1. Implement `main()` function that
   1. reads `n` from the keyboard
   1. calls `nthTerm` function
   1. prints a value returned from the function.
1. Create a file with tests for the program. The tests must cover all possible
   program scenarios.
   
## Example 1
   
```
nthTerm(1) -> 1
```
   
**Explanation**:

`nthTerm(1)` must return `1`, because <img src="https://render.githubusercontent.com/render/math?math=a_1 = 1">.
 
## Example 2
 
```
nthTerm(4) -> 20
```
    
The function must return `4`, because
    
<img src="https://render.githubusercontent.com/render/math?math=\begin{align*}a_0 %26= 0\\a_1 %26= 1\\ a_2 %26 = 3a_0 %2b 2a_1 = 3 \cdot 0 %2b + 2 \cdot 1 = 2\\a_3 %26 = 3a_1 %2b 2a_2 = 3 \cdot 1 %2b 2 \cdot 2 = 7\\a_4 %26 = 3a_2 %2b 2a_3 = 3 \cdot 2 %2b 2 \cdot 7 = 20\end{align*}">
