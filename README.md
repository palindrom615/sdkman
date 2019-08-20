## sdkman-cli

![travis ci status](https://travis-ci.org/palindrom615/sdkman-cli.svg?branch=master)

**⚠This repo lacks a lot of features now and does not guarantee compatibility with official sdkman-cli.⚠**

**cli stands for client!**

This is unofficial client of [SDKMAN!](https://sdkman.io/) which is amazing service and the only sane way to control whole java version as far as I know.

Though SDKMAN is convenient and make you smile on linux and macOS environment, whole client of SDKMAN is written in bash script. It's a blessing🙏 to be able to use powerful coreutils instead of implementing everything. But I recently migrated my environment from linux to windows, which means excommunication 😈 from the blessing.

There was [another project to use sdkman in windows env](https://github.com/flofreud/posh-gvm), but it is not well maintained. I gladly wanted to contribute the repo, but I am not good at powershell. also, I am using both wsl, git bash and powershell in windows. Why do I even need to be bound with a specific shell? I even had to install another package to use SDKMAN with fish shell!

There's [plan to migrate SDKMAN client](https://github.com/sdkman/sdk) with rust which is platform-agnostic language, and this repository is just alternative before that releases.

## build

```bash
go build main.go
```  
