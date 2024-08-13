# WordGenerator

## Description

This Golang project, WordGenerator, is designed to generate all possible permutations of a given set of characters and saves them into a SQLite database. It's implemented in Go and provides a robust solution for tasks requiring the generation of all possible combinations of a set of characters, such as in cryptographic applications, testing, or combinatorial analysis.

## Features

- Generate permutations for any set of characters.
- Stores permutations in a SQLite database.
- Command line based, easy to use.
- Highly customizable for different user needs.

## Requirements

- Go programming language
- Task Quick command manager

## Installation

1. Ensure Go is installed on your system. You can download and install it from [Go's official site](https://golang.org/dl/).
2. Ensure you have Task Quick command manager installed. Installation instructions can be found [here](https://taskfile.dev/#/installation).
3. Clone this repository or download the source code to your local machine.

## Usage

1. Navigate to the directory where the project is located.
2. To set up the database, run the following command:
    ```sh
    task quick:migrate
    ```
    This command migrates and creates the database file needed for storing the permutations.
3. To run the generator, execute:
    ```sh
    go run .
    ```
    This will start the generation process based on the parameters set in the source code.

## Customization

- You can customize the set of characters for which permutations are generated by editing the `main.go` file. Modify the `characters` slice to include your desired characters.
- Adjust other configuration settings directly in the code to fit your specific requirements, such as length or database configurations.

## Database Structure

- The SQLite database file generated will contain a table named `words`.
- The `words` table has the following structure:
    - `id` INTEGER PRIMARY KEY AUTOINCREMENT: a unique identifier for each permutation.
    - `word` TEXT: the generated permutation.

## Create Index for faster search
```
CREATE INDEX word_name ON words (name)
```

## Example

Here is a quick example of how to use the project:

1. Set your desired characters in `main.go`:
    ```go
    characters := "abcdefghijklmnopqrstuvwxyz"
    ```
2. Run the migration command:
    ```sh
    task quick:migrate
    ```
3. Execute the generator:
    ```sh
    go run .
    ```
4. Check the database for the generated permutations.

## Contributing

Contributions are welcome! If you would like to contribute to this project, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This project is inspired by similar tools in other programming languages, like Python, enhancing them with the performance benefits of Go.
- Uses the SQLite database for efficient data handling.

## Contact

If you have any questions or suggestions, please feel free to open an issue or contact the repository owner.

---
