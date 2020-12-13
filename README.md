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

#### 第 8 章 実装を隠す

##### 第 8 章の振り返り

> - 重複を除去できる状態に一歩近づけるために、Dollar と Franc にある 2 つの times メソッドのシグニチャを合わせた。
> - Factory Method パターンを導入して、テストコードから 2 つのサブクラスの存在を隠した。
> - サブクラスを隠した結果、いくつかのテストが冗長なものになったことに気がついたが、いまはそのままにしておいた。

- Go には抽象クラスの概念が無いため、Times メソッドについては一足先に実装もろとも Money に移行した
- それにより、Dollar, Franc の 2 つの構造体が使われなくなったが、一旦取っておくことにする

##### 第 8 章の TODO リスト

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
> - [x] times の一般化
> - [x] Franc と Dollar を比較する
> - [ ] 通貨の概念
> - [ ] testFrancMultiplication を削除する？

##### 第 8 章終了時のコード

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter8)

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

// NewDollar is constructor of Dollar.
func NewDollar(a int) Money {
	return Money{
		amount: a,
		name:   "Dollar",
	}
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) Money {
	return Money{
		amount: a,
		name:   "Franc",
	}
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (m Money) Times(multiplier int) Money {
	return Money{
		amount: m.amount * multiplier,
		name:   m.name,
	}
}

// Equals checks if the amount of the receiver and the argument are the same
func (m Money) Equals(a Accessor) bool {
	return m.amount == a.Amount() && m.name == a.Name()
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

```dollar.go
package money

// Dollar is a struct that handles dollar money.
type Dollar struct {
	Money
}
```

#### 第 9 章 歩幅の調整

##### 第 9 章の振り返り

> - 大きめの設計変更にのめり込みそうになったので、その前に手前にある小さな変更に着手した。
> - 差異を呼び出し側（FactoryMethod 側）に移動することによって、2 つのサブクラスのコンストラクタを近づけていった。
> - リファクタリングの途中で少し寄り道して、times メソッドの中で FactoryMethod を使うように変更した。
> - Franc に行ったリファクタリングを Dollar にも同様に、今度は大きい歩幅で一気に適用した。
> - 完全に同じ内容になった 2 つのコンストラクタを親クラスに引き上げた。

currency field は第 7 章で作成した name field を currency に改名しただけになる

##### 第 9 章の TODO リスト

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
> - [x] times の一般化
> - [x] Franc と Dollar を比較する
> - [x] **通貨の概念**
> - [ ] testFrancMultiplication を削除する？

##### 第 9 章終了時のコード

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter9)

```money.go
package money

// Accessor is a accessor of Money
type Accessor interface {
	Amount() int
	Currency() string
}

// Money is a struct that handles money.
type Money struct {
	amount   int
	currency string
}

// NewMoney is constructor of Money.
func NewMoney(a int, c string) Money {
	return Money{
		amount:   a,
		currency: c,
	}
}

// NewDollar is constructor of Dollar.
func NewDollar(a int) Money {
	return NewMoney(a, "USD")
}

// NewFranc is constructor of Dollar.
func NewFranc(a int) Money {
	return NewMoney(a, "CHF")
}

// Times multiplies the amount of the receiver by a multiple of the argument
func (m Money) Times(multiplier int) Money {
	return Money{
		amount:   m.amount * multiplier,
		currency: m.currency,
	}
}

// Equals checks if the amount of the receiver and the argument are the same
func (m Money) Equals(a Accessor) bool {
	return m.amount == a.Amount() && m.currency == a.Currency()
}

// Amount returns amount field
func (m Money) Amount() int {
	return m.amount
}

// Currency returns name field
func (m Money) Currency() string {
	return m.currency
}
```

```money_test.go
t.Run("通貨テスト", func(t *testing.T) {
	assert.Equal(t, "USD", money.NewDollar(1).Currency())
	assert.Equal(t, "CHF", money.NewFranc(1).Currency())
})
```

#### 第 10 章 テストに聞いてみる

##### 第 10 章の振り返り

times メソッドについては既に共通化しているため、ログ出力用の String メソッドのみ実装した。

##### 第 10 章終了時のコード

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter10)

```money.go
func (m Money) String() string {
	return fmt.Sprintf("{Amount: %v, Currency: %v}", m.amount, m.currency)
}
```

#### 第 11 章 不要になったら消す

##### 第 11 章の振り返り

> - サブクラスの仕事を減らし続け、とうとう消すところまでたどり着いた。
> - サブクラス削除前の構造では意味があるものの、削除後は冗長になってしまうテストたちを消した。

##### 第 11 章の TODO リスト

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
> - [x] **Dollar と Franc の重複**
> - [x] equals の一般化
> - [x] times の一般化
> - [x] Franc と Dollar を比較する
> - [x] 通貨の概念
> - [x] **testFrancMultiplication を削除する？**

##### 第 11 章終了時のコード

dollar.go, franc.go,ファイルを削除した。

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter11)

#### 第 12 章 設計とメタファー

##### 第 12 章の振り返り

> - 大きいテスト($5 + 10 CHF)を分解して。進み具合がわかる小さいテスト($5+$5)を作成した。
> - これから行う計算のためのメタファーについて深く考えた。
> - テストがコンパイルできるところまで早足で進んだ。
> - テストを通した。
> - 本当の実装を導くためのリファクタリングを楽しみにしつつ、少し不安も感じている。

##### 第 12 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [ ] $5 + $5 = $10

##### 第 12 章終了時のコード

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter12)

```money.go
// Plus adds an argument to the amount of receiver
func (m Money) Plus(added Money) Expression {
	return NewMoney(m.amount+added.amount, m.currency)
}
```

```bank.go
package money

// Bank calculates using exchange rates
type Bank struct {
}

// Reduce applies the exchange rate to the argument expression
func (b Bank) Reduce(source Expression, to string) Money {
	return NewDollar(10)
}
```

```expression.go
package money

// Expression shows the formula of currency (regardless of the difference in exchange rate)
type Expression interface {
}
```

```money_test.go
t.Run("ドル同士の足し算が可能である", func(t *testing.T) {
	five := money.NewDollar(5)
	sum := five.Plus(five)
	bank := money.Bank{}
	reduced := bank.Reduce(sum, "USD")
	assert.Equal(t, money.NewDollar(10), reduced)
})
```

#### 第 13 章 実装を導くテスト

##### 第 13 章の振り返り

> - 重複を除去出来ていないので、TODO リストの項目を「済」にしなかった。
> - 実装の着想を得るためにさらに先に進むことにした。
> - 速やかに実装を行った(Sum のコンストラクタ)
> - キャストを使って 1 カ所で実装した後で、テストが通る馬で本来あるべき場所にコードを移した。
> - ポリモフィズムを使って、明示的なクラスチェックを置き換えた。

##### 第 13 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [ ] $5 + $5 = $10
> - [ ] $5 + $5 が Money を返す
> - [x] Bank.reduce(Money)
> - [ ] Money を変換して換算を行う
> - [ ] Reduce(Bank, String)

##### 第 13 章終了時のコード

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter13)

```maoney_test.go
t.Run("ドル同士の足し算が可能である", func(t *testing.T) {
	five := money.NewDollar(5)
	result := five.Plus(five)
	sum := result.(money.Sum)
	assert.Equal(t, five, sum.Augend)
	assert.Equal(t, five, sum.Added)
})
t.Run("Sumで足されるお金の通貨が同じなら、足し算の結果が同じになる", func(t *testing.T) {
	sum := money.Sum{
		Augend: money.NewDollar(3),
		Added:  money.NewDollar(4),
	}
	bank := money.Bank{}
	result := bank.Reduce(sum, "USD")
	assert.Equal(t, money.NewDollar(7), result)
})
t.Run("moneyをreduceしても、reduceに渡す通貨が同じであれば同じ値が返る", func(t *testing.T) {
	bank := money.Bank{}
	result := bank.Reduce(money.NewDollar(1), "USD")
	assert.Equal(t, money.NewDollar(1), result)
})
```

```money.go
// Plus adds an argument to the amount of receiver.
func (m Money) Plus(added Money) Expression {
	return Sum{
		Augend: m,
		Added:  added,
	}
}

// Reduce applies the exchange rate to receiver.
func (m Money) Reduce(to string) Money {
	return m
}
```

```bank.go
// Reduce applies the exchange rate to the argument expression
func (b Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(to)
}
```

```expression.go
// Expression shows the formula of currency (regardless of the difference in exchange rate)
type Expression interface {
	Reduce(string) Money
```

```sum.go
package money

// Sum 合計は通貨の加算を行います
type Sum struct {
	Augend Money
	Added  Money
}

// Reduce applies the exchange rate to the result of the addition
func (s Sum) Reduce(to string) Money {
	amount := s.Augend.amount + s.Added.amount
	return NewMoney(amount, to)
}
```

#### 第 14 章 学習用テストと回帰テスト

##### 第 14 章の振り返り

> - 必要になると予想されたパラメータ追加をすぐに行った。
> - コードとテストの間のデータ重複をくくりだした。
> - 内部実装で使うためだけのヘルパークラスを個別のテスト無しで作成した。
> - リファクタリング中にミスを犯したが、問題を再現するテストを追加して、着実に前進した。

今回、本にあるような Pair.go ファイルを作成せず、Pair struct のみ Bank ファイルに記述した。

Java オブジェクトの`equals`による比較は等値比較だが、Go の struct の`==`による比較は等価比較になるため。

ただし、struct に map や slice といった等価比較できないフィールドが存在する場合は、コンパイル時にエラーになる。

そういった場合は`reflect.DeepEqual`を使って比較を行う

参考:[How to compare if two structs, slices or maps are equal?](https://stackoverflow.com/questions/24534072/how-to-compare-if-two-structs-slices-or-maps-are-equal)

```go
type s struct {
	a int
}
s1 := s{1}
s2 := s{1}
fmt.Println(s1 == s2) // true
```

```go

type s struct {
	a int
	b []int
}
s1 := s{1, make([]int, 0)}
s2 := s{1, make([]int, 0)}
fmt.Println(reflect.DeepEqual(s1, s2)) // true
fmt.Println(s1 == s2) // invalid operation: s1 == s2 (struct containing []int cannot be compared)
```

ちなみに、`assert.Equal`も内部で`reflect.DeepEqual`を使用している。

```go
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual) // ここ
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}
```

##### 第 14 章の TODO リスト

> - [ ] \$5+10CHF=$10（レートが 2:1 の場合）
> - [x] $5 + $5 = $10
> - [ ] $5 + $5 が Money を返す
> - [x] Bank.reduce(Money)
> - [x] Money を変換して換算を行う
> - [x] Reduce(Bank, String)

##### 第 14 章終了時のコード

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter14)

```money_test.go
t.Run("1 CHF = $2", func(t *testing.T) {
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(money.NewFranc(2), "USD")
	assert.Equal(t, money.NewDollar(1), result)
})
t.Run("同量テスト", func(t *testing.T) {
	bank := money.NewBank()
	assert.Equal(t, 1, bank.Rate("USD", "USD"))
})
```

```money.go
// Reduce applies the exchange rate to receiver.
func (m Money) Reduce(b Bank, to string) Money {
	rate := b.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
}
```

```bank.go
// Bank calculates using exchange rates
type Bank struct {
	rates map[Pair]int
}

// Pair associates two currencies
type Pair struct {
	from, to string
}

// NewBank is a constructor of Bank
func NewBank() Bank {
	b := Bank{}
	b.rates = make(map[Pair]int)
	return b
}

// Reduce applies the exchange rate to the argument expression
func (b *Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(*b, to)
}

// AddRate adds exchange rate
func (b *Bank) AddRate(from, to string, rate int) {
	b.rates[Pair{from: from, to: to}] = rate
}

// Rate adds exchange rate
func (b *Bank) Rate(from, to string) int {
	if from == to {
		return 1
	}
	p := Pair{from: from, to: to}
	return b.rates[p]
}
```

```expression.go
// Expression shows the formula of currency (regardless of the difference in exchange rate)
type Expression interface {
	Reduce(Bank, string) Money
}
```

```sum.go
// Reduce applies the exchange rate to the result of the addition
func (s Sum) Reduce(b Bank, to string) Money {
	amount := s.Augend.amount + s.Added.amount
	return NewMoney(amount, to)
}
```

#### 第 15 章 テスト任せとコンパイラ任せ

##### 第 15 章の振り返り

> - こうなったら良いというテストを書き、次にまず一歩で動かせるところまでそのテストを少し後退させた。
> - 一般化(より抽象度の高い型で宣言する)作業を、末端から開始して頂点(テストケース)まで到達させた。
> - 変更の際にコンパイラに従い(fiveBucks 変数の Expression 型への変更)、変更の連鎖を 1 つずつ仕留めた(Expression インターフェースへの Plus メソッドの追加等)。

どうでもいいが、Dollar の代わりに Bucks が使われることを初めて知った。

##### 第 15 章の TODO リスト

> - [x] \$5+10CHF=$10（レートが 2:1 の場合）
> - [x] $5 + $5 = $10
> - [ ] $5 + $5 が Money を返す
> - [x] Bank.reduce(Money)
> - [x] Money を変換して換算を行う
> - [x] Reduce(Bank, String)
> - [ ] Sum.Plus
> - [ ] Expression.Times

##### 第 15 章終了時のコード

全文: [github](https://github.com/eyuta/golang-tdd/tree/part1_chapter15)

```money_test.go
t.Run("$5 + 10 CHF = $10 (レートが2:1の場合)", func(t *testing.T) {
	fiveBucks := money.Expression(money.NewDollar(5))
	tenFrancs := money.Expression(money.NewFranc(10))
	bank := money.NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveBucks.Plus(tenFrancs), "USD")
	assert.Equal(t, money.NewDollar(10), result)
})
```

```money.go
// Times multiplies the amount of the receiver by a multiple of the argument
func (m Money) Times(multiplier int) Expression {
	return NewMoney(m.amount*multiplier, m.currency)
}

// Plus adds an argument to the amount of receiver.
func (m Money) Plus(added Expression) Expression {
	return Sum{
		Augend: m,
		Added:  added,
	}
}
```

```expression.go
// Expression shows the formula of currency (regardless of the difference in exchange rate)
type Expression interface {
	Reduce(Bank, string) Money
	Plus(Expression) Expression
}
```

```sum.go
// Reduce applies the exchange rate to the result of the addition
func (s Sum) Reduce(b Bank, to string) Money {
	amount := s.Augend.Reduce(b, to).amount + s.Added.Reduce(b, to).amount
	return NewMoney(amount, to)
}

// Plus adds an argument to the amount of receiver.
func (s Sum) Plus(added Expression) Expression {
	return Sum{}
}
```

### 第 II 部「xUnit」

### 第 III 部「テスト駆動開発のパターン」

## 参考文献

- [ＫｅｎｔＢｅｃｋ. テスト駆動開発 Kindle 版](https://www.amazon.co.jp/dp/B077D2L69C/ref=dp-kindle-redirect?_encoding=UTF8&btkr=1)
