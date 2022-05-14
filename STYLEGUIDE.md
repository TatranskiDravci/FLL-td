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
else if (something_else) call(...);
else                     call(...);
```
In this case, the parameters should, if possible, be also aligned by the column. This isn't necessarily the case for loops or if-else clauses with different calls

```c
if (something) call(...);
else if (something_else) otherCall(...);
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
## Comments
### Simple comments
Comments describing a block of calls should be placed a line above that "block", separated by an empty line with the previous. Comments, describing effects of specific calls should be placed besides the corresponding call and should be aligned in the same column. There should always be (at least) a single space between `//` and the actual comment.
```c
something(...);
somethingElse(...);

// description of the call block
something(...);                    // description of the effect of `something`
somethingElse(...);                // description of the effect of `somethingElse`
whatever(...);

```
### Multi-line comments
When describing more complicated behaviour, multi-line comments can be used. A multi-line comment should be structured as such
```c
/*
    general description of some action
    can be written in sentences, if too long,
    should be divided into multiple lines
        - lists
        - can be
        - utilized too
*/
```
The `/* */` marks should always be placed at the current indentation level. 1 soft-tab should be used for the description, 2 soft-tabs should be used for lists.

### Commenting out function calls
When commenting out function calls, two slashes with one space (`// `) must be used. This is because many text editors support `Ctrl + /` shortcut for quick commenting and uncommenting selected calls. This feature sadly doesn't work with multi-line comments, thus single line comments should always be used in this case.
```c
// inactive block
// something(...);
// somethingElse(...);

// description of the call block
something(...);                    // description of the effect of `something`
somethingElse(...);                // description of the effect of `somethingElse`
whatever(...);
```
### Documenting functions
When writing a reference comment for a defined function, the following template should be used
```c
/*
    function description (what it does)
        parameter1 - description of parameter1
        parameter2 - description of parameter2
        ...
        @return    - description of the returned value (can be left out if the return type is void)
*/
```
Suppose a function for getting `bar` from `foo` is defined (`foo` being its only parameter), then the definition in some `.h` file with the reference comment would be
```c
/*
    gets bar from foo and returns it
        f       - foo object
        @return - bar object
*/
bar fooGetBar(foo f);

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
