# go-utils
Goで自作した便利な関数を定義するモジュールです。

## ディレクトリ、ファイルにおける命名規則
- ディレクトリ
  - ディレクトリ名をパッケージ名と合わせる
  - 基本一単語にする
  - もし２単語以上にしたいなら、ケバブケース(xxx-xxx)にする
- ファイル
  - 基本ファイル名はパッケージ名と合わせる
  - 基本一単語にする
  - もし2単語以上にしたいなら、スネークケース(xxx_xxxx)を使う

## 参考記事
- [pkgのディレクトリ構造で参考にしたリポジトリ](https://github.com/solo-io/squash/tree/master/pkg)
- [cmdのディレクトリ構造で参考にしたリポジトリ](https://github.com/ethereum/go-ethereum/tree/master/cmd)
- [Goプロジェクトのディレクトリ構造で参考にした記事](https://qqq.ninja/blog/post/go-structure/)
- [Goプロジェクトのディレクトリ構造で参考にしたリポジトリ](https://github.com/golang-standards/project-layout)
- [Goの命名で参考にした記事その1](https://zenn.dev/keitakn/articles/go-naming-rules)
- [Goの命名で参考にした記事その2](https://onuma-ryota.com/posts/golang-best-practice/)
- [パッケージインポートについてのスクラップ](https://zenn.dev/yukihaga/scraps/c2fbaac12235aa)