# CLI Password Generator
[![Commits since last release](https://img.shields.io/github/commits-since/nico-castell/genpass/latest?label=Commits%20since%20last%20release&color=informational&logo=Git&logoColor=white&style=flat-square)](https://github.com/nico-castell/genpass/commits)
[![Release](https://img.shields.io/github/v/release/nico-castell/genpass?label=Release&color=informational&logo=GitHub&logoColor=white&style=flat-square)](https://github.com/nico-castell/genpass/releases)
[![License](https://img.shields.io/github/license/nico-castell/genpass?label=License&color=informational&logo=Open%20Source%20Initiative&logoColor=white&style=flat-square)](LICENSE)
[![Lines of code](https://img.shields.io/tokei/lines/github/nico-castell/genpass?label=Lines%20of%20code&color=informational&logo=Go&logoColor=white&style=flat-square)](https://github.com/nico-castell/genpass)
[![CodeQL](https://img.shields.io/github/workflow/status/nico-castell/genpass/CodeQL?label=CodeQL&logo=GitHub%20Actions&logoColor=white&style=flat-square)](https://github.com/nico-castell/genpass/actions/workflows/codeql-analysis.yml)

This is a cryptographically secure password generator for the command line, built using [***Golang***](https://golang.org/).

The program generates cryptographically secure numbers that transform into the characters in the output.

Usage:
- **Standalone:** You can run it without arguments, and it will default to an 8 character password, or you can specify the length as follows:
    ```shell
    genpass 12
    ```
  If you input a negative number (such as `-128`), the program will just flip it to be positive (it would become `128`).
- **Variable:** This is the recommended way. You store the output in a variable, so the password never appears on screen.
    ```shell
    MY_PASS=$(genpass 12)
    ```

---
To get the program run the following command:
```shell
go get github.com/nico-castell/genpass
```

Alternatively, you can clone the repository and install it from there:
```shell
git clone https://github.com/nico-castell/genpass.git  && \
    cd genpass                                         && \
    go install .
```

---
## About
This program and this repository are availabe under an [MIT License](LICENSE).
