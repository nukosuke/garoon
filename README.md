[![Build Status](https://travis-ci.org/nukosuke/garoon.svg?branch=master)](https://travis-ci.org/nukosuke/garoon)
[![Build status](https://ci.appveyor.com/api/projects/status/xjj9su4a1wde9bac/branch/master?svg=true)](https://ci.appveyor.com/project/nukosuke/garoon/branch/master)

## 環境

クラウド版Cybozu Garoon

## インストール

``` shell
$ go install github.com/nukosuke/garoon
```

## 設定

認証情報はコマンドラインオプションで指定することもできますが、設定ファイルに記述しておくと
いちいち入力しなくてすむので便利です。デフォルトでは `$HOME/.config/garoon/config.yml` を参照
します。設定ファイルのパスは `--config` オプションで指定することもできます。

``` shell
$ cat ~/.config/garoon/config.yml
subdomain: <Garoonサブドメイン>
username: <Garoonユーザ名>
password: <Garoonパスワード>
```

## 使い方

基本的な使用方法はヘルプで確認できます。

```
$ garoon -h

garoon v0.0.0 --- Cybozu Garoonのコマンドラインツール

        バグレポートはこちらにお願いします。
        https://github.com/nukosuke/garoon

Usage:
  garoon [command]

Available Commands:
  event       予定の取得コマンド
  help        Help about any command

Flags:
      --config string      設定ファイルのパス (デフォルトは $HOME/.config/garoon/config.yml)
  -h, --help               help for garoon
  -p, --password string    ログインユーザのパスワード (設定ファイルがある場合は無視されます)
  -d, --subdomain string   Garoonのサブドメイン (設定ファイルがある場合は無視されます)
  -u, --username string    ログインユーザ名 (設定ファイルがある場合は無視されます)

Use "garoon [command] --help" for more information about a command.
```

応用例として[peco](https://github.com/peco/peco)と一緒に使用すると便利です。 `.bashrc` などに下記のようなエイリアスを張ります。

``` shell
alias ge='garoon event info $(garoon event ls | peco | cut -f1)'
```

## Copyright

Copyright (C) 2019 nukosuke.

This software is distributed under MIT License.  
See the file [LICENSE](./LICENSE) for details.
