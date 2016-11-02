# go言語の学習

## インストール@mac
`brew install go`

## コンパイル
go build <ファイル名.go>

## コンパイルして実行
go run <ファイル名.go>

## package
Go言語は何らかのパッケージに属している必要があります。
`package main`
とするとこちらのファイルの`func main()`から実行してくれます

## hello world

```
package main 
import "fmt"
import "./myfunc"

func main() {
    fmt.Println("hello world")
}

```

