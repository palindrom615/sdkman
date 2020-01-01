## sdk

![travis ci status](https://travis-ci.org/palindrom615/sdkman.svg?branch=master)

**cli stands for client!**

This is unofficial client of [SDKMAN!](https://sdkman.io/) which is a service to manage multiple versions of jdk, gradle, scala and etc.

Though SDKMAN is convenient and makes you smile on linux and macOS environment, whole client of SDKMAN is written in bash script and that means SDKMAN is very platform specific. It will not work properly on your fish shell or windows environment.

There were notable efforts to use SDKMAN [on powershell](https://github.com/flofreud/posh-gvm) or [fish](https://github.com/reitzig/sdkman-for-fish) but the only solution might be rewriting whole client in platform-agnostic way like go or rust.

There's [plan to migrate SDKMAN client](https://github.com/sdkman/sdk) with rust which is platform-agnostic language, and this repository is just alternative before that releases.

## Installation

### with go

If go is already installed on your computer, one line is enough to get executable:

```bash
# for linux & macOS
env GO111MODULE=on go install github.com/palindrom615/sdkman/cmd/sdk
```

```powershell
# for powershell
$env:GO111MODULE=on; go install github.com/palindrom615/sdkman/cmd/sdk
```

### Windows

For setting up permanent enviornment variable, open powershell and type

```powershell
Invoke-Expression (sdk export windows)
```

### powershell

For using only on powershell,

```powershell
Add-Content $Profile "Invoke-Expression (sdk export posh)"
```

### bash, fish, zsh

```bash
# You should add "eval $(sdk" export bash)" on your .bashrc file
echo "eval \$(sdk export bash)" >> ~/.bashrc
echo "eval \$(sdk export zsh)" >> ~/.zshrc
echo "eval (sdk export fish)" >> ~/.config/config.fish
```

## Usage

### Install SDK

```
sdk install java
sdk install gradle@6.0.1
```

## Build

```bash
go build cmd/sdk/main.go
```
