# monkey

[![Go](https://github.com/gustavodiasag/monkey/actions/workflows/go.yml/badge.svg)](https://github.com/gustavodiasag/monkey/actions/workflows/go.yml)

Interpreter for the [Monkey](https://interpreterbook.com/) programming language, just as an exercise to learn Go.

This implementation specifies its own lexer and follows a really common top-down style of parsing using [Pratt's technique](https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html), along with an AST-based approach for representing parsed code, whose evaluation is done by converting the representation to an internal object system. The project also aims to explore Test-Driven Development, where failing unit tests are constructed before the actual targeted implementation.

# License

The project is licensed under the [MIT License](LICENSE).
