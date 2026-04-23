# Comments

Riya supports single-line comments only. Comments are used to annotate code for readability and are ignored by the compiler.

## Lexical Rules
- A comment starts with a *#* and ends with a *newline* or *EOF*
- Anything between, including a *#*, is ignored by the compiler(lexer)

## Syntax and Placement Rules
- Full-line comments
```
# this is a full line comment
hello();

```
- Inline comments
```
hello();  # this is an inline comment

```