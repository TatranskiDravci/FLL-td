# Styleguide

## Parentheses
Parentheses after function calls are to be placed right after function call without spaces
```c
functionName(parameters);
```
Parentheses after control statements (`if`, `while`, `for`, `switch`) are to be separated from these statements by a single space
```c
if (something) ...;
while (something) ...;
for (something) ...;
switch (comething) ...;
```

## Code blocks
Two styles are used for blocks of code (enclosed by the curly braces `{}`). For if-else clauses and loops with only one call within, single-line style is used, (for the if-else clauses) usually with the matching calls indented in such way, that they start at the same column
```c
if (something)           call(parameters);
else if (something else) call(parameters);
else                     call(parameters);
```
This isn't necessarily the case for loops or if-else clauses with different calls

```c
if (something) call(parameters);
else if (something else) otherCall(parameters);
else yetAnotherCall(parameters);
```
```c
while (something) call(parameters);
for (something) call(parameters);
```
When a single line statement isn't sufficient, or a function needs to be defined, the Allman style is used
```c
if (something)
{
    call1(parameters);
    call2(parameters);
}
else
{
    call3(parameters);
    call4(parameters);
}
```
```c
while (something)
{
    call1(parameters);
    call2(parameters);
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
    call1(parameters);
    call2(parameters);
}
else if (something else) call(parameters);
else                     call(parameters);
```
