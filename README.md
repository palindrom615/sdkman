## sdkman-cli

![travis ci status](https://travis-ci.org/palindrom615/sdkman-cli.svg?branch=master)

**cli stands for client!**

This is unofficial client of [SDKMAN!](https://sdkman.io/) which is amazing service and the only sane way to control whole java version as far as I know.

Though SDKMAN is convenient and makes you smile on linux and macOS environment, whole client of SDKMAN is written in bash script. It means SDKMAN is very platform specific. It will not work properly on fish shell or windows environment.

There were notable efforts to use SDKMAN [on powershell](https://github.com/flofreud/posh-gvm) or [fish](https://github.com/reitzig/sdkman-for-fish) but the only solution for these might be writing whole client in platform-agnostic way like go or rust. 

There's [plan to migrate SDKMAN client](https://github.com/sdkman/sdk) with rust which is platform-agnostic language, and this repository is just alternative before that releases.

## Installation

```bash
go get github.com/palindrom615/sdkman-cli
```

## Build

```bash
go build main.go
```  
