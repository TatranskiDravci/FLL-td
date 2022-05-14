# Styleguide

## Parentheses
Parentheses after function calls are to be placed right after function call without any spaces
```c
functionName(parameters);
```
Parentheses after control statements (`if`, `while`, `for`, `switch`) are to be separated from these statements by a single space
```c
if (something) ...;
while (something) ...;
for (something) ...;
switch (something) ...;
```

## Code blocks
Two styles are used for blocks of code (enclosed by the curly braces `{}`). For if-else clauses and loops with only one call within, single-line style is used, (for the if-else clauses) usually with the matching calls indented in such way, that they start at the same column
```c
if (something)           call(...);
else if (something else) call(...);
else                     call(...);
```
In this case, the parameters should, if possible, be also aligned by the column. This isn't necessarily the case for loops or if-else clauses with different calls

```c
if (something) call(...);
else if (something else) otherCall(...);
else yetAnotherCall(...);
```
```c
while (something) call(...);
for (something) call(...);
```
When a single line statement isn't sufficient, or a function/struct needs to be defined, the Allman style is used
```c
if (something)
{
    call1(...);
    call2(...);
}
else
{
    call3(...);
    call4(...);
}
```
```c
while (something)
{
    call1(...);
    call2(...);
}
```
```c
type functionName(type parameter1, type parameter2)
{
    // do something
}
```
placing the curly brace on a new line.

These styles can be mixed and matched
```c
if (something)
{
    call1(...);
    call2(...);
}
else if (something else) call(...);
else                     call(...);
```
## Naming
### Functions
To name functions, `camelCase` is used. The first character of the name is to be lower cased. Any other word, appended to the name is to have its first character capitalized. When naming functions, tightly related to specific types, the name of the type should be the first word to appear in the function name.

Suppose a type `foo` is declared, then the constructor for `foo` is to be declared as
```c
foo fooNew(...);
```
and a function doing action `get bar` (returning type `bar`) on `foo` should be declared as
```c
bar fooGetBar(foo f, ...);
```
taking some `foo` (or a pointer to some `foo`) as its first parameter.
### Structs and types
Ideally, the struct and type names be single words. If such is not possibl, the struct name should always be expressed in a `PascalCase`, having the first character in its name capitalized. The corresponding type name should be then expressed as a `camelCase` of the struct name (to comply with function naming scheme).

Suppose a foo bar structure is to be defined. With the given conventions, it would be defined as such
```c
typedef struct FooBar      // struct name
{
    // fields
}
fooBar;                    // type name
```

### Variables and fields
Variables and fields are perhaps the most disorganized category in this code. There are only three rules (but mostly just recommendations), really,
 - variables must not be named using `camelCase`, as one could accidentally name a variable identically to an already defined function,
 - when using single letter names, the meaning of the variable must be either painfully obvious from the context or must be explained using a comment,
 - for longer variable names with multiple words, `snake_case` should be used (e.g. `very_long_variable_name_please_try_to_avoid_these`).
