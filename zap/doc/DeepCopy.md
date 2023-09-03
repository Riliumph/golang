# DeepCopy

Go言語に於けるDeepCopyの実装・調査。

## zapの実装

zapには`clone`メソッドの実装を見てみる。

```go
func (log *Logger) clone() *Logger {
 copy := *log
 return &copy
}
```

実体をコピーしてそのアドレスを返す実装になっている。
Go言語は、参照がなくなった時にメモリスイープされるのでC言語にはできない書き方で実現している。

## メンバはコピーされる？

オブジェクトAの実体をコピーしてDeepCopyしたのは分かった。  
だが、メンバ変数はDeepCopyされたのかを確認してみる。

`zap.Logger`はPrivate変数なのでオリジナルでやってみる。

```go
package main

import (
        "fmt"
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


func main() {
    d1 := New(1, "hoge")
    fmt.Printf("d1(%p): i(%p): %d, s(%p): %s\n", &d1, &d1.i, d1.i, &d1.s, d1.s)
    d2 := d1
    fmt.Printf("d2(%p): i(%p): %d, s(%p): %s\n", &d2, &d2.i, d2.i, &d2.s, d2.s)
}
```

```console
$ go run .
d1(0xc00000e028): i(0xc00000c030): 1, s(0xc00000c038): hoge
d2(0xc00000e038): i(0xc00000c030): 1, s(0xc00000c038): hoge
```

確かに、Data型変数`d1`と`d2`はアドレスが変わったので、別の変数になった。  
しかし、int型メンバ変数`i`もstring型メンバ変数`s`もアドレスが変わっていないので参照のままである。

しっかり全メンバをクローンするならC++のように自前実装が必要になることが分かる。

## メモリマップ
