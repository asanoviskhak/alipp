# Alipp

Alipp `(from alippe (алиппе) - translates from Kyrgyz as "alphabet")` is a new programming language that combines the power of JavaScript with the beauty of the Kyrgyz language. With its familiar syntax and rich features, Alipp aims to make programming accessible to Kyrgyz-speaking developers.

## Features

- **JavaScript-like Syntax**: Alipp adopts a syntax similar to JavaScript, making it easy for developers familiar with JavaScript to transition to Alipp.

- **Kyrgyz Language Support**: Alipp is designed to support the Kyrgyz language, allowing developers to write code using Kyrgyz keywords, variable names, and comments.

- **Interoperability**: Alipp seamlessly integrates with existing JavaScript code, allowing developers to leverage the vast ecosystem of JavaScript libraries and frameworks.

- **Easy to Learn**: With its intuitive syntax and clear documentation, Alipp is beginner-friendly, making it a great choice for aspiring Kyrgyz-speaking programmers.

## Getting Started

To start using Alipp with REPL, follow these simple steps:

1. Pull this repository to your local machine by running the following command:

    ```
    git clone git@github.com:asanoviskhak/alipp.git
    ```

2. Run main.go file:

    ```
    go run main.go
    ```

3. Write your Alipp code using the familiar JavaScript syntax, but with Kyrgyz keywords and variable names.

      ```
      сакта облустарСаны = 7;
      ```
Currently, the REPL will only show the corresponding tokens of the input code. The next steps will be to implement the parser and the interpreter to run the code.

(coming soon) 4. Compile your Alipp code to JavaScript by running the following command:

    ```
    TOWRITE
    ```

(coming soon) 5. Use the generated JavaScript file in your projects, just like any other JavaScript file.

## Example

Here's a simple "Hello, World!" program written in Alipp:

```kyrgyz
көрсөтүү("Салам, Дүйнө!");
```

When compiled to JavaScript, it becomes:

```javascript
console.log("Салам, Дүйнө!");
```

## Contributing

We welcome contributions from the Kyrgyz programming community. If you have any ideas, bug reports, or feature requests, please open an issue on our [GitHub repository](https://github.com/asanoviskhak/alipp).

## License

Alipp is released under the [MIT License](https://opensource.org/licenses/MIT).
