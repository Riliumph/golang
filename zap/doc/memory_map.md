# メモリマップ

Go言語は面白いメモリ管理を行っている。

## 例１

```go
package main

import (
 "fmt"
)

type Data struct {
 i int
 s string
}

func main() {
 d := Data{
  i: 42,
  s: "Hello Go",
 }

 fmt.Printf("dのアドレス: %p\n", &d)
 fmt.Printf("iのアドレス: %p\n", &d.i)
 fmt.Printf("sのアドレス: %p\n", &d.s)
}
```

```console
$ go run .
dのアドレス: 0xc00000c030
iのアドレス: 0xc00000c030
sのアドレス: 0xc00000c038
```

構造体変数`d`と最初のメンバ変数`i`のアドレスが同じである。

C++などでは、変数のヘッダには仮想関数テーブルなどの構造体変数特有にヘッダ情報が記載されている。  
しかし、Go言語にはクラスの概念はない。  
結果として、ヘッダ情報などというメモリ管理はなく、構造体の先頭アドレスと最初のメンバ変数のアドレスは一致する。
