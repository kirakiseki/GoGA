# GoGA
An implementation of genetic algorithm in Golang

# Usage
run in **Bash(Linux/macOS)/CMD(Windows)**. It would print optimized value in each generation and generate a GIF(GA.gif) in the same folder to visualized the evolve process.

GA.gif:

![](https://s-gz-4165-image.oss.dogecdn.com/2022/12/15/GA1.gif)

# Release

You can get pre-build binary release at [GoGA/releases](https://github.com/kirakiseki/GoGA/releases), which contains various packages for different targets.

# Command-line Arguments
|Param|Type|Default Value|Description|
|--|--|--|--|
|**min**|Float64|-5.0|minimum x value|
|**max**|Float64|5.0|maximum x value|
|**g**|Int|30|maximum generation|
|**s**|Int|100|population size|
|**m**|Float64|0.1|MutationProbability|
|**c**|Float64|0.6|CrossoverProbability|
|**p**|Float64|0.01|Precision|
|**f**|String|r|function to select individuals (r for Roulette, t for Tournament)|

You can get this usage sheet in command line using '-h'

![](https://s-gz-4165-image.oss.dogecdn.com/2022/12/15/20221215095414.png)
# Dev

## Environment
Go 1.19.3

## Build

```bash
go build main.go
```

## Run
```bash
./main <CLI Args>
```