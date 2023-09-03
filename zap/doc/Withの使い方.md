# Withの使い方

`zap`のLoggerを使う方法はドキュメントを見ろ。  
今回の説明は、`zap.Logger`をラップしたオリジナルなLoggerでどうするかを考える。

## Withメソッドの実装

```go
func (log *Logger) With(fields ...Field) *Logger {
 if len(fields) == 0 {
  return log
 }
 l := log.clone()
 l.core = l.core.With(fields)
 return l
}
```

Go言語はレシーバを`*Logger`とすることで初めてコール元の変数を書き換えることができる。  
`clone`メソッドを使って新しいインスタンスのアドレスを返し、DeepCopyをしている

> DeepCopyの対象はあくまで、そのインスタンスのみである。メンバのポインタ変数まではDeepCopyされない。

## 永久に項目を追加する

```go
// zap only
zapLogger := zapLogger.With(zap.Any("request_id", "foobar"))
// original logger
orgLogger.logger = orgLogger.logger.With(zap.Any("request_id", "foobar"))
```

`With`メソッドはインスタンスの設定を変えるのではなく、設定が変わったインスタンスを返す作りである。  

> 「元に戻す」という処理を作るよりも、元のインスタンスを使った方がいいなどといった理由で作り直しているのかもしれない。

## 一時的にネストした項目を追加する

`Namespace`メソッドを使う。ただし、このメソッドは実行後は、すべての動作がこのセクションを起点に実行される。つまり、`NewSection`を作ると、以降のログはすべて`NewSection`の下にネストされて書き足される。

そうではなく、特定の内容だけネストしたい場合がある。  
あくまで一時的な処置として`NewSection`を作りたいのである。

その場合、Loggerを更新してしまうと以降の処理が汚染されてしまうので、Clone処理を行って別のインスタンスを用いることで、元のインスタンスを守る方法がある。

```go
// zap
zapLogger_tmp := zapLogger.With(zap.Namespace("NewSection"))
// original logger
newLogger := orgLogger.Clone()
newLogger.logger = orgLogger.logger.With(zap.Namespace("NewSection"))
```
