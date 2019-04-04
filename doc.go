/*
Cybozu Garoon CLIはコマンドラインからガルーンの操作を行うためのツールです。

Requirements

- クラウド版Cybozu Garoon

Installation

下記コマンドでインストールします。

  $ go install github.com/nukosuke/garoon

Configuration

認証情報はコマンドラインオプションで指定することもできますが、
設定ファイルを記述しておくといちいち入力しなくてすむので便利です。
デフォルトでは $HOME/.config/garoon/config.yml を参照します。
設定ファイルのパスは --config オプションで指定することもできます。

  $ cat ~/.config/garoon/config.yml
  subdomain: <Garoonサブドメイン>
  username: <Garoonユーザ名>
  password: <Garoonパスワード>

Usage

基本的な使用方法はヘルプで確認できます。

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

応用例としてpecoと一緒に使用すると便利です。 .bashrc などに下記のようなエイリアスを張ります。

  alias ge='garoon event info $(garoon event ls | peco | cut -f1)'

Copyright

Copyright (C) 2019 nukosuke.
本ソフトウェアはMITライセンスで配布されています。
詳しくはこちらをご覧ください。(https://github.com/nukosuke/garoon/blob/master/LICENSE)

*/
package main
