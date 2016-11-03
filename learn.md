# Go言語

## Go言語の特徴

### クロスコンパイルのサポート
Goは，コンパイル時にOSとCPUアーキテクチャを指定し，その環境に合わせた実行ファイルを生成できます。

### 並行処理のサポート
Goでは，ゴルーチンという軽量スレッドを用いて処理を並行に実施し，同時に実行されているゴルーチンの間ではチャネルという機能でデータをやりとりするしくみが備わっています。

## Installation
for macOS  
`brew install go`

for Linux  
`sudo yum install golnag`

## コマンドラインツール
| コマンド     | 用途                                                   |
|--------------|--------------------------------------------------------|
| go build     | プログラムのビルド                                     |
| go fmt       | Goの規約に合わせてプログラムを整形                     |
| go get       | 外部パッケージの取得                                   |
| go install   | プログラムのビルドとインストール                       |
| go run       | プログラムのビルドと実行                               |
| go test      | テストやベンチマークの実行                             |
| go tool yacc | パーサをGoで出力するGo実装のyacc（パーサジェネレータ） |
| godoc        | ソースからドキュメントの生成                           |

## what is package
Go言語は何らかのパッケージに属している必要があります。  
`package main`  
とするとこちらのファイルの`func main()`から実行してくれます。

## Getting Started
```go
package main 
import(
  "fmt"
)

func main() {
    fmt.Println("hello world")
}
```

## Golang言語仕様

### package
Goのコードはパッケージの宣言から始まります。  
`package main`

プログラムをコンパイルして実行すると，まずmainパッケージの中にあるmain()関数が実行されるため，ここに処理を記述しています。 

```go
func main(){}
```

### import

importは，プログラム内にほかのパッケージを取り込むために使用する。

```go
import (
    "fmt"
)
```

インポートしたパッケージ内へは，パッケージ名にドットを付けてアクセスできる。

### 複数のimport

複数のパッケージを取り込む場合は次のように縦に並べて記述する。

```go
import (
    "fmt"
    "strings"
)
```

Goの標準パッケージはインストールしたGoの中に含まれているため自動でパスが解決され，それ以外のパッケージはGOPATH環境変数からパスを解決する。

### オプション
インポートするパッケージ名の前には，いくつかのオプションが指定できる。

```go
import (
    f "fmt"
    _ "github.com/konojunya/golang_hoge"
    . "strings"
)

func main() {
    // fmt -> fになる
    // string.ToUpper() -> stringを省略できる
    f.Println(ToUpper("hello world"))
}
```

_を付けた場合は，インポートしたパッケージを使用しないことをコンパイラに伝えます。

### 変数

型を指定して、変数に代入する方法と型を自動で推測してくれるものとあります。  
宣言には`var`を使います。
```go
var message string = "hello world"
message := "hello world" // 同じ
```

### 定数

`var`に対して`const`で定義します。

```go
const PI = 3.14.15926535
```

### 初期化

Goでは変数を定義して何も入れない場合に数値なら`0`や文字列なら`''`、真偽値なら`false`などが入ります。

### if

Goでは，if文の条件部に丸括弧は必要ありません。

```go
func main() {
    a, b := 10, 100
    if a > b {
        fmt.Println("a is larger than b")
    } else if a < b {
        fmt.Println("a is smaller than b")
    } else {
        fmt.Println("a equals b")
    }
}
```

なお，if文の処理が1行の場合に波括弧を省略可能な言語もありますが，Goではそうした省略はコンパイルエラーになります。

```go
if n == 10
    fmt.Println(n)
    // syntax error: missing { after if clause
```

また，三項演算子は存在しない。

### for

Goにはには、ループ文は`for`しかありません。　　
Goのfor文では，CやJavaでは必要な条件部の丸括弧が必要ありません。

```go
func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
```

Goで繰り返しを表現する方法はfor文以外になく，ほかの言語におけるwhile文やdo/while文，loop文といったような表現は，すべてfor文を用いて行います。

```go
// C
int n = 0;
while (n < 10) {
    printf("n = %d\n", n);
    n++;
}

// Go
n := 0
for n < 10 {
    fmt.Printf("n = %d\n", n)
    n++
}
```

#### 無限ループ

Cの場合，for文を使って次のように無限ループを表現します。

```c
for (;;) {
    doSomething();
}
```

Goでは，for文の条件部を省略することで同様の表現ができます。

```go
for {
    doSomething()
}
```

#### break, continue

繰り返し制御にはCやJavaと同様に，ループを終了するbreak，ループの最初から処理を再開するcontinueを使用できます。

```go
func main() {
    n := 0
    for {
        n++
        if n > 10 {
            break // ループを抜ける
        }
        if n%2 == 0 {
            continue // 偶数なら次の繰り返しに移る
        }
        fmt.Println(n) // 奇数のみ表示
    }
}
```

### switch

if/else文が繰り返す場合は，switch文を用いたほうがスッキリ書ける場合があります。Goのswitch文は非常に柔軟であり，値の比較だけでなく条件分岐にも使用できます。

```go
func main() {
    n := 10
    switch n {
    case 15:
        fmt.Println("FizzBuzz")
    case 5, 10:
        fmt.Println("Buzz")
    case 3, 6, 9:
        fmt.Println("Fizz")
    default:
        fmt.Println(n)
    }
}
```

switch 文に指定した値に一致するcaseが実行され，どのcaseにも一致しなかった場合はdefaultが実行されます。caseには単一の値だけでなく，カンマで区切った複数の値も指定できます。

#### fallthrough

CやJavaなどのswitch文は，1つのcaseが実行されるとその次のcaseに処理が移るため，単一のcaseの実行で終わらせたい場合に，caseごとにbreakを書く必要がありました。しかしGoのswitch文では，caseが1つ実行されると次のcaseに実行が移ることなくswitch文が終了するため，breakをいちいち書く必要はありません。

ただ，caseの処理が終わったあとに，次のcaseに処理を進めたい場合もあります。そうした場合はcase内にfallthroughを書くことで，明示的に次のcaseに処理を進めることができます。

```go
func main() {
    n := 3
    switch n {
    case 3:
        n = n - 1
        fallthrough
    case 2:
        n = n - 1
        fallthrough
    case 1:
        n = n - 1
        fmt.Println(n) // 0
    }
}
```

Goではcaseに式を持ってきても評価されます。

```go
func main() {
    n := 10
    switch {
    case n%15 == 0:
        fmt.Println("FizzBuzz")
    case n%5 == 0:
        fmt.Println("Buzz")
    case n%3 == 0:
        fmt.Println("Fizz")
    default:
        fmt.Println(n)
    }
}
```

### 関数

```go
func hello() {
    fmt.Println("hello")
}

func main() {
    hello() // hello
}
```

引数がある場合は変数と型を指定します。複数の同じ型が続く場合は，型の宣言は最後の1つにまとめることができます。

```go
func sum(i, j int) { // func sum(i int, j int) と同じ
    fmt.Println(i + j)
}

func main() {
    sum(1, 2) // 3
}
```

戻り値がある場合は引数の次に指定します。

```go
func sum(i, j int) int {
    return i + j
}

func main() {
    n := sum(1, 2)
    fmt.Println(n) // 3
}
```

Goは関数で複数の値を同時に返すことができます。

```go
func swap(i, j int) (int, int) {
    return j, i
}

func main() {
    x, y := 3, 4
    x, y = swap(x, y)
    fmt.Println(x, y) // 4, 3

    x = swap(x, y) // コンパイルエラー

    x, _ = swap(x, y) // 第二戻り値を無視
    fmt.Println(x) // 3

    swap(x, y) // コンパイル，実行ともに可能
}
```

関数の実行時には，戻り値を格納する変数を必要な数だけ用意する必要があります。関数が返す値の数と，受け取る変数の数が合わないとコンパイルエラーになります。ただし，無視したい戻り値がある場合は_で明示的に無視することで，戻り値を受け取らなくてもコンパイル，実行ともに可能です。

#### エラーを返す関数

Goでは，関数が多値を返せることを利用して，内部で発生したエラーを戻り値で表現します。関数の処理に成功した場合はエラーはnilにし，異常があった場合はエラーだけに値が入り，他方はゼロ値になります。

たとえばファイルを開くos.Open()は，1つ目の戻り値に*os.File，2つ目にerrorを返します。

```go
func main() {
    file, err := os.Open("hello.go")
    if err != nil {
        // エラー処理
        // returnなどで処理を別の場所に抜ける
    }
    // fileを用いた処理
}
```

自作のエラーは，errorsパッケージを用いて作ることができます。

```go
package main

import (
    "errors"
    "fmt"
    "log"
)

func div(i, j int) (int, error) {
    if j == 0 {
        // 自作のエラーを返す
        return 0, errors.New("divied by zero")
    }
    return i / j, nil
}

func main() {
    n, err := div(10, 0)
    if err != nil {
        // エラーを出力しプログラムを終了する
        log.Fatal(err)
    }
    fmt.Println(n)
}
```

複数の値を返す場合もエラーを最後にする慣習があるため，自分でAPIを設計する場合もエラーを最後にするほうがよいでしょう。

異常を戻り値で表現できない場合については，パニックとリカバで解説します。

#### 名前付き戻り値

Goでは，戻り値にあらかじめ名前を付けることができます。先ほどの関数の戻り値に，次のように名前を付けてみます。

```go
func div(i, j int) (result int, err error)
```

名前付き戻り値は，関数内ではゼロ値で初期化された変数として扱うことができます。また，変数に名前を付けている場合は，returnのあとに返す値を明示する必要がなく，returnされた時点での名前付き戻り値の値が自動的に返されることになります。

```go
func div(i, j int) (result int, err error) {
    if j == 0 {
        err = errors.New("divied by zero")
        return // return 0, errと同じ
    }
    result = i / j
    return // return result, nilと同じ
}
```

名前付き戻り値を用いることで，関数の宣言から戻り値の意味が読み取りやすくなると同時に，戻り値のための変数の初期化が不要になり，同じ型の戻り値が多かった場合のreturnの書き間違えなどを防ぐこともできます。

ただし，戻り値に名前を付けても，returnのあとに戻す値を明示することは可能です。プログラムのわかりやすさを重視して使い分けるとよいでしょう。

#### 関数リテラル

関数リテラルを用いると，無名関数を作ることができます。即時に実行する関数は次のように記述できます。

```go
func main() {
    func(i, j int) {
        fmt.Println(i + j)
    }(2, 4)
}
```

Goにおける関数は第一級オブジェクトであるため，関数を変数に代入したり関数の引数に渡すことができます。たとえば，上記の関数を代入する変数の型は次のようになります。

```go
var sum func(i, j int) = func(i, j int) {
    fmt.Println(i + j)
}

func main() {
    sum(2, 4)
}
```

### 配列

Goの配列は固定長です。可変長配列は後述するスライスがそれにあたります。

```go
var arr1 [4]string
```

宣言と同時に初期化することも可能で，その場合は[...]を用いることで，必要な配列の長さを暗黙的に指定できます。

```go
// どちらも同じ型
arr := [4]string{"a", "b", "c", "d"}
arr := [...]string{"a", "b", "c", "d"}
```

配列の型は長さも情報として含むため，次のarr1とarr2は，要素の型は同じstringですが，長さが違うため配列としては別の型です。関数fnは[4]string型を引数にとるため，型の合わないarr2を渡すとコンパイルエラーになります。

```go
func fn(arr [4]string) {
    fmt.Println(arr)
}

func main() {
    var arr1 [4]string
    var arr2 [5]string

    fn(arr1) // ok
    fn(arr2) // コンパイルエラー
}
```

また，関数に配列を渡す場合は値渡しとなり，配列のコピーが渡されます。次のfn()の中で配列に対して行った変更は，main()側には反映されません。

### スライス

スライスは，可変長配列として扱うことができます。配列を直接使うのは，シビアなメモリ管理が必要な一部のプログラムだけなので，同じ性質のデータを束ねて扱うという用途であれば，基本的にはスライスを用います。

宣言

```go
var s []string
```

このように，スライスの型には配列のように長さの情報はありません。

初期化を同時に行う場合は，配列と同じように書くことができます。またスライスも，配列同様に添字でアクセスできます。

```go
s := []string{"a", "b", "c", "d"}
fmt.Println(s[0]) // "a"
```

#### append

スライスの末尾に値を追加する場合はappend()を使用します。append()は，スライスの末尾に値を追加し，その結果を返す組込み関数です。複数の値を追加することもできます。

```go
var s []string
s = append(s, "a") // 追加した結果を返す
s = append(s, "b")
s = append(s, "c", "d")
fmt.Println(s) // [a b c d]
```

次のように指定すれば，スライスに別のスライスの中身を展開して追加することもできます。

```go
s1 := []string{"a", "b"}
s2 := []string{"c", "d"}
s1 = append(s1, s2...) // s1にs2を追加
fmt.Println(s1)        // [a b c d]
```

#### range

配列やスライスに格納された値を，先頭から順番に処理するような場合は，添字によるアクセスの代わりにrangeを使用できます。

for文の中でrangeを用いると，添字と値の両方が取得できます。

```go
var arr [4]string

arr[0] = "a"
arr[1] = "b"
arr[2] = "c"
arr[3] = "d"

for i, s := range arr {
    // i = 添字, s = 値
    fmt.Println(i, s)
}
```

rangeは配列やスライスのほかに，string，マップ，チャネルに対しても使用できます。マップについては本章で，チャネルについては5章で解説します。

#### 値の切り出し

string，配列，スライスから，値を部分的に切り出すことができます。次のように始点と終点をコロンで挟んで指定すると，その範囲の値を切り出すことができます。始点，終点を省略した場合，それぞれ先頭，末尾になります。

```go
s := []int{0, 1, 2, 3, 4, 5}
fmt.Println(s[2:4])      // [2 3]
fmt.Println(s[0:len(s)]) // [0 1 2 3 4 5]
fmt.Println(s[:3])       // [0 1 2]
fmt.Println(s[3:])       // [3 4 5]
fmt.Println(s[:])        // [0 1 2 3 4 5]
```

#### 可変長引数

関数において引数を次のように指定すると，可変長引数として，任意の数の引数をその型のスライスとして受け取ることができます。

```go
func sum(nums ...int) (result int) {
    // numsは[]int型
    for _, n := range nums {
        result += n
    }
    return
}

func main() {
    fmt.Println(sum(1, 2, 3, 4))  // 10
}
```

### マップ

マップは，値をKey-Valueの対応で保存するデータ構造です。

#### 宣言と初期化

たとえばintのキーにstringの値を格納するマップは次のように宣言します。

```go
var month map[int]string = map[int]string{}
```

次のようにキーを指定して値を保存します。

```go
month[1] = "January"
month[2] = "February"
fmt.Println(month) // map[1:January 2:February]
```

宣言と初期化を一緒に行う場合は次のように書きます。

```go
month := map[int]string{
    1: "January",
    2: "February",
}
fmt.Println(month) // map[1:January 2:February]
```

#### マップの操作

マップから値を取り出す場合は，次のようにキーを指定し，戻り値として受け取ります。

```go
jan := month[1]
fmt.Println(jan) // January
```

このとき2つ目の戻り値も受け取ると，指定したキーがこのマップに格納されているかをboolで返します。マップ内のキーの存在を調べるような場合には，値を無視して次のようにします。

```go
_, ok := month[1]
if ok {
    // データがあった場合
}
```

マップからデータを消す場合は組込み関数のdelete()を使用します。

```go
delete(month, 1)
fmt.Println(month) // map[1:January]
```

スライス同様，rangeを用いるとfor文でKey-Valueをそれぞれ受け取りながら処理を進めることができます。ただし，マップの場合は取り出される順番は保証されない点に注意してください。

```go
for key, value := range month {
    fmt.Printf("%d %s\n", key, value)
}
```

### ポインタ

Goはポインタを扱うことができます。ポインタ型の変数は，型の前に`*`を付けます。アドレスは変数の前に`&`を付けて取得できるため，Cと似たような形で表現できます。

```go
func callByValue(i int) {
    i = 20 // 値を上書きする
}

func callByRef(i *int) {
    *i = 20 // 参照先を上書きする
}

func main() {
    var i int = 10
    callByValue(i) // 値を渡す
    fmt.Println(i) // 10
    callByRef(&i) // アドレスを渡す
    fmt.Println(i) // 20
}
```

しかし，Cなどと違い，Goはポインタ演算を認めていません。ポインタをデータサイズ分ずつずらして，メモリ上からデータを読み込むといったことは基本的にはできません。

### defer

ファイル操作などを行う場合，使用後のファイルは必ず閉じる必要があります。次の例では関数の最後にファイルのクローズ処理を記述していますが，その前に関数を抜ける処理があったり，後述するパニックが起こってしまうと，Close()まで到達しない場合が発生してしまいます。

```go
func main() {
    file, err := os.Open("./error.go")
    if err != nil {
        // エラー処理
    }
    // 正常処理
    file.Close()
}
```

こうした処理はdeferを用いて記述できます。先の例ではfile.Close()の関数呼び出しをdeferの後ろに記述すると，この処理がmain()を抜ける直前に必ず実行されるようになります。

```go
func main() {
    file, err := os.Open("./error.go")
    if err != nil {
        // エラー処理
    }
    // 関数を抜ける前に必ず実行される
    defer file.Close()
    // 正常処理
}
```

ファイルのClose()などは，deferを用いて記述するほうが安全です。

### パニック

エラーは戻り値によって表現するのが基本ですが，そうではない場合もあります。たとえば配列やスライスの範囲外にアクセスした場合や，ゼロ除算をしてしまった場合などです。こうした処理はエラーを返すことができないため，代わりにパニックという方法でエラーが発生します。

このパニックで発生したエラーはrecover()という組込み関数で取得し，そこでエラー処理を実施できます。recover()をdeferの中に書くことで，パニックで発生したエラーの処理を実施してから，関数を抜けることができます。

```go
func main() {
    defer func() {
        err := recover()
        if err != nil {
            // runtime error: index out of range
            log.Fatal(err)
        }
    }()

    a := []int{1, 2, 3}
    fmt.Println(a[10]) // パニックが発生
}
```

#### panic

パニックは組込み関数panic()を用いて自分で発生させることもできます。先ほどの例を自分でパニックにする場合は次のように書けます。

```go
a := []int{1, 2, 3}
for i := 0; i < 10; i++ {
    if i >= len(a) {
        panic(errors.New("index out of range"))
    }
    fmt.Println(a[i])
}
```

ただしパニックを用いるのは，エラーを戻り値として表現できない場合や，回復が不可能なシステムエラー，やむを得ず大域脱出が必要な場合などであり，基本的にエラーは関数の戻り値として呼び出し側に返すようにしましょう。

### type

次のような，IDと優先度を取得してタスクを処理する関数を考えてみます。2つの情報は両方とも数値で表すため，intとして宣言しています。

```go
func ProcessTask(id, priority int) {
}
```

この関数を呼び出す側は次のようになります。

```go
id := 3 // int
priority := 5 // int
ProcessTask(id, priority)
```

正しく呼び出せていますが，もし次のように順番を間違えたらどうなるでしょうか。

```go
id := 3
priority := 5
ProcessTask(priority, id) // コンパイルは通る
```

引数の型が合っているため，コンパイルは通ってしまいます。こうしたミスはテストによって発見することもできますが，それぞれが単なるintではなく，意味に応じた型を持っていれば，コンパイル時に間違いを知ることができます。

このような場合，Goではtypeを用いて既存の型を拡張した独自の型を定義できます。

```go
type ID int
type Priority int

func ProcessTask(id ID, priority Priority) {
}
```

typeのあとには，型の名前，その型の定義が続きます。上記では，intを拡張してIDと優先度それぞれに型を定義し，この型を用いて関数の定義を書き換えています。

呼び出す際には，型が適合していないとコンパイルエラーになります。

```go
var id ID = 3
var priority Priority = 5
ProcessTask(priority, id) // コンパイルエラー
```

このように適切な型を用意することで，型レベルの整合性をコンパイル時にチェックでき，堅牢なプログラムを記述できます。IDE（Integrated Development Environment，統合開発環境）のサポートも得やすくなり，リファクタリング時のリグレッションなども防ぎやすくなります。

### 構造体（struct）

Goには，構造体というデータ構造があります。構造体は複数のデータを1つにまとめることが基本的な役割ですが，後述するメソッドを持つことができ，RubyやJavaでのクラスに近い役割も担います。

#### 宣言

ここでは，id，detail（タスクの詳細），done（完了フラグ）の3つのフィールドを持つ，Taskという型を定義してみます。

```go
type Task struct {
    ID int
    Detail string
    done bool
}
```

構造体型もtypeを用いて宣言し，構造体名のあとにそのフィールドを記述します。各フィールドの可視性は名前で決まり，大文字で始まる場合は`パブリック`，小文字の場合はパッケージ内に閉じたスコープ`(プライベート)`となります。

この型から値を生成するには，次のように各フィールドに値を割り当てます。

```go
func main() {
    var task Task = Task{
        ID: 1,
        Detail: "buy the milk",
        done: true,
    }
    fmt.Println(task.ID) // 1
    fmt.Println(task.Detail) // "buy the milk"
    fmt.Println(task.done) // true
}
```

変数taskには，生成された構造体が格納され，各フィールドにはドットでアクセスできます。

構造体に定義した順でパラメータを渡すことで，フィールド名を省略することもできます。

```go
func main() {
    var task Task = Task{
        1, "buy the milk", true,
    }
    fmt.Println(task.ID) // 1
    fmt.Println(task.Detail) // "buy the milk"
    fmt.Println(task.done) // true
} 
```

構造体の生成時に値を明示的に指定しなかった場合は，ゼロ値で初期化されます。

```go
func main() {
    task := Task{}
    fmt.Println(task.ID) // 0
    fmt.Println(task.Detail) // ""
    fmt.Println(task.done) // false
}
```

#### ポインタ型

構造体型もアドレスを取得し，ポインタ型で扱うことができます。構造体から値を生成するときに，構造体の名前の前に&を付けると，変数には構造体の値ではなくアドレスが格納されます。Taskのポインタ型は*Taskという型になります。

```go
var task Task = Task{} // Task型
var task *Task = &Task{} // Taskのポインタ型
```

たとえば，関数に対して構造体を値渡しするとデータはコピーされるため，関数内での構造体への変更は呼び出し側には反映されません。

```go
type Task struct {
    ID int
    Detail string
    done bool
}

func Finish(task Task) {
    task.done = true
}

func main() {
    task := Task{done: false}
    Finish(task)
    fmt.Println(task.done) // falseのまま
}
```

この関数の引数をポインタ型にするには，引数の型を*Taskとします。ポインタを渡すことで，渡した側に関数内での変更が反映されます。

```go
func Finish(task *Task) {
    task.done = true
}

func main() {
    task := &Task{done: false}
    Finish(task)
    fmt.Println(task.done) // true
}
```

このように，Goでは値とポインタを用途に応じて使い分けることができます。

#### new

構造体は，組込みの関数new()を用いて初期化することもできます。new()は，構造体のフィールドをすべてゼロ値で初期化し，そのポインタを返します。

```go
type Task struct {
    ID int
    Detail string
    done bool
}

func main() {
    var task *Task = new(Task)
    task.ID = 1
    task.Detail = "buy the milk"
    fmt.Println(task.done) // false
}
```

#### コンストラクタ

Goには構造体のコンストラクタにあたる構文がありません。代わりにNewで始まる関数を定義し，その内部で構造体を生成するのが通例です。たとえばTaskをNewする関数はNewTask()という関数にし，内部でTaskを生成し，そのポインタを返します。

```go
func NewTask(id int, detail string) *Task {
    task := &Task{
        ID: id,
        Detail: detail,
        done: false,
    }
    return task
}

func main() {
    task := NewTask(1, "buy the milk")
    // &{ID:1 Detail:buy the milk done:false}
    fmt.Printf("%+v", task)
}
```

#### メソッド

型にはメソッドを定義できます。メソッドは，そのメソッドを実行した対象の型をレシーバとして受け取り，メソッドの内部で使用できます。たとえば，Taskの文字列表現を返すString()というメソッドをTaskに定義してみます。

```go
func (task Task) String() string {
    str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
    return str
}

func main() {
    task := NewTask(1, "buy the milk")
    fmt.Printf(“%s”, task) // 1) buy the milk
}
```

このメソッドは，レシーバのフィールド値から文字列を生成し，それを戻り値として呼び出し側に返します。このとき，Taskのコピーがレシーバとして渡されるため，もしメソッド内部でレシーバの中身を変更していても，呼び出し側には反映されません。

呼び出し側に変更を反映したい場合は，レシーバをポインタとして受け取るようにします。たとえば，タスクを終了済みにするFinish()というメソッドをTaskに定義してみます。

```go
func (task *Task) Finish() {
    task.done = true
}

func main() {
    task := NewTask(1, "buy the milk")
    task.Finish()
    // &{ID:1 Detail:buy the milk done:true}
    fmt.Printf("%+v", task)
}
```

#### インタフェース

##### 宣言

たとえば，先ほどTaskに実装したString()という振る舞いが規定されていることを表すインタフェースは，次のように定義します。

```go
type Stringer interface {
    String() string
}
```

インタフェースの名前は，実装すべき関数名が単純な場合は，その関数名にerを加えた名前を付ける慣習があります。よってString()を実装するインタフェースはStringerとなります。

##### 実装

Goでは，Javaのimplements構文のように，インタフェースを実装していることを明示的に宣言する構文はありません。Goは，型がインタフェースに定義されたメソッドを実装していれば，インタフェースを満たしているとみなします。たとえば先ほどTaskにはString()メソッドを実装しているため，Stringerを引数に取る次のような関数に渡すことができます。

```go
func Print(stringer Stringer) {
    fmt.Println(stringer.String())
}

Print(task)
```

実は，このStringerインタフェースはGoのfmtパッケージに標準で定義されており，レシーバの文字列表現を取得するためのAPIになっています。

#### interface{}

```go
type Any interface {
}
```

このインタフェースは，実装すべきメソッドを指定していません。つまり，すべての型はこのインタフェースを実装していることになります。

これを利用すると，次のようにどんな型も受け取ることができる関数を定義できます。

```go
func Do(e Any) {
  // do something
}

Do("a") // どのような型も渡すことができる
```

また，Anyのような型を定義しなくても，次のように直接記述することもできます。

```go
func Do(e interface{}) {
  // do something
}

Do("a") // どのような型も渡すことができる
```

たとえば，fmt.Println()などのいわゆるプリント関数は，これを用いて次のように定義されているため，型を気にせずに複数の値を渡すことができるのです。

```go
// fmt.Printlnの定義
// 任意の型を可変長引数で受け取る
func Println(a ...interface{}) (n int, err error)
```

#### 型の埋め込み

Goでは，継承はサポートされていません。代わりにほかの型を「埋め込む」（Embed）という方式で，構造体やインタフェースの振る舞いを拡張できます。

##### 構造体の埋め込み

例として，先ほどのTaskに対して，Userの情報を埋め込んでみましょう。

User構造体の定義は以下とし，メソッドとしてFullName()とコンストラクタ関数を実装します。

```go
type User struct {
    FirstName string
    LastName string
}

func (u *User) FullName() string {
    fullname := fmt.Sprintf("%s %s",
        u.FirstName, u.LastName)
    return fullname
}

func NewUser(firstName, lastName string) *User {
    return &User{
        FirstName: firstName,
        LastName: lastName,
    }
}
```

これをTaskに埋め込みます。Taskの構造体型宣言時に，フィールドではなく型のみを記述することで，その型を埋め込むことができます。

```go
type Task struct {
    ID int
    Detail string
    done bool
    *User // Userを埋め込む
}

func NewTask(id int, detail,
    firstName, lastName string) *Task {
    task := &Task{
        ID: id,
        Detail: detail,
        done: false,
        User: NewUser(firstName, lastName),
    }
    return task
}
```

埋め込まれたUserのフィールドやメソッドは，Taskが実装しているかのように振る舞います。また，埋め込まれた型の実体にもアクセスできます。

```go
func main() {
    task := NewTask(1, "buy the milk", "Jxck", "Daniel")
    // TaskにUserのフィールドが埋め込まれている
    fmt.Println(task.FirstName)
    fmt.Println(task.LastName)
    // TaskにUserのメソッドが埋め込まれている
    fmt.Println(task.FullName())
    // Taskから埋め込まれたUser自体にもアクセス可能
    fmt.Println(task.User)
}
```

##### インタフェースの埋め込み

インタフェースも埋め込み可能です。主な用途は，複数のインタフェースを埋め込んで新たなインタフェースを定義することです。

たとえばioパッケージでは，Reader，Writerといったインタフェースが定義されています。

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

このとき，Read()とWrite()を両方定義したインタフェースであるReadWriterは，2つのインタフェースを埋め込んで次のように定義できます。

```go
type ReadWriter interface {
    Reader
    Writer
}
```

ioパッケージには同様に，Closerインタフェースを埋め込んだReadCloser，WriteCloser，ReadWrite Closerなども定義されています。このように既存のインタフェースをさらに拡張したインタフェースを定義するのにも，埋め込みの概念が使われます。

### 型の変換

Goでは，暗黙的な型変換が起こることはありません。しかし型を変換できないわけではなく，明示的に変換する方法がいくつか提供されています。

#### キャスト

キャストは，キャストしたい型を指定して次のように行います。

```go
var i uint8 = 3
var j uint32 = uint32(i) // uint8 -> uint32
fmt.Println(j)           // 3

var s string = "abc"
var b []byte = []byte(s) // string -> []byte
fmt.Println(b)           // [97 98 99]

// cannot convert "a" (type string) to type int
a := int("a")
```

キャストに失敗した場合はパニックが発生します。

### 並列処理

複数の処理を効率良く行うために，Goは言語自体が並行処理に必要な機能をサポートしています。特に本章で扱うゴルーチンやチャネルの機能などは，Goで並行処理プログラミングをするうえで必要不可欠な知識であり，これらを適切に使うことで，マルチコアが一般的になった近年のマシンリソースを最大限に引き出す，パフォーマンスの良いプログラムを作成できるようになります。

本節では，ゴルーチンやチャネルを用いた並行処理の考え方と，それらと合わせてよく使うsyncパッケージの使い方などについて解説します。

#### ゴルーチン

Goには，ゴルーチン（Goroutine）という軽量スレッドのしくみがあります。ここまで行っていたmain()関数も，1つのゴルーチンの中で実行されています。go構文を用いて，任意の関数を別のゴルーチンとして起動することで，処理を並行して走らせることができます。

ここではHTTPへのアクセス処理を用いて，ゴルーチンの使い方とチャネルによるメッセージのやりとりの方法を見てみましょう。

たとえば，3つのWebサイトにアクセスし，そのステータスコードを表示するプログラムを考えます。

***ゴルーチンを使わない場合***

```go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    urls := []string{
        "http://example.com",
        "http://example.net",
        "http://example.org",
    }
    for _, url := range urls {
        res, err := http.Get(url)
        if err != nil {
            log.Fatal(err)
        }
        defer res.Body.Close()
        fmt.Println(url, res.Status)
    }
}
```

http.Get()は同期処理であるため，この方法では別々のサイトにリクエストしているにもかかわらず，前のレスポンスが返らないと次のリクエストに進むことができません。

しかし，それぞれのリクエストは互いに依存がなく独立しているため，順番に実行する必要はありません。こうした場合，ゴルーチンを用いることで3つのリクエストを並行して行うことができます。

***ゴルーチンを使った場合***

先ほどのプログラムを，リクエストが別々のゴルーチンで行われるように修正してみます。実際にリクエストを発行している部分を無名関数化し，関数の前に`go`というキーワードを加えると，それだけで関数の実行が別のゴルーチンで行われるようになります。

```go
func main() {
    urls := []string{
        "http://example.com",
        "http://example.net",
        "http://example.org",
    }
    for _, url := range urls {
        // 取得処理をゴルーチンで実行する
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            fmt.Println(url, res.Status)
        }(url)
    }
    // main()が終わらないように待ち合わせる
    time.Sleep(time.Second)
}
```

ここではmain()が実行されたときに内部で3つのゴルーチンを起動していますが，起動が終わってゴルーチン内で処理が行われている間もmain()は先に進んでしまうため，待ち合わせのためにtime.Sleep()を呼んで1秒間main()を止めています。

##### sync.WaitGroup

先ほどの例ではtime.Sleep()でmain()を1秒間待たせていましたが，実際に待ちたいのはhttp.Get()を行っているすべてのゴルーチンの終了です。

起動したすべてのゴルーチンの終了を待ち合わせるにはsync.WaitGroupが利用できます。sync.WaitGroupは，Add()でカウントを増やしDone()でカウントを減らし，Wait()でカウントがゼロになるまで待ち合わせます。

```go
func main() {
    wait := new(sync.WaitGroup)
    urls := []string{
        "http://example.com",
        "http://example.net",
        "http://example.org",
    }
    for _, url := range urls {
        // waitGroupに追加
        wait.Add(1)
        // 取得処理をゴルーチンで実行する
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            fmt.Println(url, res.Status)
            // waitGroupから削除
            wait.Done()
        }(url)
    }
    // 待ち合わせ
    wait.Wait()
}
```

#### チャネル

複数のゴルーチン間でデータをやりとりしたい場合，組込みのチャネル（channel）という機能を用いることで，メッセージパッシング（情報をメッセージとして送受信する）によってデータを送受信できます。チャネルはmake()関数に型を指定して生成することで，その型のデータの書き込みと読み出しができます。

```go
// stringを扱うチャネルを生成
ch := make(chan string)

// チャネルにstringを書き込む
ch <- "a"

// チャネルからstringを読み出す
message := <- ch
```

今回の場合は，ゴルーチン内で取得したステータスコードをチャネルに書き込み，それをmain()のゴルーチンで読み出すことで，ゴルーチン間でデータを受け渡すことができます。

```go
func main() {
    urls := []string{
        "http://example.com",
        "http://example.net",
        "http://example.org",
    }
    statusChan := make(chan string)
    for _, url := range urls {
        // 取得処理をゴルーチンで実行する
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            statusChan <- res.Status
        }(url)
    }
    for i := 0; i < len(urls); i++ {
        fmt.Println(<-statusChan)
    }
}
```

ゴルーチンの中でstatusChanに値が書き込まれるまで，main()の中では値を読み出すことができません。この場合，main()内ではstatusChanの読み出しが3回完了するまで処理がブロックされるため，waitGroupのような待ち合わせ処理は必要ありません。

##### チャネルを返すパターン

先ほどはmain() 内の匿名関数でHTTPのGETを実行していましたが，この処理をgetStatus()という別の関数にし，関数が内部で生成したチャネルを返すように実装してみます。

```go
func getStatus(urls []string) <-chan string {
    // 関数でチャネルを生成
    statusChan := make(chan string)
    for _, url := range urls {
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            statusChan <- res.Status
        }(url)
    }
    return statusChan // チャネルを返す
}

func main() {
    urls := []string{
        "http://example.com",
        "http://example.net",
        "http://example.org",
    }
    statusChan := getStatus(urls)

    for i := 0; i < len(urls); i++ {
        fmt.Println(<-statusChan)
    }
}
```

まず，getStatus()内で結果を渡すためのstatus Chanを生成します。次に非同期に行う処理を匿名関数にし，リクエストをそれぞれ別のゴルーチンで実行します。関数自体はstatusChanを返して終了し，起動されたゴルーチンが内部でstatusChanに結果を書き込んでいきます。

main()は，関数を呼び出すと同時に結果を受信するチャネルを受け取り，それをfor文内で読み出します。これにより，main()側が非常にスッキリと記述でき，ロジックの大半はgetStatus()に隠蔽（いんぺい）できました。

また，このときgetStatus()はmain()がチャネルに値を書き込むことを想定していません。こうした場合は，getStatus()の戻り値を`<-chan string`と読み出し専用のチャネルにすることで，main()がこのチャネルに値を書き込むことを型レベルで防ぐことができます。

### select文を用いたイベント制御

複数のチャネルに対する読み出しや書き込みを同時に制御するためにはselect文を用います。select文はfor文と組み合わせて使う場合が多くあります。

#### case

複数の操作をselect文のcaseに指定しておくと，いずれかのcaseの操作が実行されたときに，該当する処理が実行されます。どれか1つcaseが実行されない限りは，select文はブロックします。

```go
ch1 := make(chan string)
ch2 := make(chan string)
for {
    select {
    case c1 := <-ch1:
        // ch1からデータを読み出したときに実行される
    case c2 := <-ch2:
        // ch2からデータを読み出したときに実行される
    case ch2 <- "c":
        // ch2にデータを書き込んだときに実行される
    }
}
```

#### default

caseの最後にdefaultを記述すると，実行するcaseがなかった場合はdefaultが実行されます。defaultの実行が終わるとselect文の処理が終わるため，select文がブロックされなくなります。

```go
ch1 := make(chan string)
ch2 := make(chan string)
for {
    select {
    case c1 := <-ch1:
        // ch1からデータを読み出したときに実行される
    case c2 := <-ch2:
        // ch2からデータを読み出したときに実行される
    case ch2 <- "c":
        // ch2にデータを書き込んだときに実行される
    default:
        // caseが実行されなかった場合に実行される
    }
}
```

#### タイムアウト

for/select文とbreakを用いて実装する代表的な例はタイムアウト処理です。

```go
func main() {
    // 1秒後に値が読み出せるチャネル
    timeout := time.After(time.Second)
    urls := []string{
        "http://example.com",
        "http://example.net",
        "http://example.org",
    }
    statusChan := getStatus(urls)

LOOP: // ラベル名は任意
    for {
        select {
        case status := <-statusChan:
            fmt.Println(status) // 受信したデータを表示
        case <-timeout:
            break LOOP // このfor/selectを抜ける
        }
    }
}
```

timeパケージにあるtime.After()関数は，時間を指定するとその時間後にデータを書き込むチャネルを返す関数です。このチャネルの読み出しをselect文に登録することで，タイムアウト処理を実現できます。

先ほどのstatusChanの読み出しを無限forループ内のselect文で受け取るようにします。ステータスを受け取った場合はそれが表示されますが，1秒後にtimeoutから値が読み出せると，そこでfor/select文を抜けて，HTTPリクエストがすべて終わっている／いないにかかわらず，プログラムを終わらせることができます。

注意点として，caseの中でbreakを呼ぶと，select文は抜けられますが，その外側のfor文は抜けられません。そこでfor文にラベルを付け，breakでそのラベルを指定することで，caseからfor/select文を抜けることができます。returnで関数ごと抜けることもできますが，ラベルのbreakもよく使うパターンなので覚えておくとよいでしょう。

### チャネルバッファ

make()でチャネルを生成するときに，バッファを指定できます。バッファとは，そのチャネルが内部に保持できるデータの数です。

#### バッファなしチャネル

バッファを指定せずにmake()で生成したチャネルは，内部に値を保持しておくことができません。

次の場合はmain()内でチャネルの読み出す側に先に到達しますが，ゴルーチン内で値が書き込まれるまでそこで1秒間ブロックします。

```go
func main() {
    ch := make(chan string)
    go func() {
        time.Sleep(time.Second)
        ch <- "a" // 1秒後にデータを書き込む
    }()
    <-ch // 1秒後にデータが書き込まれるまでブロック
}
```

逆に次の場合は，main()内でチャネルに書き込む側に先に到達しますが，ゴルーチン内でその値が読み出されるまでそこで1秒間ブロックします。

```go
func main() {
    ch := make(chan string)
    go func() {
        time.Sleep(time.Second)
        <-ch // 1秒後にデータを読み出す
    }()
    ch <- "a" // 1秒後にデータが読み出されるまでブロック
}
```

このことを利用してバッファゼロのチャネルをゴルーチン間の同期制御に使うこともできますが，ブロックしないほうが都合の良い場合もあります。

#### バッファ付きチャネル

チャネルのバッファサイズはmake()の引数で指定します。たとえば次のようにバッファを3にして生成したチャネルは，同時に3つまでは値を内部に保持できます。そのため，3つまでの書き込みはブロックしません。しかし，4つ目の書き込みは，先にチャネルから値が読み出されないかぎりブロックします。

```go
func main() {
    ch := make(chan string, 3)
    go func() {
        time.Sleep(time.Second)
        <-ch // 1秒後にデータを読み出す
    }()
    ch <- "a" // ブロックしない
    ch <- "b" // ブロックしない
    ch <- "c" // ブロックしない
    ch <- "d" // 1秒後にデータが読み出されるまでブロック
}
```

このように，バッファ付きのチャネルはメッセージキューのような挙動になります。

たとえば先ほどのstatus取得の例では，いずれのゴルーチンもstatusChanに値を書き込むことで終了するのですが，もしstatusChanからデータを受け取るmain()側の処理が非常に遅かった場合，ゴルーチンはステータスの取得が終わっているのに，書き込みでブロックして閉じることができません。この場合は，statusChanにバッファを付けることで，main()側の処理が遅くても，ゴルーチンはチャネルに値を書き込んで終了することができ，メモリへの負荷を下げることができます。

```go
func getStatus(urls []string) <-chan string {
    // バッファをURLの数（3）に
    statusChan := make(chan string, len(urls))
    for _, url := range urls {
        go func(url string) {
            res, err := http.Get(url)
            if err != nil {
                log.Fatal(err)
            }
            defer res.Body.Close()
            // main()の読み出しが遅くても
            // 3つのゴルーチンはすぐに終わる
            statusChan <- res.Status
        }(url)
    }
    return statusChan
}
```

#### ゴルーチンの同時起動数制御

getStatus()にURLが複数渡ってきた場合に，ここまでの実装ではURLの数だけゴルーチンを起動していました。しかし，これではURLの数が多かった場合にゴルーチンが起動し過ぎてしまい，メモリを圧迫する可能性があります。

ここではlimitというバッファ付きのチャネルを用いて，このチャネルに値を書き込める場合はゴルーチンを起動し，ゴルーチンが終わったらlimitから値を読み出すことで，ゴルーチンの同時起動数を制御してみます。この場合，チャネルのバッファにあるデータの数が重要であり，データ自体には意味がないため，サイズがゼロの構造体を用います。

```go
var empty struct{} // サイズがゼロの構造体

func getStatus(urls []string) <-chan string {
    statusChan := make(chan string, 3)
    // バッファを5に指定して生成
    limit := make(chan struct{}, 5)
    go func() {
        for _, url := range urls {
            select {
            case limit <- empty:
                // limitに書き込みが可能な場合は取得処理を実施
                go func(url string) {
                    // このゴルーチンは同時に5つしか起動しない
                    res, err := http.Get(url)
                    if err != nil {
                        log.Fatal(err)
                    }
                    statusChan <- res.Status
                    // 終わったら1つ読み出して空きを作る
                    <-limit
                }(url)
            }
        }
    }()
    return statusChan
}
```