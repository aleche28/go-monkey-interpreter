# Sparse notes while reading the book

What we are goin to build in thi book is a so called 'tree-walking' interpreter, meaning that the interpreter parses the source code and builds an abstract syntax tree (AST), then evalutes the tree. Called 'tree-walking' because it walks the AST and interprets it.

going to build:
    - lexer
    - parser
    - tree representation
    - evaluator

without using pre-existing ones.

Monkey lang features:
- C-like syntax;
- variable bindings
- integers and booleans
- arithmetic expr
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

```Monkey
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
let myArray = [1, 2, 3, 4, 5];
let thorsten = { "name": "Thorsten", "age": 28 }

myArray[0] // accessing element in array
thorsten["name"] // accessing element in hash


// let used to bind functions to names
let add = fn(a, b) { return a + b; };

// implicit return values
let add = fn(a, b) { a + b; };

// call function
add(1, 2);

// more complex function
let fibonacci = fn(x) {
    if (x == 0) {
        0 // impl. return
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

// support higher order functions, that take other functions as arguments
let twice = fn(f, x) {
    return f(f(x));
};

let addTwo = fn(x) {
    return x + 2;
};


twice(addTwo, 2); // => 6
```

The interpreter will tokenize and parse Monkey source code in a REPL, building up an internal representation of the code called AST and then evaluate this tree. Major parts:
- lexer
- parser
- AST
- internal object system
- evaluator

We're gonna build them in this order, bottom up.

## Steps

Source code -1-> Tokens -2-> AST

1 - lexical analysis (lexer, scanner, tokenizer)
2 - parsing (done by the parser)

In Monkey language, spaces are not important, so we don't care how many spaces are between tokens, and spaces are not treated as tokens.

Also, we won't add line, col or filename to the tokens for simplicity, but it can be done to have precise error messages.

## 1.2

let's start with a subset:

```Monkey
let five = 5;
let ten = 10;

let add = fn(x, y) {
    x + y;
};

let result = add(five, ten);
```

We create the token type with two fields: Type, which is of type TokenType, and Literal, that contains the value of the token.

there are some tokens that are constant and do not change, so we can actually define them as const:

```go
const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"
    ...
)
```

## 1.3

The lexer doesn't need to buffer or store tokens, cause it will just output the next token it reads when the function NextToken() is called by the consumer (in our case the parser).
So the lexer is initialized with the input text and it will work with a loop that calls the NextToken() until EOF is reached (or error encountered, if any).


Note that our implementation only supports ASCII characters, not UTF-8. To support it we would need to read runes instead of bytes. That can be a possible improvement for later.


