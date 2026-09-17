# go-fundamentals-lab

## セットアップ

[mise](https://mise.jdx.dev/) でGo・golangci-lint・lefthookのバージョンを管理しています。

```sh
mise install       # ツール一式をインストール
mise run hooks:install  # git hooks (lefthook) を有効化
```

## タスク

```sh
mise run fmt    # gofmt
mise run vet    # go vet
mise run lint   # golangci-lint
mise run test   # go test
mise run build  # go build
mise run tidy   # go mod tidy
```

`mise tasks` で一覧を確認できます。

## git hooks

lefthookで以下を自動実行します。

- `pre-commit`: シークレット検知(git-secrets) → fmt/vet/lint(変更されたGoファイルのみ)
- `pre-push`: go test
