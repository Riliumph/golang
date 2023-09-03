# Receiverとは

Go言語にはクラスという概念が存在しないため、それを実現するための機能。  
Pythonのように第一引数にレシーバが入り込む仕組み。

レシーバは別変数であるという意識が必要。  
変数レシーバもポインタレシーバも、別変数だが以下の違いがある。

- 変数レシーバは、値を実体の値渡しでDeepCopyする
- ポインタレシーバは、値を参照渡しでShallowCopyする

## 変数レシーバ

```go
package main

import (
        "fmt"
        "unsafe"
)

type Data struct{
    i int
    s string
}

func New(i int, s string) *Data {
    return &Data{
        i: i,
        s: s,
    }
}

func (d Data) Detail(){
    d.i = 100
    fmt.Printf("instance(%p)\n", &d)
    fmt.Printf("i(%p - %d): %d\n", &d.i, unsafe.Sizeof(d.i), d.i)
    fmt.Printf("s(%p - %d): %s\n", &d.s, unsafe.Sizeof(d.s), d.s)
}

func main() {
    d1 := New(1, "hoge")
    d1.Detail()
    fmt.Printf("d1(%p)\ni(%p): %d", &d1, &(d1).i, d1.i)
}
```

```console
$ go run .
instance(0xc00000c048)
i(0xc00000c048 - 8): 100
s(0xc00000c050 - 16): hoge
d1(0xc00000e028)
i(0xc00000c030): 1
```

呼び出し元は、`0xc00000e028`。
変数レシーバは、`0xc00000c048`。

レシーバは別物の変数であることが分かる。

また、Detailメソッドでは100が出力されているが、呼び出し元は1のままなので、値を変えても呼び出し元に影響しないことも分かる。

## ポインタレシーバ

```go
package main

import (
        "fmt"
        "unsafe"
)

type Data struct{
    i int
    s string
}

func New(i int, s string) *Data {
    return &Data{
        i: i,
        s: s,
    }
}

func (d *Data) Detail(){
    d.i = 100
    fmt.Printf("instance(%p)\n", &d)
    fmt.Printf("i(%p - %d): %d\n", &d.i, unsafe.Sizeof(d.i), d.i)
    fmt.Printf("s(%p - %d): %s\n", &d.s, unsafe.Sizeof(d.s), d.s)
}

func main() {
    d1 := New(1, "hoge")
    d1.Detail()
    fmt.Printf("d1(%p)\ni(%p): %d", &d1, &(d1).i, d1.i)
}
```

```console
$ go run .
instance(0xc00000e030)
i(0xc00000c030 - 8): 100
s(0xc00000c038 - 16): hoge
d1(0xc00000e028)
i(0xc00000c030): 100
```

呼び出し元は、`0xc00000e028`。
ポインタレシーバは、`0xc00000e030`。

ポインタレシーバも変数レシーバ同様に別のメモリが割り当てられていることがわかる。

しかし、Detailメソッドで書き換えた値が呼び出し元にも反映されていることから、ポインタレシーバはアドレスを受け取っているので呼び出し元にも影響を与えられる存在であることが分かる。
