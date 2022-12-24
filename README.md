# AtCoder Practice

### Language Version

- Go 1.14.1
- Kotlin 1.3.71
- Ruby 2.7.1

## Usage

### Login

login to AtCoder using [atcoder-cli](https://github.com/Tatamo/atcoder-cli)

```sh
acc login
```

linkage with [online-judge-tools](https://github.com/kmyk/online-judge-tools)

```sh
oj login https://beta.atcoder.jp/
```

### New

show user templates in the config directory

```sh
acc templates
```

create new contest project directory

```sh
cd tasks
acc n <contest-id> --template <name>
cd <contest-id>/a
```

get contest information

```sh
acc contest
```

add new directory for the task in the project directory

```sh
cd ../
acc a --template <name>
```

### Test

test your code

#### Go

```sh
oj t -c "go run main.go"
```

```sh
go run main.go
```

#### Kotlin

```sh
kotlinc Main.kt && oj t -c "kotlin MainKt"
```

```sh
kotlinc Main.kt && kotlin MainKt
```

#### Ruby

```sh
oj t -c "ruby main.rb"
```

```sh
ruby main.rb
```

### Submit

submit the program

#### Go

```sh
acc s main.go
```

#### Kotlin

```sh
acc s Main.kt
```

#### Ruby

```sh
acc s main.rb
```

## Useful links

- [AtCoder](https://atcoder.jp/contests/)
- [AtCoder Problems](https://kenkoooo.com/atcoder/#/table/)
