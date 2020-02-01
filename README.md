# switchecker
switchecker is tool of checking for missing switch case about go enum(Actually, this is not enum. It is constant value) in go file.
It can be take safety when using `switch` word in go source code.
It is ignore for dirty `case` (e.g case 5, case "strings"). 
So, switcher only support to check the missing `case`.

## Install
```bash
go get github.com/bannzai/switchecker
```

## Usage

```bash
$ switchecker --help
Usage of switchecker:
-source string
	Source go files are containing enum definition. Multiple specifications can be specified separated by ,. e.g) *.go,pkg/**/*.go  (default "*.go")
-target string
	Target go files are containing to use enum. Multiple specifications can be specified separated by ,. e.g) *.go,pkg/**/*.go  (default "*.go")
-verbose
	Enabled verbose log
```

### Example 
#### Valid pattern
Assuming go file exists. It has enum and call this enum.

```go
package main

type language int

const (
	golang language = iota
	swift
)

func main() {
	lang := golang
	switch lang {
	case golang:
		println("golang")
	case swift:
		println("swift")
	default:
		println("default")
	}
}
```

Then try running **switchecker**.
```bash
$ switchecker -source=main.go -target=main.go
Succesfull!! $ switchecker -source=main.go -target=main.go
```
Output Succesfull!! and actually execute command to console when go file is valid :tada: 

#### Invalid pattern
Next example is invalid pattern about **switchecker**.
Commentout for `case swift`.
```go
package main

type language int

const (
	golang language = iota
	swift
)

func main() {
	lang := golang
	switch lang {
	case golang:
		println("golang")
//	case swift:
//		println("swift")
	default:
		println("default")
	}
}
```

And exec **switchecker**.
```bash
$ switchecker -source=main.go -target=main.go
2020/02/02 01:18:08 missing enum pattern for main.language.swift. at /Users/bannzai/go/src/github.com/bannzai/switchecker/example/multipy_file/main.go:12:107
```

When input go file exists missing switch case, you got error information.

If you want to use **switchecker** immediately, [see more example, and play it](https://github.com/bannzai/switchecker/tree/master/example).

## LICENSE
**switchecker** is available under the MIT license. See the LICENSE file for more info.


