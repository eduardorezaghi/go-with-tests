# Go with Tests

This is a repository to learn Go with tests.  
The idea is to learn Go by writing tests first and then writing the code to make the tests pass (TDD).

```
go-with-tests/
├── hello-go/
│   ├── go.mod
│   ├── hello.go
│   └── hello_test.go
└── README.md
```

## Table of Contents
- [Go with Tests](#go-with-tests)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
  - [Running the tests](#running-the-tests)
  - [References](#references)


## Getting Started
Install Go on your machine.
```bash
sudo apt install golang-go # Debian/Ubuntu
sudo pacman -S go # Arch Linux
sudo dnf install golang # Fedora
nix-env -i go # NixOS
nix profile install go # Nix Package Manager
```

Check if Go is installed correctly.
```bash
go version
```

Run a sample Go program.
```bash
go run hello-go/hello.go
# Hello, Go!
```

## Running the tests
To run the tests, navigate to the directory of the test and run the following command.
```bash
cd hello-go; go test
```


## References
If you stumbled upon this repository and found it interesting, please take a look on the original material.
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)
- [Go by Example](https://gobyexample.com/)
- [Go Documentation](https://golang.org/doc/)