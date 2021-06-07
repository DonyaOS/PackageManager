# Donya Package Manager
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-4-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

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
$ cd $(go env GOPATH)/src/github.com/DonyaOS/PackageManager
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

## Install on debian chroot environment

1. Create a simple chroot in debian based distro 
[Create debian chroot](https://gist.github.com/esmaeelE/ab35177313793d342174c28ff4bcbc07)
2. activate chroot environment 

2. Download and Install go inside chroot env

[Official download link](https://golang.org/dl/)

extract downloaded package and place under `/usr/local`

`tar -C /usr/local -xzvf go1.13.linux-amd64.tar.gz`

Change name of package `go1.13.linux-amd64.tar.gz` as it changes version.

Once the file is extracted, edit the `$PATH` environment variable so that the system knows where the Go executable binaries are located. You can do this either by appending the following line to the /etc/profile file (for a system-wide installation) or to the `$HOME/.profile` file (for a current user installation):`~/.profile`

`export PATH=$PATH:/usr/local/go/bin`

Save the file, and apply the new PATH environment variable to the current shell session by typing:

`source ~/.profile`

To verify that Go has been successfully installed run the following command which will print the Go version:

`go version`

[Original Resource](https://linuxize.com/post/how-to-install-go-on-debian-10/)

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://maxbase.org/"><img src="https://avatars.githubusercontent.com/u/2658040?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Max Base</b></sub></a><br /><a href="https://github.com/DonyaOS/PackageManager/commits?author=BaseMax" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/jbampton"><img src="https://avatars.githubusercontent.com/u/418747?v=4?s=100" width="100px;" alt=""/><br /><sub><b>John Bampton</b></sub></a><br /><a href="https://github.com/DonyaOS/PackageManager/commits?author=jbampton" title="Code">💻</a></td>
    <td align="center"><a href="https://1995parham.me"><img src="https://avatars.githubusercontent.com/u/8181240?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Parham Alvani</b></sub></a><br /><a href="https://github.com/DonyaOS/PackageManager/commits?author=1995parham" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/esmaeelE"><img src="https://avatars.githubusercontent.com/u/22000310?v=4?s=100" width="100px;" alt=""/><br /><sub><b>EEC</b></sub></a><br /><a href="https://github.com/DonyaOS/PackageManager/commits?author=esmaeelE" title="Documentation">📖</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!