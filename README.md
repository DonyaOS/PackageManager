# Donya Package Manager

[![MIT License](https://img.shields.io/github/license/DonyaOS/PackageManager?color=brightgreen&style=flat-square)](LICENSE)
[![GitHub Linter Workflow Status](https://img.shields.io/github/workflow/status/DonyaOS/PackageManager/Lint?label=Linter&logo=github&style=flat-square)](#donya-package-manager)
[![IRC chat on freenode](https://img.shields.io/badge/IRC%20chat%20on%20freenode-%23DonyaOS-brightgreen?style=flat-square)](#donya-package-manager)
[![Drone (cloud) with branch](https://img.shields.io/drone/build/DonyaOS/PackageManager/master?logo=drone&style=flat-square)](https://cloud.drone.io/DonyaOS/PackageManager)

- [Donya or d](#donya-or-d)
- [Donya Package Manager Commands](#donya-package-manager-commands)
  - [install or `i`](#install-or-i)
  - [search or `s`](#search-or-s)
  - [remove or `r`](#remove-or-r)
  - [list or `l`](#list-or-l)
- [Contribution to Donya Package Manager](#contribution-to-donya-package-manager)
- [License](#license)

Donya Package System

![Donya Package System](https://user-images.githubusercontent.com/2658040/91432025-65355380-e876-11ea-8b4c-400b0aa77a4d.jpg)

## Donya or d

We will set `d` as an alias of `donya` later.

```
./donya install php ; install php version 7.1
./donya i php ; install php version 7.1

./donya i php7.4 ; install php version 7.4
./donya i gcc ; install gcc

./donya s php ; search all package with php prefix
./donya search php ; search all package with php prefix

./donya r php ; remove php package
./donya remove php ; remove php package

./donya r php* ; remove all php prefix package

./donya i php* ; install all php prefix package
```

## Compiling Donya Package Manager

```
$ go get github.com/DonyaOS/PackageManager
$ go env GOPATH
$ cd (paste GOPATH here)\src\github.com\DonyaOS\PackageManager
$ go build
$ ./PackageManager
```

## Donya Package Manager Commands

#### install or `i`

#### search or `s`

#### remove or `r`

#### list or `l`

`./donya list` : List of all installed packages

`./donya list all` : List of all packages available in the repository

-----------

## Contribution to Donya Package Manager

Please make sure to read the [Contributing Guide](CONTRIBUTING.md) before making a pull request. If you have a Donya-related project/feature/tool, add it with a pull request to this curated list!

Thank you to all the people who already contributed to DonyaOS!

## License

[MIT License](LICENSE)
