# go-rockmanbot

## 概要

C3のmastodonで稼働しているbot「ロックマンBot」をGoで改修したリポジトリ。

## フォルダ構成

- main.go
  - 実行するmainファイル
- apps
  - ドメイン層に関連するファイル
- interfaces
  - ドメイン層で利用するインターフェースファイル
- reactions
  - TLの投稿に反応して投稿する処理のファイル
- tests
  - interfacesを継承したテスト用の構造体を利用したテスト処理ファイル
- utils
  - 具体的な処理を行う関数をまとめたファイル
