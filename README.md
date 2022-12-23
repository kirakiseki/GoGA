# GoGA
An implementation of genetic algorithm in Golang

# Usage
run in **Bash(Linux/macOS)/CMD(Windows)**. It would print optimized value in each generation and generate a GIF(GA.gif) in the same folder to visualized the evolve process.

# Preview
## GA.gif
![GA](https://user-images.githubusercontent.com/38367158/209323654-41954dd8-3234-49b9-8e5f-4e87dd2abd30.gif)

## Windows 
![Windows](https://user-images.githubusercontent.com/38367158/209323679-6e496936-35a1-47e5-996f-addec25c25f7.png)

## macOS
![macOS](https://user-images.githubusercontent.com/38367158/209323702-daa9115c-2447-4504-80d4-b5250c205f8d.png)

## Linux
![Linux](https://user-images.githubusercontent.com/38367158/209323720-6b61e84f-9125-4831-9f83-14cd623f2774.png)

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

![CLi Args](https://user-images.githubusercontent.com/38367158/209323754-9534749e-7881-49cf-b836-1178dfa53c0b.png)

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
