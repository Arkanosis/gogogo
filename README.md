# gogogo [![](https://img.shields.io/badge/go-v0.0.1-green.svg)](https://go.arkanosis.net/gogogo) [![License](https://img.shields.io/badge/license-ISC-blue.svg)](/LICENSE) [![Build status](https://travis-ci.org/Arkanosis/gogogo.svg?branch=master)](https://travis-ci.org/Arkanosis/gogogo)

**gogogo** is a command line utility to boostrap a project.

## Features

gogogo is still under active design, and not yet ready for mainstream usage.

However, it aims at doing the following:
 - create a new directory for the project;
 - initialize git versioning in it;
 - create a bare git repository on GitHub, GitLab or whatever forge the user wants to use;
 - add the forge repository as remote in the local repository;
 - instantiate a project template, depending on the language (golang, rust, python…).
 
Additionaly, for:
 - golang projects:
  - link to the project directory from `$GOPATH/src/$PROJECT_URI`;
  - create a `go-import` redirect from `$PROJECT_URI` to the GitHub, GitLab… repository;
 - rust projects:
  - publish a boostrap release on crates.io;
 - python projects:
  - publish a boostrap release on pypi.org.

## Installation

### On any distribution with golang available

```bash
go get go.arkanosis.net/gogogo
```

## Usage

```
Usage: gogogo <project>
       gogogo -h | --help
       gogogo --version

Arguments:
    project        Project name.

Options:
    -h, --help    Show this screen.
    --version     Show version.
```

## Compiling

Run `go build` in your working copy.

## Contributing and reporting bugs

Contributions are welcome through [GitHub pull requests](https://github.com/Arkanosis/gogogo/pulls).

Please report bugs and feature requests on [GitHub issues](https://github.com/Arkanosis/gogogo/issues).

## License

gogogo is copyright (C) 2018 Jérémie Roquet <jroquet@arkanosis.net> and
licensed under the ISC license.
