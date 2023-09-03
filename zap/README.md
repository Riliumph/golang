# [zap](https://github.com/uber-go/zap)のサンプル

## goの実行方法

### modを使わない場合

```console
$ go run (すべての.goファイル)
```

C言語同様、すべてのファイルを指定する必要がある。  
非現実的なのでC言語はMakefileを用いるわけだが、Go言語には言語仕様の中にMakefileに相当する機能が入っている。

### modを使う場合

`go.mod`ファイルが実行ファイルとモジュールファイルを認識している。  
`go.mod`ファイルの直下で以下のコマンドを実行する。

```console
$ go run .
```

## テストの実行

```console
$ go test -v ./...
... ... ...
... ...
...
=== RUN   Test_GetLevel
--- PASS: Test_GetLevel (0.00s)
PASS
ok      how_to_zap/pkg/logger   0.002s
```

## モジュールのインストール

```console
$ go get go.uber.org/zap
$ go get github.com/stretchr/testify
```
