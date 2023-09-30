# gh-sosuke

GitHub CLI extension for sosukesuzuki's OSS development.

## Install

```
gh extension install sosukesuzuki/gh-sosuke
```

## Requirements

- Peco

## Usage

### `gh sosuke issue list`

```sh
gh sosuke issue list
```

`gh issue list` の結果を peco でインクリメンタルサーチし、その結果 issue 番号を標準出力に流します。

### `gh sosuke issue view`, ...etc

```sh
gh sosuke issue view
```

issue 番号を標準入力から読み取り、`gh issue view` に渡して実行します。`view` 以外にも https://cli.github.com/manual/gh_issue で説明されている Targeted commands をすべてサポートしています。
