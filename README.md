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

`go version go1.15.6 windows/amd64`

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

### 第 II 部「xUnit」

### 第 III 部「テスト駆動開発のパターン」

## 参考文献

- [ＫｅｎｔＢｅｃｋ.テスト駆動開発(Kindle の位置 No.200-201).Kindle 版](https://www.amazon.co.jp/dp/B077D2L69C/ref=dp-kindle-redirect?_encoding=UTF8&btkr=1)
