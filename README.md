# go-fundamentals-lab

## セットアップ

[mise](https://mise.jdx.dev/) でGo・golangci-lint・lefthookのバージョンを管理しています。

```sh
mise install       # ツール一式をインストール
mise run hooks:install  # git hooks (lefthook) を有効化
```

シークレット検知には[git-secrets](https://github.com/awslabs/git-secrets)を使用しています。検知パターンはリポジトリの`.git/config`に保存されるため、クローンごとに以下を実行してください（未実行の場合、`pre-commit`のsecretsチェックは何も検知しません）。

```sh
brew install git-secrets   # 未インストールの場合
git secrets --register-aws # AWSキー等の標準パターンを登録
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

- `pre-commit`: シークレット検知(git-secrets) → fmt/lint(変更されたGoファイルのみ)
- `pre-push`: go test
