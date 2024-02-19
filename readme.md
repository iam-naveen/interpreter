## A Programming Language aiming to make programming Easy

### This language is based on tamil words and sentance stucture
### It contains tanglish characters which break the barier of typing in tamil
<hr>
### Syntax of the langauge
<hr>

## Variables
```
yen a = 10
sol name = "naveen"
```

## Conditional Statement
```
yen c = 0
a < b irundha {
    c = a + b
}
illana a > b irundha {
    c = a - b
}
illana {
    c sollu
}
```

## Loops
```

// loop iterates for 10 times
10 murai {
    "Hello" sollu
}

// prints the given string a*10 times
// each iteration a is decremented by one
yen a = 0;
a*10 murai a-- {
    "variable a = {a}" sollu
}

// basically a while loop,
// runs while the expression is true
a >= b irukum varai {
    a -= 1
}

// similar to while loop,
// but loops until the expression becomes true
a < b aagum varai {
    "running..." sollu
}
```

## Functions
```
// declare the function
add seiyal | yen a, yen b | yen {
    yen c = a + b
    c ->
}

// calling the function
(1, 2 -> add) sollu
```

## Other Control Flows
```
yen a = 100

// similar to switch statements
// matches the given variable to the case patterns
indha a {
  10 bothu {
    "hello" sollu
  }
  11 bothu {
    "hi" sollu
  }
}
```
