# yatte-mi

AIが「やりたいこと」と「期間」からタスクを分解し、自動でカレンダーに登録するWebアプリのバックエンド（Go + Gin）。

---

## 🚀 概要

`yatte-mi` は、以下の入力からあなたの「やりたいこと」を実現するためのタスクを自動でAIが分解・プランニングし、Googleカレンダーなどに自動反映するサービスです。  
本リポジトリはそのバックエンドAPI部分（Go製）のコードです。

---

## 🔧 開発環境構築手順

### ✅ 1. Goのインストール（Mac）

```bash
brew install go
```

※ 初回の場合は以下のライセンス同意も必要です：

```bash
sudo xcodebuild -license accept
```
Goのバージョン確認：

```bash
go version
```
 Goモジュールの初期化
```bash
go mod init github.com/exampleuser/yatte-mi
```

 Ginなど依存パッケージの導入
```bash
go get github.com/gin-gonic/gin
go mod tidy
```