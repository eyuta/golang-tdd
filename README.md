# テスト駆動開発を Go 言語で取り組んでみる

#

## はじめに

先日、t_wada さんが弊社に公演に来てくださいました。
それに先駆け、以前購入した t_wada さんが訳された[テスト駆動開発](https://www.amazon.co.jp/dp/B077D2L69C/ref=dp-kindle-redirect?_encoding=UTF8&btkr=1)を、現在学習中の Go 言語で取り組んでみました。
本記事中の引用は、特に断りがない限りこの本の引用になります。

## リポジトリ

<https://github.com/eyuta/golang-tdd>

## 目的

## 前提

### バージョン

`go version go1.15.6 windows/amd64`

### テスト方法について

今回は、Go の標準の [testing](https://golang.org/pkg/testing/) パッケージと、こちらサードパーティの[assert](https://godoc.org/github.com/stretchr/testify/assert)パッケージを使用しています。

Go の標準の testing パッケージには、Assert が含まれておらず、推奨もされていません。
理由については、以下の記事が詳しいです。
ただ、今回は testing としてのテストではなく、checking としてのテストがメインであることから、手軽にテストケースを記述できる Assert パッケージを使用しています。

参考記事

- [Go の Test に対する考え方](https://qiita.com/Jxck_/items/8717a5982547cfa54ebc#assert-%E3%81%8C%E7%84%A1%E3%81%84%E7%90%86%E7%94%B1)
- [[公式]Why does Go not have assertions?](https://golang.org/doc/faq#assertions)

testing パッケージの使い方は以下を参照しました。

- [Go でテストを書く(テストの実装パターン集)](https://qiita.com/atotto/items/f6b8c773264a3183a53c)

## 本編

### TDD について

TDD のルール

> コードを書く前に、失敗する自動テストコードを必ず書く。
> 重複を除去する。

TDD のリズム

> 1. まずはテストを 1 つ書く
> 2. すべてのテストを走らせ、新しいテストの失敗を確認する
> 3. 小さな変更を行う
> 4. すべてのテストを走らせ、すべて成功することを確認する
> 5. リファクタリングを行って重複を除去する

### 第 I 部 多国通貨

#### 第 1 章 仮実装

##### 第 1 章の振り返り

> - 書くべきテストのリストを作った。
> - どうなったら嬉しいかを小さいテストコードで表現した。
> - 空実装を使ってコンパイラを通した。
> - 大罪を犯しながらテストを通した。
> - 動くコードをだんだんと共通化し、ベタ書きの値を変数に置き換えていった。
> - TODO リストに項目を追加するに留め、一度に多くのものを相手にすることを避けた。

##### 第 1 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [x] **\$5\*2=$10**
> - [x] **amount を private にする**
> - [ ] Dollar の副作用どうする？
> - [ ] Money の丸め処理どうする？

##### 第 1 章終了時のコード

```multiCurrencyMoney.go
type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) {
	d.amount *= multiplier
}
```

```multiCurrencyMoney_test.go
func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("$5 * 2 = $10", func(t *testing.T) {
		five := Dollar{5}
		five.times(2)
		assert.Equal(t, 10, five.amount)
	})
}
```

#### 第 2 章 明確な実装

##### 第 2 章の振り返り

> - 設計の問題点（今回は副作用）をテストコードに写し取り、その問題点のせいでテストが失敗するのを確認した。
> - 空実装でさっさとコンパイルを通した。
> - 正しいと思える実装をすぐに行い、テストを通した。

##### 第 2 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [x] \$5\*2=$10
> - [x] amount を private にする
> - [x] **Dollar の副作用どうする？**
> - [ ] Money の丸め処理どうする？

##### 第 2 章終了時のコード

```multiCurrencyMoney.go
type Dollar struct {
	amount int
}

func (d *Dollar) times(multiplier int) Dollar {
	return Dollar{d.amount * multiplier}
}
```

```multiCurrencyMoney_test.go
func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("何度でもドルの掛け算が可能である", func(t *testing.T) {
		five := Dollar{5}
		product := five.times(2)
		assert.Equal(t, 10, product.amount)
		product = five.times(3)
		assert.Equal(t, 15, product.amount)
	})
}
```

#### 第 3 章 三角測量

型のコンバージョンは、[Type Assertion](https://go-tour-jp.appspot.com/methods/15)を利用した。

##### 第 3 章の振り返り

> - Value Object パターンを満たす条件がわかった。
> - その条件を満たすテストを書いた。
> - シンプルな実装を行った。
> - すぐにリファクタリングを行うのではなく、もう 1 つテストを書いた。
> - 2 つのテストを同時に通すリファクタリングを行った。

##### 第 3 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [x] \$5\*2=$10
> - [x] amount を private にする
> - [x] Dollar の副作用どうする？
> - [ ] Money の丸め処理どうする？
> - [x] **equals()**
> - [ ] hashCode()
> - [ ] null との等価性比較
> - [ ] 他のオブジェクトとの等価性比較

##### 第 3 章終了時のコード

```multiCurrencyMoney.go
type Object interface{}

type Dollar struct {
	amount int
}

func (d Dollar) times(multiplier int) Dollar {
	return Dollar{d.amount * multiplier}
}

func (d Dollar) equals(object Object) bool {
	dollar := object.(Dollar)
	return d.amount == dollar.amount
}
```

```multiCurrencyMoney_test.go
func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("何度でもドルの掛け算が可能である", func(t *testing.T) {
		five := Dollar{5}
		product := five.times(2)
		assert.Equal(t, 10, product.amount)
		product = five.times(3)
		assert.Equal(t, 15, product.amount)
	})

	t.Run("同じ金額が等価である", func(t *testing.T) {
		assert.True(t, Dollar{5}.equals(Dollar{5}))
		assert.False(t, Dollar{5}.equals(Dollar{6}))
	})
}
```

#### 第 4 章 意図を語るテスト

##### 第 4 章の振り返り

> - 作成したばかりの機能を使って、テストを改善した。
> - そもそも正しく検証できていないテストが 2 つあったら、もはやお手上げだと気づいた。
> - そのようなリスクを受け入れて先に進んだ。
> - テスト対象オブジェクトの新しい機能を使い、テストコードとプロダクトコードの間の結合度を下げた。

##### 第 4 章終了時のコード

```multiCurrencyMoney.go
// 変化なし
```

```multiCurrencyMoney_test.go
func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("ドルの掛け算が可能である", func(t *testing.T) {
		five := Dollar{5}
		assert.Equal(t, Dollar{10}, five.times(2))
		assert.Equal(t, Dollar{15}, five.times(3))
	})

	t.Run("同じ金額が等価である", func(t *testing.T) {
		assert.True(t, Dollar{5}.equals(Dollar{5}))
		assert.False(t, Dollar{5}.equals(Dollar{6}))
	})
}
```

#### 第 5 章 原則をあえて破るとき

##### 第 5 章の振り返り

> - 大きいテストに立ち向かうにはまだ早かったので、次の一歩を進めるために小さなテストをひねり出した。
> - 恥知らずにも既存のテストをコピー&ペーストして、テストを作成した。
> - さらに恥知らずにも、既存のモデルコードを丸ごとコピー&ペーストして、テストを通した。
> - この重複を排除するまでは家に帰らないと心に決めた。

##### 第 5 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [x] \$5\*2=$10
> - [x] amount を private にする
> - [x] Dollar の副作用どうする？
> - [ ] Money の丸め処理どうする？
> - [x] equals()
> - [ ] hashCode()
> - [ ] null との等価性比較
> - [ ] 他のオブジェクトとの等価性比較
> - [x] **5CHF\*2=10CHF**
> - [ ] Dollar と Franc の重複
> - [ ] equals の一般化
> - [ ] times の一般化

##### 第 5 章終了時のコード

```multiCurrencyMoney.go

type Franc struct {
	amount int
}

func (f Franc) times(multiplier int) Franc {
	return Franc{f.amount * multiplier}
}

func (f Franc) equals(object Object) bool {
	franc := object.(Franc)
	return f.amount == franc.amount
}
```

```multiCurrencyMoney_test.go
	t.Run("フランの掛け算が可能である", func(t *testing.T) {
		five := Franc{5}
		assert.Equal(t, Franc{10}, five.times(2))
		assert.Equal(t, Franc{15}, five.times(3))
	})
	t.Run("同じ金額のフランが等価である", func(t *testing.T) {
		assert.True(t, Franc{5}.equals(Franc{5}))
		assert.False(t, Franc{5}.equals(Franc{6}))
	})
```

#### 第 6 章 テスト不足に気づいたら

- Go には継承の概念が無いため、本章では [composition](https://www.geeksforgeeks.org/inheritance-in-golang/) を用いて実装する。
- Dollar, Franc を生成するためのコンストラクタにあたるものを用意した(以下を参考にした)。
  [Constructors and composite literals](https://golang.org/doc/effective_go.html#composite_literals)
- `multiCurrencyMoney.go`を`money.go`, `dollar.go`, `franc.go`に分割した
- `multiCurrencyMoney_test.go` を`money_test.go`に改名し、package 名を`money_test`とした
- `money.go`と`money_test.go`の package 名が異なるため、プライベートメソッド(小文字のメソッド)が参照できなくなったので、`equals`, `times`メソッドをパブリックメソッドに変更した
- パブリックメソッドにはコメントが必要になるので、簡単なコメントを追加した
  参考: [Godoc: documenting Go code](https://blog.golang.org/godoc)

##### 第 6 章の振り返り

このあたりから、言語仕様の違いによりコーディング内容が本と異なってくる

> - Dollar クラスから親クラス Money へ段階的にメソッドを移動した。
> - 2 つ目のクラス（Franc）も同様にサブクラス化した。
> - 2 つの equals メソッドの差異をなくしてから、サブクラス側の実装を削除した。

##### 第 6 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [x] \$5\*2=$10
> - [x] amount を private にする
> - [x] Dollar の副作用どうする？
> - [ ] Money の丸め処理どうする？
> - [x] equals()
> - [ ] hashCode()
> - [ ] null との等価性比較
> - [ ] 他のオブジェクトとの等価性比較
> - [x] 5CHF\*2=10CHF
> - [ ] Dollar と Franc の重複
> - [x] equals の一般化
> - [ ] times の一般化
> - [ ] Franc と Dollar を比較する

##### 第 6 章終了時のコード

量が増えてきたので、主な変更点のみ抜粋して表示する
全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter6)

```money.go
package money

// AmountGetter is a wrapper of amount.
type AmountGetter interface {
	getAmount() int
}

// Money is a struct that handles money.
type Money struct {
	amount int
}

// Equals checks if the amount of the receiver and the argument are the same
func (m Money) Equals(a AmountGetter) bool {
	return m.getAmount() == a.getAmount()
}

func (m Money) getAmount() int {
	return m.amount
}
```

```money_test.go
package money_test

import (
	"testing"

	"github.com/eyuta/golang-tdd/money"
	"github.com/stretchr/testify/assert"
)

func TestMultiCurrencyMoney(t *testing.T) {
	t.Run("ドルの掛け算が可能である", func(t *testing.T) {
		five := money.NewDollar(5)
		assert.Equal(t, money.NewDollar(10), five.Times(2))
		assert.Equal(t, money.NewDollar(15), five.Times(3))
	})
}
```

#### 第 7 章 疑念をテストに翻訳する

- 今回の修正とは関係ないが、`Go: Test On Save`Setting をチェックすることで保存時にテストが走るようになった。
  ![image.png](https://qiita-image-store.s3.ap-northeast-1.amazonaws.com/0/110860/c2ce12bb-9c25-542b-62dc-ba9197949a46.png)

##### 第 7 章の振り返り

> - 頭の中にある悩みをテストとして表現した。完璧ではないものの、まずまずのやり方（getClass）でテストを通した。
> - さらなる設計は、本当に必要になるときまで先延ばしにすることにした

- Money に新しく Name フィールドを追加した
  - 上記の`getClass`の代替。struct の入れ子の場合、レシーバは常に Money になるので、type の比較ができないため

##### 第 7 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [x] \$5\*2=$10
> - [x] amount を private にする
> - [x] Dollar の副作用どうする？
> - [ ] Money の丸め処理どうする？
> - [x] equals()
> - [ ] hashCode()
> - [ ] null との等価性比較
> - [ ] 他のオブジェクトとの等価性比較
> - [x] 5CHF\*2=10CHF
> - [ ] Dollar と Franc の重複
> - [x] equals の一般化
> - [ ] times の一般化
> - [x] Franc と Dollar を比較する
> - [ ] 通貨の概念

##### 第 7 章終了時のコード

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter7)

```money.go
package money

// Accessor is a accessor of Money
type Accessor interface {
	Amount() int
	Name() string
}

// Money is a struct that handles money.
type Money struct {
	amount int
	name   string
}

// Equals checks if the amount of the receiver and the argument are the same
func (m Money) Equals(a Accessor) bool {
	return m.Amount() == a.Amount() && m.Name() == a.Name()
}

// Amount returns amount field
func (m Money) Amount() int {
	return m.amount
}

// Name returns name field
func (m Money) Name() string {
	return m.name
}
```

```money_test.go
t.Run("同じ金額のドルとフランが等価ではない", func(t *testing.T) {
	assert.False(t, money.NewFranc(5).Equals(money.NewDollar(5)))
})
```

```dollar.go
func NewDollar(a int) Dollar {
	return Dollar{Money{amount: a, name: "Dollar"}}
}
```

### 第 II 部「xUnit」

### 第 III 部「テスト駆動開発のパターン」

## 参考文献

- [ＫｅｎｔＢｅｃｋ.テスト駆動開発(Kindle の位置 No.200-201).Kindle 版](https://www.amazon.co.jp/dp/B077D2L69C/ref=dp-kindle-redirect?_encoding=UTF8&btkr=1)
