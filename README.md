# Challenge 1

**Difficulty**: Easy

Write a function that will take two numbers, calculate the sum of all numbers between them (inclusive), and divide that sum by the sum of the two entered numbers.

---

### Example

For the numbers **4** and **10**:

1. Calculate the sum of all numbers between 4 and 10:

   ```
   4 + 5 + 6 + 7 + 8 + 9 + 10 = 49
   ```

2. Divide this sum by the sum of the two entered numbers:

   ```
   49 / (4 + 10)
   ```

3. The result will be:

   ```
   3.5
   ```

---

# Challenge 2

**Difficulty**: Easy

Write a function that takes a number `n` and returns an output like this:

The `generatePattern` function takes an integer `n` as input. It first checks if `n` is greater than 10, and if so, it prints an error message:

If `n` is less than 10, it generates and prints the following pattern:

---

### Example

For the number `8`, the output will be:\
12345678\
1234567\
123456\
12345\
1234\
123\
12\
1

---

# Challenge 3

**Dificullty**: Easy

In this challenge, we are going to define a function that is called GetDeveloper which will take in 2 interface{} arguments.

Within this function, you will have to declare a new Developer instance and use type assertion to populate the values correctly before then returning this new Developer instance.

---

# Challenge 4

**Difficulty**: Easy

You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N. Write a method that takes the array as an argument and returns this "outlier" N.

---

### Example

[2, 4, 0, 100, 4, 11, 2602, 36] --> 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21] --> 160 (the only even number)

---

# Challenge 5

**Difficulty**: Easy

Create a function that returns the name of the winner in a fight between two fighters.

Each fighter takes turns attacking the other and whoever kills the other first is victorious. Death is defined as having health <= 0.

Each fighter will be a Fighter object/instance. See the Fighter class below in your chosen language.

Both health and damagePerAttack (damage_per_attack for python) will be integers larger than 0. You can mutate the Fighter objects.

Your function also receives a third argument, a string, with the name of the fighter that attacks first.

---

### Example

declare_winner(Fighter("Lew", 10, 2), Fighter("Harry", 5, 4), "Lew") => "Lew"

Lew attacks Harry; Harry now has 3 health.\
Harry attacks Lew; Lew now has 6 health.\
Lew attacks Harry; Harry now has 1 health.\
Harry attacks Lew; Lew now has 2 health.\
Lew attacks Harry: Harry now has -1 health and is dead. Lew wins.

---

# Challenge 6

**Difficulty**: Easy

Complete the solution so that it reverses the string passed into it.

---

### Example

'world' => 'dlrow'
'word' => 'drow'

---

#### References:

1. https://coddy.tech/courses/golang_challenges__level_1/sum_and_divide
2. https://coddy.tech/courses/golang_challenges__level_1/generate_pattern
3. https://tutorialedge.net/challenges/go/type-assertions/
4. https://www.codewars.com/kata/5526fc09a1bbd946250002dc/train/go
5. https://www.codewars.com/kata/577bd8d4ae2807c64b00045b/train/go
6. https://www.codewars.com/kata/5168bb5dfe9a00b126000018/go
