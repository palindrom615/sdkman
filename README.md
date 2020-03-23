# sdkman

[![Build Status](https://travis-ci.org/palindrom615/sdkman.svg?branch=master)](https://travis-ci.org/palindrom615/sdkman) [![release](https://github.com/palindrom615/sdkman/workflows/release/badge.svg)](https://github.com/palindrom615/sdkman/releases)

**cli stands for client!**

This is unofficial client of [SDKMAN!](https://sdkman.io/) which is a tool for managing parallel versions of multiple Software Development Kits like jdk, gradle, scala and etc.

## Installation

### Windows & scoop

```powershell
scoop bucket add palindrom615 https://github.com/palindrom615/scoop-bucket
scoop install sdkman
```

### macOS & homebrew

```zsh
brew tap palindrom615/homebrew-tap
brew install sdkman
```

### Arch linux & aur

with yay:

```bash
yay -S sdkman
```

without yay:

```bash
git clone https://aur.archlinux.org/sdkman
cd sdkman
makepkg -si
```

### with go

If go is already installed on your computer, one line is enough to get executable:

```sh
go install --mod=mod github.com/palindrom615/sdkman
```

## Activation

### Windows

For setting up permanent enviornment variables, execute below after install sdks.

```powershell
Invoke-Expression (sdk export windows)
```

### powershell

For using only on powershell, add line below in your `$PROFILE`

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

### Show list of SDK

```
sdk list java
```

### Set default version of SDK

```
# should be installed first!
sdk use java@8.0.242-amzn
```

### Show currently using SDKs

```
sdk current
```
